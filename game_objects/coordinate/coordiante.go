package coordinate

import (
	"github.com/TrashPony/veliri-lib/game_objects/access_manager"
	"strconv"
	"sync"
)

type Coordinate struct {
	ID        int    `json:"id,omitempty"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Radius    int    `json:"radius,omitempty"`
	UUID      string `json:"UUID,omitempty"`
	Collision bool   `json:"collision"`

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
	Hack        bool `json:"-"`

	/* соотвественно место куда попадает игрок после ивента */
	Positions     []*Coordinate `json:"positions,omitempty"`
	ToBaseID      int           `json:"to_base_id,omitempty"`
	ToMapID       int           `json:"to_map_id,omitempty"`
	Find          bool          `json:"-"`
	attr          map[string]interface{}
	Road          bool
	AccessManager access_manager.AccessManager
	key           string
	mx            sync.Mutex
}

type AccessPoint struct {
	Type int
	ID   int
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

func (coor *Coordinate) AddAttr(attr map[string]interface{}) {
	coor.attr = attr
}

func (coor *Coordinate) GetAttr() map[string]interface{} {
	return coor.attr
}
