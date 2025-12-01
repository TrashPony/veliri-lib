package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Lobby = map[string]map[string]string{
	// BaseStatus
	// ExitMenu
	"window_name_1": {
		_const.RU:   `Выход с базы`,
		_const.EN:   `Leaving the base`,
		_const.ZhCN: `离开基地`,
	},
	"error_1": {
		_const.RU:   `Грузовой отсек переполнен`,
		_const.EN:   `The hold is full`,
		_const.ZhCN: `货舱已满`,
	},
	"error_2": {
		_const.RU:   `Не выбран транспорт`,
		_const.EN:   `No transport selected`,
		_const.ZhCN: `未选择运输工具`,
	},
	"text_1": {
		_const.RU:   `Корпус:`,
		_const.EN:   `Body:`,
		_const.ZhCN: `船体:`,
	},
	"button_1": {
		_const.RU:   `Починить`,
		_const.EN:   `Repair`,
		_const.ZhCN: `修理`,
	},
	"text_2": {
		_const.RU:   `Оружие:`,
		_const.EN:   `Weapon:`,
		_const.ZhCN: `武器:`,
	},
	"button_2": {
		_const.RU:   `Зарядить`,
		_const.EN:   `Arm`,
		_const.ZhCN: `装备`,
	},
	"text_3": {
		_const.RU:   `Топливо:`,
		_const.EN:   `Fuel:`,
		_const.ZhCN: `燃料:`,
	},
	"button_3": {
		_const.RU:   `Заправить`,
		_const.EN:   `Refuel`,
		_const.ZhCN: `加注燃料`,
	},
	"button_4": {
		_const.RU:   `Отмена`,
		_const.EN:   `Cancel`,
		_const.ZhCN: `取消`,
	},
	"button_5": {
		_const.RU:   `Покинуть базу`,
		_const.EN:   `Leave base`,
		_const.ZhCN: `离开基地`,
	},
	"body_state_1": {
		_const.RU:   `готов`,
		_const.EN:   `ready`,
		_const.ZhCN: `准备就绪`,
	},
	"body_state_2": {
		_const.RU:   `поврежден`,
		_const.EN:   `damaged`,
		_const.ZhCN: `受损`,
	},
	"body_state_3": {
		_const.RU:   `крит. состояние`,
		_const.EN:   `Critical condition`,
		_const.ZhCN: `严重损坏`,
	},
	"body_state_4": {
		_const.RU:   `не установлен`,
		_const.EN:   `not installed`,
		_const.ZhCN: `未安装`,
	},
	"text_13": {
		_const.RU:   `Не хватает места в грузовом отсеке!`,
		_const.EN:   `Cargo hold full!`,
		_const.ZhCN: `货舱已满！`,
	},
	// Lobby
	"text_4": {
		_const.RU:   `Сектор атакован!`,
		_const.EN:   `Sector under attack!`,
		_const.ZhCN: `区域遭受攻击！`,
	},
	"text_5": {
		_const.RU:   `Фракция`,
		_const.EN:   `Fraction`,
		_const.ZhCN: `派系`,
	},
	"text_6": {
		_const.RU:   `Корпорация`,
		_const.EN:   `Corporation`,
		_const.ZhCN: `公司`,
	},
	"menu_1": {
		_const.RU:   `Выйти с базы`,
		_const.EN:   `Exit base`,
		_const.ZhCN: `离开基地`,
	},
	"menu_2": {
		_const.RU:   `Наблюдение`,
		_const.EN:   `Observation`,
		_const.ZhCN: `观察`,
	},
	"menu_3": {
		_const.RU:   `Мини игра`,
		_const.EN:   `Mini game`,
		_const.ZhCN: `小游戏`,
	},
	"menu_4": {
		_const.RU:   `Имущество`,
		_const.EN:   `Property`,
		_const.ZhCN: `财产`,
	},
	"menu_5": {
		_const.RU:   `Ангар`,
		_const.EN:   `Hangar`,
		_const.ZhCN: `机库`,
	},
	"menu_6": {
		_const.RU:   `Склад / Грузовой отсек`,
		_const.EN:   `Storage / Cargo Hold`,
		_const.ZhCN: `仓储 / 货舱`,
	},
	"menu_7": {
		_const.RU:   `Рынок`,
		_const.EN:   `Market`,
		_const.ZhCN: `市场`,
	},
	"menu_8": {
		_const.RU:   `Торговля`,
		_const.EN:   `Commerce`,
		_const.ZhCN: `贸易`,
	},
	"menu_9": {
		_const.RU:   `Награды`,
		_const.EN:   `Awards`,
		_const.ZhCN: `奖励`,
	},
	"menu_10": {
		_const.RU:   `Производство`,
		_const.EN:   `Production`,
		_const.ZhCN: `生产`,
	},
	"menu_11": {
		_const.RU:   `Переработчик`,
		_const.EN:   `Recycler`,
		_const.ZhCN: `回收机`,
	},
	"menu_12": {
		_const.RU:   `Верстак`,
		_const.EN:   `Workbench`,
		_const.ZhCN: `工作台`,
	},
	"menu_13": {
		_const.RU:   `Производство`,
		_const.EN:   `Production`,
		_const.ZhCN: `生产`,
	},
	"menu_14": {
		_const.RU:   `Связь`,
		_const.EN:   `Communications`,
		_const.ZhCN: `通讯`,
	},
	"menu_15": {
		_const.RU:   `Терминал`,
		_const.EN:   `Terminal`,
		_const.ZhCN: `终端`,
	},
	"menu_16": {
		_const.RU:   `Новости / Поиск`,
		_const.EN:   `News / Search`,
		_const.ZhCN: `新闻 / 搜索`,
	},
	"menu_17": {
		_const.RU:   `Лаборатория`,
		_const.EN:   `Laboratory`,
		_const.ZhCN: `实验室`,
	},
	"text_7": {
		_const.RU:   `Ожидайте освобождение главных ворот`,
		_const.EN:   `Expect the release of the main gate`,
		_const.ZhCN: `等待主门开启`,
	},
	"text_8": {
		_const.RU:   `База в пустошах`,
		_const.EN:   `A base in the wasteland`,
		_const.ZhCN: `废土中的基地`,
	},
	"text_9": {
		_const.RU:   `Неизвестно`,
		_const.EN:   `Unknown`,
		_const.ZhCN: `未知`,
	},
	// SessionBattleRank
	"text_10": {
		_const.RU:   `Недавно вы получали урон или нарушили закон, поэтому ваш транспорт сможет выйти из игры только через время.`,
		_const.EN:   `You have recently taken damage or broken the law, so your vehicle will be unable to leave the game for some time.`,
		_const.ZhCN: `你最近受到了伤害或违反了法律，因此你的载具将暂时无法离开游戏。`,
	},
	"text_11": {
		_const.RU:   `Недавно вы нарушили закон.`,
		_const.EN:   `You have recently broken the law.`,
		_const.ZhCN: `你最近违反了法律。`,
	},
	"text_12": {
		_const.RU:   `Вы можете поднять ранг ополчения.`,
		_const.EN:   `You can raise your militia rank.`,
		_const.ZhCN: `你可以提升民兵等级。`,
	},
	"text_14": {
		_const.RU:   `База законсервирована!`,
		_const.EN:   `The base is mothballed!`,
		_const.ZhCN: `基地已封存！`,
	},
	"exit_expedition": {
		_const.RU:   `Покинуть объект`,
		_const.EN:   `Leave the facility`,
		_const.ZhCN: `离开设施`,
	},
}
