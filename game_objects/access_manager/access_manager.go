package access_manager

import (
	"strconv"
	"sync"
)

type AccessManager struct {
	Access       bool
	AccessInvers bool // если true, то запрещает всем из мапы access
	access       map[string]bool
	mx           sync.Mutex
}

func (am *AccessManager) AddAccessByKey(key string) {
	am.mx.Lock()
	defer am.mx.Unlock()

	if am.access == nil {
		am.access = map[string]bool{}
	}

	am.access[key] = true
}

func (am *AccessManager) AddAccess(typeAccess string, id int) {
	am.mx.Lock()
	defer am.mx.Unlock()

	if am.access == nil {
		am.access = map[string]bool{}
	}

	am.access[typeAccess+strconv.Itoa(id)] = true
}

func (am *AccessManager) RemoveAccess(typeAccess string, id int) {
	am.mx.Lock()
	defer am.mx.Unlock()

	if am.access != nil {
		delete(am.access, typeAccess+strconv.Itoa(id))
	}
}

func (am *AccessManager) GetAccess(corporationKey, playerKey string) bool {
	am.mx.Lock()
	defer am.mx.Unlock()

	if am.access == nil {
		return false
	}

	corporationAccess, _ := am.access[corporationKey]
	if am.AccessInvers {
		corporationAccess = !corporationAccess
	}

	playerAccess, ok := am.access[playerKey]
	if !ok {
		return corporationAccess
	}

	if am.AccessInvers {
		playerAccess = !playerAccess
	}

	return playerAccess
}
