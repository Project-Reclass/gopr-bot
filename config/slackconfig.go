package config

// SlackConfig ...
type SlackConfig struct {
	Token     string `yaml:"token"`
	Channel   string `yaml:"channel"`
	IconEmoji string `yaml:"icon_emoji"`
	URI       string `yaml:"uri"`
}
