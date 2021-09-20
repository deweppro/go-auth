package acl

import (
	"strconv"
	"strings"
)

const max = uint8(9)

func StringToUints(data string) []uint8 {
	t := make([]uint8, len(data))
	for i, s := range strings.Split(data, "") {
		v, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			t[i] = 0
			continue
		}
		b := uint8(v)
		if b > max {
			t[i] = 9
		} else {
			t[i] = uint8(b)
		}
	}
	return t
}

func UintsToString(data ...uint8) string {
	t := ""
	for _, b := range data {
		if b > max {
			t += strconv.FormatUint(9, 10)
		} else {
			t += strconv.FormatUint(uint64(b), 10)
		}
	}
	return t
}
