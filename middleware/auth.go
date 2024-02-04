package middleware

import (
	"github.com/kataras/iris/v12"
	"haseeb.khan/event-booking/utils"
)

func Authenticate(ctx iris.Context) {
	token := ctx.GetHeader("Authorization")

	if token == "" {
		ctx.StopWithJSON(iris.StatusUnauthorized, iris.Map{"message": "Not Authorized"})
		return
	}

	uid, err := utils.VerifyToken(token)
	if err != nil {
		ctx.StopWithJSON(iris.StatusUnauthorized, iris.Map{"message": "Not Authorized"})
		return
	}

	ctx.SetID(uid)
	ctx.Next()
}
