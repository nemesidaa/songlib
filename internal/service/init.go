package service

import (
	"songlib/internal/httpclient"
	"songlib/internal/logger"
	storage_interfaces "songlib/internal/sql/storage"
	storage "songlib/internal/sql/storage/implementation"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/sirupsen/logrus"

	_ "songlib/docs"
)

type Service struct {
	app    *fiber.App
	db     storage_interfaces.Store
	logger *logrus.Logger
	httpc  *httpclient.Client
}

func NewService(dbConnString string) (*Service, error) {
	logger := logger.GetLogger()
	db, err := storage.Init(logger, dbConnString)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if err := db.Migrate(); err != nil {
		logger.Error(err)
		return nil, err
	}
	return &Service{
		app:    fiber.New(),
		db:     db,
		logger: logger,
		httpc:  httpclient.NewClient(),
	}, nil
}

func (s *Service) ConfRoutes() {
	// Получение данных библиотеки с фильтрацией по всем полям и пагинацией
	s.app.Post("/songs", s.Filter)
	// Получение текста песни с пагинацией по куплетам
	s.app.Get("/song/:id", s.GetSong) // param page offset
	// Удаление
	s.app.Delete("/song/:id", s.DeleteSong)
	// Изменение данных
	s.app.Put("/song/:id", s.UpdateSong)
	// Добавление новой песни
	s.app.Post("/song", s.CreateSong)

	s.app.Get("/swagger/*", swagger.HandlerDefault) // Обработка Swagger UI

}

func (s *Service) Listen(host string, port string) error {
	return s.app.Listen(host + ":" + port)
}
