package service

import (
	"auction/entity"
	"time"

	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func (svc *Service) CreateAuction(name string, endDate time.Time) (*entity.Auction, error) {
	artwork := &entity.Artwork{
		Name: name,
	}
	tx := svc.DB.Begin()

	if err := tx.Create(&artwork).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	auction := &entity.Auction{
		EndDate:   endDate,
		ArtworkID: artwork.ID,
	}

	if err := tx.Create(&auction).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return auction, tx.Commit().Error
}

func (svc *Service) CreateBidding(auctionID int, userID int, requestPrice int) (*entity.Bidding, error) {
	var (
		buyer   entity.Buyer
		auction entity.Auction
	)

	if err := svc.DB.First(&buyer, userID).Error; err != nil {
		return nil, err
	}
	if err := svc.DB.First(&auction, auctionID).Error; err != nil {
		return nil, err
	}

	bidding := &entity.Bidding{
		Timestamp:    time.Now(),
		BuyerID:      userID,
		AuctionID:    auctionID,
		RequestPrice: requestPrice,
	}

	if err := svc.DB.Create(&bidding).Error; err != nil {
		return nil, err
	}

	return bidding, nil
}
