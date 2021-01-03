package main

import (
	"simple-backend/server"
)

func main() {
	app := server.App{}
	app.InitialiseDatabase()
	app.InitialiseRoutes()
	app.Run()
}
