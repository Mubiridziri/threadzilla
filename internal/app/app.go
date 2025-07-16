package app

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"text/tabwriter"
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

	config.InitLogger()

	openaiClient := openai.NewClient(cfg.OpenAI.Token)
	slackClient := slack.NewClient(cfg.Slack.Token, cfg.Slack.Channel)
	imageGenerator := image_generator.NewImageGenerator(openaiClient)
	textGenerator := text_generator.NewTextGenerator(openaiClient)
	threadCreator := thread_creator.NewThreadCreator(imageGenerator, textGenerator, slackClient, cfg.WithImage)

	fmt.Println("Starting threadzilla...")

	PrintStartupConfig(*cfg)

	taskFunc := func(ctx context.Context) {
		fmt.Println("Creating deploy thread...")
		err = threadCreator.CreateDeployThread(ctx)

		if err != nil {
			log.Errorf("error creating deploy thread: %v", err)
			fmt.Println("Error creating deploy thread:", err)
		}

		fmt.Println("Creating review thread...")
		err = threadCreator.CreateReviewThread(ctx)

		if err != nil {
			log.Errorf("error creating review thread: %v", err)
			fmt.Println("Error creating review thread:", err)
		}
	}

	if cfg.DaemonMode {
		runDaemon(ctx, cfg.SendingHourAt, cfg.SendingMinuteAt, func(ctx context.Context) {
			taskFunc(ctx)
		})
	} else {
		taskFunc(ctx)
	}

	return nil
}

func runDaemon(ctx context.Context, hour, minute int, fn func(context.Context)) {
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

func PrintStartupConfig(config config.Config) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	_, _ = fmt.Fprintln(w, "===============================")
	_, _ = fmt.Fprintf(w, "Daemon mode\t%t\t\n", config.DaemonMode)
	_, _ = fmt.Fprintf(w, "Sending time\t%d:%d\t\n", config.Common.SendingHourAt, config.Common.SendingMinuteAt)
	_, _ = fmt.Fprintf(w, "Image generating\t%t\t\n", config.OpenAI.WithImage)
	_, _ = fmt.Fprintln(w, "===============================")
	_ = w.Flush()
}
