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

	// Configuration🧹🏦
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	// загружает переменные окружения из файла .env из корня
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

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

	// ❗ В достижении цели “разделения ответственности между всеми слоями приложения” нам помогает “правило
	// зависимости” (это о круговой диаграмме дяди Боба).
	// Зависимости направлены только внутрь (внутренний круг ничего не должен знать про внешний
	// и сущности внутреннего круга не могут обратиться к сущностям внешнего).
	// ❗ И вот чтобы реализовать “Правило зависимости” мы используем технику dependency injection !

	// Создаем сущности слоев (это three-layered architecture)
	// в порядке обратном обращению к ним:
	//
	// 3️⃣ Repository (DAL - Data Access Layer)
	repository := repository.NewRepository(db)
	// ⬇ Сервисам нужно то, что делает репозиторий (CRUD).
	// 2️⃣ Use case (BL - Business Logic Layer, service)
	// | Здесь внедряем зависимость с repository
	services := service.NewService(repository)
	// ⬇ Хендлерам нужно то, что делает сервис (обслуживание операций CRUD + GenerateToken, ParseToken)
	// 1️⃣ Handler (PL - Presentation Layer, controller)
	// | Здесь внедряем зависимость с services
	handlers := handler.NewHandler(services)
	// ↑ HTTP request

	// Работает в обратном раправлении!
	// HTTP запрос -> ручка -> обращение к службе -> служба к базе данных.

	// HTTP Server🧹🏦
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
