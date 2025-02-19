package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Lobby = map[string]map[string]string{
	// BaseStatus
	// ExitMenu
	"window_name_1": {
		_const.RU: `Выход с базы`,
		_const.EN: `Leaving the base`,
	},
	"error_1": {
		_const.RU: `Трюм переполнен`,
		_const.EN: `The hold is full`,
	},
	"error_2": {
		_const.RU: `Не выбран транспорт`,
		_const.EN: `No transport selected`,
	},
	"text_1": {
		_const.RU: `Корпус:`,
		_const.EN: `Body:`,
	},
	"button_1": {
		_const.RU: `Починить`,
		_const.EN: `Repair`,
	},
	"text_2": {
		_const.RU: `Оружие:`,
		_const.EN: `Weapon:`,
	},
	"button_2": {
		_const.RU: `Зарядить`,
		_const.EN: `Arm`,
	},
	"text_3": {
		_const.RU: `Топливо:`,
		_const.EN: `Fuel:`,
	},
	"button_3": {
		_const.RU: `Заправить`,
		_const.EN: `Refuel`,
	},
	"button_4": {
		_const.RU: `< Отмена`,
		_const.EN: `< Cancel`,
	},
	"button_5": {
		_const.RU: `Покинуть базу >`,
		_const.EN: `Leave base >`,
	},
	"body_state_1": {
		_const.RU: `готов`,
		_const.EN: `ready`,
	},
	"body_state_2": {
		_const.RU: `поврежден`,
		_const.EN: `damaged`,
	},
	"body_state_3": {
		_const.RU: `крит. состояние`,
		_const.EN: `Critical condition`,
	},
	"body_state_4": {
		_const.RU: `не установлен`,
		_const.EN: `not installed`,
	},
	"text_13": {
		_const.RU: `Не хватает места в трюме`,
		_const.EN: `Not enough space in the hold`,
	},
	// Lobby
	"text_4": {
		_const.RU: `Сектор атакован!`,
		_const.EN: `Sector under attack!`,
	},
	"text_5": {
		_const.RU: `Фракция`,
		_const.EN: `Fraction`,
	},
	"text_6": {
		_const.RU: `Корпорация`,
		_const.EN: `Corporation`,
	},
	"menu_1": {
		_const.RU: `Выйти с базы`,
		_const.EN: `Exit base`,
	},
	"menu_2": {
		_const.RU: `Наблюдение`,
		_const.EN: `Observation`,
	},
	"menu_3": {
		_const.RU: `Мини игра`,
		_const.EN: `Mini game`,
	},
	"menu_4": {
		_const.RU: `Имущество`,
		_const.EN: `Property`,
	},
	"menu_5": {
		_const.RU: `Ангар`,
		_const.EN: `Hangar`,
	},
	"menu_6": {
		_const.RU: `Склад / Трюм`,
		_const.EN: `Warehouse / Hold`,
	},
	"menu_7": {
		_const.RU: `Рынок`,
		_const.EN: `Market`,
	},
	"menu_8": {
		_const.RU: `Торговля`,
		_const.EN: `Commerce`,
	},
	"menu_9": {
		_const.RU: `Награды`,
		_const.EN: `Awards`,
	},
	"menu_10": {
		_const.RU: `Производство`,
		_const.EN: `Production`,
	},
	"menu_11": {
		_const.RU: `Переработчик`,
		_const.EN: `Recycler`,
	},
	"menu_12": {
		_const.RU: `Верстак`,
		_const.EN: `Workbench`,
	},
	"menu_13": {
		_const.RU: `Детали`,
		_const.EN: `Details`,
	},
	"menu_14": {
		_const.RU: `Связь`,
		_const.EN: `Communications`,
	},
	"menu_15": {
		_const.RU: `Терминал`,
		_const.EN: `Terminal`,
	},
	"menu_16": {
		_const.RU: `Новости / Поиск`,
		_const.EN: `News / Search`,
	},
	"text_7": {
		_const.RU: `Ожидайте освобождение главных ворот`,
		_const.EN: `Expect the release of the main gate`,
	},
	"text_8": {
		_const.RU: `База в пустошах`,
		_const.EN: `A base in the wasteland`,
	},
	"text_9": {
		_const.RU: `Неизвестно`,
		_const.EN: `Unknown`,
	},
	// SessionBattleRank
	"text_10": {
		_const.RU: `Недавно вы получали урон или нарушили закон, поэтому ваш транспорт сможет выйти из игры только через время.`,
		_const.EN: `You have recently taken damage or broken the law, so your vehicle will be unable to leave the game for some time.`,
	},
	"text_11": {
		_const.RU: `Недавно вы нарушили закон.`,
		_const.EN: `You have recently broken the law.`,
	},
	"text_12": {
		_const.RU: `Вы можете поднять ранг ополчения.`,
		_const.EN: `You can raise your militia rank.`,
	},
	"text_14": {
		_const.RU: `База законсервирована!`,
		_const.EN: `The base is mothballed!`,
	},
}
