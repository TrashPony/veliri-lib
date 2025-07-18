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

func (am *AccessManager) CheckAccess(corporationID, playerID int, fraction string) bool {
	// Если Access == true, доступ разрешён по умолчанию, если не запрещён явно
	// Если Access == false, доступ запрещён по умолчанию, если не разрешён явно
	return (!am.Access || am.GetAccess(
		"corporation"+strconv.Itoa(corporationID),
		"player"+strconv.Itoa(playerID),
		"fraction"+fraction,
	))
}

func (am *AccessManager) GetAccess(corporationKey, playerKey, fractionKey string) bool {
	am.mx.Lock()
	defer am.mx.Unlock()

	if am.access == nil {
		return false
	}

	// 1. Проверяем playerKey (высший приоритет)
	playerAccess, playerExists := am.access[playerKey]
	if playerExists {
		if am.AccessInvers {
			return !playerAccess // инвертируем, если AccessInvers = true
		}
		return playerAccess
	}

	// 2. Проверяем corporationKey (средний приоритет)
	corporationAccess, corporationExists := am.access[corporationKey]
	if corporationExists {
		if am.AccessInvers {
			return !corporationAccess
		}
		return corporationAccess
	}

	// 3. Проверяем fractionKey (низший приоритет)
	fractionAccess, fractionExists := am.access[fractionKey]
	if fractionExists {
		if am.AccessInvers {
			return !fractionAccess
		}
		return fractionAccess
	}

	// Если ни один ключ не найден → возвращаем false (или true, если AccessInvers)
	return am.AccessInvers
}
