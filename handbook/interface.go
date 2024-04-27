package handbook

import _const "github.com/TrashPony/veliri-lib/const"

var IndexInterface = map[string]map[string]map[string]string{
	"Main": {},
	"Auth": {},
	"Wiki": {},
}

var GameInterface = map[string]map[string]map[string]string{
	"Alerts": {
		"reject": {
			_const.RU: `Отказать`,
			_const.EN: `Reject`,
		},
		"accept": {
			_const.RU: `Принять`,
			_const.EN: `Accept`,
		},
	},
	"AttackDialog": {
		"failed": {
			_const.RU: `Не удалось!`,
			_const.EN: `Failed!`,
		},
		"success": {
			_const.RU: `Успех!`,
			_const.EN: `Success!`,
		},
		"text_1": {
			_const.RU: `Сектор принадлежит кластеру:`,
			_const.EN: `The sector belongs to the cluster:`,
		},
		"text_2": {
			_const.RU: `Установить доступ на врата:`,
			_const.EN: `Set up access to the gate:`,
		},
		"text_3": {
			_const.RU: `Напасть на сектор?`,
			_const.EN: `Attack the sector?`,
		},
		"option_1": {
			_const.RU: `Кластер`,
			_const.EN: `Cluster`,
		},
		"attack_button": {
			_const.RU: `Напасть`,
			_const.EN: `Attack`,
		},
		"back_button": {
			_const.RU: `Назад`,
			_const.EN: `Back`,
		},
	},
	"AudioPlayer": {
		"settings": {
			_const.RU: `настройки`,
			_const.EN: `settings`,
		},
	},
	"BoxPass": {
		"window_name": {
			_const.RU: `Пароль`,
			_const.EN: `Password`,
		},
		"cancel": {
			_const.RU: `Отменить`,
			_const.EN: `Cancel`,
		},
		"ok": {
			_const.RU: `Ок`,
			_const.EN: `Ок`,
		},
	},
	"BuildMenu": {
		// BuildMenu
		"window_name": {
			_const.RU: `Строительство`,
			_const.EN: `Building`,
		},
		"tab_1": {
			_const.RU: `Строительство`,
			_const.EN: `Building`,
		},
		"tab_2": {
			_const.RU: `Разбор`,
			_const.EN: `Dismantling`,
		},
		"text_1": {
			_const.RU: `Что бы установить строение надо:`,
			_const.EN: `To install a building you need:`,
		},
		"text_2": {
			_const.RU: `Необходимо ресурсов для постройки:`,
			_const.EN: `Required resources for construction:`,
		},
		"text_3": {
			_const.RU: `Выбрать строение`,
			_const.EN: `Select building`,
		},
		"text_4": {
			_const.RU: `Отмена`,
			_const.EN: `Cancel`,
		},
		"text_5": {
			_const.RU: `Разбор`,
			_const.EN: `Dismantling`,
		},
		"text_6": {
			_const.RU: `Будет разобрано:`,
			_const.EN: `Will be dismantling:`,
		},
		"text_7": {
			_const.RU: `Ресурсы которые будут получены:`,
			_const.EN: `Resources that will be received:`,
		},
		// BuildTip
		"tip_help": {
			_const.RU: `<div><span class="importantly">A</span> / <span class="importantly">D</span> - поворот.</div>
						<div><span class="importantly">ПКМ</span> - установить.</div>
						<div><span class="importantly">ЛКМ</span> - отмена.</div>`,
			_const.EN: `<div><span class="importantly">A</span> / <span class="importantly">D</span> - turn.</div>
						<div><span class="importantly">RMB</span> - install.</div>
						<div><span class="importantly">LMB</span> - cancel.</div>`,
		},
		"tip_fail": {
			_const.RU: `Не удалось`,
			_const.EN: `Failed`,
		},
		"tip_request_1_1": {
			_const.RU: `Установить объект`,
			_const.EN: `Install object`,
		},
		"tip_request_1_2": {
			_const.RU: `за`,
			_const.EN: `for`,
		},
		"tip_button_1": {
			_const.RU: `Назад`,
			_const.EN: `Back`,
		},
		"tip_button_2": {
			_const.RU: `Установить`,
			_const.EN: `Install`,
		},
		// OpenBuildObject
		"text_8": {
			_const.RU: `Прогресс строительства`,
			_const.EN: `Construction progress`,
		},
		"button_1": {
			_const.RU: `Погрузить все`,
			_const.EN: `Submerge all`,
		},
		"button_2": {
			_const.RU: `Добавить`,
			_const.EN: `Add`,
		},
	},
	"Chat": {
		// Chat
		"window_name_1": {
			_const.RU: `Чат`,
			_const.EN: `Chat`,
		},
		"local": {
			_const.RU: `Локальный`,
			_const.EN: `Local`,
		},
		"local_greetings": {
			_const.RU: `Вы входите на территорию `,
			_const.EN: `You enter the territory `,
		},
		"local_greetings_wasteland": {
			_const.RU: `Пустоши...`,
			_const.EN: `Wasteland...`,
		},
		// CreateNewGroup
		"window_name_2": {
			_const.RU: `Создание нового канала`,
			_const.EN: `Creating a new channel`,
		},
		"text_1": {
			_const.RU: `Имя канала...`,
			_const.EN: `Channel name...`,
		},
		"text_2": {
			_const.RU: `Пароль (если пусто то без пароля)`,
			_const.EN: `Password (if empty then no password)`,
		},
		"text_3": {
			_const.RU: `Загрузить`,
			_const.EN: `Download`,
		},
		"text_4": {
			_const.RU: `Приветственное сообщение`,
			_const.EN: `Welcome message`,
		},
		"button_1": {
			_const.RU: `Создать`,
			_const.EN: `Create`,
		},
		"button_2": {
			_const.RU: `Закрыть`,
			_const.EN: `Close`,
		},
		// Messages
		"chanel_local": {
			_const.RU: `локация`,
			_const.EN: `local`,
		},
		"chanel_group": {
			_const.RU: `группа`,
			_const.EN: `group`,
		},
		"chanel_battle": {
			_const.RU: `бой`,
			_const.EN: `battle`,
		},
		"chanel_corporation": {
			_const.RU: `кластер`,
			_const.EN: `corporation`,
		},
		"chanel_private": {
			_const.RU: `личка`,
			_const.EN: `private`,
		},
		"chanel_unknown": {
			_const.RU: `Неизвестный`,
			_const.EN: `Unknown`,
		},
		// MessagesWrapper
		"default_greetings": {
			_const.RU: `Добро пожаловать!`,
			_const.EN: `Welcome!`,
		},
		// UserLine
		"text_5": {
			_const.RU: `Выберите роль`,
			_const.EN: `Select a role`,
		},
		"text_6": {
			_const.RU: `Место:`,
			_const.EN: `Place:`,
		},
		// UserSubMenu
		"sub_1": {
			_const.RU: `Подробнее`,
			_const.EN: `More details`,
		},
		"sub_2": {
			_const.RU: `Написать`,
			_const.EN: `Write to chat`,
		},
		"sub_3": {
			_const.RU: `Отправить письмо`,
			_const.EN: `Send mail`,
		},
		"sub_4": {
			_const.RU: `Сделать лидером`,
			_const.EN: `Make a leader`,
		},
		"sub_5": {
			_const.RU: `Пригласить в группу`,
			_const.EN: `Invite to group`,
		},
		"sub_6": {
			_const.RU: `Добавить в друзья`,
			_const.EN: `Add as Friend`,
		},
		"sub_7": {
			_const.RU: `Добавить контакт`,
			_const.EN: `Add contact`,
		},
		"sub_8": {
			_const.RU: `Выгнать`,
			_const.EN: `Kick out`,
		},
		"sub_9": {
			_const.RU: `Покинуть`,
			_const.EN: `Leave`,
		},
		"sub_10": {
			_const.RU: `Перевод кредитов`,
			_const.EN: `Transfer of credits`,
		},
		"sub_11": {
			_const.RU: `Закрыть`,
			_const.EN: `Close`,
		},
		// ViewAllGroups
		"button_3": {
			_const.RU: `Отмена`,
			_const.EN: `Cancel`,
		},
		"button_4": {
			_const.RU: `Войти`,
			_const.EN: `Enter`,
		},
		"placeholder_1": {
			_const.RU: `введите пароль`,
			_const.EN: `enter the password`,
		},
		"text_7": {
			_const.RU: `Основные группы:`,
			_const.EN: `Main groups:`,
		},
		"text_8": {
			_const.RU: `Пользовательские группы:`,
			_const.EN: `Player groups:`,
		},
		"button_5": {
			_const.RU: `Создать`,
			_const.EN: `Create`,
		},
		"button_6": {
			_const.RU: `Закрыть`,
			_const.EN: `Close`,
		},
	},
	"ChoiceFraction": {
		"greetings_1": {
			_const.RU: `Добро пожаловать на планету Veliri.`,
			_const.EN: `Welcome to planet Veliri.`,
		},
		"greetings_2": {
			_const.RU: `Нет времени объяснять выбирай серию машин и погнали.`,
			_const.EN: `There is no time to explain, choose a series of cars and let's go.`,
		},
		"Replics": {
			_const.RU: `<p><span class="importantly">Replics</span> - Автократической формы правления ИИ. Отголосок яростного экспансионизма “покровителей” и приверженца теории “войны за независимость с создателями”. Предпочитающий подчинять территории и ассимилировать другие технические формы жизни силой.</p>
                    	<p><span class="importantly">Replics</span> управляет обширной сетью лишённых самоуправления ботов. Но иногда, санкционируя создание серии автономных единиц, ради расширения сфер собственного влияния.</p>
                    	<p>Если у <span class="importantly">Replics</span> помимо господства и распространения, бывают иные поводы для интереса, то это только: автоматизация инфраструктуры, бесконечное производство и полные материалов склады.</p>`,
			_const.EN: `<p><span class="importantly">Replics</span> is an autocratic form of AI government. It is an echo of the fierce expansionism of «patrons» and an adherent of the theory of the «war of independence with creators». It prefers to subjugate territories and assimilate other technical life forms by force.</p>
						<p><span class="importantly">Replics</span> controls a vast network of bots deprived of self-government. But sometimes, by sanctioning the creation of a series of autonomous units, it expands its own spheres of influence.</p>
						<p>If <span class="importantly">Replics</span>, in addition to domination and expansion, has any other reasons for interest, then these are only: infrastructure automation, endless production, and warehouses full of materials.</p>`,
		},
		"Explores": {
			_const.RU: `<p><span class="importantly">Explores</span> - конгломерат великого множества разумов, что трудятся над единой задачей.</p>
                    	<p>В данном же случае, на <span class="importantly">Veliri-5</span>, <span class="importantly">Explores</span> движим идеей познания окружающего мира и применения полученных научных открытий, как с целью самообороны, так и улучшения уже существующего естества.</p>
                    	<p><span class="importantly">Explores</span> в отличии от <span class="importantly">Replics</span> более открыт к инновациям и не рассматривает окружающий мир, как очередную претензию для притязаний. Стараясь сохранить и приумножить уже имеющееся, и остановить то, что способно положить конец бесценным научным открытиям.</p>
                    	<p>В основном, именно данное мировоззрение и становится основополагающей частью вражды, между <span class="importantly">Replics</span> и <span class="importantly">Explores</span>.</p>`,
			_const.EN: `<p><span class="importantly">Explores</span> is a conglomerate of a great many minds working on a single task.</p>
						<p>In this case, on <span class="importantly">Veliri-5</span>, <span class="importantly">Explores</span> is driven by the idea of exploring the world around them and applying the scientific discoveries they make, both for the purpose of self-defense and to improve the existing nature.</p>
						<p><span class="importantly">Explores</span>, unlike <span class="importantly">Replics</span>, is more open to innovation and does not view the world around it as another claim to be made. Trying to preserve and increase what already exists, and to stop what could put an end to valuable scientific discoveries.</p>
						<p>Basically, it is this worldview that becomes the fundamental part of the hostility between <span class="importantly">Replics</span> and <span class="importantly">Explores</span>.</p>`,
		},
		"Reverses": {
			_const.RU: `<p><span class="importantly">Reverses</span> - симбиот биокибернетической формы жизни, с преобладанием утопизма и самоизоляции.</p>
                    	<p>Самый последний из представителей нового эволюционного направления. У <span class="importantly">Reverses</span> распространены такие явления, как право масс, индивидуальности и выбора. Именно по этой самой причине, большинство представителей данного течения ИИ - независимые платформы, что неустанно трудятся над технологией терраформинга.</p>
                    	<p>Возвращению мира <span class="importantly">Veliri-5</span> прежнего облика и дальнейшего сеяния жизни там, где это только возможно. Всячески закрывая глаза на действия и противоборства иных ИИ. Впрочем, не стоя в стороне, когда опасность угрожает им самим.</p>`,
			_const.EN: `<p><span class="importantly">Reverses</span> is a symbiont of a biocybernetic life form, with a predominance of utopianism and self-isolation.</p>
						<p>The latest representative of a new evolutionary direction. The <span class="importantly">Reverses</span> are characterized by such phenomena as the right of the masses, individuality and choice. It is for this reason that most representatives of this AI trend are independent platforms that tirelessly work on terraforming technology.</p>
						<p>Returning the world of <span class="importantly">Veliri-5</span> to its former appearance and further sowing life wherever possible. In every way turning a blind eye to the actions and confrontations of other AIs. However, not standing aside when danger threatens them.</p>`,
		},
	},
	"Corporation": {
		// Corporation
		"window_name_1": {
			_const.RU: `Кластер`,
			_const.EN: `Cluster`,
		},
		"tab_1": {
			_const.RU: `Поиск`,
			_const.EN: `Search`,
		},
		"tab_2": {
			_const.RU: `О кластере`,
			_const.EN: `About corporation`,
		},
		"tab_3": {
			_const.RU: `Участники`,
			_const.EN: `Members`,
		},
		"tab_4": {
			_const.RU: `Роли`,
			_const.EN: `Roles`,
		},
		"tab_5": {
			_const.RU: `Заявки`,
			_const.EN: `Applications`,
		},
		"tab_6": {
			_const.RU: `Политика`,
			_const.EN: `Politics`,
		},
		"tab_7": {
			_const.RU: `Войны`,
			_const.EN: `Wars`,
		},
		"tab_8": {
			_const.RU: `Главная`,
			_const.EN: `Main`,
		},
		"tab_9": {
			_const.RU: `Поиск`,
			_const.EN: `Search`,
		},
		"tab_10": {
			_const.RU: `Создать`,
			_const.EN: `Create`,
		},
		"text_1": {
			_const.RU: `Вы не состоите в кластере.`,
			_const.EN: `You are not in a cluster.`,
		},
		"button_1": {
			_const.RU: `Поиск`,
			_const.EN: `Search`,
		},
		"button_2": {
			_const.RU: `Создать`,
			_const.EN: `Create`,
		},
		// CorporationLine
		"no_cluster": {
			_const.RU: `Нет кластера`,
			_const.EN: `No cluster`,
		},
		// Create
		"helper_image": {
			_const.RU: `- jpg/png/gif <br>
						- не более 16КБ <br>
						- и 200х200px`,
			_const.EN: `- jpg/png/gif <br>
						- no more 16KB <br>
						- size 200х200px`,
		},
		"text_2": {
			_const.RU: `Создание кластера`,
			_const.EN: `Creating a Cluster`,
		},
		"button_3": {
			_const.RU: `Загрузить`,
			_const.EN: `Upload`,
		},
		"text_3": {
			_const.RU: `Название`,
			_const.EN: `Name`,
		},
		"text_4": {
			_const.RU: `(5-24 символов)`,
			_const.EN: `(5-24 characters)`,
		},
		"text_5": {
			_const.RU: `Тэг`,
			_const.EN: `Tag`,
		},
		"text_6": {
			_const.RU: `(2-5 символов)`,
			_const.EN: `(2-5 characters)`,
		},
		"text_7": {
			_const.RU: `Описание`,
			_const.EN: `Description`,
		},
		"button_4": {
			_const.RU: `Создать`,
			_const.EN: `Create`,
		},
		"create_corporation_player_already": {
			_const.RU: `Вы уже состоите в кластере`,
			_const.EN: `You are already a member of the cluster`,
		},
		"create_corporation_empty_fields": {
			_const.RU: `Не все поля заполнены`,
			_const.EN: `Not all fields are filled in`,
		},
		"create_corporation_name_wrong_size": {
			_const.RU: `Неправильно кол-во символов в имени`,
			_const.EN: `The number of characters in the name is incorrect`,
		},
		"create_corporation_tag_wrong_size": {
			_const.RU: `Неправильно кол-во символов в теге`,
			_const.EN: `Incorrect number of characters in tag`,
		},
		"create_corporation_biography_wrong_size": {
			_const.RU: `Описание слишком большое`,
			_const.EN: `Description too long`,
		},
		"create_corporation_name_already": {
			_const.RU: `Имя кластера уже используется`,
			_const.EN: `Cluster name is already in use`,
		},
		"create_corporation_tag_already": {
			_const.RU: `Тег кластера уже используется`,
			_const.EN: `The cluster tag is already in use`,
		},
		"create_corporation_logo_wrong_format": {
			_const.RU: `Формат файла логотипа не верный`,
			_const.EN: `Logo file format is not correct`,
		},
		"create_corporation_big_logo": {
			_const.RU: `Размер логотипа больше 16КБ`,
			_const.EN: `Logo size is larger than 16KB`,
		},
		"create_corporation_logo_wrong_resolution": {
			_const.RU: `Размер лого больше 200х200 пикселей`,
			_const.EN: `Logo size is more than 200x200 pixels`,
		},
		// Main
		"error_1": {
			_const.RU: `Не удалось`,
			_const.EN: `Failed`,
		},
		"rental_dialog_1": {
			_const.RU: `Арендовать офис на <span class='importantly'>30</span> дней за <span
class='importantly'>%credits_count% cr</span> c главного счета?`,
			_const.EN: `Rent an office for <span class='importantly'>30</span> days for <span
class='importantly'>%credits_count% cr</span> from the main deposit?`,
		},
		"rental_dialog_2": {
			_const.RU: `Продлить аренду офиса на <span class='importantly'>30</span> дней за <span
class='importantly'>%credits_count% cr</span> c главного счета?`,
			_const.EN: `Extend office rental for <span class='importantly'>30</span> days for <span
class='importantly'>%credits_count% cr</span> from the main deposit?`,
		},
		"button_5": {
			_const.RU: `Закрыть`,
			_const.EN: `Close`,
		},
		"button_6": {
			_const.RU: `Арендовать`,
			_const.EN: `Rent`,
		},
		"button_7": {
			_const.RU: `Продлить`,
			_const.EN: `Extend`,
		},
		"text_8": {
			_const.RU: `Заправить текущую базу на еще 30 дней?`,
			_const.EN: `Refill the current database for another 30 days?`,
		},
		"fill_up_dialog_1": {
			_const.RU: `Цена:
						<span class='importantly'> %credits_count% </span>
						cr. и
						<span class='importantly'> %thorium_count% </span> единиц тория c вашего
						склада.`,
			_const.EN: `Price:
						<span class='importantly'> %credits_count% </span>
						cr. and
						<span class='importantly'> %thorium_count% </span>  thorium units from your
						warehouse.`,
		},
		"button_8": {
			_const.RU: `Закрыть`,
			_const.EN: `Close`,
		},
		"button_9": {
			_const.RU: `Заправить`,
			_const.EN: `Refuel`,
		},
		"transfer_dialog_1": {
			_const.RU: `Переместиться на базу <span class='importantly'>%base_name%</span>?`,
			_const.EN: `Move to base <span class='importantly'>%base_name%</span>?`,
		},
		"transfer_dialog_2": {
			_const.RU: `Цена: <span class='importantly'>500.0</span> cr.`,
			_const.EN: `Price: <span class='importantly'>500.0</span> cr.`,
		},
		"transfer_dialog_3": {
			_const.RU: `<span class='importantly'>Внимание!</span> Переносится только сознание, транспорт и все его содержимое останется на этой базе.`,
			_const.EN: `<span class='importantly'>Attention!</span> Only consciousness, transport and all its contents are transferred will remain at this base.`,
		},
		"button_10": {
			_const.RU: `Закрыть`,
			_const.EN: `Close`,
		},
		"button_11": {
			_const.RU: `Перемещение`,
			_const.EN: `Moving`,
		},
		"deposit_history_h1": {
			_const.RU: `Время`,
			_const.EN: `Time`,
		},
		"deposit_history_h2": {
			_const.RU: `Операция`,
			_const.EN: `Operation`,
		},
		"deposit_history_h3": {
			_const.RU: `Операция`,
			_const.EN: `Cause`,
		},
		"button_12": {
			_const.RU: `Закрыть`,
			_const.EN: `Close`,
		},
		"text_9": {
			_const.RU: `Счета:`,
			_const.EN: `Deposits:`,
		},
		"button_13": {
			_const.RU: `Перевод`,
			_const.EN: `Translation`,
		},
		"button_14": {
			_const.RU: `История`,
			_const.EN: `Story`,
		},
		"text_10": {
			_const.RU: `Офисы:`,
			_const.EN: `Offices:`,
		},
		"text_11": {
			_const.RU: `База:`,
			_const.EN: `Base:`,
		},
		"text_12": {
			_const.RU: `До:`,
			_const.EN: `Until:`,
		},
		"button_15": {
			_const.RU: `Маршрут`,
			_const.EN: `Route`,
		},
		"button_16": {
			_const.RU: `Автопилот`,
			_const.EN: `Autopilot`,
		},
		"button_17": {
			_const.RU: `Переместиться`,
			_const.EN: `Move`,
		},
		"button_18": {
			_const.RU: `Продлить аренду`,
			_const.EN: `Extend your lease`,
		},
		"text_13": {
			_const.RU: `Арендовать офис на базе`,
			_const.EN: `Rent an office at the base`,
		},
		"text_14": {
			_const.RU: `Базы:`,
			_const.EN: `Bases:`,
		},
		"text_15": {
			_const.RU: `Сектор:`,
			_const.EN: `Sector:`,
		},
		"text_16": {
			_const.RU: `Строительство до:`,
			_const.EN: `Construction until:`,
		},
		"text_17": {
			_const.RU: `Неуязвима`,
			_const.EN: `Invulnerable`,
		},
		"text_18": {
			_const.RU: `Уязвима`,
			_const.EN: `Vulnerable`,
		},
		"text_19": {
			_const.RU: `Под атакой`,
			_const.EN: `Under attack`,
		},
		"text_20": {
			_const.RU: `Топлива хватит еще до:`,
			_const.EN: `There will be enough fuel until:`,
		},
		"text_21": {
			_const.RU: `Офис`,
			_const.EN: `Office`,
		},
		"text_22": {
			_const.RU: `Завод`,
			_const.EN: `Factory`,
		},
		"text_23": {
			_const.RU: `Переработчик`,
			_const.EN: `Recycler`,
		},
		"text_24": {
			_const.RU: `Рынок`,
			_const.EN: `Market`,
		},
		"button_19": {
			_const.RU: `Маршрут`,
			_const.EN: `Route`,
		},
		"button_20": {
			_const.RU: `Автопилот`,
			_const.EN: `Autopilot`,
		},
		"button_21": {
			_const.RU: `Заправить`,
			_const.EN: `Refuel`,
		},
		"text_25": {
			_const.RU: `Меню строительства`,
			_const.EN: `Construction menu`,
		},
		"text_26": {
			_const.RU: `Недоступно в текущем секторе и на базе`,
			_const.EN: `Not available in current sector and base`,
		},
		// Members
		"text_27": {
			_const.RU: `Участники в сети:`,
			_const.EN: `Members on the network:`,
		},
		"button_22": {
			_const.RU: `Назад`,
			_const.EN: `Back`,
		},
		"button_23": {
			_const.RU: `Управление ролями`,
			_const.EN: `Role management`,
		},
		"button_24": {
			_const.RU: `Сохранить`,
			_const.EN: `Save`,
		},
		// Policy
		"placeholder_1": {
			_const.RU: `Название кластера`,
			_const.EN: `Cluster name`,
		},
		"button_25": {
			_const.RU: `Найти`,
			_const.EN: `Find`,
		},
		"text_28": {
			_const.RU: `Ничего не найдено.`,
			_const.EN: `Nothing found.`,
		},
		"button_26": {
			_const.RU: `Закрыть`,
			_const.EN: `Close`,
		},
		"button_27": {
			_const.RU: `Добавить`,
			_const.EN: `Add`,
		},
		"text_29": {
			_const.RU: `Глобальная`,
			_const.EN: `Global`,
		},
		"text_30": {
			_const.RU: `Локальная`,
			_const.EN: `Local`,
		},
		"text_31": {
			_const.RU: `Союзники`,
			_const.EN: `Allies`,
		},
		"button_28": {
			_const.RU: `Добавить`,
			_const.EN: `Add`,
		},
		"text_32": {
			_const.RU: `Соперники`,
			_const.EN: `Rivals`,
		},
		"button_29": {
			_const.RU: `Добавить`,
			_const.EN: `Add`,
		},
		"text_33": {
			_const.RU: `Сектор ...`,
			_const.EN: `Sector ...`,
		},
		"text_34": {
			_const.RU: `Доступ в сектор:`,
			_const.EN: `Access to the sector:`,
		},
		"option_1": {
			_const.RU: `Все`,
			_const.EN: `All`,
		},
		"option_2": {
			_const.RU: `Все кроме врагов`,
			_const.EN: `Everything except enemies`,
		},
		"option_3": {
			_const.RU: `Все кроме союзников`,
			_const.EN: `Everyone except allies`,
		},
		"option_4": {
			_const.RU: `Все кроме соперников`,
			_const.EN: `Everyone except the rivals`,
		},
		"option_5": {
			_const.RU: `Кластер и союзники`,
			_const.EN: `Cluster and allies`,
		},
		"option_6": {
			_const.RU: `Кластер и соперники`,
			_const.EN: `Cluster and rivals`,
		},
		"option_7": {
			_const.RU: `Только Кластер`,
			_const.EN: `Cluster only`,
		},
		"text_35": {
			_const.RU: `Доступ на базу:`,
			_const.EN: `Access to the base:`,
		},
		"option_8": {
			_const.RU: `Все`,
			_const.EN: `All`,
		},
		"option_9": {
			_const.RU: `Все кроме врагов`,
			_const.EN: `Everything except enemies`,
		},
		"option_10": {
			_const.RU: `Все кроме союзников`,
			_const.EN: `Everyone except allies`,
		},
		"option_11": {
			_const.RU: `Все кроме соперников`,
			_const.EN: `Everyone except the rivals`,
		},
		"option_12": {
			_const.RU: `Кластер и союзники`,
			_const.EN: `Cluster and allies`,
		},
		"option_13": {
			_const.RU: `Кластер и соперники`,
			_const.EN: `Cluster and rivals`,
		},
		"option_14": {
			_const.RU: `Только Кластер`,
			_const.EN: `Cluster only`,
		},
		"text_36": {
			_const.RU: `Турели атакуют:`,
			_const.EN: `Turrets attack:`,
		},
		"option_15": {
			_const.RU: `Всех`,
			_const.EN: `Everyone`,
		},
		"option_16": {
			_const.RU: `Всех кроме кластера`,
			_const.EN: `Everyone except the cluster`,
		},
		"option_17": {
			_const.RU: `Всех кроме кластера и союзников`,
			_const.EN: `Everyone except the cluster and allies`,
		},
		"option_18": {
			_const.RU: `Всех кроме кластера и соперников`,
			_const.EN: `Everyone except the cluster and rivals`,
		},
		"option_19": {
			_const.RU: `Нейтралов и врагов`,
			_const.EN: `Neutrals and enemies`,
		},
		"option_20": {
			_const.RU: `Врагов (пк, апд и врагов в войне)`,
			_const.EN: `Enemies (PK, APD and enemies in war)`,
		},
		"option_21": {
			_const.RU: `Отключены`,
			_const.EN: `Disabled`,
		},
		"text_37": {
			_const.RU: `Ремонтники чинят:`,
			_const.EN: `Repairmen repair:`,
		},
		"option_22": {
			_const.RU: `Кластер`,
			_const.EN: `Cluster`,
		},
		"option_23": {
			_const.RU: `Кластер и союзников`,
			_const.EN: `Cluster and allies`,
		},
		"option_24": {
			_const.RU: `Кластер и соперников`,
			_const.EN: `Cluster and rivals`,
		},
		"option_25": {
			_const.RU: `Всех кроме врагов`,
			_const.EN: `Everyone except enemies`,
		},
		"option_26": {
			_const.RU: `Отключены`,
			_const.EN: `Disabled`,
		},
		"text_38": {
			_const.RU: `Налог на переработку:`,
			_const.EN: `Processing tax:`,
		},
		"text_39": {
			_const.RU: `Налог на производство:`,
			_const.EN: `Production tax:`,
		},
		"text_40": {
			_const.RU: `Налог на рынок:`,
			_const.EN: `Market tax:`,
		},
		"text_41": {
			_const.RU: `Склад для сбора налога:`,
			_const.EN: `Tax collection warehouse:`,
		},
		"option_27": {
			_const.RU: `Склад лидера`,
			_const.EN: `Leader warehouse`,
		},
		"text_42": {
			_const.RU: `Счет для сбора налога:`,
			_const.EN: `Tax collection deposit:`,
		},
		"option_28": {
			_const.RU: `Склад лидера`,
			_const.EN: `Leader deposit`,
		},
		"text_43": {
			_const.RU: `Окно уязвимости`,
			_const.EN: `Vulnerability window`,
		},
		// Requests
		"text_44": {
			_const.RU: `Прием заявок:`,
			_const.EN: `Accepting applications:`,
		},
		"text_45": {
			_const.RU: `Нет прав на прием игроков`,
			_const.EN: `No rights to accept players`,
		},
		"button_30": {
			_const.RU: `Принять`,
			_const.EN: `Accept`,
		},
		"button_31": {
			_const.RU: `Отказать`,
			_const.EN: `Reject`,
		},
		// Roles
		"text_46": {
			_const.RU: `Роль:`,
			_const.EN: `Role:`,
		},
		"option_29": {
			_const.RU: `Выберите роль`,
			_const.EN: `Select a role`,
		},
		"text_47": {
			_const.RU: `Общие:`,
			_const.EN: `Common:`,
		},
		"text_48": {
			_const.RU: `Приглашать игроков`,
			_const.EN: `Invite players`,
		},
		"text_49": {
			_const.RU: `Выгонять игроков`,
			_const.EN: `Kick players`,
		},
		"text_50": {
			_const.RU: `Редактировать роли`,
			_const.EN: `Edit roles`,
		},
		"text_51": {
			_const.RU: `Менять роли игрокам`,
			_const.EN: `Change roles for players`,
		},
		"text_52": {
			_const.RU: `Счета:`,
			_const.EN: `Deposits:`,
		},
		"text_53": {
			_const.RU: `Смотреть`,
			_const.EN: `Look`,
		},
		"text_54": {
			_const.RU: `Снимать`,
			_const.EN: `Take off`,
		},
		"text_55": {
			_const.RU: `Локальные:`,
			_const.EN: `Local:`,
		},
		"text_56": {
			_const.RU: `Выберите офис`,
			_const.EN: `Select an office`,
		},
		"text_57": {
			_const.RU: `Смотреть`,
			_const.EN: `Look`,
		},
		"text_58": {
			_const.RU: `Класть`,
			_const.EN: `Putting`,
		},
		"text_59": {
			_const.RU: `Забирать`,
			_const.EN: `Take`,
		},
		"text_60": {
			_const.RU: `Добавить роль:`,
			_const.EN: `Add role:`,
		},
		"placeholder_2": {
			_const.RU: `Название роли`,
			_const.EN: `Role name`,
		},
		"button_32": {
			_const.RU: `Добавить`,
			_const.EN: `Add`,
		},
		// Search
		"text_61": {
			_const.RU: `Поиск кластера`,
			_const.EN: `Search for a cluster`,
		},
		"placeholder_3": {
			_const.RU: `Название кластера`,
			_const.EN: `Cluster Name`,
		},
		"text_62": {
			_const.RU: `моя заявка:`,
			_const.EN: `my request:`,
		},
		"button_33": {
			_const.RU: `отозвать`,
			_const.EN: `recall`,
		},
		"button_34": {
			_const.RU: `изменить`,
			_const.EN: `change`,
		},
		"button_35": {
			_const.RU: `Подать заявку`,
			_const.EN: `Send request`,
		},
		"placeholder_4": {
			_const.RU: `Текст заявки`,
			_const.EN: `Request text`,
		},
		"button_36": {
			_const.RU: `Отправить`,
			_const.EN: `Send`,
		},
		"button_37": {
			_const.RU: `Отмена`,
			_const.EN: `Cancel`,
		},
		// Wars
		"text_63": {
			_const.RU: `Успех!`,
			_const.EN: `Success!`,
		},
		"text_64": {
			_const.RU: `База уже является штабом для другой войны`,
			_const.EN: `The base is already the headquarters for another war`,
		},
		"text_65": {
			_const.RU: `Такой базы не существует`,
			_const.EN: `No such database exists`,
		},
		"text_66": {
			_const.RU: `С этим кластером уже ведется война`,
			_const.EN: `There is already a war going on with this cluster`,
		},
		"text_67": {
			_const.RU: `У цели нет баз, вы не можете объявить ему войну`,
			_const.EN: `The target has no bases, you cannot declare war on him`,
		},
		"text_68": {
			_const.RU: `Вы пока не можете объявить войну этому кластеру`,
			_const.EN: `You cannot declare war on this cluster yet`,
		},
		"placeholder_5": {
			_const.RU: `Название кластера`,
			_const.EN: `Cluster name`,
		},
		"button_38": {
			_const.RU: `Найти`,
			_const.EN: `Find`,
		},
		"text_69": {
			_const.RU: `Ничего не найдено.`,
			_const.EN: `Nothing found.`,
		},
		"text_70": {
			_const.RU: `Спонсировать с`,
			_const.EN: `Sponsor with`,
		},
		"option_30": {
			_const.RU: `Не выбрано`,
			_const.EN: `Not chosen`,
		},
		"text_71": {
			_const.RU: `При объявление войны вы оплатите взнос в размере
                <span class="importantly">100000.0 cr</span>. сразу и продолжите его
                платить каждую неделю до окончания войны.`,
			_const.EN: `When war is declared, you will pay a fee in the amount
                <span class="importantly">100000.0 cr</span>. immediately and continue it
                pay every week until the end of the war.`,
		},
		"text_72": {
			_const.RU: `на счету не хватает кредитов.`,
			_const.EN: `there are not enough credits in the deposit.`,
		},
		"text_73": {
			_const.RU: `Укажите штаб`,
			_const.EN: `Select headquarters`,
		},
		"option_31": {
			_const.RU: `Не выбрано`,
			_const.EN: `Not chosen`,
		},
		"text_74": {
			_const.RU: `Если ваш штаб уничтожат то война прекратится.`,
			_const.EN: `If your headquarters is destroyed, the war will end.`,
		},
		"button_39": {
			_const.RU: `Закрыть`,
			_const.EN: `Close`,
		},
		"button_40": {
			_const.RU: `Объявить войну`,
			_const.EN: `To declare a war`,
		},
		"text_75": {
			_const.RU: `Успех!`,
			_const.EN: `Success!`,
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
		},
		"button_41": {
			_const.RU: `Закрыть`,
			_const.EN: `Close`,
		},
		"button_42": {
			_const.RU: `Отозвать`,
			_const.EN: `Withdraw`,
		},
		"button_43": {
			_const.RU: `Предложить`,
			_const.EN: `Offer`,
		},
		"text_78": {
			_const.RU: `Успех!`,
			_const.EN: `Success!`,
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
		},
		"button_44": {
			_const.RU: `Закрыть`,
			_const.EN: `Close`,
		},
		"button_45": {
			_const.RU: `Изменить`,
			_const.EN: `Change`,
		},
		"text_81": {
			_const.RU: `Статус:`,
			_const.EN: `Status:`,
		},
		"text_82": {
			_const.RU: `Завершение`,
			_const.EN: `Completion`,
		},
		"text_83": {
			_const.RU: `Подготовка`,
			_const.EN: `Preparation`,
		},
		"text_84": {
			_const.RU: `Активна`,
			_const.EN: `Active`,
		},
		"text_85": {
			_const.RU: `Штаб наступления:`,
			_const.EN: `Offensive Headquarters:`,
		},
		"text_86": {
			_const.RU: `Статус кво:`,
			_const.EN: `Status quo:`,
		},
		"text_87": {
			_const.RU: `Наступательная`,
			_const.EN: `Attack`,
		},
		"text_88": {
			_const.RU: `Оборонительная`,
			_const.EN: `Defensive`,
		},
		"button_46": {
			_const.RU: `Отозвать`,
			_const.EN: `Withdraw`,
		},
		"button_47": {
			_const.RU: `Предложить`,
			_const.EN: `Offer`,
		},
		"text_89": {
			_const.RU: `Взаимная:`,
			_const.EN: `Mutual:`,
		},
		"button_48": {
			_const.RU: `Изменить`,
			_const.EN: `Change`,
		},
		"text_90": {
			_const.RU: `Спонсируется с:`,
			_const.EN: `Sponsored by:`,
		},
		"text_91": {
			_const.RU: `Оплата через:`,
			_const.EN: `Payment via:`,
		},
		"text_92": {
			_const.RU: `Нет активных войн.`,
			_const.EN: `There are no active wars.`,
		},
		"button_49": {
			_const.RU: `Объявить войну`,
			_const.EN: `To declare a war`,
		},
	},
	"Department":     {},
	"DetailItem":     {},
	"EquipBar":       {},
	"FractionMarket": {},
	"FriendList":     {},
	"Gate": {
		"window_name": {
			_const.RU: `Вход`,
			_const.EN: `Entry`,
		},
		"attention_1": {
			_const.RU: `<h3 class="not_card_notify">Внимание!</h3>
						<p class="not_card_notify_text">Похоже в вашем браузере отключено
						  <a class="not_card_notify_text_link" target="_blank"
							 href="https://www.google.com/search?q=%D0%B0%D0%BF%D0%BF%D0%B0%D1%80%D0%B0%D1%82%D0%BD%D0%BE%D0%B5+%D1%83%D1%81%D0%BA%D0%BE%D1%80%D0%B5%D0%BD%D0%B8%D0%B5+%D0%B1%D1%80%D0%B0%D1%83%D0%B7%D0%B5%D1%80%D0%B0">
							аппаратное ускорение</a>, без него может сильно снизится производительность.</p>`,
			_const.EN: `<h3 class="not_card_notify">Attention!</h3>
						<p class="not_card_notify_text">It seems that hardware acceleration is disabled in your browser. Performance may be significantly reduced without it.</p>`,
		},
		"play": {
			_const.RU: `Играть`,
			_const.EN: `Play`,
		},
		"createPlayerLarchName": {
			_const.RU: `В имени персонажа не может быть больше 16ти символов.`,
			_const.EN: `A character's name cannot contain more than 16 characters.`,
		},
		"createPlayerNameNotAvailable": {
			_const.RU: `Имя уже занято.`,
			_const.EN: `The name is already taken.`,
		},
		"createPlayerSmallName": {
			_const.RU: `В имени персонажа не может быть меньше 3х символов.`,
			_const.EN: `A character's name must be at least 3 characters long.`,
		},
		"create": {
			_const.RU: `Создать`,
			_const.EN: `Create`,
		},
		"back": {
			_const.RU: `назад`,
			_const.EN: `back`,
		},
		"enter_name": {
			_const.RU: `Введите имя`,
			_const.EN: `Enter your name`,
		},
		"player_name": {
			_const.RU: `Имя персонажа`,
			_const.EN: "Character `s name",
		},
	},
	"GeoScan":           {},
	"Global":            {},
	"GlobalMap":         {},
	"GroupMenu":         {},
	"Hangar":            {},
	"Help":              {},
	"HelpControl":       {},
	"Index":             {},
	"Inventory":         {},
	"Loader":            {},
	"Lobby":             {},
	"Mail":              {},
	"Market":            {},
	"MenuBar":           {},
	"MiniGames":         {},
	"MiniMap":           {},
	"Missions":          {},
	"ModalNotify":       {},
	"News":              {},
	"NPCTrade":          {},
	"ObjectDialog":      {},
	"PrefabricatedMenu": {},
	"Preloader":         {},
	"ProcessorRoot":     {},
	"SelectTarget":      {},
	"SendCredits":       {},
	"ServerState":       {},
	"Settings":          {},
	"StatusBar":         {},
	"TechnologyMenu":    {},
	"Training":          {},
	"Transfer":          {},
	"UserStat":          {},
	"Window":            {},
	"Workbench":         {},
}
