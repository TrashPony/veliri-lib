package descriptions

import _const "github.com/TrashPony/veliri-lib/const"

var StructureDescription = map[string]map[string]DescriptionItem{
	_const.EN: {
		"missile_defense":      {Name: "Missile defense", Description: "<p> An engineering structure for military purposes, tracks and counteracts the approach of unidentified or recognized as \"enemy\" missile-type projectiles. </p> <p> Like other structures, it functions by supplying energy. </p>"},
		"antenna_z-x":          {Name: "Antenna Z-X: 5", Description: "<p> A unique and perfectly preserved structure after the local end of the world, which by the appearance of its covering strongly resembles insect honeycombs </p>"},
		"explores_antenna":     {Name: "Antenna \"Veliri\"", Description: "<p> Part of the Veliri Explores tracking station, serves as a super long-range radar. </p>"},
		"explores_observatory": {Name: "Observatory \"Veliri\"", Description: "<p> Part of Veliri Explores Tracking Station, may serve as a watchtower. </p>"},
		"repair_station":       {Name: "Repair station", Description: "<p> An engineering structure used to carry out repair operations within its own radius of action and only allied signatures. </p> <p> Like other structures, it functions by supplying energy. </p>"},
		"laser_turret":         {Name: "Large laser turret", Description: "<p> A stationary automated defensive position that fires at enemy targets by emitting a high-energy beam. </p>"},
		"shield_generator":     {Name: "Force field generator", Description: "<p> An engineering structure for military purposes, serves as a defensive measure during military or civil conflicts. \"Force Field Generator\" - emits low-phase flows of energy, which forms a protective dome in a certain radius, which reduces any damage received during its work. </p> <p> Like other structures, it functions by supplying energy. </p>"},
		"jammer_generator":     {Name: "Jammer", Description: "<p> An engineering structure for military purposes, which, when activated, causes radio interference and interferes with the operation of any tracking electronics. </p>"},
		"extractor":            {Name: "Multipurpose mining rig", Description: "<p> An engineering structure of a civilian type, which is used purely for the extraction of mineral resources in an automated mode. The \"mining rig\" is defenseless and unable to withstand much damage. When filling the internal warehouse and not emptying it, it stops working. </p> <p> Like other structures, it functions by supplying energy. </p>"},
		"replic_gauss_gun":     {Name: "Gun \"Honchkins\"", Description: "<p> Ultra-Long Range Electromagnetic Artillery </p>"},
		"tank_turret":          {Name: "Large ballistic turret", Description: "<p> A stationary-automated defensive position firing large ballistic projectiles at enemy targets. </p>"},
		"artillery_turret":     {Name: "Artillery turret", Description: "<p> A stationary automated defensive position that tracks and attacks enemy targets with devastating artillery shells. </p>"},
		"energy_generator":     {Name: "Power generator", Description: "<p> An engineering structure of a fuel-exchange nature - powering buildings and devices around it within a certain radius of action. </p> <p> The structure of the structure is fragile, in order to maintain efficiency, constant refueling is required. </p>"},
		"radar":                {Name: "Radar", Description: "<p> An engineering structure, which dramatically increases the detection area of ​​any signatures and is able to warn in advance about the approach of unfriendly targets. </p> <p> Like other structures, it functions by supplying energy. </p>"},
		"storage":              {Name: "Warehouse", Description: "<p> An engineering structure serving with one single purpose - the storage of the cargo of its owner. The complex opening mechanism and dense walls make the \"warehouse\" an inaccessible target for pirates. </p> <p> Meanwhile, like other objects, the \"warehouse\" is vulnerable to the effect of \"destruction\". </p>"},
		"beacon":               {Name: "Faction beacon", Description: "<p> An engineering structure for military purposes, which is located on neutral, or enemy territory, which contributes to the beginning of its expansion. As a rule, in the course of conflicts, the \"lighthouse\" is the primary target of the attack, because it is the \"lighthouse\" that is the very first outpost where the attacking forces pull their reinforcements. </p>"},

		"corporation_base_1":          {Name: "Base", Description: "<p>The core of any large-scale structure in the wasteland. Prepared slipways, power drives and territory cleared for extraction of primary materials, which will later take on the appearance of a full-fledged and multifunctional base.</p><p>Like other similar structures of synthetics, it provides minimal support in the form of repair, refueling and storage.</p>"},
		"corporation_processing_1":    {Name: "Recycler", Description: "<p>An auxiliary resource processing structure — a base module responsible for analyzing items and crushing rocks mined by the player/bot to obtain raw materials.</p>"},
		"corporation_office_1":        {Name: "Office", Description: "<p>Auxiliary administrative structure — a base module responsible for management services. In particular, the 'Office' allows you to regulate the provision of warehouse work and transfer of synthetics to the base.</p>"},
		"corporation_prefabricated_1": {Name: "Factory", Description: "<p>An auxiliary production structure — a base module responsible for the manufacture of various other section modules, as well as parts/items often ordered for individual needs of the player.</p>"},
		"corporation_market_1":        {Name: "Market", Description: "<p><p>Auxiliary economic structure — a base module that is a large-scale trading platform where players make transactions for buying/selling all kinds of things.</p></p>"},
	},
	_const.RU: {
		"laser_turret":         {Name: "Большая лазерная турель", Description: "<p>Стационарно-автоматизированная оборонительная позиция, ведущая огонь по целям вражеского характера посредством испускания высокоэнергетического луча.</p>"},
		"tank_turret":          {Name: "Большая балистическая турель", Description: "<p>Стационарно-автоматизированная оборонительная позиция, ведущая огонь по целям вражеского характера крупными баллистическими снарядами.</p>"},
		"artillery_turret":     {Name: "Артиллерийская турель", Description: "<p>Стационарно-автоматизированная оборонительная позиция, отслеживающая и атакующая цели вражеского характера посредством разрушительных артиллерийских снарядов.</p>"},
		"shield_generator":     {Name: "Генератор силового поля", Description: "<p>Инженерное сооружение военного назначения, служит оборонительной мерой в ходе военных или гражданских конфликтов. «Генератор Силового Поля» - испускает малофазовые потоки энергии, чем образует в определённом радиусе защитный купол, чем снижает любой получаемый урон в ходе своей работы. </p>"},
		"energy_generator":     {Name: "Генератор энергии", Description: "<p>Инженерное сооружение топливо-обменного характера - запитывающее постройки и устройства вокруг себя в определённом радиусе действия. </p><p> Структура сооружения хрупка, ради поддержания работоспособности требуется постоянная дозаправка. </p>"},
		"jammer_generator":     {Name: "Генератор помех", Description: "<p>Инженерное сооружение военного назначения, при активации ставящее радиопомехи и мешающее работе любой электроники отслеживающего характера.  </p>"},
		"missile_defense":      {Name: "ПРО", Description: "<p>Инженерное сооружение военного назначения, отслеживает и противодействует подлёту неопознанных или распознанных как «вражеских» снарядов ракетного типа. </p><p> Как и прочие сооружения, функционирует за счёт запитывания энергией.</p>"},
		"radar":                {Name: "Радар", Description: "<p>Инженерное сооружение, что разительно увеличивает область обнаружения любых сигнатур и способно заранее предупредить о приближении недружелюбных целей. </p><p> Как и прочие сооружения, функционирует за счёт запитывания энергией.</p>"},
		"storage":              {Name: "Склад", Description: "<p>Инженерное сооружение, служащее с одной единственной целью – хранением груза его владельца. Сложный механизм открытия и плотные стенки, делают «склад» труднодоступной целью для пиратов. </p><p> Между тем, как и прочие объекты «склад» уязвим к эффекту «разрушения».</p>"},
		"beacon":               {Name: "Фракционный маяк", Description: "<p>Инженерное сооружение военного назначения, что размещается на нейтральной, или вражеской территории, чем способствует началу её экспансии.  Как правило, в ходе конфликтов «маяк» - первоочередная цель атаки, ведь именно «маяк» является тем самым первым форпостом, куда атакующие силы стягивают свои подкрепления.</p>"},
		"extractor":            {Name: "Универсально добывающая установка", Description: "<p>Инженерное сооружение гражданского образца, что используется сугубо при добыче ископаемых ресурсов в автоматизированном режиме. «Добывающая установка» - беззащитна и не способна выдержать большого урона. При заполнении внутреннего склада и его не опустошении, прекращает свою работу. </p>"},
		"repair_station":       {Name: "Ремонтная станция", Description: "<p>Инженерное сооружение, используемое для проведения ремонтных операций в радиусе собственного действия и только союзных сигнатур. </p><p> Как и прочие сооружения, функционирует за счёт запитывания энергией.</p>"},
		"replic_gauss_gun":     {Name: "Орудие \"Гончкинс\"", Description: "<p>Сверхдальняя электромагнитная артиллерия</p>"},
		"antenna_z-x":          {Name: "Антенна Z-X:5", Description: "<p>Уникальное и прекрасно сохранившееся после наступления местного конца света конструкция, что внешним видом своего покрытия, сильно напоминает соты насекомых</p>"},
		"explores_antenna":     {Name: "Антенна \"Велири\"", Description: "<p>Часть станции отслеживания Explores \"Велири\", служит как сверх дальний радар.</p>"},
		"explores_observatory": {Name: "Обсерватория \"Велири\"", Description: "<p>Часть станции отслеживания Explores \"Велири\", может служить как дозорная башня.</p>"},

		"corporation_base_1":          {Name: "База", Description: "<p>Ядро любого масштабного сооружения в пустошах. Заготовленные стапели, энергоприводы и расчищенная под добычу первичных материалов территория, что в дальнейшем примет облик полноценный и многофункциональной базы.</p><p>Как и прочие подобные сооружения синтетов, обеспечивает минимальную поддержку в виде: ремонта, дозаправки и хранилища.</p>\n"},
		"corporation_processing_1":    {Name: "Переработчик", Description: "<p>Вспомогательное сооружения переработки ресурсов – модуль базы, отвечающий за разбор предметов и измельчение добытых игроком/ботом горных пород ради поучения сырья.</p>"},
		"corporation_office_1":        {Name: "Офис", Description: "<p>Вспомогательное сооружения административного толка – модуль базы, отвечающий за управленческие услуги. В частности, «Офис» позволяет регулировать предоставление работы склада и переноса синтетов на базу.</p>"},
		"corporation_prefabricated_1": {Name: "Завод", Description: "<p>Вспомогательное сооружение производственного толка – модуль базы, отвечающий за изготовление различных иных секций-модулей, а также деталей/предметов, нередко заказываемых под индивидуальные нужды игрока.</p>"},
		"corporation_market_1":        {Name: "Рынок", Description: "<p><p>Вспомогательное сооружение экономического толка – модуль базы, представляющий из себя масштабную торговую площадку, где игроки совершают сделки по купле/продаже всевозможных вещей.</p></p>"},
	},
}