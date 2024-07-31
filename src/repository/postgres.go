package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/logger"
	//"github.com/Mamvriyskiy/shortLink/tree/develop/database/migration"
	
	// Импорт драйвера PostgreSQL для его регистрации.
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg *Config) (*sqlx.DB, error) {
	connectString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)

	db, err := sqlx.Open("postgres", connectString)
	if err != nil {
		logger.Log("Error", " sqlx.Open", "Error connect DB:", err, "postgres", connectString)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		logger.Log("Error", "Ping()", "Error check connection:", err, "")
		logger.Log("Error", " sqlx.Open", "Error connect DB:", err, "postgres", connectString)
		return nil, err
	}

	return db, nil
}

// Объяснение:

// *  `postgres://`  -  Схема подключения  к Postgres.
// *  `postgres:postgres`  -  Пользователь и пароль  для подключения.
// *  `@127.0.0.1:5432`  -  Адрес сервера и порт. 
// *  `/example`  -  Имя базы данных.

