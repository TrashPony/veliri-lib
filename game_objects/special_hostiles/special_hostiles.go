package special_hostiles

import (
	"encoding/json"
	_const "github.com/TrashPony/veliri-lib/const"
	"sort"
	"sync"
)

// SpecialHostiles отношение бота к другим ботам и игрокм
// 0 - нейтральное отношение, но сохраняется фракционная вражба во фри и батл секторах
// < 0 - дружаня, не атакует по причине фракционной вражды во фри, но атакуте в батл секторах, чаще помогает при запросах
// > 0 - враг, атакует его даже если тот в той же фракции, активно принимает участие в запросах против цели и не помогует цели
type SpecialHostiles struct {
	IgnoreHate      [][]*IgnoreHostile  `json:"ignore_hate"`
	SpecialHostiles [][]*SpecialHostile `json:"special_hostiles"`
	mx              sync.RWMutex
}

func (s *SpecialHostiles) initSpecialHostiles() {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.SpecialHostiles = make([][]*SpecialHostile, len(_const.MapBinItems)+1)
	for _, t := range _const.MapBinItems {
		s.SpecialHostiles[t] = make([]*SpecialHostile, 0)
	}

	s.IgnoreHate = make([][]*IgnoreHostile, len(_const.MapBinItems)+1)
	for _, t := range _const.MapBinItems {
		s.IgnoreHate[t] = make([]*IgnoreHostile, 0)
	}
}

func (s *SpecialHostiles) RangeHostiles() <-chan *SpecialHostile {
	s.mx.RLock()
	ch := make(chan *SpecialHostile, s.countHostiles())

	go func() {
		defer func() {
			close(ch)
			s.mx.RUnlock()
		}()

		for _, sType := range s.SpecialHostiles {
			for _, s := range sType {
				ch <- s
			}
		}
	}()

	return ch
}

func (s *SpecialHostiles) countHostiles() int {
	total := 0
	for _, objects := range s.SpecialHostiles {
		total += len(objects)
	}
	return total
}

func (s *SpecialHostiles) UnsafeRangeHostiles() ([][]*SpecialHostile, *sync.RWMutex) {
	s.mx.RLock()
	return s.SpecialHostiles, &s.mx
}

func (s *SpecialHostiles) getHostile(typeHostile string, id int) *SpecialHostile {

	if typeHostile == "meteorite" {
		return &SpecialHostile{} // отдаем фейк что бы нечего не сломать
	}

	if s.SpecialHostiles == nil {
		s.initSpecialHostiles()
	}

	s.mx.RLock()

	// Получаем числовой идентификатор типа
	typeID, ok := _const.MapBinItems[typeHostile]
	if !ok {
		panic("Недопустимый тип врага 1: " + typeHostile)
	}

	// получаем обьект
	var hostile *SpecialHostile
	objects := s.SpecialHostiles[typeID]
	if !(objects == nil || len(objects) == 0) {

		// Для маленьких слайсов линейный поиск быстрее
		if len(objects) < 20 {
			for _, obj := range objects {
				if obj.ID == id {
					hostile = obj
					break
				}
			}
		} else {
			index := sort.Search(len(objects), func(i int) bool {
				return objects[i].ID >= id
			})

			if index < len(objects) && objects[index].ID == id {
				hostile = objects[index]
			}
		}
	}

	// берем игнор
	var ignore bool
	iobjects := s.IgnoreHate[typeID]
	if !(iobjects == nil || len(iobjects) == 0) {
		if len(iobjects) < 20 {
			for _, obj := range iobjects {
				if obj.ID == id {
					ignore = true
					break
				}
			}
		} else {
			indexI := sort.Search(len(iobjects), func(i int) bool {
				return iobjects[i].ID >= id
			})

			if indexI < len(iobjects) && iobjects[indexI].ID == id {
				ignore = iobjects[indexI].ID == id
			}
		}
	}

	s.mx.RUnlock()

	if ignore {
		return &SpecialHostile{} // отдаем фейк что бы нечего не сломать
	}

	if hostile == nil {
		hostile = &SpecialHostile{ID: id, Type: typeHostile}
		hostile = s.addHostile(hostile)
	}

	return hostile
}

func (s *SpecialHostiles) addHostile(h *SpecialHostile) *SpecialHostile {
	if s.SpecialHostiles == nil {
		s.initSpecialHostiles()
	}

	s.mx.Lock()
	defer s.mx.Unlock()

	// Получаем числовой идентификатор типа
	typeID, ok := _const.MapBinItems[h.Type]
	if !ok {
		panic("Недопустимый тип объекта врага 2: " + h.Type)
	}

	// Добавляем объект в список
	objects := s.SpecialHostiles[typeID]
	objects = append(objects, h)

	// Сортируем список по ID
	sort.Slice(objects, func(i, j int) bool {
		return objects[i].ID < objects[j].ID
	})

	s.SpecialHostiles[typeID] = objects
	return h
}

