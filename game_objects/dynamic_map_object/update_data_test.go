package dynamic_map_object

import (
	"bytes"
	"testing"
)

// GetUpdateData кэшируется по тику: внутри тика один и тот же срез, на новом тике - свежие данные.
func TestGetUpdateDataCachedPerTick(t *testing.T) {
	o := &Object{ID: 1, HP: 100, Work: true, Complete: 100}

	tick := int64(32)
	a := o.GetUpdateData(tick)
	b := o.GetUpdateData(tick)
	if &a[0] != &b[0] {
		t.Fatal("внутри тика должен возвращаться один и тот же срез (без аллокации на каждый вызов)")
	}

	o.HP = 40
	if !bytes.Equal(o.GetUpdateData(tick), a) {
		t.Fatal("внутри тика данные не пересчитываются")
	}

	tick += 32
	c := o.GetUpdateData(tick)
	if bytes.Equal(c, a) {
		t.Fatal("на новом тике HP=40 обязан попасть в данные")
	}

	// эквивалентность прежнему расчёту: свежий объект с теми же полями на том же тике даёт те же байты
	fresh := &Object{ID: 1, HP: 40, Work: true, Complete: 100}
	if !bytes.Equal(fresh.GetUpdateData(tick), c) {
		t.Fatal("кэшированные данные отличаются от прямого расчёта")
	}
}
