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
	"tip_1": {
		_const.RU:   `<span class="importantly" style="color: #faca3e;">Внимание!</span> Торговля нелегальным товаром приведёт к потере репутации.`,
		_const.EN:   `<span class="importantly" style="color: #faca3e;">Warning!</span> Trading illegal goods will lead to a loss of reputation.`,
		_const.ZhCN: `<span class="importantly" style="color: #faca3e;">注意！</span> 交易非法商品将导致声望下降。`,
	},
	"tip_2": {
		_const.RU: `<p>Налогообразующий товар.</p>
        <p>Минимальный налог базы доступен, если количество всех необходимых товаров превышает 20%. Если товаров меньше 20%, налог базы начинает расти.</p>
        <p>Если необходимого товара совсем нет на базе, налог растёт по формуле: 75% / количество налогообразующих товаров.</p>`,
		_const.EN: `<p>Tax-forming goods.</p>
        <p>The minimum base tax is available if the quantity of all required goods exceeds 20%. If the goods are below 20%, the base tax begins to increase.</p>
        <p>If a required good is completely absent from the base, the tax increases according to the formula: 75% / number of tax-forming goods.</p>`,
		_const.ZhCN: `<p>税基商品。</p>
        <p>当所有必需商品的数量超过20%时，基地适用最低税率。若商品数量低于20%，税率将开始上升。</p>
        <p>如果基地完全没有某种必需商品，则税率按公式计算：75% / 税基商品数量。</p>`,
	},
}