func (s *SpecialHostiles) addIgnore(h *IgnoreHostile, typeHostile string) {
	if s.IgnoreHate == nil {
		s.initSpecialHostiles()
	}

	s.mx.Lock()
	defer s.mx.Unlock()

	// Получаем числовой идентификатор типа
	typeID, ok := _const.MapBinItems[typeHostile]
	if !ok {
		panic("Недопустимый тип объекта врага 3: " + typeHostile)
	}

	// Добавляем объект в список
	objects := s.IgnoreHate[typeID]
	objects = append(objects, h)

	// Сортируем список по ID
	sort.Slice(objects, func(i, j int) bool {
		return objects[i].ID < objects[j].ID
	})

	s.IgnoreHate[typeID] = objects
}

func (s *SpecialHostiles) SetPoints(typeHostile string, id, hatePoint int) {
	if s == nil || (typeHostile == "" && id == 0) {
		return
	}

	s.getHostile(typeHostile, id).SetPoints(hatePoint)
}

func (s *SpecialHostiles) AddPoints(typeHostile string, id, hatePoint int, mod string) {
	if s == nil || (typeHostile == "" && id == 0) {
		return
	}

	s.getHostile(typeHostile, id).AddPoints(hatePoint, mod)
}

func (s *SpecialHostiles) CheckHostile(typeHostile string, id int) (bool, int) {
	hostile := s.getHostile(typeHostile, id)
	hatePoints := hostile.GetPoints()

	return hatePoints > 100, hatePoints
}

func (s *SpecialHostiles) CheckHostileByMod(typeHostile string, id int, mod string) int {
	hostile := s.getHostile(typeHostile, id)
	hatePoints := hostile.GetPointsByMod(mod)

	return hatePoints
}

func (s *SpecialHostiles) AddIgnore(typeHostile string, id int) {
	s.addIgnore(&IgnoreHostile{ID: id}, typeHostile)
}

func (s *SpecialHostiles) GetJsonData() ([]byte, []byte) {
	s.mx.Lock()
	defer s.mx.Unlock()

	ignoreHate, _ := json.Marshal(s.IgnoreHate)
	specialHostiles, _ := json.Marshal(s.SpecialHostiles)

	return ignoreHate, specialHostiles
}

func (s *SpecialHostiles) LoadFromJson(ignoreHate, specialHostiles []byte) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.IgnoreHate = make([][]*IgnoreHostile, 0)
	s.SpecialHostiles = make([][]*SpecialHostile, 0)

	_ = json.Unmarshal(ignoreHate, &s.IgnoreHate)
	_ = json.Unmarshal(specialHostiles, &s.SpecialHostiles)
}

func (s *SpecialHostiles) LoadData(ignoreHate [][]*IgnoreHostile, specialHostiles [][]*SpecialHostile) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.IgnoreHate = ignoreHate
	s.SpecialHostiles = specialHostiles
}

func (s *SpecialHostiles) GetData() ([][]*IgnoreHostile, [][]*SpecialHostile) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	ignoreCopy := make([][]*IgnoreHostile, len(s.IgnoreHate))
	for k, v := range s.IgnoreHate {
		ignoreCopy[k] = make([]*IgnoreHostile, len(v))
		for i, v2 := range v {
			ignoreCopy[k][i] = &IgnoreHostile{v2.ID}
		}
	}

	hostilesCopy := make([][]*SpecialHostile, len(s.SpecialHostiles))
	for k, v := range s.SpecialHostiles {
		hostilesCopy[k] = make([]*SpecialHostile, len(v))
		for i, v2 := range v {
			hostilesCopy[k][i] = v2.Copy()
		}
	}

	return ignoreCopy, hostilesCopy
}

func (s *SpecialHostiles) Clear() {
	s.mx.Lock()
	defer s.mx.Unlock()

	for typeID, store := range s.SpecialHostiles {
		// Новый слайс только для не-nil элементов
		compact := make([]*SpecialHostile, 0, len(store))

		for _, h := range store {
			if h != nil && h.GetPoints() != 0 {
				compact = append(compact, h)
			}
		}

		// Заменяем старый слайс на компактный
		s.SpecialHostiles[typeID] = compact
	}
}
