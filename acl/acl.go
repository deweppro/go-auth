package acl

import (
	"sync"

	"github.com/deweppro/go-auth/storage"
	"github.com/pkg/errors"
)

type IAcl interface {
	GetAll(email string) []uint8
	Get(email string, feature uint16) uint8
	Set(email string, feature uint16, level uint8) error
	ResetCache(email string)
}

type ACL struct {
	size  uint
	store storage.IStorage
	cache map[string][]uint8

	mu sync.RWMutex
}

func New(store storage.IStorage, size uint) IAcl {
	return &ACL{
		store: store,
		size:  size,
		cache: make(map[string][]uint8),
	}
}

func (v *ACL) GetAll(email string) []uint8 {
	if !v.hasInCache(email) {
		v.loadToCache(email)
	}

	return v.getAllFromCache(email)
}

func (v *ACL) Get(email string, feature uint16) uint8 {
	if !v.hasInCache(email) {
		v.loadToCache(email)
	}

	if level, ok := v.getFromCache(email, feature); ok {
		return level
	}

	return 0
}

func (v *ACL) Set(email string, feature uint16, level uint8) error {
	if !v.hasInCache(email) {
		v.loadToCache(email)
	}

	v.setToCache(email, feature, level)
	if err := v.store.ChangeACL(email, UintsToString(v.getAllFromCache(email)...)); err != nil {
		return errors.Wrap(err, "change acl")
	}

	return nil
}

func (v *ACL) ResetCache(email string) {
	v.mu.Lock()
	defer v.mu.Unlock()

	delete(v.cache, email)
}

func (v *ACL) loadToCache(email string) {
	v.mu.Lock()
	defer v.mu.Unlock()

	if access, ok := v.store.FindACL(email); ok {
		v.saveToCache(email, StringToUints(access)...)
	} else {
		v.saveToCache(email, make([]uint8, v.size)...)
	}
}

func (v *ACL) hasInCache(email string) bool {
	v.mu.RLock()
	defer v.mu.RUnlock()

	_, ok := v.cache[email]

	return ok
}

func (v *ACL) getFromCache(email string, feature uint16) (uint8, bool) {
	v.mu.RLock()
	defer v.mu.RUnlock()

	if access, ok := v.cache[email]; ok {
		return access[feature], true
	}

	return 0, false
}

func (v *ACL) getAllFromCache(email string) []uint8 {
	v.mu.RLock()
	defer v.mu.RUnlock()

	if access, ok := v.cache[email]; ok {
		return access
	}
	return make([]uint8, v.size)
}

func (v *ACL) setToCache(email string, feature uint16, level uint8) {
	v.mu.Lock()
	defer v.mu.Unlock()

	if _, ok := v.cache[email]; !ok {
		v.cache[email] = make([]uint8, v.size)
	}

	v.cache[email][feature] = level
}

func (v *ACL) saveToCache(email string, access ...uint8) {
	v.mu.Lock()
	defer v.mu.Unlock()

	if len(access) < int(v.size) {
		access = append(access, make([]uint8, int(v.size)-len(access))...)
	}

	v.cache[email] = access[:v.size]
}
