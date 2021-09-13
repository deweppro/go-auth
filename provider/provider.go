package provider

import (
	"sync"

	"github.com/deweppro/go-auth/provider/isp"
	"github.com/pkg/errors"
)

var (
	ErrProviderFail = errors.New("provider not found")
)

type (
	IProvider interface {
		Name() string
		Config(isp.Config)
		AuthCodeURL() string
		AuthCodeKey() string
		Exchange(code string) ([]byte, error)
	}

	IProviders interface {
		Get(name string) (IProvider, error)
	}

	Providers struct {
		config *Config
		list   map[string]IProvider
		l      sync.RWMutex
	}
)

func New(c *Config) *Providers {
	p := &Providers{
		config: c,
		list:   make(map[string]IProvider),
	}

	p.Add(&isp.Google{}, &isp.Yandex{})

	return p
}

func (v *Providers) Add(p ...IProvider) {
	v.l.Lock()
	defer v.l.Unlock()

	for _, item := range p {
		for _, cp := range v.config.Provider {
			if cp.Name == item.Name() {
				item.Config(cp)
				v.list[item.Name()] = item
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
