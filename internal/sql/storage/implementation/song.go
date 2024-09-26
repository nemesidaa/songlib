package storage

import (
	"songlib/internal/sql/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SongRepository struct {
	db *gorm.DB
}

func (impl *SongRepository) Create(song *model.Song) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	song.ID = id.String()
	return impl.db.Create(song).Error
}

func (impl *SongRepository) GetList(filter map[string]interface{}, page, pageSize int) ([]*model.Song, error) {
	var songs []*model.Song
	query := impl.db.Model(&model.Song{})

	// Применение фильтров
	for key, value := range filter {
		query = query.Where(key+" = ?", value)
	}

	if err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&songs).Error; err != nil {
		return nil, err
	}

	return songs, nil
}

func (impl *SongRepository) GetByID(id string) (*model.Song, error) {
	var song model.Song
	if err := impl.db.Where("id = ?", id).First(&song).Error; err != nil {
		return nil, err
	}
	return &song, nil
}

func (impl *SongRepository) Update(id string, updatedData map[string]interface{}) error {
	var song model.Song
	// Проверяем наличие записи
	if err := impl.db.First(&song, "id = ?", id).Error; err != nil {
		return err // Возвращаем ошибку, если запись не найдена
	}

	// Обновляем запись
	if err := impl.db.Model(&song).Updates(updatedData).Error; err != nil {
		return err // Возвращаем ошибку при обновлении
	}

	return nil // Возвращаем nil, если обновление прошло успешно
}

func (impl *SongRepository) Delete(id string) error {
	return impl.db.Where("id = ?", id).Delete(&model.Song{}).Error
}
