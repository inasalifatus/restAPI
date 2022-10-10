package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	OrderId      int       `gorm:"colum:order_id;primaryKey;autoIncrement"`
	CustomerName string    `gorm:"column:customer_name"`
	OrderedAt    time.Time `gorm:"column:ordered_at"`
	Items        []Item
}

func (order *Order) BeforeCreate(tx *gorm.DB) (err error) {
	order.OrderedAt = time.Now()
	return nil
}
