package internal

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/deweppro/go-errors"
	"golang.org/x/oauth2"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const MAX_LEVEL = uint8(9)

func ValidateLevel(v uint8) uint8 {
	if v > MAX_LEVEL {
		return MAX_LEVEL
	}
	return v
}

func StringToUints(data string) []uint8 {
	t := make([]uint8, len(data))
	for i, s := range strings.Split(data, "") {
		v, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			t[i] = 0
			continue
		}
		b := uint8(v)
		if b > MAX_LEVEL {
			t[i] = 9
		} else {
			t[i] = uint8(b)
		}
	}
	return t
}

func UintsToString(data ...uint8) string {
	t := ""
	for _, v := range data {
		if v > MAX_LEVEL {
			v = MAX_LEVEL
		}
		t += strconv.FormatUint(uint64(v), 10)
	}
	return t
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ReadBody(resp *http.Response) ([]byte, error) {
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return data, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type oauth interface {
	Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
	Client(ctx context.Context, t *oauth2.Token) *http.Client
}

func Exchange(code string, uri string, srv oauth, model json.Unmarshaler) error {
	tok, err := srv.Exchange(context.Background(), code)
	if err != nil {
		return errors.WrapMessage(err, "exchange to oauth service")
	}
	client := srv.Client(context.Background(), tok)
	resp, err := client.Get(uri)
	if err != nil {
		return errors.WrapMessage(err, "client request to oauth service")
	}
	b, err := ReadBody(resp)
	if err != nil {
		return errors.WrapMessage(err, "read response from oauth service")
	}
	if err = json.Unmarshal(b, model); err != nil {
		return errors.WrapMessage(err, "decode oauth model")
	}
	return nil
}
