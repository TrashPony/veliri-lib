package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Department = map[string]map[string]string{
	"window_name_1": {
		_const.RU:   `Терминал`,
		_const.EN:   `Terminal`,
		_const.ZhCN: `终端`,
	},
	"text_1": {
		_const.RU:   `Пункт назначения:`,
		_const.EN:   `Destination:`,
		_const.ZhCN: `目的地：`,
	},
	"text_2": {
		_const.RU:   `Сектор`,
		_const.EN:   `Sector`,
		_const.ZhCN: `区域`,
	},
	"text_3": {
		_const.RU:   `База`,
		_const.EN:   `Base`,
		_const.ZhCN: `基地`,
	},
	"button_1": {
		_const.RU:   `карта`,
		_const.EN:   `map`,
		_const.ZhCN: `地图`,
	},
	"text_4": {
		_const.RU:   `Награда при завершение задания:`,
		_const.EN:   `Reward upon completion of the task:`,
		_const.ZhCN: `任务完成奖励：`,
	},
}
