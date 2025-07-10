package image_generator

import (
	"context"
	"errors"
	"testing"
)

type mockClient struct {
	GenerateImageFunc func(prompt string, ctx context.Context) ([]string, error)
}

func (m *mockClient) GenerateImage(prompt string, ctx context.Context) ([]string, error) {
	return m.GenerateImageFunc(prompt, ctx)
}

func TestGenerateDeployThreadImage_Success(t *testing.T) {
	mc := &mockClient{
		GenerateImageFunc: func(prompt string, ctx context.Context) ([]string, error) {
			return []string{"img1.png", "img2.png"}, nil
		},
	}
	ig := &ImageGenerator{client: mc}
	img, err := ig.GenerateDeployThreadImage(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if img != "img1.png" {
		t.Errorf("expected img1.png, got %s", img)
	}
}

func TestGenerateReviewThreadImage_Success(t *testing.T) {
	mc := &mockClient{
		GenerateImageFunc: func(prompt string, ctx context.Context) ([]string, error) {
			return []string{"imgX.jpg"}, nil
		},
	}
	ig := &ImageGenerator{client: mc}
	img, err := ig.GenerateReviewThreadImage(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if img != "imgX.jpg" {
		t.Errorf("expected imgX.jpg, got %s", img)
	}
}

func TestGenerateImage_Error(t *testing.T) {
	mc := &mockClient{
		GenerateImageFunc: func(prompt string, ctx context.Context) ([]string, error) {
			return nil, errors.New("fail")
		},
	}
	ig := &ImageGenerator{client: mc}
	_, err := ig.generateImage("any", context.Background())
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestGenerateImage_EmptyResult(t *testing.T) {
	mc := &mockClient{
		GenerateImageFunc: func(prompt string, ctx context.Context) ([]string, error) {
			return []string{}, nil
		},
	}
	ig := &ImageGenerator{client: mc}
	_, err := ig.generateImage("any", context.Background())

	if err == nil {
		t.Error("expected error, got nil")
	}

	if err != nil && err.Error() != "no images returned" {
		t.Errorf("expected error \"no images returned\", got %v", err)
	}
}
