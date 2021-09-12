package storage

type ConfigStorage struct {
	config *Config
}

func NewConfigStorage(c *Config) *ConfigStorage {
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
