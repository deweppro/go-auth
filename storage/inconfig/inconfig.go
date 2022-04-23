package inconfig

import (
	"github.com/deweppro/go-auth/storage"
	"github.com/deweppro/go-errors"
)

var (
	ErrChangeNotSupported = errors.New("changing ACL is not supported")
	ErrUserNotFound       = errors.New("user not found")
)

type Config struct {
	ACL map[string]string `yaml:"auth_acl"`
}

type InConfig struct {
	config *Config
}

func New(c *Config) storage.IStorage {
	return &InConfig{
		config: c,
	}
}

func (v *InConfig) FindACL(email string) (string, error) {
	if acl, ok := v.config.ACL[email]; ok {
		return acl, nil
	}
	return "", ErrUserNotFound
}

func (v *InConfig) ChangeACL(_, _ string) error {
	return ErrChangeNotSupported
}
