package google

//go:generate easyjson

import (
	json "encoding/json"

	"github.com/deweppro/go-auth/providers/isp"
)

var _ isp.IUser = (*User)(nil)

//easyjson:json
type model struct {
	Name          string `json:"name"`
	Icon          string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
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

	if tmp.EmailVerified {
		v.name = tmp.Name
		v.icon = tmp.Icon
		v.email = tmp.Email
	}

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
