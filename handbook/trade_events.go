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
