package provider

import "github.com/deweppro/go-auth/provider/isp"

//go:generate easyjson

//easyjson:json
type Config struct {
	Provider []isp.Config `json:"auth_provider" yaml:"auth_provider"`
}
