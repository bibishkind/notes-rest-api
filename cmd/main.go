package main

import (
	. "github.com/bibishkin/bi-notes-rest-api"
	"github.com/bibishkin/bi-notes-rest-api/pkg/repository"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(new(logrus.JSONFormatter))
	logger.SetReportCaller(true)

	if err := initConfig(); err != nil {
		logrus.Fatalf("error intializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.name"),
	})
	if err != nil {
		logrus.Fatalf("error connecting to database: %s", err.Error())
	}

	repo := repository.NewRepositrory(db)

	srv := new(Server)
	if err := srv.Run(viper.GetString("port"), nil); err != nil {
		logger.Fatalf("error running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
