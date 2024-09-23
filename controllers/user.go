package controllers

import (
	Config "github.com/SinergiaManager/sinergiamanager-backend/config"
	Model "github.com/SinergiaManager/sinergiamanager-backend/models"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllUsers(ctx iris.Context) {
	cursor, err := Config.DB.Collection("users").Find(ctx, bson.D{})

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}

	defer cursor.Close(ctx)

	var users []*Model.User
	if err = cursor.All(ctx, &users); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": users})

}
