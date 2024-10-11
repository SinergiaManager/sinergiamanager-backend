package controllers

import (
	"time"

	Config "github.com/SinergiaManager/sinergiamanager-backend/config"
	Model "github.com/SinergiaManager/sinergiamanager-backend/models"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllUsers(ctx iris.Context) {
	limit, err := ctx.URLParamInt("limit")
	if err != nil || limit <= 0 {
		limit = 10 // default limit
	}

	skip, err := ctx.URLParamInt("skip")
	if err != nil || skip < 0 {
		skip = 0 // default skip
	}

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(skip))

	cursor, err := Config.DB.Collection("users").Find(ctx, bson.M{}, findOptions)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	defer cursor.Close(ctx)

	var users []*Model.UserOut

	if err = cursor.All(ctx, &users); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": users})
}

func CreateUser(ctx iris.Context) {
	var user *Model.UserIns
	err := ctx.ReadBody(&user)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	user.InsertAt = time.Now().UTC()
	user.UpdateAt = time.Now().UTC()

	_, err = Config.DB.Collection("users").InsertOne(ctx, user)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"message": "User created successfully"})
}

func UpdateUser(ctx iris.Context) {
	user := ctx.Values().Get("user").(Config.UserClaims)

	id := ctx.Params().Get("id")
	if user.Id != id && user.Role != "admin" {
		ctx.StatusCode(iris.StatusForbidden)
		ctx.JSON(iris.Map{"error": "You are not allowed to update this user"})
		return
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid user ID format"})
		return
	}

	var updateData = make(map[string]interface{})
	updateData["update_at"] = time.Now().UTC()

	ctx.ReadBody(&updateData)

	update := bson.D{{Key: "$set", Value: bson.D{}}}

	setFields := bson.D{}

	for key, value := range updateData {
		setFields = append(setFields, bson.E{Key: key, Value: value})
	}

	update[0].Value = setFields

	_, err = Config.DB.Collection("users").UpdateByID(ctx, objectID, update)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}

	updatedUser := &Model.UserOut{}
	err = Config.DB.Collection("users").FindOne(ctx, bson.M{"_id": objectID}).Decode(updatedUser)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": updatedUser})
}

func DeleteUser(ctx iris.Context) {
	id := ctx.Params().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid user ID format"})
		return
	}
	filter := bson.M{"_id": objectID}
	result, err := Config.DB.Collection("users").DeleteOne(ctx, filter)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"error": "User not found"})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "User deleted successfully"})
}
