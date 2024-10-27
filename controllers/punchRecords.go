package controllers

import (
	"time"

	Config "github.com/SinergiaManager/sinergiamanager-backend/config"
	Models "github.com/SinergiaManager/sinergiamanager-backend/models"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllPunchRecords(ctx iris.Context) {
	limit, err := ctx.URLParamInt("limit")
	if err != nil || limit <= 0 {
		limit = 10
	}

	skip, err := ctx.URLParamInt("skip")
	if err != nil || skip < 0 {
		skip = 0
	}

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(skip))

	cursor, err := Config.DB.Collection("punchRecords").Find(ctx, bson.M{}, findOptions)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	defer cursor.Close(ctx)

	var punchRecords []*Models.PunchRecordDb

	if err = cursor.All(ctx, &punchRecords); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": punchRecords})
}

func GetPunchRecord(ctx iris.Context) {
	id := ctx.Params().Get("id")

	punchRecord := &Models.PunchRecordDb{}
	if err := Config.DB.Collection("punchRecords").FindOne(ctx, bson.M{"_id": id}).Decode(punchRecord); err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": punchRecord})
}

func CreatePunchRecord(ctx iris.Context) {
	punchRecord := &Models.PunchRecordIns{}
	if err := ctx.ReadJSON(punchRecord); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	punchRecord.InsertAt = time.Now().UTC()
	punchRecord.UpdateAt = time.Now().UTC()
	punchRecord.PunchTime = time.Now().UTC()

	if _, err := Config.DB.Collection("punchRecords").InsertOne(ctx, punchRecord); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"data": punchRecord})
}

func UpdatePunchRecord(ctx iris.Context) {
	var updateData = make(map[string]interface{})

	id := ctx.Params().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid user ID format"})
		return
	}

	ctx.ReadBody(&updateData)

	update := bson.D{{Key: "$set", Value: bson.D{}}}

	setFields := bson.D{}

	for key, value := range updateData {
		setFields = append(setFields, bson.E{Key: key, Value: value})
	}

	update[0].Value = setFields

	updateData["update_at"] = time.Now().UTC()

	_, err = Config.DB.Collection("punchRecords").UpdateByID(ctx, objectID, update)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}

	updatedData := Config.DB.Collection("items").FindOne(ctx, bson.M{"_id": objectID})

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": updatedData})
}

func DeletePunchRecord(ctx iris.Context) {
	id := ctx.Params().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid user ID format"})
		return
	}

	_, err = Config.DB.Collection("punchRecords").DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Punch record deleted successfully"})
}

func GetPunchRecordByUser(ctx iris.Context) {
	id := ctx.Params().Get("id")

	punchRecord := &Models.PunchRecordDb{}
	if err := Config.DB.Collection("punchRecords").FindOne(ctx, bson.M{"employee_id": id}).Decode(punchRecord); err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": punchRecord})
}

func GetPunchRecordByMe(ctx iris.Context) {
	user := ctx.Values().Get("user").(*Models.UserDb)

	punchRecord := &Models.PunchRecordDb{}
	if err := Config.DB.Collection("punchRecords").FindOne(ctx, bson.M{"employee_id": user.ID}).Decode(punchRecord); err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": punchRecord})
}

func GetPunchRecordByDate(ctx iris.Context) {
	date := ctx.Params().Get("date")

	punchRecord := &Models.PunchRecordDb{}
	if err := Config.DB.Collection("punchRecords").FindOne(ctx, bson.M{"punch_time": date}).Decode(punchRecord); err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": punchRecord})
}
