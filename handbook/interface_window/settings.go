package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Settings = map[string]map[string]string{
	"window_name": {
		_const.RU: `Настройки`,
		_const.EN: `Settings`,
	},
	"sound": {
		_const.RU: `Звук`,
		_const.EN: `Sound`,
	},
	"music_volume": {
		_const.RU: `Громкость музыки:`,
		_const.EN: `Music volume:`,
	},
	"sfx_Volume": {
		_const.RU: `Громкость эффектов:`,
		_const.EN: `Effects volume:`,
	},
	"graphics": {
		_const.RU: `Графика`,
		_const.EN: `Graphics`,
	},
	"traces": {
		_const.RU: `Следы от гусениц:`,
		_const.EN: `Traces from transport:`,
	},
}
