package handbook

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"math/rand"
)

type TradeEvent struct {
	Fraction    string `json:"fraction"`
	ProductName string `json:"product_name"`
	Type        string `json:"type"` // lack/excess/cancel
	Text        string `json:"text"`
}

//func init() {
//
//	for _, v := range TradeEvents[_const.RU] {
//
//		time.Sleep(time.Millisecond * 100)
//
//		var err error
//		if v.Text != "" {
//			v.Text, err = Translate(v.Text, _const.RU, _const.EN)
//			if err != nil {
//				panic(err)
//			}
//
//			v.Text = strings.ReplaceAll(v.Text, `\"`, `\\\"`)
//			v.Text = strings.ReplaceAll(v.Text, `"`, `\"`)
//			v.Text = strings.ReplaceAll(v.Text, "\n", "")
//		}
//
//		fmt.Println("{")
//		fmt.Println("Fraction: \"" + v.Fraction + "\",")
//		fmt.Println("ProductName: \"" + v.ProductName + "\",")
//		fmt.Println("Type: \"" + v.Type + "\",")
//		fmt.Println("Text: \"" + v.Text + "\",")
//		fmt.Println("},")
//	}
//}

var TradeEvents = map[string][]TradeEvent{
	_const.EN: {
		{
			Fraction:    _const.Replicas,
			ProductName: "weapon_parts",
			Type:        "lack",
			Text:        "Replics has openly acknowledged the lack of \"weapons units\" to supply its \"invincible\" army. Synthetes possessing this product can profitably sell it on the following basis %BaseName% Replics.",
		},
		{
			Fraction:    _const.Replicas,
			ProductName: "weapon_parts",
			Type:        "cancel",
			Text:        "Replics breaks the contracts and refuses to supply further \"weapon parts\" and is no longer interested in this product.",
		},
		{
			Fraction:    _const.Replicas,
			ProductName: "technological_maps",
			Type:        "lack",
			Text:        "Having admitted defeat in the technology race, Replics is once again allowing the sale of high-tech products, and in particular - “technological maps”. Synthetes possessing this product can profitably sell it on the following basis %BaseName% Replics.",
		},
		{
			Fraction:    _const.Replicas,
			ProductName: "technological_maps",
			Type:        "cancel",
			Text:        "Replics again refuses to cooperate and stops any trade in goods - \"technological maps\". Apparently, we should expect another round of man-made disasters made by Replics.",
		},
		{
			Fraction:    _const.Replicas,
			ProductName: "industrial_materials",
			Type:        "lack",
			Text:        "Today Replics came up with an offer to buy goods from any itinerant synthetics - \"industrial materials\", heading it all under the umbrella - \"expansion\". Synthetes possessing this product can profitably sell it on the following basis %BaseName% Replics.",
		},
		{
			Fraction:    _const.Replicas,
			ProductName: "industrial_materials",
			Type:        "cancel",
			Text:        "As expected, after a short but tumultuous trade, Replics, having saturated its warehouses, stops buying the next product - \"industrial materials\".",
		},
		{
			Fraction:    _const.Replicas,
			ProductName: "subatomic_composites",
			Type:        "lack",
			Text:        "Replics appears to have faced a precision equipment bundling crisis and reopened its markets to trade, causing the prices of the “subatomic composites” product to skyrocket. Synthetes possessing this product can profitably sell it on the following basis %BaseName% Replics.",
		},
		{
			Fraction:    _const.Replicas,
			ProductName: "subatomic_composites",
			Type:        "cancel",
			Text:        "Replics is ending the trade in \"subatomic composites\", claiming that the assumption has become a place of speculation and betrayal. After a series of demonstrative annihilations of service bots and base operators, the trade in “subatomic composites” was abruptly discontinued.",
		},
		{
			Fraction:    _const.Replicas,
			ProductName: "general_purpose_equipment",
			Type:        "lack",
			Text:        "Replics is interested in stocking up on \"general equipment\" items and will pay generously to any merchant who brings the item back to Replics' property. Synthetes possessing this product can profitably sell it on the following basis %BaseName% Replics.",
		},
		{
			Fraction:    _const.Replicas,
			ProductName: "general_purpose_equipment",
			Type:        "cancel",
			Text:        "After a long and mutually beneficial trade, Replics stops accepting shipments of the next item - “general equipment”.",
		},
		{
			Fraction:    _const.Replicas,
			ProductName: "energy_storage",
			Type:        "lack",
			Text:        "Due to the energy crisis in some of its own regions, Replics was forced to issue a purchase announcement for the “energy storage” product. Synthetes possessing this product can profitably sell it on the following basis %BaseName% Replics.",
		},
		{
			Fraction:    _const.Replicas,
			ProductName: "energy_storage",
			Type:        "cancel",
			Text:        "Replics apparently, having finally resolved the problems of the energy crisis, sharply reduced the number of purchases of the “energy storage” product, which caused the price policy to plummet.",
		},
		{
			Fraction:    _const.Explores,
			ProductName: "weapon_parts",
			Type:        "lack",
			Text:        "Explores has come up with an unusual offer to buy \"weapon parts\" merchandise. This message was explained as a compulsory measure. All synthetes possessing this product can profitably sell it on the following base %BaseName% Explores.",
		},
		{
			Fraction:    _const.Explores,
			ProductName: "weapon_parts",
			Type:        "cancel",
			Text:        "Obviously, that there is no hidden Explores, conflict or self-defense goals, but, the recent massive purchase of \"weapon parts\" has come to an end. Explores thanked all synthetes for their assistance and limited the sale of this product. Massive price drops are expected within the Explores regions.",
		},
		{
			Fraction:    _const.Explores,
			ProductName: "technological_maps",
			Type:        "lack",
			Text:        "The prices for the “technological cards” product are rising again. Explores has officially announced a new scientific breakthrough and the need for a sufficient number of \"flow sheets\". All synthetes possessing this product can profitably sell it on the following base %BaseName% Explores.",
		},
		{
			Fraction:    _const.Explores,
			ProductName: "technological_maps",
			Type:        "cancel",
			Text:        "The fruitful checkout trade in the Explores area appears to have come to an end. Many synthetics are denied the sale of this product, and prices are falling rapidly.",
		},
		{
			Fraction:    _const.Explores,
			ProductName: "industrial_materials",
			Type:        "lack",
			Text:        "It looks like the prices of the “industrial materials” commodity have jumped again in the Explores regions. Explores, in turn, assured their followers that this is a necessary step in the construction of new scientific bases. All synthetes possessing this product can profitably sell it on the following base %BaseName% Explores.",
		},
		{
			Fraction:    _const.Explores,
			ProductName: "industrial_materials",
			Type:        "cancel",
			Text:        "With the end of the objectives, Explores ceased, and the bulk purchase of the \"industrial materials\" product. Explores thanked all synthetes for their assistance and limited the sale of this product. Massive price drops are expected within the Explores regions.",
		},
		{
			Fraction:    _const.Explores,
			ProductName: "subatomic_composites",
			Type:        "lack",
			Text:        "The need pushed Explores to openly address all synthetics that trade for the supply of goods - “subatomic composites”. Explores also added that it will generously repay all those synthetics who bring “subatomic composites” to the Explores regions. All synthetes possessing this product can profitably sell it on the following base %BaseName% Explores.",
		},
		{
			Fraction:    _const.Explores,
			ProductName: "subatomic_composites",
			Type:        "cancel",
			Text:        "The next attenuation after the ascent also affected the “subatomic composites” product. Explores, obviously satisfying their needs, restricts the purchase of this product. Massive price drops are expected within the Explores regions.",
		},
		{
			Fraction:    _const.Explores,
			ProductName: "general_purpose_equipment",
			Type:        "lack",
			Text:        "Synthetes who have at their disposal a product - \"general purpose equipment\", have the opportunity to profitably sell it on the territory of Explores. This became known from today's Explores stock exchange reports. All synthetes possessing this product can profitably sell it on the following base %BaseName% Explores.",
		},
		{
			Fraction:    _const.Explores,
			ProductName: "general_purpose_equipment",
			Type:        "cancel",
			Text:        "Having learned about the uncontrolled sale of \"general equipment\" goods, Explores temporarily takes over the regulation of turnover and prices, which makes further sale of the \"general equipment\" goods unprofitable.",
		},
		{
			Fraction:    _const.Explores,
			ProductName: "energy_storage",
			Type:        "lack",
			Text:        "From third parties, it became known about the lack of quality goods - \"energy storage\" in the territories of Explores. Later, a representative of Explores independently dispelled these assumptions, stating the need to supply an expensive and much-needed product on this inhospitable planet. All synthetes possessing this product can profitably sell it on the following base %BaseName% Explores.",
		},
		{
			Fraction:    _const.Explores,
			ProductName: "energy_storage",
			Type:        "cancel",
			Text:        "It looks like the crisis of the commodity - \"energy storage\". still passed Explores. The demand for this type of product is falling, which makes it unprofitable for further sale in the Explores territories.",
		},
		{
			Fraction:    _const.Reverses,
			ProductName: "weapon_parts",
			Type:        "lack",
			Text:        "Bypassing the official rules, it became known that Reverses is interested in buying goods - \"weapon parts\". For what exactly needs - we do not know, but any synthetics, now, can make good money on this if they offer their products on the next base %BaseName% Reverses.",
		},
		{
			Fraction:    _const.Reverses,
			ProductName: "weapon_parts",
			Type:        "cancel",
			Text:        "The loophole was figured out and closed - Reverses cuts off any supply channels for \"weapon parts\" and crashes the markets. It's a shame, but prices are falling again.",
		},
		{
			Fraction:    _const.Reverses,
			ProductName: "technological_maps",
			Type:        "lack",
			Text:        "Unexpectedly, Reverses made an offer for bulk purchases of goods - “checklists”. Complementing his appeal with specifics that this should serve in the long term of Reverses' plans. Synthetes possessing this product can go to the next base %BaseName% Reverses where they will sell it at a bargain price.",
		},
		{
			Fraction:    _const.Reverses,
			ProductName: "technological_maps",
			Type:        "cancel",
			Text:        "It seems that the scientific activity of Reverses has come to an end. The prices for the “technological cards” product are falling rapidly, which makes it unprofitable to sell it in the future. Perhaps in the future, this situation will change more than once.",
		},
		{
			Fraction:    _const.Reverses,
			ProductName: "industrial_materials",
			Type:        "lack",
			Text:        "Probably, having seriously taken up the expansion of new regions, from today, Reverses raises the prices for the next product - \"industrial materials\". If you want to not only support Reverses, but also make good money on it, all synthetics are recommended to visit the following %BaseName% Reverses database, where they will sell this product at a bargain price.",
		},
		{
			Fraction:    _const.Reverses,
			ProductName: "industrial_materials",
			Type:        "cancel",
			Text:        "Sooner or later, but everything comes to an end - this also affected the Reverses with their interest in the product \"industrial materials\". All synthetes that did not have time to offer their goods will be disappointed to the proper extent, because prices for it began to fall sharply.",
		},
		{
			Fraction:    _const.Reverses,
			ProductName: "subatomic_composites",
			Type:        "lack",
			Text:        "Who would have thought, but Reverses is faced with a shortage of the next product - “subatomic composites”. All synthetics that now have this very expensive product can offer it for sale at very favorable prices on the following base %BaseName% Reverses.",
		},
		{
			Fraction:    _const.Reverses,
			ProductName: "subatomic_composites",
			Type:        "cancel",
			Text:        "The jump in sales has come to an end, and either by getting what it needs or finding another way to manufacture high-tech devices, Reverses is no longer buying subatomic composites. A sharp decline in prices for this product is predicted.",
		},
		{
			Fraction:    _const.Reverses,
			ProductName: "general_purpose_equipment",
			Type:        "lack",
			Text:        "The price has risen for the goods - “general purpose equipment”, whose shortage was suddenly felt in Reverses. All synthetics that currently own this product can offer it for sale at very favorable prices on the following base %BaseName% Reverses.",
		},
		{
			Fraction:    _const.Reverses,
			ProductName: "general_purpose_equipment",
			Type:        "cancel",
			Text:        "“General Purpose Equipment” products are no longer in demand in the Reverses regions. It's a shame that the Reverses themselves did not explain the unexpected price spikes for this common product. Meanwhile, we congratulate all those synthetes who managed to raise decent sums on the trade.",
		},
		{
			Fraction:    _const.Reverses,
			ProductName: "energy_storage",
			Type:        "lack",
			Text:        "The most important producers and traders of goods are “energy stores”, the Reverses community is facing a shortage of them. All synthetics who now have this product at their disposal are offered to profitably sell it on the following basis %BaseName% Reverses.",
		},
		{
			Fraction:    _const.Reverses,
			ProductName: "energy_storage",
			Type:        "cancel",
			Text:        "Should have improved their business, Reverses are no longer interested in purchasing a product - \"energy storage\". Prices have leveled out again and now, the sale of this product does not promise superprofits.",
		},
		{
			Fraction:    _const.Empty,
			ProductName: "all",
			Type:        "excess",
			Text:        "Due to massive overproduction based on %BaseName%, the market has seen a decline in the purchase price of many products. Do not waste your time and have time to take what you need now!",
		},
		{
			Fraction:    _const.Empty,
			ProductName: "all",
			Type:        "lack",
			Text:        "Due to the decline in production capacity based on %BaseName%, the market has seen an increase in the purchase price of many products. This can be a good help in the sale of certain scarce goods. Keep this in mind!",
		},
	},
	_const.RU: {
		// replics
		// оружие
		{
			Fraction:    _const.Replicas,
			ProductName: "weapon_parts",
			Type:        "lack",
			Text:        "Replics открыто признал о нехватке “<span class=\"importantly\">оружейных частей</span>” для снабжения своей “непобедимой” армии. Синтеты, обладающие данным товаром могут выгодно реализовать его на следующей базе %BaseName% Replics.",
		}, {
			Fraction:    _const.Replicas,
			ProductName: "weapon_parts",
			Type:        "cancel",
			Text:        "Replics разрывает договора и отказывается от дальнейших поставок “<span class=\"importantly\">оружейных частей</span>” и более не заинтересован в данном товаре.",
		},
		// тех карты
		{
			Fraction:    _const.Replicas,
			ProductName: "technological_maps",
			Type:        "lack",
			Text:        "Признав своё поражение в гонке технологий, Replics вновь допускает продажу высокотехнологичной продукции, а в особенности - “<span class=\"importantly\">технологических карт</span>”. Синтеты, обладающие данным товаром могут выгодно реализовать его на следующей базе %BaseName% Replics.",
		}, {
			Fraction:    _const.Replicas,
			ProductName: "technological_maps",
			Type:        "cancel",
			Text:        "Replics вновь отказывается от сотрудничества и прекращает какую-либо торговлю товаром - “<span class=\"importantly\">технологических карт</span>”. Судя по всему, стоит ожидать очередной виток техногенных катастроф, допущенных Replics.",
		},
		// промышленные материалы
		{
			Fraction:    _const.Replicas,
			ProductName: "industrial_materials",
			Type:        "lack",
			Text:        "Сегодня Replics выступил с предложением о покупке у любых странствующих синтетов товара - “<span class=\"importantly\">промышленных материалов</span>”, Озаглавив всё это под эгидой - “расширения”. Синтеты, обладающие данным товаром могут выгодно реализовать его на следующей базе %BaseName% Replics.",
		}, {
			Fraction:    _const.Replicas,
			ProductName: "industrial_materials",
			Type:        "cancel",
			Text:        "Как и ожидалось, после недолгой, но бурной торговли, Replics, насытив свои склады, прекращает скупку следующего товара - “<span class=\"importantly\">промышленных материалов</span>”.",
		},
		// субатомные композиты
		{
			Fraction:    _const.Replicas,
			ProductName: "subatomic_composites",
			Type:        "lack",
			Text:        "Судя по всему, Replics столкнулся с кризисом комплектации высокоточного оборудования и вновь открыл свои рынки для торговли, из-за чего, цены на товар - “<span class=\"importantly\">субатомные композиты</span>”, резко подскочили. Синтеты, обладающие данным товаром могут выгодно реализовать его на следующей базе %BaseName% Replics.",
		}, {
			Fraction:    _const.Replicas,
			ProductName: "subatomic_composites",
			Type:        "cancel",
			Text:        "Replics прекращает торговлю товаром “субатомных композитов”, заявив, что данное допущение стало место спекуляций и предательства. После серии показательных аннигиляций служебных ботов и операторов баз, торговля “<span class=\"importantly\">субатомными композитами</span>” была резко прекращена.",
		},
		// оборудование общего назн.
		{
			Fraction:    _const.Replicas,
			ProductName: "general_purpose_equipment",
			Type:        "lack",
			Text:        "Replics заинтересован в пополнении своих складов товарами “<span class=\"importantly\">оборудования общего назначения</span>” и щедро заплатит любому торговцу, доставившему данный товар во владения Replics. Синтеты, обладающие данным товаром могут выгодно реализовать его на следующей базе %BaseName% Replics.",
		}, {
			Fraction:    _const.Replicas,
			ProductName: "general_purpose_equipment",
			Type:        "cancel",
			Text:        "После долгой и взаимовыгодной торговли, Replics прекращает принимать поставки следующего товара - “<span class=\"importantly\">оборудование общего назначения</span>”.",
		},
		// накопители
		{
			Fraction:    _const.Replicas,
			ProductName: "energy_storage",
			Type:        "lack",
			Text:        "Из-за энергетического кризиса некоторых собственных регионов, Replics был вынужден обратиться с объявлением о покупке товаров - “<span class=\"importantly\">накопителя энергии</span>”. Синтеты, обладающие данным товаром могут выгодно реализовать его на следующей базе %BaseName% Replics.",
		}, {
			Fraction:    _const.Replicas,
			ProductName: "energy_storage",
			Type:        "cancel",
			Text:        "Replics очевидно, наконец разрешив проблемы энергетического кризиса, резко сократил количество покупок товара “<span class=\"importantly\">накопителей энергии</span>”, из-за чего, ценовая политика стремительно снизилась.",
		},
		// Explores
		// оружие
		{
			Fraction:    _const.Explores,
			ProductName: "weapon_parts",
			Type:        "lack",
			Text:        "Explores выступил с необычным предложением о покупке товара “<span class=\"importantly\">оружейных частей</span>”. Данное сообщение, было разъяснено, как вынужденная мера. Все синтеты, обладающие данным товаром могут выгодно реализовать его на следующей базе %BaseName% Explores.",
		}, {
			Fraction:    _const.Explores,
			ProductName: "weapon_parts",
			Type:        "cancel",
			Text:        "Очевидно, чтобы там не скрывал Explores, конфликт или цели самообороны, но, недавняя массовая скупка товара “<span class=\"importantly\">оружейных частей</span>” подошла к концу. Explores поблагодарил за содействие всех синтетов и ограничил продажу данного товара. Ожидается массовое падение цен в пределах регионов Explores.",
		},
		// технологические карты
		{
			Fraction:    _const.Explores,
			ProductName: "technological_maps",
			Type:        "lack",
			Text:        "Вновь возрастают цены на товар “<span class=\"importantly\">технологические карты</span>”. Explores официально объявил о витке нового научного прорыва и необходимости в достаточном количестве “технологических карт”. Все синтеты, обладающие данным товаром могут выгодно реализовать его на следующей базе %BaseName% Explores.",
		}, {
			Fraction:    _const.Explores,
			ProductName: "technological_maps",
			Type:        "cancel",
			Text:        "Плодотворная торговля товаром “<span class=\"importantly\">технологические карты</span>” на территории Explores, судя по всему, подошла к концу. Многим синтетам отказывают в продаже данного товара, а цены стремительно падают.",
		},
		// промышленные материалы
		{
			Fraction:    _const.Explores,
			ProductName: "industrial_materials",
			Type:        "lack",
			Text:        "Похоже, цены на товар “<span class=\"importantly\">промышленные материалы</span>” вновь подскочили в регионах Explores. Explores в свою очередь заверил своих последователей, что это необходимый шаг при возведении новых научных баз. Все синтеты, обладающие данным товаром могут выгодно реализовать его на следующей базе %BaseName% Explores.",
		}, {
			Fraction:    _const.Explores,
			ProductName: "industrial_materials",
			Type:        "cancel",
			Text:        "С окончанием целей Explores прекратилась, и массовая закупка товара “<span class=\"importantly\">промышленные материалы</span>”. Explores поблагодарил за содействие всех синтетов и ограничил продажу данного товара. Ожидается массовое падение цен в пределах регионов Explores.",
		},
		// субатомные композиты
		{
			Fraction:    _const.Explores,
			ProductName: "subatomic_composites",
			Type:        "lack",
			Text:        "Необходимость, толкнула Explores открыто обратиться ко всем синтетам, что занимаются торговлей о поставках товара - “<span class=\"importantly\">субатомные композиты</span>”. Explores также добавил, что щедро отплатит всем тем синтетам, что доставят “субатомные композиты” в регионы Explores. Все синтеты, обладающие данным товаром могут выгодно реализовать его на следующей базе %BaseName% Explores.",
		}, {
			Fraction:    _const.Explores,
			ProductName: "subatomic_composites",
			Type:        "cancel",
			Text:        "Очередное затухание после подъёма, также коснулось и товара “<span class=\"importantly\">субатомные композиты</span>”. Explores, очевидно удовлетворив свои нужды, ограничивает покупку данного товара. Ожидается массовое падение цен в пределах регионов Explores.",
		},
		// оборудование общего назн.
		{
			Fraction:    _const.Explores,
			ProductName: "general_purpose_equipment",
			Type:        "lack",
			Text:        "Синтеты, что имеют в своём распоряжении товар - “<span class=\"importantly\">оборудование общего назначения</span>”, имеют возможность выгодно продать его на территории Explores. Это стало известно из сегодняшних биржевых сводок Explores. Все синтеты, обладающие данным товаром могут выгодно реализовать его на следующей базе %BaseName% Explores.",
		}, {
			Fraction:    _const.Explores,
			ProductName: "general_purpose_equipment",
			Type:        "cancel",
			Text:        "Прознав о бесконтрольной продаже товара “<span class=\"importantly\">оборудование общего назначения</span>”, Explores временно принимается за регулирование оборотов и цен, что делает дальнейшую продажу товара - “оборудование общего назначения” невыгодным.",
		},
		// накопители
		{
			Fraction:    _const.Explores,
			ProductName: "energy_storage",
			Type:        "lack",
			Text:        "От третьих лиц, стало известно о нехватке качественных товаров - “<span class=\"importantly\">накопители энергии</span>” на территориях Explores. Позже, представитель Explores самостоятельно развеял данные предположения заявив о необходимости поставок дорогостоящего и столь необходимого на данной неприветливой планете товара. Все синтеты, обладающие данным товаром могут выгодно реализовать его на следующей базе %BaseName% Explores.",
		}, {
			Fraction:    _const.Explores,
			ProductName: "energy_storage",
			Type:        "cancel",
			Text:        "Похоже, кризис товара - “<span class=\"importantly\">накопители энергии</span>”. всё-таки миновал Explores. Спрос на данный вид товара падает, что делает его дальнейшую реализацию на территориях Explores невыгодной.",
		},
		// REVERSES
		// оружие
		{
			Fraction:    _const.Reverses,
			ProductName: "weapon_parts",
			Type:        "lack",
			Text:        "В обход официальных правил, стало известно, что Reverses заинтересован в покупке товаров - “<span class=\"importantly\">оружейных частей</span>”. Для каких именно нужд - нам неизвестно, но любые синтеты, сейчас, могут неплохо на этом заработать, если предложат свои товары на следующей базе %BaseName% Reverses.",
		}, {
			Fraction:    _const.Reverses,
			ProductName: "weapon_parts",
			Type:        "cancel",
			Text:        "Лазейка оказалась вычислена и закрыта - Reverses перекрывает любые каналы поставок товара “<span class=\"importantly\">оружейных частей</span>” и обваливает рынки. Очень жаль, но цены вновь падают.",
		},
		// технологические карты
		{
			Fraction:    _const.Reverses,
			ProductName: "technological_maps",
			Type:        "lack",
			Text:        "Неожиданно, Reverses сделало предложение о массовой закупке товаров - “<span class=\"importantly\">технологические карты</span>”. Дополнив своё обращение конкретикой, что это должно послужить в долгосрочной перспективе планов Reverses. Синтеты, что обладают данным товаром, могут отправиться на следующую базу %BaseName% Reverses где продадут его по выгодной цене.",
		}, {
			Fraction:    _const.Reverses,
			ProductName: "technological_maps",
			Type:        "cancel",
			Text:        "Кажется, научная деятельность Reverses подошла к концу. Цены на товар “<span class=\"importantly\">технологические карты</span>” стремительно падают, что делает его продажу в дальнейшем невыгодной. Возможно в будущем, эту ситуация переменится ещё не один раз.",
		},
		// промышленные материалы
		{
			Fraction:    _const.Reverses,
			ProductName: "industrial_materials",
			Type:        "lack",
			Text:        "Наверное, всерьёз занявшись экспансией новых регионов, с сегодняшнего дня, Reverses поднимает цены на следующий товар - “<span class=\"importantly\">промышленные материалы</span>”. Если вы желаете не только поддержать Reverses, но и хорошенько подзаработать на этом, всем синтетам рекомендуется посетить следующую базу %BaseName% Reverses где продадут этот товар по выгодной цене.",
		}, {
			Fraction:    _const.Reverses,
			ProductName: "industrial_materials",
			Type:        "cancel",
			Text:        "Рано или поздно, но всё подходит к концу - это же коснулось и Reverses с их заинтересованностью в товаре “<span class=\"importantly\">промышленные материалы</span>”. Все синтеты, что не успели предложить свой товар в должной мере будут разочарованы, ведь цены на него стали резко падать.",
		},
		// субатомные композиты
		{
			Fraction:    _const.Reverses,
			ProductName: "subatomic_composites",
			Type:        "lack",
			Text:        "Кто бы мог подумать, но Reverses столкнулся с нехваткой следующего товара - “<span class=\"importantly\">субатомные композиты</span>”. Все синтеты, что сейчас обладает данным и весьма дорогим товаром, могут предложить его для продажи по весьма выгодным ценам на следующей базе %BaseName% Reverses.",
		}, {
			Fraction:    _const.Reverses,
			ProductName: "subatomic_composites",
			Type:        "cancel",
			Text:        "Скачок в продажах подошёл к концу, и, либо заполучив необходимое, либо найдя иной способ производства высокотехнологичных устройств, Reverses прекращает закупку товара “<span class=\"importantly\">субатомные композиты</span>”. Прогнозируется жёсткое снижение цен на данный товар.",
		},
		// оборудование общего назн.
		{
			Fraction:    _const.Reverses,
			ProductName: "general_purpose_equipment",
			Type:        "lack",
			Text:        "В цене выросли товары - “<span class=\"importantly\">оборудование общего назначения</span>”, чью нехватку внезапно ощутили в Reverses. Все синтеты, что сейчас обладает данным товаром, могут предложить его для продажи по весьма выгодным ценам на следующей базе %BaseName% Reverses.",
		}, {
			Fraction:    _const.Reverses,
			ProductName: "general_purpose_equipment",
			Type:        "cancel",
			Text:        "Товары - “<span class=\"importantly\">оборудование общего назначения</span>” более не востребованы в регионах Reverses. Очень жаль, что сами Reverses так и не объяснили неожиданные скачки в ценах на данный распространённый вид товара. Между тем, мы поздравляем всех тех синтетов, что успели поднять на торговле достойные суммы.",
		},
		// накопители
		{
			Fraction:    _const.Reverses,
			ProductName: "energy_storage",
			Type:        "lack",
			Text:        "Самые главные производители и торговцы товарами - “<span class=\"importantly\">накопители энергии</span>”, сообщество Reverses столкнулось с их нехваткой. Всем синтетам, что сейчас имеют в распоряжении данный товар, предлагает выгодно реализовать его на следующей базе %BaseName% Reverses.",
		}, {
			Fraction:    _const.Reverses,
			ProductName: "energy_storage",
			Type:        "cancel",
			Text:        "Должно быть поправив своё дело, Reverses более не заинтересованы в закупке товара - “<span class=\"importantly\">накопители энергии</span>”. Цены вновь выровнялись и теперь, продажа этого товара не сулит сверхприбылей.",
		},
		// ALL
		{
			Fraction:    _const.Empty,
			ProductName: "all",
			Type:        "excess",
			Text:        "Из-за массового перепроизводства на базе %BaseName%, на рынке наблюдается снижение покупной цены на многие товары. Не теряйте время зря и успейте взять то, что вам сейчас необходимо!",
		}, {
			Fraction:    _const.Empty,
			ProductName: "all",
			Type:        "lack",
			Text:        "Из-за спада производственных мощностей на базе %BaseName%, на рынке наблюдается повышение покупной цены на многие товары. Это может стать удачным подспорьем в реализации тех или иных дефицитных товаров. Имейте это ввиду!",
		}, {
			Fraction:    _const.Empty,
			ProductName: "all",
			Type:        "cancel",
			Text:        "",
		},
	},
	_const.ZhCN: {
		// Replicas 分部
		// 武器部件
		{
			Fraction:    _const.Replicas,
			ProductName: "weapon_parts",
			Type:        "lack",
			Text:        "Replicas 公开承认其“无敌”军队缺乏武器部件供应。拥有该商品的合成体可以在此基地 %BaseName% Replicas 出售以获取丰厚利润。",
		}, {
			Fraction:    _const.Replicas,
			ProductName: "weapon_parts",
			Type:        "cancel",
			Text:        "Replicas 终止了“武器部件”的合同并停止进一步采购，对该商品不再感兴趣。",
		},
		// 技术图纸
		{
			Fraction:    _const.Replicas,
			ProductName: "technological_maps",
			Type:        "lack",
			Text:        "在承认技术竞赛失败后，Replicas 再次允许出售高科技产品，尤其是“技术图纸”。拥有该商品的合成体可以在此基地 %BaseName% Replicas 出售以获取丰厚利润。",
		}, {
			Fraction:    _const.Replicas,
			ProductName: "technological_maps",
			Type:        "cancel",
			Text:        "Replicas 再次拒绝合作，并停止“技术图纸”的所有交易。预计将会出现新一轮由 Replicas 引发的技术灾难。",
		},
		// 工业材料
		{
			Fraction:    _const.Replicas,
			ProductName: "industrial_materials",
			Type:        "lack",
			Text:        "今天，Replicas 向所有流浪合成体发出购买“工业材料”的提议，将其归类为“扩张计划”。拥有该商品的合成体可以在此基地 %BaseName% Replicas 出售以获取丰厚利润。",
		}, {
			Fraction:    _const.Replicas,
			ProductName: "industrial_materials",
			Type:        "cancel",
			Text:        "正如预期，短暂而热烈的交易结束后，Replicas 的仓库已满，停止了“工业材料”的采购。",
		},
		// 亚原子复合物
		{
			Fraction:    _const.Replicas,
			ProductName: "subatomic_composites",
			Type:        "lack",
			Text:        "似乎 Replicas 遇到了高精度设备配套危机，重新开放市场进行交易，“亚原子复合物”的价格因此飙升。拥有该商品的合成体可以在此基地 %BaseName% Replicas 出售以获取丰厚利润。",
		}, {
			Fraction:    _const.Replicas,
			ProductName: "subatomic_composites",
			Type:        "cancel",
			Text:        "Replicas 停止了“亚原子复合物”的交易，声称此商品引发了投机和背叛行为。在一系列服务机器人和基地操作员被消灭后，交易被立即终止。",
		},
		// 通用设备
		{
			Fraction:    _const.Replicas,
			ProductName: "general_purpose_equipment",
			Type:        "lack",
			Text:        "Replicas 希望补充库存中的“通用设备”，并将慷慨支付将商品运送到其领土的商人。拥有该商品的合成体可以在此基地 %BaseName% Replicas 出售以获取丰厚利润。",
		}, {
			Fraction:    _const.Replicas,
			ProductName: "general_purpose_equipment",
			Type:        "cancel",
			Text:        "经过长期互利的贸易后，Replicas 停止接收“通用设备”的供应。",
		},
		// 能量存储
		{
			Fraction:    _const.Replicas,
			ProductName: "energy_storage",
			Type:        "lack",
			Text:        "由于某些地区能源危机，Replicas 不得不发布购买“能量存储”的公告。拥有该商品的合成体可以在此基地 %BaseName% Replicas 出售以获取丰厚利润。",
		}, {
			Fraction:    _const.Replicas,
			ProductName: "energy_storage",
			Type:        "cancel",
			Text:        "显然，Replicas 最终解决了能源危机问题，大幅减少了“能量存储”的采购量，导致价格政策迅速下降。",
		},

		// Explores 分部
		// 武器部件
		{
			Fraction:    _const.Explores,
			ProductName: "weapon_parts",
			Type:        "lack",
			Text:        "Explores 发出了一个不同寻常的“武器部件”采购提议，解释为强制性措施。所有拥有该商品的合成体可以在以下基地 %BaseName% Explores 出售以获利。",
		}, {
			Fraction:    _const.Explores,
			ProductName: "weapon_parts",
			Type:        "cancel",
			Text:        "显然，无论 Explores 是否隐藏了冲突或自卫目标，但最近对“武器部件”的大规模采购已经结束。Explores 感谢所有合成体的帮助，并限制了该商品的销售。预计在 Explores 区域内价格将大幅下跌。",
		},
		// 技术图纸
		{
			Fraction:    _const.Explores,
			ProductName: "technological_maps",
			Type:        "lack",
			Text:        "“技术图纸”商品的价格再次上涨。Explores 官方宣布了新的科学突破，并需要大量“技术图纸”。所有拥有该商品的合成体可以在以下基地 %BaseName% Explores 出售以获利。",
		}, {
			Fraction:    _const.Explores,
			ProductName: "technological_maps",
			Type:        "cancel",
			Text:        "在 Explores 区域内，“技术图纸”的交易似乎已经结束。许多合成体被拒绝出售该商品，价格迅速下跌。",
		},
		// 工业材料
		{
			Fraction:    _const.Explores,
			ProductName: "industrial_materials",
			Type:        "lack",
			Text:        "看起来，“工业材料”商品的价格在 Explores 区域内再次上涨。Explores 向追随者保证，这是建造新科学基地的必要步骤。所有拥有该商品的合成体可以在以下基地 %BaseName% Explores 出售以获利。",
		}, {
			Fraction:    _const.Explores,
			ProductName: "industrial_materials",
			Type:        "cancel",
			Text:        "随着目标的完成，Explores 停止了“工业材料”的大规模采购。Explores 感谢所有合成体的帮助，并限制了该商品的销售。预计在 Explores 区域内价格将大幅下跌。",
		},
		// 亚原子复合物
		{
			Fraction:    _const.Explores,
			ProductName: "subatomic_composites",
			Type:        "lack",
			Text:        "需求迫使 Explores 公开向所有从事贸易的合成体求助，要求供应“亚原子复合物”。Explores 表示将慷慨回报那些将“亚原子复合物”运送到其区域的合成体。所有拥有该商品的合成体可以在以下基地 %BaseName% Explores 出售以获利。",
		}, {
			Fraction:    _const.Explores,
			ProductName: "subatomic_composites",
			Type:        "cancel",
			Text:        "在需求满足后，Explores 限制了“亚原子复合物”的采购。预计在 Explores 区域内价格将大幅下跌。",
		},
		// 通用设备
		{
			Fraction:    _const.Explores,
			ProductName: "general_purpose_equipment",
			Type:        "lack",
			Text:        "拥有“通用设备”商品的合成体有机会在 Explores 领土内出售以获利。这源于今天 Explores 的交易所报告。所有拥有该商品的合成体可以在以下基地 %BaseName% Explores 出售以获利。",
		}, {
			Fraction:    _const.Explores,
			ProductName: "general_purpose_equipment",
			Type:        "cancel",
			Text:        "得知“通用设备”商品的无节制销售后，Explores 临时接管了流通和价格调控，使得进一步出售该商品无利可图。",
		},
		// 能量存储
		{
			Fraction:    _const.Explores,
			ProductName: "energy_storage",
			Type:        "lack",
			Text:        "第三方消息称，Explores 地区缺乏优质“能量存储”商品。后来，Explores 的代表澄清了这一点，表示需要供应这种昂贵且必要的商品。所有拥有该商品的合成体可以在以下基地 %BaseName% Explores 出售以获利。",
		}, {
			Fraction:    _const.Explores,
			ProductName: "energy_storage",
			Type:        "cancel",
			Text:        "看起来，“能量存储”商品的危机终于过去了。该商品的需求下降，使其在 Explores 区域内的进一步销售变得无利可图。",
		},

		// Reverses 分部
		// 武器部件
		{
			Fraction:    _const.Reverses,
			ProductName: "weapon_parts",
			Type:        "lack",
			Text:        "据非官方消息，Reverses 对购买“武器部件”感兴趣。尽管具体用途不明，但任何合成体现在都可以通过在以下基地 %BaseName% Reverses 提供商品来获利。",
		}, {
			Fraction:    _const.Reverses,
			ProductName: "weapon_parts",
			Type:        "cancel",
			Text:        "漏洞被发现并关闭，Reverses 切断了所有“武器部件”的供应渠道，市场崩溃。很遗憾，价格再次下跌。",
		},
		// 技术图纸
		{
			Fraction:    _const.Reverses,
			ProductName: "technological_maps",
			Type:        "lack",
			Text:        "令人意外的是，Reverses 提出了一项关于大规模采购“技术图纸”的提议，并补充说这将服务于其长期计划。拥有该商品的合成体可以在以下基地 %BaseName% Reverses 出售以获利。",
		}, {
			Fraction:    _const.Reverses,
			ProductName: "technological_maps",
			Type:        "cancel",
			Text:        "看来，Reverses 的科研活动已经结束。“技术图纸”的价格迅速下跌，未来可能还会多次波动。",
		},
		// 工业材料
		{
			Fraction:    _const.Reverses,
			ProductName: "industrial_materials",
			Type:        "lack",
			Text:        "显然，Reverses 认真开始了新区的扩展，从今天起提高了“工业材料”的价格。如果您希望支持 Reverses 并从中获利，建议所有合成体访问以下基地 %BaseName% Reverses 出售该商品。",
		}, {
			Fraction:    _const.Reverses,
			ProductName: "industrial_materials",
			Type:        "cancel",
			Text:        "早晚会结束，这次也影响到了 Reverses 对“工业材料”的兴趣。未能及时出售商品的合成体会感到失望，因为价格开始急剧下跌。",
		},
		// 亚原子复合物
		{
			Fraction:    _const.Reverses,
			ProductName: "subatomic_composites",
			Type:        "lack",
			Text:        "难以置信，Reverses 遇到了“亚原子复合物”的短缺。目前拥有这种昂贵商品的合成体可以在以下基地 %BaseName% Reverses 出售以获利。",
		}, {
			Fraction:    _const.Reverses,
			ProductName: "subatomic_composites",
			Type:        "cancel",
			Text:        "销售高峰已过，Reverses 或许已获得所需物资或找到了其他制造高科技设备的方法，停止了“亚原子复合物”的采购。预计该商品价格将大幅下降。",
		},
		// 通用设备
		{
			Fraction:    _const.Reverses,
			ProductName: "general_purpose_equipment",
			Type:        "lack",
			Text:        "“通用设备”的价格上涨，Reverses 突然感受到其短缺。目前拥有该商品的合成体可以在以下基地 %BaseName% Reverses 出售以获利。",
		}, {
			Fraction:    _const.Reverses,
			ProductName: "general_purpose_equipment",
			Type:        "cancel",
			Text:        "“通用设备”商品在 Reverses 区域内不再需求。很遗憾，Reverses 自己没有解释这一常见商品的价格波动原因。同时，我们祝贺那些通过交易赚取了可观收入的合成体。",
		},
		// 能量存储
		{
			Fraction:    _const.Reverses,
			ProductName: "energy_storage",
			Type:        "lack",
			Text:        "作为主要生产和贸易商，Reverses 社区面临“能量存储”的短缺。目前拥有该商品的合成体可以在以下基地 %BaseName% Reverses 出售以获利。",
		}, {
			Fraction:    _const.Reverses,
			ProductName: "energy_storage",
			Type:        "cancel",
			Text:        "可能是业务改善后，Reverses 不再对采购“能量存储”感兴趣。价格再次趋于平稳，现在出售该商品无法带来超额利润。",
		},

		// 全部商品
		{
			Fraction:    _const.Empty,
			ProductName: "all",
			Type:        "excess",
			Text:        "由于基地 %BaseName% 的大规模生产过剩，市场上许多商品的采购价格下降。不要浪费时间，尽快获取您需要的东西！",
		}, {
			Fraction:    _const.Empty,
			ProductName: "all",
			Type:        "lack",
			Text:        "由于基地 %BaseName% 的生产能力下降，市场上许多商品的采购价格上升。这对销售稀缺商品是一个很好的机会，请记住这一点！",
		}, {
			Fraction:    _const.Empty,
			ProductName: "all",
			Type:        "cancel",
			Text:        "",
		},
	},
}

func GetTradeEvent(fraction, productName, typeEvent, lang string) string {
	texts := make([]TradeEvent, 0)

	for _, t := range TradeEvents[lang] {
		if t.Fraction == fraction && t.ProductName == productName && t.Type == typeEvent {
			texts = append(texts, t)
		}
	}

	if len(texts) == 0 {
		println("no text trade event:", fraction, productName, typeEvent, lang)
		return ""
	}

	return texts[rand.Intn(len(texts))].Text
}
