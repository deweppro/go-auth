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

func (g Google) Name() string {
	return "google"
}

func (g *Google) Config(c Config) {
	g.oauth = &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		RedirectURL:  c.RedirectURL,
		Endpoint:     google.Endpoint,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}
	g.config = cfg{
		State:       "state",
		AuthCodeKey: "code",
		RequestURL:  "https://openidconnect.googleapis.com/v1/userinfo",
	}
}

func (g *Google) AuthCodeURL() string {
	return g.oauth.AuthCodeURL(g.config.State)
}

func (g *Google) AuthCodeKey() string {
	return g.config.AuthCodeKey
}

func (g *Google) Exchange(code string) ([]byte, error) {
	tok, err := g.oauth.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}
	client := g.oauth.Client(context.Background(), tok)
	resp, err := client.Get(g.config.RequestURL)
	if err != nil {
		return nil, err
	}
	return readBody(resp)
}
