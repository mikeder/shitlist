package database

import (
	"errors"
	"sync"
)

// enforce PersistentDataStore inplements required interfaces
var _ ClickStore = (*VolatileDataStore)(nil)
var _ UserStore = (*VolatileDataStore)(nil)

type VolatileDataStore struct {
	clickers map[string]uint64
	m        *sync.Mutex
}

func NewVolatileClickStore() *VolatileDataStore {
	return &VolatileDataStore{
		clickers: make(map[string]uint64),
		m:        new(sync.Mutex),
	}
}

func (v *VolatileDataStore) AddClick(userID string) (*Clicker, error) {
	v.m.Lock()
	v.clickers[userID]++
	c := &Clicker{
		UserID:     userID,
		ClickCount: v.clickers[userID],
	}
	v.m.Unlock()
	return c, nil
}

func (v *VolatileDataStore) GetClickers() ([]*Clicker, error) {
	var c []*Clicker
	for k, v := range v.clickers { // Reads without locking because volatile db users don't care
		c = append(c, &Clicker{
			UserID:     k,
			ClickCount: v,
		})
	}
	return c, nil
}

func (v *VolatileDataStore) AddUser(name, email string) (*User, error) {
	return nil, errors.New("not implemented")
}

func (v *VolatileDataStore) GetUserByEmail(email string) (*User, error) {
	return nil, errors.New("not implemented")
}

func (v *VolatileDataStore) GetUserByName(name string) (*User, error) {
	return nil, errors.New("not implemented")
}

func (v *VolatileDataStore) GetUserAuthentications(userID string) (*UserAuthentications, error) {
	return nil, errors.New("not implemented")
}
