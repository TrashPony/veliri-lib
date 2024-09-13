package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Workbench = map[string]map[string]string{
	"window_name": {
		_const.RU: `Производство`,
		_const.EN: `Production`,
	},
	"error_1": {
		_const.RU: `Модуль сервиса уничтожен или отключен.`,
		_const.EN: `The service module has been destroyed or disabled.`,
	},
	"error_2": {
		_const.RU: `Чертеж должен быть на "Складе"`,
		_const.EN: `The drawing must be in the "Warehouse"`,
	},
	"error_3": {
		_const.RU: `Нехватает ресурсов`,
		_const.EN: `Not enough resources`,
	},
	"not_allow": {
		_const.RU: `Мы не доверяем вам и поэтому этот сервис не доступен.`,
		_const.EN: `We do not trust you and therefore this service is not available.`,
	},
	"queue": {
		_const.RU: `Очередь производства:`,
		_const.EN: `Production queue:`,
	},
	"qty": {
		_const.RU: `Кол-во:`,
		_const.EN: `Qty:`,
	},
	"create": {
		_const.RU: `Создать`,
		_const.EN: `Create`,
	},
	"cancel": {
		_const.RU: `Отменить`,
		_const.EN: `Cancel`,
	},
	"close": {
		_const.RU: `Закрыть`,
		_const.EN: `Close`,
	},
}
