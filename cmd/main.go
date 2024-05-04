package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"

	"github.com/AndrewFilkin/for-any-test"
	"github.com/AndrewFilkin/for-any-test/pkg/handler"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	go func() {
		route := handler.InitRoutes()
		srv := new(server.Server)
		if err := srv.Run(viper.GetString("port"), route); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Article-CRUD-Test Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Article-CRUD-Test Shutting Down")
}
