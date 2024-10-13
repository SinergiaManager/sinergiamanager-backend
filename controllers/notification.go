package controllers

import (
	"time"

	Config "github.com/SinergiaManager/sinergiamanager-backend/config"
	Models "github.com/SinergiaManager/sinergiamanager-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllNotifications(ctx iris.Context) {
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

	cursor, err := Config.DB.Collection("notifications").Find(ctx, bson.M{}, findOptions)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	defer cursor.Close(ctx)

	var notifications []*Models.NotificationDb

	if err = cursor.All(ctx, &notifications); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": notifications})
}

func GetNotification(ctx iris.Context) {
	id := ctx.Params().Get("id")

	var notification Models.NotificationDb
	err := Config.DB.Collection("notifications").FindOne(ctx, bson.M{"_id": id}).Decode(&notification)
	if err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": notification})
}

func GetNotificationsByUser(ctx iris.Context) {
	userID := ctx.Params().Get("userID")

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

	cursor, err := Config.DB.Collection("notifications").Find(ctx, bson.M{"userID": userID}, findOptions)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	defer cursor.Close(ctx)

	var notifications []*Models.NotificationDb

	if err = cursor.All(ctx, &notifications); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": notifications})
}

func GetNotificationsMe(ctx iris.Context) {
	Id := ctx.Values().Get("user").(Config.UserClaims).Id

	objID, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid user ID format"})
		return
	}

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

	cursor, err := Config.DB.Collection("notifications").Find(ctx, bson.M{"userID": objID, "isRead": false}, findOptions)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	defer cursor.Close(ctx)

	var notifications []*Models.NotificationDb

	if err = cursor.All(ctx, &notifications); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": notifications})
}

func GetNotificationMe(ctx iris.Context) {
	Id := ctx.Values().Get("user").(Config.UserClaims).Id

	objID, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid user ID format"})
		return
	}

	id := ctx.Params().Get("id")

	var notification Models.NotificationDb
	err = Config.DB.Collection("notifications").FindOne(ctx, bson.M{"_id": id, "userID": objID}).Decode(&notification)
	if err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": notification})
}

func CreateNotification(ctx iris.Context) {
	var notification Models.NotificationDb
	if err := ctx.ReadJSON(&notification); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	notification.InsertAt = time.Now().UTC()
	notification.UpdateAt = time.Now().UTC()

	/* if notification types has email wait to set delivered */
	isDelivered := true
	for _, t := range notification.Types {
		if t == string(Config.EnumNotificationType.EMAIL) {
			isDelivered = false
			break
		}
	}

	if isDelivered {
		notification.IsDelivered = true
		notification.DeliveredAt = time.Now().UTC()
	}

	res, err := Config.DB.Collection("notifications").InsertOne(ctx, notification)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"data": res.InsertedID})
}

func UpdateNotification(ctx iris.Context) {
	id := ctx.Params().Get("id")

	var notification Models.NotificationDb
	if err := ctx.ReadJSON(&notification); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	notification.UpdateAt = time.Now().UTC()

	update := bson.D{{Key: "$set", Value: notification}}

	_, err := Config.DB.Collection("notifications").UpdateByID(ctx, id, update)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": id})
}

func DeleteNotification(ctx iris.Context) {
	id := ctx.Params().Get("id")

	_, err := Config.DB.Collection("notifications").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": id})
}

func ReadNotification(ctx iris.Context) {
	id := ctx.Params().Get("id")
	userID := ctx.Values().Get("user").(Config.UserClaims).Id

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid user ID format"})
		return
	}

	update := bson.M{"$set": bson.M{"isRead": true, "readAt": time.Now().UTC()}}

	_, err = Config.DB.Collection("notifications").UpdateOne(ctx, bson.M{"_id": id, "userID": objID}, update)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": id})
}
