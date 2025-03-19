package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var ObjectDialog = map[string]map[string]string{
	"window_name": {
		_const.RU:   `Содержимое`,
		_const.EN:   `Loot`,
		_const.ZhCN: `内容`,
	},
	"owner": {
		_const.RU:   `Владелец:`,
		_const.EN:   `Owner: `,
		_const.ZhCN: `所有者：`,
	},
	"hp": {
		_const.RU:   `HP:`,
		_const.EN:   `HP: `,
		_const.ZhCN: `生命值：`,
	},
	"energy": {
		_const.RU:   `Energy:`,
		_const.EN:   `Energy: `,
		_const.ZhCN: `活力：`,
	},
}
