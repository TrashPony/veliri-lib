package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Laboratory = map[string]map[string]string{
	"window_name": {
		_const.RU:   `Лаборатория`,
		_const.EN:   `Laboratory`,
		_const.ZhCN: `实验室`,
	},
	"not_allow": {
		_const.RU:   `Мы не доверяем вам и поэтому этот сервис не доступен.`,
		_const.EN:   `We do not trust you and therefore this service is not available.`,
		_const.ZhCN: `我们不信任您，因此此服务不可用。`,
	},
	"error_1": {
		_const.RU:   `Нельзя улучшить уже улучшенный чертеж.`,
		_const.EN:   `Cannot upgrade an already upgraded blueprint.`,
		_const.ZhCN: `无法升级已升级的蓝图。`,
	},
	"error_2": {
		_const.RU:   `Чертеж должен быть на <span class="importantly">Складе</span>`,
		_const.EN:   `Blueprint must be in the <span class="importantly">Storage</span>`,
		_const.ZhCN: `蓝图必须在<span class="importantly">仓库</span>中`,
	},
	"error_3": {
		_const.RU:   `Попытка улучшения чертежа, завершилась неудачей.`,
		_const.EN:   `Blueprint upgrade attempt failed.`,
		_const.ZhCN: `蓝图升级尝试失败。`,
	},
	"complete_message": {
		_const.RU:   `Чертеж был улучшен и отправлен на склад.`,
		_const.EN:   `Blueprint was upgraded and sent to storage.`,
		_const.ZhCN: `蓝图已升级并发送至仓库。`,
	},
	"head_1": {
		_const.RU:   `ЧЕРТЕЖ`,
		_const.EN:   `BLUEPRINT`,
		_const.ZhCN: `蓝图`,
	},
	"head_2": {
		_const.RU:   `СПЕЦИАЛИЗАЦИЯ`,
		_const.EN:   `SPECIALIZATION`,
		_const.ZhCN: `专业化`,
	},
	"head_3": {
		_const.RU:   `ВАРИАНТЫ`,
		_const.EN:   `VARIANTS`,
		_const.ZhCN: `变体`,
	},
	"text_1": {
		_const.RU:   `У этого чертежа нет возможности улучшения`,
		_const.EN:   `This blueprint cannot be upgraded`,
		_const.ZhCN: `此蓝图无法升级`,
	},
	"text_2": {
		_const.RU:   `Не хватает данных на складе`,
		_const.EN:   `Insufficient data in storage`,
		_const.ZhCN: `仓库数据不足`,
	},
	"text_3": {
		_const.RU:   `Лаборатория не специализируется на этой технологии`,
		_const.EN:   `Laboratory does not specialize in this technology`,
		_const.ZhCN: `实验室不专攻此技术`,
	},
	"text_4": {
		_const.RU:   `Не выбран чертеж`,
		_const.EN:   `No blueprint selected`,
		_const.ZhCN: `未选择蓝图`,
	},
	"button_1": {
		_const.RU:   `Улучшить (%percent%%)`,
		_const.EN:   `Upgrade (%percent%%)`,
		_const.ZhCN: `升级(%percent%%)`,
	},
}
