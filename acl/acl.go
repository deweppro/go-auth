package acl

import (
	"context"
	"time"

	"github.com/deweppro/go-auth/internal"
	"github.com/deweppro/go-auth/storage"
	"github.com/deweppro/go-errors"
)

var (
	ErrChangeNotFound = errors.New("change acl is not found")
)

type IAcl interface {
	GetAll(email string) ([]uint8, error)
	Get(email string, feature uint16) (uint8, error)
	Set(email string, feature uint16, level uint8) error
	Flush(email string)
	AutoFlush(ctx context.Context, interval time.Duration)
}

type ACL struct {
	cache *cache
	store storage.IStorage
}

func New(store storage.IStorage, size uint) IAcl {
	return &ACL{
		store: store,
		cache: newCache(size),
	}
}

func (v *ACL) AutoFlush(ctx context.Context, interval time.Duration) {
	tick := time.NewTicker(interval)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case ts := <-tick.C:
			v.cache.FlushByTime(ts.Unix())
		}
	}
}

func (v *ACL) GetAll(email string) ([]uint8, error) {
	if !v.cache.Has(email) {
		if err := v.loadFromStore(email); err != nil {
			return nil, err
		}
	}

	return v.cache.GetAll(email)
}

func (v *ACL) Get(email string, feature uint16) (uint8, error) {
	if !v.cache.Has(email) {
		if err := v.loadFromStore(email); err != nil {
			return 0, err
		}
	}

	return v.cache.Get(email, feature)
}

func (v *ACL) Set(email string, feature uint16, level uint8) error {
	if !v.cache.Has(email) {
		if err := v.loadFromStore(email); err != nil {
			return err
		}
	}

	if err := v.cache.Set(email, feature, level); err != nil {
		return err
	}
	return v.saveToStore(email)
}

func (v *ACL) Flush(email string) {
	v.cache.Flush(email)
}

func (v *ACL) loadFromStore(email string) error {
	access, err := v.store.FindACL(email)
	if err != nil {
		return errors.Wrap(err, ErrUserNotFound)
	}
	v.cache.SetAll(email, internal.StringToUints(access)...)
	return nil
}

func (v *ACL) saveToStore(email string) error {
	access, err := v.cache.GetAll(email)
	if err != nil {
		return err
	}

	err = v.store.ChangeACL(email, internal.UintsToString(access...))
	return errors.WrapMessage(err, "change acl")
}
