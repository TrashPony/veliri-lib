package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var PrefabricatedMenu = map[string]map[string]string{
	"window_name": {
		_const.RU:   `Производство деталей`,
		_const.EN:   `Parts production`,
		_const.ZhCN: `零件生产`,
	},
	"error_1": {
		_const.RU:   `Модуль сервиса уничтожен или отключен.`,
		_const.EN:   `The service module has been destroyed or disabled.`,
		_const.ZhCN: `服务模块已被销毁或禁用。`,
	},
	"error_2": {
		_const.RU:   `Не хватило предметов.`,
		_const.EN:   `Not enough items.`,
		_const.ZhCN: `物品不足。`,
	},
	"not_allow": {
		_const.RU:   `Мы не доверяем вам и поэтому этот сервис не доступен.`,
		_const.EN:   `We do not trust you and therefore this service is not available.`,
		_const.ZhCN: `我们不信任您，因此此服务不可用。`,
	},
	"to_storage": {
		_const.RU:   `Предметы отправлены на <span class="importantly">склад</span>`,
		_const.EN:   `Items sent to <span class="importantly">warehouse</span>`,
		_const.ZhCN: `物品已发送到<span class="importantly">仓库</span>`,
	},
	"detail_select": {
		_const.RU:   `Выбери деталь`,
		_const.EN:   `Choose a part`,
		_const.ZhCN: `选择零件`,
	},
	"tax": {
		_const.RU:   `Налог:`,
		_const.EN:   `Tax:`,
		_const.ZhCN: `税收:`,
	},
	"button_1": {
		_const.RU:   `Произвести`,
		_const.EN:   `Produce`,
		_const.ZhCN: `生产`,
	},
	"need_resources": {
		_const.RU:   `Необходимо ресурсов`,
		_const.EN:   `Resources needed`,
		_const.ZhCN: `所需资源`,
	},
	"result": {
		_const.RU:   `Результат`,
		_const.EN:   `Result`,
		_const.ZhCN: `结果`,
	},
}
