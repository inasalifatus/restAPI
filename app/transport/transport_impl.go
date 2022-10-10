package transport

import (
	"restAPI/app/usecase"
	"restAPI/app/usecase/request"
	"strconv"

	"github.com/gin-gonic/gin"
)

type tp struct {
	usecase usecase.Usecase
}

func NewTransport(usecase usecase.Usecase) *tp {
	return &tp{
		usecase: usecase,
	}
}

func (t tp) CreateOrder(c *gin.Context) {
	var body request.CreateOrderRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	result, http_code, err := t.usecase.CreateOrder(&body)
	if err != nil {
		c.JSON(http_code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http_code, result)
}

func (t tp) FindAllOrder(c *gin.Context) {
	result, http_code, err := t.usecase.FindAllOrder()
	if err != nil {
		c.JSON(http_code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http_code, result)
}

func (t tp) FindOneOrder(c *gin.Context) {
	id := c.Param("orderId")
	order_id, _ := strconv.Atoi(id)
	result, http_code, err := t.usecase.FindOneOrder(order_id)
	if err != nil {
		c.JSON(http_code, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http_code, result)
}

func (t tp) UpdateOrder(c *gin.Context) {
	var body request.UpdateOrderRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}
	id := c.Param("orderId")
	order_id, _ := strconv.Atoi(id)
	result, http_code, err := t.usecase.UpdateOrder(order_id, &body)
	if err != nil {
		c.JSON(http_code, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http_code, result)
}

func (t tp) DeleteOrder(c *gin.Context) {
	id := c.Param("orderId")
	order_id, _ := strconv.Atoi(id)
	result, http_code, err := t.usecase.DeleteOrder(order_id)
	if err != nil {
		c.JSON(http_code, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http_code, result)
}
