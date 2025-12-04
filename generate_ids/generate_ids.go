package generate_ids

import "sync"

var (
	idMx                     sync.Mutex
	itemsIDGenerate          = 0
	boxesIDGenerate          = 0
	dynamicObjectsIDGenerate = 0
	reservoirsIDGenerate     = 0
	markIDGenerate           = 0
	droneIDGenerate          = 0
	bulletIDGenerate         = 0
)

func Init(startItems, startBoxes, startDynamicObjects, startReservoirs int) {
	idMx.Lock()
	defer idMx.Unlock()

	itemsIDGenerate = startItems
	boxesIDGenerate = startBoxes
	dynamicObjectsIDGenerate = startDynamicObjects
	reservoirsIDGenerate = startReservoirs
}

func GetItemID() int {
	idMx.Lock()
	defer idMx.Unlock()

	itemsIDGenerate++
	// Защита от переполнения int
	if itemsIDGenerate < 0 {
		itemsIDGenerate = 1
	}

	return itemsIDGenerate
}

func GetBoxID() int {
	idMx.Lock()
	defer idMx.Unlock()

	boxesIDGenerate++
	if boxesIDGenerate < 0 {
		boxesIDGenerate = 1
	}

	return boxesIDGenerate
}

func GetObjectID() int {
	idMx.Lock()
	defer idMx.Unlock()

	dynamicObjectsIDGenerate++
	if dynamicObjectsIDGenerate < 0 {
		dynamicObjectsIDGenerate = 1
	}

	return dynamicObjectsIDGenerate
}

func GetReservoirID() int {
	idMx.Lock()
	defer idMx.Unlock()

	reservoirsIDGenerate++
	if reservoirsIDGenerate < 0 {
		reservoirsIDGenerate = 1
	}

	return reservoirsIDGenerate
}

func GetMarkID() int {
	idMx.Lock()
	defer idMx.Unlock()

	markIDGenerate++
	if markIDGenerate < 0 {
		markIDGenerate = 1
	}

	return markIDGenerate
}

func GetDroneID() int {
	idMx.Lock()
	defer idMx.Unlock()

	droneIDGenerate++
	if droneIDGenerate < 0 {
		droneIDGenerate = 1
	}

	return droneIDGenerate
}

func GetBulletID() int {
	idMx.Lock()
	defer idMx.Unlock()

	bulletIDGenerate++
	if bulletIDGenerate < 0 {
		bulletIDGenerate = 1
	}

	return bulletIDGenerate
}
