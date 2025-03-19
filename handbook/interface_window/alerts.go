package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Alerts = map[string]map[string]string{
	"reject": {
		_const.RU:   `Отказать`,
		_const.EN:   `Reject`,
		_const.ZhCN: `拒绝`,
	},
	"accept": {
		_const.RU:   `Принять`,
		_const.EN:   `Accept`,
		_const.ZhCN: `接受`,
	},
}
