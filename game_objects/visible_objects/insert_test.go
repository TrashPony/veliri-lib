package visible_objects

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

func checkSortedStable(t *testing.T, ids []int, order []int) {
	t.Helper()
	for i := 1; i < len(ids); i++ {
		if ids[i-1] > ids[i] {
			t.Fatalf("список не отсортирован: %v", ids)
		}
		if ids[i-1] == ids[i] && order[i-1] > order[i] {
			t.Fatalf("среди равных ID новая метка обязана вставать после старых (порядок вставки): ids=%v order=%v", ids, order)
		}
	}
}

// Бинарная вставка должна давать ровно тот же результат, что append + sort: уникальные ID - тот же порядок, что у
// отсортированного среза; дубликаты - отсортировано, а новая метка после равных.
func TestAddVisibleObjectKeepsSortedOrder(t *testing.T) {
	r := rand.New(rand.NewSource(3))

	for _, distinct := range []int{50, 5000} { // 5000 => почти без дубликатов; 50 => много равных ID
		v := &VisibleObjectsStore{}
		var want []int
		insertion := map[*VisibleObject]int{}

		for i := 0; i < 1500; i++ {
			id := r.Intn(distinct)
			o := &VisibleObject{IDObject: id, TypeObject: "unit"}
			insertion[o] = i
			v.AddVisibleObject(o)
			want = append(want, id)
		}
		sort.Ints(want)

		var got, order []int
		v.RangeVisibleObjects(func(o *VisibleObject) bool {
			got = append(got, o.IDObject)
			order = append(order, insertion[o])
			return true
		})

		if len(got) != len(want) {
			t.Fatalf("distinct=%d: len %d, want %d", distinct, len(got), len(want))
		}
		for i := range want {
			if got[i] != want[i] {
				t.Fatalf("distinct=%d: позиция %d: %d, want %d", distinct, i, got[i], want[i])
			}
		}
		checkSortedStable(t, got, order)
	}
}

func TestAddVisibleObjectGetAndRemove(t *testing.T) {
	v := &VisibleObjectsStore{}
	for _, id := range []int{5, 1, 9, 3, 7} {
		v.AddVisibleObject(&VisibleObject{IDObject: id, TypeObject: "unit"})
		v.AddVisibleObject(&VisibleObject{IDObject: id, TypeObject: "item"}) // тот же id, другой тип
	}

	for _, id := range []int{1, 3, 5, 7, 9} {
		if o := v.GetVisibleObjectByTypeAndID("unit", id); o == nil || o.IDObject != id || o.TypeObject != "unit" {
			t.Fatalf("unit %d не найден: %v", id, o)
		}
	}
	if v.GetVisibleObjectByTypeAndID("unit", 4) != nil {
		t.Fatal("несуществующий id найден")
	}

	v.RemoveVisibleObject(v.GetVisibleObjectByTypeAndID("unit", 5))
	if v.GetVisibleObjectByTypeAndID("unit", 5) != nil || v.GetVisibleObjectByTypeAndID("item", 5) == nil {
		t.Fatal("удаление unit#5 задело не то")
	}
}

func TestAddDynamicObjectKeepsSortedOrder(t *testing.T) {
	r := rand.New(rand.NewSource(4))
	v := &VisibleObjectsStore{}

	var want []int
	for i := 0; i < 1000; i++ {
		id := 1 + r.Intn(100000)
		v.AddDynamicObject(&VisibleObject{ID: id, IDObject: id})
		want = append(want, id)
	}
	sort.Ints(want)

	objs, mx := v.UnsafeRangeMapDynamicObjects()
	mx.RLock()
	defer mx.RUnlock()

	got := objs[idMemoryType]
	if len(got) != len(want) {
		t.Fatalf("len %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i].ID != want[i] {
			t.Fatalf("позиция %d: %d, want %d", i, got[i].ID, want[i])
		}
	}
}

func BenchmarkAddVisibleObject(b *testing.B) {
	for _, n := range []int{50, 500, 2000} {
		b.Run("n"+strconv.Itoa(n), func(b *testing.B) {
			v := &VisibleObjectsStore{}
			for i := 0; i < n; i++ {
				v.AddVisibleObject(&VisibleObject{IDObject: i * 2, TypeObject: "unit"})
			}
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				o := &VisibleObject{IDObject: n*2 + 1, TypeObject: "unit"}
				v.AddVisibleObject(o)
				v.RemoveVisibleObject(o)
			}
		})
	}
}

func TestHashUpdateData(t *testing.T) {
	if HashUpdateData(nil) == 0 || HashUpdateData([]byte{}) == 0 {
		t.Fatal("хэш не может быть 0: 0 зарезервирован как 'не задан'")
	}
	a, b := HashUpdateData([]byte{1, 2, 3}), HashUpdateData([]byte{1, 2, 4})
	if a == b {
		t.Fatal("разные данные дали одинаковый хэш")
	}
	if HashUpdateData([]byte{1, 2, 3}) != a {
		t.Fatal("хэш недетерминирован")
	}
}
