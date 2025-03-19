package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var SendCredits = map[string]map[string]string{
	"window_name": {
		_const.RU:   `Кредиты`,
		_const.EN:   `Credits`,
		_const.ZhCN: `信用点`,
	},
	"transfer": {
		_const.RU:   `Перевод`,
		_const.EN:   `Transfer`,
		_const.ZhCN: `转账`,
	},
	"history": {
		_const.RU:   `История`,
		_const.EN:   `History`,
		_const.ZhCN: `历史`,
	},
	"text_1": {
		_const.RU:   `Откуда:`,
		_const.EN:   `Src:`,
		_const.ZhCN: `来源:`,
	},
	"text_2": {
		_const.RU:   `Куда:`,
		_const.EN:   `Dst:`,
		_const.ZhCN: `目标:`,
	},
	"option_1": {
		_const.RU:   `Мой счет`,
		_const.EN:   `My deposit`,
		_const.ZhCN: `我的账户`,
	},
	"player": {
		_const.RU:   `Игрок`,
		_const.EN:   `Player`,
		_const.ZhCN: `玩家`,
	},
	"cluster": {
		_const.RU:   `Кластер`,
		_const.EN:   `Cluster`,
		_const.ZhCN: `集群`,
	},
	"placeholder_1": {
		_const.RU:   `Имя`,
		_const.EN:   `Name`,
		_const.ZhCN: `名称`,
	},
	"button_1": {
		_const.RU:   `Найти`,
		_const.EN:   `Search`,
		_const.ZhCN: `搜索`,
	},
	"nothing": {
		_const.RU:   `Ничего не найдено.`,
		_const.EN:   `Nothing found.`,
		_const.ZhCN: `未找到任何内容。`,
	},
	"text_3": {
		_const.RU:   `Сколько:`,
		_const.EN:   `How many:`,
		_const.ZhCN: `数量:`,
	},
	"placeholder_2": {
		_const.RU:   `причина (необяз.)`,
		_const.EN:   `reason (optional)`,
		_const.ZhCN: `原因 (可选)`,
	},
	"button_2": {
		_const.RU:   `Перевести`,
		_const.EN:   `Send`,
		_const.ZhCN: `发送`,
	},
	"text_4": {
		_const.RU:   `Операция`,
		_const.EN:   `Operation`,
		_const.ZhCN: `操作`,
	},
	"text_5": {
		_const.RU:   `Причина`,
		_const.EN:   `Cause`,
		_const.ZhCN: `原因`,
	},
	"on_1": {
		_const.RU:   `Перевод от кластера <span class="importantly">%name%</span>`,
		_const.EN:   `Transfer from cluster <span class="importantly">%name%</span>`,
		_const.ZhCN: `来自集群 <span class="importantly">%name%</span> 的转账`,
	},
	"on_2": {
		_const.RU:   `Перевод кластеру <span class="importantly">%name%</span>`,
		_const.EN:   `Transfer to cluster <span class="importantly">%name%</span>`,
		_const.ZhCN: `转账到集群 <span class="importantly">%name%</span>`,
	},
	"on_3": {
		_const.RU:   `Перевод от <span class="importantly">%name%</span>`,
		_const.EN:   `Translation from <span class="importantly">%name%</span>`,
		_const.ZhCN: `来自 <span class="importantly">%name%</span> 的转账`,
	},
	"on_4": {
		_const.RU:   `Перевод <span class="importantly">%name%</span>`,
		_const.EN:   `Translation <span class="importantly">%name%</span>`,
		_const.ZhCN: `转账给 <span class="importantly">%name%</span>`,
	},
	"on_5": {
		_const.RU:   `Налог с продажи`,
		_const.EN:   `Sales tax`,
		_const.ZhCN: `销售税`,
	},
	"on_6": {
		_const.RU:   `Стартовый капитал`,
		_const.EN:   `Start-up capital`,
		_const.ZhCN: `启动资金`,
	},
	"on_7": {
		_const.RU:   `Продажа товара`,
		_const.EN:   `Sale of goods`,
		_const.ZhCN: `商品销售`,
	},
	"on_8": {
		_const.RU:   `Покупка товара`,
		_const.EN:   `Purchase of goods`,
		_const.ZhCN: `商品购买`,
	},
	"on_9": {
		_const.RU:   `Продажа на рынке`,
		_const.EN:   `Selling on the market`,
		_const.ZhCN: `市场销售`,
	},
	"on_10": {
		_const.RU:   `Покупка на рынке`,
		_const.EN:   `Purchase at the market`,
		_const.ZhCN: `市场购买`,
	},
	"on_11": {
		_const.RU:   `Размещение заказа`,
		_const.EN:   `Placing an order`,
		_const.ZhCN: `下单`,
	},
	"on_12": {
		_const.RU:   `Отмена заказа`,
		_const.EN:   `Cancellations`,
		_const.ZhCN: `取消订单`,
	},
	"on_13": {
		_const.RU:   `Выполнение миссии`,
		_const.EN:   `Completing the mission`,
		_const.ZhCN: `完成任务`,
	},
	"on_14": {
		_const.RU:   `Покупка покраски`,
		_const.EN:   `Buying skin`,
		_const.ZhCN: `购买涂装`,
	},
	"on_15": {
		_const.RU:   `Грабеж`,
		_const.EN:   `Robbery`,
		_const.ZhCN: `抢劫`,
	},
	"on_16": {
		_const.RU:   `Починка транспорта`,
		_const.EN:   `Vehicle repair`,
		_const.ZhCN: `载具修理`,
	},
	"on_17": {
		_const.RU:   `Починка предмета`,
		_const.EN:   `Repairing an item`,
		_const.ZhCN: `物品修理`,
	},
	"on_18": {
		_const.RU:   `Покупка в фр. магазине`,
		_const.EN:   `Purchasing from a faction store`,
		_const.ZhCN: `从阵营商店购买`,
	},
	"on_19": {
		_const.RU:   `Поиск`,
		_const.EN:   `Search`,
		_const.ZhCN: `搜索`,
	},
	"on_20": {
		_const.RU:   `Устранение преступника`,
		_const.EN:   `Elimination of the criminal`,
		_const.ZhCN: `消灭罪犯`,
	},
	"on_21": {
		_const.RU:   `Устранение бойца в войне`,
		_const.EN:   `Eliminating a fighter in a war`,
		_const.ZhCN: `消灭战争中的战士`,
	},
	"on_22": {
		_const.RU:   `Устранение сил APD`,
		_const.EN:   `Elimination of APD forces`,
		_const.ZhCN: `消灭APD部队`,
	},
	"on_23": {
		_const.RU:   `Покупка боеприпасов`,
		_const.EN:   `Buying ammunition`,
		_const.ZhCN: `购买弹药`,
	},
	"on_24": {
		_const.RU:   `Заправка транспорта`,
		_const.EN:   `Refueling vehicles`,
		_const.ZhCN: `载具加油`,
	},
}
