package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Hangar = map[string]map[string]string{
	// Constructor
	"text_1": {
		_const.RU:   `Не хватает мощности`,
		_const.EN:   `Not enough power`,
		_const.ZhCN: `电力不足`,
	},
	"text_2": {
		_const.RU:   `Не хватает места в грузовом отсеке`,
		_const.EN:   `Not enough space in cargo hold`,
		_const.ZhCN: `货舱空间不足`,
	},
	"text_3": {
		_const.RU:   `Покраски`,
		_const.EN:   `Skins`,
		_const.ZhCN: `涂装`,
	},
	"text_4": {
		_const.RU:   `Корпус`,
		_const.EN:   `Body`,
		_const.ZhCN: `机身`,
	},
	"text_5": {
		_const.RU:   `Оружие`,
		_const.EN:   `Weapon`,
		_const.ZhCN: `武器`,
	},
	"text_6": {
		_const.RU:   `Хар-ки снижены`,
		_const.EN:   `Characteristics are reduced`,
		_const.ZhCN: `特性降低`,
	},
	"text_7": {
		_const.RU:   `Не хватает навыков для управления <span class="importantly">корпусом</span>.`,
		_const.EN:   `There are not enough skills to manage the <span class="importantly">corps</span>.`,
		_const.ZhCN: `没有足够的技能来管理<span class="importantly">机身</span>。`,
	},
	"text_8": {
		_const.RU:   `Не хватает навыков для управления <span class="importantly">орудием</span>.`,
		_const.EN:   `There are not enough skills to operate the <span class="importantly">weapon</span>.`,
		_const.ZhCN: `没有足够的技能来操作<span class="importantly">武器</span>。`,
	},
	"text_9": {
		_const.RU:   `Не хватает навыков для использования <span class="importantly">снарядов</span>.`,
		_const.EN:   `There are not enough skills to use <span class="importantly">shells</span>.`,
		_const.ZhCN: `没有足够的技能来使用<span class="importantly">弹药</span>。`,
	},
	"text_10": {
		_const.RU:   `Транспорт не заправлен.`,
		_const.EN:   `The transport is not fueled.`,
		_const.ZhCN: `运输工具未加油。`,
	},
	"text_11": {
		_const.RU:   `Покраски:`,
		_const.EN:   `Skins:`,
		_const.ZhCN: `涂装：`,
	},
	"button_1": {
		_const.RU:   `Назад`,
		_const.EN:   `Back`,
		_const.ZhCN: `返回`,
	},
	"button_2": {
		_const.RU:   `Купить`,
		_const.EN:   `Buy`,
		_const.ZhCN: `购买`,
	},
	// Hangar
	"window_name": {
		_const.RU:   `Ангар`,
		_const.EN:   `Hangar`,
		_const.ZhCN: `机库`,
	},
	"text_12": {
		_const.RU:   `Перенести транспорт?`,
		_const.EN:   `Move transport?`,
		_const.ZhCN: `移动运输工具？`,
	},
	"text_13": {
		_const.RU:   `<p>Перенести транспорт и все его содержимое на склад?</p><p><span class="importantly">Внимание!</span> Модификаторы установленные на транспорт будут уничтожены.</p>`,
		_const.EN:   `<p>Move the transport and all its contents to the warehouse?</p><p><span class="importantly">Attention!</span> Modifiers installed on vehicles will be destroyed.</p>`,
		_const.ZhCN: `<p>将运输工具及其所有内容移动到仓库？</p><p><span class="importantly">注意！</span>安装在运输工具上的修改器将被销毁。</p>`,
	},
	"text_19": {
		_const.RU:   `Добавление монтажного слота`,
		_const.EN:   `Adding a Mounting Slot`,
		_const.ZhCN: `添加安装槽`,
	},
	"text_20": {
		_const.RU:   `<p>Вы хотите установить дополнительную монтажную площадку <span class="importantly">%type%</span> уровня.</p><p>Это обойдется в <span class="importantly">%credits_count%</span> , а так же <span class="importantly">снизит прочность корпуса</span> на <span class="importantly">%debuf%%</span>.</p><p>Добавить?</p>`,
		_const.EN:   `<p>You want to install an additional <span class="importantly">%type%</span> level mounting pad.</p><p>This will cost <span class="importantly">%credits_count%</span> , and will also <span class="importantly">reduce the strength of the hull</span> by <span class="importantly">%debuf%%</span>.</p><p>Add?</p>`,
		_const.ZhCN: `<p>您想安装一个额外的<span class="importantly">%type%</span>级安装平台。</p><p>这将花费<span class="importantly">%credits_count%</span> ，并且还会将<span class="importantly">机身强度</span>降低<span class="importantly">%debuf%%</span>。</p><p>添加吗？</p>`,
	},
	"text_21": {
		_const.RU:   `Нельзя устанавливать 2 одинаковых модификатора`,
		_const.EN:   `You cannot install 2 identical modifiers`,
		_const.ZhCN: `不能安装两个相同的修改器`,
	},
	"text_22": {
		_const.RU:   `Не удалось!`,
		_const.EN:   `Failed!`,
		_const.ZhCN: `失败！`,
	},
	"button_3": {
		_const.RU:   `Отмена`,
		_const.EN:   `Cancel`,
		_const.ZhCN: `取消`,
	},
	"button_4": {
		_const.RU:   `Перенести`,
		_const.EN:   `Transfer`,
		_const.ZhCN: `转移`,
	},
	"button_7": {
		_const.RU:   `Добавить`,
		_const.EN:   `Add`,
		_const.ZhCN: `添加`,
	},
	"text_14": {
		_const.RU:   `Изменить имя`,
		_const.EN:   `Change name`,
		_const.ZhCN: `更改名称`,
	},
	"text_15": {
		_const.RU:   `Укажите новое имя транспорта.`,
		_const.EN:   `Specify a new transport name.`,
		_const.ZhCN: `指定新的运输工具名称。`,
	},
	"placeholder_1": {
		_const.RU:   `новое имя`,
		_const.EN:   `new name`,
		_const.ZhCN: `新名称`,
	},
	"button_5": {
		_const.RU:   `Отмена`,
		_const.EN:   `Cancel`,
		_const.ZhCN: `取消`,
	},
	"button_6": {
		_const.RU:   `Изменить`,
		_const.EN:   `Change`,
		_const.ZhCN: `更改`,
	},
	// HangarHead
	"text_16": {
		_const.RU:   `Транспорт:`,
		_const.EN:   `Transport:`,
		_const.ZhCN: `运输工具：`,
	},
	// MSParams
	"unitsec": {
		_const.RU:   `ед/с`,
		_const.EN:   `units/s`,
		_const.ZhCN: `单位/秒`,
	},
	"km": {
		_const.RU:   `км`,
		_const.EN:   `km`,
		_const.ZhCN: `公里`,
	},
	"kms": {
		_const.RU:   `км/с`,
		_const.EN:   `km/s`,
		_const.ZhCN: `公里/秒`,
	},
	"deg": {
		_const.RU:   `гр/с`,
		_const.EN:   `deg/s`,
		_const.ZhCN: `度/秒`,
	},
	"text_17": {
		_const.RU:   `Транспорты на базе:`,
		_const.EN:   `Transports at the base:`,
		_const.ZhCN: `基地中的运输工具：`,
	},
	"ch_1": {
		_const.RU:   `Атака`,
		_const.EN:   `Attack`,
		_const.ZhCN: `攻击`,
	},
	"ch_2": {
		_const.RU:   `Урон`,
		_const.EN:   `Damage`,
		_const.ZhCN: `伤害`,
	},
	"ch_3": {
		_const.RU:   `Тип атаки`,
		_const.EN:   `Attack type`,
		_const.ZhCN: `攻击类型`,
	},
	"ch_4": {
		_const.RU:   `Очередь`,
		_const.EN:   `Queue`,
		_const.ZhCN: `队列`,
	},
	"ch_5": {
		_const.RU:   `Дальность стрельбы`,
		_const.EN:   `Range attack`,
		_const.ZhCN: `射程`,
	},
	"ch_6": {
		_const.RU:   `Точность`,
		_const.EN:   `Accuracy`,
		_const.ZhCN: `精度`,
	},
	"ch_7": {
		_const.RU:   `Скорость поворота орудия`,
		_const.EN:   `Speed of rotation of the gun`,
		_const.ZhCN: `炮塔旋转速度`,
	},
	"ch_8": {
		_const.RU:   `Скорость полета снаряда`,
		_const.EN:   `Projectile flight speed`,
		_const.ZhCN: `弹丸飞行速度`,
	},
	"ch_9": {
		_const.RU:   `Скорострельность`,
		_const.EN:   `Rate of fire`,
		_const.ZhCN: `射速`,
	},
	"ch_10": {
		_const.RU:   `Время перезарядки`,
		_const.EN:   `Reload time`,
		_const.ZhCN: `重新装填时间`,
	},
	"text_18": {
		_const.RU:   `Оружие не укомплектовано`,
		_const.EN:   `Weapon is not equipped:`,
		_const.ZhCN: `武器未装备：`,
	},
	"ch_11": {
		_const.RU:   `Защита`,
		_const.EN:   `Protection`,
		_const.ZhCN: `保护`,
	},
	"ch_12": {
		_const.RU:   `Структура`,
		_const.EN:   `Structure`,
		_const.ZhCN: `结构`,
	},
	"ch_13": {
		_const.RU:   `Защита`,
		_const.EN:   `Protection`,
		_const.ZhCN: `保护`,
	},
	"ch_14": {
		_const.RU:   `Навигация`,
		_const.EN:   `Navigation`,
		_const.ZhCN: `导航`,
	},
	"ch_15": {
		_const.RU:   `Дальность обзора`,
		_const.EN:   `Range view`,
		_const.ZhCN: `视野范围`,
	},
	"ch_16": {
		_const.RU:   `Дальность радара`,
		_const.EN:   `Range radar`,
		_const.ZhCN: `雷达范围`,
	},
	"ch_17": {
		_const.RU:   `Накопитель`,
		_const.EN:   `Battery`,
		_const.ZhCN: `电池`,
	},
	"ch_18": {
		_const.RU:   `Макс. емкость`,
		_const.EN:   `Max. capacity`,
		_const.ZhCN: `最大容量`,
	},
	"ch_19": {
		_const.RU:   `Скорость зарядки`,
		_const.EN:   `Charging speed`,
		_const.ZhCN: `充电速度`,
	},
	"ch_20": {
		_const.RU:   `на месте:`,
		_const.EN:   `on site:`,
		_const.ZhCN: `在现场：`,
	},
	"ch_21": {
		_const.RU:   `Ходовая`,
		_const.EN:   `Running gear`,
		_const.ZhCN: `行走机构`,
	},
	"ch_22": {
		_const.RU:   `Макс. скорость`,
		_const.EN:   `Max speed`,
		_const.ZhCN: `最大速度`,
	},
	"ch_23": {
		_const.RU:   `Ускорение`,
		_const.EN:   `Acceleration`,
		_const.ZhCN: `加速`,
	},
	"ch_24": {
		_const.RU:   `Скорость поворота`,
		_const.EN:   `Turn speed`,
		_const.ZhCN: `转弯速度`,
	},
	"ch_25": {
		_const.RU:   `Щит`,
		_const.EN:   `Shield`,
		_const.ZhCN: `护盾`,
	},
	"ch_26": {
		_const.RU:   `Защита щита`,
		_const.EN:   `Shield protection`,
		_const.ZhCN: `护盾保护`,
	},
	"ch_27": {
		_const.RU:   `Отсутствует`,
		_const.EN:   `Missing`,
		_const.ZhCN: `缺失`,
	},
	// StateTip
	"tittle_1": {
		_const.RU:   `Навык`,
		_const.EN:   `Skill`,
		_const.ZhCN: `技能`,
	},
	"tittle_2": {
		_const.RU:   `Корпус`,
		_const.EN:   `Body`,
		_const.ZhCN: `机身`,
	},
	"tittle_3": {
		_const.RU:   `Снаряжение`,
		_const.EN:   `Equipment`,
		_const.ZhCN: `装备`,
	},
	"tittle_4": {
		_const.RU:   `Штраф`,
		_const.EN:   `Fine`,
		_const.ZhCN: `罚款`,
	},
	"tittle_5": {
		_const.RU:   `Оружие`,
		_const.EN:   `Weapon`,
		_const.ZhCN: `武器`,
	},
	"tittle_6": {
		_const.RU:   `Топливо`,
		_const.EN:   `Fuel`,
		_const.ZhCN: `燃料`,
	},
	"tittle_7": {
		_const.RU:   `Время с уток`,
		_const.EN:   `Times of Day`,
		_const.ZhCN: `时间`,
	},
	"tittle_8": {
		_const.RU:   `Синхронизация`,
		_const.EN:   `Synchronization`,
		_const.ZhCN: `同步`,
	},
	"tittle_9": {
		_const.RU:   `Дополнительный слот`,
		_const.EN:   `Additional slot`,
		_const.ZhCN: `额外插槽`,
	},
	// ViewMS
	// Synchronization
	"text_23": {
		_const.RU:   `Синхронизация с транспортом`,
		_const.EN:   `Synchronization with transport`,
		_const.ZhCN: `与运输工具同步`,
	},
	"text_24": {
		_const.RU:   `назад`,
		_const.EN:   `back`,
		_const.ZhCN: `返回`,
	},
	"text_25": {
		_const.RU:   `Настройки и обновление синхронизации доступно только на базе.`,
		_const.EN:   `Synchronization settings and updates are available only on the base.`,
		_const.ZhCN: `同步设置和更新仅在基地可用。`,
	},
	"text_26": {
		_const.RU:   `Приоритет подсистем`,
		_const.EN:   `Subsystem priority`,
		_const.ZhCN: `子系统优先级`,
	},
	"text_27": {
		_const.RU:   `Особенности`,
		_const.EN:   `Peculiarities`,
		_const.ZhCN: `特性`,
	},
	// tips
	"tip_1": {
		_const.RU:   `Чем выше синхронизация, тем больше очков навыков доступно для распределения по системам`,
		_const.EN:   `Higher synchronization means more skill points available for system allocation`,
		_const.ZhCN: `同步率越高，可分配至各系统的技能点越多`,
	},
	"tip_2": {
		_const.RU:   `Просмотр схемы бронирования и слабых точек вашего транспорта`,
		_const.EN:   `View armor layout and weak points of your vehicle`,
		_const.ZhCN: `查看载具装甲布局与弱点分布`,
	},
	"tip_3": {
		_const.RU:   `Установленные модули`,
		_const.EN:   `Installed modules`,
		_const.ZhCN: `已安装模块`,
	},
	"tip_4": {
		_const.RU:   `Отправить на склад`,
		_const.EN:   `Send to storage`,
		_const.ZhCN: `发送至仓库`,
	},
	"tip_5": {
		_const.RU:   `Изменить название транспорта`,
		_const.EN:   `Rename vehicle`,
		_const.ZhCN: `重命名载具`,
	},
	"tip_6": {
		_const.RU:   `Шаблоны сборки`,
		_const.EN:   `Build templates`,
		_const.ZhCN: `组装模板`,
	},
	"tip_7": {
		_const.RU:   `Починка`,
		_const.EN:   `Repairs`,
		_const.ZhCN: `维修`,
	},
	"tip_8": {
		_const.RU:   `Доступная мощность для снаряжения`,
		_const.EN:   `Available power for equipment`,
		_const.ZhCN: `装备可用功率`,
	},
	"tip_9": {
		_const.RU:   `Слот для топлива`,
		_const.EN:   `Fuel slot`,
		_const.ZhCN: `燃料槽`,
	},
	"tip_10": {
		_const.RU:   `Слот питания L1`,
		_const.EN:   `L1 power slot`,
		_const.ZhCN: `L1级能源槽`,
	},
	"tip_11": {
		_const.RU:   `Установить дополнительный слот L1`,
		_const.EN:   `Install additional L1 slot`,
		_const.ZhCN: `安装额外L1级槽位`,
	},
	"tip_12": {
		_const.RU:   `Слот питания L2`,
		_const.EN:   `L2 power slot`,
		_const.ZhCN: `L2级能源槽`,
	},
	"tip_13": {
		_const.RU:   `Установить дополнительный слот L2`,
		_const.EN:   `Install additional L2 slot`,
		_const.ZhCN: `安装额外L2级槽位`,
	},
	"tip_14": {
		_const.RU:   `Слот питания L3`,
		_const.EN:   `L3 power slot`,
		_const.ZhCN: `L3级能源槽`,
	},
	"tip_15": {
		_const.RU:   `Слот модификатора`,
		_const.EN:   `Modifier slot`,
		_const.ZhCN: `改装件槽位`,
	},
	"tip_16": {
		_const.RU:   `Порт оперативного дрона`,
		_const.EN:   `Combat drone port`,
		_const.ZhCN: `作战无人机端口`,
	},
	"tip_17": {
		_const.RU:   `Установка модуля ограничивает соседний слот`,
		_const.EN:   `Module installation limits adjacent slot`,
		_const.ZhCN: `安装模块将限制相邻槽位`,
	},
	"tip_18": {
		_const.RU:   `Слот питания оружия`,
		_const.EN:   `Weapon power slot`,
		_const.ZhCN: `武器能源槽`,
	},
	"tip_19": {
		_const.RU:   `снаряды`,
		_const.EN:   `ammunition`,
		_const.ZhCN: `弹药`,
	},
}
