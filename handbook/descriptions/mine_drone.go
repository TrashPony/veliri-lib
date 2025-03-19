package descriptions

import _const "github.com/TrashPony/veliri-lib/const"

var MineDrones = map[string]map[string]DescriptionItem{
	_const.EN: {
		"mine_drone":      {Name: "Mining Drone", Description: "Capable of extracting ore in mines."},
		"builder_drone":   {Name: "Builder Drone", Description: "Capable of constructing structures in mines."},
		"transport_drone": {Name: "Transport Drone", Description: "Has a large cargo hold and can carry multiple items at once."},
	},
	_const.RU: {
		"mine_drone":      {Name: "Дрон шахтер", Description: "Способен разрабатывать руду в шахтах."},
		"builder_drone":   {Name: "Дрон строитель", Description: "Способен строить сооружения в шахтах."},
		"transport_drone": {Name: "Дрон грузовик", Description: "Имеет объемный кузов и способен переносить несколько предметов за раз."},
	},
	_const.ZhCN: {
		"mine_drone":      {Name: "采矿无人机", Description: "能够在矿井中开采矿石。"},
		"builder_drone":   {Name: "建筑无人机", Description: "能够在矿井中建造建筑物。"},
		"transport_drone": {Name: "运输无人机", Description: "拥有宽敞的车厢，能够一次性运输多个物品。"},
	},
}
