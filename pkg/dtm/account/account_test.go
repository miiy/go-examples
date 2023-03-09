package account

import (
	"fmt"
	"testing"
)

func TestSyncDict_Put(t *testing.T) {
	m := SyncDict{}
	a1 := Account{
		Uid:     1,
		Balance: 10,
	}
	a2 := Account{
		Uid:     2,
		Balance: 20,
	}
	m.Put(1, a1)
	m.Put(2, a2)
	fmt.Println(m.Get(1))
}
