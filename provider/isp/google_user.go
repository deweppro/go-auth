package isp

import json "encoding/json"

var _ IUser = (*UserGoogle)(nil)

//go:generate easyjson
//easyjson:json
type tempUserGoogle struct {
	Name          string `json:"name"`
	Icon          string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

type UserGoogle struct {
	Name  string
	Icon  string
	Email string
}

func (v *UserGoogle) UnmarshalJSON(data []byte) error {
	var tmp tempUserGoogle
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	if tmp.EmailVerified {
		v.Name = tmp.Name
		v.Icon = tmp.Icon
		v.Email = tmp.Email
	}

	return nil
}

func (v *UserGoogle) GetName() string {
	return v.Name
}

func (v *UserGoogle) GetIcon() string {
	return v.Icon
}

func (v *UserGoogle) GetEmail() string {
	return v.Email
}
