package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var MiniMap = map[string]map[string]string{
	"zoom": {
		_const.RU:   `Зум`,
		_const.EN:   `Zoom`,
		_const.ZhCN: `飞涨`,
	},
	"fix_cam": {
		_const.RU:   `Фиксация камеры`,
		_const.EN:   `Camera fixation`,
		_const.ZhCN: `固定相机`,
	},
	"battle_cam": {
		_const.RU:   `Боевой обзор`,
		_const.EN:   `Combat review`,
		_const.ZhCN: `战斗回顾`,
	},
}
