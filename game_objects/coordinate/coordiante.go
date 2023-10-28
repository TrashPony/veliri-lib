package coordinate

import (
	"strconv"
	"sync"
)

type Coordinate struct {
	ID     int `json:"id,omitempty"`
	X      int `json:"x,omitempty"`
	Y      int `json:"y,omitempty"`
	Radius int `json:"radius,omitempty"`

	Rotate float64 `json:"rotate,omitempty"`
	State  int     `json:"state,omitempty"`

	MapID      int `json:"map_id,omitempty"`
	RespRotate int `json:"resp_rotate,omitempty"`

	/* мета слушателей */
	/* если тру то с течением времени или по эвенту игрока эвакуируют с этой клетки без его желания */
	Transport bool `json:"transport,omitempty"`
	/* если строка не пуста значит эта клетка прослушивается, например вход в базу (base) или переход в другой сектор (sector),
	   и когда игрок на ней происходит событие */
	Handler string `json:"handler,omitempty"`

	/* говорит работает хендлер или нет, например занята ячейка перехода и тп не работает*/
	HandlerOpen bool `json:"handler_open,omitempty"`

	/* соотвественно место куда попадает игрок после ивента */
	Positions []*Coordinate `json:"positions,omitempty"`
	ToBaseID  int           `json:"to_base_id,omitempty"`
	ToMapID   int           `json:"to_map_id,omitempty"`
	Find      bool          `json:"-"`
	Quest     bool
	access    map[int]bool
	key       string
	mx        sync.Mutex
}

func (coor *Coordinate) GetY() int {
	return coor.Y
}

func (coor *Coordinate) GetX() int {
	return coor.X
}

func (coor *Coordinate) Key() string { //создает уникальный ключ для карты "X:Y"
	if coor.key == "" {
		coor.key = strconv.Itoa(coor.X) + ":" + strconv.Itoa(coor.Y)
	}

	return coor.key
}

func (coor *Coordinate) AddAccess(id int) {
	coor.mx.Lock()
	defer coor.mx.Unlock()

	if coor.access == nil {
		coor.access = make(map[int]bool)
	}

	coor.access[id] = true
}

func (coor *Coordinate) RemoveAccess(id int) {
	coor.mx.Lock()
	defer coor.mx.Unlock()

	if coor.access != nil {
		delete(coor.access, id)
	}
}

func (coor *Coordinate) GetAccess(id int) bool {
	coor.mx.Lock()
	defer coor.mx.Unlock()

	if coor.access == nil {
		return false
	}

	return coor.access[id]
}
