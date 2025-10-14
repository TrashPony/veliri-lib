package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var GameMenu = map[string]map[string]string{
	"button_1": {
		_const.RU:   "Продолжить",
		_const.EN:   "Continue",
		_const.ZhCN: "继续",
	},
	"button_2": {
		_const.RU:   "В главное меню",
		_const.EN:   "To Main Menu",
		_const.ZhCN: "回主菜单",
	},
	"button_3": {
		_const.RU:   "На рабочий стол",
		_const.EN:   "To Desktop",
		_const.ZhCN: "回桌面",
	},
	"button_4": {
		_const.RU:   "Назад",
		_const.EN:   "Back",
		_const.ZhCN: "返回",
	},
}
