package text_generator

import (
	"context"
	"errors"
	"testing"
)

type mockClient struct {
	GenerateTextFunc func(prompt string, ctx context.Context) (string, error)
}

func (m *mockClient) GenerateText(prompt string, ctx context.Context) (string, error) {
	return m.GenerateTextFunc(prompt, ctx)
}

func TestGenerateNoImageReasonText(t *testing.T) {
	mc := &mockClient{
		GenerateTextFunc: func(prompt string, ctx context.Context) (string, error) {
			if prompt != NoImageReasonPrompt {
				t.Errorf("expected prompt %q, got %q", NoImageReasonPrompt, prompt)
			}
			return "no image today", nil
		},
	}
	tg := NewTextGenerator(mc)
	result := tg.GenerateNoImageReasonText(context.Background())
	if result != "no image today" {
		t.Errorf("expected 'no image today', got %q", result)
	}
}

func TestGenerateInterestingFactText(t *testing.T) {
	mc := &mockClient{
		GenerateTextFunc: func(prompt string, ctx context.Context) (string, error) {
			if prompt != InterestingFact {
				t.Errorf("expected prompt %q, got %q", InterestingFact, prompt)
			}
			return "an interesting fact", nil
		},
	}
	tg := NewTextGenerator(mc)
	result := tg.GenerateInterestingFactText(context.Background())
	if result != "an interesting fact" {
		t.Errorf("expected 'an interesting fact', got %q", result)
	}
}

func TestGenerateNoImageReasonText_Error(t *testing.T) {
	mc := &mockClient{
		GenerateTextFunc: func(prompt string, ctx context.Context) (string, error) {
			return "", errors.New("fail")
		},
	}
	tg := NewTextGenerator(mc)
	result := tg.GenerateNoImageReasonText(context.Background())
	if result != "" {
		t.Errorf("expected empty string on error, got %q", result)
	}
}

func TestGenerateInterestingFactText_Error(t *testing.T) {
	mc := &mockClient{
		GenerateTextFunc: func(prompt string, ctx context.Context) (string, error) {
			return "", errors.New("fail")
		},
	}
	tg := NewTextGenerator(mc)
	result := tg.GenerateInterestingFactText(context.Background())
	if result != "" {
		t.Errorf("expected empty string on error, got %q", result)
	}
}
