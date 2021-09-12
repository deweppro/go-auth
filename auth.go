package auth

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/deweppro/go-auth/provider"
	"github.com/deweppro/go-auth/storage"
)

type HttpHandler func(http.ResponseWriter, *http.Request)

type Auth struct {
	storage   storage.IStorage
	providers provider.IProviders
}

func New(s storage.IStorage, p provider.IProviders) *Auth {
	return &Auth{
		storage:   s,
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

func (v *Auth) CallBack(name string, cb func([]byte, http.ResponseWriter)) HttpHandler {
	p, err := v.providers.Get(name)
	if err != nil {
		return func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error())) //nolint: errcheck
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get(p.AuthCodeKey())
		data, err := p.Exchange(code)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error())) //nolint: errcheck
			return
		}
		cb(data, w)
	}
}

func (v *Auth) CallBackWithACL(name string, model IUser, cb func(IUser, http.ResponseWriter)) HttpHandler {
	p, err := v.providers.Get(name)
	if err != nil {
		return func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error())) //nolint: errcheck
		}
	}

	ref := reflect.TypeOf(model).Elem()

	return func(w http.ResponseWriter, r *http.Request) {
		data, err := p.Exchange(r.URL.Query().Get(p.AuthCodeKey()))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error())) //nolint: errcheck
			return
		}

		user, ok := reflect.New(ref).Interface().(IUser)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error())) //nolint: errcheck
			return
		}

		err = json.Unmarshal(data, user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error())) //nolint: errcheck
			return
		}

		acl, ok := v.storage.FindACL(user.GetEmail())
		if !ok {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		user.SetACL(acl)

		cb(user, w)
	}
}
