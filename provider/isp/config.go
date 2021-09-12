package isp

//go:generate easyjson

//easyjson:json
type Config struct {
	Name         string `json:"name" yaml:"name"`
	ClientID     string `json:"client_id" yaml:"client_id"`
	ClientSecret string `json:"client_secret" yaml:"client_secret"`
	RedirectURL  string `json:"redirect_url" yaml:"redirect_url"`
}

type cfg struct {
	State       string
	AuthCodeKey string
	RequestURL  string
}
