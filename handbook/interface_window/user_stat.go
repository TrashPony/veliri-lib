package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var UserStat = map[string]map[string]string{
	// Missions
	"no_missions": {
		_const.RU: `Нет заданий`,
		_const.EN: `No missions`,
	},
	"track": {
		_const.RU: `Отслеживать:`,
		_const.EN: `Track:`,
	},
	"log": {
		_const.RU: `Журнал`,
		_const.EN: `Journal`,
	},
	"tasks": {
		_const.RU: `Задачи`,
		_const.EN: `Tasks`,
	},
	"cancel": {
		_const.RU: `Отменить`,
		_const.EN: `Cancel`,
	},
	"confirm": {
		_const.RU: `Подтвердить`,
		_const.EN: `Confirm`,
	},
	// MyUser
	"avatar_change": {
		_const.RU: `Изменить`,
		_const.EN: `Change`,
	},
	"fraction_1": {
		_const.RU: `Фракция:`,
		_const.EN: `Fraction:`,
	},
	"rank_1": {
		_const.RU: `Ранг:`,
		_const.EN: `Rank:`,
	},
	"cluster_1": {
		_const.RU: `Кластер:`,
		_const.EN: `Cluster:`,
	},
	"no_cluster_1": {
		_const.RU: `не состоит в кластере`,
		_const.EN: `not a member of a corporation`,
	},
	"back_1": {
		_const.RU: `Назад`,
		_const.EN: `Back`,
	},
	"buy": {
		_const.RU: `Купить`,
		_const.EN: `Buy`,
	},
	// OtherUser
	"fraction_2": {
		_const.RU: `Фракция:`,
		_const.EN: `Fraction:`,
	},
	"cluster_2": {
		_const.RU: `Кластер:`,
		_const.EN: `Cluster:`,
	},
	"no_cluster_2": {
		_const.RU: `не состоит в кластере`,
		_const.EN: `not a member of a corporation`,
	},
	// Relations
	"enmity": {
		_const.RU: `Вражда`,
		_const.EN: `Enmity`,
	},
	"bad": {
		_const.RU: `Плохое`,
		_const.EN: `Bad`,
	},
	"neutral": {
		_const.RU: `Нейтральное`,
		_const.EN: `Neutral`,
	},
	"good": {
		_const.RU: `Хорошее`,
		_const.EN: `Good`,
	},
	"excellent": {
		_const.RU: `Отличное`,
		_const.EN: `Excellent`,
	},
	"chane_event_1": {
		_const.RU: `Торговля`,
		_const.EN: `Trade`,
	},
	"chane_event_2": {
		_const.RU: `Переработка ресурсов`,
		_const.EN: `Resource Recycling`,
	},
	"chane_event_3": {
		_const.RU: `Выполнение задания`,
		_const.EN: `Completing mission`,
	},
	"chane_event_4": {
		_const.RU: `Провал/отказ от задания`,
		_const.EN: `Failure/refusal of mission`,
	},
	"chane_event_5": {
		_const.RU: `Устранение пирата`,
		_const.EN: `Eliminating a pirate`,
	},
	"chane_event_6": {
		_const.RU: `Устранение мирного транспорта`,
		_const.EN: `Elimination of peaceful transport`,
	},
	"chane_event_7": {
		_const.RU: `Участие в захвате сектора`,
		_const.EN: `Participation in the capture of the sector`,
	},
	"head_1": {
		_const.RU: `Отношения с фракциями синтетов`,
		_const.EN: `Relations with Synthet factions`,
	},
	"button_1": {
		_const.RU: `история`,
		_const.EN: `log`,
	},
	"button_2": {
		_const.RU: `эффекты`,
		_const.EN: `effects`,
	},
	"head_2": {
		_const.RU: `Отношения с другими`,
		_const.EN: `Relationships with others`,
	},
	"head_3": {
		_const.RU: `Событие`,
		_const.EN: `Event`,
	},
	"head_4": {
		_const.RU: `Изменение`,
		_const.EN: `Change`,
	},
	"button_3": {
		_const.RU: `назад`,
		_const.EN: `back`,
	},
	"b_1": {
		_const.RU: `Обзор фракционных дронов`,
		_const.EN: `Faction drones view`,
	},
	"b_2": {
		_const.RU: `Снижение налога на переработку ресурсов`,
		_const.EN: `Reduced tax on resource processing`,
	},
	"b_3": {
		_const.RU: `Снижение налога на производство`,
		_const.EN: `Production tax reduction`,
	},
	"b_4": {
		_const.RU: `Снижение налога на производство деталей`,
		_const.EN: `Reduced tax on parts production`,
	},
	"b_5": {
		_const.RU: `Снижение налога на рынке`,
		_const.EN: `Market tax reduction`,
	},
	"b_6": {
		_const.RU: `Бонус кредитов за выполнение заданий`,
		_const.EN: `Bonus credits for completing tasks`,
	},
	"b_7": {
		_const.RU: `Запрет на торговлю`,
		_const.EN: `Trade ban`,
	},
	"b_8": {
		_const.RU: `Запрет на производство`,
		_const.EN: `Production ban`,
	},
	"b_9": {
		_const.RU: `Запрет на переработку ресурсов`,
		_const.EN: `Prohibition on resource processing`,
	},
	"b_10": {
		_const.RU: `Враг фракции`,
		_const.EN: `Fraction Enemy`,
	},
	// Skills
	"only_studied": {
		_const.RU: `Только изученные:`,
		_const.EN: `Only studied:`,
	},
	"experience": {
		_const.RU: `Опыт:`,
		_const.EN: `Experience:`,
	},
	"filter": {
		_const.RU: `Фильтр`,
		_const.EN: `Filter`,
	},
	"base_alert": {
		_const.RU: `Внимание! Тренировать навыки можно только на базе.`,
		_const.EN: `Attention! You can only train skills at the base.`,
	},
	"lvl": {
		_const.RU: `Уровень:`,
		_const.EN: `Level:`,
	},
	"price": {
		_const.RU: `Стоимость улучшения:`,
		_const.EN: `Upgrade cost:`,
	},
	"bonus_for_lvl": {
		_const.RU: `- бонус за уровень:`,
		_const.EN: `- level bonus:`,
	},
	"bonus_current": {
		_const.RU: `- текущий бонус:`,
		_const.EN: `- current bonus:`,
	},
	"missile": {
		_const.RU: `Ракетное`,
		_const.EN: `Missile`,
	},
	"mining": {
		_const.RU: `Добыча полезных ископаемых`,
		_const.EN: `Mining`,
	},
	"production": {
		_const.RU: `Производство`,
		_const.EN: `Production`,
	},
	"processing": {
		_const.RU: `Переработка полезных ископаемых`,
		_const.EN: `Processing`,
	},
	"bodies": {
		_const.RU: `Корпуса`,
		_const.EN: `Bodies`,
	},
	"weapon": {
		_const.RU: `Оружие`,
		_const.EN: `Weapon`,
	},
	"ammo": {
		_const.RU: `Боеприпасы`,
		_const.EN: `Ammo`,
	},
	"laser": {
		_const.RU: `Лазерное`,
		_const.EN: `Laser`,
	},
	"ballistic": {
		_const.RU: `Баллистическое`,
		_const.EN: `Ballistic`,
	},
	"small": {
		_const.RU: `малые`,
		_const.EN: `small`,
	},
	"medium": {
		_const.RU: `средние`,
		_const.EN: `medium`,
	},
	"big": {
		_const.RU: `большие`,
		_const.EN: `big`,
	},
	"light": {
		_const.RU: `легкие`,
		_const.EN: `light`,
	},
	"heavy": {
		_const.RU: `тяжелые`,
		_const.EN: `heavy`,
	},
	"other_1": {
		_const.RU: `остальные`,
		_const.EN: `other`,
	},
	"other_2": {
		_const.RU: `Остальные`,
		_const.EN: `Other`,
	},
	// UserStat
	"tab_1": {
		_const.RU: `Общие`,
		_const.EN: `Common`,
	},
	"tab_2": {
		_const.RU: `Навыки`,
		_const.EN: `Skills`,
	},
	"tab_3": {
		_const.RU: `Задания`,
		_const.EN: `Missions`,
	},
	"tab_4": {
		_const.RU: `Личное дело`,
		_const.EN: `Personal matter`,
	},
	"tab_5": {
		_const.RU: `Сводка`,
		_const.EN: `Summary`,
	},
	"tab_6": {
		_const.RU: `История игр`,
		_const.EN: `History games`,
	},
	"tab_7": {
		_const.RU: `Репутация`,
		_const.EN: `Relationship`,
	},
	// UserSummary
}
