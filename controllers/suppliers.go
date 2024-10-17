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

func GetAllSuppliers(ctx iris.Context) {
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

	cursor, err := Config.DB.Collection("suppliers").Find(ctx, bson.M{}, findOptions)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	defer cursor.Close(ctx)

	var suppliers []*Models.SupplierDb

	if err = cursor.All(ctx, &suppliers); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": suppliers})
}

func GetSupplier(ctx iris.Context) {
	supplier := &Models.SupplierDb{}
	Id := ctx.Params().Get("id")

	objID, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid user ID format"})
		return
	}

	err = Config.DB.Collection("users").FindOne(ctx, bson.M{"_id": objID}).Decode(supplier)
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

	ctx.JSON(iris.Map{"data": supplier})
}

func CreateSupplier(ctx iris.Context) {
	supplier := &Models.SupplierIns{}
	if err := ctx.ReadJSON(supplier); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	supplier.InsertAt = time.Now().UTC()
	supplier.UpdateAt = time.Now().UTC()

	_, err := Config.DB.Collection("suppliers").InsertOne(ctx, supplier)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"message": "Supplier created successfully"})
}

func UpdateSupplier(ctx iris.Context) {
	supplier := &Models.SupplierIns{}
	if err := ctx.ReadJSON(supplier); err != nil {
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

	supplier.UpdateAt = time.Now().UTC()

	_, err = Config.DB.Collection("suppliers").UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": supplier})
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Supplier updated successfully"})
}

func DeleteSupplier(ctx iris.Context) {
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
	ctx.JSON(iris.Map{"message": "Supplier deleted successfully"})
}
