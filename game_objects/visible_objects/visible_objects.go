package visible_objects

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"sort"
	"sync"
)

type VisibleObjectsStore struct {
	visibleObjects [][]*VisibleObject // [type_object][id_object]
	mx             sync.RWMutex
}

func (v *VisibleObjectsStore) InitVisibleObjects() {
	v.mx.Lock()
	defer v.mx.Unlock()

	if v.visibleObjects != nil {
		for t := range v.visibleObjects {
			for _, v := range v.visibleObjects[t] {
				v.UpdateMsg = nil
				v.Object = nil
				v.ObjectJSON = nil
			}
		}
	}

	v.visibleObjects = make([][]*VisibleObject, len(_const.MapBinItems)+1)
	for _, t := range _const.MapBinItems {
		v.visibleObjects[t] = make([]*VisibleObject, 0)
	}
}

func (v *VisibleObjectsStore) GetVisibleObjectByTypeAndID(typeObj string, id int) *VisibleObject {
	if v.visibleObjects == nil {
		v.InitVisibleObjects()
	}

	v.mx.RLock()
	defer v.mx.RUnlock()

	// Получаем числовой идентификатор типа
	typeID, ok := _const.MapBinItems[typeObj]
	if !ok {
		panic("Недопустимый тип объекта 1: " + typeObj)
	}

	objects := v.visibleObjects[typeID]
	if objects == nil || len(objects) == 0 {
		return nil // Список пуст
	}

	// Бинарный поиск по ID
	index := sort.Search(len(objects), func(i int) bool {
		return objects[i].IDObject >= id
	})

	if index < len(objects) && objects[index].IDObject == id {
		return objects[index]
	}

	return nil
}

func (v *VisibleObjectsStore) GetVisibleObjects() <-chan *VisibleObject {
	v.mx.RLock()
	objs := make(chan *VisibleObject, v.countVisibleObjects())
	go func() {
		defer func() {
			v.mx.RUnlock()
			close(objs)
		}()

		for _, objects := range v.visibleObjects {
			for _, obj := range objects {
				objs <- obj
			}
		}
	}()
	return objs
}

func (v *VisibleObjectsStore) countVisibleObjects() int {
	total := 0
	for _, objects := range v.visibleObjects {
		total += len(objects)
	}
	return total
}

func (v *VisibleObjectsStore) AddVisibleObject(newObj *VisibleObject) {
	if v.visibleObjects == nil {
		v.InitVisibleObjects()
	}

	v.mx.Lock()
	defer v.mx.Unlock()

	// Получаем числовой идентификатор типа
	typeID, ok := _const.MapBinItems[newObj.TypeObject]
	if !ok {
		panic("Недопустимый тип объекта 2: " + newObj.TypeObject)
	}

	// Добавляем объект в список
	objects := v.visibleObjects[typeID]
	objects = append(objects, newObj)

	// Сортируем список по ID
	sort.Slice(objects, func(i, j int) bool {
		return objects[i].IDObject < objects[j].IDObject
	})

	v.visibleObjects[typeID] = objects
}

func (v *VisibleObjectsStore) RemoveVisibleObject(removeObj *VisibleObject) {
	v.mx.Lock()
	defer v.mx.Unlock()

	// Получаем числовой идентификатор типа
	typeID, ok := _const.MapBinItems[removeObj.TypeObject]
	if !ok {
		panic("Недопустимый тип объекта 3: " + removeObj.TypeObject)
	}

	objects := v.visibleObjects[typeID]
	if objects == nil || len(objects) == 0 {
		return // Список пуст
	}

	// Находим индекс объекта
	index := sort.Search(len(objects), func(i int) bool {
		return objects[i].IDObject >= removeObj.IDObject
	})

	if index < len(objects) && objects[index].IDObject == removeObj.IDObject {
		// Удаляем объект
		objects = append(objects[:index], objects[index+1:]...)
		v.visibleObjects[typeID] = objects
	}
}

type VisibleObject struct {
	ID         int    `json:"id_mark"`
	IDObject   int    `json:"id"`
	TypeObject string `json:"to"`
	UUIDObject string `json:"uo"`
	//UUID       string `json:"-"`
	View   bool   `json:"-"`    // в прямой видимости
	Radar  bool   `json:"-"`    // видим только радаром
	Type   string `json:"type"` // fly(летающий), ground(наземный), structure(структура), resource(ресурс)
	update int

	HP            int         `json:"-"`
	Complete      float64     `json:"-"`
	Scale         int         `json:"-"`
	Energy        int         `json:"-"`
	MapID         int         `json:"mid"`
	X             int         `json:"-"`
	Y             int         `json:"-"`
	Rotate        int         `json:"-"`
	OwnerPlayerID int         `json:"-"`
	CorporationID int         `json:"-"`
	OwnerID       int         `json:"-"`
	Object        interface{} `json:"-"`
	ObjectJSON    []byte      `json:"-"`
	UpdateChecker []byte      `json:"-"`
	Work          bool        `json:"-"`
	ForceView     bool        `json:"-"`

	UpdateMsg *UpdateObjMap `json:"-"`

	Pool string `json:"-"`

	mx sync.RWMutex
}

type UpdateObjMap struct {
	UpdateData []byte
	Update     bool
	JSON       string
	ServerTime int64
	mx         sync.Mutex
}

func (v *VisibleObject) LockUpdateObjMap() {
	v.UpdateMsg.mx.Lock()
}

func (v *VisibleObject) UnlockUpdateObjMap() {
	v.UpdateMsg.mx.Unlock()
}

func (v *VisibleObject) GetUpdateData(mapTime int64) []byte {
	return v.UpdateChecker
}

func (v *VisibleObject) GetUpdate() int {
	return v.update
}

func (v *VisibleObject) SetUpdate(update int) {
	v.update = update
}

func (v *VisibleObject) GetView() bool {
	return v.View
}

func (v *VisibleObject) SetView(view bool) {
	v.View = view
}

func (v *VisibleObject) GetRadar() bool {
	return v.Radar
}

func (v *VisibleObject) SetRadar(radar bool) {
	v.Radar = radar
}
