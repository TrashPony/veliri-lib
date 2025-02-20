package visible_objects

import (
	"fmt"
	"sort"
	"sync"
)

const idMemoryType = 0

func (v *VisibleObjectsStore) AddDynamicObject(object *VisibleObject) {
	if v.visibleObjects == nil {
		v.InitVisibleObjects()
	}

	v.mx.Lock()
	defer v.mx.Unlock()

	objects := v.visibleObjects[idMemoryType]
	objects = append(objects, object)

	// Сортируем список по ID
	sort.Slice(objects, func(i, j int) bool {
		return objects[i].ID < objects[j].ID
	})

	v.visibleObjects[idMemoryType] = objects
}
func (v *VisibleObjectsStore) AddDynamicMemoryObject(object *VisibleObject) {
	v.AddDynamicObject(object)
}

func (v *VisibleObjectsStore) RemoveDynamicObject(id int) {
	v.mx.Lock()
	defer v.mx.Unlock()

	objects := v.visibleObjects[idMemoryType]

	if objects == nil || len(objects) == 0 {
		return
	}

	// Находим индекс объекта
	index := sort.Search(len(objects), func(i int) bool {
		return objects[i].ID >= id
	})

	if index < len(objects) && objects[index].ID == id {
		// Удаляем объект
		objects = append(objects[:index], objects[index+1:]...)
		v.visibleObjects[idMemoryType] = objects
	}
}

func (v *VisibleObjectsStore) GetMapDynamicObjectByID(id int) *VisibleObject {
	if v.visibleObjects == nil {
		v.InitVisibleObjects()
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in GetMapDynamicObjectByID", r)
		}
	}()
	v.mx.RLock()
	defer v.mx.RUnlock()

	objects := v.visibleObjects[idMemoryType]

	if objects == nil || len(objects) == 0 {
		return nil
	}

	// Бинарный поиск по ID
	index := sort.Search(len(objects), func(i int) bool {
		return objects[i].ID >= id
	})

	if index < len(objects) && objects[index].ID == id {
		return objects[index]
	}

	return nil
}

func (v *VisibleObjectsStore) GetMapDynamicObjects(mapID int) <-chan *VisibleObject {
	if v.visibleObjects == nil {
		v.InitVisibleObjects()
	}

	v.mx.RLock()
	objChan := make(chan *VisibleObject, v.countVisibleObjects())

	go func() {
		defer func() {
			v.mx.RUnlock()
			close(objChan)
		}()

		objects := v.visibleObjects[idMemoryType]
		if objects == nil {
			return
		}

		for _, obj := range objects {
			if obj.MapID == mapID {
				objChan <- obj
			}
		}
	}()
	return objChan
}

func (v *VisibleObjectsStore) UnsafeRangeMapDynamicObjects() ([][]*VisibleObject, *sync.RWMutex) {
	return v.visibleObjects, &v.mx
}
