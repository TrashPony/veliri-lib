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
	"tip_1": {
		_const.RU:   `Сектор с пониженным уровнем безопасности`,
		_const.EN:   `Low-security sector`,
		_const.ZhCN: `低安全等级区域`,
	},
	"tip_2": {
		_const.RU:   `Количество воинов вашей фракции в этом секторе`,
		_const.EN:   `Number of your faction's warriors in this sector`,
		_const.ZhCN: `本区域中你方势力的战士数量`,
	},
	"tip_3": {
		_const.RU:   `В этом секторе находится аванпост фракции: <span class='importantly'>%FRACTION%</span>`,
		_const.EN:   `This sector contains an outpost of the <span class='importantly'>%FRACTION%</span> faction`,
		_const.ZhCN: `此区域有<span class='importantly'>%FRACTION%</span>势力的前哨站`,
	},
	"tip_4": {
		_const.RU:   `Сектор находится под контролем фракции: <span class='importantly'>%FRACTION%</span>`,
		_const.EN:   `Sector controlled by the <span class='importantly'>%FRACTION%</span> faction`,
		_const.ZhCN: `此区域由<span class='importantly'>%FRACTION%</span>势力控制`,
	},
	"tip_5": {
		_const.RU:   `На этот сектор готовится нападение`,
		_const.EN:   `This sector is about to be attacked`,
		_const.ZhCN: `此区域即将遭受攻击`,
	},
	"tip_6": {
		_const.RU:   `Этот сектор является приоритетной целью`,
		_const.EN:   `This sector is a priority target`,
		_const.ZhCN: `此区域是优先目标`,
	},
	"tip_7": {
		_const.RU:   `В этом секторе ведется сражение`,
		_const.EN:   `Battle in progress in this sector`,
		_const.ZhCN: `此区域正在发生战斗`,
	},
}
