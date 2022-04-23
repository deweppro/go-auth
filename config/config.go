package config

type ConfigItem struct {
	Code         string `yaml:"code"`
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	RedirectURL  string `yaml:"redirect_url"`
}

type Config struct {
	Provider []ConfigItem `yaml:"oauth_providers"`
}
