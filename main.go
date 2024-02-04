package main

import (
	"github.com/kataras/iris/v12"
	"haseeb.khan/event-booking/database"
	"haseeb.khan/event-booking/routes"
)

func main() {
	database.InitDB()

	server := iris.Default()

	routes.InitializeRoutes(server)

	server.Run(iris.Addr(":8080"))
}
