package account

import (
	"sync"
)

type SyncDict struct {
	m sync.Map
}

type Account struct {
	Uid     int     `json:"uid"`
	Balance float64 `json:"balance"`
}

func (d *SyncDict) Get(uid int) *Account {
	account, ok := d.m.Load(uid)
	if ok {
		item := account.(Account)
		return &item
	}
	return nil
}

func (d *SyncDict) Put(key int, a Account) {
	d.m.Store(key, a)
}

func (d *SyncDict) Remove(key int) {
	d.m.Delete(key)
}

func (a *Account) AddBalance(amount float64) {
	a.Balance += amount
}

func (a *Account) SubBalance(amount float64) {
	a.Balance -= amount
}
