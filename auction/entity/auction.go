package entity

import (
	"time"
)

type Auction struct {
	ID        int
	EndDate   time.Time
	Status    string `gorm:"default:ACTIVE"`
	ArtworkID int    `gorm:"not null"`
}
