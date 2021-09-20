package isp

import (
	json "encoding/json"
	"fmt"
)

var _ IUser = (*UserYandex)(nil)

//go:generate easyjson

//easyjson:json
type tempUserYandex struct {
	Name  string `json:"display_name"`
	Icon  string `json:"default_avatar_id"`
	Email string `json:"default_email"`
}

type UserYandex struct {
	Name  string
	Icon  string
	Email string
}

func (v *UserYandex) UnmarshalJSON(data []byte) error {
	var tmp tempUserYandex
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	if len(tmp.Icon) > 0 {
		v.Icon = fmt.Sprintf("https://avatars.yandex.net/get-yapic/%s/islands-retina-50", tmp.Icon)
	}
	v.Name = tmp.Name
	v.Email = tmp.Email

	return nil
}

func (v *UserYandex) GetName() string {
	return v.Name
}

func (v *UserYandex) GetIcon() string {
	return v.Icon
}

func (v *UserYandex) GetEmail() string {
	return v.Email
}
