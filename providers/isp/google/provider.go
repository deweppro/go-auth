package google

import (
	"github.com/deweppro/go-auth/config"
	"github.com/deweppro/go-auth/internal"
	"github.com/deweppro/go-auth/providers/isp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const CODE = "google"

type Provider struct {
	oauth  *oauth2.Config
	config internal.Config
}

func (v Provider) Code() string {
	return CODE
}

func (v *Provider) Config(c config.ConfigItem) {
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
	v.config = internal.Config{
		State:       "state",
		AuthCodeKey: "code",
		RequestURL:  "https://openidconnect.googleapis.com/v1/userinfo",
	}
}

func (v *Provider) AuthCodeURL() string {
	return v.oauth.AuthCodeURL(v.config.State)
}

func (v *Provider) AuthCodeKey() string {
	return v.config.AuthCodeKey
}

func (v *Provider) Exchange(code string) (isp.IUser, error) {
	model := &User{}
	if err := internal.Exchange(code, v.config.RequestURL, v.oauth, model); err != nil {
		return nil, err
	}
	return model, nil
}
