package main

import (
	"log"
	"os"
	"songlib/internal/logger"
	"songlib/internal/service"

	"github.com/joho/godotenv"
)

// @title Song Library API
// @version 1.0
// @description This is a sample API for managing a song library
// @contact.name API Support
// @contact.url tg:https://t.me/nowerealldied
// @contact.email egor200619@gmail.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8081
// @BasePath /

func main() {
	// Загружаем переменные окружения из .env файла
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	logger.InitLogger(os.Getenv("LOG_DEST"), os.Getenv("LOG_LEVEL"))
	logger := logger.GetLogger()
	logger.Info("Starting service...")
	service, err := service.NewService(os.Getenv("DB_CONNSTR"), os.Getenv("API_URL"))
	if err != nil {
		logger.Fatal(err)
	}
	defer service.CloseDBConn()

	service.ConfRoutes()

	err = service.Listen(os.Getenv("HOST"), os.Getenv("PORT"))
	if err != nil {
		logger.Fatal(err)
	}
}
