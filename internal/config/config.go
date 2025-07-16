package config

import "fmt"
import "threadzilla/internal/utils"

const OpenAIToken = "OPENAI_TOKEN"
const SlackToken = "SLACK_TOKEN"
const SlackChannel = "SLACK_CHANNEL"
const SendingHourAt = "SENDING_HOUR_AT"
const DaemonMode = "DAEMON_MODE"
const WithImage = "OPENAI_GENERATING_WITH_IMAGE"

type Loader struct {
}

type Config struct {
	Common
	OpenAI
	Slack
}

type Common struct {
	SendingHourAt   int
	SendingMinuteAt int
	DaemonMode      bool
}

type OpenAI struct {
	Token     string
	WithImage bool
}

type Slack struct {
	Token   string
	Channel string
}

func (loader *Loader) createConfig() *Config {
	return &Config{
		Common{
			DaemonMode: utils.GetEnvBool(DaemonMode, false),
		},
		OpenAI{
			Token:     utils.GetEnvStr(OpenAIToken, ""),
			WithImage: utils.GetEnvBool(WithImage, true),
		},
		Slack{
			Token:   utils.GetEnvStr(SlackToken, ""),
			Channel: utils.GetEnvStr(SlackChannel, ""),
		},
	}
}

func (loader *Loader) LoadConfig() (*Config, error) {
	var cfg = loader.createConfig()

	sendingHourAt, sendingMinuteAt, err := utils.ParseTime(utils.GetEnvStr(SendingHourAt, "12:00"))

	if err != nil {
		return nil, fmt.Errorf("invalid format for env variable SENDING_HOUR_AT: %w", err)
	}

	cfg.SendingHourAt = sendingHourAt
	cfg.SendingMinuteAt = sendingMinuteAt

	if err := loader.validate(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (loader *Loader) validate(cfg *Config) error {
	if cfg.OpenAI.Token == "" {
		return loader.createNotNullEnvError(OpenAIToken)
	}

	if cfg.Slack.Token == "" {
		return loader.createNotNullEnvError(SlackToken)
	}

	if cfg.Slack.Channel == "" {
		return loader.createNotNullEnvError(SlackToken)
	}

	return nil
}

func (loader *Loader) createNotNullEnvError(envName string) error {
	return fmt.Errorf("env variable %v cannot be null", envName)
}
