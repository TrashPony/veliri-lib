package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

// WarBoard тексты войны и ее экономики (docs/WAR_ECONOMY.md): окно WarBoard и подсказки в лобби, магазине, верстаке, лаборатории, карте.
// Подстановки в формате %name% (значения подставляет static/src/components/War/war_texts.js).
var WarBoard = map[string]map[string]string{
	"title": {
		_const.RU:   `Война: сезон и бонусы`,
		_const.EN:   `War: season and bonuses`,
		_const.ZhCN: `战争:赛季与加成`,
	},
	"tab_summary": {
		_const.RU:   `Сводка`,
		_const.EN:   `Summary`,
		_const.ZhCN: `概览`,
	},
	"tab_bases": {
		_const.RU:   `Базы`,
		_const.EN:   `Bases`,
		_const.ZhCN: `基地`,
	},
	"tab_leaders": {
		_const.RU:   `Лидеры`,
		_const.EN:   `Leaders`,
		_const.ZhCN: `排行榜`,
	},
	"season": {
		_const.RU:   `Сезон`,
		_const.EN:   `Season`,
		_const.ZhCN: `赛季`,
	},
	"ends_in": {
		_const.RU:   `до конца`,
		_const.EN:   `ends in`,
		_const.ZhCN: `剩余`,
	},
	"prev_winner": {
		_const.RU:   `Победитель прошлого сезона`,
		_const.EN:   `Last season winner`,
		_const.ZhCN: `上赛季冠军`,
	},
	"prev_underdog": {
		_const.RU:   `утешительный бонус`,
		_const.EN:   `consolation bonus`,
		_const.ZhCN: `安慰奖励`,
	},
	"leaders": {
		_const.RU:   `Лидеры сезона`,
		_const.EN:   `Season leaders`,
		_const.ZhCN: `赛季排行榜`,
	},
	"best_supplier": {
		_const.RU:   `Лучший снабженец`,
		_const.EN:   `Top supplier`,
		_const.ZhCN: `最佳补给员`,
	},
	"best_fighter": {
		_const.RU:   `Лучший боец`,
		_const.EN:   `Top fighter`,
		_const.ZhCN: `最佳战士`,
	},
	"no_leaders": {
		_const.RU:   `пока никого`,
		_const.EN:   `nobody yet`,
		_const.ZhCN: `暂无`,
	},
	"my_deliveries": {
		_const.RU:   `Ваше участие`,
		_const.EN:   `Your participation`,
		_const.ZhCN: `您的参与`,
	},
	"part_action": {
		_const.RU:   `Что сделано`,
		_const.EN:   `Action`,
		_const.ZhCN: `行动`,
	},
	"part_points": {
		_const.RU:   `Очки участия`,
		_const.EN:   `Participation points`,
		_const.ZhCN: `参与点数`,
	},
	"part_total": {
		_const.RU:   `Всего`,
		_const.EN:   `Total`,
		_const.ZhCN: `合计`,
	},
	"part_none": {
		_const.RU:   `Пока ничего: воюйте, захватывайте точки и возите снабжение.`,
		_const.EN:   `Nothing yet: fight, capture points and deliver supplies.`,
		_const.ZhCN: `暂无:参与战斗、占领据点、运送补给。`,
	},
	"my_sources": {
		_const.RU:   `откуда`,
		_const.EN:   `from`,
		_const.ZhCN: `来源`,
	},
	"fraction_points": {
		_const.RU:   `Фракционные очки`,
		_const.EN:   `Fraction points`,
		_const.ZhCN: `阵营点数`,
	},
	"how_score": {
		_const.RU:   `Рейтинг: +1/мин за удержание базы | +10 за захват базы | +50 за захват сектора`,
		_const.EN:   `Score: +1/min for holding a base | +10 for capturing a base | +50 for capturing a sector`,
		_const.ZhCN: `积分:每分钟持有据点+1 | 占领据点+10 | 占领区域+50`,
	},
	"bonuses": {
		_const.RU:   `Бонусы вашей фракции`,
		_const.EN:   `Your fraction bonuses`,
		_const.ZhCN: `您阵营的加成`,
	},
	"no_bonuses": {
		_const.RU:   `Пока нет: захватывайте сектора, чтобы открыть бонусы.`,
		_const.EN:   `None yet: capture sectors to unlock bonuses.`,
		_const.ZhCN: `暂无:占领区域以解锁加成。`,
	},
	"sectors": {
		_const.RU:   `Сектора войны`,
		_const.EN:   `War sectors`,
		_const.ZhCN: `战争区域`,
	},
	"nobody": {
		_const.RU:   `никто`,
		_const.EN:   `nobody`,
		_const.ZhCN: `无人`,
	},
	"bases": {
		_const.RU:   `Военные базы`,
		_const.EN:   ``,
		_const.ZhCN: ``,
	},
	"supply": {
		_const.RU:   `снабжение`,
		_const.EN:   `supply`,
		_const.ZhCN: `补给`,
	},
	"supply_hint": {
		_const.RU:   `Товары пополняют снабжение базы и засчитывается вам как участие.`,
		_const.EN:   ``,
		_const.ZhCN: ``,
	},
	"supply_freeze": {
		_const.RU:   `Укрепление не растет`,
		_const.EN:   `Fortification does not grow`,
		_const.ZhCN: `防御工事不再成长`,
	},
	"supply_no_repair": {
		_const.RU:   `ремонтные станции отключены`,
		_const.EN:   `repair stations are offline`,
		_const.ZhCN: `维修站已关闭`,
	},
	"supply_decay": {
		_const.RU:   `прочность строений падает до 25%`,
		_const.EN:   `structure HP drops down to 25%`,
		_const.ZhCN: `建筑耐久降至25%`,
	},
	"garrison": {
		_const.RU:   `Гарнизон`,
		_const.EN:   `Garrison`,
		_const.ZhCN: `驻军`,
	},
	"garrison_hint": {
		_const.RU:   `Гарнизон базы: бойцов сейчас / положено.<br>Размер зависит от снабжения:<br>выше 66% - 3 бойца<br>выше 33% - 2 бойца<br>ниже - 1 боец`,
		_const.EN:   `Base garrison: fighters now / allowed.<br>Size depends on supply:<br>above 66% - 3 fighters<br>above 33% - 2 fighters<br>below - 1 fighter`,
		_const.ZhCN: `基地驻军:当前 / 上限人数。<br>规模取决于补给:<br>高于66% - 3人<br>高于33% - 2人<br>否则 - 1人`,
	},
	"caravan": {
		_const.RU:   `Караван`,
		_const.EN:   `Caravan`,
		_const.ZhCN: `商队`,
	},
	"caravan_hint": {
		_const.RU:   `Караван снабжения: возит товары с мирных баз фракции на эту базу.<br>Бойцов и хуллеров сейчас / в составе.<br>Если караван уничтожен, новый появится через 10 минут.`,
		_const.EN:   `Supply caravan: carries goods from peaceful fraction bases to this base.<br>Fighters and haulers now / in the group.<br>If the caravan is destroyed, a new one appears in 10 minutes.`,
		_const.ZhCN: `补给商队:把货物从阵营和平基地运往此基地。<br>当前 / 编制人数。<br>商队被摧毁后,10分钟后会出现新的商队。`,
	},
	"caravan_respawn": {
		_const.RU:   `через %n% мин`,
		_const.EN:   `in %n% min`,
		_const.ZhCN: `%n%分钟后`,
	},
	"needs": {
		_const.RU:   `Нужно`,
		_const.EN:   `Needed`,
		_const.ZhCN: `需要`,
	},
	"you_here": {
		_const.RU:   `вы здесь`,
		_const.EN:   `you are here`,
		_const.ZhCN: `您在此处`,
	},
	"src_kill": {
		_const.RU:   `убийства`,
		_const.EN:   `kills`,
		_const.ZhCN: `击杀`,
	},
	"src_assist": {
		_const.RU:   `помощь в бою`,
		_const.EN:   `assists`,
		_const.ZhCN: `助攻`,
	},
	"src_damage": {
		_const.RU:   `урон`,
		_const.EN:   `damage`,
		_const.ZhCN: `伤害`,
	},
	"src_capture": {
		_const.RU:   `захват точек`,
		_const.EN:   `point capture`,
		_const.ZhCN: `占领据点`,
	},
	"src_capture_sector": {
		_const.RU:   `захват секторов`,
		_const.EN:   `sector capture`,
		_const.ZhCN: `占领区域`,
	},
	"src_destroy_structure": {
		_const.RU:   `разрушение строений`,
		_const.EN:   `structures destroyed`,
		_const.ZhCN: `摧毁建筑`,
	},
	"src_war_supply": {
		_const.RU:   `снабжение`,
		_const.EN:   `supply`,
		_const.ZhCN: `补给`,
	},
	"src_dialog": {
		_const.RU:   `задания`,
		_const.EN:   `quests`,
		_const.ZhCN: `任务`,
	},
	"src_battle": {
		_const.RU:   `бой`,
		_const.EN:   `combat`,
		_const.ZhCN: `战斗`,
	},
	"fraction_Replics": {
		_const.RU:   `Replics`,
		_const.EN:   `Replics`,
		_const.ZhCN: `Replics`,
	},
	"fraction_Explores": {
		_const.RU:   `Explores`,
		_const.EN:   `Explores`,
		_const.ZhCN: `Explores`,
	},
	"fraction_Reverses": {
		_const.RU:   `Reverses`,
		_const.EN:   `Reverses`,
		_const.ZhCN: `Reverses`,
	},
	"fraction_APD": {
		_const.RU:   `APD`,
		_const.EN:   `APD`,
		_const.ZhCN: `APD`,
	},
	"spec_title": {
		_const.RU:   `Специализация`,
		_const.EN:   `Specialization`,
		_const.ZhCN: `专业化`,
	},
	"spec_industry": {
		_const.RU:   `Промышленность`,
		_const.EN:   `Industry`,
		_const.ZhCN: `工业`,
	},
	"spec_industry_hint": {
		_const.RU:   `Снижает налог верстака, сборочного цеха и переработки и дает скидку на очки в фракционном магазине. Часть позиций фракционного магазина (например чертежи) открывается, только когда фракция держит нужное число таких секторов.`,
		_const.EN:   ``,
		_const.ZhCN: ``,
	},
	"spec_science": {
		_const.RU:   `Наука`,
		_const.EN:   `Science`,
		_const.ZhCN: `科研`,
	},
	"spec_science_hint": {
		_const.RU:   `Владение удешевляет RFF для апгрейда чертежей в лаборатории и повышает шанс успеха.`,
		_const.EN:   `Owning it lowers the RFF cost of blueprint upgrades in the laboratory and raises the success chance.`,
		_const.ZhCN: `占领后降低实验室升级蓝图所需的RFF并提高成功率。`,
	},
	"spec_shipyard": {
		_const.RU:   `Машиностроение`,
		_const.EN:   `Engineering`,
		_const.ZhCN: `机械制造`,
	},
	"spec_shipyard_hint": {
		_const.RU:   `Владение открывает T2 детали в сборочном цехе.`,
		_const.EN:   `Owning it unlocks T2 details in the prefab plant.`,
		_const.ZhCN: `占领后解锁装配车间的T2部件。`,
	},
	"spec_trade": {
		_const.RU:   `Торговля`,
		_const.EN:   `Trade`,
		_const.ZhCN: `贸易`,
	},
	"spec_trade_hint": {
		_const.RU:   `Владение снижает комиссию рыночных ордеров и повышает награду за задания.`,
		_const.EN:   `Owning it lowers the market order fee and raises mission rewards.`,
		_const.ZhCN: `占领后降低市场订单手续费并提高任务奖励。`,
	},
	"mod_work_tax": {
		_const.RU:   `Налог верстака`,
		_const.EN:   `Workbench tax`,
		_const.ZhCN: `工作台税`,
	},
	"mod_detail_tax": {
		_const.RU:   `Налог сборочного цеха`,
		_const.EN:   `Prefab plant tax`,
		_const.ZhCN: `装配车间税`,
	},
	"mod_recycle_tax": {
		_const.RU:   `Налог переработки`,
		_const.EN:   `Recycling tax`,
		_const.ZhCN: `回收税`,
	},
	"mod_market_fee": {
		_const.RU:   `Комиссия рынка`,
		_const.EN:   `Market fee`,
		_const.ZhCN: `市场手续费`,
	},
	"mod_store_discount": {
		_const.RU:   `Скидка в фракционном магазине`,
		_const.EN:   `Fraction store discount`,
		_const.ZhCN: `阵营商店折扣`,
	},
	"mod_lab_cost": {
		_const.RU:   `Стоимость RFF в лаборатории`,
		_const.EN:   `Laboratory RFF cost`,
		_const.ZhCN: `实验室RFF成本`,
	},
	"mod_mission_reward": {
		_const.RU:   `Награда за задания`,
		_const.EN:   `Mission reward`,
		_const.ZhCN: `任务奖励`,
	},
	"mod_lab_chance": {
		_const.RU:   `Шанс апгрейда в лаборатории`,
		_const.EN:   `Laboratory upgrade chance`,
		_const.ZhCN: `实验室升级成功率`,
	},
	"mod_unlock_prefab_t2": {
		_const.RU:   `T2 детали в сборочном цехе открыты`,
		_const.EN:   `T2 details in the prefab plant unlocked`,
		_const.ZhCN: `装配车间T2部件已解锁`,
	},
	"unit_pp": {
		_const.RU:   `%`,
		_const.EN:   ` `,
		_const.ZhCN: ``,
	},
	"unit_pct": {
		_const.RU:   `%`,
		_const.EN:   `%`,
		_const.ZhCN: `%`,
	},
	"source_season_winner": {
		_const.RU:   `победа в прошлом сезоне`,
		_const.EN:   `won the last season`,
		_const.ZhCN: `上赛季获胜`,
	},
	"source_season_underdog": {
		_const.RU:   `утешительный бонус прошлого сезона`,
		_const.EN:   `last season consolation bonus`,
		_const.ZhCN: `上赛季安慰奖励`,
	},
	"source_sectors": {
		_const.RU:   `сектора "%spec%": %owned% из %total%`,
		_const.EN:   `"%spec%" sectors: %owned% of %total%`,
		_const.ZhCN: `"%spec%"区域:%owned%/%total%`,
	},
	"source_no_sectors": {
		_const.RU:   `секторов "%spec%" нет, открыто всем`,
		_const.EN:   `no "%spec%" sectors, open to all`,
		_const.ZhCN: `没有"%spec%"区域,对所有人开放`,
	},
	"tax_corp": {
		_const.RU:   `Налог задан политикой корпорации базы: %base%%`,
		_const.EN:   `Tax is set by the base corporation policy: %base%%`,
		_const.ZhCN: `税率由基地所属公司政策决定:%base%%`,
	},
	"tax_base": {
		_const.RU:   `Налог базы: %base%%. Растет, когда на складе базы не хватает товаров (меньше 1000 ед.), в столицах выше.`,
		_const.EN:   `Base tax: %base%%. Grows when the base stock runs low (under 1000 units), higher in capitals.`,
		_const.ZhCN: `基地税:%base%%。基地库存不足(少于1000件)时上涨,首都更高。`,
	},
	"tax_extra": {
		_const.RU:   `+%n% % - надбавка самого чертежа`,
		_const.EN:   `+%n% pp - blueprint surcharge`,
		_const.ZhCN: `+%n%个百分点 - 蓝图附加费`,
	},
	"tax_relation": {
		_const.RU:   `-%n% % - отношения с фракцией "%fraction%" (каждые 20 очков отношений -1)`,
		_const.EN:   `-%n% pp - relations with "%fraction%" (-1 per 20 relation points)`,
		_const.ZhCN: `-%n%个百分点 - 与"%fraction%"的关系(每20点关系-1)`,
	},
	"tax_war": {
		_const.RU:   `-%n% % - бонус войны`,
		_const.EN:   `-%n% pp - war bonus`,
		_const.ZhCN: `-%n%个百分点 - 战争加成`,
	},
	"tax_war_why": {
		_const.RU:   `-%n% % - бонус войны (%why%)`,
		_const.EN:   `-%n% pp - war bonus (%why%)`,
		_const.ZhCN: `-%n%个百分点 - 战争加成(%why%)`,
	},
	"tax_floor": {
		_const.RU:   `Бонус войны не опускает налог ниже %floor%%`,
		_const.EN:   `The war bonus never lowers tax below %floor%%`,
		_const.ZhCN: `战争加成不会使税率低于%floor%%`,
	},
	"tax_skill": {
		_const.RU:   `-%n% % - ваш навык производства (при расчете материалов)`,
		_const.EN:   `-%n% pp - your production skill (applied to materials)`,
		_const.ZhCN: `-%n%个百分点 - 您的生产技能(计入材料)`,
	},
	"tax_total_work": {
		_const.RU:   `Итого расход материалов: %n%%`,
		_const.EN:   `Total materials cost: %n%%`,
		_const.ZhCN: `材料消耗合计:%n%%`,
	},
	"tax_total_detail": {
		_const.RU:   `Итого расход ресурсов: %n%%`,
		_const.EN:   `Total resource cost: %n%%`,
		_const.ZhCN: `资源消耗合计:%n%%`,
	},
	"tax_total": {
		_const.RU:   `Итого: %n%%`,
		_const.EN:   `Total: %n%%`,
		_const.ZhCN: `合计:%n%%`,
	},
	"lab_chance_base": {
		_const.RU:   `Базовый шанс: %n%%`,
		_const.EN:   `Base chance: %n%%`,
		_const.ZhCN: `基础成功率:%n%%`,
	},
	"lab_chance_pirate": {
		_const.RU:   `Базовый шанс пиратской лаборатории: %n%%`,
		_const.EN:   `Pirate laboratory base chance: %n%%`,
		_const.ZhCN: `海盗实验室基础成功率:%n%%`,
	},
	"lab_chance_relation": {
		_const.RU:   `+%n% % - отношения с фракцией`,
		_const.EN:   `+%n% pp - fraction relations`,
		_const.ZhCN: `+%n%个百分点 - 阵营关系`,
	},
	"lab_chance_war": {
		_const.RU:   `+%n% % - бонус войны (%why%)`,
		_const.EN:   `+%n% pp - war bonus (%why%)`,
		_const.ZhCN: `+%n%个百分点 - 战争加成(%why%)`,
	},
	"lab_chance_total": {
		_const.RU:   `Итого: %n%%`,
		_const.EN:   `Total: %n%%`,
		_const.ZhCN: `合计:%n%%`,
	},
	"lab_cost_full": {
		_const.RU:   `Полный рецепт RFF: %n%%`,
		_const.EN:   `Full RFF recipe: %n%%`,
		_const.ZhCN: `完整RFF配方:%n%%`,
	},
	"lab_cost_pirate": {
		_const.RU:   `Пиратская лаборатория берет %n%% от рецепта RFF`,
		_const.EN:   `Pirate laboratory takes %n%% of the RFF recipe`,
		_const.ZhCN: `海盗实验室只需RFF配方的%n%%`,
	},
	"lab_cost_war": {
		_const.RU:   `-%n%% - бонус войны (%why%)`,
		_const.EN:   `-%n%% - war bonus (%why%)`,
		_const.ZhCN: `-%n%% - 战争加成(%why%)`,
	},
	"lab_cost_total": {
		_const.RU:   `Итого: %n%% от рецепта`,
		_const.EN:   `Total: %n%% of the recipe`,
		_const.ZhCN: `合计:配方的%n%%`,
	},
	"lab_rff_cost": {
		_const.RU:   `Стоимость RFF`,
		_const.EN:   `RFF cost`,
		_const.ZhCN: `RFF成本`,
	},
	"war_discount": {
		_const.RU:   `Скидка от войны`,
		_const.EN:   `War discount`,
		_const.ZhCN: `战争折扣`,
	},
	"empty_category": {
		_const.RU:   `В этой категории пока ничего нет`,
		_const.EN:   `Nothing in this category yet`,
		_const.ZhCN: `此分类暂无商品`,
	},
	"need_sectors_one": {
		_const.RU:   `Необходим %n% сектор "%name%"`,
		_const.EN:   `Requires %n% "%name%" sector`,
		_const.ZhCN: `需要%n%个"%name%"区域`,
	},
	"need_sectors_few": {
		_const.RU:   `Необходимо %n% сектора "%name%"`,
		_const.EN:   `Requires %n% "%name%" sectors`,
		_const.ZhCN: `需要%n%个"%name%"区域`,
	},
	"need_sectors_many": {
		_const.RU:   `Необходимо %n% секторов "%name%"`,
		_const.EN:   `Requires %n% "%name%" sectors`,
		_const.ZhCN: `需要%n%个"%name%"区域`,
	},
	"sector_yours": {
		_const.RU:   `ваш ✓`,
		_const.EN:   `yours ✓`,
		_const.ZhCN: `您的 ✓`,
	},
	"sector_not_captured": {
		_const.RU:   `не захвачен`,
		_const.EN:   `not captured`,
		_const.ZhCN: `未占领`,
	},
	"sector_points": {
		_const.RU:   `точки`,
		_const.EN:   `points`,
		_const.ZhCN: `据点`,
	},
	"locked_by_war": {
		_const.RU:   `закрыто войной`,
		_const.EN:   `locked by war`,
		_const.ZhCN: `被战争锁定`,
	},
	"locked_tip": {
		_const.RU:   `Закрыто войной: фракции нужно удерживать сектор "%name%"`,
		_const.EN:   `Locked by war: your fraction must hold a "%name%" sector`,
		_const.ZhCN: `被战争锁定:阵营需要占领"%name%"区域`,
	},
	"lobby_supply": {
		_const.RU:   `Снабжение`,
		_const.EN:   `Supply`,
		_const.ZhCN: `补给`,
	},
	"lobby_war_bonuses": {
		_const.RU:   `Бонусы войны`,
		_const.EN:   `War bonuses`,
		_const.ZhCN: `战争加成`,
	},
	"lobby_supply_tip": {
		_const.RU:   `Снабжение базы войны: %n%%`,
		_const.EN:   `War base supply: %n%%`,
		_const.ZhCN: `战争基地补给:%n%%`,
	},
	"lobby_supply_garrison": {
		_const.RU:   `Гарнизон: >66% - 3 бойца, >33% - 2, иначе 1`,
		_const.EN:   `Garrison: >66% - 3 fighters, >33% - 2, otherwise 1`,
		_const.ZhCN: `驻军:>66%为3人,>33%为2人,否则1人`,
	},
	"lobby_supply_freeze": {
		_const.RU:   `Укрепление не растет, ремонтные станции отключены`,
		_const.EN:   `Fortification does not grow, repair stations are offline`,
		_const.ZhCN: `防御工事停止成长,维修站已关闭`,
	},
	"lobby_supply_decay": {
		_const.RU:   `Прочность строений падает до 25%`,
		_const.EN:   `Structure HP drops down to 25%`,
		_const.ZhCN: `建筑耐久降至25%`,
	},
	"lobby_supply_sell": {
		_const.RU:   `Сдайте снабженческие товары на этой базе (продажа NPC)`,
		_const.EN:   `Sell supply goods at this base (NPC trade)`,
		_const.ZhCN: `请在此基地出售补给货物(NPC交易)`,
	},
	"map_supply_request": {
		_const.RU:   `Запрос снабжения: базам войны вашей фракции в секторе (%n%) не хватает припасов. Сдайте товар на базе: без снабжения укрепление не растет, а гарнизон слабеет.`,
		_const.EN:   `Supply request: your fraction war bases in this sector (%n%) are short on supplies. Deliver goods to the base: without supplies fortifications stop growing and the garrison weakens.`,
		_const.ZhCN: `补给请求:本区域您阵营的战争基地(%n%)补给不足。请向基地送货:没有补给防御工事不会成长,驻军也会变弱。`,
	},
	"src_kill_hint": {
		_const.RU:   `Уничтожение юнита: 10% от его максимальной прочности.`,
		_const.EN:   `Destroying a unit: 10% of its max HP.`,
		_const.ZhCN: `击毁单位:其最大耐久的10%。`,
	},
	"src_assist_hint": {
		_const.RU:   `Помощь в бою: 10% урона, нанесенного уничтоженному юниту за последние 30 секунд.`,
		_const.EN:   `Assist: 10% of the damage dealt to the destroyed unit in the last 30 seconds.`,
		_const.ZhCN: ``,
	},
	"src_damage_hint": {
		_const.RU:   `Урон: 5% от каждого нанесенного попадания врагам.`,
		_const.EN:   `Damage: 5% of every hit dealt to enemies.`,
		_const.ZhCN: `伤害:对敌人每次命中伤害的5%。`,
	},
	"src_capture_hint": {
		_const.RU:   `Захват точки: 500 очков каждому участнику осады, чей отряд стоял на точке.`,
		_const.EN:   `Point capture: 500 points for each participant of the siege whose squad stood on the point.`,
		_const.ZhCN: `占领据点:每位驻守据点的围攻参与者获得250点。`,
	},
	"src_capture_sector_hint": {
		_const.RU:   `Захват сектора: 3000 очков каждому, кто участвовал в осаде.`,
		_const.EN:   `Sector capture: 3000 points for everyone who took part in the siege.`,
		_const.ZhCN: `占领区域:每位参与围攻者获得1000点。`,
	},
	"src_destroy_structure_hint": {
		_const.RU:   `Разрушение вражеского строения: 10% от его максимальной прочности (как ее видит игрок).`,
		_const.EN:   `Destroying an enemy structure: 10% of its max HP (as the player sees it).`,
		_const.ZhCN: `摧毁敌方建筑:其最大耐久(玩家所见数值)的10%。`,
	},
	"src_war_supply_hint": {
		_const.RU:   `Снабжение: 1% от суммы сделки в кредитах при сдаче товара на базе войны своей фракции.`,
		_const.EN:   ``,
		_const.ZhCN: ``,
	},
	"src_dialog_hint": {
		_const.RU:   `Награды за задания и диалоги.`,
		_const.EN:   `Rewards for quests and dialogs.`,
		_const.ZhCN: `任务和对话奖励。`,
	},
	"src_battle_hint": {
		_const.RU:   `Прочие бои.`,
		_const.EN:   `Other combat.`,
		_const.ZhCN: `其他战斗。`,
	},
}
