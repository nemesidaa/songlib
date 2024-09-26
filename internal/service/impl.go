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

// Filter godoc
// @Summary Get filtered songs with pagination
// @Description Retrieves a list of songs filtered by specified fields with pagination support
// @Tags songs
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param size query int false "Number of items per page" default(10)
// @Param filter body ListRequest true "Filter parameters"
// @Success 200 {object} ListResponse
// @Failure 500 {object} ErrorResponse
// @Router /songs [post]
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
		return c.Status(500).JSON(&ErrorResponse{
			Message: err.Error(),
		})
	}
	s.logger.Infof("Got filter %v in /songs", req.Filtermap)
	songs, err := s.db.Song().GetList(req.Filtermap, page, size)
	if err != nil {
		s.logger.Errorf("Error getting songs in /songs: %v", err)
		return c.Status(500).JSON(&ErrorResponse{
			Message: err.Error(),
		})
	}
	s.logger.Infof("Got %d songs in /songs, handled", len(songs))
	return c.Status(200).JSON(songs)
}

// GetSong godoc
// @Summary Get song details
// @Description Retrieves details of a song by ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path string true "Song ID"
// @Success 200 {object} model.Song
// @Failure 500 {object} ErrorResponse
// @Router /song/{id} [get]
func (s *Service) GetSong(c *fiber.Ctx) error {
	id := c.Params("id")
	song, err := s.db.Song().GetByID(id)
	if err != nil {
		s.logger.Errorf("Error getting song in /song: %v", err)
		return c.Status(500).JSON(&ErrorResponse{
			Message: err.Error(),
		})
	}
	s.logger.Infof("Got song %s in /song, handled", id)
	return c.Status(200).JSON(song)
}

// DeleteSong godoc
// @Summary Delete a song
// @Description Deletes a song by ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path string true "Song ID"
// @Success 200
// @Failure 500 {object} ErrorResponse
// @Router /song/{id} [delete]
func (s *Service) DeleteSong(c *fiber.Ctx) error {
	id := c.Params("id")
	err := s.db.Song().Delete(id)
	if err != nil {
		s.logger.Errorf("Error deleting song in /song: %v", err)
		return c.Status(500).JSON(&ErrorResponse{
			Message: err.Error(),
		})
	}

	s.logger.Infof("Deleted song %s in /song, handled", id)
	return c.SendStatus(200)
}

// UpdateSong godoc
// @Summary Update a song
// @Description Updates song details by ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path string true "Song ID"
// @Param data body UpdateRequest true "Updated song data"
// @Success 200
// @Failure 500 {object} ErrorResponse
// @Router /song/{id} [put]
func (s *Service) UpdateSong(c *fiber.Ctx) error {
	id := c.Params("id")
	req := new(UpdateRequest)
	if err := c.BodyParser(req); err != nil {
		s.logger.Errorf("Error parsing body in /song: %v", err)
		return c.Status(500).JSON(&ErrorResponse{
			Message: err.Error(),
		})
	}

	if err := s.db.Song().Update(id, req.Data); err != nil {
		s.logger.Errorf("Error updating song in /song: %v", err)
		return c.Status(500).JSON(&ErrorResponse{
			Message: err.Error(),
		})
	}
	s.logger.Infof("Updated song %s in /song, handled", id)
	return c.SendStatus(200)
}

// CreateSong godoc
// @Summary Create a new song
// @Description Adds a new song to the library
// @Tags songs
// @Accept json
// @Produce json
// @Param data body CreateRequest true "Song creation data"
// @Success 200 {object} model.Song
// @Failure 500 {object} ErrorResponse
// @Router /song [post]
func (s *Service) CreateSong(c *fiber.Ctx) error {
	req := new(CreateRequest)
	if err := c.BodyParser(req); err != nil {
		s.logger.Errorf("Error parsing body in /song: %v", err)
		return c.Status(500).JSON(&ErrorResponse{
			Message: err.Error(),
		})
	}
	song := &model.Song{
		Group: req.Group,
		Song:  req.Song,
	}
	data, err := s.httpc.DataMock()
	if err != nil {
		s.logger.Errorf("Error getting data in /song: %v", err)
		return c.Status(500).JSON(&ErrorResponse{
			Message: err.Error(),
		})
	}
	song.Link = data.Link
	song.Text = data.Text
	song.ReleaseDate = data.ReleaseDate

	err = s.db.Song().Create(song)
	if err != nil {
		s.logger.Errorf("Error creating song in /song: %v", err)
		return c.Status(500).JSON(&ErrorResponse{
			Message: err.Error(),
		})
	}
	s.logger.Infof("Created song %s in /song, handled", req.Song)
	return c.Status(200).JSON(song)
}
