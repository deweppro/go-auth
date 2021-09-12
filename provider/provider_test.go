package provider_test

import (
	"testing"

	"github.com/deweppro/go-auth/provider"
	"github.com/deweppro/go-auth/provider/isp"
	"github.com/stretchr/testify/require"
)

func TestUnit_New(t *testing.T) {
	conf := provider.Config{
		Provider: []isp.Config{
			{
				Name:         "google",
				ClientID:     "123",
				ClientSecret: "456",
				RedirectURL:  "https://example.com",
			},
		},
	}
	prov := provider.New(&conf)

	v, err := prov.Get("demo")
	require.Error(t, err)
	require.Nil(t, v)

	v, err = prov.Get("google")
	require.NoError(t, err)
	require.Implements(t, (*provider.IProvider)(nil), v)

	require.Equal(t, "code", v.AuthCodeKey())
	require.Equal(t,
		"https://accounts.google.com/o/oauth2/auth?client_id=123&redirect_uri=https%3A%2F%2Fexample.com&response_type=code&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile&state=state",
		v.AuthCodeURL())
}

func TestUnit_NewEmpty(t *testing.T) {
	conf := provider.Config{
		Provider: []isp.Config{},
	}
	prov := provider.New(&conf)

	v, err := prov.Get("demo")
	require.Error(t, err)
	require.Nil(t, v)

	v, err = prov.Get("google")
	require.Error(t, err)
	require.Nil(t, v)
}
