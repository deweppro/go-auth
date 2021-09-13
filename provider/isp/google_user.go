package isp

import json "encoding/json"

//go:generate easyjson
//easyjson:json
type tempUserGoogle struct {
	ID            string `json:"sub"`
	Name          string `json:"name"`
	Icon          string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

type UserGoogle struct {
	ID    string
	Name  string
	Icon  string
	Email string
	ACL   string
}

func (v *UserGoogle) UnmarshalJSON(data []byte) error {
	var tmp tempUserGoogle
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	if tmp.EmailVerified {
		v.ID = tmp.ID
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

func (v *UserGoogle) GetACL() string {
	return v.ACL
}

func (v *UserGoogle) SetACL(acl string) {
	v.ACL = acl
}
