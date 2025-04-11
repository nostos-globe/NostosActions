package db

import (
	"main/internal/models"

	"gorm.io/gorm"
)

type AlbumRepository struct {
	DB *gorm.DB
}

func (repo *AlbumRepository) CreateAlbum(album models.Album) (any, error) {
	result := repo.DB.Table("albums.albums").Create(&album)

	if result.Error != nil {
		return nil, result.Error
	}

	return album, nil
}