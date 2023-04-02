package reservoir

import (
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/game_objects/obstacle_point"
	"strconv"
	"sync"
	"time"
)

type Reservoir struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Texture     string    `json:"texture"`
	Type        string    `json:"type"`
	TypeID      int       `json:"type_id"`
	ResourceID  int       `json:"resource_id"`
	Count       int       `json:"count"`
	MapID       int       `json:"map_id"`
	X           int       `json:"x"`
	Y           int       `json:"y"`
	Rotate      float64   `json:"rotate"`
	DestroyTime time.Time `json:"-"`
	MaxCount    int       `json:"max_count"`
	MinCount    int       `json:"min_count"`

	Scale   int                             `json:"scale"`
	Shadow  bool                            `json:"shadow"`
	GeoData []*obstacle_point.ObstaclePoint `json:"geo_data"`
	Height  float64                         `json:"height"`

	Mining     sync.Mutex `json:"-"`
	Complexity int        `json:"complexity"`
	// все пользователи которые добывают эту руду отображаются тут [user_id] progress_points
	// progress_points - % завершения циклы добычи
	miningUsers   map[int]float64
	miningUsersMX sync.Mutex

	CacheJson      []byte `json:"-"`
	CreateJsonTime int64  `json:"-"`

	mx sync.RWMutex
}

func (m *Reservoir) AddMiningUser(userID int, count float64) {
	m.miningUsersMX.Lock()
	defer m.miningUsersMX.Unlock()

	// копание руды, пополняем виртуальную еденицу руды до 100, когда она стала 100 то мы добавляем в инвентарь,
	// если мышку отпустили/промахнулись прогрес сохраняется, в конкретной руде
	if m.miningUsers == nil {
		m.miningUsers = make(map[int]float64)
	}

	m.miningUsers[userID] += count
}

func (m *Reservoir) ResetMiningUser(userID int) {
	m.miningUsersMX.Lock()
	defer m.miningUsersMX.Unlock()

	// копание руды, пополняем виртуальную еденицу руды до 100, когда она стала 100 то мы добавляем в инвентарь,
	// если мышку отпустили/промахнулись прогрес сохраняется, в конкретной руде
	if m.miningUsers == nil {
		m.miningUsers = make(map[int]float64)
	}
	m.miningUsers[userID] = 0
}

func (m *Reservoir) GetMiningUser(userID int) float64 {
	m.miningUsersMX.Lock()
	defer m.miningUsersMX.Unlock()

	// копание руды, пополняем виртуальную еденицу руды до 100, когда она стала 100 то мы добавляем в инвентарь,
	// если мышку отпустили/промахнулись прогрес сохраняется, в конкретной руде
	if m.miningUsers == nil {
		m.miningUsers = make(map[int]float64)
	}

	return m.miningUsers[userID]
}

func (m *Reservoir) SetGeoData() {
	for _, geoPoint := range m.GeoData {
		// заполняем мета данные
		geoPoint.SetParentType("reservoir")
		geoPoint.SetParentID(m.ID)

		if geoPoint.GetResource() {
			geoPoint.SetKey(geoPoint.GetParentType() + strconv.Itoa(m.ID) + strconv.Itoa(geoPoint.GetX()) + strconv.Itoa(geoPoint.GetY()) + "resource")
		} else {
			geoPoint.SetKey(geoPoint.GetParentType() + strconv.Itoa(m.ID) + strconv.Itoa(geoPoint.GetX()) + strconv.Itoa(geoPoint.GetY()))
		}

		geoPoint.SetHeight(m.Height)

		// поворачиваем геодату на угол обьекта
		newX, newY := game_math.RotatePoint(float64(m.GetX()), float64(m.GetY()), float64(m.GetX()), float64(m.GetY()), m.GetRotate())

		geoPoint.SetX(int(newX))
		geoPoint.SetY(int(newY))
	}
}

func (m *Reservoir) GetGeoData() []*obstacle_point.ObstaclePoint {
	return m.GeoData
}

func (m *Reservoir) GetX() int {
	m.mx.RLock()
	defer m.mx.RUnlock()

	return m.X
}

func (m *Reservoir) SetX(x int) {
	m.mx.Lock()
	defer m.mx.Unlock()

	m.X = x
}

func (m *Reservoir) GetY() int {
	m.mx.RLock()
	defer m.mx.RUnlock()

	return m.Y
}

func (m *Reservoir) SetY(y int) {
	m.mx.Lock()
	defer m.mx.Unlock()

	m.Y = y
}

func (m *Reservoir) GetRotate() float64 {
	m.mx.RLock()
	defer m.mx.RUnlock()

	return m.Rotate
}

func (m *Reservoir) SetRotate(rotate float64) {
	m.mx.Lock()
	defer m.mx.Unlock()

	m.Rotate = rotate
}

func (m *Reservoir) GetCount() int {
	m.mx.RLock()
	defer m.mx.RUnlock()

	return m.Count
}

func (m *Reservoir) SetCount(count int) {
	m.mx.Lock()
	defer m.mx.Unlock()

	m.Count = count
}

func (m *Reservoir) GetJSON(mapTime int64) []byte {
	m.mx.RLock()
	defer m.mx.RUnlock()

	if m.CreateJsonTime == mapTime && len(m.CacheJson) > 0 {
		return m.CacheJson
	}

	if m.CacheJson == nil {
		m.CacheJson = []byte{}
	}
	m.CacheJson = m.CacheJson[:0]

	m.CacheJson = append(m.CacheJson, game_math.GetIntBytes(m.ID)...)
	m.CacheJson = append(m.CacheJson, game_math.GetIntBytes(m.GetX())...)
	m.CacheJson = append(m.CacheJson, game_math.GetIntBytes(m.GetY())...)
	m.CacheJson = append(m.CacheJson, game_math.GetIntBytes(int(m.GetRotate()))...)
	m.CacheJson = append(m.CacheJson, game_math.GetIntBytes(m.Scale)...)
	m.CacheJson = append(m.CacheJson, game_math.BoolToByte(m.Shadow))
	m.CacheJson = append(m.CacheJson, game_math.GetIntBytes(m.MapID)...)
	m.CacheJson = append(m.CacheJson, game_math.GetIntBytes(int(m.Height))...)
	m.CacheJson = append(m.CacheJson, game_math.GetIntBytes(m.Count)...)
	m.CacheJson = append(m.CacheJson, game_math.GetIntBytes(m.Complexity)...)

	m.CacheJson = append(m.CacheJson, byte(len([]byte(m.Texture))))
	m.CacheJson = append(m.CacheJson, []byte(m.Texture)...)

	m.CacheJson = append(m.CacheJson, byte(len([]byte(m.Name))))
	m.CacheJson = append(m.CacheJson, []byte(m.Name)...)

	m.CreateJsonTime = mapTime

	return m.CacheJson
}

func (m *Reservoir) GetUpdateData(mapTime int64) []byte {
	command := []byte{}

	game_math.ReuseByteSlice(&command, 0, game_math.GetIntBytes(m.GetCount()))

	return command
}
