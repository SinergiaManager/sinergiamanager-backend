package controllers

import (
	"time"

	ConfigDb "github.com/SinergiaManager/sinergiamanager-backend/config/database"
	Model "github.com/SinergiaManager/sinergiamanager-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllItems(ctx iris.Context) {
	cursor, err := ConfigDb.DB.Collection("items").Find(ctx, bson.M{})

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	defer cursor.Close(ctx)

	var items []*Model.ItemDB

	if err = cursor.All(ctx, &items); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": items})

}

func CreateItem(ctx iris.Context) {
	var item *Model.ItemIns
	err := ctx.ReadJSON(&item)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	item.InsertAt = time.Now().UTC()
	item.UpdateAt = time.Now().UTC()

	_, err = ConfigDb.DB.Collection("items").InsertOne(ctx, item)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"message": "Item created successfully"})
}

func UpdateItem(ctx iris.Context) {
	var updateData map[string]interface{}

	id := ctx.Params().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid user ID format"})
		return
	}

	err = ctx.ReadBody(&updateData)

	update := bson.D{{"$set", bson.D{}}}

	setFields := bson.D{}

	for key, value := range updateData {
		setFields = append(setFields, bson.E{key, value})
	}

	update[0].Value = setFields

	_, err = Config.DB.Collection("items").UpdateByID(ctx, objectID, update)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Item updated successfully"})
}

func DeleteItem(ctx iris.Context) {
	id := ctx.Params().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	filter := bson.M{"_id": objectID}

	result, err := ConfigDb.DB.Collection("items").DeleteOne(ctx, filter)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"error": "Item does not exist"})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Item deleted successfully"})
}
