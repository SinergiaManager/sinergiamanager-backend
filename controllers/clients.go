package controllers

import (
	"time"

	Config "github.com/SinergiaManager/sinergiamanager-backend/config"
	Models "github.com/SinergiaManager/sinergiamanager-backend/models"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllClients(ctx iris.Context) {
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

	cursor, err := Config.DB.Collection("clients").Find(ctx, bson.M{}, findOptions)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	defer cursor.Close(ctx)

	var clients []*Models.ClientDb

	if err = cursor.All(ctx, &clients); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": clients})
}

func GetClient(ctx iris.Context) {
	client := &Models.ClientDb{}
	Id := ctx.Params().Get("id")

	objID, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid user ID format"})
		return
	}

	err = Config.DB.Collection("users").FindOne(ctx, bson.M{"_id": objID}).Decode(client)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.StatusCode(iris.StatusNotFound)
			ctx.JSON(iris.Map{"error": "User not found"})
		} else {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"error": "Failed to fetch user"})
		}
		return
	}

	ctx.JSON(iris.Map{"data": client})
}

func CreateClient(ctx iris.Context) {
	client := &Models.ClientIns{}
	if err := ctx.ReadJSON(client); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	client.InsertAt = time.Now().UTC()
	client.UpdateAt = time.Now().UTC()

	_, err := Config.DB.Collection("suppliers").InsertOne(ctx, client)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"message": "Client created successfully"})
}

func UpdateClient(ctx iris.Context) {
	client := &Models.ClientIns{}
	if err := ctx.ReadJSON(client); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	Id := ctx.Params().Get("id")
	objID, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid user ID format"})
		return
	}

	client.UpdateAt = time.Now().UTC()

	_, err = Config.DB.Collection("suppliers").UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": client})
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Client updated successfully"})
}

func DeleteClient(ctx iris.Context) {
	Id := ctx.Params().Get("id")
	objID, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid user ID format"})
		return
	}

	_, err = Config.DB.Collection("suppliers").DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Client deleted successfully"})
}
