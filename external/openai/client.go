package openai

import (
	"context"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/packages/param"
	"github.com/openai/openai-go/responses"
)

const imageGenerationToolName = "image_generation_call"

type Client struct {
	openai.Client
}

func NewClient(token string) Client {
	client := openai.NewClient(
		option.WithAPIKey(token),
	)

	return Client{client}
}

func (c Client) GenerateImage(prompt string, ctx context.Context) ([]string, error) {
	response, err := c.Responses.New(ctx, responses.ResponseNewParams{
		Input: responses.ResponseNewParamsInputUnion{OfString: param.Opt[string]{Value: prompt}},
		Model: openai.ChatModelGPT4_1,
		Tools: []responses.ToolUnionParam{
			{OfImageGeneration: &responses.ToolImageGenerationParam{
				Size: "1536x1024",
			}},
		},
	})

	if err != nil {
		return []string{}, err
	}

	var images []string

	for _, output := range response.Output {
		if output.Type == imageGenerationToolName {
			images = append(images, output.Result)
		}
	}

	return images, nil
}

func (c Client) GenerateText(prompt string, ctx context.Context) (string, error) {
	response, err := c.Responses.New(ctx, responses.ResponseNewParams{
		Input: responses.ResponseNewParamsInputUnion{OfString: param.Opt[string]{Value: prompt}},
		Model: openai.ChatModelGPT4_1,
	})

	if err != nil {
		return "", err
	}

	return response.OutputText(), nil
}
