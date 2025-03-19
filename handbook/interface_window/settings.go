package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Settings = map[string]map[string]string{
	"window_name": {
		_const.RU:   `Настройки`,
		_const.EN:   `Settings`,
		_const.ZhCN: `设置`,
	},
	"sound": {
		_const.RU:   `Звук`,
		_const.EN:   `Sound`,
		_const.ZhCN: `声音`,
	},
	"music_volume": {
		_const.RU:   `Громкость музыки:`,
		_const.EN:   `Music volume:`,
		_const.ZhCN: `音乐音量:`,
	},
	"sfx_Volume": {
		_const.RU:   `Громкость эффектов:`,
		_const.EN:   `Effects volume:`,
		_const.ZhCN: `音效音量:`,
	},
	"graphics": {
		_const.RU:   `Графика`,
		_const.EN:   `Graphics`,
		_const.ZhCN: `图形`,
	},
	"traces": {
		_const.RU:   `Следы от гусениц:`,
		_const.EN:   `Traces from transport:`,
		_const.ZhCN: `载具痕迹:`,
	},
}
