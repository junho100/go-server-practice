package entity

import (
	"time"
)

type Bidding struct {
	Timestamp    time.Time `gorm:"primaryKey"`
	BuyerID      int       `gorm:"primaryKey"`
	AuctionID    int       `gorm:"primaryKey"`
	RequestPrice int
}
