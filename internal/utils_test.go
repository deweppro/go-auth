package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringToUints(t *testing.T) {
	v := StringToUints("123456789")

	require.Equal(t, []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}, v)
}

func TestUintsToString(t *testing.T) {
	v := UintsToString([]uint8{0, 1, 5, 3, 8, 9, 10}...)

	require.Equal(t, "0153899", v)
}
