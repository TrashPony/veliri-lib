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
						<div><span class="importantly">ЛКМ</span> - установить.</div>
						<div><span class="importantly">ПКМ</span> - отмена.</div>`,
			_const.EN: `<div><span class="importantly">A</span> / <span class="importantly">D</span> - turn.</div>
						<div><span class="importantly">LMB</span> - install.</div>
						<div><span class="importantly">RMB</span> - cancel.</div>`,
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
		"button_1": {
			_const.RU: `Выбрать`,
			_const.EN: `Choose`,
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
		"tab_11": {
			_const.RU: `Аудит`,
			_const.EN: `Audit`,
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
			_const.RU: `Комментарий`,
			_const.EN: `Comment`,
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
			_const.RU: `Счет лидера`,
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
		"build": {
			_const.RU: `Строительство`,
			_const.EN: `Building`,
		},
		"dismantling": {
			_const.RU: `Демонтаж строений`,
			_const.EN: `Dismantling of buildings`,
		},
		"change_local_policy": {
			_const.RU: `Изменение локальной политики`,
			_const.EN: `Change local policy`,
		},
		// Search
		"text_61": {
			_const.RU: `Фильтры:`,
			_const.EN: `Filters:`,
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
		"button_50": {
			_const.RU: `Поиск`,
			_const.EN: `Search`,
		},
		"option_32": {
			_const.RU: `Все`,
			_const.EN: `All`,
		},
		"option_33": {
			_const.RU: `Открыт`,
			_const.EN: `Open`,
		},
		"option_34": {
			_const.RU: `Закрыт`,
			_const.EN: `Close`,
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
		// Audit
		"corporation": {
			_const.RU: `Кластер`,
			_const.EN: `Cluster`,
		},
		"tax": {
			_const.RU: `Налог`,
			_const.EN: `Tax`,
		},
	},
	"Department": {
		"window_name_1": {
			_const.RU: `Терминал`,
			_const.EN: `Terminal`,
		},
		"text_1": {
			_const.RU: `Пункт назначения:`,
			_const.EN: `Destination:`,
		},
		"text_2": {
			_const.RU: `Сектор`,
			_const.EN: `Sector`,
		},
		"text_3": {
			_const.RU: `База`,
			_const.EN: `Base`,
		},
		"button_1": {
			_const.RU: `карта`,
			_const.EN: `map`,
		},
		"text_4": {
			_const.RU: `Награда при завершение задания:`,
			_const.EN: `Reward upon completion of the task:`,
		},
	},
	"DetailItem": {
		// BonusTab
		"button_1": {
			_const.RU: `Навыки`,
			_const.EN: `Skills`,
		},
		// CharacteristicTab
		"small": {
			_const.RU: `малый`,
			_const.EN: `small	`,
		},
		"average": {
			_const.RU: `средний`,
			_const.EN: `average	`,
		},
		"large": {
			_const.RU: `большой`,
			_const.EN: `large`,
		},
		"firearms": {
			_const.RU: `огнестрельное`,
			_const.EN: `firearms`,
		},
		"missile": {
			_const.RU: `ракетное`,
			_const.EN: `missile`,
		},
		"laser": {
			_const.RU: `лазерное`,
			_const.EN: `laser`,
		},
		"m3": {
			_const.RU: `м^3`,
			_const.EN: `m^3	`,
		},
		"km": {
			_const.RU: `км/с`,
			_const.EN: `km/s`,
		},
		"insec": {
			_const.RU: `в сек`,
			_const.EN: `per sec`,
		},
		"unitsec": {
			_const.RU: `ед/с`,
			_const.EN: `units/s`,
		},
		"percentsec": {
			_const.RU: `%/с`,
			_const.EN: `%/s`,
		},
		"s": {
			_const.RU: `c`,
			_const.EN: `s`,
		},
		"deg": {
			_const.RU: `гр/с`,
			_const.EN: `deg/s`,
		},
		"rate_fire": {
			_const.RU: `выс./мин`,
			_const.EN: `r/min`,
		},
		"ch_1": {
			_const.RU: `Занимаемый объем`,
			_const.EN: `Occupied volume`,
		},
		"ch_2": {
			_const.RU: `Вместимость`,
			_const.EN: `Capacity`,
		},
		"ch_3": {
			_const.RU: `Прочность`,
			_const.EN: `Strength`,
		},
		"ch_4": {
			_const.RU: `Занимаемый объем в трюме`,
			_const.EN: `Occupied volume in the hold`,
		},
		"ch_5": {
			_const.RU: `Защищен паролем`,
			_const.EN: `Password protected`,
		},
		"ch_6": {
			_const.RU: `Находится под землей`,
			_const.EN: `Located underground`,
		},

		"ch_7": {
			_const.RU: `Размер корпуса`,
			_const.EN: `Body size`,
		},
		"ch_8": {
			_const.RU: `Вместимость трюма`,
			_const.EN: `Hold capacity`,
		},
		"ch_9": {
			_const.RU: `Дополнительные отсеки:`,
			_const.EN: `Additional compartments:`,
		},
		"ch_10": {
			_const.RU: `Вместимость`,
			_const.EN: `Capacity`,
		},
		"ch_11": {
			_const.RU: `Допустимые предметы`,
			_const.EN: `Acceptable Items`,
		},
		"ch_12": {
			_const.RU: `Ходовая`,
			_const.EN: `Running gear`,
		},
		"ch_13": {
			_const.RU: `Макс. скорость`,
			_const.EN: `Max speed`,
		},
		"ch_14": {
			_const.RU: `Скорость поворота`,
			_const.EN: `Turning speed`,
		},
		"ch_15": {
			_const.RU: `Навигация`,
			_const.EN: `Navigation`,
		},
		"ch_16": {
			_const.RU: `Дальность обзора`,
			_const.EN: `View range`,
		},
		"ch_17": {
			_const.RU: `Дальность радара`,
			_const.EN: `Radar range`,
		},
		"ch_18": {
			_const.RU: `Накопитель`,
			_const.EN: `Battery`,
		},
		"ch_19": {
			_const.RU: `Вместимость накопителя`,
			_const.EN: `Energy storage capacity`,
		},
		"ch_20": {
			_const.RU: `Скорость зарядки`,
			_const.EN: `Charging speed`,
		},
		"ch_21": {
			_const.RU: `Выживаемость:`,
			_const.EN: `Survival:`,
		},
		"ch_22": {
			_const.RU: `Структура`,
			_const.EN: `Structure`,
		},
		"ch_23": {
			_const.RU: `Защита:`,
			_const.EN: `Protection:`,
		},
		"ch_24": {
			_const.RU: `Броня и уязвимые места`,
			_const.EN: `Armor and vulnerabilities`,
		},
		"ch_164": {
			_const.RU: `Щит`,
			_const.EN: `Shield`,
		},
		"ch_165": {
			_const.RU: `Защита щита`,
			_const.EN: `Shield protection`,
		},
		"ch_166": {
			_const.RU: `Отсутствует`,
			_const.EN: `Missing`,
		},
		"ch_25": {
			_const.RU: `Размер боеприпаса`,
			_const.EN: `Ammunition size`,
		},
		"ch_26": {
			_const.RU: `Тип оружия`,
			_const.EN: `Weapon type`,
		},
		"ch_27": {
			_const.RU: `Урон`,
			_const.EN: `Damage`,
		},
		"ch_28": {
			_const.RU: `Тип атаки`,
			_const.EN: `Attack type`,
		},
		"ch_29": {
			_const.RU: `Радиус поражения`,
			_const.EN: `Damage radius`,
		},
		"ch_30": {
			_const.RU: `Скорость полета пули`,
			_const.EN: `Bullet speed`,
		},
		"ch_31": {
			_const.RU: `Занимаемый объем`,
			_const.EN: `Occupied volume`,
		},
		"ch_32": {
			_const.RU: `Особенность:`,
			_const.EN: `Feature:`,
		},
		"ch_33": {
			_const.RU: `Ручной захват цели`,
			_const.EN: `Target lock`,
		},
		"ch_34": {
			_const.RU: `Управляемый снаряд`,
			_const.EN: `Guided projectile`,
		},
		"ch_35": {
			_const.RU: `Автоматический захват цели`,
			_const.EN: `Automatic target lock`,
		},
		"ch_36": {
			_const.RU: `Макс. прочность`,
			_const.EN: `Max strength`,
		},
		"ch_37": {
			_const.RU: `Радиус щита`,
			_const.EN: `Shield radius`,
		},
		"ch_38": {
			_const.RU: `Восстановление прочности`,
			_const.EN: `Restoring Strength`,
		},
		"ch_39": {
			_const.RU: `Потребляемая энергия`,
			_const.EN: `Energy consumption`,
		},
		"ch_40": {
			_const.RU: `Потребляемая энергия в простое`,
			_const.EN: `Energy consumption when idle`,
		},
		"ch_41": {
			_const.RU: `Время восстановления после урона`,
			_const.EN: `Recovery time after damage`,
		},
		"ch_42": {
			_const.RU: `Радиус мемех`,
			_const.EN: `Interference radius`,
		},
		"ch_43": {
			_const.RU: `Потребляемая энергия`,
			_const.EN: `Energy consumption`,
		},
		"ch_44": {
			_const.RU: `Радиус защиты`,
			_const.EN: `Protection radius`,
		},
		"ch_45": {
			_const.RU: `Потребляемая энергия за сбитие`,
			_const.EN: `Energy consumption per shoot down`,
		},
		"ch_46": {
			_const.RU: `Время перезарядки`,
			_const.EN: `Recharge time`,
		},
		"ch_47": {
			_const.RU: `Дальность ремонта`,
			_const.EN: `Repair range`,
		},
		"ch_48": {
			_const.RU: `Время перезарядки`,
			_const.EN: `Recharge time`,
		},
		"ch_49": {
			_const.RU: `Потребляемая энергия`,
			_const.EN: `Energy consumption`,
		},
		"ch_50": {
			_const.RU: `Кол-во восстанавливаемой прочности`,
			_const.EN: `Amount of restored strength`,
		},
		"ch_51": {
			_const.RU: `Объем добываемой руды`,
			_const.EN: `Volume of ore mined`,
		},
		"ch_52": {
			_const.RU: `Разработка руды`,
			_const.EN: `Ore development`,
		},
		"ch_53": {
			_const.RU: `Дальность добычи`,
			_const.EN: `Mining range`,
		},
		"ch_54": {
			_const.RU: `Потребление энергии`,
			_const.EN: `Power consumption`,
		},
		"ch_55": {
			_const.RU: `Наполнение бочки`,
			_const.EN: `Filling the barrel`,
		},
		"ch_56": {
			_const.RU: `Дальность шланга`,
			_const.EN: `Hose range`,
		},
		"ch_57": {
			_const.RU: `Скорость перезарядки`,
			_const.EN: `Recharge speed`,
		},
		"ch_58": {
			_const.RU: `Потребление энергии за такт`,
			_const.EN: `Energy consumption per cycle`,
		},
		"ch_59": {
			_const.RU: `Дальность управления`,
			_const.EN: `Control range`,
		},
		"ch_60": {
			_const.RU: `Радиус раскопок`,
			_const.EN: `Excavation radius`,
		},
		"ch_61": {
			_const.RU: `Перезарядка дрона`,
			_const.EN: `Recharging the drone`,
		},
		"ch_62": {
			_const.RU: `Энергия для активации дрона`,
			_const.EN: `Energy to activate the drone`,
		},
		"ch_63": {
			_const.RU: `Дальность`,
			_const.EN: `Range`,
		},
		"ch_64": {
			_const.RU: `Скорость перезарядки`,
			_const.EN: `Recharge speed`,
		},
		"ch_65": {
			_const.RU: `Потребление энергии`,
			_const.EN: `Power consumption`,
		},
		"ch_66": {
			_const.RU: `Влияние на прочность`,
			_const.EN: `Impact on strength`,
		},
		"ch_67": {
			_const.RU: `Дальность передачи энергии`,
			_const.EN: `Energy transmission range`,
		},
		"ch_68": {
			_const.RU: `Передаваемая энергия за такт`,
			_const.EN: `Transmitted energy per cycle`,
		},
		"ch_69": {
			_const.RU: `Скорость перезарядки`,
			_const.EN: `Recharge speed`,
		},
		"ch_70": {
			_const.RU: `Чувствительность датчика`,
			_const.EN: `Sensor sensitivity`,
		},
		"ch_71": {
			_const.RU: `Потребление энергии за сканирование`,
			_const.EN: `Energy consumption per scan`,
		},
		"ch_72": {
			_const.RU: `Перезарядка`,
			_const.EN: `Reload`,
		},
		"ch_73": {
			_const.RU: `Урон`,
			_const.EN: `Damage`,
		},
		"ch_74": {
			_const.RU: `Радиус поражения`,
			_const.EN: `Damage radius`,
		},
		"ch_75": {
			_const.RU: `Задержка детонации после установки`,
			_const.EN: `Detonation delay after installation`,
		},
		"ch_76": {
			_const.RU: `Задержка детонации после срабатывания`,
			_const.EN: `Detonation delay after actuation`,
		},
		"ch_77": {
			_const.RU: `Радиус обнаружения цели`,
			_const.EN: `Target detection radius`,
		},
		"ch_78": {
			_const.RU: `Кол-во зарядов`,
			_const.EN: `Number of charges`,
		},
		"ch_79": {
			_const.RU: `Потребляемая энергия за использование`,
			_const.EN: `Energy consumption per use`,
		},
		"ch_80": {
			_const.RU: `Время перезарядки`,
			_const.EN: `Recharge time`,
		},
		"ch_81": {
			_const.RU: `Дальность обзора`,
			_const.EN: `View range`,
		},
		"ch_82": {
			_const.RU: `Структура`,
			_const.EN: `Structure`,
		},
		"ch_83": {
			_const.RU: `Время жизни`,
			_const.EN: `Lifetime`,
		},
		"ch_84": {
			_const.RU: `Дальность полета капсулы с турелью`,
			_const.EN: `Flight range of a capsule with a turret`,
		},
		"ch_85": {
			_const.RU: `Кол-во зарядов`,
			_const.EN: `Number of charges`,
		},
		"ch_86": {
			_const.RU: `Потребляемая энергия за использование`,
			_const.EN: `Energy consumption per use`,
		},
		"ch_87": {
			_const.RU: `Время перезарядки`,
			_const.EN: `Recharge time`,
		},
		"ch_88": {
			_const.RU: `Радиус закрытия обзора одного заряда`,
			_const.EN: `Radius of closing the view of one charge`,
		},
		"ch_89": {
			_const.RU: `Кол-во зарядов`,
			_const.EN: `Number of charges`,
		},
		"ch_90": {
			_const.RU: `Потребляемая энергия за использование`,
			_const.EN: `Energy consumption per use`,
		},
		"ch_91": {
			_const.RU: `Время перезарядки`,
			_const.EN: `Recharge time`,
		},
		"ch_92": {
			_const.RU: `Радиус закрытия обзора одного заряда`,
			_const.EN: `Radius of closing the view of one charge`,
		},
		"ch_93": {
			_const.RU: `Дальность полета капсул`,
			_const.EN: `Capsule flight range`,
		},
		"ch_94": {
			_const.RU: `Кол-во зарядов`,
			_const.EN: `Number of charges`,
		},
		"ch_95": {
			_const.RU: `Кол-во капсул в одном заряде`,
			_const.EN: `Number of capsules in one charge`,
		},
		"ch_96": {
			_const.RU: `Потребляемая энергия за использование`,
			_const.EN: `Energy consumption per use`,
		},
		"ch_97": {
			_const.RU: `Время перезарядки`,
			_const.EN: `Recharge time`,
		},
		"ch_98": {
			_const.RU: `Дальность обзора`,
			_const.EN: `View range`,
		},
		"ch_99": {
			_const.RU: `Время жизни`,
			_const.EN: `Lifetime`,
		},
		"ch_100": {
			_const.RU: `Потребляемая энергия за использование`,
			_const.EN: `Energy consumption per use`,
		},
		"ch_101": {
			_const.RU: `Время перезарядки`,
			_const.EN: `Recharge time`,
		},
		"ch_102": {
			_const.RU: `Дальность обзора`,
			_const.EN: `View range`,
		},
		"ch_103": {
			_const.RU: `Время жизни`,
			_const.EN: `Lifetime`,
		},
		"ch_104": {
			_const.RU: `Потребляемая энергия за использование`,
			_const.EN: `Energy consumption per use`,
		},
		"ch_105": {
			_const.RU: `Время перезарядки`,
			_const.EN: `Recharge time`,
		},
		"ch_106": {
			_const.RU: `Дальность захвата целей`,
			_const.EN: `Target acquisition range`,
		},
		"ch_107": {
			_const.RU: `Структура`,
			_const.EN: `Structure`,
		},
		"ch_108": {
			_const.RU: `Время жизни ловушки`,
			_const.EN: `Trap lifetime`,
		},
		"ch_109": {
			_const.RU: `Время жизни ловушки после активации`,
			_const.EN: `Trap lifetime after activation`,
		},
		"ch_110": {
			_const.RU: `Кол-во зарядов`,
			_const.EN: `Number of charges`,
		},
		"ch_111": {
			_const.RU: `Потребляемая энергия за использование`,
			_const.EN: `Energy consumption per use`,
		},
		"ch_112": {
			_const.RU: `Время перезарядки`,
			_const.EN: `Recharge time`,
		},
		"ch_113": {
			_const.RU: `Дальность полета`,
			_const.EN: `Range of flight`,
		},
		"ch_114": {
			_const.RU: `Кол-во зарядов`,
			_const.EN: `Number of charges`,
		},
		"ch_115": {
			_const.RU: `Потребляемая энергия за использование`,
			_const.EN: `Energy consumption per use`,
		},
		"ch_116": {
			_const.RU: `Время перезарядки`,
			_const.EN: `Recharge time`,
		},
		"ch_117": {
			_const.RU: `Дальность полета`,
			_const.EN: `Range of flight`,
		},
		"ch_118": {
			_const.RU: `Сила притяжения`,
			_const.EN: `Force of gravity`,
		},
		"ch_119": {
			_const.RU: `Радиус`,
			_const.EN: `Radius`,
		},
		"ch_120": {
			_const.RU: `Время жизни`,
			_const.EN: `Lifetime`,
		},
		"ch_121": {
			_const.RU: `Кол-во зарядов`,
			_const.EN: `Number of charges`,
		},
		"ch_122": {
			_const.RU: `Потребляемая энергия за использование`,
			_const.EN: `Energy consumption per use`,
		},
		"ch_123": {
			_const.RU: `Время перезарядки`,
			_const.EN: `Recharge time`,
		},
		"ch_124": {
			_const.RU: `Сила прыжка`,
			_const.EN: `Jump Power`,
		},
		"ch_125": {
			_const.RU: `Потребляемая энергия за использование`,
			_const.EN: `Energy consumption per use`,
		},
		"ch_126": {
			_const.RU: `Время перезарядки`,
			_const.EN: `Recharge time`,
		},
		"ch_127": {
			_const.RU: `Время активации`,
			_const.EN: `Activation time`,
		},
		"ch_128": {
			_const.RU: `Потребляемая энергия`,
			_const.EN: `Energy consumption`,
		},
		"ch_129": {
			_const.RU: `Время перезарядки`,
			_const.EN: `Recharge time`,
		},
		"ch_130": {
			_const.RU: `Структура`,
			_const.EN: `Structure`,
		},
		"ch_131": {
			_const.RU: `Дальность полета капсулы со стеной`,
			_const.EN: `Flight range of a capsule with a wall`,
		},
		"ch_132": {
			_const.RU: `Кол-во зарядов`,
			_const.EN: `Number of charges`,
		},
		"ch_133": {
			_const.RU: `Потребляемая энергия за использование`,
			_const.EN: `Energy consumption per use`,
		},
		"ch_134": {
			_const.RU: `Время перезарядки`,
			_const.EN: `Recharge time`,
		},
		"ch_135": {
			_const.RU: `Макс. прочность`,
			_const.EN: `Max strength`,
		},
		"ch_136": {
			_const.RU: `Радиус щита`,
			_const.EN: `Shield radius`,
		},
		"ch_137": {
			_const.RU: `Время жизни`,
			_const.EN: `Lifetime`,
		},
		"ch_138": {
			_const.RU: `Восстановление прочности`,
			_const.EN: `Restoring Strength`,
		},
		"ch_139": {
			_const.RU: `Время восстановления после урона`,
			_const.EN: `Recovery time after damage`,
		},
		"ch_140": {
			_const.RU: `Структура`,
			_const.EN: `Structure`,
		},
		"ch_141": {
			_const.RU: `Дальность полета капсулы`,
			_const.EN: `Capsule flight range`,
		},
		"ch_142": {
			_const.RU: `Кол-во зарядов`,
			_const.EN: `Number of charges`,
		},
		"ch_143": {
			_const.RU: `Потребляемая энергия за использование`,
			_const.EN: `Energy consumption per use`,
		},
		"ch_144": {
			_const.RU: `Время перезарядки`,
			_const.EN: `Recharge time`,
		},
		"ch_145": {
			_const.RU: `Радиус`,
			_const.EN: `Radius`,
		},
		"ch_146": {
			_const.RU: `Потребляемая энергия за использование`,
			_const.EN: `Energy consumption per use`,
		},
		"ch_147": {
			_const.RU: `Потребляя мощность`,
			_const.EN: `Consuming power`,
		},
		"ch_148": {
			_const.RU: `Занимаемый объем`,
			_const.EN: `Occupied volume`,
		},
		"ch_149": {
			_const.RU: `Размер оружия`,
			_const.EN: `Weapon size`,
		},
		"ch_150": {
			_const.RU: `Тип оружия`,
			_const.EN: `Weapon type`,
		},
		"ch_151": {
			_const.RU: `Скорость поворота`,
			_const.EN: `Swing speed`,
		},
		"ch_152": {
			_const.RU: `Модификатор урона`,
			_const.EN: `Damage modifier`,
		},
		"ch_153": {
			_const.RU: `Доступные углы атаки`,
			_const.EN: `Available angles of attack`,
		},
		"ch_154": {
			_const.RU: `Влияние на скорость пули`,
			_const.EN: `Effect on bullet speed`,
		},
		"ch_155": {
			_const.RU: `Дальность стрельбы`,
			_const.EN: `Range attack`,
		},
		"ch_156": {
			_const.RU: `зависит от снаряда`,
			_const.EN: `depends on the projectile`,
		},
		"ch_157": {
			_const.RU: `Очередь`,
			_const.EN: `Queue`,
		},
		"ch_158": {
			_const.RU: `Скорострельность`,
			_const.EN: `Rate of fire`,
		},
		"ch_159": {
			_const.RU: `Боезапас`,
			_const.EN: `Ammunition`,
		},
		"ch_160": {
			_const.RU: `Время перезарядки`,
			_const.EN: `Reload time`,
		},
		"ch_161": {
			_const.RU: `Потребляя мощность`,
			_const.EN: `Consuming power`,
		},
		"ch_162": {
			_const.RU: `Занимаемый объем`,
			_const.EN: `Occupied volume`,
		},
		"ch_163": {
			_const.RU: `артиллерия`,
			_const.EN: `artillery`,
		},
		// DetailItem
		"text_1": {
			_const.RU: `примерная цена: <span class="importantly">%credits_count%</span> cr.`,
			_const.EN: `approximate price: <span class="importantly">%credits_count%</span> cr.`,
		},
		"text_2": {
			_const.RU: `Описание`,
			_const.EN: `Description`,
		},
		"text_3": {
			_const.RU: `Особенности`,
			_const.EN: `Peculiarities`,
		},
		"text_4": {
			_const.RU: `Хар-ки`,
			_const.EN: `Specs`,
		},
		"text_5": {
			_const.RU: `Хар-ки оружия`,
			_const.EN: `Weapon specs.`,
		},
		"text_6": {
			_const.RU: `Хар-ки снаряда`,
			_const.EN: `Projectile specs.`,
		},
		"text_7": {
			_const.RU: `Варианты`,
			_const.EN: `Options`,
		},
		"text_8": {
			_const.RU: `Требования`,
			_const.EN: `Requirements`,
		},
		"text_9": {
			_const.RU: `Оснащение`,
			_const.EN: `Equipment`,
		},
		"text_10": {
			_const.RU: `Произ-во`,
			_const.EN: `Production`,
		},
		"broken_tip": {
			_const.RU: `Предмет сломан из за этого:<br>
          - его нельзя переработать или продать<br>
          - если выкинуть из трюма он уничтожиться<br>
          - его характеристики могут быть ухудшены<br>`,
			_const.EN: `The item is broken, which means that:<br>
          - it cannot be recycled or sold<br>
          - if you throw it away, it will be destroyed<br>
          - its characteristics may be reduced<br>`,
		},
		"resource_tip_1_1": {
			_const.RU: `Этот ресурс можно получить`,
			_const.EN: `This resource can be obtained by `,
		},
		"resource_tip_1_2": {
			_const.RU: `переработав`,
			_const.EN: `processing`,
		},
		"resource_tip_2_1": {
			_const.RU: `Этот ресурс можно`,
			_const.EN: `This resource can be`,
		},
		"resource_tip_2_2": {
			_const.RU: `добыть в пустошах`,
			_const.EN: `mined in the wasteland`,
		},
		"blueprints_tip_1_1": {
			_const.RU: `Чертежи можно получить из обломков или исследуя`,
			_const.EN: `Blueprints can be obtained from debris or by exploring`,
		},
		"blueprints_tip_1_2": {
			_const.RU: `аномалии в пустошах`,
			_const.EN: `anomalies in the wasteland`,
		},
		"detail_tip_1": {
			_const.RU: `Детали можно получить:`,
			_const.EN: `Details can be obtained:`,
		},
		"detail_tip_1_1": {
			_const.RU: `создав их на заводе полуфабрикатов`,
			_const.EN: `creating them at a semi-finished products factory`,
		},
		"detail_tip_1_2": {
			_const.RU: `создав их по чертежам`,
			_const.EN: `creating them according to drawings`,
		},
		"detail_tip_1_3": {
			_const.RU: `переработав обломки`,
			_const.EN: `recycling the wreckage`,
		},
		// EquipmentTab
		"text_11": {
			_const.RU: `Ячейки для топлива`,
			_const.EN: `Fuel cells`,
		},
		"text_12": {
			_const.RU: `Лимит мощности`,
			_const.EN: `Power limit`,
		},
		"text_13": {
			_const.RU: `Слотов оружия`,
			_const.EN: `Weapon slots`,
		},
		"text_14": {
			_const.RU: `Разъемы большой мощности`,
			_const.EN: `High Power Connectors`,
		},
		"text_15": {
			_const.RU: `Разъемы средней мощности`,
			_const.EN: `Medium Power Connectors`,
		},
		"text_16": {
			_const.RU: `Разъемы малой мощности`,
			_const.EN: `Low Power Connectors`,
		},
		"text_17": {
			_const.RU: `Боеприпасы`,
			_const.EN: `Ammunition`,
		},
		// ProductionTab
		"text_18": {
			_const.RU: `Продукция`,
			_const.EN: `Products`,
		},
		"text_19": {
			_const.RU: `Требуемые ресурсы <span style="font-size: 9px">(при 100% эф.)</span>`,
			_const.EN: `Required resources <span style="font-size: 9px">(at 100% ef.)</span>`,
		},
		"text_20": {
			_const.RU: `Можно получить переработав`,
			_const.EN: `Can be obtained by recycling`,
		},
		"text_21": {
			_const.RU: `Чертежи`,
			_const.EN: `Blueprints`,
		},
		"text_22": {
			_const.RU: `Компоненты после переработки`,
			_const.EN: `Components after processing`,
		},
		"text_24": {
			_const.RU: `Переработав <span style="color: yellow">%count%</span> ед. выйдет:`,
			_const.EN: `Having processed <span style="color: yellow">%count%</span> units. will come out:`,
		},
		// Requirements
		"text_23": {
			_const.RU: `Навыки`,
			_const.EN: `Skills`,
		},
	},
	"EquipBar": {},
	"FractionMarket": {
		"window_name": {
			_const.RU: `Фракционный магазин`,
			_const.EN: `Faction Store`,
		},
		"to_storage": {
			_const.RU: `Предметы отправлены на <span class="importantly">склад</span>`,
			_const.EN: `Items sent to <span class="importantly">warehouse</span>`,
		},
		"error_1": {
			_const.RU: `Не удалось`,
			_const.EN: `Failed`,
		},
		"text_1": {
			_const.RU: `Фракционный магазин`,
			_const.EN: `Faction Store`,
		},
		"text_2": {
			_const.RU: `Кол-во очков фракции`,
			_const.EN: `Faction Points`,
		},
		"text_3": {
			_const.RU: `оф.`,
			_const.EN: `p.`,
		},
		"text_4": {
			_const.RU: `Магазин доступен только членам ополчения.`,
			_const.EN: `The store is only available to militia members.`,
		},
		"text_5": {
			_const.RU: `Предмет`,
			_const.EN: `Item`,
		},
		"text_6": {
			_const.RU: `Кол-во`,
			_const.EN: `Qty.`,
		},
		"text_7": {
			_const.RU: `Необходимый ранг`,
			_const.EN: `Required rank`,
		},
		"text_8": {
			_const.RU: `Стоимость`,
			_const.EN: `Price`,
		},
		"button_1": {
			_const.RU: `Купить`,
			_const.EN: `Buy`,
		},
		"text_9": {
			_const.RU: `Приобрести`,
			_const.EN: `Price`,
		},
		"text_10": {
			_const.RU: `за`,
			_const.EN: `for`,
		},
		"text_11": {
			_const.RU: `и`,
			_const.EN: `and`,
		},
		"button_2": {
			_const.RU: `Отмена`,
			_const.EN: `Cancel`,
		},
		"button_3": {
			_const.RU: `Купить`,
			_const.EN: `Buy`,
		},
	},
	"FriendList": {
		"window_name": {
			_const.RU: `Контакты`,
			_const.EN: `Contacts`,
		},
		"tab_1": {
			_const.RU: `Друзья`,
			_const.EN: `Friends`,
		},
		"tab_2": {
			_const.RU: `Контакты`,
			_const.EN: `Contacts`,
		},
		"tab_3": {
			_const.RU: `Заблокированные`,
			_const.EN: `Blocked`,
		},
		"placeholder_1": {
			_const.RU: `Имя игрока`,
			_const.EN: `Username`,
		},
		"button_1": {
			_const.RU: `Отправить запрос`,
			_const.EN: `Send request`,
		},
		"button_2": {
			_const.RU: `Добавить контакт`,
			_const.EN: `Add contact`,
		},
	},
	"Game": {},
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
	"GeoScan": {
		"window_name": {
			_const.RU: `Сканер`,
			_const.EN: `Scanner`,
		},
	},
	"Global": {
		"text_1": {
			_const.RU: `Помощь`,
			_const.EN: `Help`,
		},
		"text_2": {
			_const.RU: `Доберитесь до указанной на <span style="color: yellow">мини-карте</span> зоне.`,
			_const.EN: `Get to the area indicated on the <span style="color: yellow">minimap</span>.`,
		},
		"text_3": {
			_const.RU: `Примечание:`,
			_const.EN: `Note:`,
		},
		"text_4": {
			_const.RU: `движение.`,
			_const.EN: `movement.`,
		},
		"text_5": {
			_const.RU: `ускорение.`,
			_const.EN: `acceleration.`,
		},
		"text_6": {
			_const.RU: `тормоз.`,
			_const.EN: `brake.`,
		},
		"text_7": {
			_const.RU: `Мышь`,
			_const.EN: `Mouse`,
		},
		"text_8": {
			_const.RU: `направление движения.`,
			_const.EN: `direction of movement.`,
		},
		"text_9": {
			_const.RU: `Уничтожьте бочки и баки на пути <span style="color: grey; font-size: 12px">(рекомендуется активировать щит)</span>.`,
			_const.EN: `Destroy the barrels and tanks on the way <span style="color: grey; font-size: 12px">(it is recommended to activate the shield)</span>.`,
		},
		"text_10": {
			_const.RU: `Примечание:`,
			_const.EN: `Note:`,
		},
		"text_11": {
			_const.RU: `Что бы активировать оружие/снаряжение его надо выбрать на панели в <span
            style="color: yellow">центре с низу
            экрана</span> или нажать на указанной на ячейке клавишу (<span style="color: yellow">2</span> или '<span
            style="color: yellow">E</span>').`,
			_const.EN: `To activate the weapons/equipment, you need to select it on the panel in the <span style="color: yellow">center
			at the bottom of the screen</span> or press the key indicated in the cell (<span style="color: yellow">2</span>
			or '<span style="color: yellow">E</span>').`,
		},
		"text_12": {
			_const.RU: `прицеливание.`,
			_const.EN: `aiming.`,
		},
		"text_13": {
			_const.RU: `ЛКМ`,
			_const.EN: `LMB`,
		},
		"text_14": {
			_const.RU: `атака.`,
			_const.EN: `attack.`,
		},
		"text_15": {
			_const.RU: `Красный круг`,
			_const.EN: `Red circle`,
		},
		"text_16": {
			_const.RU: `зона поражения.`,
			_const.EN: `affected area.`,
		},
		"text_17": {
			_const.RU: `перезарядка.`,
			_const.EN: `reloading.`,
		},
		"text_18": {
			_const.RU: `Неизвестная угроза!`,
			_const.EN: `Unknown threat!`,
		},
		"text_19": {
			_const.RU: `Вы являетесь врагом в этом секторе!`,
			_const.EN: `You are the enemy in this sector!`,
		},
		"text_20": {
			_const.RU: `включить фары.`,
			_const.EN: `turn off the headlights.`,
		},
		"text_21": {
			_const.RU: `Что бы выбрать цель, <span style="color: yellow">наведите мышку</span> на транспорт и нажмите <span style="color: yellow">TAB</span>.`,
			_const.EN: `To select a target, <span style="color: yellow">move the mouse</span> over the vehicle and press <span style="color: yellow">TAB</span>.`,
		},
		"text_22": {
			_const.RU: `Что бы убрать цель, нажмите <span style="color: yellow">TAB</span> в месте где нет транспорта.`,
			_const.EN: `To remove a target, press <span style="color: yellow">TAB</span> in a place where there is no vehicle.`,
		},
	},
	"GlobalMap": {
		// GlobalMap
		"window_name": {
			_const.RU: `Карта мира`,
			_const.EN: `World map`,
		},
		// MapPin
		"button_1": {
			_const.RU: `Автопилот`,
			_const.EN: `Autopilot`,
		},
		"button_2": {
			_const.RU: `Установить цель`,
			_const.EN: `Set a path`,
		},
		"text_1": {
			_const.RU: `В секторе идет бой!`,
			_const.EN: `There is a battle in the sector!`,
		},
		"text_2": {
			_const.RU: `Атака возможна:`,
			_const.EN: `Attack possible:`,
		},
		"button_3": {
			_const.RU: `Закрыть`,
			_const.EN: `Close`,
		},
	},
	"GroupMenu": {
		"window_name": {
			_const.RU: `Отряд`,
			_const.EN: `Squad`,
		},
		"text_1": {
			_const.RU: `Вы не состоите в отряде...`,
			_const.EN: `You are not a member of the squad...`,
		},
	},
	"Hangar": {
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
	},
	"Help": {
		"window_name": {
			_const.RU: `Помощь`,
			_const.EN: `Help`,
		},
		"text_1": {
			_const.RU: `Добро подалось в базу вспомогательных материалов Veliri!`,
			_const.EN: `Welcome to the Veliri support materials database!`,
		},
	},
	"HelpControl": {
		"text_1": {
			_const.RU: `Управление:`,
			_const.EN: `Control:`,
		},
		"text_2": {
			_const.RU: `Стрельба:`,
			_const.EN: `Shooting:`,
		},
		"text_3": {
			_const.RU: `Меню:`,
			_const.EN: `Menu:`,
		},
	},
	"Inventory": {
		// ArmModal
		"text_1": {
			_const.RU: `за %count%:`,
			_const.EN: `for %count%:`,
		},
		"text_2": {
			_const.RU: `Всего:`,
			_const.EN: `Total:`,
		},
		"text_3": {
			_const.RU: `Кол-во:`,
			_const.EN: `Qty:`,
		},
		"button_1": {
			_const.RU: `Закрыть`,
			_const.EN: `Close`,
		},
		"button_2": {
			_const.RU: `Зарядить`,
			_const.EN: `Arm`,
		},
		// BackgroundItemCell
		// CellSubMenu
		"button_3": {
			_const.RU: `Установить оружие`,
			_const.EN: `Install weapon`,
		},
		"button_4": {
			_const.RU: `Установить снаряды`,
			_const.EN: `Install ammo`,
		},
		"button_5": {
			_const.RU: `Активировать корпус`,
			_const.EN: `Activate the body`,
		},
		"button_6": {
			_const.RU: `Заправить корпус`,
			_const.EN: `Fill body`,
		},
		"button_7": {
			_const.RU: `Подробнее`,
			_const.EN: `More details`,
		},
		"button_8": {
			_const.RU: `Разделить`,
			_const.EN: `Split`,
		},
		"button_9": {
			_const.RU: `Выбросить`,
			_const.EN: `Throw`,
		},
		"button_10": {
			_const.RU: `Уничтожить`,
			_const.EN: `Destroy`,
		},
		"button_11": {
			_const.RU: `Установить`,
			_const.EN: `Install`,
		},
		"button_12": {
			_const.RU: `Продать`,
			_const.EN: `Sell`,
		},
		"button_13": {
			_const.RU: `Починить`,
			_const.EN: `Repair`,
		},
		"button_14": {
			_const.RU: `Снять`,
			_const.EN: `Take off`,
		},
		"button_15": {
			_const.RU: `В трюм`,
			_const.EN: `To hold`,
		},
		"button_16": {
			_const.RU: `На склад`,
			_const.EN: `To storage`,
		},
		"button_17": {
			_const.RU: `Произвести`,
			_const.EN: `Produce`,
		},
		"button_18": {
			_const.RU: `Отменить`,
			_const.EN: `Cancel`,
		},
		"text_4": {
			_const.RU: `Кол-во:`,
			_const.EN: `Qty:`,
		},
		"text_5": {
			_const.RU: `Сr. за шт.`,
			_const.EN: `Cr. per pce`,
		},
		"text_6": {
			_const.RU: `Всего:`,
			_const.EN: `Total:`,
		},
		"text_7": {
			_const.RU: `Налог (%tax%%)`,
			_const.EN: `Tax (%tax%%)`,
		},
		"text_8": {
			_const.RU: `Вы получите:`,
			_const.EN: `You'll get:`,
		},
		"button_19": {
			_const.RU: `Отмена`,
			_const.EN: `Cancel`,
		},
		"button_20": {
			_const.RU: `Продать`,
			_const.EN: `Sell`,
		},
		"button_21": {
			_const.RU: `Отмена`,
			_const.EN: `Cancel`,
		},
		"button_22": {
			_const.RU: `Разделить`,
			_const.EN: `Divide`,
		},
		// FillUpModal
		"text_9": {
			_const.RU: `Сr. за шт.`,
			_const.EN: `Cr. per pce`,
		},
		"text_10": {
			_const.RU: `Всего:`,
			_const.EN: `Total:`,
		},
		"text_11": {
			_const.RU: `Кол-во:`,
			_const.EN: `Qty:`,
		},
		"button_23": {
			_const.RU: `Закрыть`,
			_const.EN: `Close`,
		},
		"button_24": {
			_const.RU: `Заправить`,
			_const.EN: `Refuel`,
		},
		// Inventory
		"window_name": {
			_const.RU: `Инвентарь`,
			_const.EN: `Inventory`,
		},
		"text_12": {
			_const.RU: `Склад заблокирован. Продлите аренду офиса или восстановите модуль базы.`,
			_const.EN: `The warehouse is locked. Please extend the office rental or restore the database module.`,
		},
		"text_13": {
			_const.RU: `Нет прав.`,
			_const.EN: `You have no rights.`,
		},
		"text_14": {
			_const.RU: `Офис`,
			_const.EN: `Office`,
		},
		"text_15": {
			_const.RU: `Склады`,
			_const.EN: `Warehouses`,
		},
		"button_25": {
			_const.RU: `Автопилот`,
			_const.EN: `Autopilot`,
		},
		"button_26": {
			_const.RU: `Проложить маршрут`,
			_const.EN: `Plot a route.`,
		},
		"text_16": {
			_const.RU: `Руда`,
			_const.EN: `Ore`,
		},
		"text_20": {
			_const.RU: `Органика`,
			_const.EN: `Organic`,
		},
		"storage": {
			_const.RU: `Склад`,
			_const.EN: `Storage`,
		},
		"squadInventory": {
			_const.RU: `Трюм`,
			_const.EN: `Hold`,
		},
		"scanner": {
			_const.RU: `Сканер`,
			_const.EN: `Scanner`,
		},
		"box": {
			_const.RU: `Ящик`,
			_const.EN: `Box`,
		},
		// ItemCell
		"cell_tax": {
			_const.RU: `Налог:`,
			_const.EN: `Tax:`,
		},
		// PlaceDialog
		"text_17": {
			_const.RU: `Не все влезает:`,
			_const.EN: `Not everything fits:`,
		},
		// Range
		// RepairModal
		"text_18": {
			_const.RU: `Починить`,
			_const.EN: `Repair`,
		},
		"for": {
			_const.RU: `за`,
			_const.EN: `for`,
		},
		"text_19": {
			_const.RU: `Не нуждается в ремонте.`,
			_const.EN: `Does not need to be repaired.`,
		},
		"button_27": {
			_const.RU: `Закрыть`,
			_const.EN: `Close`,
		},
		"button_28": {
			_const.RU: `Починить`,
			_const.EN: `Repair`,
		},
		// Size
		// WarehouseTab
	},
	"Loader": {
		"window_name": {
			_const.RU: `Загрузка...`,
			_const.EN: `Loading...`,
		},
		"text_1": {
			_const.RU: `Происходит что то неведомое... ¯\_(ツ)_/¯`,
			_const.EN: `Something mysterious is happening... ¯\_(ツ)_/¯`,
		},
	},
	"Lobby": {
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
	},
	"Mail": {
		"window_name": {
			_const.RU: `Почта`,
			_const.EN: `Mail`,
		},
		"error_1": {
			_const.RU: `При отправке произошла ошибка.`,
			_const.EN: `There was an error while sending.`,
		},
		"tab_1": {
			_const.RU: `Входящие`,
			_const.EN: `Inbox`,
		},
		"tab_2": {
			_const.RU: `Отправленные`,
			_const.EN: `Sent`,
		},
		"tab_3": {
			_const.RU: `Написать`,
			_const.EN: `Write`,
		},
		"tab_4": {
			_const.RU: `Корзина`,
			_const.EN: `Trash`,
		},
		"from": {
			_const.RU: `от`,
			_const.EN: `from`,
		},
		"date": {
			_const.RU: `дата`,
			_const.EN: `date`,
		},
		"button_1": {
			_const.RU: `Закрыть`,
			_const.EN: `Close`,
		},
		"button_2": {
			_const.RU: `Удалить`,
			_const.EN: `Delete`,
		},
		"button_3": {
			_const.RU: `Ответить`,
			_const.EN: `Answer`,
		},
		"placeholder_1": {
			_const.RU: `Адресат`,
			_const.EN: `Destination`,
		},
		"button_4": {
			_const.RU: `Контакты`,
			_const.EN: `Contacts`,
		},
		"button_5": {
			_const.RU: `Найти`,
			_const.EN: `Search`,
		},
		"text_1": {
			_const.RU: `Ничего не найдено.`,
			_const.EN: `Nothing found.`,
		},
		"placeholder_2": {
			_const.RU: `Тема`,
			_const.EN: `Subject`,
		},
		"placeholder_3": {
			_const.RU: `Сообщение`,
			_const.EN: `Message`,
		},
		"button_6": {
			_const.RU: `Отправить`,
			_const.EN: `Send`,
		},
		"text_2": {
			_const.RU: `Сообщение отправлено`,
			_const.EN: `The message has been sent.`,
		},
	},
	"Market": {
		// Buy
		"text_1": {
			_const.RU: `Кол-во:`,
			_const.EN: `Qty:`,
		},
		"text_2": {
			_const.RU: `Сr. за шт.`,
			_const.EN: `Cr. per pce`,
		},
		"text_3": {
			_const.RU: `Всего:`,
			_const.EN: `Total:`,
		},
		"button_1": {
			_const.RU: `Отменить`,
			_const.EN: `Cancel`,
		},
		"button_2": {
			_const.RU: `Купить`,
			_const.EN: `Buy`,
		},
		// Filter
		"cat_1": {
			_const.RU: `Боеприпасы`,
			_const.EN: `Ammunition`,
		},
		"cat_2": {
			_const.RU: `Оружие`,
			_const.EN: `Weapon`,
		},
		"cat_3": {
			_const.RU: `Корпуса`,
			_const.EN: `Bodies`,
		},
		"cat_4": {
			_const.RU: `Снаряжение`,
			_const.EN: `Equipment`,
		},
		"cat_5": {
			_const.RU: `Ресурсы`,
			_const.EN: `Resources`,
		},
		"cat_6": {
			_const.RU: `Чертежи`,
			_const.EN: `Blueprints`,
		},
		"cat_7": {
			_const.RU: `Ящики`,
			_const.EN: `Boxes`,
		},
		"cat_8": {
			_const.RU: `Хлам`,
			_const.EN: `Rubbish`,
		},
		"cat_9": {
			_const.RU: `Товары`,
			_const.EN: `Products`,
		},
		"size_t1": {
			_const.RU: `Малые`,
			_const.EN: `Small`,
		},
		"size_t2": {
			_const.RU: `Средние`,
			_const.EN: `Medium`,
		},
		"size_t3": {
			_const.RU: `Большие`,
			_const.EN: `Large`,
		},
		"equip_t3": {
			_const.RU: `Модификаторы`,
			_const.EN: `Modifiers`,
		},
		"body_t1": {
			_const.RU: `Легкие`,
			_const.EN: `Light`,
		},
		"body_t2": {
			_const.RU: `Средние`,
			_const.EN: `Medium`,
		},
		"body_t3": {
			_const.RU: `Тяжелые`,
			_const.EN: `Heavy`,
		},
		"bp_t1": {
			_const.RU: `Боеприпасы`,
			_const.EN: `Ammunition`,
		},
		"bp_t2": {
			_const.RU: `Корпуса`,
			_const.EN: `Bodies`,
		},
		"bp_t3": {
			_const.RU: `Оружие`,
			_const.EN: `Weapons`,
		},
		"bp_t4": {
			_const.RU: `Снаряжение`,
			_const.EN: `Equipments`,
		},
		"bp_t5": {
			_const.RU: `Детали`,
			_const.EN: `Details`,
		},
		"bp_t6": {
			_const.RU: `Ящики`,
			_const.EN: `Boxes`,
		},
		"bp_t7": {
			_const.RU: `Структуры`,
			_const.EN: `Structures`,
		},
		"r_t1": {
			_const.RU: `Ископаемые`,
			_const.EN: `Fossils`,
		},
		"r_t2": {
			_const.RU: `Сырье`,
			_const.EN: `Raw`,
		},
		"r_t3": {
			_const.RU: `Детали`,
			_const.EN: `Details`,
		},
		"a_t1": {
			_const.RU: `Лазерные`,
			_const.EN: `Laser`,
		},
		"a_t2": {
			_const.RU: `Ракетные`,
			_const.EN: `Missile`,
		},
		"a_t3": {
			_const.RU: `Баллистические`,
			_const.EN: `Ballistic`,
		},
		"w_t1": {
			_const.RU: `Лазерное`,
			_const.EN: `Laser`,
		},
		"w_t2": {
			_const.RU: `Ракетное`,
			_const.EN: `Missile`,
		},
		"w_t3": {
			_const.RU: `Баллистическое`,
			_const.EN: `Ballistic`,
		},
		// Market
		"window_name": {
			_const.RU: `Рынок`,
			_const.EN: `Market`,
		},
		"error_1": {
			_const.RU: `Модуль сервиса на безе уничтожен или отключен.`,
			_const.EN: `The service module based on the base has been destroyed or disabled.`,
		},
		"sub_button_1": {
			_const.RU: `Подробно`,
			_const.EN: `Details`,
		},
		"sub_button_2": {
			_const.RU: `Продать`,
			_const.EN: `Sell`,
		},
		"sub_button_3": {
			_const.RU: `Купить`,
			_const.EN: `Buy`,
		},
		"sub_button_4": {
			_const.RU: `Автопилот`,
			_const.EN: `Autopilot`,
		},
		"sub_button_5": {
			_const.RU: `Проложить маршрут`,
			_const.EN: `Get directions`,
		},
		"sub_button_6": {
			_const.RU: `Закрыть`,
			_const.EN: `Close`,
		},
		"text_4": {
			_const.RU: `Мой баланс:`,
			_const.EN: `My balance:`,
		},
		"tab_1": {
			_const.RU: `Рынок`,
			_const.EN: `Market`,
		},
		"tab_2": {
			_const.RU: `Мои запросы/предложения`,
			_const.EN: `My requests/suggestions`,
		},
		"tab_3": {
			_const.RU: `История сделок`,
			_const.EN: `Transaction history`,
		},
		"button_3": {
			_const.RU: `Купить`,
			_const.EN: `Buy`,
		},
		"dist_1": {
			_const.RU: `Недостижимо`,
			_const.EN: `Unattainable`,
		},
		"dist_2": {
			_const.RU: `База`,
			_const.EN: `Base`,
		},
		"dist_3": {
			_const.RU: `Сектор`,
			_const.EN: `Sector`,
		},
		"head_table_1": {
			_const.RU: `Продавцы`,
			_const.EN: `Sellers`,
		},
		"head_col_table_1": {
			_const.RU: `Расстояние`,
			_const.EN: `Distance`,
		},
		"head_col_table_2": {
			_const.RU: `Кол-во`,
			_const.EN: `Qty`,
		},
		"head_col_table_3": {
			_const.RU: `Цена`,
			_const.EN: `Price`,
		},
		"head_col_table_4": {
			_const.RU: `Тип`,
			_const.EN: `Type`,
		},
		"head_col_table_5": {
			_const.RU: `Название`,
			_const.EN: `Name`,
		},
		"head_col_table_6": {
			_const.RU: `Место`,
			_const.EN: `Place`,
		},
		"head_table_2": {
			_const.RU: `Покупатели`,
			_const.EN: `Buyers`,
		},
		"head_col_table_7": {
			_const.RU: `Расстояние`,
			_const.EN: `Distance`,
		},
		"head_col_table_8": {
			_const.RU: `Кол-во`,
			_const.EN: `Qty`,
		},
		"head_col_table_9": {
			_const.RU: `Цена`,
			_const.EN: `Price`,
		},
		"head_col_table_10": {
			_const.RU: `Тип`,
			_const.EN: `Type`,
		},
		"head_col_table_11": {
			_const.RU: `Название`,
			_const.EN: `Name`,
		},
		"head_col_table_12": {
			_const.RU: `Место`,
			_const.EN: `Place`,
		},
		"head_col_table_13": {
			_const.RU: `Расстояние`,
			_const.EN: `Distance`,
		},
		"head_col_table_14": {
			_const.RU: `Кол-во`,
			_const.EN: `Qty`,
		},
		"head_col_table_15": {
			_const.RU: `Цена`,
			_const.EN: `Price`,
		},
		"head_col_table_16": {
			_const.RU: `Тип`,
			_const.EN: `Type`,
		},
		"head_col_table_17": {
			_const.RU: `Название`,
			_const.EN: `Name`,
		},
		"head_col_table_18": {
			_const.RU: `Место`,
			_const.EN: `Place`,
		},
		"head_col_table_19": {
			_const.RU: `Тип сделки`,
			_const.EN: `Transaction type`,
		},
		"head_col_table_20": {
			_const.RU: `Общая стоимость`,
			_const.EN: `Total cost`,
		},
		"head_col_table_21": {
			_const.RU: `Название`,
			_const.EN: `Name`,
		},
		"head_col_table_22": {
			_const.RU: `Кол-во`,
			_const.EN: `Qty`,
		},
		"head_col_table_23": {
			_const.RU: `Тип сделки`,
			_const.EN: `Transaction type`,
		},
		"head_col_table_24": {
			_const.RU: `Стоимость`,
			_const.EN: `Total cost`,
		},
		"head_col_table_25": {
			_const.RU: `Место`,
			_const.EN: `Place`,
		},
		"deal_t1": {
			_const.RU: `Продажа`,
			_const.EN: `Sell`,
		},
		"deal_t2": {
			_const.RU: `Покупка`,
			_const.EN: `Buy`,
		},
		"text_14": {
			_const.RU: `~ цена: <span class="importantly">%credits_count%</span> cr.`,
			_const.EN: `~ price: <span class="importantly">%credits_count%</span> cr.`,
		},
		// NewBuy
		"text_5": {
			_const.RU: `Кол-во:`,
			_const.EN: `Qty:`,
		},
		"text_6": {
			_const.RU: `Сr. за шт.`,
			_const.EN: `Cr. per pce`,
		},
		"text_7": {
			_const.RU: `Всего:`,
			_const.EN: `Total:`,
		},
		"button_4": {
			_const.RU: `Отменить`,
			_const.EN: `Cancel`,
		},
		"button_5": {
			_const.RU: `Разместить`,
			_const.EN: `Place an order`,
		},
		// Sell
		"text_8": {
			_const.RU: `Кол-во:`,
			_const.EN: `Qty:`,
		},
		"text_9": {
			_const.RU: `Сr. за шт.`,
			_const.EN: `Cr. per pce`,
		},
		"text_10": {
			_const.RU: `Всего:`,
			_const.EN: `Total:`,
		},
		"text_11": {
			_const.RU: `Налог (%tax%%):`,
			_const.EN: `Tax (%tax%%):`,
		},
		"text_12": {
			_const.RU: `Вы получите:`,
			_const.EN: `You'll get:`,
		},
		"button_6": {
			_const.RU: `Отменить`,
			_const.EN: `Cancel`,
		},
		"button_7": {
			_const.RU: `Продать`,
			_const.EN: `Sell`,
		},
		"text_13": {
			_const.RU: `На базе где размещен заказ нет предметов для продажи.`,
			_const.EN: `There are no items for sale at the base where the order was placed.`,
		},
		"button_8": {
			_const.RU: `Назад`,
			_const.EN: `Back`,
		},
	},
	"MenuBar": {
		// MenuBar
		"text_1": {
			_const.RU: `Общение`,
			_const.EN: `Communication`,
		},
		"sub_1": {
			_const.RU: `Личное дело`,
			_const.EN: `Personal information`,
		},
		"sub_2": {
			_const.RU: `Контакты`,
			_const.EN: `Contacts`,
		},
		"sub_3": {
			_const.RU: `Группа`,
			_const.EN: `Group`,
		},
		"sub_4": {
			_const.RU: `Кластер`,
			_const.EN: `Corporation`,
		},
		"text_2": {
			_const.RU: `Помощь`,
			_const.EN: `Help`,
		},
		"sub_5": {
			_const.RU: `Управление`,
			_const.EN: `Control`,
		},
		"sub_6": {
			_const.RU: `Помощь`,
			_const.EN: `Help`,
		},
		"sub_7": {
			_const.RU: `Технологии`,
			_const.EN: `Technologies`,
		},
		"sub_8": {
			_const.RU: `Группа ВК`,
			_const.EN: `VK group`,
		},
		"sub_9": {
			_const.RU: `Discord`,
			_const.EN: `Discord`,
		},
		"text_3": {
			_const.RU: `Автопилот`,
			_const.EN: `Autopilot`,
		},
		"text_4": {
			_const.RU: `Эвакуация`,
			_const.EN: `Evacuation`,
		},
		"text_5": {
			_const.RU: `Топливо`,
			_const.EN: `Fuel`,
		},
		// Notifys
	},
	"MiniGames": {
		"window_name": {
			_const.RU: `Мини-игра`,
			_const.EN: `Mini-game`,
		},
		"text_1": {
			_const.RU: `Победа!`,
			_const.EN: `Win!`,
		},
		"button_1": {
			_const.RU: `Завершить`,
			_const.EN: `Complete`,
		},
	},
	"MiniMap": {
		"zoom": {
			_const.RU: `Зум`,
			_const.EN: `Zoom`,
		},
		"fix_cam": {
			_const.RU: `Фиксация камеры`,
			_const.EN: `Camera fixation`,
		},
		"battle_cam": {
			_const.RU: `Боевой обзор`,
			_const.EN: `Combat review`,
		},
	},
	"Missions": {
		// Missions
		"window_name": {
			_const.RU: `Активные задания`,
			_const.EN: `Active tasks`,
		},
		"button_1": {
			_const.RU: `задания`,
			_const.EN: `missions`,
		},
		// Points
		"text_1": {
			_const.RU: `Обновлено`,
			_const.EN: `Updated`,
		},
		"button_2": {
			_const.RU: `помощь`,
			_const.EN: `help`,
		},
		"button_3": {
			_const.RU: `потерял предмет?`,
			_const.EN: `lost an item?`,
		},
	},
	"ModalNotify": {
		"window_name": {
			_const.RU: `Наблюдение`,
			_const.EN: `Observation`,
		},
	},
	"News": {
		// News
		"window_name": {
			_const.RU: `Новости`,
			_const.EN: `News`,
		},
		"tab_1": {
			_const.RU: `Новости`,
			_const.EN: `News`,
		},
		"tab_2": {
			_const.RU: `Поиск`,
			_const.EN: `Search`,
		},
		"n_tab_1": {
			_const.RU: `Все`,
			_const.EN: `Everything`,
		},
		"n_tab_2": {
			_const.RU: `Торговля`,
			_const.EN: `Trade`,
		},
		"n_tab_3": {
			_const.RU: `Война`,
			_const.EN: `War`,
		},
		"n_tab_4": {
			_const.RU: `Экспедиция`,
			_const.EN: `Expedition`,
		},
		"n_tab_5": {
			_const.RU: `Аномалии`,
			_const.EN: `Anomalies`,
		},
		"n_tab_6": {
			_const.RU: `Кластеры`,
			_const.EN: `Clusters`,
		},
		"search_t_1": {
			_const.RU: `Персона`,
			_const.EN: `Persona`,
		},
		"search_t_2": {
			_const.RU: `База`,
			_const.EN: `Base`,
		},
		"search_t_3": {
			_const.RU: `Товары`,
			_const.EN: `Product`,
		},
		"button_1": {
			_const.RU: `Найти (0.3 cr.)`,
			_const.EN: `Find (0.3 cr.)`,
		},
		"base_t_1": {
			_const.RU: `Перерабатывающий завод`,
			_const.EN: `Processing plant`,
		},
		"base_t_2": {
			_const.RU: `Основная база`,
			_const.EN: `Main base`,
		},
		"base_t_3": {
			_const.RU: `Завод расходных материалов`,
			_const.EN: `Factory`,
		},
		"range_1": {
			_const.RU: `Налог до`,
			_const.EN: `Tax up to`,
		},
		"range_2": {
			_const.RU: `Расстояние до`,
			_const.EN: `Distance to`,
		},
		"text_1": {
			_const.RU: `Кол-во:`,
			_const.EN: `Qty:`,
		},
		"text_2": {
			_const.RU: `Макс. цена продажи:`,
			_const.EN: `Max selling price:`,
		},
		"text_3": {
			_const.RU: `Мин. цена покупки:`,
			_const.EN: `Min. purchase price:`,
		},
		"nothing": {
			_const.RU: `Ничего не найдено`,
			_const.EN: `Nothing found`,
		},
		"target_1": {
			_const.RU: `Регион:`,
			_const.EN: `Region:`,
		},
		"target_2": {
			_const.RU: `Сектор:`,
			_const.EN: `Sector:`,
		},
		"target_3": {
			_const.RU: `База:`,
			_const.EN: `Base:`,
		},
		"button_2": {
			_const.RU: `карта`,
			_const.EN: `map`,
		},
		"text_4": {
			_const.RU: `Налог:`,
			_const.EN: `Tax:`,
		},
		"text_5": {
			_const.RU: `расстояние:`,
			_const.EN: `distance:`,
		},
		"text_6": {
			_const.RU: `Расстояние:`,
			_const.EN: `Distance:`,
		},
		"text_7": {
			_const.RU: `кол-во:`,
			_const.EN: `qty:`,
		},
		"text_8": {
			_const.RU: `Покупка:`,
			_const.EN: `Purchase:`,
		},
		"text_9": {
			_const.RU: `Продажа:`,
			_const.EN: `Sale:`,
		},
		// Notify
		"text_10": {
			_const.RU: `Завершено производство`,
			_const.EN: `Production completed`,
		},
		"on_base": {
			_const.RU: `На базе`,
			_const.EN: `On the base`,
		},
		"sector": {
			_const.RU: `сектор:`,
			_const.EN: `sector:`,
		},
		"text_11": {
			_const.RU: `завершено производство:`,
			_const.EN: `production completed:`,
		},
		"text_12": {
			_const.RU: `Созданные вещи ожидают вас на складе базы.`,
			_const.EN: `The created items are waiting for you in the base warehouse.`,
		},
		"text_13": {
			_const.RU: `Совершена продажа`,
			_const.EN: `Sale completed`,
		},
		"text_14": {
			_const.RU: `Совершена покупка`,
			_const.EN: `Purchase completed`,
		},
		"text_15": {
			_const.RU: `было продано:`,
			_const.EN: `sold:`,
		},
		"text_16": {
			_const.RU: `было куплено:`,
			_const.EN: `was purchased:`,
		},
		"text_17": {
			_const.RU: `Вы заработали:`,
			_const.EN: `You've earned:`,
		},
		"text_18": {
			_const.RU: `Купленные вещи ожидают вас на складе базы.`,
			_const.EN: `The purchased items are waiting for you at the base warehouse.`,
		},
		"dialog": {
			_const.RU: `Диалог`,
			_const.EN: `Dialog`,
		},
		"from": {
			_const.RU: `От`,
			_const.EN: `From`,
		},
		"to": {
			_const.RU: `Кому`,
			_const.EN: `To`,
		},
		"attention": {
			_const.RU: `Внимание!`,
			_const.EN: `Attention!`,
		},
		"request": {
			_const.RU: `Запрос`,
			_const.EN: `Request`,
		},
		"transfer": {
			_const.RU: `Перевод`,
			_const.EN: `Transfer`,
		},
		"relation_text": {
			_const.RU: `Отношения с фракцией <span class="importantly">%fraction%</span> улучшены.`,
			_const.EN: `Relations with the <span class="importantly">%fraction%</span> faction deteriorated.`,
		},
		"close": {
			_const.RU: `Закрыть`,
			_const.EN: `Close`,
		},
	},
	"NPCTrade": {
		"window_name": {
			_const.RU: `Торговля`,
			_const.EN: `Commerce`,
		},
		"error_1": {
			_const.RU: `Не удалось`,
			_const.EN: `Failed`,
		},
		"button_1": {
			_const.RU: `Отмена`,
			_const.EN: `Cancel`,
		},
		"text_1": {
			_const.RU: `плохое отношение`,
			_const.EN: `bad attitude`,
		},
		"qty": {
			_const.RU: `Кол-во`,
			_const.EN: `Qty`,
		},
		"price": {
			_const.RU: `Цена`,
			_const.EN: `Price`,
		},
		"text_2": {
			_const.RU: `Синтет`,
			_const.EN: `Synthet`,
		},
		"sell": {
			_const.RU: `Продать`,
			_const.EN: `Sell`,
		},
		"buy": {
			_const.RU: `Купить`,
			_const.EN: `Buy`,
		},
	},
	"ObjectDialog": {
		"window_name": {
			_const.RU: `Объект: `,
			_const.EN: `An object: `,
		},
		"owner": {
			_const.RU: `Владелец:`,
			_const.EN: `Owner: `,
		},
		"hp": {
			_const.RU: `HP:`,
			_const.EN: `HP: `,
		},
		"energy": {
			_const.RU: `Energy:`,
			_const.EN: `Energy: `,
		},
	},
	"PrefabricatedMenu": {
		"window_name": {
			_const.RU: `Производство деталей`,
			_const.EN: `Parts production`,
		},
		"error_1": {
			_const.RU: `Модуль сервиса уничтожен или отключен.`,
			_const.EN: `The service module has been destroyed or disabled.`,
		},
		"not_allow": {
			_const.RU: `Мы не доверяем вам и поэтому этот сервис не доступен.`,
			_const.EN: `We do not trust you and therefore this service is not available.`,
		},
		"to_storage": {
			_const.RU: `Предметы отправлены на <span class="importantly">склад</span>`,
			_const.EN: `Items sent to <span class="importantly">warehouse</span>`,
		},
		"detail_select": {
			_const.RU: `Выбери деталь`,
			_const.EN: `Choose a part`,
		},
		"tax": {
			_const.RU: `Налог:`,
			_const.EN: `Tax:`,
		},
		"button_1": {
			_const.RU: `Произвести`,
			_const.EN: `Produce`,
		},
		"need_resources": {
			_const.RU: `Необходимо ресурсов`,
			_const.EN: `Resources needed`,
		},
	},
	"Preloader": {},
	"ProcessorRoot": {
		// ItemsDropZone
		"head": {
			_const.RU: `Переработать:`,
			_const.EN: `Recycler:`,
		},
		"text_1": {
			_const.RU: `Перетащите сюда предметы которые хотите переработать.`,
			_const.EN: `Drag the items you want to recycle here.`,
		},
		// ProcessorRoot
		"window_name": {
			_const.RU: `Переработчик`,
			_const.EN: `Processor`,
		},
		"to_storage": {
			_const.RU: `Предметы отправлены на <span class="importantly">склад</span>`,
			_const.EN: `Items sent to <span class="importantly">warehouse</span>`,
		},
		"text_2": {
			_const.RU: `Модуль сервиса уничтожен или отключен.`,
			_const.EN: `The service module has been destroyed or disabled.`,
		},
		"not_allow": {
			_const.RU: `Мы не доверяем вам и поэтому этот сервис не доступен.`,
			_const.EN: `We do not trust you and therefore this service is not available.`,
		},
		"lost_items": {
			_const.RU: `Потери:`,
			_const.EN: `Losses:`,
		},
		"tax": {
			_const.RU: `Налог:`,
			_const.EN: `Tax:`,
		},
		"result": {
			_const.RU: `Результат:`,
			_const.EN: `Result:`,
		},
		"button_1": {
			_const.RU: `Закрыть`,
			_const.EN: `Close`,
		},
		"button_2": {
			_const.RU: `Переработать`,
			_const.EN: `Recycle`,
		},
		"text_3": {
			_const.RU: `Переработать`,
			_const.EN: `Recycle`,
		},
	},
	"SelectTarget": {
		"title_1": {
			_const.RU: `диалог`,
			_const.EN: `dialog`,
		},
		"chanel_unknown": {
			_const.RU: `Неизвестный`,
			_const.EN: `Unknown`,
		},
	},
	"SendCredits": {
		"window_name": {
			_const.RU: `Кредиты`,
			_const.EN: `Credits`,
		},
		"transfer": {
			_const.RU: `Перевод`,
			_const.EN: `Transfer`,
		},
		"history": {
			_const.RU: `История`,
			_const.EN: `History`,
		},
		"text_1": {
			_const.RU: `Откуда:`,
			_const.EN: `Src:`,
		},
		"text_2": {
			_const.RU: `Куда:`,
			_const.EN: `Dst:`,
		},
		"option_1": {
			_const.RU: `Мой счет`,
			_const.EN: `My deposit`,
		},
		"player": {
			_const.RU: `Игрок`,
			_const.EN: `Player`,
		},
		"cluster": {
			_const.RU: `Кластер`,
			_const.EN: `Cluster`,
		},
		"placeholder_1": {
			_const.RU: `Имя`,
			_const.EN: `Name`,
		},
		"button_1": {
			_const.RU: `Найти`,
			_const.EN: `Search`,
		},
		"nothing": {
			_const.RU: `Ничего не найдено.`,
			_const.EN: `Nothing found.`,
		},
		"text_3": {
			_const.RU: `Сколько:`,
			_const.EN: `How many:`,
		},
		"placeholder_2": {
			_const.RU: `причина (необяз.)`,
			_const.EN: `reason (optional)`,
		},
		"button_2": {
			_const.RU: `Перевести`,
			_const.EN: `Send`,
		},
		"text_4": {
			_const.RU: `Операция`,
			_const.EN: `Operation`,
		},
		"text_5": {
			_const.RU: `Причина`,
			_const.EN: `Cause`,
		},
		"on_1": {
			_const.RU: `Перевод от кластера <span class="importantly">%name%</span>`,
			_const.EN: `Transfer from cluster <span class="importantly">%name%</span>`,
		},
		"on_2": {
			_const.RU: `Перевод кластеру <span class="importantly">%name%</span>`,
			_const.EN: `Transfer to cluster <span class="importantly">%name%</span>`,
		},
		"on_3": {
			_const.RU: `Перевод от <span class="importantly">%name%</span>`,
			_const.EN: `Translation from <span class="importantly">%name%</span>`,
		},
		"on_4": {
			_const.RU: `Перевод <span class="importantly">%name%</span>`,
			_const.EN: `Translation <span class="importantly">%name%</span>`,
		},
		"on_5": {
			_const.RU: `Налог с продажи`,
			_const.EN: `Sales tax`,
		},
		"on_6": {
			_const.RU: `Стартовый капитал`,
			_const.EN: `Start-up capital`,
		},
		"on_7": {
			_const.RU: `Продажа товара`,
			_const.EN: `Sale of goods`,
		},
		"on_8": {
			_const.RU: `Покупка товара`,
			_const.EN: `Purchase of goods`,
		},
		"on_9": {
			_const.RU: `Продажа на рынке`,
			_const.EN: `Selling on the market`,
		},
		"on_10": {
			_const.RU: `Покупка на рынке`,
			_const.EN: `Purchase at the market`,
		},
		"on_11": {
			_const.RU: `Размещение заказа`,
			_const.EN: `Placing an order`,
		},
		"on_12": {
			_const.RU: `Отмена заказа`,
			_const.EN: `Cancellations`,
		},
		"on_13": {
			_const.RU: `Выполнение миссии`,
			_const.EN: `Completing the mission`,
		},
		"on_14": {
			_const.RU: `Покупка покраски`,
			_const.EN: `Buying skin`,
		},
		"on_15": {
			_const.RU: `Грабеж`,
			_const.EN: `Robbery`,
		},
		"on_16": {
			_const.RU: `Починка транспорта`,
			_const.EN: `Vehicle repair`,
		},
		"on_17": {
			_const.RU: `Починка предмета`,
			_const.EN: `Repairing an item`,
		},
		"on_18": {
			_const.RU: `Покупка в фр. магазине`,
			_const.EN: `Purchasing from a faction store`,
		},
		"on_19": {
			_const.RU: `Поиск`,
			_const.EN: `Search`,
		},
		"on_20": {
			_const.RU: `Устранение преступника`,
			_const.EN: `Elimination of the criminal`,
		},
		"on_21": {
			_const.RU: `Устранение бойца в войне`,
			_const.EN: `Eliminating a fighter in a war`,
		},
		"on_22": {
			_const.RU: `Устранение сил APD`,
			_const.EN: `Elimination of APD forces`,
		},
		"on_23": {
			_const.RU: `Покупка боеприпасов`,
			_const.EN: `Buying ammunition`,
		},
		"on_24": {
			_const.RU: `Заправка транспорта`,
			_const.EN: `Refueling vehicles`,
		},
	},
	"ServerState": {},
	"Settings": {
		"window_name": {
			_const.RU: `Настройки`,
			_const.EN: `Settings`,
		},
		"sound": {
			_const.RU: `Звук`,
			_const.EN: `Sound`,
		},
		"music_volume": {
			_const.RU: `Громкость музыки:`,
			_const.EN: `Music volume:`,
		},
		"sfx_Volume": {
			_const.RU: `Громкость эффектов:`,
			_const.EN: `Effects volume:`,
		},
		"graphics": {
			_const.RU: `Графика`,
			_const.EN: `Graphics`,
		},
		"traces": {
			_const.RU: `Следы от гусениц:`,
			_const.EN: `Traces from transport:`,
		},
	},
	"StatusBar": {},
	"TechnologyMenu": {
		"window_name": {
			_const.RU: `Дерево технологий`,
			_const.EN: `Wood technologies`,
		},
		"tab_1": {
			_const.RU: `Корпуса`,
			_const.EN: `Bodies`,
		},
		"blu": {
			_const.RU: `Легкие универсалы`,
			_const.EN: `Lightweight universal`,
		},
		"bmm": {
			_const.RU: `Средние шахтеры`,
			_const.EN: `Medium miners`,
		},
		"bhm": {
			_const.RU: `Тяжелые шахтеры`,
			_const.EN: `Heavy miners`,
		},
		"blf": {
			_const.RU: `Легкие бойцы`,
			_const.EN: `Light fighters`,
		},
		"bmf": {
			_const.RU: `Средние бойцы`,
			_const.EN: `Medium fighters`,
		},
		"bhf": {
			_const.RU: `Тяжелые бойцы`,
			_const.EN: `Heavy fighters`,
		},
		"bmt": {
			_const.RU: `Средние грузовики`,
			_const.EN: `Medium trucks`,
		},
		"bht": {
			_const.RU: `Тяжелые грузовики`,
			_const.EN: `Heavy trucks`,
		},
		"bl": {
			_const.RU: `Легкие`,
			_const.EN: `Light`,
		},
		"bm": {
			_const.RU: `Средние`,
			_const.EN: `Medium`,
		},
		"bh": {
			_const.RU: `Тяжелые`,
			_const.EN: `Heavy`,
		},
	},
	"Training": {},
	"Transfer": {},
	"UserStat": {
		// Missions
		"no_missions": {
			_const.RU: `Нет заданий`,
			_const.EN: `No missions`,
		},
		"track": {
			_const.RU: `отслеживать:`,
			_const.EN: `track:`,
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
			_const.RU: `Синтет`,
			_const.EN: `Synthet`,
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
	},
	"Window": {},
	"Workbench": {
		"window_name": {
			_const.RU: `Производство`,
			_const.EN: `Production`,
		},
		"error_1": {
			_const.RU: `Модуль сервиса уничтожен или отключен.`,
			_const.EN: `The service module has been destroyed or disabled.`,
		},
		"error_2": {
			_const.RU: `Чертеж должен быть на "Складе"`,
			_const.EN: `The drawing must be in the "Warehouse"`,
		},
		"error_3": {
			_const.RU: `Нехватает ресурсов`,
			_const.EN: `Not enough resources`,
		},
		"not_allow": {
			_const.RU: `Мы не доверяем вам и поэтому этот сервис не доступен.`,
			_const.EN: `We do not trust you and therefore this service is not available.`,
		},
		"queue": {
			_const.RU: `Очередь производства:`,
			_const.EN: `Production queue:`,
		},
		"qty": {
			_const.RU: `Кол-во:`,
			_const.EN: `Qty:`,
		},
		"create": {
			_const.RU: `Создать`,
			_const.EN: `Create`,
		},
		"cancel": {
			_const.RU: `Отменить`,
			_const.EN: `Cancel`,
		},
		"close": {
			_const.RU: `Закрыть`,
			_const.EN: `Close`,
		},
	},
}
