package main

import (

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
		Username: "Misfio32",
		Password: "Smena",
		DBName: "shortlink",
		SSLMode: "disable",
	})

	if err != nil {
		//logger.Log("Error", "initCongig", "Error config DB:", err, "")
		//TODO: error
		return
	}

	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handler := handler.NewHandler(services)

	_ = handler
}

