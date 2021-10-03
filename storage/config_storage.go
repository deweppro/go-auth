package storage

import "github.com/pkg/errors"

var (
	ErrChangeNotSupported = errors.New("changing ACL is not supported")
)

type ConfigStorage struct {
	config *Config
}

func NewConfigStorage(c *Config) IStorage {
	return &ConfigStorage{
		config: c,
	}
}

func (v *ConfigStorage) FindACL(email string) (string, bool) {
	if acl, ok := v.config.ACL[email]; ok {
		return acl, true
	}
	return "", false
}

func (v *ConfigStorage) ChangeACL(_, _ string) error {
	return ErrChangeNotSupported
}
