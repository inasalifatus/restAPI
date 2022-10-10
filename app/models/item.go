package models

type Item struct {
	ItemId      int    `gorm:"column:item_id;primaryKey;autoIncrement"`
	ItemCode    string `gorm:"column:item_code"`
	Description string `gorm:"column:description"`
	Quantity    int    `gorm:"column:quantity"`
	OrderId     int    `gorm:"column:order_id"`
}
