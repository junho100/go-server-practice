package service

import (
	"auction/entity"
	"errors"
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

func (svc *Service) GetExpiredAuctionByTime(nowTime time.Time) ([]entity.Auction, error) {
	var auctions []entity.Auction

	err := svc.DB.Where("end_date < ? AND status = ?", nowTime, "ACTIVE").Find(&auctions).Error

	if err != nil {
		return []entity.Auction{}, err
	}

	return auctions, nil
}

func (svc *Service) TerminateAuction(auction *entity.Auction) {
	var (
		bidding entity.Bidding
		buyer   entity.Buyer
		artwork entity.Artwork
	)
	tx := svc.DB.Begin()

	err := svc.DB.Order("request_price desc").Where("auction_id = ?", auction.ID).First(&bidding).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		auction.Status = "FAILED"
		svc.DB.Save(&auction)
		return
	}

	svc.DB.First(&buyer, bidding.BuyerID)
	buyer.Balance -= bidding.RequestPrice
	svc.DB.Save(&buyer)

	svc.DB.First(&artwork, auction.ArtworkID)
	artwork.Buyer = &buyer
	artwork.OwnedBy = &buyer.ID
	svc.DB.Save(&artwork)

	auction.Status = "TERMINATED"
	svc.DB.Save(&auction)

	tx.Commit()
}
