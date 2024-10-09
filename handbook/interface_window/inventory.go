package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Inventory = map[string]map[string]string{
	// ArmModal
	"text_1": {
		_const.RU: `за %count%:`,
		_const.EN: `for %count%:`,
	},
	"text_2": {
		_const.RU: `Всего:`,
		_const.EN: `Total:`,
	},
	"text_3": {
		_const.RU: `Кол-во:`,
		_const.EN: `Qty:`,
	},
	"button_1": {
		_const.RU: `Закрыть`,
		_const.EN: `Close`,
	},
	"button_2": {
		_const.RU: `Зарядить`,
		_const.EN: `Arm`,
	},
	// BackgroundItemCell
	// CellSubMenu
	"button_3": {
		_const.RU: `Установить оружие`,
		_const.EN: `Install weapon`,
	},
	"button_4": {
		_const.RU: `Установить снаряды`,
		_const.EN: `Install ammo`,
	},
	"button_5": {
		_const.RU: `Активировать корпус`,
		_const.EN: `Activate the body`,
	},
	"button_6": {
		_const.RU: `Заправить корпус`,
		_const.EN: `Fill body`,
	},
	"button_7": {
		_const.RU: `Подробнее`,
		_const.EN: `More details`,
	},
	"button_8": {
		_const.RU: `Разделить`,
		_const.EN: `Split`,
	},
	"button_9": {
		_const.RU: `Выбросить`,
		_const.EN: `Throw`,
	},
	"button_10": {
		_const.RU: `Уничтожить`,
		_const.EN: `Destroy`,
	},
	"button_11": {
		_const.RU: `Установить`,
		_const.EN: `Install`,
	},
	"button_12": {
		_const.RU: `Продать`,
		_const.EN: `Sell`,
	},
	"button_13": {
		_const.RU: `Починить`,
		_const.EN: `Repair`,
	},
	"button_14": {
		_const.RU: `Снять`,
		_const.EN: `Take off`,
	},
	"button_15": {
		_const.RU: `В трюм`,
		_const.EN: `To hold`,
	},
	"button_16": {
		_const.RU: `На склад`,
		_const.EN: `To storage`,
	},
	"button_17": {
		_const.RU: `Произвести`,
		_const.EN: `Produce`,
	},
	"button_18": {
		_const.RU: `Отменить`,
		_const.EN: `Cancel`,
	},
	"text_4": {
		_const.RU: `Кол-во:`,
		_const.EN: `Qty:`,
	},
	"text_5": {
		_const.RU: `Сr. за шт.`,
		_const.EN: `Cr. per pce`,
	},
	"text_6": {
		_const.RU: `Всего:`,
		_const.EN: `Total:`,
	},
	"text_7": {
		_const.RU: `Налог (%tax%%)`,
		_const.EN: `Tax (%tax%%)`,
	},
	"text_8": {
		_const.RU: `Вы получите:`,
		_const.EN: `You'll get:`,
	},
	"button_19": {
		_const.RU: `Отмена`,
		_const.EN: `Cancel`,
	},
	"button_20": {
		_const.RU: `Продать`,
		_const.EN: `Sell`,
	},
	"button_21": {
		_const.RU: `Отмена`,
		_const.EN: `Cancel`,
	},
	"button_22": {
		_const.RU: `Разделить`,
		_const.EN: `Divide`,
	},
	// FillUpModal
	"text_9": {
		_const.RU: `Сr. за шт.`,
		_const.EN: `Cr. per pce`,
	},
	"text_10": {
		_const.RU: `Всего:`,
		_const.EN: `Total:`,
	},
	"text_11": {
		_const.RU: `Кол-во:`,
		_const.EN: `Qty:`,
	},
	"button_23": {
		_const.RU: `Закрыть`,
		_const.EN: `Close`,
	},
	"button_24": {
		_const.RU: `Заправить`,
		_const.EN: `Refuel`,
	},
	// Inventory
	"window_name": {
		_const.RU: `Инвентарь`,
		_const.EN: `Inventory`,
	},
	"text_12": {
		_const.RU: `Склад заблокирован. Продлите аренду офиса или восстановите модуль базы.`,
		_const.EN: `The warehouse is locked. Please extend the office rental or restore the database module.`,
	},
	"text_13": {
		_const.RU: `Нет прав.`,
		_const.EN: `You have no rights.`,
	},
	"text_14": {
		_const.RU: `Офис`,
		_const.EN: `Office`,
	},
	"text_15": {
		_const.RU: `Склады`,
		_const.EN: `Warehouses`,
	},
	"button_25": {
		_const.RU: `Автопилот`,
		_const.EN: `Autopilot`,
	},
	"button_26": {
		_const.RU: `Проложить маршрут`,
		_const.EN: `Plot a route.`,
	},
	"text_16": {
		_const.RU: `Руда`,
		_const.EN: `Ore`,
	},
	"text_20": {
		_const.RU: `Органика`,
		_const.EN: `Organic`,
	},
	"text_21": {
		_const.RU: `Нефть`,
		_const.EN: `Oil`,
	},
	"text_22": {
		_const.RU: `Гранулы`,
		_const.EN: `Granules`,
	},
	"text_23": {
		_const.RU: `Топливо`,
		_const.EN: `Fuel`,
	},
	"storage": {
		_const.RU: `Склад`,
		_const.EN: `Storage`,
	},
	"squadInventory": {
		_const.RU: `Трюм`,
		_const.EN: `Hold`,
	},
	"scanner": {
		_const.RU: `Сканер`,
		_const.EN: `Scanner`,
	},
	"box": {
		_const.RU: `Ящик`,
		_const.EN: `Box`,
	},
	// ItemCell
	"cell_tax": {
		_const.RU: `Налог:`,
		_const.EN: `Tax:`,
	},
	// PlaceDialog
	"text_17": {
		_const.RU: `Не все влезает:`,
		_const.EN: `Not everything fits:`,
	},
	// Range
	// RepairModal
	"text_18": {
		_const.RU: `Починить`,
		_const.EN: `Repair`,
	},
	"for": {
		_const.RU: `за`,
		_const.EN: `for`,
	},
	"text_19": {
		_const.RU: `Не нуждается в ремонте.`,
		_const.EN: `Does not need to be repaired.`,
	},
	"button_27": {
		_const.RU: `Закрыть`,
		_const.EN: `Close`,
	},
	"button_28": {
		_const.RU: `Починить`,
		_const.EN: `Repair`,
	},
	// Size
	// WarehouseTab
	// ThrowItemsDialog
	"text_24": {
		_const.RU: `Уничтожить эти предметы?`,
		_const.EN: `Destroy these items?`,
	},
	"button_29": {
		_const.RU: `Уничтожить`,
		_const.EN: `Destroy`,
	},
}
