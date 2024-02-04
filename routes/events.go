package routes

import (
	"github.com/kataras/iris/v12"
	"haseeb.khan/event-booking/models"
)

func getEvents(ctx iris.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "Could not fetch events."})
		return
	}
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(events)
}

func getEvent(ctx iris.Context) {
	id, err := ctx.Params().GetInt64("id")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "invalid request."})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "Could not fetch event by provided id."})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(event)
}

func createEvent(ctx iris.Context) {
	/* token := ctx.GetHeader("Authorization")

	if token == "" {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"message": "Not Authorized"})
		return
	}

	uid, err := utils.VerifyToken(token)
	if err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"message": "Not Authorized"})
		return
	} */

	var event models.Event
	err := ctx.ReadJSON(&event)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "invalid request body"})
		return
	}

	uid := ctx.GetID().(int64)

	event.UserId = uid
	err = event.Save()
	if err != nil {
		// fmt.Println(err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "unable to create event", "cause": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"message": "Event Created", "event": event})
}

func updateEvent(ctx iris.Context) {
	id, err := ctx.Params().GetInt64("id")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "invalid request."})
		return
	}

	uid := ctx.GetID().(int64)
	event, err := models.GetEventById(id)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "Could not fetch event by provided id."})
		return
	}

	if event.UserId != uid {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"message": "Not Authorised to update: Event created by another user."})
		return
	}

	var updatedEvent models.Event
	err = ctx.ReadJSON(&updatedEvent)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "invalid request body"})
		return
	}

	updatedEvent.ID = id
	err = updatedEvent.UpdateEvent()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "Could not fetch event by provided id."})
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"message": "Event Updated"})
}

func deleteEvent(ctx iris.Context) {
	id, err := ctx.Params().GetInt64("id")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "invalid request."})
		return
	}

	uid := ctx.GetID().(int64)
	event, err := models.GetEventById(id)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "Could not fetch event by provided id."})
		return
	}

	if event.UserId != uid {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"message": "Not Authorised to delete: Event created by another user."})
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "Could not delete event by provided id."})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Event Deleted"})
}
