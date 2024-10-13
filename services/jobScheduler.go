package services

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

// Esempio di funzione di notifica che sarà schedulata
func sendScheduledNotifications() {
	// Esegui ogni invio in una goroutine separata
	go func() {
		fmt.Println("Job iniziato:", time.Now())
		defer fmt.Println("Job completato:", time.Now())
		// Logica per inviare notifiche
	}()
}

func SetupJobScheduler() {
	scheduler := gocron.NewScheduler(time.UTC)

	scheduler.Every(1).Minute().Do(sendScheduledNotifications)

	scheduler.StartAsync()

	select {}
}
