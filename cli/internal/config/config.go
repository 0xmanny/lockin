package config

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	BlockedWebsites []string `yaml:"blocked_websites"`
	BlockedApps     []string `yaml:"blocked_apps"`
	Settings        Settings `yaml:"settings"`
}

type Settings struct {
	PollInterval time.Duration `yaml:"poll_interval"`
}

func Dir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".lockin"), nil
}

func FilePath() (string, error) {
	dir, err := Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.yaml"), nil
}

func PidPath() (string, error) {
	dir, err := Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "lockin.pid"), nil
}

// true if the config file has already been created
func Exists() bool {
	p, err := FilePath()
	if err != nil {
		return false
	}
	_, err = os.Stat(p)
	return err == nil
}

func Load() (*Config, error) {
	dir, err := Dir()
	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	cfgPath := filepath.Join(dir, "config.yaml")

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		cfg := Defaults()
		if err := Save(cfg); err != nil {
			return nil, err
		}
		return cfg, nil
	}

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	if cfg.Settings.PollInterval == 0 {
		cfg.Settings.PollInterval = 3 * time.Second
	}

	return &cfg, nil
}

func Save(cfg *Config) error {
	cfgPath, err := FilePath()
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(cfgPath, data, 0644)
}

var KnownWebsites = []string{
	"twitter.com",
	"x.com",
	"facebook.com",
	"instagram.com",
	"tiktok.com",
	"reddit.com",
	"youtube.com",
	"netflix.com",
	"twitch.tv",
	"linkedin.com",
	"pinterest.com",
	"tumblr.com",
	"hulu.com",
	"threads.net",
	"bsky.app",
	"truthsocial.com",
}

var KnownApps = []string{
	"Messages",
	"Telegram",
	"Discord",
	"WhatsApp",
	"Slack",
	"Messenger",
	"Signal",
	"Spotify",
	"Music",
	"TV",
	"FaceTime",
}

func Defaults() *Config {
	return &Config{
		BlockedWebsites: []string{
			"twitter.com",
			"x.com",
			"facebook.com",
			"instagram.com",
			"tiktok.com",
			"reddit.com",
			"youtube.com",
			"netflix.com",
			"twitch.tv",
		},
		BlockedApps: []string{
			"Messages",
			"Telegram",
			"Discord",
			"WhatsApp",
		},
		Settings: Settings{
			PollInterval: 3 * time.Second,
		},
	}
}

// adds www. variants
func ExpandWebsites(domains []string) []string {
	expanded := make([]string, 0, len(domains)*2)
	for _, d := range domains {
		expanded = append(expanded, d)
		if len(d) > 0 && !strings.HasPrefix(d, "www.") {
			expanded = append(expanded, "www."+d)
		}
	}
	return expanded
}
