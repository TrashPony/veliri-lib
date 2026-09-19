package drone

import "sync/atomic"

// droneViewRange - закэшированная дальность обзора дрона для набора эффектов version и базовой дальности base.
type droneViewRange struct {
	version uint64
	base    int
	value   int
}

type droneViewRangeCache = atomic.Pointer[droneViewRange]
