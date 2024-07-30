package service

import (
	"auction/entity"
	"time"

	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func (svc *Service) CreateAuction(name string, endDate time.Time) error {
	artwork := &entity.Artwork{
		Name: name,
	}
	tx := svc.DB.Begin()

	if err := tx.Create(&artwork).Error; err != nil {
		tx.Rollback()
		return err
	}

	auction := &entity.Auction{
		EndDate:   endDate,
		ArtworkID: artwork.ID,
	}

	if err := tx.Create(&auction).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}