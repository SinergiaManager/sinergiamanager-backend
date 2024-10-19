package services

import (
	"context"
	"fmt"
	"log"
	"time"

	gomail "gopkg.in/gomail.v2"

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

			message := gomail.NewMessage()

			message.SetHeader("From", "support@singergiamanager.com")
			message.SetHeader("To", notification["email"].(string))
			message.SetHeader("Subject", "Notifica da SinergiaManager - "+notification["title"].(string))

			message.SetBody("text/html", fmt.Sprintf("<h1>%v</h1><p>%v</p>", notification["title"], notification["message"]))

			dialer := gomail.NewDialer("mailserver", 2500, "", "") // No username/password for now..
			if err := dialer.DialAndSend(message); err != nil {
				fmt.Println("Error:", err)
			}

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

func TestEmail() {
	notification := map[string]interface{}{
		"email":   "fabio04musitelli@gmail.com",
		"title":   "Test Notification",
		"message": "This is a test message from SinergiaManager.",
	}

	message := gomail.NewMessage()
	message.SetHeader("From", "support@singergiamanager.com")
	message.SetHeader("To", notification["email"].(string))
	message.SetHeader("Subject", "Notifica da SinergiaManager - "+notification["title"].(string))
	message.SetBody("text/html", fmt.Sprintf("<h1>%v</h1><p>%v</p>", notification["title"], notification["message"]))

	dialer := gomail.NewDialer("mailserverTest", 1025, "", "") // No username/password for MailHog

	if err := dialer.DialAndSend(message); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}
