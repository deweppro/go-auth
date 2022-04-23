package auth

import (
	"net/http"

	"github.com/deweppro/go-auth/providers"
	"github.com/deweppro/go-auth/providers/isp"
)

type HttpHandler func(http.ResponseWriter, *http.Request)

type Auth struct {
	providers providers.IProviders
}

func New(p providers.IProviders) *Auth {
	return &Auth{
		providers: p,
	}
}

func (v *Auth) Request(name string) HttpHandler {
	p, err := v.providers.Get(name)
	if err != nil {
		return func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error())) //nolint: errcheck
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, p.AuthCodeURL(), http.StatusMovedPermanently)
	}
}

func (v *Auth) CallBack(name string, cb func(http.ResponseWriter, *http.Request, isp.IUser)) HttpHandler {
	p, err := v.providers.Get(name)
	if err != nil {
		return func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error())) //nolint: errcheck
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get(p.AuthCodeKey())
		u, err := p.Exchange(code)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error())) //nolint: errcheck
			return
		}
		cb(w, r, u)
	}
}
