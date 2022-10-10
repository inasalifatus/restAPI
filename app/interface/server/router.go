package server

import (
	"restAPI/app/transport"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /orders
// @schemes http
func SetupRouter(transport *transport.Tp, app *gin.Engine) {
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	orderRoute := app.Group("/orders")
	{
		orderRoute.POST("/", transport.Transport.CreateOrder)
		orderRoute.GET("/", transport.Transport.FindAllOrder)
		orderRoute.GET("/:orderId", transport.Transport.FindOneOrder)
		orderRoute.PUT("/:orderId", transport.Transport.UpdateOrder)
		orderRoute.DELETE("/:orderId", transport.Transport.DeleteOrder)
	}
}
