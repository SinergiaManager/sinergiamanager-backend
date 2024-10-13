package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	Config "github.com/SinergiaManager/sinergiamanager-backend/config"
)

// Esempio di funzione di notifica che sarà schedulata per l'esec invio di notifiche via email
func SendScheduledNotifications(ctx context.Context) {
	go func() {
		fmt.Println("Job iniziato:", time.Now())
		defer fmt.Println("Job completato:", time.Now())
		fmt.Println("Simulazione di invio di notifiche...")

		filter := bson.M{"isDelivered": false}

		cursor, err := Config.DB.Collection("notifications").Find(ctx, filter)
		if err != nil {
			log.Println("Error finding notifications:", err)
			return
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var notification bson.M
			if err := cursor.Decode(&notification); err != nil {
				log.Println("Error decoding notification:", err)
				continue
			}

			fmt.Printf("Inviando notifica a: %v\n", notification["email"])

			update := bson.M{"$set": bson.M{"isDelivered": true, "deliveredAt": time.Now()}}
			_, err = Config.DB.Collection("notifications").UpdateOne(ctx, bson.M{"_id": notification["_id"]}, update)
			if err != nil {
				log.Println("Error updating notification status:", err)
				continue
			}

			fmt.Printf("Notifica consegnata a: %v\n", notification["email"])
		}

		if err := cursor.Err(); err != nil {
			log.Println("Cursor error:", err)
		}
	}()
}
