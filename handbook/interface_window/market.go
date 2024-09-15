package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Market = map[string]map[string]string{
	// Buy
	"text_1": {
		_const.RU: `Кол-во:`,
		_const.EN: `Qty:`,
	},
	"text_2": {
		_const.RU: `Сr. за шт.`,
		_const.EN: `Cr. per pce`,
	},
	"text_3": {
		_const.RU: `Всего:`,
		_const.EN: `Total:`,
	},
	"button_1": {
		_const.RU: `Отменить`,
		_const.EN: `Cancel`,
	},
	"button_2": {
		_const.RU: `Купить`,
		_const.EN: `Buy`,
	},
	// Filter
	"cat_1": {
		_const.RU: `Боеприпасы`,
		_const.EN: `Ammunition`,
	},
	"cat_2": {
		_const.RU: `Оружие`,
		_const.EN: `Weapon`,
	},
	"cat_3": {
		_const.RU: `Корпуса`,
		_const.EN: `Bodies`,
	},
	"cat_4": {
		_const.RU: `Снаряжение`,
		_const.EN: `Equipment`,
	},
	"cat_5": {
		_const.RU: `Ресурсы`,
		_const.EN: `Resources`,
	},
	"cat_6": {
		_const.RU: `Чертежи`,
		_const.EN: `Blueprints`,
	},
	"cat_7": {
		_const.RU: `Ящики`,
		_const.EN: `Boxes`,
	},
	"cat_8": {
		_const.RU: `Хлам`,
		_const.EN: `Rubbish`,
	},
	"cat_9": {
		_const.RU: `Товары`,
		_const.EN: `Products`,
	},
	"cat_10": {
		_const.RU: `Топливо`,
		_const.EN: `Fuel`,
	},
	"size_t1": {
		_const.RU: `Малые`,
		_const.EN: `Small`,
	},
	"size_t2": {
		_const.RU: `Средние`,
		_const.EN: `Medium`,
	},
	"size_t3": {
		_const.RU: `Большие`,
		_const.EN: `Large`,
	},
	"equip_t3": {
		_const.RU: `Модификаторы`,
		_const.EN: `Modifiers`,
	},
	"body_t1": {
		_const.RU: `Легкие`,
		_const.EN: `Light`,
	},
	"body_t2": {
		_const.RU: `Средние`,
		_const.EN: `Medium`,
	},
	"body_t3": {
		_const.RU: `Тяжелые`,
		_const.EN: `Heavy`,
	},
	"bp_t1": {
		_const.RU: `Боеприпасы`,
		_const.EN: `Ammunition`,
	},
	"bp_t2": {
		_const.RU: `Корпуса`,
		_const.EN: `Bodies`,
	},
	"bp_t3": {
		_const.RU: `Оружие`,
		_const.EN: `Weapons`,
	},
	"bp_t4": {
		_const.RU: `Снаряжение`,
		_const.EN: `Equipments`,
	},
	"bp_t5": {
		_const.RU: `Детали`,
		_const.EN: `Details`,
	},
	"bp_t6": {
		_const.RU: `Ящики`,
		_const.EN: `Boxes`,
	},
	"bp_t7": {
		_const.RU: `Структуры`,
		_const.EN: `Structures`,
	},
	"r_t1": {
		_const.RU: `Ископаемые`,
		_const.EN: `Fossils`,
	},
	"r_t2": {
		_const.RU: `Сырье`,
		_const.EN: `Raw`,
	},
	"r_t3": {
		_const.RU: `Детали`,
		_const.EN: `Details`,
	},
	"a_t1": {
		_const.RU: `Лазерные`,
		_const.EN: `Laser`,
	},
	"a_t2": {
		_const.RU: `Ракетные`,
		_const.EN: `Missile`,
	},
	"a_t3": {
		_const.RU: `Баллистические`,
		_const.EN: `Ballistic`,
	},
	"w_t1": {
		_const.RU: `Лазерное`,
		_const.EN: `Laser`,
	},
	"w_t2": {
		_const.RU: `Ракетное`,
		_const.EN: `Missile`,
	},
	"w_t3": {
		_const.RU: `Баллистическое`,
		_const.EN: `Ballistic`,
	},
	// Market
	"window_name": {
		_const.RU: `Рынок`,
		_const.EN: `Market`,
	},
	"error_1": {
		_const.RU: `Модуль сервиса на безе уничтожен или отключен.`,
		_const.EN: `The service module based on the base has been destroyed or disabled.`,
	},
	"sub_button_1": {
		_const.RU: `Подробно`,
		_const.EN: `Details`,
	},
	"sub_button_2": {
		_const.RU: `Продать`,
		_const.EN: `Sell`,
	},
	"sub_button_3": {
		_const.RU: `Купить`,
		_const.EN: `Buy`,
	},
	"sub_button_4": {
		_const.RU: `Автопилот`,
		_const.EN: `Autopilot`,
	},
	"sub_button_5": {
		_const.RU: `Проложить маршрут`,
		_const.EN: `Get directions`,
	},
	"sub_button_6": {
		_const.RU: `Закрыть`,
		_const.EN: `Close`,
	},
	"text_4": {
		_const.RU: `Мой баланс:`,
		_const.EN: `My balance:`,
	},
	"tab_1": {
		_const.RU: `Рынок`,
		_const.EN: `Market`,
	},
	"tab_2": {
		_const.RU: `Мои запросы/предложения`,
		_const.EN: `My requests/suggestions`,
	},
	"tab_3": {
		_const.RU: `История сделок`,
		_const.EN: `Transaction history`,
	},
	"button_3": {
		_const.RU: `Купить`,
		_const.EN: `Buy`,
	},
	"dist_1": {
		_const.RU: `Недостижимо`,
		_const.EN: `Unattainable`,
	},
	"dist_2": {
		_const.RU: `База`,
		_const.EN: `Base`,
	},
	"dist_3": {
		_const.RU: `Сектор`,
		_const.EN: `Sector`,
	},
	"head_table_1": {
		_const.RU: `Продавцы`,
		_const.EN: `Sellers`,
	},
	"head_col_table_1": {
		_const.RU: `Расстояние`,
		_const.EN: `Distance`,
	},
	"head_col_table_2": {
		_const.RU: `Кол-во`,
		_const.EN: `Qty`,
	},
	"head_col_table_3": {
		_const.RU: `Цена`,
		_const.EN: `Price`,
	},
	"head_col_table_4": {
		_const.RU: `Тип`,
		_const.EN: `Type`,
	},
	"head_col_table_5": {
		_const.RU: `Название`,
		_const.EN: `Name`,
	},
	"head_col_table_6": {
		_const.RU: `Место`,
		_const.EN: `Place`,
	},
	"head_table_2": {
		_const.RU: `Покупатели`,
		_const.EN: `Buyers`,
	},
	"head_col_table_7": {
		_const.RU: `Расстояние`,
		_const.EN: `Distance`,
	},
	"head_col_table_8": {
		_const.RU: `Кол-во`,
		_const.EN: `Qty`,
	},
	"head_col_table_9": {
		_const.RU: `Цена`,
		_const.EN: `Price`,
	},
	"head_col_table_10": {
		_const.RU: `Тип`,
		_const.EN: `Type`,
	},
	"head_col_table_11": {
		_const.RU: `Название`,
		_const.EN: `Name`,
	},
	"head_col_table_12": {
		_const.RU: `Место`,
		_const.EN: `Place`,
	},
	"head_col_table_13": {
		_const.RU: `Расстояние`,
		_const.EN: `Distance`,
	},
	"head_col_table_14": {
		_const.RU: `Кол-во`,
		_const.EN: `Qty`,
	},
	"head_col_table_15": {
		_const.RU: `Цена`,
		_const.EN: `Price`,
	},
	"head_col_table_16": {
		_const.RU: `Тип`,
		_const.EN: `Type`,
	},
	"head_col_table_17": {
		_const.RU: `Название`,
		_const.EN: `Name`,
	},
	"head_col_table_18": {
		_const.RU: `Место`,
		_const.EN: `Place`,
	},
	"head_col_table_19": {
		_const.RU: `Тип сделки`,
		_const.EN: `Transaction type`,
	},
	"head_col_table_20": {
		_const.RU: `Общая стоимость`,
		_const.EN: `Total cost`,
	},
	"head_col_table_21": {
		_const.RU: `Название`,
		_const.EN: `Name`,
	},
	"head_col_table_22": {
		_const.RU: `Кол-во`,
		_const.EN: `Qty`,
	},
	"head_col_table_23": {
		_const.RU: `Тип сделки`,
		_const.EN: `Transaction type`,
	},
	"head_col_table_24": {
		_const.RU: `Стоимость`,
		_const.EN: `Total cost`,
	},
	"head_col_table_25": {
		_const.RU: `Место`,
		_const.EN: `Place`,
	},
	"deal_t1": {
		_const.RU: `Продажа`,
		_const.EN: `Sell`,
	},
	"deal_t2": {
		_const.RU: `Покупка`,
		_const.EN: `Buy`,
	},
	"text_14": {
		_const.RU: `~ цена: <span class="importantly">%credits_count%</span> cr.`,
		_const.EN: `~ price: <span class="importantly">%credits_count%</span> cr.`,
	},
	// NewBuy
	"text_5": {
		_const.RU: `Кол-во:`,
		_const.EN: `Qty:`,
	},
	"text_6": {
		_const.RU: `Сr. за шт.`,
		_const.EN: `Cr. per pce`,
	},
	"text_7": {
		_const.RU: `Всего:`,
		_const.EN: `Total:`,
	},
	"button_4": {
		_const.RU: `Отменить`,
		_const.EN: `Cancel`,
	},
	"button_5": {
		_const.RU: `Разместить`,
		_const.EN: `Place an order`,
	},
	// Sell
	"text_8": {
		_const.RU: `Кол-во:`,
		_const.EN: `Qty:`,
	},
	"text_9": {
		_const.RU: `Сr. за шт.`,
		_const.EN: `Cr. per pce`,
	},
	"text_10": {
		_const.RU: `Всего:`,
		_const.EN: `Total:`,
	},
	"text_11": {
		_const.RU: `Налог (%tax%%):`,
		_const.EN: `Tax (%tax%%):`,
	},
	"text_12": {
		_const.RU: `Вы получите:`,
		_const.EN: `You'll get:`,
	},
	"button_6": {
		_const.RU: `Отменить`,
		_const.EN: `Cancel`,
	},
	"button_7": {
		_const.RU: `Продать`,
		_const.EN: `Sell`,
	},
	"text_13": {
		_const.RU: `На базе где размещен заказ нет предметов для продажи.`,
		_const.EN: `There are no items for sale at the base where the order was placed.`,
	},
	"button_8": {
		_const.RU: `Назад`,
		_const.EN: `Back`,
	},
}
