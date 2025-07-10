package thread_creator

import "context"

type Channel interface {
	SendMessageWithImage(base64Image, text string, ctx context.Context) error
	SendMessage(text string, ctx context.Context) error
}

type ImageGenerator interface {
	GenerateDeployThreadImage(ctx context.Context) (string, error)
	GenerateReviewThreadImage(ctx context.Context) (string, error)
}

type ThreadCreator struct {
	imageGenerator ImageGenerator
	channel        Channel
}

func NewThreadCreator(channel Channel) *ThreadCreator {
	return &ThreadCreator{
		channel: channel,
	}
}

func (tc *ThreadCreator) CreateDeployThread(ctx context.Context) error {

	image, err := tc.imageGenerator.GenerateDeployThreadImage(ctx)

	if err != nil {
		// todo: make funny message via OpenAI about we cant generate image today
		return tc.channel.SendMessage("Deploy thread", ctx)
	}

	return tc.channel.SendMessageWithImage(image, "Deploy thread", ctx)
}

func (tc *ThreadCreator) CreateReviewThread(ctx context.Context) error {
	image, err := tc.imageGenerator.GenerateReviewThreadImage(ctx)

	if err != nil {
		// todo: make funny message via OpenAI about we cant generate image today
		return tc.channel.SendMessage("Review thread", ctx)
	}

	return tc.channel.SendMessageWithImage(image, "Review thread", ctx)
}
