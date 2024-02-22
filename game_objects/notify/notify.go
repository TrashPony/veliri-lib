package notify

import "github.com/TrashPony/veliri-lib/game_objects/inventory"

// todo много лишних полей, весь пейлоад должен быть в поле Data, а тут только служебные данные
type Notify struct {
	Name              string              `json:"name,omitempty"`
	ToUserID          int                 `json:"to_user_id,omitempty"`
	ModalText         bool                `json:"modal_text,omitempty"`
	Text              map[string]string   `json:"text,omitempty"`
	Off               bool                `json:"off,omitempty"`
	UUID              string              `json:"uuid,omitempty"`
	UUIDBattle        string              `json:"uuid_battle,omitempty"`
	Event             string              `json:"event,omitempty"`
	Send              bool                `json:"send,omitempty"`
	Data              interface{}         `json:"data,omitempty"`
	Time              int64               `json:"time,omitempty"` // unix nano utc
	TimeOut           int                 `json:"time_out,omitempty"`
	Destroy           bool                `json:"destroy,omitempty"`
	Count             int                 `json:"count,omitempty"`
	Price             int                 `json:"price,omitempty"`
	UpdateInventory   bool                `json:"update_inventory,omitempty"`
	UpdateChat        bool                `json:"update_chat,omitempty"`
	UpdateMission     bool                `json:"update_mission,omitempty"`
	SkillUp           bool                `json:"skill_up,omitempty"`
	Item              *inventory.Slot     `json:"item,omitempty"`
	X                 int                 `json:"x,omitempty"`
	Y                 int                 `json:"y,omitempty"`
	BroadcastMapID    int                 `json:"broadcast_map_id,omitempty"`
	BroadcastFraction string              `json:"broadcast_fraction,omitempty"`
	DeadID            int                 `json:"dead_id,omitempty"`
	KillerID          int                 `json:"killer_id,omitempty"`
	DeadName          string              `json:"dead_name,omitempty"`
	KillerName        string              `json:"killer_name,omitempty"`
	DeadFraction      string              `json:"dead_fraction,omitempty"`
	KillerFraction    string              `json:"killer_fraction,omitempty"`
	AmmoID            int                 `json:"ammo_id,omitempty"`
	From              int                 `json:"from,omitempty"` // unit id
	To                int                 `json:"to,omitempty"`   // unit id
	FromUserName      string              `json:"from_user_name,omitempty"`
	ToUserName        string              `json:"to_user_name,omitempty"`
	Answers           []map[string]Answer `json:"answers,omitempty"`
	MapID             int                 `json:"map_id,omitempty"`
	BaseID            int                 `json:"base_id,omitempty"`
	Chat              bool                `json:"chat,omitempty"`
	ChatID            int                 `json:"chat_id,omitempty"`
}

type Answer struct {
	Text string `json:"text"`
	ID   int    `json:"id"`
}
