package order

import "restAPI/app/models"

type OrderRepository interface {
	CreateOrder(order *models.Order) (result *models.Order, err error)
	CreateItem(item *models.Item) (result *models.Item, err error)
	FindAllOrder() (result *[]models.Order, err error)
	FindOneOrder(order_id int) (result *models.Order, err error)
	UpdateOrder(order_id int, order *models.Order) (result *models.Order, err error)
	UpdateItem(item_id int, item *models.Item) (result *models.Item, err error)
	DeleteOrder(order_id int) (err error)
}
