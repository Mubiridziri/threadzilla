package app

import (
	"context"
	"fmt"
	"threadzilla/internal/config"
	"time"
)
import "threadzilla/external/openai"

type Application struct {
}

func (a *Application) Run(ctx context.Context) error {
	configLoader := config.Loader{}

	cfg, err := configLoader.LoadConfig()
	if err != nil {
		return err
	}

	_ = openai.NewClient(cfg.OpenAI.Token)

	runDailyJob(ctx, cfg.SendingHourAt, cfg.SendingMinuteAt, func(ctx context.Context) {})

	return nil
}

func runDailyJob(ctx context.Context, hour, minute int, fn func(context.Context)) {
	for {
		now := time.Now()
		next := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())
		if !next.After(now) {
			next = next.Add(24 * time.Hour)
		}

		sleep := time.Until(next)
		fmt.Printf("Next run: %s (in %s)\n", next.Format(time.RFC1123), sleep.Round(time.Second))

		select {
		case <-ctx.Done():
			fmt.Println("Stopping before start.")
			return
		case <-time.After(sleep):
			fn(ctx)
		}
	}
}
