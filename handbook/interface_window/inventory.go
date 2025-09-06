package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Inventory = map[string]map[string]string{
	// ArmModal
	"text_1": {
		_const.RU:   `за %count%:`,
		_const.EN:   `for %count%:`,
		_const.ZhCN: `为 %count%:`,
	},
	"text_2": {
		_const.RU:   `Всего:`,
		_const.EN:   `Total:`,
		_const.ZhCN: `总计:`,
	},
	"text_3": {
		_const.RU:   `Кол-во:`,
		_const.EN:   `Qty:`,
		_const.ZhCN: `数量:`,
	},
	"button_1": {
		_const.RU:   `Закрыть`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
	"button_2": {
		_const.RU:   `Зарядить`,
		_const.EN:   `Arm`,
		_const.ZhCN: `装备`,
	},
	// BackgroundItemCell
	// CellSubMenu
	"button_3": {
		_const.RU:   `Установить оружие`,
		_const.EN:   `Install weapon`,
		_const.ZhCN: `安装武器`,
	},
	"button_4": {
		_const.RU:   `Установить снаряды`,
		_const.EN:   `Install ammo`,
		_const.ZhCN: `安装弹药`,
	},
	"button_5": {
		_const.RU:   `Активировать корпус`,
		_const.EN:   `Activate the body`,
		_const.ZhCN: `激活船体`,
	},
	"button_6": {
		_const.RU:   `Заправить корпус`,
		_const.EN:   `Fill body`,
		_const.ZhCN: `填充船体`,
	},
	"button_7": {
		_const.RU:   `Подробнее`,
		_const.EN:   `More details`,
		_const.ZhCN: `更多详情`,
	},
	"button_8": {
		_const.RU:   `Разделить`,
		_const.EN:   `Split`,
		_const.ZhCN: `拆分`,
	},
	"button_9": {
		_const.RU:   `Выбросить`,
		_const.EN:   `Throw`,
		_const.ZhCN: `丢弃`,
	},
	"button_10": {
		_const.RU:   `Уничтожить`,
		_const.EN:   `Destroy`,
		_const.ZhCN: `销毁`,
	},
	"button_11": {
		_const.RU:   `Установить`,
		_const.EN:   `Install`,
		_const.ZhCN: `安装`,
	},
	"button_12": {
		_const.RU:   `Продать`,
		_const.EN:   `Sell`,
		_const.ZhCN: `出售`,
	},
	"button_13": {
		_const.RU:   `Починить`,
		_const.EN:   `Repair`,
		_const.ZhCN: `修理`,
	},
	"button_14": {
		_const.RU:   `Снять`,
		_const.EN:   `Take off`,
		_const.ZhCN: `卸下`,
	},
	"button_15": {
		_const.RU:   `В грузовой отсек`,
		_const.EN:   `To cargo hold`,
		_const.ZhCN: `装入货舱`,
	},
	"button_16": {
		_const.RU:   `На склад`,
		_const.EN:   `To storage`,
		_const.ZhCN: `放入仓库`,
	},
	"button_17": {
		_const.RU:   `Произвести`,
		_const.EN:   `Produce`,
		_const.ZhCN: `生产`,
	},
	"button_18": {
		_const.RU:   `Отменить`,
		_const.EN:   `Cancel`,
		_const.ZhCN: `取消`,
	},
	"text_4": {
		_const.RU:   `Кол-во:`,
		_const.EN:   `Qty:`,
		_const.ZhCN: `数量:`,
	},
	"text_5": {
		_const.RU:   `Сr. за шт.`,
		_const.EN:   `Cr. per pce`,
		_const.ZhCN: `每件价格`,
	},
	"text_6": {
		_const.RU:   `Всего:`,
		_const.EN:   `Total:`,
		_const.ZhCN: `总计:`,
	},
	"text_7": {
		_const.RU:   `Налог (%tax%%)`,
		_const.EN:   `Tax (%tax%%)`,
		_const.ZhCN: `税 (%tax%%)`,
	},
	"text_8": {
		_const.RU:   `Вы получите:`,
		_const.EN:   `You'll get:`,
		_const.ZhCN: `你将获得:`,
	},
	"button_19": {
		_const.RU:   `Отмена`,
		_const.EN:   `Cancel`,
		_const.ZhCN: `取消`,
	},
	"button_20": {
		_const.RU:   `Продать`,
		_const.EN:   `Sell`,
		_const.ZhCN: `出售`,
	},
	"button_21": {
		_const.RU:   `Отмена`,
		_const.EN:   `Cancel`,
		_const.ZhCN: `取消`,
	},
	"button_22": {
		_const.RU:   `Разделить`,
		_const.EN:   `Divide`,
		_const.ZhCN: `拆分`,
	},
	"button_30": {
		_const.RU:   `Распаковать`,
		_const.EN:   `Unpack`,
		_const.ZhCN: `解包`,
	},
	// FillUpModal
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
		_const.RU:   `Кол-во:`,
		_const.EN:   `Qty:`,
		_const.ZhCN: `数量:`,
	},
	"button_23": {
		_const.RU:   `Закрыть`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
	"button_24": {
		_const.RU:   `Заправить`,
		_const.EN:   `Refuel`,
		_const.ZhCN: `加注燃料`,
	},
	// Inventory
	"window_name": {
		_const.RU:   `Склад / Грузовой отсек`,
		_const.EN:   `Storage / Cargo Hold`,
		_const.ZhCN: `仓储 / 货舱`,
	},
	"text_12": {
		_const.RU:   `Склад заблокирован. Продлите аренду офиса или восстановите модуль базы.`,
		_const.EN:   `The warehouse is locked. Please extend the office rental or restore the database module.`,
		_const.ZhCN: `仓库已锁定。请延长办公室租期或恢复数据库模块。`,
	},
	"text_13": {
		_const.RU:   `Нет прав.`,
		_const.EN:   `You have no rights.`,
		_const.ZhCN: `没有权限。`,
	},
	"text_14": {
		_const.RU:   `Офис`,
		_const.EN:   `Office`,
		_const.ZhCN: `办公室`,
	},
	"text_15": {
		_const.RU:   `Склады`,
		_const.EN:   `Warehouses`,
		_const.ZhCN: `仓库`,
	},
	"button_25": {
		_const.RU:   `Автопилот`,
		_const.EN:   `Autopilot`,
		_const.ZhCN: `自动驾驶`,
	},
	"button_26": {
		_const.RU:   `Проложить маршрут`,
		_const.EN:   `Plot a route.`,
		_const.ZhCN: `规划路线`,
	},
	"text_16": {
		_const.RU:   `Руда`,
		_const.EN:   `Ore`,
		_const.ZhCN: `矿石`,
	},
	"text_20": {
		_const.RU:   `Органика`,
		_const.EN:   `Organic`,
		_const.ZhCN: `有机物`,
	},
	"text_21": {
		_const.RU:   `Нефть`,
		_const.EN:   `Oil`,
		_const.ZhCN: `石油`,
	},
	"text_22": {
		_const.RU:   `Гранулы`,
		_const.EN:   `Granules`,
		_const.ZhCN: `颗粒`,
	},
	"text_23": {
		_const.RU:   `Топливо`,
		_const.EN:   `Fuel`,
		_const.ZhCN: `燃料`,
	},
	"storage": {
		_const.RU:   `Склад`,
		_const.EN:   `Storage`,
		_const.ZhCN: `仓库`,
	},
	"squadInventory": {
		_const.RU:   `Грузовой отсек`,
		_const.EN:   `Cargo Hold`,
		_const.ZhCN: `货舱`,
	},
	"scanner": {
		_const.RU:   `Сканер`,
		_const.EN:   `Scanner`,
		_const.ZhCN: `扫描仪`,
	},
	"box": {
		_const.RU:   `Ящик`,
		_const.EN:   `Box`,
		_const.ZhCN: `箱子`,
	},
	// ItemCell
	"cell_tax": {
		_const.RU:   `Налог:`,
		_const.EN:   `Tax:`,
		_const.ZhCN: `税:`,
	},
	// PlaceDialog
	"text_17": {
		_const.RU:   `Не все влезает:`,
		_const.EN:   `Not everything fits:`,
		_const.ZhCN: `无法全部放入:`,
	},
	// Range
	// RepairModal
	"text_18": {
		_const.RU:   `Починить`,
		_const.EN:   `Repair`,
		_const.ZhCN: `修理`,
	},
	"for": {
		_const.RU:   `за`,
		_const.EN:   `for`,
		_const.ZhCN: `为`,
	},
	"text_19": {
		_const.RU:   `Не нуждается в ремонте.`,
		_const.EN:   `Does not need to be repaired.`,
		_const.ZhCN: `不需要修理。`,
	},
	"button_27": {
		_const.RU:   `Закрыть`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
	"button_28": {
		_const.RU:   `Починить`,
		_const.EN:   `Repair`,
		_const.ZhCN: `修理`,
	},
	// Size
	// WarehouseTab
	// ThrowItemsDialog
	"text_24": {
		_const.RU:   `Уничтожить эти предметы?`,
		_const.EN:   `Destroy these items?`,
		_const.ZhCN: `销毁这些物品吗？`,
	},
	"button_29": {
		_const.RU:   `Уничтожить`,
		_const.EN:   `Destroy`,
		_const.ZhCN: `销毁`,
	},
	"text_25": {
		_const.RU:   `Ангар`,
		_const.EN:   `Hangar`,
		_const.ZhCN: `机库`,
	},
	"text_26": {
		_const.RU:   `Нельзя выбросить/продать <span class="importantly">последний корпус</span>.`,
		_const.EN:   `You cannot drop/sell <span class="importantly">the last hull</span>.`,
		_const.ZhCN: `不能丢弃/出售<span class="importantly">最后一个船体</span>。`,
	},
	"text_27": {
		_const.RU:   `Не поместилось.`,
		_const.EN:   `Didn't fit.`,
		_const.ZhCN: `无法放入。`,
	},
	"text_28": {
		_const.RU:   `Не известная ошибка.`,
		_const.EN:   `Unknown error.`,
		_const.ZhCN: `未知错误。`,
	},
	// modal names
	"arm_modal_head": {
		_const.RU:   `Покупка снарядов`,
		_const.EN:   `Buying shells`,
		_const.ZhCN: `买贝壳`,
	},
	"repair_modal_head": {
		_const.RU:   `Ремонт`,
		_const.EN:   `Repair`,
		_const.ZhCN: `维修`,
	},
	"fill_up_modal_head": {
		_const.RU:   `Заправка`,
		_const.EN:   `Refueling`,
		_const.ZhCN: `加油`,
	},
	// tips
	"tip_1": {
		_const.RU:   `Фильтр: %filter_name%`,
		_const.EN:   `Filter: %filter_name%`,
		_const.ZhCN: `筛选器: %filter_name%`,
	},
	"tip_2": {
		_const.RU:   `Перенести все подходящие предметы из грузового отсека`,
		_const.EN:   `Transfer all matching items from cargo hold`,
		_const.ZhCN: `从货舱转移所有匹配物品`,
	},
	"tip_3": {
		_const.RU:   `Объединить топливные ячейки`,
		_const.EN:   `Merge fuel cells`,
		_const.ZhCN: `合并燃料单元`,
	},
	"tip_4": {
		_const.RU:   `Развернуть грузовой отсек`,
		_const.EN:   `Expand cargo hold`,
		_const.ZhCN: `展开货舱`,
	},
	"tip_5": {
		_const.RU:   `Свернуть грузовой отсек`,
		_const.EN:   `Collapse cargo hold`,
		_const.ZhCN: `收起货舱`,
	},
	"tip_6": {
		_const.RU:   `Грузовые отсеки транспортов, которые сейчас не используются`,
		_const.EN:   `Cargo holds of currently unused vehicles`,
		_const.ZhCN: `当前未使用载具的货舱`,
	},
	"tip_7": {
		_const.RU:   `Склады на других базах (только просмотр)`,
		_const.EN:   `Warehouses at other bases (view only)`,
		_const.ZhCN: `其他基地仓库(仅查看)`,
	},
	"tip_8": {
		_const.RU: `<div>
        <div>Состояние: Износ <b>%durability%%</b><br><br></div>
        <div>► При достижении 100% предмет будет уничтожен.</div>
        <div>► Каждая смерть увеличивает износ на 2-5%.</div>
        <div>► Для модификаторов: 6-15% за смерть.<br><br></div>
        <div>Предметы с износом нельзя продать на рыноке, но можно передать другому игроку, бросив на землю, или разобрать на ресурсы.</div>
    </div>`,
		_const.EN: `<div>
        <div>Condition: Wear <b>%durability%%</b><br><br></div>
        <div>► Item will be destroyed at 100%.</div>
        <div>► Each death increases wear by 2-5%.</div>
        <div>► For modifiers: 6-15% per death.<br><br></div>
        <div>Worn items cannot be sold on the market, but can be dropped for other players or recycled for resources.</div>
    </div>`,
		_const.ZhCN: `<div>
        <div>状态: 磨损度 <b>%durability%%</b><br><br></div>
        <div>► 达到100%时物品将被销毁。</div>
        <div>► 每次死亡增加2-5%磨损度。</div>
        <div>► 改装件: 每次死亡增加6-15%。<br><br></div>
        <div>磨损物品无法在市场出售，但可丢弃给其他玩家或分解为资源。</div>
    </div>`,
	},
	"tip_9": {
		_const.RU: `<div>
        <div class="tooltip-header">Состояние корпуса: Износ <b>%durability%%</b><br><br></div>
        <div>► Каждые 5% износа дают <b>-1%</b> к скорости и HP.</div>
        <div>► Каждая смерть увеличивает износ на 2-5%.</div>
        <div>► Чем выше износ, тем дороже ремонт.<br><br></div>
        <div>Предметы с износом нельзя продать на рыноке, но можно передать другому игроку, бросив на землю, или разобрать на ресурсы.</div>
    </div>`,
		_const.EN: `<div>
        <div class="tooltip-header">Hull Condition: Wear <b>%durability%%</b><br><br></div>
        <div>► Every 5% wear reduces speed and HP by <b>-1%</b>.</div>
        <div>► Each death increases wear by 2-5%.</div>
        <div>► Higher wear increases repair costs.<br><br></div>
        <div>Worn items cannot be sold on the market, but can be dropped for other players or recycled for resources.</div>
    </div>`,
		_const.ZhCN: `<div>
        <div class="tooltip-header">机体状态: 磨损度 <b>%durability%%</b><br><br></div>
        <div>► 每5%磨损度降低速度和HP <b>-1%</b>。</div>
        <div>► 每次死亡增加2-5%磨损度。</div>
        <div>► 磨损度越高，维修费用越贵。<br><br></div>
        <div>磨损物品无法在市场出售，但可丢弃给其他玩家或分解为资源。</div>
    </div>`,
	},
}
