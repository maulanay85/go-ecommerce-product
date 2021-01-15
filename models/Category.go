package models

import "time"

type Category struct {
	ID           uint
	CategoryName string
	IsDeleted    bool
	CreatedAt    time.Time
	updateAt     time.Time
}

func (Category) TableName() string {
	return "products.category"
}
