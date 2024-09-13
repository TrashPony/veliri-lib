package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Alerts = map[string]map[string]string{
	"reject": {
		_const.RU: `Отказать`,
		_const.EN: `Reject`,
	},
	"accept": {
		_const.RU: `Принять`,
		_const.EN: `Accept`,
	},
}
