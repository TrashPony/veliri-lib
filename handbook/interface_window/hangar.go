package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Hangar = map[string]map[string]string{
	// Constructor
	"text_1": {
		_const.RU: `Не хватает мощности`,
		_const.EN: `Not enough power`,
	},
	"text_2": {
		_const.RU: `Не хватает места в трюме`,
		_const.EN: `Not enough space in the hold`,
	},
	"text_3": {
		_const.RU: `Покраски`,
		_const.EN: `Skins`,
	},
	"text_4": {
		_const.RU: `Корпус`,
		_const.EN: `Body`,
	},
	"text_5": {
		_const.RU: `Оружие`,
		_const.EN: `Weapon`,
	},
	"text_6": {
		_const.RU: `Хар-ки снижены`,
		_const.EN: `Characteristics are reduced`,
	},
	"text_7": {
		_const.RU: `Не хватает навыков для управления <span class="importantly">корпусом</span>.`,
		_const.EN: `There are not enough skills to manage the <span class="importantly">corps</span>.`,
	},
	"text_8": {
		_const.RU: `Не хватает навыков для управления <span class="importantly">орудием</span>.`,
		_const.EN: `There are not enough skills to operate the <span class="importantly">weapon</span>.`,
	},
	"text_9": {
		_const.RU: `Не хватает навыков для использования <span class="importantly">снарядов</span>.`,
		_const.EN: `There are not enough skills to use <span class="importantly">shells</span>.`,
	},
	"text_10": {
		_const.RU: `Транспорт не заправлен.`,
		_const.EN: `The transport is not fueled.`,
	},
	"text_11": {
		_const.RU: `Покраски:`,
		_const.EN: `Skins:`,
	},
	"button_1": {
		_const.RU: `Назад`,
		_const.EN: `Back`,
	},
	"button_2": {
		_const.RU: `Купить`,
		_const.EN: `Buy`,
	},
	// Hangar
	"window_name": {
		_const.RU: `Конструктор снаряжения`,
		_const.EN: `Equipment designer`,
	},
	"text_12": {
		_const.RU: `Перенести транспорт?`,
		_const.EN: `Move transport?`,
	},
	"text_13": {
		_const.RU: `<p>Перенести транспорт и все его содержимое на склад?</p><p><span class="importantly">Внимание!</span> Модификаторы установленные на транспорт будут уничтожены.</p>`,
		_const.EN: `<p>Move the transport and all its contents to the warehouse?</p><p><span class="importantly">Attention!</span> Modifiers installed on vehicles will be destroyed.</p>`,
	},
	"text_19": {
		_const.RU: `Добавление монтажного слота`,
		_const.EN: `Adding a Mounting Slot`,
	},
	"text_20": {
		_const.RU: `<p>Вы хотите установить дополнительную монтажную площадку <span class="importantly">%type%</span> уровня.</p><p>Это обойдется в <span class="importantly">%credits_count%</span> cr., а так же <span class="importantly">снизит прочность корпуса</span> на <span class="importantly">%debuf%%</span>.</p><p>Добавить?</p>`,
		_const.EN: `<p>You want to install an additional <span class="importantly">%type%</span> level mounting pad.</p><p>This will cost <span class="importantly">%credits_count%</span> cr., and will also <span class="importantly">reduce the strength of the hull</span> by <span class="importantly">%debuf%%</span>.</p><p>Add?</p>`,
	},
	"text_21": {
		_const.RU: `Нельзя устанавливать 2 одинаковых модификатора`,
		_const.EN: `You cannot install 2 identical modifiers`,
	},
	"text_22": {
		_const.RU: `Не удалось!`,
		_const.EN: `Failed!`,
	},
	"button_3": {
		_const.RU: `Отмена`,
		_const.EN: `Cancel`,
	},
	"button_4": {
		_const.RU: `Перенести`,
		_const.EN: `Transfer`,
	},
	"button_7": {
		_const.RU: `Добавить`,
		_const.EN: `Add`,
	},
	"text_14": {
		_const.RU: `Изменить имя`,
		_const.EN: `Change name`,
	},
	"text_15": {
		_const.RU: `Укажите новое имя транспорта.`,
		_const.EN: `Specify a new transport name.`,
	},
	"placeholder_1": {
		_const.RU: `новое имя`,
		_const.EN: `new name`,
	},
	"button_5": {
		_const.RU: `Отмена`,
		_const.EN: `Cancel`,
	},
	"button_6": {
		_const.RU: `Изменить`,
		_const.EN: `Change`,
	},
	// HangarHead
	"text_16": {
		_const.RU: `Транспорт:`,
		_const.EN: `Transport:`,
	},
	// MSParams
	"unitsec": {
		_const.RU: `ед/с`,
		_const.EN: `units/s`,
	},
	"km": {
		_const.RU: `км`,
		_const.EN: `km`,
	},
	"kms": {
		_const.RU: `км/с`,
		_const.EN: `km/s`,
	},
	"deg": {
		_const.RU: `гр/с`,
		_const.EN: `deg/s`,
	},
	"text_17": {
		_const.RU: `Транспорты на базе:`,
		_const.EN: `Transports at the base:`,
	},
	"ch_1": {
		_const.RU: `Атака`,
		_const.EN: `Attack`,
	},
	"ch_2": {
		_const.RU: `Урон`,
		_const.EN: `Damage`,
	},
	"ch_3": {
		_const.RU: `Тип атаки`,
		_const.EN: `Attack type`,
	},
	"ch_4": {
		_const.RU: `Очередь`,
		_const.EN: `Queue`,
	},
	"ch_5": {
		_const.RU: `Дальность стрельбы`,
		_const.EN: `Range attack`,
	},
	"ch_6": {
		_const.RU: `Точность`,
		_const.EN: `Accuracy`,
	},
	"ch_7": {
		_const.RU: `Скорость поворота орудия`,
		_const.EN: `Speed of rotation of the gun`,
	},
	"ch_8": {
		_const.RU: `Скорость полета снаряда`,
		_const.EN: `Projectile flight speed`,
	},
	"ch_9": {
		_const.RU: `Скорострельность`,
		_const.EN: `Rate of fire`,
	},
	"ch_10": {
		_const.RU: `Время перезарядки`,
		_const.EN: `Reload time`,
	},
	"text_18": {
		_const.RU: `Оружие не укомплектовано`,
		_const.EN: `Weapon is not equipped:`,
	},
	"ch_11": {
		_const.RU: `Защита`,
		_const.EN: `Protection`,
	},
	"ch_12": {
		_const.RU: `Структура`,
		_const.EN: `Structure`,
	},
	"ch_13": {
		_const.RU: `Защита`,
		_const.EN: `Protection`,
	},
	"ch_14": {
		_const.RU: `Навигация`,
		_const.EN: `Navigation`,
	},
	"ch_15": {
		_const.RU: `Дальность обзора`,
		_const.EN: `Range view`,
	},
	"ch_16": {
		_const.RU: `Дальность радара`,
		_const.EN: `Range radar`,
	},
	"ch_17": {
		_const.RU: `Накопитель`,
		_const.EN: `Battery`,
	},
	"ch_18": {
		_const.RU: `Макс. емкость`,
		_const.EN: `Max. capacity`,
	},
	"ch_19": {
		_const.RU: `Скорость зарядки`,
		_const.EN: `Charging speed`,
	},
	"ch_20": {
		_const.RU: `на месте:`,
		_const.EN: `on site:`,
	},
	"ch_21": {
		_const.RU: `Ходовая`,
		_const.EN: `Running gear`,
	},
	"ch_22": {
		_const.RU: `Макс. скорость`,
		_const.EN: `Max speed`,
	},
	"ch_23": {
		_const.RU: `Ускорение`,
		_const.EN: `Acceleration`,
	},
	"ch_24": {
		_const.RU: `Скорость поворота`,
		_const.EN: `Turn speed`,
	},
	"ch_25": {
		_const.RU: `Щит`,
		_const.EN: `Shield`,
	},
	"ch_26": {
		_const.RU: `Защита щита`,
		_const.EN: `Shield protection`,
	},
	"ch_27": {
		_const.RU: `Отсутствует`,
		_const.EN: `Missing`,
	},
	// StateTip
	"tittle_1": {
		_const.RU: `Навык`,
		_const.EN: `Skill`,
	},
	"tittle_2": {
		_const.RU: `Корпус`,
		_const.EN: `Body`,
	},
	"tittle_3": {
		_const.RU: `Снаряжение`,
		_const.EN: `Equipment`,
	},
	"tittle_4": {
		_const.RU: `Штраф`,
		_const.EN: `Fine`,
	},
	"tittle_5": {
		_const.RU: `Оружие`,
		_const.EN: `Weapon`,
	},
	"tittle_6": {
		_const.RU: `Топливо`,
		_const.EN: `Fuel`,
	},
	"tittle_7": {
		_const.RU: `Время с уток`,
		_const.EN: `Times of Day`,
	},
	"tittle_8": {
		_const.RU: `Синхронизация`,
		_const.EN: `Synchronization`,
	},
	"tittle_9": {
		_const.RU: `Дополнительный слот`,
		_const.EN: `Additional slot`,
	},
	// ViewMS
	// Synchronization
	"text_23": {
		_const.RU: `Синхронизация с транспортом`,
		_const.EN: `Synchronization with transport`,
	},
	"text_24": {
		_const.RU: `назад`,
		_const.EN: `back`,
	},
	"text_25": {
		_const.RU: `Настройки и обновление синхронизации доступно только на базе.`,
		_const.EN: `Synchronization settings and updates are available only on the base.`,
	},
	"text_26": {
		_const.RU: `Приоритет подсистем`,
		_const.EN: `Subsystem priority`,
	},
	"text_27": {
		_const.RU: `Особенности`,
		_const.EN: `Peculiarities`,
	},
}
