package entity

type Artwork struct {
	ID       int
	Name     string `gorm:"unique"`
	OwnedBy  int
	Buyer    Buyer `gorm:"foreignKey:OwnedBy"`
	Auctions []Auction
}