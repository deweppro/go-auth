package acl_test

import (
	"testing"

	"github.com/deweppro/go-auth/acl"
	"github.com/stretchr/testify/require"
)

func TestUnit_StringToUints(t *testing.T) {
	v := acl.StringToUints("123456789")

	require.Equal(t, []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}, v)
}

func TestUnit_UintsToString(t *testing.T) {
	v := acl.UintsToString([]uint8{0, 1, 5, 3, 8, 9, 10}...)

	require.Equal(t, "0153899", v)
}
