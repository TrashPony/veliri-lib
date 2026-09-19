package dynamic_map_object

import "sync/atomic"

// objectUpdateData - данные обновления объекта, посчитанные для тика time (см. Object.GetUpdateData).
type objectUpdateData struct {
	time int64
	data []byte
}

type objectUpdateDataCache = atomic.Pointer[objectUpdateData]
