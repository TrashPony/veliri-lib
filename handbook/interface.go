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
		// CorporationLine
		// Create
		// Main
		// Members
		// Policy
		// Requests
		// Roles
		// Search
		// Wars
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
