package migration

import (
	//"github.com/jmoiron/sqlx"
	"database/sql"
	"fmt"
	"embed"
	"errors"

	// Импорт драйвера PostgreSQL для его регистрации.
	// _ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4/database/postgres"

	// Импорт драйвера File для его регистрации.
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

const migrationsDir = "migrations"

//go:embed migrations/*.sql
var MigrationsFS embed.FS

func Migration(connectString string) error {
	migrator := MustGetNewMigrator(MigrationsFS, migrationsDir)

	// Get the DB instance
	conn, err := sql.Open("postgres", connectString)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()

	err = migrator.ApplyMigrations(conn)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//fmt.Printf("Migrations applied!!")

	return err
}

type Migrator struct {
	srcDriver source.Driver // Драйвер источника миграций.
}
   
func MustGetNewMigrator(sqlFiles embed.FS, dirName string) *Migrator {
	// Создаем новый драйвер источника миграций с встроенными SQL-файлами.
	d, err := iofs.New(sqlFiles, dirName)
	if err != nil {
	 	panic(err)
	}

	return &Migrator{
	 	srcDriver: d,
	}
}

func (m *Migrator) ApplyMigrations(db *sql.DB) error {
	// Создаем экземпляр драйвера базы данных для PostgreSQL.
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("unable to create db instance: %v", err)
	}

	// Создаем новый экземпляр мигратора с использованием драйвера источника и драйвера базы данных PostgreSQL.
	migrator, err := migrate.NewWithInstance("migration_embeded_sql_files", m.srcDriver, "psql_db", driver)
	if err != nil {
		return fmt.Errorf("unable to create migration: %v", err)
	}

	// Закрываем мигратор в конце работы функции.
	defer func() {
		migrator.Close()
	}()

	// Применяем миграции.
	if err = migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("unable to apply migrations %v", err)
	}

	return nil
}
