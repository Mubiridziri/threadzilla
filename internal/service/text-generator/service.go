package text_generator

import "context"

type Client interface {
	GenerateText(prompt string, ctx context.Context) (string, error)
}

type TextGenerator struct {
	client Client
}

func NewTextGenerator(client Client) *TextGenerator {
	return &TextGenerator{
		client: client,
	}
}

func (s *TextGenerator) GenerateNoImageReasonText(ctx context.Context) string {
	text, _ := s.client.GenerateText(NoImageReasonPrompt, ctx)

	return text
}

func (s *TextGenerator) GenerateInterestingFactText(ctx context.Context) string {
	text, _ := s.client.GenerateText(InterestingFact, ctx)

	return text
}
