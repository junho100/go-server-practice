package entity

type Buyer struct {
	ID       int
	Balance  int
	Artworks []Artwork `gorm:"foreignKey:OwnedBy"`
}
