package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var AttackDialog = map[string]map[string]string{
	"failed": {
		_const.RU: `Не удалось!`,
		_const.EN: `Failed!`,
	},
	"success": {
		_const.RU: `Успех!`,
		_const.EN: `Success!`,
	},
	"text_1": {
		_const.RU: `Сектор принадлежит кластеру:`,
		_const.EN: `The sector belongs to the cluster:`,
	},
	"text_2": {
		_const.RU: `Установить доступ на врата:`,
		_const.EN: `Set up access to the gate:`,
	},
	"text_3": {
		_const.RU: `Напасть на сектор?`,
		_const.EN: `Attack the sector?`,
	},
	"option_1": {
		_const.RU: `Кластер`,
		_const.EN: `Cluster`,
	},
	"attack_button": {
		_const.RU: `Напасть`,
		_const.EN: `Attack`,
	},
	"back_button": {
		_const.RU: `Назад`,
		_const.EN: `Back`,
	},
}
