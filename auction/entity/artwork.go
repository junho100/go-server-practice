package entity

type Artwork struct {
	ID       int
	Name     string
	OwnedBy  int
	Buyer    Buyer `gorm:"foreignKey:OwnedBy"`
	Auctions []Auction
}
