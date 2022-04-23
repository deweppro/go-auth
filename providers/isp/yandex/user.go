package yandex

//go:generate easyjson

import (
	json "encoding/json"
	"fmt"

	"github.com/deweppro/go-auth/providers/isp"
)

var _ isp.IUser = (*User)(nil)

//easyjson:json
type model struct {
	Name  string `json:"display_name"`
	Icon  string `json:"default_avatar_id"`
	Email string `json:"default_email"`
}

type User struct {
	name  string
	icon  string
	email string
}

func (v *User) UnmarshalJSON(data []byte) error {
	var tmp model
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	if len(tmp.Icon) > 0 {
		v.icon = fmt.Sprintf("https://avatars.yandex.net/get-yapic/%s/islands-retina-50", tmp.Icon)
	}
	v.name = tmp.Name
	v.email = tmp.Email

	return nil
}

func (v *User) GetName() string {
	return v.name
}

func (v *User) GetIcon() string {
	return v.icon
}

func (v *User) GetEmail() string {
	return v.email
}
