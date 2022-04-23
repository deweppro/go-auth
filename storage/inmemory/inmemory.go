package inmemory

import (
	"sync"

	"github.com/deweppro/go-auth/storage"
	"github.com/deweppro/go-errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type InMemory struct {
	data map[string]string
	mux  sync.Mutex
}

func New() storage.IStorage {
	return &InMemory{
		data: map[string]string{},
	}
}

func (v *InMemory) FindACL(email string) (string, error) {
	v.mux.Lock()
	defer v.mux.Unlock()

	if acl, ok := v.data[email]; ok {
		return acl, nil
	}
	return "", ErrUserNotFound
}

func (v *InMemory) ChangeACL(email, data string) error {
	v.mux.Lock()
	defer v.mux.Unlock()

	v.data[email] = data
	return nil
}
