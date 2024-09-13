package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var ProcessorRoot = map[string]map[string]string{
	// ItemsDropZone
	"head": {
		_const.RU: `Переработать:`,
		_const.EN: `Recycler:`,
	},
	"text_1": {
		_const.RU: `Перетащите сюда предметы которые хотите переработать.`,
		_const.EN: `Drag the items you want to recycle here.`,
	},
	// ProcessorRoot
	"window_name": {
		_const.RU: `Переработчик`,
		_const.EN: `Processor`,
	},
	"to_storage": {
		_const.RU: `Предметы отправлены на <span class="importantly">склад</span>`,
		_const.EN: `Items sent to <span class="importantly">warehouse</span>`,
	},
	"text_2": {
		_const.RU: `Модуль сервиса уничтожен или отключен.`,
		_const.EN: `The service module has been destroyed or disabled.`,
	},
	"not_allow": {
		_const.RU: `Мы не доверяем вам и поэтому этот сервис не доступен.`,
		_const.EN: `We do not trust you and therefore this service is not available.`,
	},
	"lost_items": {
		_const.RU: `Потери:`,
		_const.EN: `Losses:`,
	},
	"tax": {
		_const.RU: `Налог:`,
		_const.EN: `Tax:`,
	},
	"result": {
		_const.RU: `Результат:`,
		_const.EN: `Result:`,
	},
	"button_1": {
		_const.RU: `Закрыть`,
		_const.EN: `Close`,
	},
	"button_2": {
		_const.RU: `Переработать`,
		_const.EN: `Recycle`,
	},
	"text_3": {
		_const.RU: `Переработать`,
		_const.EN: `Recycle`,
	},
}
