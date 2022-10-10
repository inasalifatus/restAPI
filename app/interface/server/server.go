package server

import (
	"restAPI/app/interface/container"
	"restAPI/app/shared/config"
	"restAPI/app/transport"

	"github.com/gin-gonic/gin"
)

func SetupServer(container container.Container) {
	app := gin.Default()
	transport := transport.SetupTransport(container)
	SetupRouter(transport, app)

	app.Run(config.Server.PORTHTTP)
}
