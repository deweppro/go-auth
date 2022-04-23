package acl_test

import (
	"testing"

	"github.com/deweppro/go-auth/acl"
	"github.com/deweppro/go-auth/storage/inmemory"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {

	store := inmemory.New()
	acls := acl.New(store, 3)

	email := "test@dewep.pro"

	t.Log("user not exist")

	levels, err := acls.GetAll(email)
	require.Error(t, err)
	require.Nil(t, levels)

	require.Error(t, acls.Set(email, 10, 1))

	t.Log("user exist")

	require.NoError(t, store.ChangeACL(email, ""))

	require.Error(t, acls.Set(email, 10, 1))

	levels, err = acls.GetAll(email)
	require.NoError(t, err)
	require.Equal(t, []uint8{0, 0, 0}, levels)

	require.NoError(t, acls.Set(email, 2, 10))

	levels, err = acls.GetAll(email)
	require.NoError(t, err)
	require.Equal(t, []uint8{0, 0, 9}, levels)
}
