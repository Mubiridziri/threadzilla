package app

import (
	"context"
	"fmt"
	"threadzilla/external/slack"
	"threadzilla/internal/config"
	image_generator "threadzilla/internal/service/image-generator"
	text_generator "threadzilla/internal/service/text-generator"
	thread_creator "threadzilla/internal/service/thread-creator"
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

	openaiClient := openai.NewClient(cfg.OpenAI.Token)
	slackClient := slack.NewClient(cfg.Slack.Token, cfg.Slack.Channel)
	imageGenerator := image_generator.NewImageGenerator(openaiClient)
	textGenerator := text_generator.NewTextGenerator(openaiClient)
	threadCreator := thread_creator.NewThreadCreator(imageGenerator, textGenerator, slackClient)

	runDailyJob(ctx, cfg.SendingHourAt, cfg.SendingMinuteAt, func(ctx context.Context) {
		err = threadCreator.CreateDeployThread(ctx)

		if err != nil {
			fmt.Println("Error creating deploy thread:", err)
		}

		err = threadCreator.CreateReviewThread(ctx)

		if err != nil {
			fmt.Println("Error creating review thread:", err)
		}
	})

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
