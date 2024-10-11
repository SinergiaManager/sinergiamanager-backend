package controllers

import (
	"time"

	Config "github.com/SinergiaManager/sinergiamanager-backend/config"
	Models "github.com/SinergiaManager/sinergiamanager-backend/models"
	"github.com/go-playground/validator/v10"
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
	err = Config.DB.Collection("users").FindOne(ctx, bson.M{"email": credentials.Email}).Decode(user)
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

	token, err := Config.GenerateToken(Config.Signer, user)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"token": string(token)})
}

func Logout(ctx iris.Context) {
	Config.Logout(ctx)
}

func Register(ctx iris.Context) {
	user := &Models.UserIns{}

	err := ctx.ReadJSON(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
			return
		}
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	var count int64 = 0
	count, err = Config.DB.Collection("users").CountDocuments(ctx, bson.M{"$or": []bson.M{{"email": user.Email}, {"username": user.Username}}})
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	if count > 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Email or username already exists"})
		return
	}

	user.InsertAt = time.Now().UTC()
	user.UpdateAt = time.Now().UTC()
	user.Role = string(Config.EnumUserRole.USER)

	_, err = Config.DB.Collection("users").InsertOne(ctx, user)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"message": "User created successfully"})
}
