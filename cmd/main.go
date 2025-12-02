package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"todo-app1"
	"todo-app1/pkg/handler"
	"todo-app1/pkg/repository"
	"todo-app1/pkg/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title Todo App API
// @version 1.0
// @description API Server for TodoList Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// 1️⃣ Configuration🧹🏦
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	// загружает переменные окружения из файла .env из корня
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	// 2️⃣ Repository🧹🏦
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	fmt.Printf("DB Host: %s, Port: %s, User: %s\n",
		viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.username"))

	repos := repository.NewRepository(db)

	// 3️⃣ Use case🧹🏦
	services := service.NewService(repos)

	// Router
	handlers := handler.NewHandler(services)

	// 4️⃣ HTTP Server🧹🏦
	srv := new(todo.Server)

	// Отдельная горутина: сервер запускается в своей собственной горутине.
	// Это необходимо, так как ListenAndServe() является блокирующим вызовом.
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")

	// ❗Graceful shutdown
	// quit: Это наш "стоп-кран".
	// Это буферизованный канал, который будет ожидать системные сигналы.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	// Корректное завершение (?)
	// Используем корневой контекст Background
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	// Close storage
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
