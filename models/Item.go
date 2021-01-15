package models

type Item struct {
	ID         uint
	ItemName   string
	Price      float64
	Stock      int32
	CategoryID uint
	Category   Category
}

func (Item) TableName() string {
	return "products.item"
}
