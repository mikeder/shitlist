package database

import "sync"

type VolatileClickStore struct {
	clickers map[string]int64
	m        *sync.Mutex
}

func NewVolatileClickStore() *VolatileClickStore {
	return &VolatileClickStore{
		clickers: make(map[string]int64),
		m:        new(sync.Mutex),
	}
}

func (v *VolatileClickStore) AddClick(userID string) (Clicker, error) {
	v.m.Lock()
	v.clickers[userID]++
	c := Clicker{
		UserID:     userID,
		ClickCount: v.clickers[userID],
	}
	v.m.Unlock()
	return c, nil
}

func (v *VolatileClickStore) GetClickers() ([]Clicker, error) {
	var c []Clicker
	for k, v := range v.clickers { // Reads without locking because volatile db users don't care
		c = append(c, Clicker{
			UserID:     k,
			ClickCount: v,
		})
	}
	return c, nil
}
