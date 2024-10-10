package controllers

import (
	Auth "github.com/SinergiaManager/sinergiamanager-backend/config/auth"
	ConfigDb "github.com/SinergiaManager/sinergiamanager-backend/config/database"
	Models "github.com/SinergiaManager/sinergiamanager-backend/models"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(ctx iris.Context) {
	credentials := &Models.Login{}
	err := ctx.ReadJSON(credentials)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	user := &Models.UserDb{}
	err = ConfigDb.DB.Collection("users").FindOne(ctx, bson.M{"email": credentials.Email}).Decode(user)
	if err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"error": "Invalid credentials"})
		return
	}

	if user.Password != credentials.Password {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"error": "Invalid credentials"})
		return
	}

	token := Auth.GenerateToken(Auth.Signer, user)

	ctx.JSON(iris.Map{"token": token})
}

func Logout(ctx iris.Context) {
	Auth.Logout(ctx)
}

func Register() {

}
