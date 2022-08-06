package main

import (
	. "github.com/bibishkin/bi-notes-rest-api"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(new(logrus.JSONFormatter))
	logger.SetReportCaller(true)

	srv := new(Server)
	if err := srv.Run("8080111", nil); err != nil {
		logger.Fatalf("error occured whlie running a http server: %s", err.Error())
	}
}
