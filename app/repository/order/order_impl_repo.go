package order

import (
	"fmt"
	"restAPI/app/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (or orderRepository) CreateOrder(order *models.Order) (result *models.Order, err error) {
	err = or.db.Create(order).Error
	if err != nil {
		return result, err
	}
	return order, nil
}

func (or orderRepository) CreateItem(item *models.Item) (result *models.Item, err error) {
	err = or.db.Create(item).Error
	if err != nil {
		return result, err
	}
	return item, nil
}

func (or orderRepository) FindAllOrder() (result *[]models.Order, err error) {
	var order []models.Order
	err = or.db.Find(&order).Error
	for key, val := range order {
		var items []models.Item
		err = or.db.Model(&models.Item{}).Where("order_id = ?", val.OrderId).Find(&items).Error
		val.Items = items
		order[key] = val
	}
	result = &order
	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}
	return result, nil
}

func (or orderRepository) FindOneOrder(order_id int) (result *models.Order, err error) {
	var order models.Order
	var items []models.Item
	err = or.db.Where("order_id = ?", order_id).First(&order).Error
	if err != nil {
		return result, err
	}
	err = or.db.Where("order_id = ?", order_id).Find(&items).Error
	if err != nil {
		return result, err
	}
	order.Items = items
	fmt.Println(order, "order")
	result = &order
	return result, nil
}

func (or orderRepository) UpdateOrder(order_id int, order *models.Order) (result *models.Order, err error) {
	var orders models.Order
	err = or.db.Model(&orders).Clauses(clause.Returning{}).Where("order_id = ?", order_id).Updates(order).Error
	if err != nil {
		return result, err
	}
	err = or.db.Where("order_id = ?", order_id).First(&orders).Error
	if err != nil {
		return result, err
	}
	fmt.Println(orders, "orders")
	return &orders, nil
}

func (or orderRepository) UpdateItem(item_id int, item *models.Item) (result *models.Item, err error) {
	var items models.Item
	err = or.db.Model(&items).Clauses(clause.Returning{}).Where("item_id = ?", item_id).Updates(item).Error
	if err != nil {
		return result, err
	}
	err = or.db.Where("item_id = ?", item_id).First(&items).Error
	if err != nil {
		return result, err
	}
	fmt.Println(items, "items")
	return &items, nil
}

func (or orderRepository) DeleteOrder(order_id int) (err error) {
	err = or.db.Where("order_id = ?", order_id).Delete(&models.Item{}).Error
	if err != nil {
		return err
	}
	err = or.db.Delete(&models.Order{}, order_id).Error
	if err != nil {
		return err
	}
	return nil
}
