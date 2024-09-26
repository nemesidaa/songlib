package storage

import (
	"songlib/internal/sql/model"
	interfaces "songlib/internal/sql/storage"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	logger   *logrus.Logger
	songRepo interfaces.SongRepository
	db       *gorm.DB
}

/*
dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",

	os.Getenv("DB_HOST"),
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"),
	os.Getenv("DB_PORT"),

)
*/
func Init(logger *logrus.Logger, dsn string) (*Store, error) {

	var err error
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return &Store{
		logger: logger,
		db:     DB,
	}, nil
}

func (impl *Store) Migrate() error {
	return impl.db.AutoMigrate(&model.Song{})
}

func (impl *Store) Song() interfaces.SongRepository {
	if impl.songRepo == nil {
		impl.songRepo = &SongRepository{
			db: impl.db,
		}
	}
	return impl.songRepo
}

func (impl *Store) Close() error {
	sqlDB, err := impl.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
