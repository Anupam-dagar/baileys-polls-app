package main

import (
	"github.com/Anupam-dagar/baileys/server"
	"polls/route"
)

func main() {
	server.NewGinEngine().InitGinApp(route.SetupRoutes).RunServer()
}
