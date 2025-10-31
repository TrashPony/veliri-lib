package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var SellChipsModal = map[string]map[string]string{
	"window_name": {
		_const.RU:   `Переработка чипов`,
		_const.EN:   `Chip Recycling`,
		_const.ZhCN: `芯片回收`,
	},
	"text_1": {
		_const.RU:   `Доступные чипы:`,
		_const.EN:   `Available chips:`,
		_const.ZhCN: `可用芯片：`,
	},
	"text_2": {
		_const.RU:   `На переработку:`,
		_const.EN:   `For recycling:`,
		_const.ZhCN: `回收数量：`,
	},
	"text_3": {
		_const.RU:   `Получишь опыта:`,
		_const.EN:   `Experience gained:`,
		_const.ZhCN: `获得经验：`,
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
}

var ChangeFractionModal = map[string]map[string]string{
	"window_name": {
		_const.RU:   `Смена фракции`,
		_const.EN:   `Faction Change`,
		_const.ZhCN: `派系更换`,
	},
	"text_1": {
		_const.RU:   `Выберите фракцию к которой хотите присоеденится.`,
		_const.EN:   `Select the faction you want to join.`,
		_const.ZhCN: `选择您想要加入的派系。`,
	},
	"text_2": {
		_const.RU:   `Цена:`,
		_const.EN:   `Price:`,
		_const.ZhCN: `价格：`,
	},
	"text_3": {
		_const.RU:   `Присоеденится?`,
		_const.EN:   `Join?`,
		_const.ZhCN: `确认加入？`,
	},
	"text_4": {
		_const.RU:   `Недостаточно кредитов на складе.`,
		_const.EN:   `Not enough credits in storage.`,
		_const.ZhCN: `仓库信用点不足。`,
	},
	"button_1": {
		_const.RU:   `Закрыть`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
	"button_2": {
		_const.RU:   `Присоеденится`,
		_const.EN:   `Join`,
		_const.ZhCN: `加入`,
	},
}

var ClearRelationModal = map[string]map[string]string{
	"window_name": {
		_const.RU:   `Очистка репутации`,
		_const.EN:   `Reputation Reset`,
		_const.ZhCN: `声望重置`,
	},
	"text_1": {
		_const.RU: `Репутация со всеми фракциями изменится на
					"<span class="highlight">нейтрально</span>"
					(<span class="highlight">0:0</span>).`,
		_const.EN: `Reputation with all factions will change to
					"<span class="highlight">neutral</span>"
					(<span class="highlight">0:0</span>).`,
		_const.ZhCN: `与所有派系的声望将变为
					"<span class="highlight">中立</span>"
					(<span class="highlight">0:0</span>)。`,
	},
	"text_2": {
		_const.RU:   `Цена:`,
		_const.EN:   `Price:`,
		_const.ZhCN: `价格：`,
	},
	"text_3": {
		_const.RU:   `Обнулить?`,
		_const.EN:   `Reset?`,
		_const.ZhCN: `确认重置？`,
	},
	"text_4": {
		_const.RU:   `Недостаточно кредитов на складе.`,
		_const.EN:   `Not enough credits in storage.`,
		_const.ZhCN: `仓库信用点不足。`,
	},
	"button_1": {
		_const.RU:   `Закрыть`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
	"button_2": {
		_const.RU:   `Обнулить`,
		_const.EN:   `Reset`,
		_const.ZhCN: `重置`,
	},
}
