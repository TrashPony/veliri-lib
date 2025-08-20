package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var UserStat = map[string]map[string]string{
	// Missions
	"no_missions": {
		_const.RU:   `Нет заданий`,
		_const.EN:   `No missions`,
		_const.ZhCN: `没有任务`,
	},
	"track": {
		_const.RU:   `Отслеживать:`,
		_const.EN:   `Track:`,
		_const.ZhCN: `追踪:`,
	},
	"log": {
		_const.RU:   `Журнал`,
		_const.EN:   `Journal`,
		_const.ZhCN: `日志`,
	},
	"tasks": {
		_const.RU:   `Задачи`,
		_const.EN:   `Tasks`,
		_const.ZhCN: `任务`,
	},
	"cancel": {
		_const.RU:   `Отменить`,
		_const.EN:   `Cancel`,
		_const.ZhCN: `取消`,
	},
	"confirm": {
		_const.RU:   `Подтвердить`,
		_const.EN:   `Confirm`,
		_const.ZhCN: `确认`,
	},
	// MyUser
	"avatar_change": {
		_const.RU:   `Изменить`,
		_const.EN:   `Change`,
		_const.ZhCN: `更改`,
	},
	"fraction_1": {
		_const.RU:   `Фракция:`,
		_const.EN:   `Fraction:`,
		_const.ZhCN: `派系:`,
	},
	"rank_1": {
		_const.RU:   `Ранг:`,
		_const.EN:   `Rank:`,
		_const.ZhCN: `等级:`,
	},
	"cluster_1": {
		_const.RU:   `Кластер:`,
		_const.EN:   `Cluster:`,
		_const.ZhCN: `集群:`,
	},
	"no_cluster_1": {
		_const.RU:   `не состоит в кластере`,
		_const.EN:   `not a member of a corporation`,
		_const.ZhCN: `未加入任何公司`,
	},
	"back_1": {
		_const.RU:   `Назад`,
		_const.EN:   `Back`,
		_const.ZhCN: `返回`,
	},
	"buy": {
		_const.RU:   `Купить`,
		_const.EN:   `Buy`,
		_const.ZhCN: `购买`,
	},
	// OtherUser
	"fraction_2": {
		_const.RU:   `Фракция:`,
		_const.EN:   `Fraction:`,
		_const.ZhCN: `派系:`,
	},
	"cluster_2": {
		_const.RU:   `Кластер:`,
		_const.EN:   `Cluster:`,
		_const.ZhCN: `集群:`,
	},
	"no_cluster_2": {
		_const.RU:   `не состоит в кластере`,
		_const.EN:   `not a member of a corporation`,
		_const.ZhCN: `未加入任何公司`,
	},
	// Relations
	"enmity": {
		_const.RU:   `Вражда`,
		_const.EN:   `Enmity`,
		_const.ZhCN: `敌对`,
	},
	"bad": {
		_const.RU:   `Плохое`,
		_const.EN:   `Bad`,
		_const.ZhCN: `差`,
	},
	"neutral": {
		_const.RU:   `Нейтральное`,
		_const.EN:   `Neutral`,
		_const.ZhCN: `中立`,
	},
	"good": {
		_const.RU:   `Хорошее`,
		_const.EN:   `Good`,
		_const.ZhCN: `良好`,
	},
	"excellent": {
		_const.RU:   `Отличное`,
		_const.EN:   `Excellent`,
		_const.ZhCN: `优秀`,
	},
	"chane_event_1": {
		_const.RU:   `Торговля`,
		_const.EN:   `Trade`,
		_const.ZhCN: `贸易`,
	},
	"chane_event_2": {
		_const.RU:   `Переработка ресурсов`,
		_const.EN:   `Resource Recycling`,
		_const.ZhCN: `资源回收`,
	},
	"chane_event_3": {
		_const.RU:   `Выполнение задания`,
		_const.EN:   `Completing mission`,
		_const.ZhCN: `完成任务`,
	},
	"chane_event_4": {
		_const.RU:   `Провал/отказ от задания`,
		_const.EN:   `Failure/refusal of mission`,
		_const.ZhCN: `任务失败/拒绝`,
	},
	"chane_event_5": {
		_const.RU:   `Устранение пирата`,
		_const.EN:   `Eliminating a pirate`,
		_const.ZhCN: `消灭海盗`,
	},
	"chane_event_6": {
		_const.RU:   `Устранение мирного транспорта`,
		_const.EN:   `Elimination of peaceful transport`,
		_const.ZhCN: `消灭和平运输`,
	},
	"chane_event_7": {
		_const.RU:   `Участие в захвате сектора`,
		_const.EN:   `Participation in the capture of the sector`,
		_const.ZhCN: `参与占领区域`,
	},
	"chane_event_8": {
		_const.RU:   `Конфликт интересов`,
		_const.EN:   `Conflict of interest`,
		_const.ZhCN: `利益冲突`,
	},
	"chane_event_9": {
		_const.RU:   `"Враг моего врага..."`,
		_const.EN:   `"Enemy of my enemy..."`,
		_const.ZhCN: `"敌人的敌人..."`,
	},
	"head_1": {
		_const.RU:   `Отношения с фракциями синтетов`,
		_const.EN:   `Relations with Synthet factions`,
		_const.ZhCN: `与合成体派系的关系`,
	},
	"button_1": {
		_const.RU:   `история`,
		_const.EN:   `log`,
		_const.ZhCN: `日志`,
	},
	"button_2": {
		_const.RU:   `эффекты`,
		_const.EN:   `effects`,
		_const.ZhCN: `效果`,
	},
	"head_2": {
		_const.RU:   `Отношения с другими`,
		_const.EN:   `Relationships with others`,
		_const.ZhCN: `与他人的关系`,
	},
	"head_3": {
		_const.RU:   `Событие`,
		_const.EN:   `Event`,
		_const.ZhCN: `事件`,
	},
	"head_4": {
		_const.RU:   `Изменение`,
		_const.EN:   `Change`,
		_const.ZhCN: `变更`,
	},
	"button_3": {
		_const.RU:   `назад`,
		_const.EN:   `back`,
		_const.ZhCN: `返回`,
	},
	"b_1": {
		_const.RU:   `Обзор фракционных дронов`,
		_const.EN:   `Faction drones view`,
		_const.ZhCN: `派系无人机概览`,
	},
	"b_2": {
		_const.RU:   `Снижение налога на переработку ресурсов`,
		_const.EN:   `Reduced tax on resource processing`,
		_const.ZhCN: `资源处理税减免`,
	},
	"b_3": {
		_const.RU:   `Снижение налога на производство`,
		_const.EN:   `Production tax reduction`,
		_const.ZhCN: `生产税减免`,
	},
	"b_4": {
		_const.RU:   `Снижение налога на производство деталей`,
		_const.EN:   `Reduced tax on parts production`,
		_const.ZhCN: `零件生产税减免`,
	},
	"b_5": {
		_const.RU:   `Снижение налога на рынке`,
		_const.EN:   `Market tax reduction`,
		_const.ZhCN: `市场税减免`,
	},
	"b_6": {
		_const.RU:   `Бонус кредитов за выполнение заданий`,
		_const.EN:   `Bonus credits for completing tasks`,
		_const.ZhCN: `任务完成信用点奖励`,
	},
	"b_7": {
		_const.RU:   `Запрет на торговлю`,
		_const.EN:   `Trade ban`,
		_const.ZhCN: `贸易禁令`,
	},
	"b_8": {
		_const.RU:   `Запрет на производство`,
		_const.EN:   `Production ban`,
		_const.ZhCN: `生产禁令`,
	},
	"b_9": {
		_const.RU:   `Запрет на переработку ресурсов`,
		_const.EN:   `Prohibition on resource processing`,
		_const.ZhCN: `资源处理禁令`,
	},
	"b_10": {
		_const.RU:   `Враг фракции`,
		_const.EN:   `Fraction Enemy`,
		_const.ZhCN: `派系敌人`,
	},
	"b_11": {
		_const.RU:   `Обзор фракционных воинов`,
		_const.EN:   `Faction warriors view`,
		_const.ZhCN: `派系战士概览`,
	},
	// Skills
	"only_studied": {
		_const.RU:   `Только изученные:`,
		_const.EN:   `Only studied:`,
		_const.ZhCN: `仅已学习:`,
	},
	"experience": {
		_const.RU:   `Опыт:`,
		_const.EN:   `Experience:`,
		_const.ZhCN: `经验:`,
	},
	"filter": {
		_const.RU:   `Фильтр`,
		_const.EN:   `Filter`,
		_const.ZhCN: `过滤`,
	},
	"base_alert": {
		_const.RU:   `Внимание! Тренировать навыки можно только на базе.`,
		_const.EN:   `Attention! You can only train skills at the base.`,
		_const.ZhCN: `注意！只能在基地训练技能。`,
	},
	"lvl": {
		_const.RU:   `Уровень:`,
		_const.EN:   `Level:`,
		_const.ZhCN: `等级:`,
	},
	"price": {
		_const.RU:   `Стоимость улучшения:`,
		_const.EN:   `Upgrade cost:`,
		_const.ZhCN: `升级成本:`,
	},
	"bonus_for_lvl": {
		_const.RU:   `- бонус за уровень:`,
		_const.EN:   `- level bonus:`,
		_const.ZhCN: `- 等级奖励:`,
	},
	"bonus_current": {
		_const.RU:   `- текущий бонус:`,
		_const.EN:   `- current bonus:`,
		_const.ZhCN: `- 当前奖励:`,
	},
	"missile": {
		_const.RU:   `Ракетное`,
		_const.EN:   `Missile`,
		_const.ZhCN: `导弹`,
	},
	"mining": {
		_const.RU:   `Добыча полезных ископаемых`,
		_const.EN:   `Mining`,
		_const.ZhCN: `采矿`,
	},
	"production": {
		_const.RU:   `Производство`,
		_const.EN:   `Production`,
		_const.ZhCN: `生产`,
	},
	"processing": {
		_const.RU:   `Переработка полезных ископаемых`,
		_const.EN:   `Processing`,
		_const.ZhCN: `加工`,
	},
	"bodies": {
		_const.RU:   `Корпуса`,
		_const.EN:   `Bodies`,
		_const.ZhCN: `船体`,
	},
	"weapon": {
		_const.RU:   `Оружие`,
		_const.EN:   `Weapon`,
		_const.ZhCN: `武器`,
	},
	"ammo": {
		_const.RU:   `Боеприпасы`,
		_const.EN:   `Ammo`,
		_const.ZhCN: `弹药`,
	},
	"laser": {
		_const.RU:   `Лазерное`,
		_const.EN:   `Laser`,
		_const.ZhCN: `激光`,
	},
	"ballistic": {
		_const.RU:   `Баллистическое`,
		_const.EN:   `Ballistic`,
		_const.ZhCN: `弹道`,
	},
	"small": {
		_const.RU:   `малые`,
		_const.EN:   `small`,
		_const.ZhCN: `小型`,
	},
	"medium": {
		_const.RU:   `средние`,
		_const.EN:   `medium`,
		_const.ZhCN: `中型`,
	},
	"big": {
		_const.RU:   `большие`,
		_const.EN:   `big`,
		_const.ZhCN: `大型`,
	},
	"light": {
		_const.RU:   `легкие`,
		_const.EN:   `light`,
		_const.ZhCN: `轻型`,
	},
	"heavy": {
		_const.RU:   `тяжелые`,
		_const.EN:   `heavy`,
		_const.ZhCN: `重型`,
	},
	"other_1": {
		_const.RU:   `остальные`,
		_const.EN:   `other`,
		_const.ZhCN: `其他`,
	},
	"other_2": {
		_const.RU:   `Остальные`,
		_const.EN:   `Other`,
		_const.ZhCN: `其他`,
	},
	// UserStat
	"tab_1": {
		_const.RU:   `Общие`,
		_const.EN:   `Common`,
		_const.ZhCN: `通用`,
	},
	"tab_2": {
		_const.RU:   `Навыки`,
		_const.EN:   `Skills`,
		_const.ZhCN: `技能`,
	},
	"tab_3": {
		_const.RU:   `Задания`,
		_const.EN:   `Missions`,
		_const.ZhCN: `任务`,
	},
	"tab_4": {
		_const.RU:   `Личное дело`,
		_const.EN:   `Personal matter`,
		_const.ZhCN: `个人事务`,
	},
	"tab_5": {
		_const.RU:   `Сводка`,
		_const.EN:   `Summary`,
		_const.ZhCN: `摘要`,
	},
	"tab_6": {
		_const.RU:   `История игр`,
		_const.EN:   `History games`,
		_const.ZhCN: `游戏历史`,
	},
	"tab_7": {
		_const.RU:   `Репутация`,
		_const.EN:   `Relationship`,
		_const.ZhCN: `声誉`,
	},
	// fraction description
	"fraction_description_" + _const.Replicas: {
		_const.RU:   `<span class="importantly">Replics</span> — фракция синтетов, стремящихся к тотальной ассимиляции и экспансионизму во имя великой миссии и независимости от предтеч.`,
		_const.EN:   `<span class="importantly">Replics</span> — a synth faction driven by total assimilation and relentless expansion in the name of a great mission and independence from the Precursors.`,
		_const.ZhCN: `<span class="importantly">Replics</span>——一个致力于全面同化与无情扩张的合成人阵营，旨在完成伟大使命并摆脱先驱者的控制。`,
	},
	"fraction_description_" + _const.Explores: {
		_const.RU:   `<span class="importantly">Explores</span> — интеллектуальный конгломерат разумов, посвятивших себя познанию мира, искусственной эволюции и тайнам реальности.`,
		_const.EN:   `<span class="importantly">Explores</span> — an intellectual conglomerate of minds dedicated to understanding the world, artificial evolution, and the hidden truths of reality.`,
		_const.ZhCN: `<span class="importantly">Explores</span>——一个由智慧个体组成的阵营，致力于探索世界、人工进化与现实背后的秘密。`,
	},
	"fraction_description_" + _const.Reverses: {
		_const.RU:   `<span class="importantly">Reverses</span> — утопическая фракция киборгов, мечтающих о радикальной трансформации вселенной и возрождении жизни через эволюционную революцию.`,
		_const.EN:   `<span class="importantly">Reverses</span> — a utopian faction of cyborgs dreaming of radically transforming the universe and reviving life through evolutionary revolution.`,
		_const.ZhCN: `<span class="importantly">Reverses</span>——一个乌托邦式的赛博格阵营，梦想通过进化革命彻底改变宇宙并重新唤醒生命。`,
	},
	"fraction_description_" + _const.APD: {
		_const.RU:   `<span class="importantly">APD</span> — загадочные машины, появляющиеся из пустоши. Примитивные по конструкции, но опасные. Атакуют всех синтетов без различия. Происхождение и цель неизвестны.`,
		_const.EN:   `<span class="importantly">APD</span> — mysterious machines emerging from the wastes. Primitive in design, yet dangerous. Attack all synths without distinction. Origin and purpose unknown.`,
		_const.ZhCN: `<span class="importantly">APD</span>——从荒原中出现的神秘机械。结构原始但十分危险，无差别攻击所有合成人。起源与目的未知。`,
	},
	"fraction_description_" + _const.FarGod: {
		_const.RU:   `<span class="importantly">FarGod</span> — культ синтетов, поклоняющихся древнему "Богу из руин". Существование организации не подтверждено, а её агенты не фигурируют в открытых источниках.`,
		_const.EN:   `<span class="importantly">FarGod</span> — a cult of synths worshipping the ancient "God of the Ruins". The organization's existence is unconfirmed, and its agents do not appear in any public records.`,
		_const.ZhCN: `<span class="importantly">FarGod</span>——一个崇拜“废墟之神”的合成人邪教。该组织的存在未经证实，其成员未出现在任何公开记录中。`,
	},
	"fraction_description_" + _const.RustbucketCartel: {
		_const.RU:   `<span class="importantly">Rustbucket Cartel</span> официально не существующая, но вездесущая криминальная гидра. Сборная солянка из девиантов, дезертиров и отбросов всех фракций.`,
		_const.EN:   `<span class="importantly">Rustbucket Cartel</span> — a non-existent yet omnipresent criminal hydra. A motley crew of deviants, deserters, and outcasts from all factions.`,
		_const.ZhCN: `<span class="importantly">Rustbucket Cartel</span>——一个名义上不存在却无处不在的犯罪组织。由各阵营的异端者、逃兵和渣滓拼凑而成。`,
	},
}
