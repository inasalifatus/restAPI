package usecase

import (
	"fmt"
	"restAPI/app/models"
	"restAPI/app/repository/order"
	"restAPI/app/usecase/request"
	"restAPI/app/usecase/response"
)

type usecase struct {
	orderRepository order.OrderRepository
}

func NewUsecase(order order.OrderRepository) Usecase {
	return &usecase{
		orderRepository: order,
	}
}

// Create One Order
// @Summary Mendaftarkan/membuat satu order
// @Description Mendaftarkan/membuat satu order
// @Tags Orders
// @Accept  */*
// @Produce  json
// @Param data body request.CreateOrderRequest true "Order"
// @Success 200 {object} response.CreateOrderResponse
// @Failure 500 {object} string "error"
// @Router /orders [post]
func (uc usecase) CreateOrder(body *request.CreateOrderRequest) (result *response.CreateOrderResponse, http_code int, err error) {
	var arrItem []response.Item
	order := &models.Order{
		CustomerName: body.CustomerName,
	}
	resOrder, err := uc.orderRepository.CreateOrder(order)
	if err != nil {
		return result, 400, err
	}
	for _, val := range body.Items {
		item := &models.Item{
			ItemCode:    val.ItemCode,
			Description: val.Description,
			Quantity:    val.Quantity,
			OrderId:     resOrder.OrderId,
		}
		resItem, err := uc.orderRepository.CreateItem(item)
		if err != nil {
			return result, 400, err
		}
		itemOne := response.Item{
			ItemCode:    resItem.ItemCode,
			Description: resItem.Description,
			Quantity:    resItem.Quantity,
			ItemId:      resItem.ItemId,
		}
		arrItem = append(arrItem, itemOne)
	}
	resTemp := &response.CreateOrderResponse{
		Order: response.Order{
			CustomerName: resOrder.CustomerName,
			OrderedAt:    resOrder.OrderedAt,
			OrderId:      resOrder.OrderId,
			Items:        arrItem,
		},
	}

	result = resTemp

	return result, 200, nil
}

// Get All Orders
// @Summary Mendapatkan semua order yang telah di generate
// @Description Mendapatkan semua order yang telah di generate
// @Tags Orders
// @Accept  */*
// @Produce  json
// @Success 200 {object} response.FindAllOrderResponse
// @Failure 500 {object} string "error"
// @Router /orders [get]
func (uc usecase) FindAllOrder() (result *response.FindAllOrderResponse, http_code int, err error) {
	var resp []response.Order
	res, err := uc.orderRepository.FindAllOrder()
	if err != nil {
		return result, 400, nil
	}
	for _, val := range *res {
		fmt.Println(res)
		var items []response.Item
		item := val.Items
		for _, valItem := range item {
			itemVal := response.Item{
				ItemCode:    valItem.ItemCode,
				Description: valItem.Description,
				Quantity:    valItem.Quantity,
				ItemId:      valItem.ItemId,
			}
			items = append(items, itemVal)
		}
		order := response.Order{
			OrderId:      val.OrderId,
			CustomerName: val.CustomerName,
			OrderedAt:    val.OrderedAt,
			Items:        items,
		}
		resp = append(resp, order)
	}

	result = &response.FindAllOrderResponse{
		Order: resp,
	}

	return result, 200, nil
}

// Get One Order
// @Summary Mendapatkan satu order yang telah dibuat dengan id
// @Description Mendapatkan satu order yang telah dibuat dengan id
// @Tags Orders
// @Accept  */*
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} response.FindOneOrderResponse
// @Failure 500 {object} string "error"
// @Router /orders/{id} [get]
func (uc usecase) FindOneOrder(id int) (result *response.FindOneOrderResponse, http_code int, err error) {
	var resp []response.Item
	res, err := uc.orderRepository.FindOneOrder(id)
	if err != nil {
		return result, 400, err
	}
	for _, val := range res.Items {
		itemVal := response.Item{
			ItemCode:    val.ItemCode,
			Description: val.Description,
			Quantity:    val.Quantity,
			ItemId:      val.ItemId,
		}
		resp = append(resp, itemVal)
	}
	order := response.Order{
		CustomerName: res.CustomerName,
		Items:        resp,
		OrderId:      res.OrderId,
		OrderedAt:    res.OrderedAt,
	}
	result = &response.FindOneOrderResponse{
		Order: order,
	}
	return result, 200, nil
}

// Update One Order
// @Summary Mengupdate satu Order yang sudah terbuat berdasarkan id
// @Description Mengupdate satu Order yang sudah terbuat berdasarkan id
// @Tags Orders
// @Accept  */*
// @Produce  json
// @Param id path int true "id"
// @Param data body request.UpdateOrderRequest true "Order"
// @Success 200 {object} response.UpdateOrderResponse
// @Failure 500 {object} string "error"
// @Router /orders/{id} [put]
func (uc usecase) UpdateOrder(id int, body *request.UpdateOrderRequest) (result *response.UpdateOrderResponse, http_code int, err error) {
	var items []response.Item
	order := &models.Order{
		CustomerName: body.CustomerName,
	}
	resOrder, err := uc.orderRepository.UpdateOrder(id, order)
	if err != nil {
		return result, 400, err
	}
	fmt.Println(resOrder, "res order")
	for _, val := range body.Items {
		item := &models.Item{
			ItemCode:    val.ItemCode,
			Description: val.Description,
			Quantity:    val.Quantity,
			ItemId:      val.ItemId,
		}
		resItem, err := uc.orderRepository.UpdateItem(val.ItemId, item)
		fmt.Println(resItem, "res item")
		if err != nil {
			return result, 400, err
		}
		responsItem := response.Item{
			ItemCode:    resItem.ItemCode,
			Description: resItem.Description,
			Quantity:    resItem.Quantity,
			ItemId:      resItem.ItemId,
		}
		items = append(items, responsItem)
	}
	responseOrder := &response.Order{
		CustomerName: resOrder.CustomerName,
		OrderedAt:    resOrder.OrderedAt,
		OrderId:      resOrder.OrderId,
		Items:        items,
	}
	result = &response.UpdateOrderResponse{
		Order: *responseOrder,
	}
	return result, 200, nil
}

// Delete
// @Summary Menghapus Order yang sudah dibuat berdasarkan id
// @Description Menghapus Order yang sudah dibuat berdasarkan id
// @Tags Orders
// @Accept  */*
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} response.DeleteOrderResponse
// @Failure 400 {object} string "error"
// @Router /orders/{id} [delete]
func (uc usecase) DeleteOrder(id int) (result *response.DeleteOrderResponse, http_code int, err error) {
	err = uc.orderRepository.DeleteOrder(id)
	if err != nil {
		return result, 400, err
	}
	result = &response.DeleteOrderResponse{
		Message: "Success to delete",
	}
	return result, 200, nil
}
