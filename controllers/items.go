package controllers

import (
	Config "github.com/SinergiaManager/sinergiamanager-backend/config"
	Model "github.com/SinergiaManager/sinergiamanager-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllItems(ctx iris.Context) {
	cursor, err := Config.DB.Collection("items").Find(ctx, bson.M{})

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

	_, err = Config.DB.Collection("items").InsertOne(ctx, item)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"message": "Item created successfully"})
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

	result, err := Config.DB.Collection("items").DeleteOne(ctx, filter)
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
