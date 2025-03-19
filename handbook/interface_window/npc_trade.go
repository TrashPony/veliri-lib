package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var NPCTrade = map[string]map[string]string{
	"window_name": {
		_const.RU:   `Торговля`,
		_const.EN:   `Commerce`,
		_const.ZhCN: `贸易`,
	},
	"error_1": {
		_const.RU:   `Не удалось`,
		_const.EN:   `Failed`,
		_const.ZhCN: `失败`,
	},
	"button_1": {
		_const.RU:   `Отмена`,
		_const.EN:   `Cancel`,
		_const.ZhCN: `取消`,
	},
	"text_1": {
		_const.RU:   `плохое отношение`,
		_const.EN:   `bad attitude`,
		_const.ZhCN: `态度恶劣`,
	},
	"qty": {
		_const.RU:   `Кол-во`,
		_const.EN:   `Qty`,
		_const.ZhCN: `数量`,
	},
	"price": {
		_const.RU:   `Цена`,
		_const.EN:   `Price`,
		_const.ZhCN: `价格`,
	},
	"text_2": {
		_const.RU:   `Синтет`,
		_const.EN:   `Synthet`,
		_const.ZhCN: `合成体`,
	},
	"sell": {
		_const.RU:   `Продать`,
		_const.EN:   `Sell`,
		_const.ZhCN: `出售`,
	},
	"buy": {
		_const.RU:   `Купить`,
		_const.EN:   `Buy`,
		_const.ZhCN: `购买`,
	},
}
