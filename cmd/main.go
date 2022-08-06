package main

import (
	. "github.com/bibishkin/bi-notes-rest-api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(new(logrus.JSONFormatter))
	logger.SetReportCaller(true)

	if err := initConfig(); err != nil {
		logrus.Fatalf("error intializing configs: %s", err.Error())
	}

	srv := new(Server)
	if err := srv.Run(viper.GetString("port"), nil); err != nil {
		logger.Fatalf("error occured whlie running a http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
