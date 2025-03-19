package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var AttackDialog = map[string]map[string]string{
	"failed": {
		_const.RU:   `Не удалось!`,
		_const.EN:   `Failed!`,
		_const.ZhCN: `失败！`,
	},
	"success": {
		_const.RU:   `Успех!`,
		_const.EN:   `Success!`,
		_const.ZhCN: `成功！`,
	},
	"text_1": {
		_const.RU:   `Сектор принадлежит кластеру:`,
		_const.EN:   `The sector belongs to the cluster:`,
		_const.ZhCN: `该区域属于集群：`,
	},
	"text_2": {
		_const.RU:   `Установить доступ на врата:`,
		_const.EN:   `Set up access to the gate:`,
		_const.ZhCN: `设置门禁权限：`,
	},
	"text_3": {
		_const.RU:   `Напасть на сектор?`,
		_const.EN:   `Attack the sector?`,
		_const.ZhCN: `攻击该区域？`,
	},
	"option_1": {
		_const.RU:   `Кластер`,
		_const.EN:   `Cluster`,
		_const.ZhCN: `集群`,
	},
	"attack_button": {
		_const.RU:   `Напасть`,
		_const.EN:   `Attack`,
		_const.ZhCN: `攻击`,
	},
	"back_button": {
		_const.RU:   `Назад`,
		_const.EN:   `Back`,
		_const.ZhCN: `返回`,
	},
}
