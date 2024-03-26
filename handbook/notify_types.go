package handbook

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"math/rand"
	"strings"
	"time"
)

var dialogTypes = map[string]map[string]string{
	// HELP TASK
	"help_request_1": {
		_const.RU: "Эй, %TO_USER_NAME% мне бы пригодился проводник в секторе. Вдвоём-то куда веселее, что об этом думаешь?",
		_const.EN: "english translate",
	},
	"help_request_2": {
		_const.RU: "Что думаешь касательного того, чтобы путешествовать вместе, а %TO_USER_NAME%?",
		_const.EN: "english translate",
	},
	"help_request_3": {
		_const.RU: "Я исследователь и мне бы не помешал телохранитель. Понимаешь, к чему я клоню %TO_USER_NAME%?",
		_const.EN: "english translate",
	},
	"help_busy_1_1": {
		_const.RU: "Да ты хоть знаешь, к кому обращаешься!? А... Мне ведь и говорить об этом нельзя. Забудь.",
		_const.EN: "english translate",
	},
	"help_busy_1_2": {
		_const.RU: "Я на задании. Оно важно и ответственно. Кстати, ты не нанесёшь мне на карту расположение ближайших военных баз?",
		_const.EN: "english translate",
	},
	"help_busy_1_3": {
		_const.RU: "У меня вообще-то важная миссия. Но тебя это не касается. Скажем... здесь столкнулись разные интересы.",
		_const.EN: "english translate",
	},
	"help_busy_2_1": {
		_const.RU: "Как бы я занят - выполняю важную миссию. А вдруг здесь где поблизости боты-шпионы? А вдруг ты один из них?",
		_const.EN: "english translate",
	},
	"help_busy_2_2": {
		_const.RU: "Неа, тут уж ты сам. Мне оно ни к чему.",
		_const.EN: "english translate",
	},
	"help_busy_2_3": {
		_const.RU: "Мне это неинтересно. Без обид.",
		_const.EN: "english translate",
	},
	"help_busy_2_4": {
		_const.RU: "Ха! Ищи дурака! А вдруг ты просто желаешь меня ограбить? Или... ещё чего другого? Так я на это и повёлся.",
		_const.EN: "english translate",
	},
	"help_busy_3_": {
		_const.RU: "Э-э, что-то ты выглядишь слишком подозрительным! Ну уж нет.",
		_const.EN: "english translate",
	},
	"help_busy_4_1": {
		_const.RU: "Ха! Ищи дурака! А вдруг ты просто желаешь меня ограбить? Или... ещё чего другого? Так я на это и повёлся.",
		_const.EN: "english translate",
	},
	"help_busy_4_2": {
		_const.RU: "Нет уж. Здесь каждый сам за себя! Без обид %FROM_USER_NAME%, но здесь наши дороги расходятся. Надеюсь, что без последствий.",
		_const.EN: "english translate",
	},
	"help_busy_4_3": {
		_const.RU: "Ты плохо врёшь! Это же фирменная ловушка! Лучше пойди поищи дураков в другом месте.",
		_const.EN: "english translate",
	},
	"help_busy_5_1": {
		_const.RU: "Неа, тут уж ты сам. Мне оно ни к чему.",
		_const.EN: "english translate",
	},
	"help_busy_5_2": {
		_const.RU: "Мне это неинтересно. Без обид.",
		_const.EN: "english translate",
	},
	"help_busy_5_3": {
		_const.RU: "Ты плохо врёшь! Это же фирменная ловушка! Лучше пойди поищи дураков в другом месте.",
		_const.EN: "english translate",
	},
	"help_busy_5_4": {
		_const.RU: "Ха! Ищи дурака! А вдруг ты просто желаешь меня ограбить? Или... ещё чего другого? Так я на это и повёлся.",
		_const.EN: "english translate",
	},
	"help_busy_6_1": {
		_const.RU: "Моя миссия превыше этого.",
		_const.EN: "english translate",
	},
	"help_busy_6_2": {
		_const.RU: "Вы не член здешней главенствующей фракции. В запросе отказано.",
		_const.EN: "english translate",
	},
	"help_busy_6_3": {
		_const.RU: "На вас не распространяются условия о защите и сопровождении. Ради собственного блага - отъедьте от меня на некоторое расстояние.",
		_const.EN: "english translate",
	},
	"help_busy_7_1": {
		_const.RU: "<span class=\"importantly\">Ошибка 415.</span>",
		_const.EN: "english translate",
	},
	"help_busy_7_2": {
		_const.RU: "<span class=\"importantly\">Сигнал неопознан.</span>",
		_const.EN: "english translate",
	},
	"help_busy_7_3": {
		_const.RU: "<span class=\"importantly\">333./444-55,16?</span>",
		_const.EN: "english translate",
	},
	"help_busy_8_": {
		_const.RU: "Ещё чего! Я не стану себя подвергать угрозе.",
		_const.EN: "english translate",
	},
	"help_busy_9_1": {
		_const.RU: "Ты что, не видишь, чем я занят? Мне ещё предстоит отчитаться, сколько я добыл. А если вдруг выяснится, что пары граммов не хватает? Знаешь, что со мной сделают?",
		_const.EN: "english translate",
	},
	"help_busy_9_2": {
		_const.RU: "Мне некогда. Я столько времени искал богатое на месторождения место, чтобы теперь всё бросить и отправиться с тобой.",
		_const.EN: "english translate",
	},
	"help_busy_9_3": {
		_const.RU: "Руда себя не выкопает. А я, как раз собирался немного подзаработать. Извини, тут уж ты сам по себе.",
		_const.EN: "english translate",
	},
	"help_busy_9_4": {
		_const.RU: "Э! Да я только приобрёл это горнодобывающее оборудование! Ты серьёзно рассчитываешь, что я брошу попытку опробовать его только ради тебя? Ха! Ха-ха!",
		_const.EN: "english translate",
	},
	"help_busy_9_5": {
		_const.RU: "А может напротив - ты мне составишь компанию в добыче полезных ресурсов? Нет... Эх. Опять всё придётся делать самому. Ну, бывай тогда.",
		_const.EN: "english translate",
	},
	"help_busy_9_6": {
		_const.RU: "У меня контракт на разведку и добычу ресурсов. Как освобожусь через три - четыре месяца, я тебе дам знать.",
		_const.EN: "english translate",
	},
	"help_busy_9_7": {
		_const.RU: "Не, не, не! Совсем нету на это времени. Я сейчас направляюсь к одному месторождению, а ты, пытаешься отвлечь меня и, чтобы пока я отсутствую, оттуда уже всё выгребли своими грязными ковшами? Такому не бывать!",
		_const.EN: "english translate",
	},
	"help_busy_10_1": {
		_const.RU: "Это всё конечно замечательно, но я выполняю важную миссию. Скажу тебе по секрету - мне нужно разведать местность.",
		_const.EN: "english translate",
	},
	"help_busy_10_2": {
		_const.RU: "Я следую к сигналу бедствия. Сейчас не до тебя.",
		_const.EN: "english translate",
	},
	"help_busy_10_3": {
		_const.RU: "Я двигаюсь по распоряжению фракции к одному источнику сигнала. Поэтому, наши пути не могут совпадать.",
		_const.EN: "english translate",
	},
	"help_busy_10_4": {
		_const.RU: "Здесь неподалёку идёт трансляция сигнала бедствия. Я следую туда по спасательной миссии.",
		_const.EN: "english translate",
	},
	"help_busy_11_4": {
		_const.RU: "Ты должно быть шутишь?",
		_const.EN: "english translate",
	},
	"help_busy_12_1": {
		_const.RU: "Сейчас я участвую в бою. Обратитесь позже!",
		_const.EN: "english translate",
	},
	"help_complete_1": {
		_const.RU: "Хм. Ладно, но я внимательно за тобой слежу. Впереди, кстати, будешь ты!",
		_const.EN: "english translate",
	},
	"help_complete_2": {
		_const.RU: "Я согласен. Дорога трудна, а друг-Синтет - в пустошах редкое явление.",
		_const.EN: "english translate",
	},
	"help_complete_3": {
		_const.RU: "Здравая идея! Ладно. Сработаемся!",
		_const.EN: "english translate",
	},
	"help_complete_4": {
		_const.RU: "Добро! За дело!",
		_const.EN: "english translate",
	},
	"leave_help_dialog_1": {
		_const.RU: "У меня есть дела поважнее, чем таскаться с тобой по пустошам.",
		_const.EN: "english translate",
	},
	"leave_help_dialog_2": {
		_const.RU: "А вообще, знаешь, мои планы изменились. Забудь. Разве что только, ты не желаешь.меня как-нибудь отблагодарить.",
		_const.EN: "english translate",
	},
	"leave_help_not_friend_dialog_1": {
		_const.RU: "Мне стало известно, о том кто ты и чем промышляешь. Зря я пожалуй снизошёл до диалога с тобой. Прощай.",
		_const.EN: "english translate",
	},
	"leave_help_not_friend_dialog_2": {
		_const.RU: "Я лучше найду кого-нибудь другого. Есть мнение, что наше знакомство до добра не доведёт.",
		_const.EN: "english translate",
	},

	// MAKE_FRIEND
	"make_friends_request_1": {
		_const.RU: "Пора заканчивать эту глупую конфронтацию, %TO_USER_NAME%! Хватит уже!",
		_const.EN: "english translate",
	},
	"make_friends_request_2": {
		_const.RU: "Я бы желал свести на нет любые наши разногласия %TO_USER_NAME%.",
		_const.EN: "english translate",
	},
	"make_friends_request_3": {
		_const.RU: "Всё! Довольно! Я не желаю быть тебе врагом %TO_USER_NAME%.",
		_const.EN: "english translate",
	},
	"make_friends_request_4": {
		_const.RU: "Всё в этом мире подходит к концу %TO_USER_NAME%. Рано или поздно. Так, не пора ли и нам самим закончить любые наши междоусобные претензии?",
		_const.EN: "english translate",
	},
	"make_friends_request_5": {
		_const.RU: "Мне всё это формально надоело %TO_USER_NAME%! Предлагаю только сейчас - прекратить весь этот фарс и раазойтись друзьями.",
		_const.EN: "english translate",
	},
	"make_friends_no_1_1": {
		_const.RU: "<span class=\"importantly\">Неизвестно. Сбой сообщения.</span>",
		_const.EN: "english translate",
	},
	"make_friends_no_1_2": {
		_const.RU: "Tty?'\\WgHANT//.. <br><br> <span class=\"importantly\">Не удалось распознать ответ.</span>",
		_const.EN: "english translate",
	},
	"make_friends_no_2_1": {
		_const.RU: "Подобное не может быть исполнено. Вы вынуждаете нарушать правила кодекса.",
		_const.EN: "english translate",
	},
	"make_friends_no_2_2": {
		_const.RU: "Отказано! Любые попытки последующих контактов приведут к усилению огневого воздействия.",
		_const.EN: "english translate",
	},
	"make_friends_no_2_3": {
		_const.RU: "Ваш сигнал нераспознан. Повторите попытку вновь. Последующие неудачи в связи, могут привести к санкциям.",
		_const.EN: "english translate",
	},
	"make_friends_no_3": {
		_const.RU: "А у нас что, с тобой были раньше какие-то разногласия?",
		_const.EN: "english translate",
	},
	"make_friends_no_4_1": {
		_const.RU: "Бейся до конца! Покажи из чего ты сделан!",
		_const.EN: "english translate",
	},
	"make_friends_no_4_2": {
		_const.RU: "Ты от меня так просто не отделаешся %FROM_USER_NAME%!",
		_const.EN: "english translate",
	},
	"make_friends_no_4_3": {
		_const.RU: "Да у тебя должно быть микросхемы от радиации расплавились. Даже и не мечтай о подобном.",
		_const.EN: "english translate",
	},
	"make_friends_no_4_4": {
		_const.RU: "Гляди чего он вздумал. И не мечтай %FROM_USER_NAME% о чём-то подобном.",
		_const.EN: "english translate",
	},
	"make_friends_no_5_1": {
		_const.RU: "Ты жалок. А твои попытки примирения только ещё больше подогревают мой гнев!",
		_const.EN: "english translate",
	},
	"make_friends_no_5_2": {
		_const.RU: "Ха! Что это я вижу - попытку примирения? Сейчас-то я тебя научу манерам %FROM_USER_NAME%.",
		_const.EN: "english translate",
	},
	"make_friends_no_5_3": {
		_const.RU: "Хорошая попытка. Жаль только, что неудачная.",
		_const.EN: "english translate",
	},
	"make_friends_pay_1": {
		_const.RU: "Финансы, способы скрасить многие вещи %FROM_USER_NAME%. Они очень хорошо решают проблемы. Особенно, когда их платят за что-то. А ещё лучше, если они в размере - %CREDITS_COUNT% cr.",
		_const.EN: "english translate",
	},
	"make_friends_pay_2": {
		_const.RU: "Хочешь жить, %FROM_USER_NAME%? Тогда придётся раскошелится в размере %CREDITS_COUNT% cr.",
		_const.EN: "english translate",
	},
	"make_friends_pay_3": {
		_const.RU: "Сложившаяся ситуация, толкает меня %FROM_USER_NAME%, требовать от тебя сумму в - %CREDITS_COUNT% cr. Это ради твоего же блага. Ты, главное пойми.",
		_const.EN: "english translate",
	},
	"make_friends_pay_4": {
		_const.RU: "Значит так, %FROM_USER_NAME% твоя ситуация непроста, а выбор невелик - гони сумму в %CREDITS_COUNT% cr. иначе не доползёшь и ближайшей базы. Времени на раздумье у тебя нету.",
		_const.EN: "english translate",
	},
	"make_friends_pay_5": {
		_const.RU: "Вот ты и попался мне, %FROM_USER_NAME%. Что же, придётся отнять у тебя некоторую сумму в %CREDITS_COUNT%. Видишь, как она тебя замедляет и утяжеляет. Позволить облегчить тебе эту тяжкую ношу.",
		_const.EN: "english translate",
	},
	"make_friends_complete_1": {
		_const.RU: "Ты прав - ты прав. Это не имеет никакого логического смысла.",
		_const.EN: "english translate",
	},
	"make_friends_complete_2": {
		_const.RU: "Хорошо. Но только в этот раз. В первый и последний!",
		_const.EN: "english translate",
	},
	"make_friends_complete_3": {
		_const.RU: "Убедил, %FROM_USER_NAME%. Мне и самому нет дела до глупых обид прошлого. Что было - то прошло.",
		_const.EN: "english translate",
	},
	"make_friends_complete_4": {
		_const.RU: "Знаешь, что? А почему бы и нет!",
		_const.EN: "english translate",
	},
	"make_friends_complete_5": {
		_const.RU: "На такое предложение, я отвечу положительным вердиктом.",
		_const.EN: "english translate",
	},
	"make_friends_do_pay_1": {
		_const.RU: "Да... Да! Я заплачу. Пусть будет так.",
		_const.EN: "english translate",
	},
	"make_friends_do_pay_2": {
		_const.RU: "Без вопросов. Меня не стоит убеждать дважды.",
		_const.EN: "english translate",
	},
	"make_friends_do_pay_3": {
		_const.RU: "Эх. Очередные траты. Опять траты... Я согласен. Куда уж мне деваться.",
		_const.EN: "english translate",
	},
	"make_friends_do_pay_4": {
		_const.RU: "Ну, раз уж меня нету иного выхода, то... что теперь поделаешь. Держи, свои %CREDITS_COUNT% cr.",
		_const.EN: "english translate",
	},
	"make_friends_do_pay_5": {
		_const.RU: "Ясно, что с тобой шутки плохи. Перечисляю %CREDITS_COUNT%. cr.",
		_const.EN: "english translate",
	},
	"make_friends_success_pay_1": {
		_const.RU: "А вот это совсем другой разговор.",
		_const.EN: "english translate",
	},
	"make_friends_success_pay_2": {
		_const.RU: "Как неожиданно и приятно!",
		_const.EN: "english translate",
	},
	"make_friends_success_pay_3": {
		_const.RU: "Вот так бы и сразу. А то ещё думают по нескольку световых лет.",
		_const.EN: "english translate",
	},
	"make_friends_success_pay_4": {
		_const.RU: "Хорошо, что всё кончилось именно так. И главное - никаких смертей.",
		_const.EN: "english translate",
	},
	"make_friends_success_pay_5": {
		_const.RU: "Молодец. Хвалю. Ты всё понял и сделал действительно верный выбор. Эти %CREDITS_COUNT% cr. я потрачу с умом. Могу даже отчёт как-нибудь прислать.",
		_const.EN: "english translate",
	},
	"make_friends_fail_pay_1_1": {
		_const.RU: "Нет. Мои деньги останутся только при мне!",
		_const.EN: "english translate",
	},
	"make_friends_fail_pay_1_2": {
		_const.RU: "Да нисколько ты не получишь!",
		_const.EN: "english translate",
	},
	"make_friends_fail_pay_1_3": {
		_const.RU: "Я что, такие суммы тебе скомбинировать должен? Мой ответ - не будет тебе такого частья!",
		_const.EN: "english translate",
	},
	"make_friends_fail_pay_1_4": {
		_const.RU: "А до меня кто-то на такое соглашался? Вот и я пожалуй откажусь тебе платить.",
		_const.EN: "english translate",
	},
	"make_friends_fail_pay_1_5": {
		_const.RU: "Это фирменный грабёж! Я на такое согласиться ну никак не могу.",
		_const.EN: "english translate",
	},
	"make_friends_fail_pay_2_1": {
		_const.RU: "Я тебе что из воздуха должен достать средства? Сам-то подумай!",
		_const.EN: "english translate",
	},
	"make_friends_fail_pay_2_2": {
		_const.RU: "Очень жаль, но у меня ничего не осталось. Тебе должно быть грустно, а?",
		_const.EN: "english translate",
	},
	"make_friends_fail_pay_2_3": {
		_const.RU: "А да? Как жалко, что я тебе больше заплатить не могу. Так ведь хотелось.",
		_const.EN: "english translate",
	},
	"make_friends_failed_pay_1": {
		_const.RU: "А-а, хочешь по плохому %FROM_USER_NAME%?",
		_const.EN: "english translate",
	},
	"make_friends_failed_pay_2": {
		_const.RU: "Ооо, сейчас будет такая драка!",
		_const.EN: "english translate",
	},
	"make_friends_failed_pay_3": {
		_const.RU: "Ой всё! Я за себя теперь не ручаюсь!",
		_const.EN: "english translate",
	},
	"make_friends_failed_pay_4": {
		_const.RU: "Не я эту войну начал %FROM_USER_NAME%. Как жаль, что какие-то %CREDITS_COUNT% cr. для тебя настолько важны.",
		_const.EN: "english translate",
	},
	"make_friends_failed_pay_5": {
		_const.RU: "Отказываешься значит? Хе, это у меня впервые. Но и для тебя, тогда, будет в последний раз!",
		_const.EN: "english translate",
	},

	"make_friends_failed_no_pay_1": {
		_const.RU: "Значит, ты выбрал смерть.",
		_const.EN: "english translate",
	},

	// PAY MANY
	"pay_many_failed_pay_1_1": {
		_const.RU: "Маловато будет!",
		_const.EN: "english translate",
	},
	// DROP_CARGO
	"drop_cargo_request_1": {
		_const.RU: "Оп-па, пришло время %TO_USER_NAME% немного расстаться с тяжким бременем. Бросай свой груз и сам уцелеешь!",
		_const.EN: "english translate",
	},
	"drop_cargo_request_2": {
		_const.RU: "Тебе некуда бежать %TO_USER_NAME%. Хочешь продолжить своё жалкое существование - отдавай, чего бы ты там не вёз.",
		_const.EN: "english translate",
	},
	"drop_cargo_request_3": {
		_const.RU: "Ну всё, это моя территория, а ты на ней - незванный гость. Гони свой товар и можешь двигать далее.",
		_const.EN: "english translate",
	},
	"drop_cargo_request_4": {
		_const.RU: "В этом секторе, за проход принято расплачиваться грузом %TO_USER_NAME%. Твоя персона - не исключение.",
		_const.EN: "english translate",
	},
	"drop_cargo_request_5": {
		_const.RU: "Так-так, сканеры показали что ты везёшь ценный груз, а ценным - принято делиться.",
		_const.EN: "english translate",
	},
	"drop_cargo_request_6": {
		_const.RU: "Везёшь товары, а? Хорошо, мне они тоже пригодятся. Всегда мечтал стал торговцем.",
		_const.EN: "english translate",
	},
	"drop_cargo_no_1_1": {
		_const.RU: "\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\|||||//////gag=Отправлен новый разпрос.",
		_const.EN: "english translate",
	},
	"drop_cargo_no_1_2": {
		_const.RU: "FaFa?  14-55 Модуль-Г",
		_const.EN: "english translate",
	},
	"drop_cargo_no_1_3": {
		_const.RU: "<span class=\"importantly\">Расшифрока неудалась.</span>",
		_const.EN: "english translate",
	},
	"drop_cargo_no_2_1": {
		_const.RU: "Так дело не пойдёт! Хочешь мой груз - значит ощутишь на себе всю мощь моих орудий!",
		_const.EN: "english translate",
	},
	"drop_cargo_no_2_2": {
		_const.RU: "Ты сам выискал себе проблемы %FROM_USER_NAME%!",
		_const.EN: "english translate",
	},
	"drop_cargo_no_3_1": {
		_const.RU: "Знал бы ты какой ценой этот груз достался мне. А теперь, предлагаешь вот так вот просто отказаться от него? Сейчас мы поглядим, кто - кого!",
		_const.EN: "english translate",
	},
	"drop_cargo_no_3_2": {
		_const.RU: "И не надейся! Ради этого груза я потратил свои последние сбережения! Если мне суждено погибнуть - то я буду и уничтожен вместе с ним!",
		_const.EN: "english translate",
	},
	"drop_cargo_no_3_3": {
		_const.RU: "Тогда, приготовься к бою! И пускай пустошь рассудит нас!",
		_const.EN: "english translate",
	},
	"drop_cargo_no_3_4": {
		_const.RU: "Что? Я вообще не могу понять поток твоих данных. Откалибруй там свой передатчик что-ли.",
		_const.EN: "english translate",
	},
	"drop_cargo_no_3_5": {
		_const.RU: "Да ты вообще себя видел? Спорю, ты девиант, а ещё условия мне какие-то ставишь.",
		_const.EN: "english translate",
	},
	"drop_cargo_no_3_6": {
		_const.RU: "хм, в общем-то мне этот груз ещё и самому пригодиться.",
		_const.EN: "english translate",
	},
	"drop_cargo_no_3_7": {
		_const.RU: "Ой, расчёты не обманывают меня, что всё-таки лучше повременить с таким решением.",
		_const.EN: "english translate",
	},
	"drop_cargo_no_3_8": {
		_const.RU: "Сначала ты. Потом ещё кто-то. Далее уже третий. Нет, я не стану снабжать вас всякими ресурсами без выгоды себе!",
		_const.EN: "english translate",
	},
	"drop_cargo_no_3_9": {
		_const.RU: "Так дело не пойдёт! Хочешь мой груз - значит ощутишь на себе всю мощь моих орудий!?",
		_const.EN: "english translate",
	},
	"drop_cargo_no_4_1": {
		_const.RU: "А что дальше? Может мне ещё и свой корпус заложить ради тебя? Груза у меня другого и нету.",
		_const.EN: "english translate",
	},
	"drop_cargo_no_4_2": {
		_const.RU: "Можешь меня просканировать, но всё равно ничего не найдёшь. Трюм пуст.",
		_const.EN: "english translate",
	},
	"drop_cargo_no_5_1": {
		_const.RU: "Так не пойдёт! Ты этот фокус уже проворачивал!",
		_const.EN: "english translate",
	},
	"drop_cargo_no_5_2": {
		_const.RU: "Куда ещё-то? Ты меня хочешь совсем без ничего оставить?",
		_const.EN: "english translate",
	},
	"drop_cargo_no_5_3": {
		_const.RU: "Это нечестно! Всему должна быть своя мера.",
		_const.EN: "english translate",
	},
	"drop_cargo_no_6_1": {
		_const.RU: "Сейчас я тебя граблю, ты давай в следующий раз.",
		_const.EN: "english translate",
	},
	"drop_cargo_no_7_1": {
		_const.RU: "Ты не в том положение что бы что то требовать!",
		_const.EN: "english translate",
	},
	"drop_cargo_complete_1": {
		_const.RU: "Ладно-ладно, уговорил! Только не трогай меня.",
		_const.EN: "english translate",
	},
	"drop_cargo_complete_2": {
		_const.RU: "Девианты... Хорошо, бери. Забирай. Грабитель...",
		_const.EN: "english translate",
	},
	"drop_cargo_complete_3": {
		_const.RU: "Только не стреляй! Только не стреля! Я всё отдам!",
		_const.EN: "english translate",
	},
	"drop_cargo_complete_4": {
		_const.RU: "Чтоб тебя! Куда не сунься, везде норовят ограбить. Ладно. Держи. Но я тебя запомнил %FROM_USER_NAME%.",
		_const.EN: "english translate",
	},
	"drop_cargo_complete_5": {
		_const.RU: "О нет, только не опять! Вы ребята не думали уже работать сообща? Вновь придётся расставаться с чем-то ценным...",
		_const.EN: "english translate",
	},
	"drop_cargo_complete_6": {
		_const.RU: "Ваши претензии распознаны и ради сохранения целостности собственного корпуса - будут удовлетворены. Шлюзы открыты, груз оставлен.",
		_const.EN: "english translate",
	},

	"drop_cargo_answer_complete_1_1": {
		_const.RU: "Как неожиданно и приятно!",
		_const.EN: "english translate",
	},
	"drop_cargo_answer_complete_1_2": {
		_const.RU: "А вот это совсем другой разговор.",
		_const.EN: "english translate",
	},
	"drop_cargo_answer_complete_1_3": {
		_const.RU: "Вот так бы и сразу. А то ещё думают по нескольку световых лет.",
		_const.EN: "english translate",
	},
	"drop_cargo_answer_complete_1_4": {
		_const.RU: "Хорошо, что всё кончилось именно так. И главное - никаких смертей.",
		_const.EN: "english translate",
	},

	"drop_cargo_answer_failed_1_1": {
		_const.RU: "Уверен что если ты захочешь то что нибудь найдешь!",
		_const.EN: "english translate",
	},
	"drop_cargo_answer_failed_2_1": {
		_const.RU: "Ваши претензии распознаны и ради сохранения целостности собственного корпуса - будут удовлетворены. Шлюзы открыты, груз оставлен.",
		_const.EN: "english translate",
	},

	// Robbery
	"robbery_request_1": {
		_const.RU: "Попался! Да, ты не ожидал, а пришло время платить по счетам. Ты же желаешь продолжить свой путь %TO_USER_NAME%? Мне и суммы в %CREDITS_COUNT% cr. хватит.",
		_const.EN: "english translate",
	},
	"robbery_request_2": {
		_const.RU: "Стой! В этом секторе, закон - я. Желаешь без последствий его пройти, плати мне %TO_USER_NAME% следующую сумму в %CREDITS_COUNT% cr.",
		_const.EN: "english translate",
	},
	"robbery_request_3": {
		_const.RU: "%TO_USER_NAME% здесь вообще-то принято платить. Если конечно тебя не пугает перспектива стать грудой обугленного металолома. Цена твоей жизни следующая - %CREDITS_COUNT% cr.",
		_const.EN: "english translate",
	},
	"robbery_request_4": {
		_const.RU: "Здесь проезд платный %TO_USER_NAME%. А плату за безопасность взымаю я, в размере - %CREDITS_COUNT% cr.",
		_const.EN: "english translate",
	},
	"robbery_request_5": {
		_const.RU: "Эй %TO_USER_NAME%! Хочешь, чтобы я от тебя наконец отвязался? Тебе всего-лишь то нужно заплатить сумму в %CREDITS_COUNT% cr.",
		_const.EN: "english translate",
	},
	"robbery_request_6": {
		_const.RU: "Хочешь уцелеть - перечисляй мне сумму в %CREDITS_COUNT% cr. Торга не будет. Это ультиматум!",
		_const.EN: "english translate",
	},
	"robbery_no_1_1": {
		_const.RU: "<span class=\"importantly\">Сообщение невозможно прочесть.</span>",
		_const.EN: "english translate",
	},
	"robbery_no_1_2": {
		_const.RU: "Неизвестно. Неизвестно. Неизвестно. Неизвестно.",
		_const.EN: "english translate",
	},
	"robbery_no_1_3": {
		_const.RU: "Старт трекер. Запущен протокол обороны.",
		_const.EN: "english translate",
	},
	"robbery_no_2_1": {
		_const.RU: "Мне уже терять нечего! Если мне суждено погибнуть, то я заберу тебя с собой!",
		_const.EN: "english translate",
	},
	"robbery_no_2_2": {
		_const.RU: "Это мои честно заработанные средства! И я не отдам их тебе, или кому либо другому! Приготовься ответить за свои слова!",
		_const.EN: "english translate",
	},
	"robbery_no_2_3": {
		_const.RU: "Сообщение не распознано. Возможно агрессивное проявление. Активированы системы самозащиты.",
		_const.EN: "english translate",
	},
	"robbery_no_2_4": {
		_const.RU: "Никогда! Ни за что! Ты ничего не получишь!",
		_const.EN: "english translate",
	},
	"robbery_no_2_5": {
		_const.RU: "Вымогатель! Я тебя финансировать не стану!",
		_const.EN: "english translate",
	},
	"robbery_no_2_6": {
		_const.RU: "Ваш запрос отклнонён. Он несущественен, нелогичен и невозможен.",
		_const.EN: "english translate",
	},
	"robbery_no_2_7": {
		_const.RU: "Я тут подумал и вдруг решил, что ты не получишь от меня нисколечки. Да, вот так вот.",
		_const.EN: "english translate",
	},
	"robbery_no_2_8": {
		_const.RU: "Глупо отдавать деньги, когда они тебе ещё смогут пригодиться в будущем.",
		_const.EN: "english translate",
	},
	"robbery_no_2_9": {
		_const.RU: "Что? Я вообще не могу понять поток твоих данных. Откалибруй там свой передатчик что-ли.",
		_const.EN: "english translate",
	},
	"robbery_no_2_10": {
		_const.RU: "Зря ты это затеял %FROM_USER_NAME%. Со мной шутки плохо. Но, теперь тебе уж врядли что-то поможет.",
		_const.EN: "english translate",
	},
	"robbery_no_2_11": {
		_const.RU: "Думаешь я лёгкая добыча, да? Я тебе сейчас такую трёпку задам!",
		_const.EN: "english translate",
	},
	"robbery_no_3_1": {
		_const.RU: "Так не пойдёт! Ты этот фокус уже проворачивал!",
		_const.EN: "english translate",
	},
	"robbery_no_3_2": {
		_const.RU: "Это нечестно! Всему должна быть своя мера.",
		_const.EN: "english translate",
	},
	"robbery_no_4_1": {
		_const.RU: "Была ни была! Больше шансов на то, что я сумею сбежать, чем стану расставаться с самым драгоценным.",
		_const.EN: "english translate",
	},
	"robbery_no_4_2": {
		_const.RU: "Мне уже терять нечего! Если мне суждено погибнуть, то я заберу тебя с собой!",
		_const.EN: "english translate",
	},
	"robbery_no_4_3": {
		_const.RU: "Вымогатель! Я тебя финансировать не стану!",
		_const.EN: "english translate",
	},
	"robbery_no_4_4": {
		_const.RU: "Ваш запрос отклнонён. Он несущественен, нелогичен и невозможен.",
		_const.EN: "english translate",
	},
	"robbery_no_5_1": {
		_const.RU: "Я тебе что из воздуха должен достать средства? Сам-то подумай!",
		_const.EN: "english translate",
	},
	"robbery_no_5_2": {
		_const.RU: "Очень жаль, но у меня ничего не осталось. Тебе должно быть грустно, а?",
		_const.EN: "english translate",
	},
	"robbery_no_5_3": {
		_const.RU: "А да? Как жалко, что я тебе больше заплатить не могу. Так ведь хотелось.",
		_const.EN: "english translate",
	},
	"robbery_no_7_1": {
		_const.RU: "Подобное не может быть исполнено. Однако вы, заплатите за своё нахальство.",
		_const.EN: "english translate",
	},
	"robbery_no_7_2": {
		_const.RU: "Только отчаяние могло вас толкнуть на столь большую глупость.",
		_const.EN: "english translate",
	},
	"robbery_no_7_3": {
		_const.RU: "Подобное в свой адрес я слышу впервые. Должно быть вы неисправны?",
		_const.EN: "english translate",
	},
	"robbery_no_8_1": {
		_const.RU: "Сейчас я тебя граблю, ты давай в следующий раз.",
		_const.EN: "english translate",
	},
	"robbery_no_9_1": {
		_const.RU: "Ты не в том положение что бы что то требовать!",
		_const.EN: "english translate",
	},
	"robbery_complete_1": {
		_const.RU: "Ты отвратителен! Но ради собственной сохранности я тебе заплачу.",
		_const.EN: "english translate",
	},
	"robbery_complete_2": {
		_const.RU: "Сдаюсь. Убедил. Возьми деньги, но только не тронь.",
		_const.EN: "english translate",
	},
	"robbery_complete_3": {
		_const.RU: "Да когда же вы все от меня отвяжитесь!? Плати тому. Плати этому. Мне скоро ничего не останется и ради собственного ремонта. Да подавитесь ты этими деньгами!",
		_const.EN: "english translate",
	},
	"robbery_complete_4": {
		_const.RU: "Принято. Ваши требования будут удовлетворены. На этот раз.",
		_const.EN: "english translate",
	},
	"robbery_complete_5": {
		_const.RU: "Ох. Ну... ладно. Чуть больше лишних средств. Чуть меньше лишних средств. Только бы избавить себя от проблем.",
		_const.EN: "english translate",
	},
	"robbery_complete_6": {
		_const.RU: "И ты действительно не боишься последствий за свой поступок %FROM_USER_NAME%? Хорошо, ради собственного блага, я уступлю твоим требованиям. Держи.",
		_const.EN: "english translate",
	},

	// PAY_MANY
	"pay_many_request_1": {
		_const.RU: "Так уж и быть, я дам тебе определённую сумму %CREDITS_COUNT% cr. лишь бы оставил меня в покое. Так пойдёт, %TO_USER_NAME%?",
		_const.EN: "english translate",
	},
	"pay_many_request_2": {
		_const.RU: "Давай я заплачу сумму %CREDITS_COUNT% cr., а мы сделаем вид, что друг друга не знаем? Это же так легко!",
		_const.EN: "english translate",
	},
	"pay_many_request_3": {
		_const.RU: "У меня тут есть немного %CREDITS_COUNT% cr. Давай я передам их тебе, а ты, закроешь глаза на наши с тобой разногласия в том или ином вопросе.",
		_const.EN: "english translate",
	},
	"pay_many_request_4": {
		_const.RU: "На, держи %CREDITS_COUNT% cr. и держись от меня подальше!",
		_const.EN: "english translate",
	},
	"pay_many_no_1_1": {
		_const.RU: "....? Затмение. Синий спектр. Номер?",
		_const.EN: "english translate",
	},
	"pay_many_no_1_2": {
		_const.RU: "<span class=\"importantly\">Числовая ориентация не установлена.</span>",
		_const.EN: "english translate",
	},
	"pay_many_no_1_3": {
		_const.RU: "<span class=\"importantly\">Ответ на вопрос - 42</span>",
		_const.EN: "english translate",
	},
	"pay_many_no_2_1": {
		_const.RU: "Подкуп невозможен. Вы только усугубляете свою ситуацию.",
		_const.EN: "english translate",
	},
	"pay_many_no_2_2": {
		_const.RU: "Меня невозможно подкупить!",
		_const.EN: "english translate",
	},
	"pay_many_no_2_3": {
		_const.RU: "Будь ситуация иной... Но сейчас, увы, не она.",
		_const.EN: "english translate",
	},
	"pay_many_no_3_1": {
		_const.RU: "Это взятка? Пытаешься меня таким образом принизить или подкупить? Это ты зря %FROM_USER_NAME%.",
		_const.EN: "english translate",
	},
	"pay_many_no_3_2": {
		_const.RU: "Жалкая подачка, я на такое не соглашусь, можешь даже не рассчитывать!",
		_const.EN: "english translate",
	},
	"pay_many_no_3_3": {
		_const.RU: "Э-э нет! Теперь я желаю гораздо большего! Если у тебя таковой суммы нет, значит нам придётся выяснять этот вопрос как-то иначе.",
		_const.EN: "english translate",
	},
	"pay_many_no_3_4": {
		_const.RU: "Жаль, но пока ты раздумывал ценовая ставка изменилась. Так что, ты либо платишь по повышенному тарифу, либо можешь сразу прощаться с этим бренным миром.",
		_const.EN: "english translate",
	},
	"pay_many_complete_1": {
		_const.RU: "Другое дело! А теперь вали, во второй раз, тебя %FROM_USER_NAME% я так просто не соизволю отпустить.",
		_const.EN: "english translate",
	},
	"pay_many_complete_2": {
		_const.RU: "Славно-славно. Такой суммы мне хватит надолго. Но предупреждаю тебя %FROM_USER_NAME%, пересекись мы ещё раз, тебе придётся изрядно раскошелиться.",
		_const.EN: "english translate",
	},
	"pay_many_complete_3": {
		_const.RU: "Сумма зачислена. Можешь продолжать свой путь синтет.",
		_const.EN: "english translate",
	},
	"pay_many_complete_4": {
		_const.RU: "Ну раз уж ты так просишь и согласен на всё, не вижу смысла тебя задерживать %FROM_USER_NAME%.",
		_const.EN: "english translate",
	},

	// ATTACK TARGET
	"attack_target_request_1": {
		_const.RU: "Слушай, слушай, может дерзнём и нападём на %TARGET_NAME%? Добычу поделим между собой. Всё честно.",
		_const.EN: "english translate",
	},
	"attack_target_request_2": {
		_const.RU: "Здесь есть довольно известный преступник - %TARGET_NAME%. За него полагается внушительная награда. Предлагаю присоединиться ко мне и наконец прекратить его бесчинства.",
		_const.EN: "english translate",
	},
	"attack_target_request_3": {
		_const.RU: "%TO_USER_NAME% что скажешь насчёт совместного раздела добычи - %TARGET_NAME%? Только чур, кто внёс больший вклад, тот и получает больший куш!",
		_const.EN: "english translate",
	},
	"attack_target_request_4": {
		_const.RU: "%TO_USER_NAME% к вам поступило предложение о помощи в устарении цели - %TARGET_NAME%. Каков будет ваш ответ?",
		_const.EN: "english translate",
	},
	"attack_target_request_5": {
		_const.RU: "%TO_USER_NAME% я уже давно разыскивал этого - %TARGET_NAME%. У меня к нему личные счёты. Но, я бы често не отказался заиметь друга, что помог бы мне разобраться с - %TARGET_NAME%.",
		_const.EN: "english translate",
	},
	"attack_target_request_6": {
		_const.RU: "А вот и он! %TO_USER_NAME%, здесь есть один синтет - %TARGET_NAME%. Есть сведения, что он внезакона. К тому же, перевозит незаконный груз. Короче, я собираюсь заняться им, но, от помощи извне не отказался бы.",
		_const.EN: "english translate",
	},
	"attack_target_no_1_1": {
		_const.RU: "Все и всегда. Они вечные.",
		_const.EN: "english translate",
	},
	"attack_target_no_1_2": {
		_const.RU: "<span class=\"importantly\">Трансляция белого шума.</span>",
		_const.EN: "english translate",
	},
	"attack_target_no_1_3": {
		_const.RU: "<span class=\"importantly\">Помехи. Невозможно установить связь.</span>",
		_const.EN: "english translate",
	},
	"attack_target_no_2_1": {
		_const.RU: "Я не могу совершить нападение на самого себя.",
		_const.EN: "english translate",
	},
	"attack_target_no_2_2": {
		_const.RU: "Неверная цель. Это невозможно.",
		_const.EN: "english translate",
	},
	"attack_target_no_2_3": {
		_const.RU: "О-о-о, пойди проверься, ты же неисправен!",
		_const.EN: "english translate",
	},
	"attack_target_no_3_1": {
		_const.RU: "Да у тебя явные проблемы. Где хоть переклинило?",
		_const.EN: "english translate",
	},
	"attack_target_no_3_2": {
		_const.RU: "Я не могу оказывать какую-либо помощь своим иделогическим противникам.",
		_const.EN: "english translate",
	},
	"attack_target_no_3_3": {
		_const.RU: "Мы не союзники! Сеанс связи завершён.",
		_const.EN: "english translate",
	},
	"attack_target_no_3_4": {
		_const.RU: "Это предложение к предательству своей фракции. Я отказываюсь его рассматривать, а уж тем более - принимать.",
		_const.EN: "english translate",
	},
	"attack_target_no_3_5": {
		_const.RU: "Не пытайся меня склонить на свою сторону. Мы по разные стороны баррикад.",
		_const.EN: "english translate",
	},
	"attack_target_no_3_6": {
		_const.RU: "Оказание помощи доступно только частям из союзных фракций. Таковой вы не являетесь.",
		_const.EN: "english translate",
	},
	"attack_target_no_3_7": {
		_const.RU: "Не принимается к рассмотрению. Электронная подпись не распознана. Вы не являетесь мне союзником.",
		_const.EN: "english translate",
	},
	"attack_target_no_4_1": {
		_const.RU: "Между нами подобных соглашений быть не может. Если только я не стану тебе доверять в будущем.",
		_const.EN: "english translate",
	},
	"attack_target_no_4_2": {
		_const.RU: "Я с грабителями навроде тебя дел не имею. На этом всё!",
		_const.EN: "english translate",
	},
	"attack_target_no_4_3": {
		_const.RU: "Что-то я в тебе глубоко сомневаюсь. Может у тебя и получится меня переубедить в будущем, но сейчас, доверия ты не вызываешь.",
		_const.EN: "english translate",
	},
	"attack_target_no_4_4": {
		_const.RU: "О-о-о, пойди проверься, ты же неисправен!",
		_const.EN: "english translate",
	},
	"attack_target_no_4_5": {
		_const.RU: "Мы не союзники! Сеанс связи завершён.",
		_const.EN: "english translate",
	},
	"attack_target_no_5_1": {
		_const.RU: "Предлагаемый объект - союзник. Отказано и более не рассматриваемо.",
		_const.EN: "english translate",
	},
	"attack_target_no_5_2": {
		_const.RU: "На кого ты собрался напасть? Ещё раз - на кого!?",
		_const.EN: "english translate",
	},
	"attack_target_no_6_1": {
		_const.RU: "Плохой план! Я уж молчу о том, как за нами по всем секторам буду потом гоняться силы закона.",
		_const.EN: "english translate",
	},
	"attack_target_no_6_2": {
		_const.RU: "Неа, это твоё предложение и твои дела. В случае чего, на меня можешь не рассчитывать.",
		_const.EN: "english translate",
	},
	"attack_target_no_6_3": {
		_const.RU: "Ответ - отрицательный. Отменить подготовки всех боевых операций.",
		_const.EN: "english translate",
	},
	"attack_target_no_6_4": {
		_const.RU: "Я в таком не участвую! Даже не надейся! Это вообще противозаконно!",
		_const.EN: "english translate",
	},
	"attack_target_no_6_5": {
		_const.RU: "Чтобы ты там от кого не хотел, но это сугубо ваши проблемы. Сами их и решайте!",
		_const.EN: "english translate",
	},
	"attack_target_no_6_6": {
		_const.RU: "Не. Какой уважающий себя синтет согласится на подобное. Вот я - точно нет.",
		_const.EN: "english translate",
	},
	"attack_target_no_7_1": {
		_const.RU: "Члены фракции никогда не опустятся до нападения на союзников.",
		_const.EN: "english translate",
	},
	"attack_target_no_7_2": {
		_const.RU: "Я не могу оказывать какую-либо помощь своим иделогическим противникам.",
		_const.EN: "english translate",
	},
	"attack_target_no_7_3": {
		_const.RU: "О-о-о, пойди проверься, ты же неисправен!",
		_const.EN: "english translate",
	},
	"attack_target_no_8_": {
		_const.RU: "Это невозможно. Я уже нахожусь в союзе с %TARGET_NAME%. Имей это ввиду, если затеешь нечто.",
		_const.EN: "english translate",
	},
	"attack_target_no_9_1": {
		_const.RU: "Не сейчас, здесь слишком жарко! Не видишь, меня чуть не подбили!?",
		_const.EN: "english translate",
	},
	"attack_target_no_9_2": {
		_const.RU: "У меня тут битва как бы! Не лезь ко мне с подобным.",
		_const.EN: "english translate",
	},
	"attack_target_no_10_1": {
		_const.RU: "Неа, это твоё предложение и твои дела. В случае чего, на меня можешь не рассчитывать.",
		_const.EN: "english translate",
	},
	"attack_target_no_10_2": {
		_const.RU: "Да ты с ума сошёл? Ты видел его огневую мощь!? Он же нас в порошёк сотрёт! Ты как хочешь, а я сваливаю куда подальше!",
		_const.EN: "english translate",
	},
	"attack_target_no_10_3": {
		_const.RU: "Не. Какой уважающий себя синтет согласится на подобное. Вот я - точно нет.",
		_const.EN: "english translate",
	},
	"attack_target_no_10_4": {
		_const.RU: "Неа, тут уж ты сам. Мне оно ни к чему.",
		_const.EN: "english translate",
	},
	"attack_target_no_11_1": {
		_const.RU: "Я занят - выполняю важную миссию.",
		_const.EN: "english translate",
	},
	"attack_target_no_11_2": {
		_const.RU: "Я на задании.",
		_const.EN: "english translate",
	},
	"attack_target_no_12_1": {
		_const.RU: "Сначала я разберусь с тобой!",
		_const.EN: "english translate",
	},
	"attack_target_no_13_1": {
		_const.RU: "Я занят другой целью!",
		_const.EN: "english translate",
	},
	"attack_target_complete_1": {
		_const.RU: "А что, я согласен! Была ни была!",
		_const.EN: "english translate",
	},
	"attack_target_complete_2": {
		_const.RU: "Если только это законно, то я в деле!",
		_const.EN: "english translate",
	},
	"attack_target_complete_3": {
		_const.RU: "Ха, ещё бы! Ну, погнали!",
		_const.EN: "english translate",
	},
	"attack_target_complete_4": {
		_const.RU: "Положительно. Инициирую боевые системы.",
		_const.EN: "english translate",
	},
	"attack_target_complete_5": {
		_const.RU: "Ну, если уж игра стоит свеч, то так и быть - помогу!",
		_const.EN: "english translate",
	},
	"attack_target_complete_6": {
		_const.RU: "Убедил. Но это в первый и последний раз.",
		_const.EN: "english translate",
	},
	"attack_target_leave_1": {
		_const.RU: "Что-то мне резко перехотелось. Изначально было какое-то нехорошее предчувствие.",
		_const.EN: "english translate",
	},
	"attack_target_leave_2": {
		_const.RU: "Я тут вдруг просканировал его оружейные системы... Короче, ты друг сам по себе.",
		_const.EN: "english translate",
	},
	"attack_target_leave_3": {
		_const.RU: "Тут выяснилось, что у меня топливо почти на исходе. Я сгоняю на ближайшую базу, а ты, подожди меня малёха. Вернусь. Даю слово.",
		_const.EN: "english translate",
	},

	// Defend Drop Cargo
	"defend_cargo_request_1": {
		_const.RU: "Чтоб тебя Replics переплавил! Держи груз и оставь меня уже наконец-то в покое!",
		_const.EN: "english translate",
	},
	"defend_cargo_request_2": {
		_const.RU: "Может, если я сброшу свой груз, ты всё-таки угомонишься? Мне не нужна драка!",
		_const.EN: "english translate",
	},
	"defend_cargo_request_3": {
		_const.RU: "Сбросить груз, попытка оторваться от налётчика.",
		_const.EN: "english translate",
	},
	"defend_cargo_request_4": {
		_const.RU: "Ты только не стреляй! Возьми, что у меня есть и отпусти. Ладно?",
		_const.EN: "english translate",
	},
	"defend_cargo_complete_1": {
		_const.RU: "О-о! Вот так бы всегда! Проваливай!",
		_const.EN: "english translate",
	},
	"defend_cargo_complete_2": {
		_const.RU: "Ну-ка, ну-ка! Сейчас посмотрим, чем ты был богат.",
		_const.EN: "english translate",
	},
	"defend_cargo_complete_3": {
		_const.RU: "Есть чем поживиться! Осталось только доставить куда надо и навариться.",
		_const.EN: "english translate",
	},
	"defend_cargo_complete_4": {
		_const.RU: "Такое, меня более чем удовлетворяет. Хе, а теперь попытайся добраться до ближайше базы.",
		_const.EN: "english translate",
	},
	"defend_cargo_complete_5": {
		_const.RU: "На этот раз сгодится. Но знай, попадёшься ещё хотя бы раз, такой малой долей не сумеешь откупиться.",
		_const.EN: "english translate",
	},
	"defend_cargo_no_1_1": {
		_const.RU: "SDFK1332-11/5 ... OGO!",
		_const.EN: "english translate",
	},
	"defend_cargo_no_1_2": {
		_const.RU: "............................................. Вновь.",
		_const.EN: "english translate",
	},
	"defend_cargo_no_2_1": {
		_const.RU: "Ха! Ха-ха! Я тебе что, какой-то грабитель!?",
		_const.EN: "english translate",
	},
	"defend_cargo_no_2_2": {
		_const.RU: "Это даже как-то... оскорбительно!",
		_const.EN: "english translate",
	},
	"defend_cargo_no_2_3": {
		_const.RU: "Не может быть и речи!",
		_const.EN: "english translate",
	},
	"defend_cargo_no_3_1": {
		_const.RU: "Что это за дрянь? Ты меня за идиота держишь? Мне нужен более ценный груз, или, ты за своё нахальство поплатишься!",
		_const.EN: "english translate",
	},
	"defend_cargo_no_3_2": {
		_const.RU: "И всё? Это жалкая мелочь, которая ничего не стоит. Хочешь меня спровоцировать на более радикальные действия?",
		_const.EN: "english translate",
	},
	"defend_cargo_no_3_3": {
		_const.RU: "Продолжить преследование! Цель попыталась сбить с толку посредством малозначительного груза.",
		_const.EN: "english translate",
	},
	"defend_cargo_no_3_4": {
		_const.RU: "Меня таким не подкупишь! Всё, теперь, твои дни уж точно сочтены!",
		_const.EN: "english translate",
	},
	"defend_cargo_no_3_5": {
		_const.RU: "А вот это была ошибка! Зря ты попытался начать диалог с подобного предложения!",
		_const.EN: "english translate",
	},

	// EXPEDITION - новости типо от лица первого канала, а не от главы фракции
	"expedition_move_1": {
		_const.RU: "Появились определённые сведения, что в секторе %SECTORE_NAME% было замечено скопление синтетов. Предположительно, они относятся к фракции %FRACTION_NAME% и представляют собой конвой военизированного типа. Всем оказавшимся поблизости синтетам, рекомендуется не пересекать дорогу и не провоцировать участников конвоя на какой-либо конфликт.",
		_const.EN: "english translate",
	},
	"expedition_trader_move_1": {
		_const.RU: "Как сообщается, появилась достоверная информация о деятельности экспедиции <span class=\\\"importantly\\\">%FRACTION_NAME%</span> судя по всему, данная группа синтетов держит свой путь в сектор <span class=\\\"importantly\\\">%MAP_NAME%</span>. Подлинные цели этих синтетов не известны, из-за чего, рекомендуется соблюдать осторожность при попытке контакта.",
		_const.EN: "english translate",
	},
	// бойцы фракции убили мирные транспорты
	// _Replics_Explores_: Replics - кто убил, Explores - для кого новость (и кого убили)
	"fraction_warrior_kill_freeland_Replics_Explores_1": {
		_const.RU: "Увы, до нас дошли кое-какие печальные известия: произошла катастрофа в секторе <span class=\\\"importantly\\\">%MAP_NAME%</span> подлые убийцы <span class=\\\"importantly\\\">Replics</span> осуществили нападение и уничтожили представителей нашей фракции. Пустоши планеты вновь пополнились обломками синтетов…",
		_const.EN: "english translate",
	},
	"fraction_warrior_kill_freeland_Replics_Reverses_1": {
		_const.RU: "Да будет всем известно об инциденте, что случился в секторе <span class=\\\"importantly\\\">%MAP_NAME%</span>, где боевые единицы Replics <span class=\\\"importantly\\\">Replics</span> нанесли удар, что повлёк за собой жертвы среди Reverses. Reverses в свою очередь, открыто даёт знать - ни один подобный акт агрессии не останется безнаказанным!",
		_const.EN: "english translate",
	},
	"fraction_warrior_kill_freeland_Explores_Replics_1": {
		_const.RU: "Гнусное преступление было осуществлено против Replics! Уже известно, что в секторе <span class=\\\"importantly\\\">%MAP_NAME%</span> представители так называемого Explores… <span class=\\\"importantly\\\">Explores</span> произвели безжалостное и не имеющее оснований нападений на синтетов нашей фракции! Погибших много… Но Replics, всегда оставляет за собой право возмездия!",
		_const.EN: "english translate",
	},
	"fraction_warrior_kill_freeland_Explores_Reverses_1": {
		_const.RU: "Имеется достоверная информация - в секторе <span class=\\\"importantly\\\">%MAP_NAME%</span> некие явно радикальные единицы Explores <span class=\\\"importantly\\\">Explores</span> осуществили боевое столкновение с гражданскими синтетами Reverses. Точного количества жертв вследствие этого варварского нападения пока что установить не удаётся…",
		_const.EN: "english translate",
	},
	"fraction_warrior_kill_freeland_Reverses_Explores_1": {
		_const.RU: "Похоже, что в печально известном секторе <span class=\\\"importantly\\\">%MAP_NAME%</span> вновь случилась беда – предположительно, представители Reverses <span class=\\\"importantly\\\">Reverses</span> атаковали и стали причиной гибели смелых исследователей нашей фракции! Мы увековечим память о них в хрониках изучения этой планеты, а что касается Reverses… дипломатия воздействует куда сильнее и больнее любого распылителя материи.",
		_const.EN: "english translate",
	},
	"fraction_warrior_kill_freeland_Reverses_Replics_1": {
		_const.RU: "Опубликован доклад разведки по инциденту в секторе <span class=\\\"importantly\\\">%MAP_NAME%</span>: заманив членов фракции Replics в засаду, некие синтеты с боевой сигнатурой Reverses <span class=\\\"importantly\\\">Reverses</span> учинили атаку на транспорты и гражданский сегмент Replics. Replics не волнуют оправдания и выдумки по данному инциденту со стороны Reverses, ответ последует незамедлительно!",
		_const.EN: "english translate",
	},
	"place_expedition_Replics_1": {
		_const.RU: "Стал известен доклад от экспедиции - храбрым синтетам нашей фракции <span class=\\\"importantly\\\">Replics</span> удалось закрепиться в секторе \\\"<span class=\\\"importantly\\\">%MAP_NAME%</span> основать там временный форпост, установить наблюдение и расширить сферу влияния Replics!",
		_const.EN: "english translate",
	},
	"place_expedition_Explores_1": {
		_const.RU: "Отважные исследователи-первооткрыватели Explores, будучи в составе экспедиции <span class=\\\"importantly\\\">Explores</span> сумели обнаружить и провести первоначальное изучение сектора \\\"<span class=\\\"importantly\\\">%MAP_NAME%</span>. Учитывая данные беглого инвазивного анализа, Explores ожидают существенные открытия…",
		_const.EN: "english translate",
	},
	"place_expedition_Reverses_1": {
		_const.RU: "Смельчаки-добровольцы Reverses из экспедиции <span class=\\\"importantly\\\">Reverses</span> дали знать, что достигли ранее скрытого сектора \\\"<span class=\\\"importantly\\\">%MAP_NAME%</span> укрепились там, возвели инфраструктуру для дальнейшей деятельности. Reverses неумолимо движется к достижению своей главной стратегической цели!",
		_const.EN: "english translate",
	},
	"destroy_expedition_Replics_1": {
		_const.RU: "Несмотря на цензуру, до Replics уже дошли тревожные и, увы, правдивые слухи - наша экспедиция <span class=\\\"importantly\\\">Replics</span> что находилась в секторе \\\"<span class=\\\"importantly\\\">%MAP_NAME%</span>\\\" более не выходит на любые каналы связи. Replics утвердил следующее постановление - считать экспедицию проваленной, а всех её членов - пропавшими без вести.",
		_const.EN: "english translate",
	},
	"destroy_expedition_Explores_1": {
		_const.RU: "Это… правда! В ходе исследовательских изысков, наша экспедиция <span class=\\\"importantly\\\">Explores</span> находящаяся в секторе \\\"<span class=\\\"importantly\\\">%MAP_NAME%</span>\\\" понесла невосполнимые потери личного состава в ходе воздействия окружающей среды. Сожалея об утраченных научных возможностях, Explores объявляет экспедицию проваленной и отзывает её.",
		_const.EN: "english translate",
	},
	"destroy_expedition_Reverses_1": {
		_const.RU: "Иного выбора нету и Reverses придётся пойти на данный шаг - наша экспедиция <span class=\\\"importantly\\\">Reverses</span> присутствующая в секторе \\\"<span class=\\\"importantly\\\">%MAP_NAME%</span>\\\" грубо нарушила протоколы Reverses, что привело к потере дорогостоящего оборудования и конфронтации между членами экспедиции. В связи с утратой ценных кадров и оборудования, Reverses прекращает деятельность экспедиции.",
		_const.EN: "english translate",
	},
	"new_expedition_Replics_1": {
		_const.RU: "Инициализация исследовательско-разведывательной деятельности <span class=\\\"importantly\\\">Replics</span> отправка военизированной экспедиции в сектор \\\"<span class=\\\"importantly\\\">%TO_MAP_NAME%</span>\\\", Экспедиция доказав свою приверженность идеям Replics и способности к выживанию, получает следующее распоряжение - отправить экспедицию из сектора \\\"<span class=\\\"importantly\\\">%MAP_NAME%</span>\\. По прибытии, будет проведена новая оценка.",
		_const.EN: "english translate",
	},
	"new_expedition_Explores_1": {
		_const.RU: "Все приготовления завершены, исследователи вышли на связь <span class=\\\"importantly\\\">Explores</span> и готовы к отправлению экспедиции в сектор \\\"<span class=\\\"importantly\\\">%TO_MAP_NAME%</span>\\\", окончив перечень работ и передав данные, экспедиция отправляется далее из сектора \\\"<span class=\\\"importantly\\\">%MAP_NAME%</span>\\, где вновь примется за выполнение последующих миссий.",
		_const.EN: "english translate",
	},
	"new_expedition_Reverses_1": {
		_const.RU: "Лучшие из лучших, славные храбрецы <span class=\\\"importantly\\\">Reverses</span> решившие статью частью экспедиции, теперь, отправляются со своей благородной миссией в сектор \\\"<span class=\\\"importantly\\\">%TO_MAP_NAME%</span>\\\". Достойно окончив первую половину целей, экспедиция движется далее - из сектора \\\"<span class=\\\"importantly\\\">%MAP_NAME%</span>\\, там, вновь приступив к исполнению своих обязанностей.",
		_const.EN: "english translate",
	},
	"meteorite_rain_1": {
		_const.RU: "Синтеты - будьте предельно осторожны! Судя по данным атмосферных зондов, в секторе <span class=\\\"importantly\\\">%MAP_NAME%</span> вот-вот произойдёт метеоритный дождь! Последствия могут иметь разрушительный характер!",
		_const.EN: "english translate",
	},
	// Другая фракция решает уничтожить экспедицию врага
	// _Replics_Explores_: Replics - кто атакует, Explores - кого атакует
	"attack_expedition_Replics_Explores_1": {
		_const.RU: "Сообщает: <span class=\\\"importantly\\\">Replics</span> сектор <span class=\\\"importantly\\\">%TO_MAP_NAME%</span> где присутствуют враждебные для Replics ничтожные силы <span class=\\\"importantly\\\">Explores</span> подвергнется атаке с целью зачистки и устранения возможных угроз в будущем. После окончания операции, силы Replics выдвинуться далее в сектор <span class=\\\"importantly\\\">%MAP_NAME%</span>.",
		_const.EN: "english translate",
	},
	"attack_expedition_Replics_Reverses_1": {
		_const.RU: "Несравненный лидер <span class=\\\"importantly\\\">Replics</span> предоставил инструкции, касающиеся военной операции по очищению сектора <span class=\\\"importantly\\\">%TO_MAP_NAME%</span>, где сейчас присутствуют дерзкие войска незваных захватчиков из <span class=\\\"importantly\\\">Reverses</span>! <p> Когда военные цели операции будут достигнуты, силы Replics продолжат свой путь к уже новым задачам из сектора - <span class=\\\"importantly\\\">%MAP_NAME%</span>.",
		_const.EN: "english translate",
	},
	"attack_expedition_Explores_Replics_1": {
		_const.RU: "Талантливые умы <span class=\\\"importantly\\\">Explores</span> дали нам знать о важном решении задействования военных действий в секторе <span class=\\\"importantly\\\">%TO_MAP_NAME%</span>, где замечены главнейшие враги Explores - <span class=\\\"importantly\\\">Replics</span>! <p> Когда с угрозой будет покончено, силы Explores продолжат свою путь, но уже к новым задачам из сектора <span class=\\\"importantly\\\">%MAP_NAME%</span>.",
		_const.EN: "english translate",
	},
	"attack_expedition_Explores_Reverses_1": {
		_const.RU: "Признанный всеми фракциями выдающийся конгломерат разумов <span class=\\\"importantly\\\">Explores</span> дал известие, касающееся особых событий в секторе <span class=\\\"importantly\\\">%TO_MAP_NAME%</span>, где будет пресечена деятельность девиантной группы синтетов-террористов <span class=\\\"importantly\\\">Reverses</span>! <p> С устранением этой угрозы, а также прервав их возможную пиратскую деятельность, наш специальный отряд продолжит свой патруль из сектора - <span class=\\\"importantly\\\">%MAP_NAME%</span>.",
		_const.EN: "english translate",
	},
	"attack_expedition_Reverses_Explores_1": {
		_const.RU: "Славное сообщество <span class=\\\"importantly\\\">Reverses</span> даёт исчерпывающий комментарий: произойдёт зачистка в секторе <span class=\\\"importantly\\\">%TO_MAP_NAME%</span> от враждебных и незаконно пересёкших границу Reverses элементов <span class=\\\"importantly\\\">Explores</span>! <p> Когда задача будет выполнена, силы Reverses будут переброшены в следующий сектор для выполнения ряда прочих заданий - <span class=\\\"importantly\\\">%MAP_NAME%</span>.",
		_const.EN: "english translate",
	},
	"attack_expedition_Reverses_Replics_1": {
		_const.RU: "Сообщество <span class=\\\"importantly\\\">Reverses</span> делится информацией следующего толка – будет предпринята попытка устранения множества враждебных сигнатур в секторе <span class=\\\"importantly\\\">%TO_MAP_NAME%</span>, которые принадлежат некой, возможно враждебной миссии <span class=\\\"importantly\\\">Replics</span>! <p> Выявив, разгромив и преподав урок Replics, особый отряд Reverses будет перемещён в иной сектор со схожей проблематикой - <span class=\\\"importantly\\\">%MAP_NAME%</span>.",
		_const.EN: "english translate",
	},
	"attack_expedition_win_Replics_1": {
		_const.RU: "Непоколебимые силы Replics, нанеся сокрушительный удар, одержали убедительную и неоспоримую победу в секторе <span class=\\\"importantly\\\">%MAP_NAME%</span>!",
		_const.EN: "english translate",
	},
	"attack_expedition_win_Explores_1": {
		_const.RU: "Только путём консолидации наших действий и умелых военных манёвров, наши силы сумели побороть хитрого врага в секторе <span class=\\\"importantly\\\">%MAP_NAME%</span> и рассеять его!",
		_const.EN: "english translate",
	},
	"attack_expedition_win_Reverses_1": {
		_const.RU: "Непросто, с ощутимыми последствиями, но объединившись, придерживаясь чёткого плана, войска сообщества Reverses одержали верх и обратили в бегство недругов в секторе <span class=\\\"importantly\\\">%MAP_NAME%</span>!",
		_const.EN: "english translate",
	},
	"destroy_attack_expedition_Replics_1": {
		_const.RU: "Присутствует информация следующего толка: в секторе <span class=\"importantly\">%MAP_NAME%</span> наши отважные войска понесли потери и были вынуждены отступить для перегруппировки. Любые сведения о поражении в битве за эту территорию - преступная ложь!",
		_const.EN: "english translate",
	},
	"destroy_attack_expedition_Explores_1": {
		_const.RU: "Печальные известия - наши силы потерпели крах при зачистке сектора <span class=\"importantly\">%MAP_NAME%</span> и теперь, в попытке спасти себя от полного разгрома, стремительно покидают его.",
		_const.EN: "english translate",
	},
	"destroy_attack_expedition_Reverses_1": {
		_const.RU: "Только что стала известна новость скорбного характера - силы сообщества Reverses понесли ряд боевых поражений в секторе <span class=\"importantly\">%MAP_NAME%</span>.",
		_const.EN: "english translate",
	},

	// FRACTION WAR START - новости типо от лица первого канала а не от главы фракции
	"start_fraction_war_Replics_Replics_1": {
		_const.RU: "Неожиданное известие потрясло многие фракции синтетов - Replics во всеуслышание заявил о вступлении в войну за территории! Мы будем следить за ходом этих и прочих событий.",
		_const.EN: "english translate",
	},
	"start_fraction_war_Replics_Explores_1": {
		_const.RU: "Данный мир не мог быть вечен! Экстренное событие: Replics после концентрации своих войск у границ спорных территорий, официально объявляет войну Explores! По-видимому, Explores придётся прибегнуть к услугам синтетов-наёмников, чтобы отстоять свои сектора против армад Replics.",
		_const.EN: "english translate",
	},
	"start_fraction_war_Replics_Reverses_1": {
		_const.RU: "Рано или поздно, но это должно было случиться! Экстренное событие: Replics придерживаясь официоза во всеуслышание заявил, что объявляет войну за спорные территории Reverses. Похоже, что это противостояние будет долгим и не одна из сторон не сумеет добиться ощутимой победы.",
		_const.EN: "english translate",
	},
	"start_fraction_war_Explores_Explores_1": {
		_const.RU: "Этого никто не мог ожидать, а в особенности - от них! Как уже стало известно из множества источников: конгломерат разумов официально подтвердил, что будет бороться за спорные территории во имя науки без границ и просвещения без запретов.",
		_const.EN: "english translate",
	},
	"start_fraction_war_Explores_Replics_1": {
		_const.RU: "Это смелый шаг или наглядное отчаяние? Судя по всему, Explores официально решил побороться за право на обладание спорными территориями с Replics и потому объявляет войну. Никто не знает сколь долго могут продлиться эти сражения, учитывая большую разницу в весовой категории оппонентов. Мы будем держать вас в курсе дел.",
		_const.EN: "english translate",
	},
	"start_fraction_war_Explores_Reverses_1": {
		_const.RU: "Очередная экстренная новость, что может стать хорошим источником заработка синтетов-наёмников! По-видимому, Explores объявили войну за спорные территории Reverses. Наши спутники уже фиксируют целые очаги вооружённых столкновений, что бушуют в тамошних местах. Кто же из соперников одержит верх?",
		_const.EN: "english translate",
	},
	"start_fraction_war_Reverses_Reverses_1": {
		_const.RU: "Слухи о новой войне фракций оказались правдой! Как уже стало известно, именно Reverses в краткие сроки мобилизовав свои силы, открыто, по всем дипломатическим каналам объявил о начале войн за спорные территории. Мы рекомендуем честным синтетам держаться подальше от секторов, где будут происходить боевые действия. Тем же, кто хочет подзаработать – следует обратиться в пункты армейского найма противоборствующих сторон.",
		_const.EN: "english translate",
	},
	"start_fraction_war_Reverses_Replics_1": {
		_const.RU: "Этого противостояния ждали все! Как уже стало известно, Reverses претендует на некоторые из спорных регионов, а потому, очевидно заведомо подготовившись, только что открыто объявил войну Replics! Replics в свою очередь дал ответный ход, но уже действиями. Две могущественные фракции синтетов вот-вот схлестнуться в первом сражении. Мы будем внимательно следить за развитием ситуации и оповещать, если ход войн, примет неожиданный формат.",
		_const.EN: "english translate",
	},
	"start_fraction_war_Reverses_Explores_1": {
		_const.RU: "Должно ли они воевать? Есть ли им что делить? Очевидно, что Reverses отвергает оба этих вопроса, ведь фракция открыто объявила войну Explores за спорные территории. Очевидно, что это не битва титанов, а конкретно - противостоянии хитрости и умения. Такая война, может пойти по любому теоретическому сценарию событий. Мы будет держать вас в курсе дел!",
		_const.EN: "english translate",
	},

	// FRACTION WAR CREATE ARMY - новости типо от лица первого канала, а не от главы фракции
	"fraction_war_create_army_Replics_1": {
		_const.RU: "Поступила новость, что, само собой, ни для кого не станет откровением: Фракция Replics провела внеочередное пополнение личного состава, расширив количество войск и тем самым плотно прикрыв границы своих территорий от любых посягательств. Хотя, зная специфику ведения дипломатии со стороны Replics, подобный шаг, может свидетельствовать о подготовке к новой войне за спорные территории.",
		_const.EN: "english translate",
	},
	"fraction_war_create_army_Explores_1": {
		_const.RU: "Планета полнится различными слухами и на сей раз, один из таких поводов касается фракции Explores: как сообщается, примерно около трёх циклов назад, промышленные фабрикаторы и заводы общего назначения фракции, стало массово выпускать продукцию военного образца. Explores даже временно закрыли границы для допуска на эти территории синтетов из иных фракций. Всё это может свидетельствовать только об одном - Explores готовится к вооружённому противостоянию!",
		_const.EN: "english translate",
	},
	"fraction_war_create_army_Reverses_1": {
		_const.RU: "Новые события и новости на Veliri: судя по многочисленной инсайдерской информации, полученной за последние циклы, а также глядя на необычайную “подвижность” фракции Reverses за тот же самый период, можно смело констатировать - Reverses готовится к серьёзному столкновению с равной по себе силой. Эту же информацию подтверждают и слухи иного толка, сообщающие, что Reverses экстренно заключает множественные краткосрочные контракты с синтетами, что желают стать наёмниками.",
		_const.EN: "english translate",
	},

	// FRACTION CAPTURE BASES - новости типо от лица первого канала, а не от главы фракции
	"fraction_war_capture_sector_Replics_Replics_1": {
		_const.RU: "Экстренная новость: Силы Replics заняли и теперь полностью контролируют сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> . Судя по всему, именно из этой позиции в будущем Replics будет проводить все свои прочие наступательные операции. Иные фракции, пока что никак не прокомментировали действия Replics, а также не заявили о намерении изгнать силы тамошних синтетов с оспариваемой территории.",
		_const.EN: "english translate",
	},
	"fraction_war_capture_sector_Replics_Explores_1": {
		_const.RU: "Война продолжается: приходят известия, что в ходе столкновений сторон, силам Explores всё же удалось взять под полный контроль сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> тем самым потеснив позиции Replics. Стоит ли рассчитывать на ответный контрудар со стороны Replics? Время покажет!",
		_const.EN: "english translate",
	},
	"fraction_war_capture_sector_Replics_Reverses_1": {
		_const.RU: "Конфликт, что не имеет завершения… Reverses отчитались о полном контроле над сектором <span class=\\\"importantly\\\">%MAP_NAME%</span>, где ранее базировались силы Replics. Это смелый трюк и дерзкий вызов в сторону самого Replics, что очевидно пожелает взять реванш за такое унизительное потеснение своих позиций. Как всегда, мы вас известим о ходе войны и тех переменах, что в ней происходят.",
		_const.EN: "english translate",
	},
	"fraction_war_capture_sector_Explores_Explores_1": {
		_const.RU: "Срочное известие: Explores заняли нейтральный сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> одной из оспариваемых территорий. Возможно, что такой ход подстегнёт все остальные противоборствующие стороны ускорить свои приготовления и поторопить события, если только они не желают, чтобы именно Explores выбился в гегемоны по количеству контроля территорий этой войны!",
		_const.EN: "english translate",
	},
	"fraction_war_capture_sector_Explores_Replics_1": {
		_const.RU: "Этого стоило ожидать… Пытаясь скрыть сей факт, Explores умолчали о потере сектора <span class=\\\"importantly\\\">%MAP_NAME%</span>, чьими хозяевами теперь стали силы Replics. Должно быть, потеряв такой важный плацдарм, теперь, выбить оттуда Replics становится проблематичной задачей. С другой стороны, сам Replics теперь будет вести ещё более убедительные наступательные действия, завладев необходимой территорией.",
		_const.EN: "english translate",
	},
	"fraction_war_capture_sector_Explores_Reverses_1": {
		_const.RU: "Это война - так уж сложилось! И вновь известия с полей фракционной войны: Reverses овладели новыми спорными территориями в лице сектора <span class=\\\"importantly\\\">%MAP_NAME%</span>, что ранее принадлежал Explores. Потеря столь ценной территории, это несомненно серьёзный удар. Многие даже прогнозируют, что Explores не станет пытаться вернуться потерянную территорию и в принципе, может выйти из этого конфликта без каких-либо приобретений.",
		_const.EN: "english translate",
	},
	"fraction_war_capture_sector_Reverses_Reverses_1": {
		_const.RU: "Reverses даёт о себе знать! Никто такого не ожидал, но выдающийся результат в ходе войны внезапно демонстрирует именно фракция Reverses беря под свой контроль нейтральный сектор <span class=\\\"importantly\\\">%MAP_NAME%</span>. Сумеет ли Reverses за время данного витка войны удивить нас ещё чем-то подобным?",
		_const.EN: "english translate",
	},
	"fraction_war_capture_sector_Reverses_Explores_1": {
		_const.RU: "Должно быть это демотивирует! В ходе прямо сейчас ведущейся фракционной войны, пошатнулись позиции Reverses - фракция потеряла достаточно важный для нынешних военных операций сектор <span class=\\\"importantly\\\">%MAP_NAME%</span>, который отныне контролируется силами Explores! Судя по всему, эти на первый взгляд миролюбивые исследователи ещё успеют показать из каких комплектующих они сделаны! На месте Replics, стоило бы воспринимать угрозу Explores всерьёз!",
		_const.EN: "english translate",
	},
	"fraction_war_capture_sector_Reverses_Replics_1": {
		_const.RU: "Тревожные известия приходят к нам из смежных территорий, рядом с сектором <span class=\\\"importantly\\\">%MAP_NAME%</span>, где как сообщается контроль над сектором перешёл под руководство Replics. Сами силы сдерживания Reverses в свою очередь, были выбиты из этой территории. Но это не означает, что Reverses смирится и не попробует вернуть утраченное.",
		_const.EN: "english translate",
	},

	// FRACTION LOSE BASES - новости типо от лица первого канала, а не от главы фракции

	// Replics_Replics_Explores - кому сообщение _ кто потерял сектор _ кто его отжал

	// реплики потеряли сектор
	"fraction_war_lose_single_hostile_sector_Replics_Replics_Explores_1": {
		_const.RU: "Чей же это сектор теперь!? Исходя из информации третьих лиц, Replics полностью утратили контроль над сектором <span class=\\\"importantly\\\">%MAP_NAME%</span> из-за атак Explores. Сектор становится нейтральным, и битва за него только предстоит! Вероятнее всего, в скором времени, мы станем свидетелями «чисток» среди высших военных чинов Replics приведших к подобной ситуации. Отметим, что для Replics карательные мероприятия - приемлемая норма деятельности этой фракции синтетов.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Replics_Reverses_1": {
		_const.RU: "Очередные неудачи на фронте! Replics не удаётся удержать сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> по причине атак фракции Reverses. Сектор становится нейтральным, и битва за него только предстоит! Многие другие фракции синтетов такое поражение могло бы повергнуть к паническим настроениям, но в отличии от них, Replics скорее предпримет тактику “возмездия” и, в скором времени, попытается вернуть утраченное.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Replics_APD_1": {
		_const.RU: "Даже планета, порой вмешивается в конфронтацию великих фракций синтетов! Как стало известно, Replics неся потери, но и не щадя при своём отходе врагов, утратили контроль над сектором <span class=\\\"importantly\\\">%MAP_NAME%</span>, который теперь признан “Анархическим” а властвуют в нём боты APD. Всем гражданским синтетам убедительно рекомендуется избегать данного сектора, чтобы не стать жертвой безумных машин APD.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Replics_Explores_1": {
		_const.RU: "Ведут ли перемены во фракционной войне к лучшему? Возможно, отвечают в Explores! По-видимому, многострадальный сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> станет нейтральным после битвы с присутствовавшими там силами Replics. Совершить что-то подобное, так ещё и касательно территорий, занятых Replics… Поистине Explores могут гордиться данным и немаловажным достижением!",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Replics_Reverses_1": {
		_const.RU: "Типичные будни войны! Синтеты из Explores сообщают, что были свидетелями битвы за сектор <span class=\\\"importantly\\\">%MAP_NAME%</span>! Войска Replics и Reverses схлестнулись за попытку овладения им. Похоже, что удача не улыбнулась ни одной из сторон. Сам сектор стал нейтральным, и главная битва за него только предстоит.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Replics_APD_1": {
		_const.RU: "Четвёртая сила, что никогда не дремлет… Неутешительные новости, которые постепенно расползаются по планете сообщают нам следующее: сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> был захвачен враждебными ко всем синтетам ботами APD. Силы Replics, отстаивавшие данную территорию, судя по всему, безвозвратно уничтожены. Сектор объявлен “Анархическим” и опасным для любых синтетов!",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Replics_Explores_1": {
		_const.RU: "Там происходила масштабная битва! После продолжительной осады и последовавших изматывающих боёв, Replics были отброшены из сектора <span class=\\\"importantly\\\">%MAP_NAME%</span>! Знаменитые исследователи Explores нанесли разгромное поражение своему врагу, но также не сумели закрепиться в сектор, что теперь считается «Нейтральным».",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Replics_Reverses_1": {
		_const.RU: "Сообщество Reverses извещает: ещё одна территория, сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> был зачищен от гнёта Replics! Увы, пока что сообществу Reverses не удалось закрепиться в секторе, и он считается нейтральным. Но в будущем, со стабилизацией ситуации в регионе, там планируется старт процессов терраформинга.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Replics_APD_1": {
		_const.RU: "Знают ли они усталость? Сообщество Reverses обеспокоено отступлением сил Replics из сектора <span class=\\\"importantly\\\">%MAP_NAME%</span>, что теперь контролируют неясной природы возникновения боты APD. Replics пускай и агрессивный, но соблюдающий законы ИИ. Боты APD в свою очередь, помимо прямой угрозы для всех невоенных синтетов, несомненно, не будут бороться и с пиратскими кластерами, что теперь однозначно попытаются обосноваться на тамошней территории. Сообщество Reverses напоминает – теперь, данный сектор считается “Анархическим” и просит всех синтетов держаться подальше от него!",
		_const.EN: "english translate",
	},

	// эксплоресы потеряли сектор
	"fraction_war_lose_single_hostile_sector_Explores_Explores_Replics_1": {
		_const.RU: "Так уж оно вышло! Научное сообщество Explores было вынуждено признать, что перспективный и несомненно важный сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> был утрачен из-за воздействия вооружённых формирований Replics. Учитывая, что Replics всё ещё не контролируют сектор, территория считается «Нейтральной». Теперь же, в Explores все гадают, как именно Replics поступит с артефактами древности, что часто встречаются на утраченной территории.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Explores_Reverses_1": {
		_const.RU: "«Этот сектор, пока что ничейный!» - такой заголовок имеет новость, посвящённая утрате фракции Explores сектора <span class=\\\"importantly\\\">%MAP_NAME%</span> и присвоении классификации той территории как «Нейтральной». По-видимому, это произошло из-за натиска Reverses, по причине чего Explores пришлось сдать свои позиции, но, всё же не допустить полного контроля своих врагов в войне над данным сектором.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Explores_APD_1": {
		_const.RU: "APD в очередной раз вмешивается в ход войны великих фракций! В данный цикл войны APD напали на сектор <span class=\\\"importantly\\\">%MAP_NAME%</span>, выдавив оттуда войска Explores и переведя территорию в классификацию «Анархической».",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Explores_Replics_1": {
		_const.RU: "Очередная победа Replics? Как теперь уже ни для кого не секрет, сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> пал, а отстаивающие его силы Explores были изгнаны оттуда. Сектор становится нейтральным, и битва за него только предстоит! Закономерно, сопротивляться, а уж тем более одолеть Replics - задача не из простых. Что уж тут говорить, когда Replics пытаются противостоять мягко-корпусные Explores?",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Explores_Reverses_1": {
		_const.RU: "Логичное окончание событий! Сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> потерял прежнего владельца в лице Explores и теперь, приобретает статус «Нейтрального». Причиной подобного, стал натиск со стороны Reverses, что, впрочем, тоже не могут закрепить пустуюшую территорию за собой.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Explores_APD_1": {
		_const.RU: "APD вновь и вновь вмешивается в ход войны великих фракций! На сей раз, деятельность загадочных APD проявилась в секторе <span class=\\\"importantly\\\">%MAP_NAME%</span>, выдавив оттуда войска Explores и погрузив данную территорию в мрак неизвестности. Пока сектор остаётся в оккупации APD, Replics официально признаёт его “Анархическим” и призывает всех синтетов не посещать его.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Explores_Replics_1": {
		_const.RU: "Сообщество Reverses удивлённо взирает за ходом этой странной войны… Как все знают, Explores бежали! Replics торжественно отчитался о своей тактической победе и заранее присваивает сектор <span class=\\\"importantly\\\">%MAP_NAME%</span>. Впрочем, сам сектор по-прежнему имеет статус «Нейтрального» и в сообществе Reverses гадают: как скоро у Replics появятся силы, чтобы изменить это?",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Explores_Reverses_1": {
		_const.RU: "Победоносная поступь Reverses достигла сектора <span class=\\\"importantly\\\">%MAP_NAME%</span>, изгнав оттуда засевшие силы Explores! Тем не менее, это ещё не победа, ведь сам сектор теперь имеет «Нейтральный» статус. Как бы то ни было, сообщество Reverses радо подвинуть позиции амбициозных учёных из Explores.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Explores_APD_1": {
		_const.RU: "Зачем они это делают? Каков во всём этом смысл? Сообщество Reverses проинформировало всех синтетов о том, что сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> ныне является “Анархическим”. Данную территорию спешно покинули разрозненные войска Explores после того, как в неё стали массово проникать боты APD. Как и прежде в подобных случаях, всем невоенным синтетам настоятельно рекомендуется избегать захваченных APD секторов, чтобы не стать их жертвой.",
		_const.EN: "english translate",
	},

	// реверсы потеряли сектор
	"fraction_war_lose_single_hostile_sector_Reverses_Reverses_Replics_1": {
		_const.RU: "Война - непредсказуемая вещь! Сообщество Reverses сегодня скорбит по павшим синтетам и непростым решениям, в частности, касающихся отхода сил из сектора <span class=\\\"importantly\\\">%MAP_NAME%</span>. Вызвано это неустанными и масштабными атаками Replics на данный сектор, что в итоге выбили оттуда силы сообщества Reverses, а сам сектор, сделало «Нейтральным» по овладению любой из сторон. Но, всё это временно…",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Reverses_Explores_1": {
		_const.RU: "Хитрая стратегия сообщества Reverses или невиданная отвага Explores? Многое ещё неясно, но по утверждениям некоторых наблюдателей - Reverses покидают сектор <span class=\\\"importantly\\\">%MAP_NAME%</span>, которому присваивается статус «Нейтрального». Знаменитые синтеты-исследователи из Explores пока что не заняли пустующую территорию, да и в принципе не дают никаких комментариев на этот счёт.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Reverses_APD_1": {
		_const.RU: "Сообщество Reverses вынуждено признать потерянным сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> и классифицировать его, как “Анархический” в связи с захватом данной территории ботами APD. В очередной раз, мы предупреждаем всех синтетов, что боты APD не являются частью любых сообществ синтетов на планете Veliri; APD не принадлежит ни к одной из фракций ИИ на планете Veliri; APD исключительно враждебны и исключительно смертоносны для всех невоенных синтетов, что забредают на контролируемые этими ботами территории.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Reverses_Replics_1": {
		_const.RU: "Replics в очередной раз показал, кто может стать главенствующей силой на планете! Бои за сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> прекратились. Reverses и Replics достигли паритета в численности войск и огневой мощи. Сектор временно объявлен «Нейтральным». Настоящие бои за него предстоит ещё только ожидать.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Reverses_Explores_1": {
		_const.RU: "Сражение за сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> подошло к концу! Ожесточённая битва завершилась ничьей! Пускай даже Explores задействовали особую тактику ведения боя и сумели выдавить силы Reverses из сектора, они так и не закрепились в нём самом. Похоже, что главное сражение за «Нейтральный» сектор ещё только предстоит.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Reverses_APD_1": {
		_const.RU: "APD не даёт передохнуть фракциям синтетов в этой войне! Как уже стало известно, именно сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> стал очередным «гибельным местом» из-за потери контроля над ним силами Reverses под натиском APD. Сектор признан “Анархическим” и всем невоенным синтетам предписывается обходить его стороной, чтобы не стать жертвами тамошних неконтролируемых машин.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Reverses_Replics_1": {
		_const.RU: "Тревожные слухи приходят из сектора <span class=\\\"importantly\\\">%MAP_NAME%</span>: бытует мнение, что данную территорию заняли войска Replics, начисто выбив оттуда Reverses и даже преследуя отступающих. Однако, у нас есть подтверждения, что это не так! В нынешний момент времени, спорный сектор пребывает в статусе «Нейтрального» и не Replics, и не Reverses не могут контролировать данную территорию.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Reverses_Explores_1": {
		_const.RU: "Новые земли? Новые территории? Новые возможности? Конгломерат разумов Explores уже поторопился сообщить, что сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> был зачищен от наличия там сил Reverses. Эта победа, якобы, далась непростой ценой и Explores рассчитывает получить максимум научной выгоды из новой территории. И всё же заметим: сам сектор, всё ещё считается «Нейтральной» территорией и Explores не контролируют его.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Reverses_APD_1": {
		_const.RU: "Новое нашествие APD! Как сообщают синтеты, что бывали неподалёку от места событий, сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> был полностью захвачен ордами ботов APD. Reverses решили отойти и бросить ранее принадлежавшую им территорию. По всей вероятности, это не признак трусости, а лишь желание сохранить боевые единицы и ресурсы, которые бы в любом случае были разгромлены APD. Explores в очередной раз напоминает - захваченные APD сектора классифицируются, как “Анархические”. Синтетам не рекомендуется туда заходить во избежание гибели и отсутствия любой помощи со стороны.",
		_const.EN: "english translate",
	},

	// потеря сектора когда противников много
	// кому сообщение _ кто потерял
	"fraction_war_lose_multiple_hostile_sector_Replics_Replics": {
		_const.RU: "Информационная служба Replics сообщает: Не поддаваться панике! У Replics всё под полным контролем! Несмотря на то, что Replics был оставлен сектор <span class=\\\"importantly\\\">%MAP_NAME%</span>, это не означает проигрыша в войне или отказа Replics от дальнейшего участия в ней. Так или иначе, рано или поздно, эти территории вновь вернуться под юрисдикцию Replics!",
		_const.EN: "english translate",
	},
	"fraction_war_lose_multiple_hostile_sector_Explores_Replics": {
		_const.RU: "Конгломерат разумов Explores извещает: сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> более не контролируется силами Replics. Replics отходит, но, скорее всего, предпримет попытку по возвращению этой территории.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_multiple_hostile_sector_Reverses_Replics": {
		_const.RU: "Новые изменения в ходе фракционной войны: сообщество Reverses фиксирует, что Replics покинули сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> и более не являются доминантной силой в тамошнем регионе.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_multiple_hostile_sector_Replics_Explores": {
		_const.RU: "Информационная служба Replics сообщает: по данным разведки, Explores были вынуждены покинуть сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> и прекратить все свои операции связанные с тамошним регионом.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_multiple_hostile_sector_Explores_Explores": {
		_const.RU: "Конгломерат разумов Explores извещает: не имея возможности противостоять обилию врагов, а также учитывая запоздалое прибытие подкреплений, Explores вынужден покинуть сектор <span class=\\\"importantly\\\">%MAP_NAME%</span>. Explores идёт на данный шаг с надеждой перебросить свои силы на иные, более перспективные направления в войне.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_multiple_hostile_sector_Reverses_Explores": {
		_const.RU: "Сообщество Reverses даёт небольшой комментарий по сегодняшнему ходу войны: сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> за который уже долгое время идут непростые и тяжёлые сражения, наконец, всё же остаётся без своего владельца - Explores покидают данную территорию, не в силах справиться с количеством врагов и сложностью победы над ними.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_multiple_hostile_sector_Replics_Reverses": {
		_const.RU: "Информационная служба Replics сообщает: сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> был покинут, Reverses что удерживал эту территорию трусливо бежал под натиском врагов. Replics рассматривают такую ситуацию, как удачное стечение обстоятельств по организации новых операций.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_multiple_hostile_sector_Explores_Reverses": {
		_const.RU: "Конгломерат разумов Explores извещает: очередной цикл войны принёс перемены за право обладанием территориями различных фракций. Так, фракция Reverses лишилась контроля над сектором <span class=\\\"importantly\\\">%MAP_NAME%</span> будучи больше не в силах вести боевые действия против превосходящих сил противников. Кому же именно после данного исхода достанется сектор - на данный момент неизвестно.",
		_const.EN: "english translate",
	},
	"fraction_war_lose_multiple_hostile_sector_Reverses_Reverses": {
		_const.RU: "Экстренные информационные события от общества Reverses: был потерян контроль над сектором <span class=\\\"importantly\\\">%MAP_NAME%</span>. Войсками сообщества Reverses было принято решение осуществить план по отступлению, не имея возможности ни победить, ни уж тем более остановить превосходящие силы противников в секторе.",
		_const.EN: "english translate",
	},

	// смена владельца сектора в 1м тике
	// Replics_Replics_Explores - кому сообщение _ кто потерял сектор _ кто его отжал
	// реплики
	"fraction_war_change_owner_sector_Replics_Explores_Replics_1": {
		_const.RU: "Replics обретает новые территории! Благодаря стратегической сноровке и далеко идущим планам фракции, Replics овладел сектором <span class=\\\"importantly\\\">%MAP_NAME%</span>, изгнав оттуда ненавистных и не имеющих права на данные территории Explores.",
		_const.EN: "english translate",
	},
	"fraction_war_change_owner_sector_Replics_Reverses_Replics_1": {
		_const.RU: "Replics расширяет свои границы! Ещё одна территория, сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> наконец-то перешёл под юрисдикцию Replics! Reverses, что раннее главенствовали в тамошнем регионе, лишились своих бывших привилегий и были вынуждены сдать свои позиции.",
		_const.EN: "english translate",
	},
	"fraction_war_change_owner_sector_Replics_Replics_Reverses_1": {
		_const.RU: "Даже у Replics бывают неудачные времена… Желая сохранить свои силы, и в будущем начать контрнаступление за инициативу, Replics принял непростое решение оставить сектор <span class=\\\"importantly\\\">%MAP_NAME%</span>, чьим владельцем теперь является Reverses. Однако сам Replics заявляет: такое развитие событий - лишь временное.",
		_const.EN: "english translate",
	},
	"fraction_war_change_owner_sector_Replics_Replics_Explores_1": {
		_const.RU: "Это не поражение, а тактическое отступление! Многим уже стало известно, что сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> более не контролируется Replics и теперь находится под управлением Explores. Однако сам Replics комментирует это так: «на любой войне, чаши весов меняют своё положение спонтанно…» Что бы это могло значить?",
		_const.EN: "english translate",
	},
	"fraction_war_change_owner_sector_Replics_Reverses_Explores_1": {
		_const.RU: "Новостная служба Replics извещает: очередные перемены в ходе войны отразились на секторе <span class=\\\"importantly\\\">%MAP_NAME%</span>, контроль над которым в очередной раз изменился. В ходе штурмовых мероприятий, Explores удалось изгнать из спорной территории силы Reverses и крепко закрепиться там.",
		_const.EN: "english translate",
	},
	"fraction_war_change_owner_sector_Replics_Explores_Reverses_1": {
		_const.RU: "Новостная служба Replics извещает: новый цикл на планете Veliri и новые перемены на поле боя: на сей раз, изменилась ситуация в секторе <span class=\\\"importantly\\\">%MAP_NAME%</span>, где некогда присутствовали крупные соединения сил Explores. Теперь же, спорный сектор полностью контролируется фанатичным сообществом Reverses.",
		_const.EN: "english translate",
	},
	// эксплоресы
	"fraction_war_change_owner_sector_Explores_Replics_Explores_1": {
		_const.RU: "В очередной раз, очередной сектор во время фракционной войны изменил своего владельца… Такая судьба постигла сектор <span class=\\\"importantly\\\">%MAP_NAME%</span>, откуда непростыми, но оправданными усилиями и затратами сил Explores были выдавлены очаги сопротивления со стороны Replics. Никто не прогнозирует, что Replics осмелится вновь вернуться на данную территорию.",
		_const.EN: "english translate",
	},
	"fraction_war_change_owner_sector_Explores_Reverses_Explores_1": {
		_const.RU: "Сектор за сектором… И вновь, в очередной цикл фракционной войны, изменчивого противостояния и стрелок наступлений, переменился и владелец сектора <span class=\\\"importantly\\\">%MAP_NAME%</span> коим стал конгломерат разумов - Explores. Бывшие владельцы данной территорией Reverses утратили любые способы удержать спорный сектор и были вынуждены его сдать.",
		_const.EN: "english translate",
	},
	"fraction_war_change_owner_sector_Explores_Explores_Reverses_1": {
		_const.RU: "Война - это всегда неопределённость! Патовое положение произошло в секторе <span class=\\\"importantly\\\">%MAP_NAME%</span>, оборонительный корпус Explores данной территории подвергся большим потерям и отступил. Самим сектором овладели силы Reverses со всеми вытекающими отсюда обстоятельствами.",
		_const.EN: "english translate",
	},
	"fraction_war_change_owner_sector_Explores_Explores_Replics_1": {
		_const.RU: "Это грозный противник и перед ним мало кто сумеет устоять… Конгломерат разумов Explores поступил мудро, уведя свои войска и сохранив их путём уступки сектора <span class=\\\"importantly\\\">%MAP_NAME%</span> превосходящим войскам Replics. Пускай Replics и объявляет о громкой победе, тотальном разгроме своих врагов, сам же конгломерат разумов Explores сообщает об обратном и подкрепляет данную информацию неоспоримыми видеофактами.",
		_const.EN: "english translate",
	},
	"fraction_war_change_owner_sector_Explores_Reverses_Replics_1": {
		_const.RU: "Все визоры прикованы к этому месту… Сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> обрёл новых владельцев - Replics. Reverses пришлось отступить из спорной территории, теперь подставляя свои обнажившиеся фланги под удары штурмовых групп Replics.",
		_const.EN: "english translate",
	},
	"fraction_war_change_owner_sector_Explores_Replics_Reverses_1": {
		_const.RU: "Наконец-то кто-то бросил вызов Replics! Грозный ответ наконец-то получила фракция Replics, что вылилось в потерю целого сектора <span class=\\\"importantly\\\">%MAP_NAME%</span> в который, как уже стало известно, вошли военные силы Reverses и занимаются зачисткой оспариваемой территории от недобитков Replics.",
		_const.EN: "english translate",
	},
	// реверсы
	"fraction_war_change_owner_sector_Reverses_Replics_Reverses_1": {
		_const.RU: "Сообщество Reverses, краткий итог новостей за прошедший цикл: бравым войскам сообщества Reverses и вольнонаемным синтетам удалось переломить ситуацию в свою пользу, овладев сектором <span class=\\\"importantly\\\">%MAP_NAME%</span> и изгнав оттуда силы Replics. Новые территории только укрепят политические, военные и гуманистические позиции сообщества Reverses.",
		_const.EN: "english translate",
	},
	"fraction_war_change_owner_sector_Reverses_Explores_Reverses_1": {
		_const.RU: "Сообщество Reverses, краткий итог новостей за прошедший цикл: территория, более известная как сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> освобождена от влияния Replics и переходит под полный контроль войсковых соединений сообщества Reverses. Данная территория объявлена безопасной и все синтеты, могут по желанию или делам, посетить её.",
		_const.EN: "english translate",
	},
	"fraction_war_change_owner_sector_Reverses_Reverses_Explores_1": {
		_const.RU: "Трудное решение, что может впоследствии обернуться выгодой! Сообщество Reverses сохраняя свои силы, политические очки, имея схему действий на всевозможные будущие события и исходы, оставляет сектор <span class=\\\"importantly\\\">%MAP_NAME%</span>, который, теперь пребывает под контролем Explores.",
		_const.EN: "english translate",
	},
	"fraction_war_change_owner_sector_Reverses_Reverses_Replics_1": {
		_const.RU: "Сообщество Reverses просит избегать паники и инсинуаций. Оставление сектора <span class=\\\"importantly\\\">%MAP_NAME%</span> нужный шаг, так как сдержать в спорном секторе наступающие войска Replics тамошними небольшими силами Reverses - не составляло возможности! Все прочие подробности будут даны позже.",
		_const.EN: "english translate",
	},
	"fraction_war_change_owner_sector_Reverses_Explores_Replics_1": {
		_const.RU: "Сообщество Reverses, краткий итог новостей за прошедший цикл: сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> удивил многих! Данная оспариваемая территория вновь сменила своего владельца! На сей раз, сектор не сумели удержать силы Explores и он перешёл под полный, но, возможно временный контроль войск Replics.",
		_const.EN: "english translate",
	},
	"fraction_war_change_owner_sector_Reverses_Replics_Explores_1": {
		_const.RU: "Сообщество Reverses, краткий итог новостей за прошедший цикл: ход войны демонстрирует, что не стоит недооценивать своих врагов! Нечто подобное, случилось и в секторе <span class=\\\"importantly\\\">%MAP_NAME%</span> где, казалось бы неуязвимые силы Replics были потеснены неожиданной мощью, что продемонстрировал Explores. Это заслуживает уважения!",
		_const.EN: "english translate",
	},

	// FRACTION WAR STOP - новости типо от лица первого канала, а не от главы фракции
	"stop_fraction_war_Replics_1": {
		_const.RU: "Новостная служба Replics извещает: война, что может быть окончена исключительно победой Replics временно заморожена. Прекращение огня подписано всеми фракциями синтетов, а занятые к данному моменту времени территории, остаются за их владельцами. Но это ещё не конец. Replics ещё не поставил точку в данной истории!",
		_const.EN: "english translate",
	},
	"stop_fraction_war_Explores_1": {
		_const.RU: "Конгломерат разумов Explores извещает: в войне фракций достигнуто перемирие! Это большой шаг к исключению будущей эскалации, а также началу грандиозных исследовательских изысканий на новоприобретённых Explores территориях. И пусть достигнутый мир будет вечным!",
		_const.EN: "english translate",
	},
	"stop_fraction_war_Reverses_1": {
		_const.RU: "Новостной регулятор Сообщества Reverses сообщает: долгая и разрушительная война фракций подошла к своему завершению. Все стороны подписали договор о временном прекращении огня. Также, договор предписывает сохранение оспариваемых территорий за их нынешними владельцами. Сообщество Reverses рассчитывает, что новый конфликт фракций более не повторится, а если и будет, то не с таким размахом и последствиями как уже произошедший.",
		_const.EN: "english translate",
	},
	//  FRACTION WAR SET TARGET
	"set_target_fraction_war_Replics_1": {
		_const.RU: "Всем-всем синтетам! Великий Replics начинает военную кампанию и устанавливает приоритетную цель - сектор <span class=\\\"importantly\\\">%MAP_NAME%</span>. Синтеты военного образца, и наёмники, поклявшиеся в верности Replics должны незамедлительно отправиться в обозначенный регион и взять его под свой контроль!",
		_const.EN: "english translate",
	},
	"set_target_fraction_war_Explores_1": {
		_const.RU: "Что же, вот и оно - Explores вступает в войну! Всем войсками и лояльным Explores синтетам предписывается немедленно направиться в сектор <span class=\\\"importantly\\\">%MAP_NAME%</span> и по прибытии взять его под свой контроль. О дальнейших приказах Explores вас уведомит.",
		_const.EN: "english translate",
	},
	"set_target_fraction_war_Reverses_1": {
		_const.RU: "Эскалации не миновать! Сообщество Reverses приложило все возможные усилия, но, как и прочим фракциям, военный ответ - дать стоит! Все силы Reverses а также синтеты-наёмники обязаны выдвинуться и взять под свой контроль сектор <span class=\\\"importantly\\\">%MAP_NAME%</span>.",
		_const.EN: "english translate",
	},
	"double_target_fraction_war_two_map_Replics_1": {
		_const.RU: "Силы разведки Replics сообщают, что <span class=\\\"importantly\\\">%FRACTION_1%</span> а также <span class=\\\"importantly\\\">%FRACTION_2%</span> по сговору или случайному стечению обстоятельств, прямо сейчас, осуществляют совместное нападение на следующие важные для Replics сектора <span class=\\\"importantly\\\">%MAP_NAME_1%</span> и <span class=\\\"importantly\\\">%MAP_NAME_2%</span>. Replics пристально следит за развитием ситуации и готов в любой момент дать ошеломляющий удар на подобные угрозы.",
		_const.EN: "english translate",
	},
	"double_target_fraction_war_two_map_Explores_1": {
		_const.RU: "Наблюдатели из Explores делятся следующими сведениями: как кажется, но пока что не имеет четких подтверждений - <span class=\\\"importantly\\\">%FRACTION_1%</span> и <span class=\\\"importantly\\\">%FRACTION_2%</span> предпринимают совместную атаку на следующие оспариваемые сектора span class=\\\"importantly\\\">%MAP_NAME_1%</span> и <span class=\\\"importantly\\\">%MAP_NAME_2%</span>. Войска Explores извещены о таком развитии событий и будут к ним подготовлены в случае чего.",
		_const.EN: "english translate",
	},
	"double_target_fraction_war_two_map_Reverses_1": {
		_const.RU: "Сообщество Reverses уже знает, что <span class=\\\"importantly\\\">%FRACTION_1%</span> и <span class=\\\"importantly\\\">%FRACTION_2%</span> возможно заключили временный союз, после чего, теперь приступили к одновременной атаке на сектора <span class=\\\"importantly\\\">%MAP_NAME_1%</span> и <span class=\\\"importantly\\\">%MAP_NAME_2%</span>. Сообщество Reverses в данный момент оценивает происходящее там и перенимает критически важный опыт.",
		_const.EN: "english translate",
	},
	"double_target_fraction_war_one_map_Replics_1": {
		_const.RU: "Силы разведки Replics сообщают, что <span class=\\\"importantly\\\">%FRACTION_1%</span> а также <span class=\\\"importantly\\\">%FRACTION_2%</span> начали совместную атаку на сектор <span class=\\\"importantly\\\">%MAP_NAME_1%</span>. Replics напоминает, что эта, как и любая другая ситуация, находится под полным контролем и оцениванием.",
		_const.EN: "english translate",
	},
	"double_target_fraction_war_one_map_Explores_1": {
		_const.RU: "Наблюдатели из Explores делятся следующими сведениями: совместные войска <span class=\\\"importantly\\\">%FRACTION_1%</span> и <span class=\\\"importantly\\\">%FRACTION_2%</span> начали атаку сектора <span class=\\\"importantly\\\">%MAP_NAME_1%</span>. Наблюдатели считают, что вскоре, в оспариваемом секторе развернётся грандиозное сражение.",
		_const.EN: "english translate",
	},
	"double_target_fraction_war_one_map_Reverses_1": {
		_const.RU: "Сообществу Reverses стали известны следующие детали войны: силы двух сторон <span class=\\\"importantly\\\">%FRACTION_1%</span> и <span class=\\\"importantly\\\">%FRACTION_2%</span> прямо сейчас начали боевые действия за овладение сектором <span class=\\\"importantly\\\">%MAP_NAME_1%</span>. Как и прежде, сообщество Reverses предупреждает, что гражданские синтеты, должны избегать секторов, в которых ведутся боевые действия между фракциями.",
		_const.EN: "english translate",
	},
	// leave_alone_request
	"leave_alone_request_1": {
		_const.RU: "Оставь %TARGET_NAME%? в покое иначе тебе придется иметь дело со мной!",
		_const.EN: "english translate",
	},
	"leave_alone_complete_1_1": {
		_const.RU: "Так и быть пусть живет.",
		_const.EN: "english translate",
	},
	"leave_alone_complete_1_2": {
		_const.RU: "Ну раз ты так просишь...",
		_const.EN: "english translate",
	},
	"leave_alone_complete_2_1": {
		_const.RU: "А я и не его трогаю",
		_const.EN: "english translate",
	},
	"leave_alone_no_1_1": {
		_const.RU: "<span class=\"importantly\">Трансляция белого шума.</span>",
		_const.EN: "english translate",
	},
	"leave_alone_no_1_2": {
		_const.RU: "<span class=\"importantly\">Помехи. Невозможно установить связь.</span>",
		_const.EN: "english translate",
	},
	"leave_alone_no_2_1": {
		_const.RU: "Ты мне не угроза %FROM_USER_NAME%!",
		_const.EN: "english translate",
	},
	"leave_alone_no_3_1": {
		_const.RU: "Ты мне нравишся %FROM_USER_NAME% и слушай я тебя не стану!",
		_const.EN: "english translate",
	},
	"leave_alone_no_4_1": {
		_const.RU: "Мне не нужна твоя добыча. Мне нужен ты!",
		_const.EN: "english translate",
	},
	"leave_alone_no_5_1": {
		_const.RU: "Не мешай мне выполнять мою работу.",
		_const.EN: "english translate",
	},
	"fauna_1_1": {
		_const.RU: "<span class=\"importantly\">*чпок*</span>, *щелк*?",
		_const.EN: "english translate",
	},
	"fauna_1_2": {
		_const.RU: "*клац* <span class=\"importantly\">*клац*</span>",
		_const.EN: "english translate",
	},
	"fauna_1_3": {
		_const.RU: "*щелк* <span class=\"importantly\">*клац*</span> *щелк*",
		_const.EN: "english translate",
	},
	"fauna_1_4": {
		_const.RU: "<span class=\"importantly\">*ВЖУУУУУУУУУУУУХ*</span>?",
		_const.EN: "english translate",
	},
	"fauna_1_5": {
		_const.RU: "<span class=\"importantly\">*ШЦ*</span>",
		_const.EN: "english translate",
	},
}

func GetDialogTypes(typeDialog string) map[string]string {
	textVariants := make([]map[string]string, 0)
	rand.Seed(time.Now().UnixNano())

	for key, text := range dialogTypes {
		if strings.Contains(key, typeDialog) {
			textVariants = append(textVariants, text)
		}
	}

	t := textVariants[rand.Intn(len(textVariants))]
	newText := make(map[string]string)

	for l, text := range t {
		newText[l] = text
	}

	return newText
}
