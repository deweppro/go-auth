package providers

import (
	"sync"

	"github.com/deweppro/go-auth/config"
	"github.com/deweppro/go-auth/providers/isp"
	"github.com/deweppro/go-auth/providers/isp/google"
	"github.com/deweppro/go-auth/providers/isp/yandex"
	"github.com/deweppro/go-errors"
)

var (
	ErrProviderFail = errors.New("provider not found")
)

type (
	IProvider interface {
		Code() string
		Config(config.ConfigItem)
		AuthCodeURL() string
		AuthCodeKey() string
		Exchange(string) (isp.IUser, error)
	}

	IProviders interface {
		Get(string) (IProvider, error)
	}

	Providers struct {
		config *config.Config
		list   map[string]IProvider
		l      sync.RWMutex
	}
)

func New(c *config.Config) *Providers {
	p := &Providers{
		config: c,
		list:   make(map[string]IProvider),
	}

	p.Add(
		&google.Provider{},
		&yandex.Provider{},
	)

	return p
}

func (v *Providers) Add(p ...IProvider) {
	v.l.Lock()
	defer v.l.Unlock()

	for _, item := range p {
		for _, cp := range v.config.Provider {
			if cp.Code == item.Code() {
				item.Config(cp)
				v.list[item.Code()] = item
			}
		}
	}
}

func (v *Providers) Get(name string) (IProvider, error) {
	v.l.RLock()
	defer v.l.RUnlock()

	p, ok := v.list[name]
	if !ok {
		return nil, ErrProviderFail
	}
	return p, nil
}
