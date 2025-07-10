package slack

import (
	"context"
	"github.com/slack-go/slack"
	"os"
	"path"
)

type Client struct {
	*slack.Client
	channel string
}

func NewClient(token string, channel string) Client {
	client := slack.New(token)

	return Client{client, channel}
}

func (c Client) SendMessageWithImage(filepath string, title string, description string, ctx context.Context) error {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	fileSize := fileInfo.Size()

	params := slack.UploadFileV2Parameters{
		File:           filepath,
		Filename:       path.Base(filepath),
		FileSize:       int(fileSize),
		Title:          description,
		Channel:        c.channel,
		InitialComment: title,
	}

	_, err = c.UploadFileV2Context(ctx, params)

	return err
}

func (c Client) SendMessage(text string, ctx context.Context) error {
	_, _, err := c.PostMessageContext(ctx, c.channel, slack.MsgOptionText(text, false))

	return err
}
