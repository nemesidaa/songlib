package storage_interfaces

import (
	"songlib/internal/sql/model"
)

type Store interface {
	Migrate() error
	Song() SongRepository
	Close() error
}

type SongRepository interface {
	Create(song *model.Song) error
	GetList(filter map[string]interface{}, page, pageSize int) ([]*model.Song, error)
	GetByID(id string) (*model.Song, error)
	Update(id string, updatedData map[string]interface{}) error
	Delete(id string) error
}
