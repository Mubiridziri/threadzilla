package config

import "fmt"
import "threadzilla/internal/utils"

const OpenAIToken = "OPENAI_TOKEN"
const SlackToken = "SLACK_TOKEN"
const SendingHourAt = "SENDING_HOUR_AT"

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
}

type OpenAI struct {
	Token string
}

type Slack struct {
	Token string
}

func (loader *Loader) createConfig() *Config {
	return &Config{
		Common{},
		OpenAI{
			Token: utils.GetEnvStr(OpenAIToken, ""),
		},
		Slack{
			Token: utils.GetEnvStr(SlackToken, ""),
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

	return nil
}

func (loader *Loader) createNotNullEnvError(envName string) error {
	return fmt.Errorf("env variable %v cannot be null", envName)
}
