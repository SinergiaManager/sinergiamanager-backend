package services

import (
	"context"
	"time"

	"github.com/go-co-op/gocron"
)

func SetupJobScheduler(ctx context.Context) {
	scheduler := gocron.NewScheduler(time.UTC)

	//scheduler.Every(1).Minute().Do(SendScheduledNotificationsByEmail, ctx)
	scheduler.StartAsync()

	//TestEmail()
	select {}
}
