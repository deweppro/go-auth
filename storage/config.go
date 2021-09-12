package storage

//go:generate easyjson

//easyjson:json
type Config struct {
	ACL map[string]string `json:"auth_acl" yaml:"auth_acl"`
}
