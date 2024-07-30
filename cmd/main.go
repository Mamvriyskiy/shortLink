package main

import (
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/handler"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/service"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/repository"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/logger"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src"
	"github.com/spf13/viper"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		logger.Log("Error", "initCongig", "Error config DB:", err)
		return
	}
	logger.Log("Info", "", "InitConfig", nil)

	if err := godotenv.Load(); err != nil {
		logger.Log("Error", "Load", "Load env file:", err, "")
		return
	}
	logger.Log("Info", "", "Load env", nil)

	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logger.Log("Error", "initCongig", "Error config DB:", err, "")
		return
	}

	logger.Log("Info", "", "Connect to db", nil)
	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	srv := new(src.Server)
	if err := srv.Run(viper.GetString("server.port"), handlers.InitRouters()); err != nil {
		logger.Log("Error", "Run", "Error occurred while running http server:", err, "")
		return
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
