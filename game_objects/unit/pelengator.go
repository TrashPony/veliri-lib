package unit

import (
	"github.com/TrashPony/veliri-lib/game_objects/visible_anomaly"
)

type Pelengator struct {
	pings         []*visible_anomaly.VisibleAnomaly
	geoScanRadius int
	pelengRadius  int
	pelengMods    []string
	on            bool
}

func (u *Unit) AddPing(ping *visible_anomaly.VisibleAnomaly) {
	u.mx.Lock()
	defer u.mx.Unlock()
	u.checkPelengator()
	u.pelengator.pings = append(u.pelengator.pings, ping)
}

func (u *Unit) FlushPings() []*visible_anomaly.VisibleAnomaly {
	u.mx.Lock()
	defer u.mx.Unlock()
	u.checkPelengator()

	// Сохраняем ссылку на текущий слайс
	pings := u.pelengator.pings

	// Очищаем буфер, но сохраняем capacity
	u.pelengator.pings = u.pelengator.pings[:0]

	return pings
}

func (u *Unit) SetPelengatorState(geoScanRadius, pelengRadius int, pelengMods []string) {
	u.mx.Lock()
	defer u.mx.Unlock()
	u.checkPelengator()

	u.pelengator.geoScanRadius = geoScanRadius
	u.pelengator.pelengRadius = pelengRadius
	u.pelengator.pelengMods = pelengMods
}

func (u *Unit) GetPelengatorOn() bool {
	u.mx.Lock()
	defer u.mx.Unlock()
	u.checkPelengator()

	return u.pelengator.on
}

func (u *Unit) SetPelengatorOn(on bool) {
	u.mx.Lock()
	defer u.mx.Unlock()
	u.checkPelengator()

	u.pelengator.on = on
}

func (u *Unit) GetPelengatorState() (int, int, []string) {
	u.mx.Lock()
	defer u.mx.Unlock()
	u.checkPelengator()
	return u.pelengator.geoScanRadius, u.pelengator.pelengRadius, u.pelengator.pelengMods
}

func (u *Unit) checkPelengator() {
	if u.pelengator.pings == nil {
		u.pelengator.pings = make([]*visible_anomaly.VisibleAnomaly, 0, 16)
	}
}
