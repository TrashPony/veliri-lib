package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var BpEditor = map[string]map[string]string{
	// ===== Шапка =====
	"window_name": {
		_const.RU:   `Баланс производства`,
		_const.EN:   `Production Balance`,
		_const.ZhCN: `生产平衡`,
	},
	"tab_ores": {
		_const.RU:   `Руда`,
		_const.EN:   `Ores`,
		_const.ZhCN: `矿石`,
	},
	"tab_recycle": {
		_const.RU:   `Сырьё`,
		_const.EN:   `Raw materials`,
		_const.ZhCN: `原材料`,
	},
	"tab_details": {
		_const.RU:   `Детали`,
		_const.EN:   `Parts`,
		_const.ZhCN: `零件`,
	},
	"tab_blueprints": {
		_const.RU:   `Чертежи`,
		_const.EN:   `Blueprints`,
		_const.ZhCN: `蓝图`,
	},
	"tab_loot": {
		_const.RU:   `Лут`,
		_const.EN:   `Loot`,
		_const.ZhCN: `战利品`,
	},

	// ===== Общие =====
	"btn_export": {
		_const.RU:   `JSON`,
		_const.EN:   `JSON`,
		_const.ZhCN: `JSON`,
	},
	"label_batch": {
		_const.RU:   `Партия`,
		_const.EN:   `Batch`,
		_const.ZhCN: `批次`,
	},
	"label_yields": {
		_const.RU:   `Выход при переработке:`,
		_const.EN:   `Processing yields:`,
		_const.ZhCN: `加工产出：`,
	},

	// ===== Вкладка Руда =====
	"section_ores_title": {
		_const.RU:   `Руда`,
		_const.EN:   `Ores`,
		_const.ZhCN: `矿石`,
	},
	"section_mining": {
		_const.RU:   `Руда в месторождениях`,
		_const.EN:   `Ores in deposits`,
		_const.ZhCN: `矿床中的矿石`,
	},
	"section_oil": {
		_const.RU:   `Жидкость`,
		_const.EN:   `Liquids`,
		_const.ZhCN: `液体`,
	},
	"section_organic": {
		_const.RU:   `Органика`,
		_const.EN:   `Organics`,
		_const.ZhCN: `有机物`,
	},
	"badge_combine": {
		_const.RU:   `COMBINE`,
		_const.EN:   `COMBINE`,
		_const.ZhCN: `合成`,
	},
	"section_granules": {
		_const.RU:   `Гранулы`,
		_const.EN:   `Granules`,
		_const.ZhCN: `颗粒`,
	},

	// ===== Вкладка Сырьё =====
	"section_recycle_title": {
		_const.RU:   `Сырьё`,
		_const.EN:   `Raw materials`,
		_const.ZhCN: `原材料`,
	},

	// ===== Вкладка Детали =====
	"section_details_title": {
		_const.RU:   `Детали-константы`,
		_const.EN:   `Constant parts`,
		_const.ZhCN: `常量零件`,
	},
	"details_t1": {
		_const.RU:   `Детали T1`,
		_const.EN:   `Parts T1`,
		_const.ZhCN: `零件 T1`,
	},
	"details_t2": {
		_const.RU:   `Детали T2`,
		_const.EN:   `Parts T2`,
		_const.ZhCN: `零件 T2`,
	},
	"label_tech": {
		_const.RU:   `Tech`,
		_const.EN:   `Tech`,
		_const.ZhCN: `科技`,
	},
	"label_price": {
		_const.RU:   `price:`,
		_const.EN:   `price:`,
		_const.ZhCN: `价格：`,
	},
	"label_requires": {
		_const.RU:   `Требует:`,
		_const.EN:   `Requires:`,
		_const.ZhCN: `需要：`,
	},
	"empty_details_t1": {
		_const.RU:   `Нет деталей T1`,
		_const.EN:   `No T1 parts`,
		_const.ZhCN: `没有 T1 零件`,
	},
	"empty_details_t2": {
		_const.RU:   `Нет деталей T2`,
		_const.EN:   `No T2 parts`,
		_const.ZhCN: `没有 T2 零件`,
	},

	// ===== Вкладка Чертежи =====
	"section_blueprints_title": {
		_const.RU:   `Чертежи предметов`,
		_const.EN:   `Item blueprints`,
		_const.ZhCN: `物品蓝图`,
	},
	"placeholder_search": {
		_const.RU:   `Поиск...`,
		_const.EN:   `Search...`,
		_const.ZhCN: `搜索...`,
	},
	"filter_all_tech": {
		_const.RU:   `Все Tech`,
		_const.EN:   `All Tech`,
		_const.ZhCN: `所有科技`,
	},
	"filter_all_types": {
		_const.RU:   `Все типы`,
		_const.EN:   `All types`,
		_const.ZhCN: `所有类型`,
	},
	"section_total": {
		_const.RU:   `ИТОГО`,
		_const.EN:   `TOTAL`,
		_const.ZhCN: `总计`,
	},
	"badge_bpo": {
		_const.RU:   `BPO`,
		_const.EN:   `BPO`,
		_const.ZhCN: `BPO`,
	},
	"label_resources_first": {
		_const.RU:   `Ресурсы первого уровня:`,
		_const.EN:   `First level resources:`,
		_const.ZhCN: `一级资源：`,
	},
	"label_primitives_total": {
		_const.RU:   `Итог (примитивы):`,
		_const.EN:   `Total (primitives):`,
		_const.ZhCN: `总计（原始）：`,
	},

	// ===== Вкладка Лут =====
	"section_loot_title": {
		_const.RU:   `Система лута`,
		_const.EN:   `Loot system`,
		_const.ZhCN: `战利品系统`,
	},
	"section_sectors": {
		_const.RU:   `Сектора (Аномалии)`,
		_const.EN:   `Sectors (Anomalies)`,
		_const.ZhCN: `区域（异常）`,
	},
	"label_max_points": {
		_const.RU:   `Max Points:`,
		_const.EN:   `Max Points:`,
		_const.ZhCN: `最大点数：`,
	},
	"label_resources_reservoirs": {
		_const.RU:   `Ресурсы (Reservoirs)`,
		_const.EN:   `Resources (Reservoirs)`,
		_const.ZhCN: `资源（储层）`,
	},
	"label_treasure": {
		_const.RU:   `Сокровища (Аномалии)`,
		_const.EN:   `Treasure (Anomalies)`,
		_const.ZhCN: `宝藏（异常）`,
	},
	"label_parts_t": {
		_const.RU:   `Детали T`,
		_const.EN:   `Parts T`,
		_const.ZhCN: `零件 T`,
	},
	"unit_pieces": {
		_const.RU:   `шт`,
		_const.EN:   `pcs`,
		_const.ZhCN: `件`,
	},
	"label_goods_grade": {
		_const.RU:   `Товары Grade`,
		_const.EN:   `Goods Grade`,
		_const.ZhCN: `商品等级`,
	},
	"label_blueprints_t": {
		_const.RU:   `Чертежи T`,
		_const.EN:   `Blueprints T`,
		_const.ZhCN: `蓝图 T`,
	},
	"label_science": {
		_const.RU:   `Научные данные`,
		_const.EN:   `Science data`,
		_const.ZhCN: `科研数据`,
	},
	"section_outposts": {
		_const.RU:   `Форпосты`,
		_const.EN:   `Outposts`,
		_const.ZhCN: `哨站`,
	},
	"details_t1_loot": {
		_const.RU:   `Детали Т1`,
		_const.EN:   `Parts T1`,
		_const.ZhCN: `零件 T1`,
	},
	"label_blueprints": {
		_const.RU:   `Чертежи`,
		_const.EN:   `Blueprints`,
		_const.ZhCN: `蓝图`,
	},
	"label_currency": {
		_const.RU:   `Валюта`,
		_const.EN:   `Currency`,
		_const.ZhCN: `货币`,
	},
	"section_bots": {
		_const.RU:   `Боты`,
		_const.EN:   `Bots`,
		_const.ZhCN: `机器人`,
	},
	"label_drop_chances": {
		_const.RU:   `Шансы дропа`,
		_const.EN:   `Drop chances`,
		_const.ZhCN: `掉落概率`,
	},
	"section_pools": {
		_const.RU:   `Справочник пулов`,
		_const.EN:   `Pool handbook`,
		_const.ZhCN: `池手册`,
	},
	"section_details_pools": {
		_const.RU:   `Детали`,
		_const.EN:   `Parts`,
		_const.ZhCN: `零件`,
	},
	"label_parts_tier": {
		_const.RU:   `Тир деталей:`,
		_const.EN:   `Parts tier:`,
		_const.ZhCN: `零件等级：`,
	},
	"section_blueprints_t0": {
		_const.RU:   `Чертежи T0`,
		_const.EN:   `Blueprints T0`,
		_const.ZhCN: `蓝图 T0`,
	},
	"section_blueprints_t1": {
		_const.RU:   `Чертежи T1`,
		_const.EN:   `Blueprints T1`,
		_const.ZhCN: `蓝图 T1`,
	},
	"unit_blueprints": {
		_const.RU:   `чертежей`,
		_const.EN:   `blueprints`,
		_const.ZhCN: `个蓝图`,
	},
	"section_frr": {
		_const.RU:   `Научные данные (FRR)`,
		_const.EN:   `Science data (FRR)`,
		_const.ZhCN: `科研数据 (FRR)`,
	},
	"section_goods": {
		_const.RU:   `Товары`,
		_const.EN:   `Goods`,
		_const.ZhCN: `商品`,
	},
	"unit_grade": {
		_const.RU:   `грейд`,
		_const.EN:   `grade`,
		_const.ZhCN: `等级`,
	},
	"loot_empty": {
		_const.RU:   `Данные лута не загружены или ручка недоступна`,
		_const.EN:   `Loot data not loaded or handler unavailable`,
		_const.ZhCN: `战利品数据未加载或接口不可用`,
	},

	// ===== Экспорт панель =====
	"section_export": {
		_const.RU:   `Экспорт`,
		_const.EN:   `Export`,
		_const.ZhCN: `导出`,
	},
	"unit_records": {
		_const.RU:   `записей`,
		_const.EN:   `records`,
		_const.ZhCN: `条记录`,
	},
	"btn_copy": {
		_const.RU:   `Копировать`,
		_const.EN:   `Copy`,
		_const.ZhCN: `复制`,
	},
	"alert_copied": {
		_const.RU:   `Скопировано`,
		_const.EN:   `Copied`,
		_const.ZhCN: `已复制`,
	},
	// Названия ресурсов для системы лута
	"resource_copper": {
		_const.RU:   `Медная руда`,
		_const.EN:   `Copper ore`,
		_const.ZhCN: `铜矿石`,
	},
	"resource_iron": {
		_const.RU:   `Железная руда`,
		_const.EN:   `Iron ore`,
		_const.ZhCN: `铁矿石`,
	},
	"resource_silicon": {
		_const.RU:   `Кремниевая руда`,
		_const.EN:   `Silicon ore`,
		_const.ZhCN: `硅矿石`,
	},
	"resource_lithium": {
		_const.RU:   `Литиева руда`,
		_const.EN:   `Lithium ore`,
		_const.ZhCN: `锂矿石`,
	},
	"resource_titanium": {
		_const.RU:   `Титановая руда`,
		_const.EN:   `Titanium ore`,
		_const.ZhCN: `钛矿石`,
	},
	"resource_oil": {
		_const.RU:   `Нефть`,
		_const.EN:   `Oil`,
		_const.ZhCN: `石油`,
	},
	"resource_unknown": {
		_const.RU:   `Ресурс #`,
		_const.EN:   `Resource #`,
		_const.ZhCN: `资源 #`,
	},
}
