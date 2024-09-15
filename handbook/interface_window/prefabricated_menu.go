package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var PrefabricatedMenu = map[string]map[string]string{
	"window_name": {
		_const.RU: `Производство деталей`,
		_const.EN: `Parts production`,
	},
	"error_1": {
		_const.RU: `Модуль сервиса уничтожен или отключен.`,
		_const.EN: `The service module has been destroyed or disabled.`,
	},
	"error_2": {
		_const.RU: `Не хватило предметов.`,
		_const.EN: `Not enough items.`,
	},
	"not_allow": {
		_const.RU: `Мы не доверяем вам и поэтому этот сервис не доступен.`,
		_const.EN: `We do not trust you and therefore this service is not available.`,
	},
	"to_storage": {
		_const.RU: `Предметы отправлены на <span class="importantly">склад</span>`,
		_const.EN: `Items sent to <span class="importantly">warehouse</span>`,
	},
	"detail_select": {
		_const.RU: `Выбери деталь`,
		_const.EN: `Choose a part`,
	},
	"tax": {
		_const.RU: `Налог:`,
		_const.EN: `Tax:`,
	},
	"button_1": {
		_const.RU: `Произвести`,
		_const.EN: `Produce`,
	},
	"need_resources": {
		_const.RU: `Необходимо ресурсов`,
		_const.EN: `Resources needed`,
	},
	"result": {
		_const.RU: `Результат`,
		_const.EN: `Result`,
	},
}
