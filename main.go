package main

import "restAPI/app/interface/container"

func main() {
	container := container.SetupContainer()
	server.SetupServer(container)
}
