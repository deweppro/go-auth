package isp

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Google struct {
	oauth  *oauth2.Config
	config cfg
}

func (v Google) Name() string {
	return "google"
}

func (v *Google) Config(c Config) {
	v.oauth = &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		RedirectURL:  c.RedirectURL,
		Endpoint:     google.Endpoint,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}
	v.config = cfg{
		State:       "state",
		AuthCodeKey: "code",
		RequestURL:  "https://openidconnect.googleapis.com/v1/userinfo",
	}
}

func (v *Google) AuthCodeURL() string {
	return v.oauth.AuthCodeURL(v.config.State)
}

func (v *Google) AuthCodeKey() string {
	return v.config.AuthCodeKey
}

func (v *Google) Exchange(code string) ([]byte, error) {
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
