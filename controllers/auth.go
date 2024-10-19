package controllers

import (
	"strings"
	"time"

	Config "github.com/SinergiaManager/sinergiamanager-backend/config"
	Models "github.com/SinergiaManager/sinergiamanager-backend/models"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
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

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
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
	credentials := &Models.Register{}

	err := ctx.ReadJSON(credentials)
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

	user := &Models.UserIns{}
	user.Email = credentials.Email
	user.Password = credentials.Password
	user.Username = credentials.Username
	user.Name = credentials.Name
	user.Surname = credentials.Surname

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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	user.Password = string(hashedPassword)

	_, err = Config.DB.Collection("users").InsertOne(ctx, user)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"message": "User created successfully"})
}

func RefreshToken(ctx iris.Context) {
	tokenAuth := ctx.GetHeader("Authorization")
	if tokenAuth == "" {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"error": "Not authenticated"})
		return
	}

	tokenParts := strings.Split(tokenAuth, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"error": "Invalid token format"})
		return
	}

	token := tokenParts[1]

	newToken, err := Config.RefreshToken(Config.Signer, []byte(token))
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"token": string(newToken)})
}
