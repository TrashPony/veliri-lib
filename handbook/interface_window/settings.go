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
	"interface": {
		_const.RU:   `Интерфейс`,
		_const.EN:   `Interface`,
		_const.ZhCN: `界面`,
	},
	"traces": {
		_const.RU:   `Следы от гусениц:`,
		_const.EN:   `Traces from transport:`,
		_const.ZhCN: `载具痕迹:`,
	},
	"external_scaling": {
		_const.RU:   `Внешнее масштабирование`,
		_const.EN:   `External scaling`,
		_const.ZhCN: `外部缩放`,
	},
	"internal_scaling": {
		_const.RU:   `Внутреннее масштабирование`,
		_const.EN:   `Internal scaling`,
		_const.ZhCN: `内部缩放`,
	},
	"control": {
		_const.RU:   `Управление`,
		_const.EN:   `Control`,
		_const.ZhCN: `控制`,
	},
	"full": {
		_const.RU:   `Относительно корпуса`,
		_const.EN:   `Relative to body`,
		_const.ZhCN: `相对于身体`,
	},
	"simplified": {
		_const.RU:   `Относительно экрана`,
		_const.EN:   `Relative to screen`,
		_const.ZhCN: `相对于屏幕`,
	},
	"alert_reload": {
		_const.RU:   `Рекомендуется перезагрузить игру.`,
		_const.EN:   `It is recommended to restart the game.`,
		_const.ZhCN: `建议重新开始游戏。`,
	},
	"reload_button": {
		_const.RU:   `Перезагрузить`,
		_const.EN:   `Reload`,
		_const.ZhCN: `重新启动`,
	},
	"postprocess_intensity": {
		_const.RU:   `Постобработка`,
		_const.EN:   `Post-processing`,
		_const.ZhCN: `后处理`,
	},
	"camera_leading_travel": {
		_const.RU:   `Опережение походной камеры`,
		_const.EN:   `Travel camera lead`,
		_const.ZhCN: `行进摄像机预置`,
	},
	"camera_combat_smoothing": {
		_const.RU:   `Сглаживание боевой камеры`,
		_const.EN:   `Combat camera smoothing`,
		_const.ZhCN: `战斗摄像机平滑度`,
	},
}
