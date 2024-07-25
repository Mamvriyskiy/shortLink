package main

import (
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/handler"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/service"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/repository"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src"
	"fmt"
)

func main() {
	// db, err := repository.NewPostgresDB(&repository.Config{
	// 	Host:     viper.GetString("db.host"),
	// 	Port:     viper.GetString("db.port"),
	// 	Username: viper.GetString("db.username"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// 	DBName:   viper.GetString("db.dbname"),
	// 	SSLMode:  viper.GetString("db.sslmode"),
	// })

	db, err := repository.NewPostgresDB(&repository.Config{
		Host: "localhost",
		Port: "5432",
		Username: "Mamre32",
		Password: "Smena",
		DBName: "postgres",
		SSLMode: "disable",
	})


	if err != nil {
		//fmt.Println("No connect")
		//logger.Log("Error", "initCongig", "Error config DB:", err, "")
		//TODO: error
		return
	}

	fmt.Println("DB connect")
	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	srv := new(src.Server)
	if err := srv.Run("8000", handlers.InitRouters()); err != nil {
		fmt.Println(err)
		//logger.Log("Error", "Run", "Error occurred while running http server:", err, "")
		return
	}
}

