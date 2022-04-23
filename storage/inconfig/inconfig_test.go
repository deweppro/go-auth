package inconfig_test

import (
	"testing"

	"github.com/deweppro/go-auth/storage/inconfig"
	"github.com/stretchr/testify/require"
)

func TestNewInConfig(t *testing.T) {
	conf := &inconfig.Config{
		ACL: map[string]string{
			"test": "000000",
		},
	}
	store := inconfig.New(conf)

	v, err := store.FindACL("demo")
	require.Error(t, err)
	require.Equal(t, "", v)

	v, err = store.FindACL("test")
	require.NoError(t, err)
	require.Equal(t, "000000", v)
}
