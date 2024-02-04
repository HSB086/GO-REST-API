package routes

import (
	"github.com/kataras/iris/v12"
	"haseeb.khan/event-booking/middleware"
)

func InitializeRoutes(server *iris.Application) {
	server.Get("/events", getEvents)
	server.Get("/events/{id:int64}", getEvent)

	protected := server.Party("/")
	protected.Use(middleware.Authenticate)
	protected.Post("/events", createEvent)
	protected.Put("/events/{id:int64}", updateEvent)
	protected.Delete("/events/{id:int64}", deleteEvent)
	protected.Post("/events/{id:int64}/register", registerForEvent)
	protected.Delete("/events/{id:int64}/unregister", unregisterForEvent)

	/* server.Post("/events", middleware.Authenticate, createEvent)
	server.Put("/events/{id:int64}", updateEvent)
	server.Delete("/events/{id:int64}", deleteEvent) */
	server.Post("/signup", signup)
	server.Post("/login", login)
}
