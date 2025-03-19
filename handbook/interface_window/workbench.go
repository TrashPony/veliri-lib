package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Workbench = map[string]map[string]string{
	"window_name": {
		_const.RU:   `Производство`,
		_const.EN:   `Production`,
		_const.ZhCN: `生产`,
	},
	"error_1": {
		_const.RU:   `Модуль сервиса уничтожен или отключен.`,
		_const.EN:   `The service module has been destroyed or disabled.`,
		_const.ZhCN: `服务模块已被销毁或禁用。`,
	},
	"error_2": {
		_const.RU:   `Чертеж должен быть на "Складе"`,
		_const.EN:   `The drawing must be in the "Warehouse"`,
		_const.ZhCN: `图纸必须在“仓库”中`,
	},
	"error_3": {
		_const.RU:   `Нехватает ресурсов`,
		_const.EN:   `Not enough resources`,
		_const.ZhCN: `资源不足`,
	},
	"not_allow": {
		_const.RU:   `Мы не доверяем вам и поэтому этот сервис не доступен.`,
		_const.EN:   `We do not trust you and therefore this service is not available.`,
		_const.ZhCN: `我们不信任您，因此此服务不可用。`,
	},
	"queue": {
		_const.RU:   `Очередь производства:`,
		_const.EN:   `Production queue:`,
		_const.ZhCN: `生产队列:`,
	},
	"qty": {
		_const.RU:   `Кол-во:`,
		_const.EN:   `Qty:`,
		_const.ZhCN: `数量:`,
	},
	"create": {
		_const.RU:   `Создать`,
		_const.EN:   `Create`,
		_const.ZhCN: `创建`,
	},
	"cancel": {
		_const.RU:   `Отменить`,
		_const.EN:   `Cancel`,
		_const.ZhCN: `取消`,
	},
	"close": {
		_const.RU:   `Закрыть`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
}
