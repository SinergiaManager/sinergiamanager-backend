package controllers

import (
	"time"

	Config "github.com/SinergiaManager/sinergiamanager-backend/config"
	Models "github.com/SinergiaManager/sinergiamanager-backend/models"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllWarehouses(ctx iris.Context) {
	limit, err := ctx.URLParamInt("limit")
	if err != nil || limit <= 0 {
		limit = 10
	}

	skip, err := ctx.URLParamInt("skip")
	if err != nil || skip < 0 {
		skip = 0
	}

	pipeline := mongo.Pipeline{
		{{"$lookup", bson.D{
			{"from", "items"},
			{"localField", "items.item_id"},
			{"foreignField", "_id"},
			{"as", "itemsDetails"},
		}},
			{"$skip", skip},
			{"$limit", limit}}}

	cursor, err := Config.DB.Collection("warehouses").Aggregate(ctx, pipeline)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	defer cursor.Close(ctx)

	var warehouses []*bson.M
	if err = cursor.All(ctx, &warehouses); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": warehouses})

}

func GetWarehouseById(ctx iris.Context) {
	id := ctx.Params().Get("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	pipeline := mongo.Pipeline{
		{{"$match", bson.D{{"_id", objectID}}}},
		{{"$lookup", bson.D{
			{"from", "items"},
			{"localField", "items.item_id"},
			{"foreignField", "_id"},
			{"as", "itemsDetails"},
		}}}}

	var warehouse *bson.M

	cursor, err := Config.DB.Collection("warehouses").Aggregate(ctx, pipeline)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		if err = cursor.Decode(&warehouse); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}
	}

	ctx.JSON(iris.Map{"data": warehouse})
}

func CreateWarehouse(ctx iris.Context) {
	var warehouse *Models.WarehouseIns
	err := ctx.ReadBody(&warehouse)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	warehouse.InsertAt = time.Now().UTC()
	warehouse.UpdateAt = time.Now().UTC()

	_, err = Config.DB.Collection("warehouses").InsertOne(ctx, warehouse)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"message": "Warehouse created successfully"})

}

func UpdateWarehouse(ctx iris.Context) {
	var updateData = make(map[string]interface{})

	id := ctx.Params().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid warehouse ID format"})
		return
	}

	ctx.ReadBody(&updateData)

	update := bson.D{{Key: "$set", Value: bson.D{}}}

	setFields := bson.D{}

	for key, value := range updateData {
		setFields = append(setFields, bson.E{Key: key, Value: value})
	}

	update[0].Value = setFields

	updateData["updateAt"] = time.Now().UTC()

	_, err = Config.DB.Collection("warehouses").UpdateOne(ctx, objectID, update)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	updatedData := Config.DB.Collection("warehouses").FindOne(ctx, bson.M{"_id": objectID})

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": updatedData})

}

func DeleteWarehouse(ctx iris.Context) {
	id := ctx.Params().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid warehouse ID format"})
		return
	}
	filter := bson.M{"_id": objectID}
	result, err := Config.DB.Collection("warehouses").DeleteOne(ctx, filter)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"error": "Warehouse not found"})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Warehouse deleted successfully"})
}
