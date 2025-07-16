package thread_creator

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	filemanager "threadzilla/internal/service/file-manager"
	"time"
)

type Channel interface {
	SendMessageWithImage(filepath string, title string, description string, ctx context.Context) error
	SendMessage(text string, ctx context.Context) error
}

type ImageGenerator interface {
	GenerateDeployThreadImage(ctx context.Context) (string, error)
	GenerateReviewThreadImage(ctx context.Context) (string, error)
}

type TextGenerator interface {
	GenerateNoImageReasonText(ctx context.Context) string
	GenerateInterestingFactText(ctx context.Context) string
}

type FileManager interface {
	SaveBase64File(filepath string, encodedContent string) error
	DeleteFile(filepath string) error
}

type ThreadCreator struct {
	imageGenerator ImageGenerator
	textGenerator  TextGenerator
	channel        Channel
	fileManager    FileManager
	withImage      bool
}

func NewThreadCreator(imageGenerator ImageGenerator, textGenerator TextGenerator, channel Channel, withImage bool) *ThreadCreator {
	return &ThreadCreator{
		imageGenerator: imageGenerator,
		textGenerator:  textGenerator,
		channel:        channel,
		fileManager:    filemanager.FileManager{},
		withImage:      withImage,
	}
}

func (tc *ThreadCreator) CreateDeployThread(ctx context.Context) error {
	currentDate := time.Now().Format("02.01.2006")
	fact := tc.textGenerator.GenerateInterestingFactText(ctx)
	title := fmt.Sprintf("Деплой тред на %s \n", currentDate)
	description := fmt.Sprintf("Интересный факт на сегодня: \n "+
		"%s", fact)

	if !tc.withImage {
		return tc.channel.SendMessage(title+"\n\n"+description, ctx)
	}

	filepath, err := tc.getImage(DeployThread, ctx)

	if err != nil {
		log.Errorf("error when generating image: %v", err)

		reason := tc.textGenerator.GenerateNoImageReasonText(ctx)

		description += "\n\n" + reason
		return tc.channel.SendMessage(title+"\n\n"+description, ctx)
	}

	return tc.channel.SendMessageWithImage(filepath, title, description, ctx)
}

func (tc *ThreadCreator) CreateReviewThread(ctx context.Context) error {
	currentDate := time.Now().Format("02.01.2006")
	fact := tc.textGenerator.GenerateInterestingFactText(ctx)
	title := fmt.Sprintf("Ревью тред на %s \n", currentDate)
	description := fmt.Sprintf("Интересный факт на сегодня: \n "+
		"%s", fact)

	if !tc.withImage {
		return tc.channel.SendMessage(title+"\n\n"+description, ctx)
	}

	filepath, err := tc.getImage(ReviewThread, ctx)

	if err != nil {
		log.Errorf("error when generating image: %v", err)

		reason := tc.textGenerator.GenerateNoImageReasonText(ctx)

		description += "\n\n" + reason

		return tc.channel.SendMessage(title+"\n\n"+description, ctx)
	}

	return tc.channel.SendMessageWithImage(filepath, title, description, ctx)
}

func (tc *ThreadCreator) getImage(action Action, ctx context.Context) (string, error) {
	handlers := map[Action]func(context.Context) (string, error){
		DeployThread: tc.imageGenerator.GenerateDeployThreadImage,
		ReviewThread: tc.imageGenerator.GenerateReviewThreadImage,
	}

	generateFunc, ok := handlers[action]
	if !ok {
		return "", fmt.Errorf("unknown action: %v", action)
	}

	image, err := generateFunc(ctx)
	if err != nil {
		return "", fmt.Errorf("error when generating image: %w", err)
	}

	pwd, err := os.Getwd()

	if err != nil {
		return "", fmt.Errorf("error when getting current directory: %w", err)
	}

	fileName := fmt.Sprintf("%s.png", action)
	filePath := path.Join(pwd, fileName)

	if err = tc.fileManager.SaveBase64File(filePath, image); err != nil {
		return "", fmt.Errorf("error when saving image: %w", err)
	}

	return filePath, nil
}
