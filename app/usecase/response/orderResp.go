package response

import "time"

type Item struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	ItemId      int    `json:"item_id"`
}

type Order struct {
	OrderId      int       `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items"`
}

type CreateOrderResponse struct {
	Order Order `json:"order"`
}

type FindAllOrderResponse struct {
	Order []Order `json:"order"`
}

type FindOneOrderResponse struct {
	Order Order `json:"order"`
}

type UpdateOrderResponse struct {
	Order Order `json:"order"`
}

type DeleteOrderResponse struct {
	Message string `json:"message"`
}
