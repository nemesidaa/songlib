package main

import (
	"log"
	"os"
	"songlib/internal/logger"
	"songlib/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные окружения из .env файла
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	logger.InitLogger()
	logger := logger.GetLogger()
	logger.Info("Starting service...")
	service, err := service.NewService(os.Getenv("DB_CONNSTR"))
	if err != nil {
		logger.Fatal(err)
	}

	service.ConfRoutes()
	err = service.Listen(os.Getenv("HOST"), os.Getenv("PORT"))
	if err != nil {
		logger.Fatal(err)
	}
}
