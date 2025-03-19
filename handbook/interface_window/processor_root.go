package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var ProcessorRoot = map[string]map[string]string{
	// ItemsDropZone
	"head": {
		_const.RU:   `Переработать:`,
		_const.EN:   `Recycler:`,
		_const.ZhCN: `回收:`,
	},
	"text_1": {
		_const.RU:   `Перетащите сюда предметы которые хотите переработать.`,
		_const.EN:   `Drag the items you want to recycle here.`,
		_const.ZhCN: `将您想要回收的物品拖到此处。`,
	},
	// ProcessorRoot
	"window_name": {
		_const.RU:   `Переработчик`,
		_const.EN:   `Processor`,
		_const.ZhCN: `处理器`,
	},
	"to_storage": {
		_const.RU:   `Предметы отправлены на <span class="importantly">склад</span>`,
		_const.EN:   `Items sent to <span class="importantly">warehouse</span>`,
		_const.ZhCN: `物品已发送到<span class="importantly">仓库</span>`,
	},
	"text_2": {
		_const.RU:   `Модуль сервиса уничтожен или отключен.`,
		_const.EN:   `The service module has been destroyed or disabled.`,
		_const.ZhCN: `服务模块已被销毁或禁用。`,
	},
	"not_allow": {
		_const.RU:   `Мы не доверяем вам и поэтому этот сервис не доступен.`,
		_const.EN:   `We do not trust you and therefore this service is not available.`,
		_const.ZhCN: `我们不信任您，因此此服务不可用。`,
	},
	"lost_items": {
		_const.RU:   `Потери:`,
		_const.EN:   `Losses:`,
		_const.ZhCN: `损失:`,
	},
	"tax": {
		_const.RU:   `Налог:`,
		_const.EN:   `Tax:`,
		_const.ZhCN: `税收:`,
	},
	"result": {
		_const.RU:   `Результат:`,
		_const.EN:   `Result:`,
		_const.ZhCN: `结果:`,
	},
	"button_1": {
		_const.RU:   `Закрыть`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
	"button_2": {
		_const.RU:   `Переработать`,
		_const.EN:   `Recycle`,
		_const.ZhCN: `回收`,
	},
	"text_3": {
		_const.RU:   `Переработать`,
		_const.EN:   `Recycle`,
		_const.ZhCN: `回收`,
	},
}
