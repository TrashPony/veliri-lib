package physical_model

import (
	"github.com/TrashPony/veliri-lib/timecache"
)

type WASD struct {
	w      byte
	a      byte
	s      byte
	d      byte
	q      bool
	e      bool
	z      bool
	sp     bool
	st     bool
	update int64
}

func (wasd *WASD) SetAllFalse() {
	wasd.w = 0
	wasd.a = 0
	wasd.s = 0
	wasd.d = 0
}

func (wasd *WASD) Set(w, a, s, d byte, sp, st, z, q, e bool) {

	wasd.w = w
	wasd.a = a
	wasd.s = s
	wasd.d = d
	wasd.sp = sp
	wasd.st = st
	wasd.z = z
	wasd.q = q
	wasd.e = e

	wasd.update = timecache.GetTimer().UnixNano()
}

func (wasd *WASD) GetW() byte {
	return wasd.w
}

func (wasd *WASD) GetA() byte {
	return wasd.a
}

func (wasd *WASD) GetS() byte {
	return wasd.s
}

func (wasd *WASD) GetD() byte {
	return wasd.d
}

func (wasd *WASD) GetWBool() bool {
	return wasd.w > 0
}

func (wasd *WASD) GetABool() bool {
	return wasd.a > 0
}

func (wasd *WASD) GetSBool() bool {
	return wasd.s > 0
}

func (wasd *WASD) GetDBool() bool {
	return wasd.d > 0
}

func (wasd *WASD) GetQ() bool {
	return wasd.q
}

func (wasd *WASD) GetE() bool {
	return wasd.e
}

func (wasd *WASD) GetZ() bool {
	return wasd.z
}

func (wasd *WASD) GetSpace() bool {
	return wasd.sp
}

func (wasd *WASD) GetShift() bool {
	return wasd.st
}

func (wasd *WASD) GetUpdate() int64 {
	return wasd.update
}
