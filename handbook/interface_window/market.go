package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Market = map[string]map[string]string{
	"text_15": {
		_const.RU:   `Выберите категорию`,
		_const.EN:   `Select category`,
		_const.ZhCN: `选择类别`,
	},
	"text_16": {
		_const.RU:   `Нельзя выбросить/продать <span class="importantly">последний корпус</span>.`,
		_const.EN:   `You cannot drop/sell <span class="importantly">the last hull</span>.`,
		_const.ZhCN: `不能丢弃/出售<span class="importantly">最后一个船体</span>。`,
	},
	"text_17": {
		_const.RU:   `Предметы отправлены на <span class="importantly">склад</span> базы <span class="importantly">%BASE_NAME%</span>.`,
		_const.EN:   `Items sent to <span class="importantly">warehouse</span> base <span class="importantly">%BASE_NAME%</span>.`,
		_const.ZhCN: `物品已发送到<span class="importantly">仓库</span>基地<span class="importantly">%BASE_NAME%</span>。`,
	},
	// Buy
	"text_1": {
		_const.RU:   `Кол-во:`,
		_const.EN:   `Qty:`,
		_const.ZhCN: `数量:`,
	},
	"text_2": {
		_const.RU:   `Сr. за шт.`,
		_const.EN:   `Cr. per pce`,
		_const.ZhCN: `每件价格`,
	},
	"text_3": {
		_const.RU:   `Всего:`,
		_const.EN:   `Total:`,
		_const.ZhCN: `总计:`,
	},
	"button_1": {
		_const.RU:   `Отменить`,
		_const.EN:   `Cancel`,
		_const.ZhCN: `取消`,
	},
	"button_2": {
		_const.RU:   `Купить`,
		_const.EN:   `Buy`,
		_const.ZhCN: `购买`,
	},
	// Filter
	"cat_1": {
		_const.RU:   `Боеприпасы`,
		_const.EN:   `Ammunition`,
		_const.ZhCN: `弹药`,
	},
	"cat_2": {
		_const.RU:   `Оружие`,
		_const.EN:   `Weapon`,
		_const.ZhCN: `武器`,
	},
	"cat_3": {
		_const.RU:   `Корпуса`,
		_const.EN:   `Bodies`,
		_const.ZhCN: `船体`,
	},
	"cat_4": {
		_const.RU:   `Снаряжение`,
		_const.EN:   `Equipment`,
		_const.ZhCN: `装备`,
	},
	"cat_5": {
		_const.RU:   `Ресурсы`,
		_const.EN:   `Resources`,
		_const.ZhCN: `资源`,
	},
	"cat_6": {
		_const.RU:   `Чертежи`,
		_const.EN:   `Blueprints`,
		_const.ZhCN: `蓝图`,
	},
	"cat_7": {
		_const.RU:   `Ящики`,
		_const.EN:   `Boxes`,
		_const.ZhCN: `箱子`,
	},
	"cat_8": {
		_const.RU:   `Разное`,
		_const.EN:   `Miscellaneous`,
		_const.ZhCN: `杂项`,
	},
	"cat_9": {
		_const.RU:   `Товары`,
		_const.EN:   `Products`,
		_const.ZhCN: `商品`,
	},
	"cat_10": {
		_const.RU:   `Топливо`,
		_const.EN:   `Fuel`,
		_const.ZhCN: `燃料`,
	},
	"cat_11": {
		_const.RU:   `Дроны для шахт`,
		_const.EN:   `Drones for mines`,
		_const.ZhCN: `矿用无人机`,
	},
	"size_t1": {
		_const.RU:   `Малые`,
		_const.EN:   `Small`,
		_const.ZhCN: `小型`,
	},
	"size_t2": {
		_const.RU:   `Средние`,
		_const.EN:   `Medium`,
		_const.ZhCN: `中型`,
	},
	"size_t3": {
		_const.RU:   `Большие`,
		_const.EN:   `Large`,
		_const.ZhCN: `大型`,
	},
	"equip_t3": {
		_const.RU:   `Модификаторы`,
		_const.EN:   `Modifiers`,
		_const.ZhCN: `修改器`,
	},
	"body_t1": {
		_const.RU:   `Легкие`,
		_const.EN:   `Light`,
		_const.ZhCN: `轻型`,
	},
	"body_t2": {
		_const.RU:   `Средние`,
		_const.EN:   `Medium`,
		_const.ZhCN: `中型`,
	},
	"body_t3": {
		_const.RU:   `Тяжелые`,
		_const.EN:   `Heavy`,
		_const.ZhCN: `重型`,
	},
	"bp_t1": {
		_const.RU:   `Боеприпасы`,
		_const.EN:   `Ammunition`,
		_const.ZhCN: `弹药`,
	},
	"bp_t2": {
		_const.RU:   `Корпуса`,
		_const.EN:   `Bodies`,
		_const.ZhCN: `船体`,
	},
	"bp_t3": {
		_const.RU:   `Оружие`,
		_const.EN:   `Weapons`,
		_const.ZhCN: `武器`,
	},
	"bp_t4": {
		_const.RU:   `Снаряжение`,
		_const.EN:   `Equipments`,
		_const.ZhCN: `装备`,
	},
	"bp_t5": {
		_const.RU:   `Детали`,
		_const.EN:   `Details`,
		_const.ZhCN: `零件`,
	},
	"bp_t6": {
		_const.RU:   `Ящики`,
		_const.EN:   `Boxes`,
		_const.ZhCN: `箱子`,
	},
	"bp_t7": {
		_const.RU:   `Структуры`,
		_const.EN:   `Structures`,
		_const.ZhCN: `结构`,
	},
	"r_t1": {
		_const.RU:   `Ископаемые`,
		_const.EN:   `Fossils`,
		_const.ZhCN: `化石`,
	},
	"r_t2": {
		_const.RU:   `Сырье`,
		_const.EN:   `Raw`,
		_const.ZhCN: `原材料`,
	},
	"r_t3": {
		_const.RU:   `Детали`,
		_const.EN:   `Details`,
		_const.ZhCN: `零件`,
	},
	"r_t4": {
		_const.RU:   `Упакованные`,
		_const.EN:   `Packed`,
		_const.ZhCN: `打包`,
	},
	"a_t1": {
		_const.RU:   `Лазерные`,
		_const.EN:   `Laser`,
		_const.ZhCN: `激光`,
	},
	"a_t2": {
		_const.RU:   `Ракетные`,
		_const.EN:   `Missile`,
		_const.ZhCN: `导弹`,
	},
	"a_t3": {
		_const.RU:   `Баллистические`,
		_const.EN:   `Ballistic`,
		_const.ZhCN: `弹道`,
	},
	"w_t1": {
		_const.RU:   `Лазерное`,
		_const.EN:   `Laser`,
		_const.ZhCN: `激光`,
	},
	"w_t2": {
		_const.RU:   `Ракетное`,
		_const.EN:   `Missile`,
		_const.ZhCN: `导弹`,
	},
	"w_t3": {
		_const.RU:   `Баллистическое`,
		_const.EN:   `Ballistic`,
		_const.ZhCN: `弹道`,
	},
	// Market
	"window_name": {
		_const.RU:   `Рынок`,
		_const.EN:   `Market`,
		_const.ZhCN: `市场`,
	},
	"error_1": {
		_const.RU:   `Модуль сервиса на безе уничтожен или отключен.`,
		_const.EN:   `The service module based on the base has been destroyed or disabled.`,
		_const.ZhCN: `基地上的服务模块已被销毁或禁用。`,
	},
	"sub_button_1": {
		_const.RU:   `Подробно`,
		_const.EN:   `Details`,
		_const.ZhCN: `详情`,
	},
	"sub_button_2": {
		_const.RU:   `Продать`,
		_const.EN:   `Sell`,
		_const.ZhCN: `出售`,
	},
	"sub_button_3": {
		_const.RU:   `Купить`,
		_const.EN:   `Buy`,
		_const.ZhCN: `购买`,
	},
	"sub_button_4": {
		_const.RU:   `Автопилот`,
		_const.EN:   `Autopilot`,
		_const.ZhCN: `自动驾驶`,
	},
	"sub_button_5": {
		_const.RU:   `Проложить маршрут`,
		_const.EN:   `Get directions`,
		_const.ZhCN: `获取路线`,
	},
	"sub_button_6": {
		_const.RU:   `Закрыть`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
	"text_4": {
		_const.RU:   `Мой баланс:`,
		_const.EN:   `My balance:`,
		_const.ZhCN: `我的余额:`,
	},
	"tab_1": {
		_const.RU:   `Рынок`,
		_const.EN:   `Market`,
		_const.ZhCN: `市场`,
	},
	"tab_2": {
		_const.RU:   `Мои запросы/предложения`,
		_const.EN:   `My requests/suggestions`,
		_const.ZhCN: `我的请求/报价`,
	},
	"tab_3": {
		_const.RU:   `История сделок`,
		_const.EN:   `Transaction history`,
		_const.ZhCN: `交易历史`,
	},
	"button_3": {
		_const.RU:   `Купить`,
		_const.EN:   `Buy`,
		_const.ZhCN: `购买`,
	},
	"dist_1": {
		_const.RU:   `Недостижимо`,
		_const.EN:   `Unattainable`,
		_const.ZhCN: `无法到达`,
	},
	"dist_2": {
		_const.RU:   `База`,
		_const.EN:   `Base`,
		_const.ZhCN: `基地`,
	},
	"dist_3": {
		_const.RU:   `Сектор`,
		_const.EN:   `Sector`,
		_const.ZhCN: `区域`,
	},
	"head_table_1": {
		_const.RU:   `Продавцы`,
		_const.EN:   `Sellers`,
		_const.ZhCN: `卖家`,
	},
	"head_col_table_1": {
		_const.RU:   `Расстояние`,
		_const.EN:   `Distance`,
		_const.ZhCN: `距离`,
	},
	"head_col_table_2": {
		_const.RU:   `Кол-во`,
		_const.EN:   `Qty`,
		_const.ZhCN: `数量`,
	},
	"head_col_table_3": {
		_const.RU:   `Цена`,
		_const.EN:   `Price`,
		_const.ZhCN: `价格`,
	},
	"head_col_table_4": {
		_const.RU:   `Тип`,
		_const.EN:   `Type`,
		_const.ZhCN: `类型`,
	},
	"head_col_table_5": {
		_const.RU:   `Название`,
		_const.EN:   `Name`,
		_const.ZhCN: `名称`,
	},
	"head_col_table_6": {
		_const.RU:   `Место`,
		_const.EN:   `Place`,
		_const.ZhCN: `地点`,
	},
	"head_table_2": {
		_const.RU:   `Покупатели`,
		_const.EN:   `Buyers`,
		_const.ZhCN: `买家`,
	},
	"head_col_table_7": {
		_const.RU:   `Расстояние`,
		_const.EN:   `Distance`,
		_const.ZhCN: `距离`,
	},
	"head_col_table_8": {
		_const.RU:   `Кол-во`,
		_const.EN:   `Qty`,
		_const.ZhCN: `数量`,
	},
	"head_col_table_9": {
		_const.RU:   `Цена`,
		_const.EN:   `Price`,
		_const.ZhCN: `价格`,
	},
	"head_col_table_10": {
		_const.RU:   `Тип`,
		_const.EN:   `Type`,
		_const.ZhCN: `类型`,
	},
	"head_col_table_11": {
		_const.RU:   `Название`,
		_const.EN:   `Name`,
		_const.ZhCN: `名称`,
	},
	"head_col_table_12": {
		_const.RU:   `Место`,
		_const.EN:   `Place`,
		_const.ZhCN: `地点`,
	},
	"head_col_table_13": {
		_const.RU:   `Расстояние`,
		_const.EN:   `Distance`,
		_const.ZhCN: `距离`,
	},
	"head_col_table_14": {
		_const.RU:   `Кол-во`,
		_const.EN:   `Qty`,
		_const.ZhCN: `数量`,
	},
	"head_col_table_15": {
		_const.RU:   `Цена`,
		_const.EN:   `Price`,
		_const.ZhCN: `价格`,
	},
	"head_col_table_16": {
		_const.RU:   `Тип`,
		_const.EN:   `Type`,
		_const.ZhCN: `类型`,
	},
	"head_col_table_17": {
		_const.RU:   `Название`,
		_const.EN:   `Name`,
		_const.ZhCN: `名称`,
	},
	"head_col_table_18": {
		_const.RU:   `Место`,
		_const.EN:   `Place`,
		_const.ZhCN: `地点`,
	},
	"head_col_table_19": {
		_const.RU:   `Тип сделки`,
		_const.EN:   `Transaction type`,
		_const.ZhCN: `交易类型`,
	},
	"head_col_table_20": {
		_const.RU:   `Общая стоимость`,
		_const.EN:   `Total cost`,
		_const.ZhCN: `总成本`,
	},
	"head_col_table_21": {
		_const.RU:   `Название`,
		_const.EN:   `Name`,
		_const.ZhCN: `名称`,
	},
	"head_col_table_22": {
		_const.RU:   `Кол-во`,
		_const.EN:   `Qty`,
		_const.ZhCN: `数量`,
	},
	"head_col_table_23": {
		_const.RU:   `Тип сделки`,
		_const.EN:   `Transaction type`,
		_const.ZhCN: `交易类型`,
	},
	"head_col_table_24": {
		_const.RU:   `Стоимость`,
		_const.EN:   `Total cost`,
		_const.ZhCN: `总成本`,
	},
	"head_col_table_25": {
		_const.RU:   `Место`,
		_const.EN:   `Place`,
		_const.ZhCN: `地点`,
	},
	"deal_t1": {
		_const.RU:   `Продажа`,
		_const.EN:   `Sell`,
		_const.ZhCN: `出售`,
	},
	"deal_t2": {
		_const.RU:   `Покупка`,
		_const.EN:   `Buy`,
		_const.ZhCN: `购买`,
	},
	"text_14": {
		_const.RU:   `~ цена: <span class="importantly">%credits_count%</span> cr.`,
		_const.EN:   `~ price: <span class="importantly">%credits_count%</span> cr.`,
		_const.ZhCN: `~ 价格: <span class="importantly">%credits_count%</span> cr.`,
	},
	// NewBuy
	"text_5": {
		_const.RU:   `Кол-во:`,
		_const.EN:   `Qty:`,
		_const.ZhCN: `数量:`,
	},
	"text_6": {
		_const.RU:   `Сr. за шт.`,
		_const.EN:   `Cr. per pce`,
		_const.ZhCN: `每件价格`,
	},
	"text_7": {
		_const.RU:   `Всего:`,
		_const.EN:   `Total:`,
		_const.ZhCN: `总计:`,
	},
	"button_4": {
		_const.RU:   `Отменить`,
		_const.EN:   `Cancel`,
		_const.ZhCN: `取消`,
	},
	"button_5": {
		_const.RU:   `Разместить`,
		_const.EN:   `Place an order`,
		_const.ZhCN: `下单`,
	},
	// Sell
	"text_8": {
		_const.RU:   `Кол-во:`,
		_const.EN:   `Qty:`,
		_const.ZhCN: `数量:`,
	},
	"text_9": {
		_const.RU:   `Сr. за шт.`,
		_const.EN:   `Cr. per pce`,
		_const.ZhCN: `每件价格`,
	},
	"text_10": {
		_const.RU:   `Всего:`,
		_const.EN:   `Total:`,
		_const.ZhCN: `总计:`,
	},
	"text_11": {
		_const.RU:   `Налог (%tax%%):`,
		_const.EN:   `Tax (%tax%%):`,
		_const.ZhCN: `税 (%tax%%):`,
	},
	"text_12": {
		_const.RU:   `Вы получите:`,
		_const.EN:   `You'll get:`,
		_const.ZhCN: `你将获得:`,
	},
	"button_6": {
		_const.RU:   `Отменить`,
		_const.EN:   `Cancel`,
		_const.ZhCN: `取消`,
	},
	"button_7": {
		_const.RU:   `Продать`,
		_const.EN:   `Sell`,
		_const.ZhCN: `出售`,
	},
	"text_13": {
		_const.RU:   `На базе где размещен заказ нет предметов для продажи.`,
		_const.EN:   `There are no items for sale at the base where the order was placed.`,
		_const.ZhCN: `下单的基地没有可出售的物品。`,
	},
	"button_8": {
		_const.RU:   `Назад`,
		_const.EN:   `Back`,
		_const.ZhCN: `返回`,
	},
	"region_head": {
		_const.RU:   `Область:`,
		_const.EN:   ``,
		_const.ZhCN: ``,
	},
	"region_option_0": {
		_const.RU:   `Регион`,
		_const.EN:   ``,
		_const.ZhCN: ``,
	},
	"region_option_1": {
		_const.RU:   `Сектор`,
		_const.EN:   ``,
		_const.ZhCN: ``,
	},
	"region_option_2": {
		_const.RU:   `База`,
		_const.EN:   ``,
		_const.ZhCN: ``,
	},
}
