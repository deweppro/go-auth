package isp

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/yandex"
)

type Yandex struct {
	oauth  *oauth2.Config
	config cfg
}

func (v Yandex) Name() string {
	return "yandex"
}

func (v *Yandex) Config(c Config) {
	v.oauth = &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		RedirectURL:  c.RedirectURL,
		Endpoint:     yandex.Endpoint,
		Scopes: []string{
			"login:email",
			"login:info",
			"login:avatar",
		},
	}
	v.config = cfg{
		State:       "state",
		AuthCodeKey: "code",
		RequestURL:  "https://login.yandex.ru/info",
	}
}

func (v *Yandex) AuthCodeURL() string {
	return v.oauth.AuthCodeURL(v.config.State)
}

func (v *Yandex) AuthCodeKey() string {
	return v.config.AuthCodeKey
}

func (v *Yandex) Exchange(code string) ([]byte, error) {
	tok, err := v.oauth.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}
	client := v.oauth.Client(context.Background(), tok)
	resp, err := client.Get(v.config.RequestURL)
	if err != nil {
		return nil, err
	}
	return readBody(resp)
}
