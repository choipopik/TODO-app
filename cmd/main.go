package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/choipopik/todo-app"
	"github.com/choipopik/todo-app/pkg/handler"
	"github.com/choipopik/todo-app/pkg/repository"
	"github.com/choipopik/todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	err := initConfig()
	if err != nil {
		logrus.Fatalf("error reading configs: %s", err.Error())
	}

	err = godotenv.Load()
	if err != nil {
		logrus.Fatalf("error getting evn file:%s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error db init: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	go func() {
		err = srv.Run(viper.GetString("port"), handlers.InitRoutes())
		if err != nil {
			logrus.Fatalf("error running http server: %s", err.Error())
		}
	}()

	logrus.Print("App started...")

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logrus.Print("App shutdown...")

	err = srv.Shutdown(context.Background())
	if err != nil {
		logrus.Errorf("error shutting down server: %s", err.Error())
	}

	err = db.Close()
	if err != nil {
		logrus.Errorf("error closing db: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
