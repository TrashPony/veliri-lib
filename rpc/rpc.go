package rpc

import (
	"encoding/gob"
	"github.com/TrashPony/veliri-lib/game_objects/rpc_request"
	"github.com/valyala/gorpc"
)

type RPC struct {
	Server           *gorpc.Server
	VeliriMainClient *gorpc.Client
	Error            bool
}

func GobRegister() {
	gob.Register(rpc_request.Request{})
	gob.Register(map[int]string{})
	gob.Register(map[int][]byte{})
}
