package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var GlobalMap = map[string]map[string]string{
	// GlobalMap
	"window_name": {
		_const.RU:   `Карта мира`,
		_const.EN:   `World map`,
		_const.ZhCN: `世界地图`,
	},
	// MapPin
	"button_1": {
		_const.RU:   `Автопилот`,
		_const.EN:   `Autopilot`,
		_const.ZhCN: `自动驾驶`,
	},
	"button_2": {
		_const.RU:   `Установить цель`,
		_const.EN:   `Set a path`,
		_const.ZhCN: `设置路径`,
	},
	"text_1": {
		_const.RU:   `В секторе идет бой!`,
		_const.EN:   `There is a battle in the sector!`,
		_const.ZhCN: `该区域正在发生战斗！`,
	},
	"text_2": {
		_const.RU:   `Атака возможна:`,
		_const.EN:   `Attack possible:`,
		_const.ZhCN: `可攻击：`,
	},
	"button_3": {
		_const.RU:   `Закрыть`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
}
