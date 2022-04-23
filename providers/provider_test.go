package providers_test

import (
	"testing"

	"github.com/deweppro/go-auth/config"
	"github.com/deweppro/go-auth/providers"
	"github.com/stretchr/testify/require"
)

func TestUnit_New(t *testing.T) {
	conf := config.Config{
		Provider: []config.ConfigItem{
			{
				Code:         "google",
				ClientID:     "123",
				ClientSecret: "456",
				RedirectURL:  "https://example.com",
			},
		},
	}
	prov := providers.New(&conf)

	v, err := prov.Get("demo")
	require.Error(t, err)
	require.Nil(t, v)

	v, err = prov.Get("google")
	require.NoError(t, err)
	require.Implements(t, (*providers.IProvider)(nil), v)

	require.Equal(t, "code", v.AuthCodeKey())
	require.Equal(t,
		"https://accounts.google.com/o/oauth2/auth?client_id=123&redirect_uri=https%3A%2F%2Fexample.com&response_type=code&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile&state=state",
		v.AuthCodeURL())
}

func TestUnit_NewEmpty(t *testing.T) {
	conf := config.Config{
		Provider: []config.ConfigItem{},
	}
	prov := providers.New(&conf)

	v, err := prov.Get("demo")
	require.Error(t, err)
	require.Nil(t, v)

	v, err = prov.Get("google")
	require.Error(t, err)
	require.Nil(t, v)
}
