package image_generator

import (
	"context"
	"errors"
	"threadzilla/external/openai"
)

type Client interface {
	GenerateImage(prompt string, ctx context.Context) ([]string, error)
}

type ImageGenerator struct {
	client Client
}

func NewImageGenerator(client *openai.Client) *ImageGenerator {
	return &ImageGenerator{
		client: client,
	}
}

func (s *ImageGenerator) GenerateDeployThreadImage(ctx context.Context) (string, error) {
	return s.generateImage(deployThreadImagePrompt, ctx)
}

func (s *ImageGenerator) GenerateReviewThreadImage(ctx context.Context) (string, error) {
	return s.generateImage(reviewThreadImagePrompt, ctx)
}

func (s *ImageGenerator) generateImage(prompt string, ctx context.Context) (string, error) {
	images, err := s.client.GenerateImage(prompt, ctx)
	if err != nil {
		return "", err
	}
	if len(images) == 0 {
		return "", errors.New("no images returned")
	}
	return images[0], nil
}
