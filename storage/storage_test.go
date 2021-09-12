package storage_test

import (
	"testing"

	"github.com/deweppro/go-auth/storage"
	"github.com/stretchr/testify/require"
)

func TestUnit_NewConfigStorage(t *testing.T) {
	conf := storage.Config{
		ACL: map[string]string{
			"test": "000000",
		},
	}
	store := storage.NewConfigStorage(&conf)

	v, ok := store.FindACL("demo")
	require.False(t, ok)
	require.Equal(t, "", v)

	v, ok = store.FindACL("test")
	require.True(t, ok)
	require.Equal(t, "000000", v)
}
