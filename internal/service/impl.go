package service

import (
	"songlib/internal/sql/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

/*
Methods:
	Получение данных библиотеки с фильтрацией по всем полям и пагинацией
	Получение текста песни с пагинацией по куплетам
	Удаление песни
	Изменение данных песни
	Добавление новой песни в формате:
	{
 	"group": "Muse",
	"song": "Supermassive Black Hole"
	}
*/

func (s *Service) Filter(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		s.logger.Warnf("Error parsing page in /songs: %v", err)
		page = 1
	}
	s.logger.Infof("Got page %d in /songs", page)
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		s.logger.Warnf("Error parsing size in /songs: %v", err)
		size = 10
	}
	s.logger.Infof("Got size %d in /songs", size)
	req := new(ListRequest)
	if err := c.BodyParser(req); err != nil {
		s.logger.Errorf("Error parsing body in /songs: %v", err)
		return c.Status(500).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	s.logger.Infof("Got filter %v in /songs", req.Filtermap)
	songs, err := s.db.Song().GetList(req.Filtermap, page, size)
	if err != nil {
		s.logger.Errorf("Error getting songs in /songs: %v", err)
		return c.Status(500).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	s.logger.Infof("Got %d songs in /songs, handled", len(songs))
	return c.Status(200).JSON(songs)
}

func (s *Service) GetSong(c *fiber.Ctx) error {
	id := c.Params("id")
	song, err := s.db.Song().GetByID(id)
	if err != nil {
		s.logger.Errorf("Error getting song in /song: %v", err)
		return c.Status(500).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	s.logger.Infof("Got song %s in /song, handled", id)
	return c.Status(200).JSON(song)
}

func (s *Service) DeleteSong(c *fiber.Ctx) error {
	id := c.Params("id")
	err := s.db.Song().Delete(id)
	if err != nil {
		s.logger.Errorf("Error deleting song in /song: %v", err)
		return c.Status(500).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	s.logger.Infof("Deleted song %s in /song, handled", id)
	return c.Status(200).JSON(&fiber.Map{
		"message": "ok",
	})
}

func (s *Service) UpdateSong(c *fiber.Ctx) error {
	id := c.Params("id")
	req := new(UpdateRequest)
	if err := c.BodyParser(req); err != nil {
		s.logger.Errorf("Error parsing body in /song: %v", err)
		return c.Status(500).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	if err := s.db.Song().Update(id, req.Data); err != nil {
		s.logger.Errorf("Error updating song in /song: %v", err)
		return c.Status(500).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	s.logger.Infof("Updated song %s in /song, handled", id)
	return c.Status(200).JSON(&fiber.Map{
		"message": "ok",
	})
}

func (s *Service) CreateSong(c *fiber.Ctx) error {
	req := new(CreateRequest)
	if err := c.BodyParser(req); err != nil {
		s.logger.Errorf("Error parsing body in /song: %v", err)
		return c.Status(500).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	song := &model.Song{
		Group: req.Group,
		Song:  req.Song,
	}
	data, err := s.httpc.DataMock()
	if err != nil {
		s.logger.Errorf("Error getting data in /song: %v", err)
		return c.Status(500).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	song.Link = data.Link
	song.Text = data.Text
	song.ReleaseDate = data.ReleaseDate

	err = s.db.Song().Create(song)
	if err != nil {
		s.logger.Errorf("Error creating song in /song: %v", err)
		return c.Status(500).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	s.logger.Infof("Created song %s in /song, handled", req.Song)
	return c.Status(200).JSON(song)
}
