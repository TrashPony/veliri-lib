package rpc_request

import (
	"github.com/TrashPony/veliri-lib/game_objects/web_socket_response"
)

type Request struct {
	Event      string                       `json:"-"`
	Api        string                       `json:"-"`
	OK         bool                         `json:"-"`
	UUID       string                       `json:"-"`
	BattleUUID string                       `json:"-"`
	ID         int                          `json:"-"`
	Response   web_socket_response.Response `json:"-"`
	TypeSlot   int                          `json:"-"`
	Slot       int                          `json:"-"`
	W          bool                         `json:"-"`
	A          bool                         `json:"-"`
	S          bool                         `json:"-"`
	D          bool                         `json:"-"`
	Z          bool                         `json:"-"`
	Sp         bool                         `json:"-"`
	St         bool                         `json:"-"`
	Fire       bool                         `json:"-"`
	X          int                          `json:"-"`
	Y          int                          `json:"-"`
	Data       interface{}                  `json:"-"`
}
