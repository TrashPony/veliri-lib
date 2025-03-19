package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Corporation = map[string]map[string]string{
	// Corporation
	"window_name_1": {
		_const.RU:   `Кластер`,
		_const.EN:   `Cluster`,
		_const.ZhCN: `集群`,
	},
	"tab_1": {
		_const.RU:   `Поиск`,
		_const.EN:   `Search`,
		_const.ZhCN: `搜索`,
	},
	"tab_2": {
		_const.RU:   `О кластере`,
		_const.EN:   `About corporation`,
		_const.ZhCN: `关于集群`,
	},
	"tab_3": {
		_const.RU:   `Участники`,
		_const.EN:   `Members`,
		_const.ZhCN: `成员`,
	},
	"tab_4": {
		_const.RU:   `Роли`,
		_const.EN:   `Roles`,
		_const.ZhCN: `角色`,
	},
	"tab_5": {
		_const.RU:   `Заявки`,
		_const.EN:   `Applications`,
		_const.ZhCN: `申请`,
	},
	"tab_6": {
		_const.RU:   `Политика`,
		_const.EN:   `Politics`,
		_const.ZhCN: `政策`,
	},
	"tab_7": {
		_const.RU:   `Войны`,
		_const.EN:   `Wars`,
		_const.ZhCN: `战争`,
	},
	"tab_8": {
		_const.RU:   `Главная`,
		_const.EN:   `Main`,
		_const.ZhCN: `主页`,
	},
	"tab_9": {
		_const.RU:   `Поиск`,
		_const.EN:   `Search`,
		_const.ZhCN: `搜索`,
	},
	"tab_10": {
		_const.RU:   `Создать`,
		_const.EN:   `Create`,
		_const.ZhCN: `创建`,
	},
	"tab_11": {
		_const.RU:   `Аудит`,
		_const.EN:   `Audit`,
		_const.ZhCN: `审计`,
	},
	"text_1": {
		_const.RU:   `Вы не состоите в кластере.`,
		_const.EN:   `You are not in a cluster.`,
		_const.ZhCN: `您未加入任何集群`,
	},
	"button_1": {
		_const.RU:   `Поиск`,
		_const.EN:   `Search`,
		_const.ZhCN: `搜索`,
	},
	"button_2": {
		_const.RU:   `Создать`,
		_const.EN:   `Create`,
		_const.ZhCN: `创建`,
	},
	// CorporationLine
	"no_cluster": {
		_const.RU:   `Нет кластера`,
		_const.EN:   `No cluster`,
		_const.ZhCN: `无集群`,
	},
	// Create
	"helper_image": {
		_const.RU: `- jpg/png/gif <br>
						- не более 16КБ <br>
						- и 200х200px`,
		_const.EN: `- jpg/png/gif <br>
						- no more 16KB <br>
						- size 200х200px`,
		_const.ZhCN: `- jpg/png/gif格式 <br>
						- 不超过16KB <br>
						- 尺寸200х200像素`,
	},
	"text_2": {
		_const.RU:   `Создание кластера`,
		_const.EN:   `Creating a Cluster`,
		_const.ZhCN: `创建集群`,
	},
	"button_3": {
		_const.RU:   `Загрузить`,
		_const.EN:   `Upload`,
		_const.ZhCN: `上传`,
	},
	"text_3": {
		_const.RU:   `Название`,
		_const.EN:   `Name`,
		_const.ZhCN: `名称`,
	},
	"text_4": {
		_const.RU:   `(5-24 символов)`,
		_const.EN:   `(5-24 characters)`,
		_const.ZhCN: `(5-24个字符)`,
	},
	"text_5": {
		_const.RU:   `Тэг`,
		_const.EN:   `Tag`,
		_const.ZhCN: `标签`,
	},
	"text_6": {
		_const.RU:   `(2-5 символов)`,
		_const.EN:   `(2-5 characters)`,
		_const.ZhCN: `(2-5个字符)`,
	},
	"text_7": {
		_const.RU:   `Описание`,
		_const.EN:   `Description`,
		_const.ZhCN: `描述`,
	},
	"button_4": {
		_const.RU:   `Создать`,
		_const.EN:   `Create`,
		_const.ZhCN: `创建`,
	},
	"create_corporation_player_already": {
		_const.RU:   `Вы уже состоите в кластере`,
		_const.EN:   `You are already a member of the cluster`,
		_const.ZhCN: `您已加入集群`,
	},
	"create_corporation_empty_fields": {
		_const.RU:   `Не все поля заполнены`,
		_const.EN:   `Not all fields are filled in`,
		_const.ZhCN: `未填写所有字段`,
	},
	"create_corporation_name_wrong_size": {
		_const.RU:   `Неправильно кол-во символов в имени`,
		_const.EN:   `The number of characters in the name is incorrect`,
		_const.ZhCN: `名称字符数不正确`,
	},
	"create_corporation_tag_wrong_size": {
		_const.RU:   `Неправильно кол-во символов в теге`,
		_const.EN:   `Incorrect number of characters in tag`,
		_const.ZhCN: `标签字符数不正确`,
	},
	"create_corporation_biography_wrong_size": {
		_const.RU:   `Описание слишком большое`,
		_const.EN:   `Description too long`,
		_const.ZhCN: `描述过长`,
	},
	"create_corporation_name_already": {
		_const.RU:   `Имя кластера уже используется`,
		_const.EN:   `Cluster name is already in use`,
		_const.ZhCN: `集群名称已被使用`,
	},
	"create_corporation_tag_already": {
		_const.RU:   `Тег кластера уже используется`,
		_const.EN:   `The cluster tag is already in use`,
		_const.ZhCN: `集群标签已被使用`,
	},
	"create_corporation_logo_wrong_format": {
		_const.RU:   `Формат файла логотипа не верный`,
		_const.EN:   `Logo file format is not correct`,
		_const.ZhCN: `徽标文件格式不正确`,
	},
	"create_corporation_big_logo": {
		_const.RU:   `Размер логотипа больше 16КБ`,
		_const.EN:   `Logo size is larger than 16KB`,
		_const.ZhCN: `徽标大小超过16KB`,
	},
	"create_corporation_logo_wrong_resolution": {
		_const.RU:   `Размер лого больше 200х200 пикселей`,
		_const.EN:   `Logo size is more than 200x200 pixels`,
		_const.ZhCN: `徽标尺寸超过200x200像素`,
	},
	// Main
	"error_1": {
		_const.RU:   `Не удалось`,
		_const.EN:   `Failed`,
		_const.ZhCN: `失败`,
	},
	"rental_dialog_1": {
		_const.RU:   `Арендовать офис на <span class='importantly'>30</span> дней за <span class='importantly'>%credits_count% cr</span> c главного счета?`,
		_const.EN:   `Rent an office for <span class='importantly'>30</span> days for <span class='importantly'>%credits_count% cr</span> from the main deposit?`,
		_const.ZhCN: `从主账户支付<span class='importantly'>%credits_count% cr</span>租用办公室<span class='importantly'>30</span>天？`,
	},
	"rental_dialog_2": {
		_const.RU:   `Продлить аренду офиса на <span class='importantly'>30</span> дней за <span class='importantly'>%credits_count% cr</span> c главного счета?`,
		_const.EN:   `Extend office rental for <span class='importantly'>30</span> days for <span class='importantly'>%credits_count% cr</span> from the main deposit?`,
		_const.ZhCN: `从主账户支付<span class='importantly'>%credits_count% cr</span>延长办公室租期<span class='importantly'>30</span>天？`,
	},
	"button_5": {
		_const.RU:   `Закрыть`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
	"button_6": {
		_const.RU:   `Арендовать`,
		_const.EN:   `Rent`,
		_const.ZhCN: `租用`,
	},
	"button_7": {
		_const.RU:   `Продлить`,
		_const.EN:   `Extend`,
		_const.ZhCN: `延长`,
	},
	"text_8": {
		_const.RU:   `Заправить текущую базу на еще 30 дней?`,
		_const.EN:   `Refill the current database for another 30 days?`,
		_const.ZhCN: `为当前基地补充30天燃料？`,
	},
	"fill_up_dialog_1": {
		_const.RU:   `Цена: <span class='importantly'> %credits_count% </span> cr. и <span class='importantly'> %thorium_count% </span> единиц тория c вашего склада.`,
		_const.EN:   `Price: <span class='importantly'> %credits_count% </span> cr. and <span class='importantly'> %thorium_count% </span> thorium units from your warehouse.`,
		_const.ZhCN: `价格：<span class='importantly'>%credits_count%</span>信用点及仓库中的<span class='importantly'>%thorium_count%</span>单位钍。`,
	},
	"button_8": {
		_const.RU:   `Закрыть`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
	"button_9": {
		_const.RU:   `Заправить`,
		_const.EN:   `Refuel`,
		_const.ZhCN: `补充`,
	},
	"transfer_dialog_1": {
		_const.RU:   `Переместиться на базу <span class='importantly'>%base_name%</span>?`,
		_const.EN:   `Move to base <span class='importantly'>%base_name%</span>?`,
		_const.ZhCN: `转移到基地<span class='importantly'>%base_name%</span>？`,
	},
	"transfer_dialog_2": {
		_const.RU:   `Цена: <span class='importantly'>500.0</span> cr.`,
		_const.EN:   `Price: <span class='importantly'>500.0</span> cr.`,
		_const.ZhCN: `价格：<span class='importantly'>500.0</span>信用点`,
	},
	"transfer_dialog_3": {
		_const.RU:   `<span class='importantly'>Внимание!</span> Переносится только сознание, транспорт и все его содержимое останется на этой базе.`,
		_const.EN:   `<span class='importantly'>Attention!</span> Only consciousness, transport and all its contents are transferred will remain at this base.`,
		_const.ZhCN: `<span class='importantly'>注意！</span>仅意识转移，载具及内容物将留在原基地。`,
	},
	"button_10": {
		_const.RU:   `Закрыть`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
	"button_11": {
		_const.RU:   `Перемещение`,
		_const.EN:   `Moving`,
		_const.ZhCN: `移动`,
	},
	"deposit_history_h1": {
		_const.RU:   `Время`,
		_const.EN:   `Time`,
		_const.ZhCN: `时间`,
	},
	"deposit_history_h2": {
		_const.RU:   `Операция`,
		_const.EN:   `Operation`,
		_const.ZhCN: `操作`,
	},
	"deposit_history_h3": {
		_const.RU:   `Комментарий`,
		_const.EN:   `Comment`,
		_const.ZhCN: `备注`,
	},
	"button_12": {
		_const.RU:   `Закрыть`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
	"text_9": {
		_const.RU:   `Счета:`,
		_const.EN:   `Deposits:`,
		_const.ZhCN: `账户：`,
	},
	"button_13": {
		_const.RU:   `Перевод`,
		_const.EN:   `Translation`,
		_const.ZhCN: `转账`,
	},
	"button_14": {
		_const.RU:   `История`,
		_const.EN:   `Story`,
		_const.ZhCN: `历史`,
	},
	"text_10": {
		_const.RU:   `Офисы:`,
		_const.EN:   `Offices:`,
		_const.ZhCN: `办公室：`,
	},
	"text_11": {
		_const.RU:   `База:`,
		_const.EN:   `Base:`,
		_const.ZhCN: `基地：`,
	},
	"text_12": {
		_const.RU:   `До:`,
		_const.EN:   `Until:`,
		_const.ZhCN: `截至：`,
	},
	"button_15": {
		_const.RU:   `Маршрут`,
		_const.EN:   `Route`,
		_const.ZhCN: `路线`,
	},
	"button_16": {
		_const.RU:   `Автопилот`,
		_const.EN:   `Autopilot`,
		_const.ZhCN: `自动驾驶`,
	},
	"button_17": {
		_const.RU:   `Переместиться`,
		_const.EN:   `Move`,
		_const.ZhCN: `移动`,
	},
	"button_18": {
		_const.RU:   `Продлить аренду`,
		_const.EN:   `Extend your lease`,
		_const.ZhCN: `延长租期`,
	},
	"text_13": {
		_const.RU:   `Арендовать офис на базе`,
		_const.EN:   `Rent an office at the base`,
		_const.ZhCN: `在基地租用办公室`,
	},
	"text_14": {
		_const.RU:   `Базы:`,
		_const.EN:   `Bases:`,
		_const.ZhCN: `基地：`,
	},
	"text_15": {
		_const.RU:   `Сектор:`,
		_const.EN:   `Sector:`,
		_const.ZhCN: `区域：`,
	},
	"text_16": {
		_const.RU:   `Строительство до:`,
		_const.EN:   `Construction until:`,
		_const.ZhCN: `建造截止：`,
	},
	"text_17": {
		_const.RU:   `Неуязвима`,
		_const.EN:   `Invulnerable`,
		_const.ZhCN: `无敌`,
	},
	"text_18": {
		_const.RU:   `Уязвима`,
		_const.EN:   `Vulnerable`,
		_const.ZhCN: `脆弱`,
	},
	"text_19": {
		_const.RU:   `Под атакой`,
		_const.EN:   `Under attack`,
		_const.ZhCN: `遭受攻击`,
	},
	"text_20": {
		_const.RU:   `Топлива хватит еще до:`,
		_const.EN:   `There will be enough fuel until:`,
		_const.ZhCN: `燃料充足至：`,
	},
	"text_21": {
		_const.RU:   `Офис`,
		_const.EN:   `Office`,
		_const.ZhCN: `办公室`,
	},
	"text_22": {
		_const.RU:   `Завод`,
		_const.EN:   `Factory`,
		_const.ZhCN: `工厂`,
	},
	"text_23": {
		_const.RU:   `Переработчик`,
		_const.EN:   `Recycler`,
		_const.ZhCN: `回收站`,
	},
	"text_24": {
		_const.RU:   `Рынок`,
		_const.EN:   `Market`,
		_const.ZhCN: `市场`,
	},
	"button_19": {
		_const.RU:   `Маршрут`,
		_const.EN:   `Route`,
		_const.ZhCN: `路线`,
	},
	"button_20": {
		_const.RU:   `Автопилот`,
		_const.EN:   `Autopilot`,
		_const.ZhCN: `自动驾驶`,
	},
	"button_21": {
		_const.RU:   `Заправить`,
		_const.EN:   `Refuel`,
		_const.ZhCN: `加油`,
	},
	"text_25": {
		_const.RU:   `Меню строительства`,
		_const.EN:   `Construction menu`,
		_const.ZhCN: `建造菜单`,
	},
	"text_26": {
		_const.RU:   `Недоступно в текущем секторе и на базе`,
		_const.EN:   `Not available in current sector and base`,
		_const.ZhCN: `在当前区域和基地不可用`,
	},
	"text_27": {
		_const.RU:   `Участники в сети:`,
		_const.EN:   `Members on the network:`,
		_const.ZhCN: `在线成员：`,
	},
	"button_22": {
		_const.RU:   `Назад`,
		_const.EN:   `Back`,
		_const.ZhCN: `返回`,
	},
	"button_23": {
		_const.RU:   `Управление ролями`,
		_const.EN:   `Role management`,
		_const.ZhCN: `角色管理`,
	},
	"button_24": {
		_const.RU:   `Сохранить`,
		_const.EN:   `Save`,
		_const.ZhCN: `保存`,
	},
	"placeholder_1": {
		_const.RU:   `Название кластера`,
		_const.EN:   `Cluster name`,
		_const.ZhCN: `集群名称`,
	},
	"button_25": {
		_const.RU:   `Найти`,
		_const.EN:   `Find`,
		_const.ZhCN: `查找`,
	},
	"text_28": {
		_const.RU:   `Ничего не найдено.`,
		_const.EN:   `Nothing found.`,
		_const.ZhCN: `未找到任何内容。`,
	},
	"button_26": {
		_const.RU:   `Закрыть`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
	"button_27": {
		_const.RU:   `Добавить`,
		_const.EN:   `Add`,
		_const.ZhCN: `添加`,
	},
	"text_29": {
		_const.RU:   `Глобальная`,
		_const.EN:   `Global`,
		_const.ZhCN: `全局`,
	},
	"text_30": {
		_const.RU:   `Локальная`,
		_const.EN:   `Local`,
		_const.ZhCN: `本地`,
	},
	"text_31": {
		_const.RU:   `Союзники`,
		_const.EN:   `Allies`,
		_const.ZhCN: `盟友`,
	},
	"button_28": {
		_const.RU:   `Добавить`,
		_const.EN:   `Add`,
		_const.ZhCN: `添加`,
	},
	"text_32": {
		_const.RU:   `Соперники`,
		_const.EN:   `Rivals`,
		_const.ZhCN: `对手`,
	},
	"button_29": {
		_const.RU:   `Добавить`,
		_const.EN:   `Add`,
		_const.ZhCN: `添加`,
	},
	"text_33": {
		_const.RU:   `Сектор ...`,
		_const.EN:   `Sector ...`,
		_const.ZhCN: `区域...`,
	},
	"text_34": {
		_const.RU:   `Доступ в сектор:`,
		_const.EN:   `Access to the sector:`,
		_const.ZhCN: `区域访问：`,
	},
	"option_1": {
		_const.RU:   `Все`,
		_const.EN:   `All`,
		_const.ZhCN: `所有人`,
	},
	"option_2": {
		_const.RU:   `Все кроме врагов`,
		_const.EN:   `Everything except enemies`,
		_const.ZhCN: `除敌人外所有人`,
	},
	"option_3": {
		_const.RU:   `Все кроме союзников`,
		_const.EN:   `Everyone except allies`,
		_const.ZhCN: `除盟友外所有人`,
	},
	"option_4": {
		_const.RU:   `Все кроме соперников`,
		_const.EN:   `Everyone except the rivals`,
		_const.ZhCN: `除对手外所有人`,
	},
	"option_5": {
		_const.RU:   `Кластер и союзники`,
		_const.EN:   `Cluster and allies`,
		_const.ZhCN: `集群和盟友`,
	},
	"option_6": {
		_const.RU:   `Кластер и соперники`,
		_const.EN:   `Cluster and rivals`,
		_const.ZhCN: `集群和对手`,
	},
	"option_7": {
		_const.RU:   `Только Кластер`,
		_const.EN:   `Cluster only`,
		_const.ZhCN: `仅集群`,
	},
	"text_35": {
		_const.RU:   `Доступ на базу:`,
		_const.EN:   `Access to the base:`,
		_const.ZhCN: `基地访问：`,
	},
	"option_8": {
		_const.RU:   `Все`,
		_const.EN:   `All`,
		_const.ZhCN: `所有人`,
	},
	"option_9": {
		_const.RU:   `Все кроме врагов`,
		_const.EN:   `Everything except enemies`,
		_const.ZhCN: `除敌人外所有人`,
	},
	"option_10": {
		_const.RU:   `Все кроме союзников`,
		_const.EN:   `Everyone except allies`,
		_const.ZhCN: `除盟友外所有人`,
	},
	"option_11": {
		_const.RU:   `Все кроме соперников`,
		_const.EN:   `Everyone except the rivals`,
		_const.ZhCN: `除对手外所有人`,
	},
	"option_12": {
		_const.RU:   `Кластер и союзники`,
		_const.EN:   `Cluster and allies`,
		_const.ZhCN: `集群和盟友`,
	},
	"option_13": {
		_const.RU:   `Кластер и соперники`,
		_const.EN:   `Cluster and rivals`,
		_const.ZhCN: `集群和对手`,
	},
	"option_14": {
		_const.RU:   `Только Кластер`,
		_const.EN:   `Cluster only`,
		_const.ZhCN: `仅集群`,
	},
	"text_36": {
		_const.RU:   `Турели атакуют:`,
		_const.EN:   `Turrets attack:`,
		_const.ZhCN: `炮塔攻击：`,
	},
	"option_15": {
		_const.RU:   `Всех`,
		_const.EN:   `Everyone`,
		_const.ZhCN: `所有人`,
	},
	"option_16": {
		_const.RU:   `Всех кроме кластера`,
		_const.EN:   `Everyone except the cluster`,
		_const.ZhCN: `除集群外所有人`,
	},
	"option_17": {
		_const.RU:   `Всех кроме кластера и союзников`,
		_const.EN:   `Everyone except the cluster and allies`,
		_const.ZhCN: `除集群和盟友外所有人`,
	},
	"option_18": {
		_const.RU:   `Всех кроме кластера и соперников`,
		_const.EN:   `Everyone except the cluster and rivals`,
		_const.ZhCN: `除集群和对手外所有人`,
	},
	"option_19": {
		_const.RU:   `Нейтралов и врагов`,
		_const.EN:   `Neutrals and enemies`,
		_const.ZhCN: `中立者和敌人`,
	},
	"option_20": {
		_const.RU:   `Врагов (пк, апд и врагов в войне)`,
		_const.EN:   `Enemies (PK, APD and enemies in war)`,
		_const.ZhCN: `敌人（PK、APD和战争中的敌人）`,
	},
	"option_21": {
		_const.RU:   `Отключены`,
		_const.EN:   `Disabled`,
		_const.ZhCN: `禁用`,
	},
	"text_37": {
		_const.RU:   `Ремонтники чинят:`,
		_const.EN:   `Repairmen repair:`,
		_const.ZhCN: `维修人员修理：`,
	},
	"option_22": {
		_const.RU:   `Кластер`,
		_const.EN:   `Cluster`,
		_const.ZhCN: `集群`,
	},
	"option_23": {
		_const.RU:   `Кластер и союзников`,
		_const.EN:   `Cluster and allies`,
		_const.ZhCN: `集群和盟友`,
	},
	"option_24": {
		_const.RU:   `Кластер и соперников`,
		_const.EN:   `Cluster and rivals`,
		_const.ZhCN: `集群和对手`,
	},
	"option_25": {
		_const.RU:   `Всех кроме врагов`,
		_const.EN:   `Everyone except enemies`,
		_const.ZhCN: `除敌人外所有人`,
	},
	"option_26": {
		_const.RU:   `Отключены`,
		_const.EN:   `Disabled`,
		_const.ZhCN: `禁用`,
	},
	"text_38": {
		_const.RU:   `Налог на переработку:`,
		_const.EN:   `Processing tax:`,
		_const.ZhCN: `加工税：`,
	},
	"text_39": {
		_const.RU:   `Налог на производство:`,
		_const.EN:   `Production tax:`,
		_const.ZhCN: `生产税：`,
	},
	"text_40": {
		_const.RU:   `Налог на рынок:`,
		_const.EN:   `Market tax:`,
		_const.ZhCN: `市场税：`,
	},
	"text_41": {
		_const.RU:   `Склад для сбора налога:`,
		_const.EN:   `Tax collection warehouse:`,
		_const.ZhCN: `税收仓库：`,
	},
	"option_27": {
		_const.RU:   `Склад лидера`,
		_const.EN:   `Leader warehouse`,
		_const.ZhCN: `领袖仓库`,
	},
	"text_42": {
		_const.RU:   `Счет для сбора налога:`,
		_const.EN:   `Tax collection deposit:`,
		_const.ZhCN: `税收账户：`,
	},
	"option_28": {
		_const.RU:   `Счет лидера`,
		_const.EN:   `Leader deposit`,
		_const.ZhCN: `领袖账户`,
	},
	"text_43": {
		_const.RU:   `Окно уязвимости`,
		_const.EN:   `Vulnerability window`,
		_const.ZhCN: `漏洞窗口`,
	},
	"text_44": {
		_const.RU:   `Прием заявок:`,
		_const.EN:   `Accepting applications:`,
		_const.ZhCN: `接受申请：`,
	},
	"text_45": {
		_const.RU:   `Нет прав на прием игроков`,
		_const.EN:   `No rights to accept players`,
		_const.ZhCN: `无权接受玩家`,
	},
	"button_30": {
		_const.RU:   `Принять`,
		_const.EN:   `Accept`,
		_const.ZhCN: `接受`,
	},
	"button_31": {
		_const.RU:   `Отказать`,
		_const.EN:   `Reject`,
		_const.ZhCN: `拒绝`,
	},
	"text_46": {
		_const.RU:   `Роль:`,
		_const.EN:   `Role:`,
		_const.ZhCN: `角色：`,
	},
	"option_29": {
		_const.RU:   `Выберите роль`,
		_const.EN:   `Select a role`,
		_const.ZhCN: `选择角色`,
	},
	"text_47": {
		_const.RU:   `Общие:`,
		_const.EN:   `Common:`,
		_const.ZhCN: `通用：`,
	},
	"text_48": {
		_const.RU:   `Приглашать игроков`,
		_const.EN:   `Invite players`,
		_const.ZhCN: `邀请玩家`,
	},
	"text_49": {
		_const.RU:   `Выгонять игроков`,
		_const.EN:   `Kick players`,
		_const.ZhCN: `踢出玩家`,
	},
	"text_50": {
		_const.RU:   `Редактировать роли`,
		_const.EN:   `Edit roles`,
		_const.ZhCN: `编辑角色`,
	},
	"text_51": {
		_const.RU:   `Менять роли игрокам`,
		_const.EN:   `Change roles for players`,
		_const.ZhCN: `更改玩家角色`,
	},
	"text_52": {
		_const.RU:   `Счета:`,
		_const.EN:   `Deposits:`,
		_const.ZhCN: `账户：`,
	},
	"text_53": {
		_const.RU:   `Смотреть`,
		_const.EN:   `Look`,
		_const.ZhCN: `查看`,
	},
	"text_54": {
		_const.RU:   `Снимать`,
		_const.EN:   `Take off`,
		_const.ZhCN: `取出`,
	},
	"text_55": {
		_const.RU:   `Локальные:`,
		_const.EN:   `Local:`,
		_const.ZhCN: `本地：`,
	},
	"text_56": {
		_const.RU:   `Выберите офис`,
		_const.EN:   `Select an office`,
		_const.ZhCN: `选择办公室`,
	},
	"text_57": {
		_const.RU:   `Смотреть`,
		_const.EN:   `Look`,
		_const.ZhCN: `查看`,
	},
	"text_58": {
		_const.RU:   `Класть`,
		_const.EN:   `Putting`,
		_const.ZhCN: `放入`,
	},
	"text_59": {
		_const.RU:   `Забирать`,
		_const.EN:   `Take`,
		_const.ZhCN: `取出`,
	},
	"text_60": {
		_const.RU:   `Добавить роль:`,
		_const.EN:   `Add role:`,
		_const.ZhCN: `添加角色：`,
	},
	"placeholder_2": {
		_const.RU:   `Название роли`,
		_const.EN:   `Role name`,
		_const.ZhCN: `角色名称`,
	},
	"button_32": {
		_const.RU:   `Добавить`,
		_const.EN:   `Add`,
		_const.ZhCN: `添加`,
	},
	"build": {
		_const.RU:   `Строительство`,
		_const.EN:   `Building`,
		_const.ZhCN: `建造`,
	},
	"dismantling": {
		_const.RU:   `Демонтаж строений`,
		_const.EN:   `Dismantling of buildings`,
		_const.ZhCN: `拆除建筑`,
	},
	"change_local_policy": {
		_const.RU:   `Изменение локальной политики`,
		_const.EN:   `Change local policy`,
		_const.ZhCN: `更改本地政策`,
	},
	"text_61": {
		_const.RU:   `Фильтры:`,
		_const.EN:   `Filters:`,
		_const.ZhCN: `过滤器：`,
	},
	"placeholder_3": {
		_const.RU:   `Название кластера`,
		_const.EN:   `Cluster Name`,
		_const.ZhCN: `集群名称`,
	},
	"text_62": {
		_const.RU:   `моя заявка:`,
		_const.EN:   `my request:`,
		_const.ZhCN: `我的申请：`,
	},
	"button_33": {
		_const.RU:   `отозвать`,
		_const.EN:   `recall`,
		_const.ZhCN: `撤回`,
	},
	"button_34": {
		_const.RU:   `изменить`,
		_const.EN:   `change`,
		_const.ZhCN: `修改`,
	},
	"button_35": {
		_const.RU:   `Подать заявку`,
		_const.EN:   `Send request`,
		_const.ZhCN: `提交申请`,
	},
	"placeholder_4": {
		_const.RU:   `Текст заявки`,
		_const.EN:   `Request text`,
		_const.ZhCN: `申请文本`,
	},
	"button_36": {
		_const.RU:   `Отправить`,
		_const.EN:   `Send`,
		_const.ZhCN: `发送`,
	},
	"button_37": {
		_const.RU:   `Отмена`,
		_const.EN:   `Cancel`,
		_const.ZhCN: `取消`,
	},
	"button_50": {
		_const.RU:   `Поиск`,
		_const.EN:   `Search`,
		_const.ZhCN: `搜索`,
	},
	"option_32": {
		_const.RU:   `Все`,
		_const.EN:   `All`,
		_const.ZhCN: `所有`,
	},
	"option_33": {
		_const.RU:   `Открыт`,
		_const.EN:   `Open`,
		_const.ZhCN: `打开`,
	},
	"option_34": {
		_const.RU:   `Закрыт`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
	"text_63": {
		_const.RU:   `Успех!`,
		_const.EN:   `Success!`,
		_const.ZhCN: `成功！`,
	},
	"text_64": {
		_const.RU:   `База уже является штабом для другой войны`,
		_const.EN:   `The base is already the headquarters for another war`,
		_const.ZhCN: `该基地已经是另一场战争的总部`,
	},
	"text_65": {
		_const.RU:   `Такой базы не существует`,
		_const.EN:   `No such database exists`,
		_const.ZhCN: `该基地不存在`,
	},
	"text_66": {
		_const.RU:   `С этим кластером уже ведется война`,
		_const.EN:   `There is already a war going on with this cluster`,
		_const.ZhCN: `已经与该集群处于战争状态`,
	},
	"text_67": {
		_const.RU:   `У цели нет баз, вы не можете объявить ему войну`,
		_const.EN:   `The target has no bases, you cannot declare war on him`,
		_const.ZhCN: `目标没有基地，您无法向其宣战`,
	},
	"text_68": {
		_const.RU:   `Вы пока не можете объявить войну этому кластеру`,
		_const.EN:   `You cannot declare war on this cluster yet`,
		_const.ZhCN: `您暂时无法向该集群宣战`,
	},
	"placeholder_5": {
		_const.RU:   `Название кластера`,
		_const.EN:   `Cluster name`,
		_const.ZhCN: `集群名称`,
	},
	"button_38": {
		_const.RU:   `Найти`,
		_const.EN:   `Find`,
		_const.ZhCN: `查找`,
	},
	"text_69": {
		_const.RU:   `Ничего не найдено.`,
		_const.EN:   `Nothing found.`,
		_const.ZhCN: `未找到任何内容。`,
	},
	"text_70": {
		_const.RU:   `Спонсировать с`,
		_const.EN:   `Sponsor with`,
		_const.ZhCN: `赞助`,
	},
	"option_30": {
		_const.RU:   `Не выбрано`,
		_const.EN:   `Not chosen`,
		_const.ZhCN: `未选择`,
	},
	"text_71": {
		_const.RU: `При объявление войны вы оплатите взнос в размере
                <span class="importantly">100000.0 cr</span>. сразу и продолжите его
                платить каждую неделю до окончания войны.`,
		_const.EN: `When war is declared, you will pay a fee in the amount
                <span class="importantly">100000.0 cr</span>. immediately and continue it
                pay every week until the end of the war.`,
		_const.ZhCN: `宣战时，您将支付
                <span class="importantly">100000.0 cr</span>的费用，并每周支付直到战争结束。`,
	},
	"text_72": {
		_const.RU:   `на счету не хватает кредитов.`,
		_const.EN:   `there are not enough credits in the deposit.`,
		_const.ZhCN: `账户中信用点不足。`,
	},
	"text_73": {
		_const.RU:   `Укажите штаб`,
		_const.EN:   `Select headquarters`,
		_const.ZhCN: `选择总部`,
	},
	"option_31": {
		_const.RU:   `Не выбрано`,
		_const.EN:   `Not chosen`,
		_const.ZhCN: `未选择`,
	},
	"text_74": {
		_const.RU:   `Если ваш штаб уничтожат то война прекратится.`,
		_const.EN:   `If your headquarters is destroyed, the war will end.`,
		_const.ZhCN: `如果您的总部被摧毁，战争将结束。`,
	},
	"button_39": {
		_const.RU:   `Закрыть`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
	"button_40": {
		_const.RU:   `Объявить войну`,
		_const.EN:   `To declare a war`,
		_const.ZhCN: `宣战`,
	},
	"text_75": {
		_const.RU:   `Успех!`,
		_const.EN:   `Success!`,
		_const.ZhCN: `成功！`,
	},
	"text_76": {
		_const.RU: `<p>Вы хотите отозвать запрос о "статус-кво" с кластером <span
			class="importantly">%corporation_name%</span>?</p>
			<p>Если оба участника конфликта пошлют запрос о статусе кво то конфликт завершится, а боевые действия
			завершатся через 24ч.</p>`,
		_const.EN: `<p>You want to withdraw the request for the "status quo" with the cluster <span
			class="importantly">%corporation_name%</span>?</p>
			<p>If both parties to the conflict send a request for the status quo, the conflict will end and the fighting will
			will end in 24 hours.</p>`,
		_const.ZhCN: `<p>您想撤回与集群 <span
			class="importantly">%corporation_name%</span> 的“现状”请求吗？</p>
			<p>如果冲突双方都发送现状请求，冲突将结束，战斗将在24小时后停止。</p>`,
	},
	"text_77": {
		_const.RU: `<p>Вы хотите предложить запрос о "статус-кво" с кластером <span
			class="importantly">%corporation_name%</span>?</p>
			<p>Если оба участника конфликта пошлют запрос о статусе кво то конфликт завершится, а боевые действия
			завершатся через 24ч.</p>`,
		_const.EN: `<p>You want to propose a query about the "status quo" with a cluster <span
			class="importantly">%corporation_name%</span>?</p>
			<p>If both parties to the conflict send a request for the status quo, the conflict will end and the fighting will
			will end in 24 hours.</p>`,
		_const.ZhCN: `<p>您想向集群 <span
			class="importantly">%corporation_name%</span> 提出“现状”请求吗？</p>
			<p>如果冲突双方都发送现状请求，冲突将结束，战斗将在24小时后停止。</p>`,
	},
	"button_41": {
		_const.RU:   `Закрыть`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
	"button_42": {
		_const.RU:   `Отозвать`,
		_const.EN:   `Withdraw`,
		_const.ZhCN: `撤回`,
	},
	"button_43": {
		_const.RU:   `Предложить`,
		_const.EN:   `Offer`,
		_const.ZhCN: `提议`,
	},
	"text_78": {
		_const.RU:   `Успех!`,
		_const.EN:   `Success!`,
		_const.ZhCN: `成功！`,
	},
	"text_79": {
		_const.RU: `<p>Вы хотите убрать статус взаимной войны с кластером <span
				class="importantly">%corporation_name%</span>?</p>
				<p>Нападающая сторона начнет платить взнос за поддержание войны и сможет закончить ее если он не будет
				оплачен.</p>`,
		_const.EN: `<p>You want to remove the status of mutual war with the cluster <span
				class="importantly">%corporation_name%</span>?</p>
				<p>The attacking side will begin to pay a fee for maintaining the war and will be able to end it if he does not
				paid.</p>`,
		_const.ZhCN: `<p>您想移除与集群 <span
				class="importantly">%corporation_name%</span> 的互战状态吗？</p>
				<p>攻击方将开始支付战争维持费用，如果未支付，则可以结束战争。</p>`,
	},
	"text_80": {
		_const.RU: `<p>Вы хотите установить статус взаимной войны с кластером <span
				class="importantly">%corporation_name%</span>?</p>
				<p>Во взаимной войне нападающая сторона перестает платить взносы на поддержание войны и потеряет
				возможность самостоятельно завершить ее.</p>`,
		_const.EN: `<p>You want to set the status of mutual war with the cluster <span
				class="importantly">%corporation_name%</span>?</p>
				<p>In a mutual war, the attacking side stops paying contributions to maintain the war and will lose
				the ability to complete it yourself.</p>`,
		_const.ZhCN: `<p>您想设置与集群 <span
				class="importantly">%corporation_name%</span> 的互战状态吗？</p>
				<p>在互战中，攻击方将停止支付战争维持费用，并失去自行结束战争的能力。</p>`,
	},
	"button_44": {
		_const.RU:   `Закрыть`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
	"button_45": {
		_const.RU:   `Изменить`,
		_const.EN:   `Change`,
		_const.ZhCN: `更改`,
	},
	"text_81": {
		_const.RU:   `Статус:`,
		_const.EN:   `Status:`,
		_const.ZhCN: `状态：`,
	},
	"text_82": {
		_const.RU:   `Завершение`,
		_const.EN:   `Completion`,
		_const.ZhCN: `完成`,
	},
	"text_83": {
		_const.RU:   `Подготовка`,
		_const.EN:   `Preparation`,
		_const.ZhCN: `准备`,
	},
	"text_84": {
		_const.RU:   `Активна`,
		_const.EN:   `Active`,
		_const.ZhCN: `活跃`,
	},
	"text_85": {
		_const.RU:   `Штаб наступления:`,
		_const.EN:   `Offensive Headquarters:`,
		_const.ZhCN: `进攻总部：`,
	},
	"text_86": {
		_const.RU:   `Статус кво:`,
		_const.EN:   `Status quo:`,
		_const.ZhCN: `现状：`,
	},
	"text_87": {
		_const.RU:   `Наступательная`,
		_const.EN:   `Attack`,
		_const.ZhCN: `进攻`,
	},
	"text_88": {
		_const.RU:   `Оборонительная`,
		_const.EN:   `Defensive`,
		_const.ZhCN: `防御`,
	},
	"button_46": {
		_const.RU:   `Отозвать`,
		_const.EN:   `Withdraw`,
		_const.ZhCN: `撤回`,
	},
	"button_47": {
		_const.RU:   `Предложить`,
		_const.EN:   `Offer`,
		_const.ZhCN: `提议`,
	},
	"text_89": {
		_const.RU:   `Взаимная:`,
		_const.EN:   `Mutual:`,
		_const.ZhCN: `互战：`,
	},
	"button_48": {
		_const.RU:   `Изменить`,
		_const.EN:   `Change`,
		_const.ZhCN: `更改`,
	},
	"text_90": {
		_const.RU:   `Спонсируется с:`,
		_const.EN:   `Sponsored by:`,
		_const.ZhCN: `赞助者：`,
	},
	"text_91": {
		_const.RU:   `Оплата через:`,
		_const.EN:   `Payment via:`,
		_const.ZhCN: `支付方式：`,
	},
	"text_92": {
		_const.RU:   `Нет активных войн.`,
		_const.EN:   `There are no active wars.`,
		_const.ZhCN: `没有活跃的战争。`,
	},
	"button_49": {
		_const.RU:   `Объявить войну`,
		_const.EN:   `To declare a war`,
		_const.ZhCN: `宣战`,
	},
	"corporation": {
		_const.RU:   `Кластер`,
		_const.EN:   `Cluster`,
		_const.ZhCN: `集群`,
	},
	"tax": {
		_const.RU:   `Налог`,
		_const.EN:   `Tax`,
		_const.ZhCN: `税收`,
	},
}
