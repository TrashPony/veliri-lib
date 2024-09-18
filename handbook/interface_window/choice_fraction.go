package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var ChoiceFraction = map[string]map[string]string{
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
	"Replics_weaponText": {
		_const.RU: `<span class="importantly">Replics</span> предпочитают баллистическое орудия. Это оружие обладает высокую дальность поражения и вариативность боеприпасов.`,
		_const.EN: `<span class="importantly">Replics</span> prefer ballistic weapons. These weapons have a long range and offer a variety of ammunition options.`,
	},
	"Replics_bodyText": {
		_const.RU: `<span class="importantly">Replics</span> используют тяжелые транспорты нагусеничном шасси. <span class="importantly">Гусеницы</span> позволяют крутиться на месте и не имеют штрафапри движении назад.`,
		_const.EN: `<span class="importantly">Replics</span> use heavy vehicles on a tracked chassis. The tracks allow them to spin in place and there is no movement penalty when moving backward.`,
	},
	"Replics_bonus": {
		_const.RU: `<span class="importantly">Защита</span> от всего урона повышена на <span class="importantly">5%</span>.`,
		_const.EN: `The <span class="importantly">protection</span> from all damage has been increased by <span class="importantly">5%</span>.`,
	},
	"Explores_weaponText": {
		_const.RU: `<span class="importantly">Explores</span> предпочитают лазерные орудия. Это оружие обладает высокою точностью и способны мгновенно поражать цель, однако их дальность оставляет желать лучшего.`,
		_const.EN: `<span class="importantly">Explores</span> prefer laser weapons. These weapons have high accuracy and can instantly hit the target, but their range leaves much to be desired.`,
	},
	"Explores_bodyText": {
		_const.RU: `<span class="importantly">Explores</span> используют легкие транспорты на антиграве. <span class="importantly">Антиграв</span> дает вам большую свободу и гибкость. Оно позволяет вам управлять направлением корпуса мышкой и двигаться боком без поворота корпуса.`,
		_const.EN: `<span class="importantly">Explores</span> use light vehicles with antigravity. The antigravity gives you great freedom and flexibility. It allows you to control the direction of the hull with the mouse and move sideways without turning the hull.`,
	},
	"Explores_bonus": {
		_const.RU: `<span class="importantly">Дальность обзора</span> увеличена на <span class="importantly">5%</span>.`,
		_const.EN: `The <span class="importantly">view range</span> has been increased by <span class="importantly">5%</span>.`,
	},
	"Reverses_weaponText": {
		_const.RU: `<span class="importantly">Reverses</span> предпочитают ракетные пусковые установки. Ракеты обладают высокой поражающей способностью, однако существует снаряжение способное их сбивать.`,
		_const.EN: `The <span class="importantly">Reverses</span> prefer rocket launchers. Rockets have a high destructive capacity, but there is equipment that can shoot them down.`,
	},
	"Reverses_bodyText": {
		_const.RU: `<span class="importantly">Reverses</span> используют транспорты на колесах. <span class="importantly">Колеса</span> не имеет особых особенностей. Оно является самым простым и надежным типом шасси. Оно подходит для тех, кто предпочитает классический стиль управления.`,
		_const.EN: `The <span class="importantly">Reverses</span> use wheeled vehicles. The wheels do not have any special features. They are the simplest and most reliable type of chassis. They suit those who prefer a classic style of control.`,
	},
	"Reverses_bonus": {
		_const.RU: `<span class="importantly">Вместимость трюма</span> увеличена на <span class="importantly">5%</span>.`,
		_const.EN: `The <span class="importantly">hold capacity</span> has been increased by <span class="importantly">5%</span>.`,
	},
	"button_1": {
		_const.RU: `ВПЕРЕД >`,
		_const.EN: `CHOOSE >`,
	},
	"text_1": {
		_const.RU: `Выбери сторону`,
		_const.EN: `Choose a side.`,
	},
}
