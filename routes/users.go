package routes

import (
	"github.com/kataras/iris/v12"
	"haseeb.khan/event-booking/models"
	"haseeb.khan/event-booking/utils"
)

func signup(ctx iris.Context) {
	var user models.User

	err := ctx.ReadJSON(&user)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "Invalid request body"})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "unable to save user"})
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"message": "sign up successful"})
}

func login(ctx iris.Context) {
	var user models.User
	err := ctx.ReadJSON(&user)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "Invalid request body"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "login success", "token": token})
}
