package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Missions = map[string]map[string]string{
	// Missions
	"window_name": {
		_const.RU:   `Активные задания`,
		_const.EN:   `Active tasks`,
		_const.ZhCN: `活跃任务`,
	},
	"button_1": {
		_const.RU:   `К заданиям`,
		_const.EN:   `To missions`,
		_const.ZhCN: `前往任务`,
	},
	// Points
	"text_1": {
		_const.RU:   `Обновлено`,
		_const.EN:   `Updated`,
		_const.ZhCN: `已更新`,
	},
	"button_2": {
		_const.RU:   `помощь`,
		_const.EN:   `help`,
		_const.ZhCN: `帮助`,
	},
	"button_3": {
		_const.RU:   `потерял предмет?`,
		_const.EN:   `lost an item?`,
		_const.ZhCN: `丢失物品？`,
	},
	"no_mission": {
		_const.RU:   `Нет заданий`,
		_const.EN:   `No Missions`,
		_const.ZhCN: `暂无任务`,
	},
}
