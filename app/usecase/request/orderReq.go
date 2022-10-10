package request

type Item struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	ItemId      int    `json:"item_id"`
}

type CreateOrderRequest struct {
	CustomerName string `json:"customer_name"`
	Items        []Item `json:"items"`
}

type UpdateOrderRequest struct {
	CustomerName string `json:"customer_name"`
	Items        []Item `json:"items"`
}
