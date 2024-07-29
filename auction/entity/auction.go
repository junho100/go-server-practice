package entity

import (
	"time"
)

type Auction struct {
	ID        int
	EndDate   time.Time
	ArtworkID int
}
