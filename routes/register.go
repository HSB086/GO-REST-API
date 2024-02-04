package routes

import (
	"github.com/kataras/iris/v12"
	"haseeb.khan/event-booking/models"
)

func registerForEvent(ctx iris.Context) {
	uId := ctx.GetID().(int64)
	eventId, err := ctx.Params().GetInt64("id")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "Invalid eventId Passed."})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"message": "No event found for eventId."})
		return
	}

	err = event.Register(uId)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "Unabale to register for event."})
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"message": "Registered Successfully !"})
}

func unregisterForEvent(ctx iris.Context) {
	uId := ctx.GetID().(int64)
	eventId, err := ctx.Params().GetInt64("id")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "Invalid eventId Passed."})
		return
	}

	var event models.Event
	event.ID = eventId
	err = event.Unregister(uId)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "Unabale to cancel registeration for event."})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Unregistered Successfully !"})
}
