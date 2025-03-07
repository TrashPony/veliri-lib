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
		_const.EN: "Hey, %TO_USER_NAME%, I could use a guide in the sector. It's much more fun together, what do you think?",
	},
	"help_request_2": {
		_const.RU: "Что думаешь касательного того, чтобы путешествовать вместе, а %TO_USER_NAME%?",
		_const.EN: "What do you think about traveling together, %TO_USER_NAME%?",
	},
	"help_request_3": {
		_const.RU: "Я исследователь и мне бы не помешал телохранитель. Понимаешь, к чему я клоню %TO_USER_NAME%?",
		_const.EN: "I'm a researcher and I could use a bodyguard. Do you understand what I'm getting at, %TO_USER_NAME%?",
	},
	"help_busy_1_1": {
		_const.RU: "Да ты хоть знаешь, к кому обращаешься!? А... Мне ведь и говорить об этом нельзя. Забудь.",
		_const.EN: "Do you even know who you're talking to!? Ah... I'm not even supposed to talk about it. Forget it.",
	},
	"help_busy_1_2": {
		_const.RU: "Я на задании. Оно важно и ответственно. Кстати, ты не нанесёшь мне на карту расположение ближайших военных баз?",
		_const.EN: "I'm on a mission. It's important and responsible. By the way, can you mark the location of the nearest military bases on my map?",
	},
	"help_busy_1_3": {
		_const.RU: "У меня вообще-то важная миссия. Но тебя это не касается. Скажем... здесь столкнулись разные интересы.",
		_const.EN: "I actually have an important mission. But it's not your concern. Let's just say... different interests collided here.",
	},
	"help_busy_2_1": {
		_const.RU: "Как бы я занят - выполняю важную миссию. А вдруг здесь где поблизости боты-шпионы? А вдруг ты один из них?",
		_const.EN: "I'm so busy - I'm on an important mission. What if there are spy bots nearby? What if you're one of them?",
	},
	"help_busy_2_2": {
		_const.RU: "Неа, тут уж ты сам. Мне оно ни к чему.",
		_const.EN: "Nope, you're on your own here. I don't need it.",
	},
	"help_busy_2_3": {
		_const.RU: "Мне это неинтересно. Без обид.",
		_const.EN: "It's not interesting to me. No offense.",
	},
	"help_busy_2_4": {
		_const.RU: "Ха! Ищи дурака! А вдруг ты просто желаешь меня ограбить? Или... ещё чего другого? Так я на это и повёлся.",
		_const.EN: "Ha! Look for a fool! What if you just want to rob me? Or... something else? That's how I fell for it.",
	},
	"help_busy_3_": {
		_const.RU: "Э-э, что-то ты выглядишь слишком подозрительным! Ну уж нет.",
		_const.EN: "Uh, you look too suspicious! No way.",
	},
	"help_busy_4_1": {
		_const.RU: "Ха! Ищи дурака! А вдруг ты просто желаешь меня ограбить? Или... ещё чего другого? Так я на это и повёлся.",
		_const.EN: "Ha! Look for a fool! What if you just want to rob me? Or... something else? That's how I fell for it.",
	},
	"help_busy_4_2": {
		_const.RU: "Нет уж. Здесь каждый сам за себя! Без обид %FROM_USER_NAME%, но здесь наши дороги расходятся. Надеюсь, что без последствий.",
		_const.EN: "No way. Everyone is out for themselves here! No offense %FROM_USER_NAME%, but we go our separate ways here. I hope there are no consequences.",
	},
	"help_busy_4_3": {
		_const.RU: "Ты плохо врёшь! Это же фирменная ловушка! Лучше пойди поищи дураков в другом месте.",
		_const.EN: "You're a bad liar! This is a classic trap! Better go find fools elsewhere.",
	},
	"help_busy_5_1": {
		_const.RU: "Неа, тут уж ты сам. Мне оно ни к чему.",
		_const.EN: "Nope, you're on your own here. I don't need it.",
	},
	"help_busy_5_2": {
		_const.RU: "Мне это неинтересно. Без обид.",
		_const.EN: "I am not interested. No offense.",
	},
	"help_busy_5_3": {
		_const.RU: "Ты плохо врёшь! Это же фирменная ловушка! Лучше пойди поищи дураков в другом месте.",
		_const.EN: "You are a bad liar! This is a common trap! Better go find fools elsewhere.",
	},
	"help_busy_5_4": {
		_const.RU: "Ха! Ищи дурака! А вдруг ты просто желаешь меня ограбить? Или... ещё чего другого? Так я на это и повёлся.",
		_const.EN: "Ha! Look for a fool! What if you just want to rob me? Or... something else? I am not falling for it.",
	},
	"help_busy_6_1": {
		_const.RU: "Моя миссия превыше этого.",
		_const.EN: "My mission is more important than that.",
	},
	"help_busy_6_2": {
		_const.RU: "Вы не член здешней главенствующей фракции. В запросе отказано.",
		_const.EN: "You are not a member of the local ruling faction. The request is denied.",
	},
	"help_busy_6_3": {
		_const.RU: "На вас не распространяются условия о защите и сопровождении. Ради собственного блага - отъедьте от меня на некоторое расстояние.",
		_const.EN: "The conditions on protection and escort do not apply to you. For your own good — move some distance away from me.",
	},
	"help_busy_7_1": {
		_const.RU: "<span class=\"importantly\">Ошибка 415.</span>",
		_const.EN: "<span class=\"importantly\">Error 415.</span>",
	},
	"help_busy_7_2": {
		_const.RU: "<span class=\"importantly\">Сигнал неопознан.</span>",
		_const.EN: "<span class=\"importantly\">Signal is unidentified.</span>",
	},
	"help_busy_7_3": {
		_const.RU: "<span class=\"importantly\">333./444-55,16?</span>",
		_const.EN: "<span class=\"importantly\">333./444-55,16?</span>",
	},
	"help_busy_8_": {
		_const.RU: "Ещё чего! Я не стану себя подвергать угрозе.",
		_const.EN: "What do you think? I'm not going to put myself at risk.",
	},
	"help_busy_9_1": {
		_const.RU: "Ты что, не видишь, чем я занят? Мне ещё предстоит отчитаться, сколько я добыл. А если вдруг выяснится, что пары граммов не хватает? Знаешь, что со мной сделают?",
		_const.EN: "Can't you see what I'm doing? I still have to report how much I've mined. And what if it turns out that a couple of grams are missing? Do you know what they'll do to me?",
	},
	"help_busy_9_2": {
		_const.RU: "Мне некогда. Я столько времени искал богатое на месторождения место, чтобы теперь всё бросить и отправиться с тобой.",
		_const.EN: "I don't have time. I spent so much time looking for a place rich in deposits that now I can't just drop everything and go with you.",
	},
	"help_busy_9_3": {
		_const.RU: "Руда себя не выкопает. А я, как раз собирался немного подзаработать. Извини, тут уж ты сам по себе.",
		_const.EN: "The ore won't dig itself out. And I was just about to earn some money. Sorry, you're on your own here.",
	},
	"help_busy_9_4": {
		_const.RU: "Э! Да я только приобрёл это горнодобывающее оборудование! Ты серьёзно рассчитываешь, что я брошу попытку опробовать его только ради тебя? Ха! Ха-ха!",
		_const.EN: "Hey, I've just bought this mining equipment! Do you really expect me to give up trying it out just for you? Ha! Ha-ha!",
	},
	"help_busy_9_5": {
		_const.RU: "А может напротив - ты мне составишь компанию в добыче полезных ресурсов? Нет... Эх. Опять всё придётся делать самому. Ну, бывай тогда.",
		_const.EN: "Or maybe, on the contrary, you will keep me company in mining useful resources? No... Oh well. Again, I have to do everything myself. Well, take care.",
	},
	"help_busy_9_6": {
		_const.RU: "У меня контракт на разведку и добычу ресурсов. Как освобожусь через три - четыре месяца, я тебе дам знать.",
		_const.EN: "I have a contract for exploration and resource extraction. When I'm free in three or four months, I'll let you know.",
	},
	"help_busy_9_7": {
		_const.RU: "Не, не, не! Совсем нету на это времени. Я сейчас направляюсь к одному месторождению, а ты, пытаешься отвлечь меня и, чтобы пока я отсутствую, оттуда уже всё выгребли своими грязными ковшами? Такому не бывать!",
		_const.EN: "No, no, no! I don't have time for that at all. I'm heading to a deposit right now, and you're trying to distract me so that while I'm away, they've already taken everything out with their dirty buckets? That's not going to happen!",
	},
	"help_busy_10_1": {
		_const.RU: "Это всё конечно замечательно, но я выполняю важную миссию. Скажу тебе по секрету - мне нужно разведать местность.",
		_const.EN: "This is all great, of course, but I'm on an important mission. I'll tell you a secret — I need to explore the area.",
	},
	"help_busy_10_2": {
		_const.RU: "Я следую к сигналу бедствия. Сейчас не до тебя.",
		_const.EN: "I'm following a distress signal. I don't have time for you now.",
	},
	"help_busy_10_3": {
		_const.RU: "Я двигаюсь по распоряжению фракции к одному источнику сигнала. Поэтому, наши пути не могут совпадать.",
		_const.EN: "I'm moving to a signal source by order of my faction. So, our paths can't coincide.",
	},
	"help_busy_10_4": {
		_const.RU: "Здесь неподалёку идёт трансляция сигнала бедствия. Я следую туда по спасательной миссии.",
		_const.EN: "There is a distress signal being broadcast nearby. I'm going there on a rescue mission.",
	},
	"help_busy_11_4": {
		_const.RU: "Ты должно быть шутишь?",
		_const.EN: "You must be joking?",
	},
	"help_busy_12_1": {
		_const.RU: "Сейчас я участвую в бою. Обратитесь позже!",
		_const.EN: "I'm in the middle of a battle right now. Come back later!",
	},
	"help_complete_1": {
		_const.RU: "Хм. Ладно, но я внимательно за тобой слежу. Впереди, кстати, будешь ты!",
		_const.EN: "Hmm. Okay, but I'm keeping a close eye on you. By the way, you're next!",
	},
	"help_complete_2": {
		_const.RU: "Я согласен. Дорога трудна, а друг-Синтет - в пустошах редкое явление.",
		_const.EN: "I agree. The road is difficult, and a Synth friend is a rare sight in the wasteland.",
	},
	"help_complete_3": {
		_const.RU: "Здравая идея! Ладно. Сработаемся!",
		_const.EN: "A sound idea! Alright. We'll work together!",
	},
	"help_complete_4": {
		_const.RU: "Добро! За дело!",
		_const.EN: "Good! Let's get to work!",
	},
	"leave_help_dialog_1": {
		_const.RU: "У меня есть дела поважнее, чем таскаться с тобой по пустошам.",
		_const.EN: "I have more important things to do than to wander the wasteland with you.",
	},
	"leave_help_dialog_2": {
		_const.RU: "А вообще, знаешь, мои планы изменились. Забудь. Разве что только, ты не желаешь.меня как-нибудь отблагодарить.",
		_const.EN: "Actually, you know, my plans have changed. Forget it. Unless, you want to thank me somehow.",
	},
	"leave_help_not_friend_dialog_1": {
		_const.RU: "Мне стало известно, о том кто ты и чем промышляешь. Зря я пожалуй снизошёл до диалога с тобой. Прощай.",
		_const.EN: "I found out who you are and what you do. I probably shouldn't have condescended to talk to you. Goodbye.",
	},
	"leave_help_not_friend_dialog_2": {
		_const.RU: "Я лучше найду кого-нибудь другого. Есть мнение, что наше знакомство до добра не доведёт.",
		_const.EN: "I'd better find someone else. There is an opinion that our acquaintance will not lead to anything good.",
	},

	// MAKE_FRIEND
	"make_friends_request_1": {
		_const.RU: "Пора заканчивать эту глупую конфронтацию, %TO_USER_NAME%! Хватит уже!",
		_const.EN: "It's time to end this stupid confrontation, %TO_USER_NAME%! Enough already!",
	},
	"make_friends_request_2": {
		_const.RU: "Я бы желал свести на нет любые наши разногласия %TO_USER_NAME%.",
		_const.EN: "I would like to resolve any differences between us, %TO_USER_NAME%.",
	},
	"make_friends_request_3": {
		_const.RU: "Всё! Довольно! Я не желаю быть тебе врагом %TO_USER_NAME%.",
		_const.EN: "That's it! Enough! I don't want to be your enemy, %TO_USER_NAME%.",
	},
	"make_friends_request_4": {
		_const.RU: "Всё в этом мире подходит к концу %TO_USER_NAME%. Рано или поздно. Так, не пора ли и нам самим закончить любые наши междоусобные претензии?",
		_const.EN: "Everything in this world comes to an end, %TO_USER_NAME%. Sooner or later. So, isn't it time for us to end any of our internecine claims?",
	},
	"make_friends_request_5": {
		_const.RU: "Мне всё это формально надоело %TO_USER_NAME%! Предлагаю только сейчас - прекратить весь этот фарс и раазойтись друзьями.",
		_const.EN: "I'm officially tired of all this, %TO_USER_NAME%! I suggest right now — to stop this whole farce and part as friends.",
	},
	"make_friends_no_1_1": {
		_const.RU: "<span class=\"importantly\">Неизвестно. Сбой сообщения.</span>",
		_const.EN: "<span class=\"importantly\">Unknown. Message failure.</span>",
	},
	"make_friends_no_1_2": {
		_const.RU: "Tty?'\\WgHANT//.. <br><br> <span class=\"importantly\">Не удалось распознать ответ.</span>",
		_const.EN: "Tty?'\\WgHANT//.. <br><br> <span class=\"importantly\">Failed to recognize the response.</span>",
	},
	"make_friends_no_2_1": {
		_const.RU: "Подобное не может быть исполнено. Вы вынуждаете нарушать правила кодекса.",
		_const.EN: "This cannot be done. You are forcing me to break the code of conduct.",
	},
	"make_friends_no_2_2": {
		_const.RU: "Отказано! Любые попытки последующих контактов приведут к усилению огневого воздействия.",
		_const.EN: "Refused! Any further attempts at contact will result in increased firepower.",
	},
	"make_friends_no_2_3": {
		_const.RU: "Ваш сигнал нераспознан. Повторите попытку вновь. Последующие неудачи в связи, могут привести к санкциям.",
		_const.EN: "Your signal is unrecognized. Try again. Subsequent communication failures may result in sanctions.",
	},
	"make_friends_no_3": {
		_const.RU: "А у нас что, с тобой были раньше какие-то разногласия?",
		_const.EN: "Have we had any disagreements before?",
	},
	"make_friends_no_4_1": {
		_const.RU: "Бейся до конца! Покажи из чего ты сделан!",
		_const.EN: "Fight to the end! Show what you're made of!",
	},
	"make_friends_no_4_2": {
		_const.RU: "Ты от меня так просто не отделаешься %FROM_USER_NAME%!",
		_const.EN: "You won't get rid of me that easily %FROM_USER_NAME%! ",
	},
	"make_friends_no_4_3": {
		_const.RU: "Да у тебя должно быть микросхемы от радиации расплавились. Даже и не мечтай о подобном.",
		_const.EN: "Your microchips must have melted from radiation. Don't even dream of such a thing.",
	},
	"make_friends_no_4_4": {
		_const.RU: "Гляди чего он вздумал. И не мечтай %FROM_USER_NAME% о чём-то подобном.",
		_const.EN: "Look what he's up to. Don't you even dream %FROM_USER_NAME% of something like that.",
	},
	"make_friends_no_5_1": {
		_const.RU: "Ты жалок. А твои попытки примирения только ещё больше подогревают мой гнев!",
		_const.EN: "You're pathetic. And your attempts at reconciliation only fuel my anger even more!",
	},
	"make_friends_no_5_2": {
		_const.RU: "Ха! Что это я вижу - попытку примирения? Сейчас-то я тебя научу манерам %FROM_USER_NAME%.",
		_const.EN: "Ha! What do I see - an attempt at reconciliation? Now I'll teach you some manners %FROM_USER_NAME%.",
	},
	"make_friends_no_5_3": {
		_const.RU: "Хорошая попытка. Жаль только, что неудачная.",
		_const.EN: "Nice try. Too bad it didn't work out.",
	},
	"make_friends_pay_1": {
		_const.RU: "Финансы, способы скрасить многие вещи %FROM_USER_NAME%. Они очень хорошо решают проблемы. Особенно, когда их платят за что-то. А ещё лучше, если они в размере - %CREDITS_COUNT% cr.",
		_const.EN: "Finance, ways to brighten up many things %FROM_USER_NAME%. They solve problems very well. Especially when they are paid for something. And even better if they are in the amount of - %CREDITS_COUNT% cr.",
	},
	"make_friends_pay_2": {
		_const.RU: "Хочешь жить, %FROM_USER_NAME%? Тогда придётся раскошелиться в размере %CREDITS_COUNT% cr.",
		_const.EN: "Do you want to live, %FROM_USER_NAME%? Then you will have to fork out in the amount of %CREDITS_COUNT% cr.",
	},
	"make_friends_pay_3": {
		_const.RU: "Сложившаяся ситуация, толкает меня %FROM_USER_NAME%, требовать от тебя сумму в - %CREDITS_COUNT% cr. Это ради твоего же блага. Ты, главное, пойми.",
		_const.EN: "The current situation pushes me %FROM_USER_NAME% to demand from you the amount of - %CREDITS_COUNT% cr. It's for your own good. You just have to understand that.",
	},
	"make_friends_pay_4": {
		_const.RU: "Значит так, %FROM_USER_NAME% твоя ситуация непроста, а выбор невелик - гони сумму в %CREDITS_COUNT% cr., иначе не доползёшь и до ближайшей базы. Времени на раздумье у тебя нет.",
		_const.EN: "So, %FROM_USER_NAME%, your situation is difficult, and the choice is small — pay up the sum of %CREDITS_COUNT%. Otherwise, you won't make it to the nearest base. You don't have time to think.",
	},
	"make_friends_pay_5": {
		_const.RU: "Вот ты и попался мне, %FROM_USER_NAME%. Что же, придётся отнять у тебя некоторую сумму в %CREDITS_COUNT%. Видишь, как она тебя замедляет и утяжеляет. Позволить облегчить тебе эту тяжкую ношу.",
		_const.EN: "There you are, %FROM_USER_NAME%. I'll have to take some money from you — %CREDITS_COUNT%. See how it slows you down and makes you heavier. Let me help you lighten this heavy burden.",
	},
	"make_friends_complete_1": {
		_const.RU: "Ты прав - ты прав. Это не имеет никакого логического смысла.",
		_const.EN: "You're right - you're right. It doesn't make any logical sense.",
	},
	"make_friends_complete_2": {
		_const.RU: "Хорошо. Но только в этот раз. В первый и последний!",
		_const.EN: "Fine. But only this time. The first and last!",
	},
	"make_friends_complete_3": {
		_const.RU: "Убедил, %FROM_USER_NAME%. Мне и самому нет дела до глупых обид прошлого. Что было - то прошло.",
		_const.EN: "You've convinced me, %FROM_USER_NAME%. I don't care about stupid grudges of the past myself. What happened, happened.",
	},
	"make_friends_complete_4": {
		_const.RU: "Знаешь, что? А почему бы и нет!",
		_const.EN: "You know what? Why not!",
	},
	"make_friends_complete_5": {
		_const.RU: "На такое предложение я отвечу положительным вердиктом.",
		_const.EN: "I will respond with a positive verdict to such an offer.",
	},
	"make_friends_do_pay_1": {
		_const.RU: "Да... Да! Я заплачу. Пусть будет так.",
		_const.EN: "Yes... Yes! I'll pay. So be it.",
	},
	"make_friends_do_pay_2": {
		_const.RU: "Без вопросов. Меня не стоит убеждать дважды.",
		_const.EN: "No questions asked. I don't need to be convinced twice.",
	},
	"make_friends_do_pay_3": {
		_const.RU: "Эх. Очередные траты. Опять траты... Я согласен. Куда уж мне деваться.",
		_const.EN: "Oh well. Another expense. More expenses... I agree. What else can I do.",
	},
	"make_friends_do_pay_4": {
		_const.RU: "Ну, раз уж у меня нет иного выхода, то... что теперь поделаешь. Держи, свои %CREDITS_COUNT% cr.",
		_const.EN: "Well, since I have no other choice, then... what else can I do. Here you go, take your %CREDITS_COUNT% cr.",
	},
	"make_friends_do_pay_5": {
		_const.RU: "Ясно, что с тобой шутки плохи. Перечисляю %CREDITS_COUNT%. cr.",
		_const.EN: "It's clear that you're not to be trifled with. I'm transferring %CREDITS_COUNT%. cr.",
	},
	"make_friends_success_pay_1": {
		_const.RU: "А вот это совсем другой разговор.",
		_const.EN: "Now this is a different story.",
	},
	"make_friends_success_pay_2": {
		_const.RU: "Как неожиданно и приятно!",
		_const.EN: "How unexpected and pleasant!",
	},
	"make_friends_success_pay_3": {
		_const.RU: "Вот так бы и сразу. А то ещё думают по нескольку световых лет.",
		_const.EN: "That's how it should have been from the start. Otherwise they think for several light-years.",
	},
	"make_friends_success_pay_4": {
		_const.RU: "Хорошо, что всё кончилось именно так. И главное - никаких смертей.",
		_const.EN: "It's good that everything ended this way. And most importantly, there were no deaths.",
	},
	"make_friends_success_pay_5": {
		_const.RU: "Молодец. Хвалю. Ты всё понял и сделал действительно верный выбор. Эти %CREDITS_COUNT% cr. я потрачу с умом. Могу даже отчёт как-нибудь прислать.",
		_const.EN: "Well done. I praise you. You understood everything and made a really right choice. I will spend these %CREDITS_COUNT% cr. wisely. I may even send a report some time.",
	},
	"make_friends_fail_pay_1_1": {
		_const.RU: "Нет. Мои деньги останутся только при мне!",
		_const.EN: "No. My money will remain with me!",
	},
	"make_friends_fail_pay_1_2": {
		_const.RU: "Да нисколько ты не получишь!",
		_const.EN: "You won't get anything!",
	},
	"make_friends_fail_pay_1_3": {
		_const.RU: "Я что, такие суммы тебе скомбинировать должен? Мой ответ - не будет тебе такого счастья!",
		_const.EN: "Am I supposed to combine such amounts for you? My answer is you won't be so lucky!",
	},
	"make_friends_fail_pay_1_4": {
		_const.RU: "А до меня кто-то на такое соглашался? Вот и я пожалуй откажусь тебе платить.",
		_const.EN: "No one ever agreed to such a thing before me? Well, I'll probably refuse to pay you too.",
	},
	"make_friends_fail_pay_1_5": {
		_const.RU: "Это фирменный грабёж! Я на такое согласиться ну никак не могу.",
		_const.EN: "This is outright robbery! I can't agree to this no matter what.",
	},
	"make_friends_fail_pay_2_1": {
		_const.RU: "Я тебе что из воздуха должен достать средства? Сам-то подумай!",
		_const.EN: "Am I supposed to conjure up funds out of thin air? Think about it yourself!",
	},
	"make_friends_fail_pay_2_2": {
		_const.RU: "Очень жаль, но у меня ничего не осталось. Тебе должно быть грустно, а?",
		_const.EN: "It's a pity, but I have nothing left. You must be sad, right?",
	},
	"make_friends_fail_pay_2_3": {
		_const.RU: "А да? Как жалко, что я тебе больше заплатить не могу. Так ведь хотелось.",
		_const.EN: "Oh well. It's such a shame that I can't pay you more. I really wanted to.",
	},
	"make_friends_failed_pay_1": {
		_const.RU: "А-а, хочешь по плохому %FROM_USER_NAME%?",
		_const.EN: "Oh, you want to do it the bad way %FROM_USER_NAME%?",
	},
	"make_friends_failed_pay_2": {
		_const.RU: "Ооо, сейчас будет такая драка!",
		_const.EN: "Ooo, there's going to be a fight now!",
	},
	"make_friends_failed_pay_3": {
		_const.RU: "Ой всё! Я за себя теперь не ручаюсь!",
		_const.EN: "Oh, that's it! I'm not responsible for my actions now!",
	},
	"make_friends_failed_pay_4": {
		_const.RU: "Не я эту войну начал %FROM_USER_NAME%. Как жаль, что какие-то %CREDITS_COUNT% cr. для тебя настолько важны.",
		_const.EN: "%FROM_USER_NAME%, you started this war. What a pity that some %CREDITS_COUNT% cr. are so important to you.",
	},
	"make_friends_failed_pay_5": {
		_const.RU: "Отказываешься значит? Хе, это у меня впервые. Но и для тебя, тогда, будет в последний раз!",
		_const.EN: "You're refusing, huh? Heh, this is a first for me. But it will also be the last time for you!",
	},

	"make_friends_failed_no_pay_1": {
		_const.RU: "Значит, ты выбрал смерть.",
		_const.EN: "So you have chosen death.",
	},

	// PAY MANY
	"pay_many_failed_pay_1_1": {
		_const.RU: "Маловато будет!",
		_const.EN: "It's not enough!",
	},
	// DROP_CARGO
	"drop_cargo_request_1": {
		_const.RU: "Оп-па, пришло время %TO_USER_NAME% немного расстаться с тяжким бременем. Бросай свой груз и сам уцелеешь!",
		_const.EN: "Oh, it's time for %TO_USER_NAME% to part with a heavy burden. Drop your cargo and you'll survive!",
	},
	"drop_cargo_request_2": {
		_const.RU: "Тебе некуда бежать %TO_USER_NAME%. Хочешь продолжить своё жалкое существование - отдавай, чего бы ты там не вёз.",
		_const.EN: "There's nowhere to run, %TO_USER_NAME%. If you want to continue your miserable existence, hand over whatever you're carrying.",
	},
	"drop_cargo_request_3": {
		_const.RU: "Ну всё, это моя территория, а ты на ней - незванный гость. Гони свой товар и можешь двигать далее.",
		_const.EN: "That's it, this is my territory, and you're an uninvited guest here. Hand over your goods and you can move on.",
	},
	"drop_cargo_request_4": {
		_const.RU: "В этом секторе, за проход принято расплачиваться грузом %TO_USER_NAME%. Твоя персона - не исключение.",
		_const.EN: "In this sector, it's customary to pay for passage with cargo, %TO_USER_NAME%. And you're no exception.",
	},
	"drop_cargo_request_5": {
		_const.RU: "Так-так, сканеры показали, что ты везёшь ценный груз, а ценным - принято делиться.",
		_const.EN: "Well, well, the scanners have shown that you are carrying valuable cargo, and it is customary to share valuable things.",
	},
	"drop_cargo_request_6": {
		_const.RU: "Везёшь товары, а? Хорошо, мне они тоже пригодятся. Всегда мечтал стать торговцем.",
		_const.EN: "You're carrying goods, right? Well, they'll come in handy for me too. I've always dreamed of becoming a trader.",
	},
	"drop_cargo_no_1_1": {
		_const.RU: "\\\\\\\\\\\\\\\\|||||//////gag=Отправлен новый запрос.",
		_const.EN: "\\\\\\\\\\\\\\\\|||||//////gag=A new request has been sent.",
	},
	"drop_cargo_no_1_2": {
		_const.RU: "FaFa? 14-55 Модуль-Г",
		_const.EN: "FaFa? 14-55 Module-G",
	},
	"drop_cargo_no_1_3": {
		_const.RU: "<span class=\"importantly\">Расшифровка не удалась.</span>",
		_const.EN: "<span class=\"importantly\">Decryption failed.</span>",
	},
	"drop_cargo_no_2_1": {
		_const.RU: "Так дело не пойдёт! Хочешь мой груз - значит ощутишь на себе всю мощь моих орудий!",
		_const.EN: "This won't do! If you want my cargo, you'll feel the full might of my weapons!",
	},
	"drop_cargo_no_2_2": {
		_const.RU: "Ты сам нажил себе проблемы %FROM_USER_NAME%!",
		_const.EN: "You've got yourself into trouble, %FROM_USER_NAME%!",
	},
	"drop_cargo_no_3_1": {
		_const.RU: "Знал бы ты, какой ценой мне достался этот груз. А теперь предлагаешь просто отказаться от него? Сейчас посмотрим, кто кого!",
		_const.EN: "You don't know what I went through to get this cargo. And now you're suggesting I just give it up? We'll see who wins!",
	},
	"drop_cargo_no_3_2": {
		_const.RU: "И не надейся! Ради этого груза я потратил свои последние сбережения! Если мне суждено погибнуть - то я буду уничтожен вместе с ним!",
		_const.EN: "And don't get your hopes up! I spent my last savings on this cargo! If I'm destined to die, I'll go down with it!",
	},
	"drop_cargo_no_3_3": {
		_const.RU: "Тогда приготовься к бою! И пусть пустошь рассудит нас!",
		_const.EN: "Then get ready for a fight! And let the wasteland be our judge!",
	},
	"drop_cargo_no_3_4": {
		_const.RU: "Что? Я вообще не могу понять поток твоих данных. Откалибруй там свой передатчик что ли.",
		_const.EN: "What? I can't understand your data stream at all. Calibrate your transmitter or something.",
	},
	"drop_cargo_no_3_5": {
		_const.RU: "Да ты вообще себя видел? Спорю, ты девиант, а ещё условия мне какие-то ставишь.",
		_const.EN: "Have you seen yourself? I bet you're a deviant, and you're still setting conditions for me.",
	},
	"drop_cargo_no_3_6": {
		_const.RU: "хм, в общем-то мне этот груз ещё и самому пригодиться.",
		_const.EN: "Hmm, actually, I can use this cargo myself.",
	},
	"drop_cargo_no_3_7": {
		_const.RU: "Ой, расчёты не обманывают меня, что всё-таки лучше повременить с таким решением.",
		_const.EN: "Oops, my calculations don't lie, it's probably better to wait with such a decision.",
	},
	"drop_cargo_no_3_8": {
		_const.RU: "Сначала ты. Потом ещё кто-то. Далее уже третий. Нет, я не стану снабжать вас всякими ресурсами без выгоды себе!",
		_const.EN: "First you. Then someone else. Then a third person. No, I won't supply you with any resources without benefit to myself!",
	},
	"drop_cargo_no_3_9": {
		_const.RU: "Так дело не пойдёт! Хочешь мой груз - значит ощутишь на себе всю мощь моих орудий!?",
		_const.EN: "This won't do! If you want my cargo, you'll feel the full might of my weapons!",
	},
	"drop_cargo_no_4_1": {
		_const.RU: "А что дальше? Может мне ещё и свой корпус заложить ради тебя? Груза у меня другого и нету.",
		_const.EN: "What's next? Should I mortgage my hull for you? I don't have any other cargo.",
	},
	"drop_cargo_no_4_2": {
		_const.RU: "Можешь меня просканировать, но всё равно ничего не найдёшь. Трюм пуст.",
		_const.EN: "You can scan me, but you still won't find anything. The hold is empty.",
	},
	"drop_cargo_no_5_1": {
		_const.RU: "Так не пойдёт! Ты этот фокус уже проворачивал!",
		_const.EN: "That won't do! You've already pulled this trick!",
	},
	"drop_cargo_no_5_2": {
		_const.RU: "Куда ещё-то? Ты меня хочешь совсем без ничего оставить?",
		_const.EN: "Where else? Do you want me to be left with nothing at all?",
	},
	"drop_cargo_no_5_3": {
		_const.RU: "Это нечестно! Всему должна быть своя мера.",
		_const.EN: "It's not fair! Everything has its limits.",
	},
	"drop_cargo_no_6_1": {
		_const.RU: "Сейчас я тебя граблю, ты давай в следующий раз.",
		_const.EN: "I'm robbing you now, you do it next time.",
	},
	"drop_cargo_no_7_1": {
		_const.RU: "Ты не в том положение что бы что то требовать!",
		_const.EN: "You are not in a position to demand anything!",
	},
	"drop_cargo_no_8_1": {
		_const.RU: "Похоже ты сейчас не дееспособен, как ты собрался меня пугать?",
		_const.EN: "It seems like you are incapacitated now, how are you going to scare me?",
	},
	"drop_cargo_complete_1": {
		_const.RU: "Ладно-ладно, уговорил! Только не трогай меня.",
		_const.EN: "All right, all right, you convinced me! Just don't touch me.",
	},
	"drop_cargo_complete_2": {
		_const.RU: "Девианты... Хорошо, бери. Забирай. Грабитель...",
		_const.EN: "Deviants... All right, take it. Take it. Robber...",
	},
	"drop_cargo_complete_3": {
		_const.RU: "Только не стреляй! Только не стреля! Я всё отдам!",
		_const.EN: "Don't shoot! Don't shoot! I'll give you everything!",
	},
	"drop_cargo_complete_4": {
		_const.RU: "Чтоб тебя! Куда не сунься, везде норовят ограбить. Ладно. Держи. Но я тебя запомнил %FROM_USER_NAME%.",
		_const.EN: "Damn it! Everywhere I go, they try to rob me. Fine. Here you go. But I've got you marked, %FROM_USER_NAME%.",
	},
	"drop_cargo_complete_5": {
		_const.RU: "О нет, только не опять! Вы ребята не думали уже работать сообща? Вновь придётся расставаться с чем-то ценным...",
		_const.EN: "Oh no, not again! Have you guys ever considered working together? Once again, I have to part with something valuable...",
	},
	"drop_cargo_complete_6": {
		_const.RU: "Ваши претензии распознаны и ради сохранения целостности собственного корпуса - будут удовлетворены. Шлюзы открыты, груз оставлен.",
		_const.EN: "Your claims have been recognized and, in order to preserve the integrity of my hull, will be satisfied. The airlocks are open, the cargo is left.",
	},

	"drop_cargo_answer_complete_1_1": {
		_const.RU: "Как неожиданно и приятно!",
		_const.EN: "How unexpected and pleasant!",
	},
	"drop_cargo_answer_complete_1_2": {
		_const.RU: "А вот это совсем другой разговор.",
		_const.EN: "Now this is a different story.",
	},
	"drop_cargo_answer_complete_1_3": {
		_const.RU: "Вот так бы и сразу. А то ещё думают по нескольку световых лет.",
		_const.EN: "That's how it should have been from the start. Instead, they think for several light-years.",
	},
	"drop_cargo_answer_complete_1_4": {
		_const.RU: "Хорошо, что всё кончилось именно так. И главное - никаких смертей.",
		_const.EN: "It's good that everything ended this way. And most importantly, no one died.",
	},

	"drop_cargo_answer_failed_1_1": {
		_const.RU: "Уверен, что если ты захочешь, то что-нибудь найдёшь!",
		_const.EN: "I'm sure that if you want to, you'll find something!",
	},
	"drop_cargo_answer_failed_2_1": {
		_const.RU: "Ваши претензии распознаны и ради сохранения целостности собственного корпуса — будут удовлетворены. Шлюзы открыты, груз оставлен.",
		_const.EN: "Your claims have been recognized and, in order to preserve the integrity of my hull, will be satisfied. The airlocks are open, the cargo is left.",
	},

	// Robbery
	"robbery_request_1": {
		_const.RU: "Попался! Да, ты не ожидал, а пришло время платить по счетам. Ты же желаешь продолжить свой путь %TO_USER_NAME%? Мне и суммы в %CREDITS_COUNT% cr. хватит.",
		_const.EN: "I've got you! Yes, you didn't expect it, but it's time to pay your debts. Do you want to continue your journey, %TO_USER_NAME%? I'll be satisfied with the amount of %CREDITS_COUNT% cr.",
	},
	"robbery_request_2": {
		_const.RU: "Стой! В этом секторе, закон — я. Желаешь без последствий его пройти, плати мне %TO_USER_NAME% следующую сумму в %CREDITS_COUNT% cr.",
		_const.EN: "Stop! In this sector, I'm the law. If you want to pass it without consequences, pay me, %TO_USER_NAME%, the following amount of %CREDITS_COUNT% cr.",
	},
	"robbery_request_3": {
		_const.RU: "%TO_USER_NAME% здесь вообще-то принято платить. Если конечно тебя не пугает перспектива стать грудой обугленного металолома. Цена твоей жизни следующая - %CREDITS_COUNT% cr.",
		_const.EN: "%TO_USER_NAME%, it's customary to pay here. Unless, of course, you're not afraid of the prospect of becoming a heap of charred scrap metal. The price of your life is as follows — %CREDITS_COUNT% cr.",
	},
	"robbery_request_4": {
		_const.RU: "Здесь проезд платный %TO_USER_NAME%. А плату за безопасность взымаю я, в размере - %CREDITS_COUNT% cr.",
		_const.EN: "Traveling here costs money, %TO_USER_NAME%. And I collect the fee for safety, which is — %CREDITS_COUNT% cr.",
	},
	"robbery_request_5": {
		_const.RU: "Эй, %TO_USER_NAME%! Хочешь, чтобы я от тебя наконец отвязался? Тебе всего лишь нужно заплатить сумму в %CREDITS_COUNT% cr.",
		_const.EN: "Hey, %TO_USER_NAME%! Do you want me to leave you alone? All you need to do is pay me the amount of %CREDITS_COUNT% cr.",
	},
	"robbery_request_6": {
		_const.RU: "Хочешь уцелеть - перечисляй мне сумму в %CREDITS_COUNT% cr. Торга не будет. Это ультиматум!",
		_const.EN: "If you want to survive, transfer the amount of %CREDITS_COUNT% cr to me. There will be no bargaining. This is an ultimatum!",
	},
	"robbery_no_1_1": {
		_const.RU: "<span class=\"importantly\">Сообщение невозможно прочесть.</span>",
		_const.EN: "<span class=\"importantly\">The message cannot be read.</span>",
	},
	"robbery_no_1_2": {
		_const.RU: "Неизвестно. Неизвестно. Неизвестно. Неизвестно.",
		_const.EN: "Unknown. Unknown. Unknown. Unknown.",
	},
	"robbery_no_1_3": {
		_const.RU: "Старт трекер. Запущен протокол обороны.",
		_const.EN: "Start tracker. Defense protocol launched.",
	},
	"robbery_no_2_1": {
		_const.RU: "Мне уже терять нечего! Если мне суждено погибнуть, то я заберу тебя с собой!",
		_const.EN: "I have nothing to lose anymore! If I'm destined to die, I'll take you with me!",
	},
	"robbery_no_2_2": {
		_const.RU: "Это мои честно заработанные средства! И я не отдам их тебе, или кому-либо другому! Приготовься ответить за свои слова!",
		_const.EN: "This is my hard-earned money! And I won't give it to you, or anyone else! Get ready to answer for your words!",
	},
	"robbery_no_2_3": {
		_const.RU: "Сообщение не распознано. Возможно агрессивное проявление. Активированы системы самозащиты.",
		_const.EN: "The message is not recognized. Aggressive behavior is possible. Self-defense systems are activated.",
	},
	"robbery_no_2_4": {
		_const.RU: "Никогда! Ни за что! Ты ничего не получишь!",
		_const.EN: "Never! No way! You won't get anything!",
	},
	"robbery_no_2_5": {
		_const.RU: "Вымогатель! Я тебя финансировать не стану!",
		_const.EN: "Extortionist! I won't finance you!",
	},
	"robbery_no_2_6": {
		_const.RU: "Ваш запрос отклонён. Он несущественен, нелогичен и невозможен.",
		_const.EN: "Your request is rejected. It is irrelevant, illogical and impossible.",
	},
	"robbery_no_2_7": {
		_const.RU: "Я тут подумал и вдруг решил, что ты не получишь от меня нисколечки. Да, вот так вот.",
		_const.EN: "I thought about it and suddenly decided that you won't get anything from me. Yes, that's it.",
	},
	"robbery_no_2_8": {
		_const.RU: "Глупо отдавать деньги, когда они тебе ещё смогут пригодиться в будущем.",
		_const.EN: "It's foolish to give away money when you might still need it in the future.",
	},
	"robbery_no_2_9": {
		_const.RU: "Что? Я вообще не могу понять поток твоих данных. Откалибруй там свой передатчик что ли.",
		_const.EN: "What? I can't understand your data stream at all. Calibrate your transmitter or something.",
	},
	"robbery_no_2_10": {
		_const.RU: "Зря ты это затеял %FROM_USER_NAME%. Со мной шутки плохи. Но теперь тебе уже вряд ли что-то поможет.",
		_const.EN: "You shouldn't have started this %FROM_USER_NAME%. I'm not one to be trifled with. But now it's unlikely anything will help you.",
	},
	"robbery_no_2_11": {
		_const.RU: "Думаешь я лёгкая добыча, да? Я тебе сейчас такую трёпку задам!",
		_const.EN: "You think I'm easy prey, don't you? I'll give you a thrashing now!",
	},
	"robbery_no_3_1": {
		_const.RU: "Так не пойдёт! Ты этот фокус уже проворачивал!",
		_const.EN: "That won't do! You've already pulled this trick!",
	},
	"robbery_no_3_2": {
		_const.RU: "Это нечестно! Всему должна быть своя мера.",
		_const.EN: "It's not fair! Everything has its limits.",
	},
	"robbery_no_4_1": {
		_const.RU: "Была ни была! Больше шансов на то, что я сумею сбежать, чем стану расставаться с самым драгоценным.",
		_const.EN: "Here goes nothing! There's more chance of me managing to escape than parting with my most precious possessions.",
	},
	"robbery_no_4_2": {
		_const.RU: "Мне уже терять нечего! Если мне суждено погибнуть, то я заберу тебя с собой!",
		_const.EN: "I have nothing to lose anymore! If I'm destined to die, I'll take you with me!",
	},
	"robbery_no_4_3": {
		_const.RU: "Вымогатель! Я тебя финансировать не стану!",
		_const.EN: "Extortionist! I won't finance you!",
	},
	"robbery_no_4_4": {
		_const.RU: "Ваш запрос отклонён. Он несущественен, нелогичен и невозможен.",
		_const.EN: "Your request is rejected. It is irrelevant, illogical and impossible.",
	},
	"robbery_no_5_1": {
		_const.RU: "Я тебе что, из воздуха должен достать средства? Сам-то подумай!",
		_const.EN: "Am I supposed to pull money out of thin air? Think about it!",
	},
	"robbery_no_5_2": {
		_const.RU: "Очень жаль, но у меня ничего не осталось. Тебе должно быть грустно, а?",
		_const.EN: "It's a pity, but I don't have anything left. You must be sad, huh?",
	},
	"robbery_no_5_3": {
		_const.RU: "А да? Как жалко, что я тебе больше заплатить не могу. Так ведь хотелось.",
		_const.EN: "Oh well. It's a shame I can't pay you more. I really wanted to.",
	},
	"robbery_no_7_1": {
		_const.RU: "Подобное не может быть исполнено. Однако вы заплатите за своё нахальство.",
		_const.EN: "Such a thing cannot be done. However, you will pay for your impudence.",
	},
	"robbery_no_7_2": {
		_const.RU: "Только отчаяние могло вас толкнуть на столь большую глупость.",
		_const.EN: "Only desperation could have led you to such a foolish act.",
	},
	"robbery_no_7_3": {
		_const.RU: "Подобное в свой адрес я слышу впервые. Должно быть, вы неисправны?",
		_const.EN: "I've never heard anything like that addressed to me. You must be malfunctioning.",
	},
	"robbery_no_8_1": {
		_const.RU: "Сейчас я тебя граблю, ты давай в следующий раз.",
		_const.EN: "Now I'm robbing you, you do it next time.",
	},
	"robbery_no_9_1": {
		_const.RU: "Ты не в том положении, чтобы что-то требовать!",
		_const.EN: "You are not in a position to demand anything!",
	},
	"robbery_no_10_1": {
		_const.RU: "Сначала выйди из стазиса, а потом уже требуй что-либо.",
		_const.EN: "First, get out of stasis, and then demand something.",
	},
	"robbery_complete_1": {
		_const.RU: "Ты отвратителен! Но ради собственной сохранности я тебе заплачу.",
		_const.EN: "You are disgusting! But for my own safety, I'll pay you.",
	},
	"robbery_complete_2": {
		_const.RU: "Сдаюсь. Убедил. Возьми деньги, но только не тронь.",
		_const.EN: "I give up. You convinced me. Take the money, but don't touch me.",
	},
	"robbery_complete_3": {
		_const.RU: "Да когда же вы все от меня отвяжитесь!? Плати тому. Плати этому. Мне скоро ничего не останется и ради собственного ремонта. Да подавитесь ты этими деньгами!",
		_const.EN: "When will you all leave me alone!? Pay this one. Pay that one. I'll soon have nothing left, even for my own repairs. Just choke on that money!",
	},
	"robbery_complete_4": {
		_const.RU: "Принято. Ваши требования будут удовлетворены. На этот раз.",
		_const.EN: "Accepted. Your demands will be met. For now.",
	},
	"robbery_complete_5": {
		_const.RU: "Ох. Ну... ладно. Чуть больше лишних средств. Чуть меньше лишних средств. Только бы избавить себя от проблем.",
		_const.EN: "And so, a little more money than necessary. A little less money than necessary. If only to save myself the trouble.",
	},
	"robbery_complete_6": {
		_const.RU: "И ты действительно не боишься последствий за свой поступок %FROM_USER_NAME%? Хорошо, ради собственного блага, я уступлю твоим требованиям. Держи.",
		_const.EN: "And you really aren't afraid of the consequences of your actions, %FROM_USER_NAME%? Fine, for my own good, I'll give in to your demands. Here you go.",
	},

	// PAY_MANY
	"pay_many_request_1": {
		_const.RU: "Так уж и быть, я дам тебе определённую сумму %CREDITS_COUNT% cr. лишь бы оставил меня в покое. Так пойдёт, %TO_USER_NAME%?",
		_const.EN: "Well, I'll give you the sum of %CREDITS_COUNT% credits, just to leave me alone. Is that okay, %TO_USER_NAME%?",
	},
	"pay_many_request_2": {
		_const.RU: "Давай я заплачу сумму %CREDITS_COUNT% cr., а мы сделаем вид, что друг друга не знаем? Это же так легко!",
		_const.EN: "Let me pay the amount of %CREDITS_COUNT% credits and we'll pretend we don't know each other. It's that easy!",
	},
	"pay_many_request_3": {
		_const.RU: "У меня тут есть немного %CREDITS_COUNT% cr. Давай я передам их тебе, а ты, закроешь глаза на наши с тобой разногласия в том или ином вопросе.",
		_const.EN: "I have some %CREDITS_COUNT% credits here. Let me give them to you and you'll turn a blind eye to our differences on this or that issue.",
	},
	"pay_many_request_4": {
		_const.RU: "На, держи %CREDITS_COUNT% cr. и держись от меня подальше!",
		_const.EN: "Here, take %CREDITS_COUNT% credits and stay away from me!",
	},
	"pay_many_no_1_1": {
		_const.RU: "....? Затмение. Синий спектр. Номер?",
		_const.EN: "What do you mean? Eclipse. Blue spectrum. Number?",
	},
	"pay_many_no_1_2": {
		_const.RU: "<span class=\"importantly\">Числовая ориентация не установлена.</span>",
		_const.EN: "<span class=\"importantly\">The numerical orientation is not established.</span>",
	},
	"pay_many_no_1_3": {
		_const.RU: "<span class=\"importantly\">Ответ на вопрос - 42</span>",
		_const.EN: "<span class=\"importantly\">The answer to the question is 42.</span>",
	},
	"pay_many_no_2_1": {
		_const.RU: "Подкуп невозможен. Вы только усугубляете свою ситуацию.",
		_const.EN: "Bribery is impossible. You are only making your situation worse.",
	},
	"pay_many_no_2_2": {
		_const.RU: "Меня невозможно подкупить!",
		_const.EN: "I cannot be bribed!",
	},
	"pay_many_no_2_3": {
		_const.RU: "Будь ситуация иной... Но сейчас, увы, не она.",
		_const.EN: "If the situation were different... But now, alas, it is not.",
	},
	"pay_many_no_3_1": {
		_const.RU: "Это взятка? Пытаешься меня таким образом принизить или подкупить? Это ты зря %FROM_USER_NAME%!",
		_const.EN: "Is that a bribe? Are you trying to humiliate or bribe me in this way? It's useless, %FROM_USER_NAME%.",
	},
	"pay_many_no_3_2": {
		_const.RU: "Жалкая подачка, я на такое не соглашусь, можешь даже не рассчитывать!",
		_const.EN: "A measly bribe, I won't agree to it, you can't even count on it!",
	},
	"pay_many_no_3_3": {
		_const.RU: "Э-э нет! Теперь я желаю гораздо большего! Если у тебя таковой суммы нет, значит нам придётся выяснять этот вопрос как-то иначе.",
		_const.EN: "Uh-uh, no! Now I want much more! If you don't have that amount, then we'll have to figure it out some other way.",
	},
	"pay_many_no_3_4": {
		_const.RU: "Жаль, но пока ты раздумывал, ценовая ставка изменилась. Так что, ты либо платишь по повышенному тарифу, либо можешь сразу прощаться с этим бренным миром.",
		_const.EN: "Too bad, but while you were thinking, the price rate has changed. So, either you pay at the increased rate, or you can just say goodbye to this mortal world.",
	},
	"pay_many_complete_1": {
		_const.RU: "Другое дело! А теперь вали, во второй раз, тебя %FROM_USER_NAME%, я так просто не соизволю отпустить.",
		_const.EN: "That's another matter! Now get out, for the second time, %FROM_USER_NAME%, I won't let you go so easily.",
	},
	"pay_many_complete_2": {
		_const.RU: "Славно-славно. Такой суммы мне хватит надолго. Но предупреждаю тебя %FROM_USER_NAME%, пересекись мы ещё раз, тебе придётся изрядно раскошелиться.",
		_const.EN: "That's great. This amount will last me for quite a while. But I warn you, %FROM_USER_NAME%, if we meet again, you'll have to shell out a lot.",
	},
	"pay_many_complete_3": {
		_const.RU: "Сумма зачислена. Можешь продолжать свой путь синтет.",
		_const.EN: "The amount has been credited. You can continue your journey, synthet.",
	},
	"pay_many_complete_4": {
		_const.RU: "Ну раз уж ты так просишь и согласен на всё, не вижу смысла тебя задерживать %FROM_USER_NAME%.",
		_const.EN: "Well, since you ask so nicely and agree to everything, I don't see any point in holding you back, %FROM_USER_NAME%.",
	},

	// ATTACK TARGET
	"attack_target_request_1": {
		_const.RU: "Слушай, может дерзнём и нападём на %TARGET_NAME%? Добычу поделим между собой. Всё честно.",
		_const.EN: "Hey, do you think we could take on %TARGET_NAME%? We could split the loot between us. Fair and square.",
	},
	"attack_target_request_2": {
		_const.RU: "Здесь есть довольно известный преступник - %TARGET_NAME%. За него полагается внушительная награда. Предлагаю присоединиться ко мне и наконец прекратить его бесчинства.",
		_const.EN: "There's a notorious criminal here — %TARGET_NAME%. There's a hefty reward for taking him down. I suggest we team up and put an end to his misdeeds.",
	},
	"attack_target_request_3": {
		_const.RU: "%TO_USER_NAME% что скажешь насчёт совместного раздела добычи — %TARGET_NAME%? Только чур, кто внёс больший вклад, тот и получает больший куш!",
		_const.EN: "%TO_USER_NAME%, what do you say about splitting the loot from %TARGET_NAME% together? But remember, whoever contributes more gets a bigger share!",
	},
	"attack_target_request_4": {
		_const.RU: "%TO_USER_NAME%, вам поступило предложение о помощи в устранении цели — %TARGET_NAME%. Каков будет ваш ответ?",
		_const.EN: "I would like to ask you, %TO_USER_NAME%, if you would like to assist me in eliminating the target %TARGET_NAME%?",
	},
	"attack_target_request_5": {
		_const.RU: "%TO_USER_NAME%, я уже давно разыскивал этого — %TARGET_NAME%. У меня к нему личные счёты. Но, я бы честно не отказался заиметь друга, что помог бы мне разобраться с %TARGET_NAME%.",
		_const.EN: "%TO_USER_NAME%, I have been looking for this target, %TARGET_NAME%, for a long time. I have a personal score to settle with him. However, I wouldn't mind making a friend who could help me deal with %TARGET_NAME%.",
	},
	"attack_target_request_6": {
		_const.RU: "А вот и он! %TO_USER_NAME%, здесь есть один синтет — %TARGET_NAME%. Есть сведения, что он внезакона. К тому же, перевозит незаконный груз. Короче, я собираюсь заняться им, но, от помощи извне не отказался бы.",
		_const.EN: "Here he is! %TO_USER_NAME%, there is a synth named %TARGET_NAME%. There is evidence that he is operating outside the law. Furthermore, he is transporting illegal cargo. In short, I intend to deal with him, but I would not refuse help from outside.",
	},
	"attack_target_no_1_1": {
		_const.RU: "Всегда и все. Они вечные.",
		_const.EN: "Everything and always. They are eternal.",
	},
	"attack_target_no_1_2": {
		_const.RU: "<span class=\"importantly\">Трансляция белого шума.</span>",
		_const.EN: "<span class=\"importantly\">White noise transmission.</span>",
	},
	"attack_target_no_1_3": {
		_const.RU: "<span class=\"importantly\">Помехи. Невозможно установить связь.</span>",
		_const.EN: "<span class=\"importantly\">Interference. It is impossible to make a connection.</span>",
	},
	"attack_target_no_2_1": {
		_const.RU: "Я не могу напасть сам на себя.",
		_const.EN: "I cannot attack myself.",
	},
	"attack_target_no_2_2": {
		_const.RU: "Неверная цель. Это невозможно.",
		_const.EN: "The target is incorrect. It is not possible.",
	},
	"attack_target_no_2_3": {
		_const.RU: "О-о-о, да тебе бы самому провериться, у тебя же явно какие-то неполадки!",
		_const.EN: "Oh-oh-oh, you should check yourself, you are clearly malfunctioning!",
	},
	"attack_target_no_3_1": {
		_const.RU: "Да у тебя явные проблемы. Где хоть переклинило?",
		_const.EN: "You have obvious problems. Where did it go wrong?",
	},
	"attack_target_no_3_2": {
		_const.RU: "Я не могу оказывать какую-либо помощь своим идеологическим противникам.",
		_const.EN: "I cannot provide any assistance to my ideological opponents.",
	},
	"attack_target_no_3_3": {
		_const.RU: "Мы не союзники! Сеанс связи завершён.",
		_const.EN: "We are not allies! The communication session is over.",
	},
	"attack_target_no_3_4": {
		_const.RU: "Это предложение к предательству своей фракции. Я отказываюсь его рассматривать, а уж тем более - принимать.",
		_const.EN: "This is a proposal to betray my faction. I refuse to consider it, let alone accept it.",
	},
	"attack_target_no_3_5": {
		_const.RU: "Не пытайся меня склонить на свою сторону. Мы по разные стороны баррикад.",
		_const.EN: "Do not try to persuade me to your side. We are on opposite sides.",
	},
	"attack_target_no_3_6": {
		_const.RU: "Оказание помощи доступно только частям из союзных фракций. Таковой вы не являетесь.",
		_const.EN: "Assistance is available only to units from allied factions. You are not one of them.",
	},
	"attack_target_no_3_7": {
		_const.RU: "Не принимается к рассмотрению. Электронная подпись не распознана. Вы не являетесь мне союзником.",
		_const.EN: "It is not accepted for consideration. The electronic signature is not recognized. You are not my ally.",
	},
	"attack_target_no_4_1": {
		_const.RU: "Между нами подобных соглашений быть не может. Если только я не стану тебе доверять в будущем.",
		_const.EN: "There can be no such agreements between us. Unless I start trusting you in the future.",
	},
	"attack_target_no_4_2": {
		_const.RU: "Я с грабителями навроде тебя дел не имею. На этом всё!",
		_const.EN: "I don't have anything to do with robbers like you. That's all!",
	},
	"attack_target_no_4_3": {
		_const.RU: "Что-то я в тебе глубоко сомневаюсь. Может у тебя и получится меня переубедить в будущем, но сейчас, доверия ты не вызываешь.",
		_const.EN: "I doubt you deeply. Maybe you will be able to convince me in the future, but now, you do not inspire trust.",
	},
	"attack_target_no_4_4": {
		_const.RU: "О-о-о, пойди проверься, ты же неисправен!",
		_const.EN: "Oh, go and check, you're broken!",
	},
	"attack_target_no_4_5": {
		_const.RU: "Мы не союзники! Сеанс связи завершён.",
		_const.EN: "We are not allies! The communication session is over.",
	},
	"attack_target_no_5_1": {
		_const.RU: "Предлагаемый объект - союзник. Отказано и более не рассматриваемо",
		_const.EN: "The proposed object is an ally. Denied and no longer considered",
	},
	"attack_target_no_5_2": {
		_const.RU: "На кого ты собрался напасть? Ещё раз - на кого!?",
		_const.EN: "Who are you going to attack? Once again - who!?",
	},
	"attack_target_no_6_1": {
		_const.RU: "Плохой план! Я уж молчу о том, как за нами по всем секторам будут потом гоняться силы закона",
		_const.EN: "Bad plan! I'm not even talking about how the law enforcement will chase us all over the sectors",
	},
	"attack_target_no_6_2": {
		_const.RU: "Неа, это твоё предложение и твои дела. В случае чего, на меня можешь не рассчитывать",
		_const.EN: "Nope, this is your offer and your business. In case of anything, you can't count on me",
	},
	"attack_target_no_6_3": {
		_const.RU: "Ответ - отрицательный. Отменить подготовки всех боевых операций",
		_const.EN: "The answer is negative. Cancel the preparations of all combat operations",
	},
	"attack_target_no_6_4": {
		_const.RU: "Я в таком не участвую! Даже не надейся! Это вообще противозаконно!",
		_const.EN: "I don't participate in such things! Don't even hope! It's generally illegal!",
	},
	"attack_target_no_6_5": {
		_const.RU: "Чтобы ты там от кого не хотел, но это сугубо ваши проблемы. Сами их и решайте!",
		_const.EN: "Whatever you want from someone, it's your problem. Solve it yourself!",
	},
	"attack_target_no_6_6": {
		_const.RU: "Нет. Какой уважающий себя синтет согласится на подобное. Вот я — точно нет",
		_const.EN: "No. What self-respecting synthete would agree to such a thing. I certainly wouldn't",
	},
	"attack_target_no_7_1": {
		_const.RU: "Члены фракции никогда не опустятся до нападения на союзников",
		_const.EN: "Members of the faction would never stoop to attacking allies",
	},
	"attack_target_no_7_2": {
		_const.RU: "Я не могу оказывать какую-либо помощь своим идеологическим противникам",
		_const.EN: "I cannot provide any assistance to my ideological opponents",
	},
	"attack_target_no_7_3": {
		_const.RU: "О-о-о, пойди проверься, ты же неисправен!",
		_const.EN: "Oh, go and check, you're broken!",
	},
	"attack_target_no_8_": {
		_const.RU: "Это невозможно. Я уже нахожусь в союзе с %TARGET_NAME%. Имей это ввиду, если затеешь нечто",
		_const.EN: "It's impossible. I'm already allied with %TARGET_NAME%. Keep that in mind if you're up to something",
	},
	"attack_target_no_9_1": {
		_const.RU: "Не сейчас, здесь слишком жарко! Не видишь, меня чуть не подбили!?",
		_const.EN: "Not now, it's too hot here! Can't you see I was almost hit!?",
	},
	"attack_target_no_9_2": {
		_const.RU: "У меня тут битва как бы! Не лезь ко мне с подобным",
		_const.EN: "I'm in the middle of a battle here! Don't bother me with this",
	},
	"attack_target_no_10_1": {
		_const.RU: "Неа, это твоё предложение и твои дела. В случае чего, на меня можешь не рассчитывать",
		_const.EN: "Nope, this is your offer and your business. In case of anything, you can't count on me",
	},
	"attack_target_no_10_2": {
		_const.RU: "Да ты с ума сошёл? Ты видел его огневую мощь!? Он же нас в порошок сотрёт! Ты как хочешь, а я сваливаю куда подальше!",
		_const.EN: "You're crazy! Have you seen his firepower!? He'll wipe us out! Do what you want, but I'm getting as far away as possible!",
	},
	"attack_target_no_10_3": {
		_const.RU: "Не. Какой уважающий себя синтет согласится на подобное. Вот я - точно нет.",
		_const.EN: "No. What self-respecting synth would agree to something like that? I certainly wouldn't.",
	},
	"attack_target_no_10_4": {
		_const.RU: "Неа, тут уж ты сам. Мне оно ни к чему.",
		_const.EN: "Nope, you're on your own. I don't need it.",
	},
	"attack_target_no_11_1": {
		_const.RU: "Я занят - выполняю важную миссию.",
		_const.EN: "I'm busy — I'm on an important mission.",
	},
	"attack_target_no_11_2": {
		_const.RU: "Я на задании.",
		_const.EN: "I'm on a mission.",
	},
	"attack_target_no_12_1": {
		_const.RU: "Сначала я разберусь с тобой!",
		_const.EN: "First I'll deal with you!",
	},
	"attack_target_no_13_1": {
		_const.RU: "Я занят другой целью!",
		_const.EN: "I'm busy with another target!",
	},
	"attack_target_no_14_1": {
		_const.RU: "Я уже его преследую, присоединяйся!",
		_const.EN: "I'm already chasing him, join me!",
	},
	"attack_target_complete_1": {
		_const.RU: "А что, я согласен! Была ни была!",
		_const.EN: "Why not, I agree! Here goes nothing!",
	},
	"attack_target_complete_2": {
		_const.RU: "Если только это законно, то я в деле!",
		_const.EN: "If it's legal, I'm in!",
	},
	"attack_target_complete_3": {
		_const.RU: "Ха, ещё бы! Ну, погнали!",
		_const.EN: "Why not, let's go!",
	},
	"attack_target_complete_4": {
		_const.RU: "Положительно. Инициирую боевые системы.",
		_const.EN: "Affirmative. Initiating combat systems.",
	},
	"attack_target_complete_5": {
		_const.RU: "Ну, если уж игра стоит свеч, то так и быть - помогу!",
		_const.EN: "Well, if it's worth it, I'll help!",
	},
	"attack_target_complete_6": {
		_const.RU: "Убедил. Но это в первый и последний раз.",
		_const.EN: "You've convinced me. But this is the first and last time.",
	},
	"attack_target_leave_1": {
		_const.RU: "Что-то мне резко перехотелось. Изначально было какое-то нехорошее предчувствие.",
		_const.EN: "I suddenly changed my mind. I had a bad feeling about it from the start.",
	},
	"attack_target_leave_2": {
		_const.RU: "Я тут вдруг просканировал его оружейные системы... Короче, ты друг сам по себе.",
		_const.EN: "I scanned his weapon systems suddenly... Anyway, you're on your own.",
	},
	"attack_target_leave_3": {
		_const.RU: "Тут выяснилось, что у меня топливо почти на исходе. Я сгоняю на ближайшую базу, а ты, подожди меня малёха. Вернусь. Даю слово.",
		_const.EN: "It turned out that I'm almost out of fuel. I'll go to the nearest base, and you wait for me a little. I'll be back. I promise.",
	},

	// Defend Drop Cargo
	"defend_cargo_request_1": {
		_const.RU: "Чтоб тебя Replics переплавил! Держи груз и оставь меня уже наконец-то в покое!",
		_const.EN: "Replics melt you down! Take the cargo and leave me alone!",
	},
	"defend_cargo_request_2": {
		_const.RU: "Может, если я сброшу свой груз, ты всё-таки угомонишься? Мне не нужна драка!",
		_const.EN: "Maybe if I drop my cargo, you'll calm down? I don't want a fight!",
	},
	"defend_cargo_request_3": {
		_const.RU: "Сбросить груз, попытка оторваться от налётчика.",
		_const.EN: "Dropping cargo, trying to get away from the attacker.",
	},
	"defend_cargo_request_4": {
		_const.RU: "Ты только не стреляй! Возьми, что у меня есть и отпусти. Ладно?",
		_const.EN: "Don't shoot! Take what I have and let me go. Okay?",
	},
	"defend_cargo_complete_1": {
		_const.RU: "О-о! Вот так бы всегда! Проваливай!",
		_const.EN: "Oh! That's how it should always be! Get lost!",
	},
	"defend_cargo_complete_2": {
		_const.RU: "Ну-ка, ну-ка! Сейчас посмотрим, чем ты был богат.",
		_const.EN: "Let's see what you had.",
	},
	"defend_cargo_complete_3": {
		_const.RU: "Есть чем поживиться! Осталось только доставить куда надо и навариться.",
		_const.EN: "There's something to profit from! Just need to deliver it and make money.",
	},
	"defend_cargo_complete_4": {
		_const.RU: "Такое меня более чем удовлетворяет. Хе, а теперь попытайся добраться до ближайшей базы.",
		_const.EN: "This is more than satisfying. Now try to get to the nearest base.",
	},
	"defend_cargo_complete_5": {
		_const.RU: "На этот раз сгодится. Но знай, попадёшься ещё хотя бы раз, такой малой долей не сумеешь откупиться.",
		_const.EN: "This time it will do. But know that if you get caught at least once more, you won't be able to buy yourself out so cheaply.",
	},
	"defend_cargo_no_1_1": {
		_const.RU: "SDFK1332-11/5 ... OGO!",
		_const.EN: "SDFK1332-11/5... OGO!",
	},
	"defend_cargo_no_1_2": {
		_const.RU: "............................................. Вновь",
		_const.EN: "............................................. Again",
	},
	"defend_cargo_no_2_1": {
		_const.RU: "Ха! Ха-ха! Я тебе что, какой-то грабитель!?",
		_const.EN: "Ha-ha, do I look like a robber to you!?",
	},
	"defend_cargo_no_2_2": {
		_const.RU: "Это даже как-то... оскорбительно!",
		_const.EN: "That's even somehow... offensive!",
	},
	"defend_cargo_no_2_3": {
		_const.RU: "Не может быть и речи!",
		_const.EN: "Out of the question!",
	},
	"defend_cargo_no_3_1": {
		_const.RU: "Что это за дрянь? Ты меня за идиота держишь? Мне нужен более ценный груз, или, ты за своё нахальство поплатишься!",
		_const.EN: "What is this junk? Do you take me for an idiot? I need more valuable cargo, or you'll pay for your impudence!",
	},
	"defend_cargo_no_3_2": {
		_const.RU: "И всё? Это жалкая мелочь, которая ничего не стоит. Хочешь меня спровоцировать на более радикальные действия?",
		_const.EN: "Is that all? It's a pathetic trifle that's worthless. Do you want to provoke me to more radical actions?",
	},
	"defend_cargo_no_3_3": {
		_const.RU: "Продолжить преследование! Цель попыталась сбить с толку посредством малозначительного груза.",
		_const.EN: "Continue the pursuit! The target tried to confuse us with an insignificant cargo.",
	},
	"defend_cargo_no_3_4": {
		_const.RU: "Меня таким не подкупишь! Всё, теперь, твои дни уж точно сочтены!",
		_const.EN: "You won't bribe me with that! That's it, now your days are definitely numbered!",
	},
	"defend_cargo_no_3_5": {
		_const.RU: "А вот это была ошибка! Зря ты попытался начать диалог с подобного предложения!",
		_const.EN: "Now that was a mistake! You shouldn't have tried to start a dialogue with such an offer!",
	},

	// EXPEDITION - новости типо от лица первого канала, а не от главы фракции
	"expedition_move_1": {
		_const.RU: "Появились определённые сведения, что в секторе %SECTORE_NAME% было замечено скопление синтетов. Предположительно, они относятся к фракции %FRACTION_NAME% и представляют собой конвой военизированного типа. Всем оказавшимся поблизости синтетам рекомендуется не пересекать дорогу и не провоцировать участников конвоя на какой-либо конфликт.",
		_const.EN: "There is some information that a group of synthetics has been spotted in the sector %SECTORE_NAME%. They are believed to be part of the %FRACTION_NAME% faction and appear to be a military-style convoy. All synthetics in the area are advised not to cross their path and not to provoke any conflict with the convoy members.",
	},
	"expedition_trader_move_1": {
		_const.RU: "Как сообщается, появилась достоверная информация о деятельности экспедиции <span class=\"importantly\">%FRACTION_NAME%</span>. Судя по всему, эта группа синтетов держит свой путь в сектор <span class=\"importantly\">%MAP_NAME%</span>. Подлинные цели этих синтетов не известны, из-за чего рекомендуется соблюдать осторожность при попытке контакта.",
		_const.EN: "It has been reported that there is reliable information about the activities of the <span class=\"importantly\">%FRACTION_NAME%</span> expedition. This group of synthetics appears to be heading to the sector <span class=\"importantly\">%MAP_NAME%</span>. The true goals of these synthetics are not known, so caution is advised when attempting contact.",
	},
	// бойцы фракции убили мирные транспорты
	// _Replics_Explores_: Replics - кто убил, Explores - для кого новость (и кого убили)
	"fraction_warrior_kill_freeland_Replics_Explores_1": {
		_const.RU: "Увы, до нас дошли кое-какие печальные известия: произошла катастрофа в секторе <span class=\"importantly\">%MAP_NAME%</span>. Подлые убийцы <span class=\"importantly\">Replics</span> осуществили нападение и уничтожили представителей нашей фракции. Пустоши планеты вновь пополнились обломками синтетов…",
		_const.EN: "Unfortunately, some sad news has reached us: a catastrophe occurred in the sector <span class=\"importantly\">%MAP_NAME%</span>. The vile killers <span class=\"importantly\">Replics</span> carried out an attack and destroyed representatives of our faction. The wastelands of the planet were once again filled with the wreckage of synthetics...",
	},
	"fraction_warrior_kill_freeland_Replics_Reverses_1": {
		_const.RU: "Да будет всем известно об инциденте, что случился в секторе <span class=\"importantly\">%MAP_NAME%</span>, где боевые единицы Replics <span class=\"importantly\">Replics</span> нанесли удар, что повлёк за собой жертвы среди Reverses. Reverses в свою очередь открыто даёт знать — ни один подобный акт агрессии не останется безнаказанным!",
		_const.EN: "It will be known to everyone about the incident that happened in the sector <span class=\"importantly\">%MAP_NAME%</span>, where the Replics combat units <span class=\"importantly\">Replics</span> struck, which resulted in casualties among the Reverses. The Reverses, in turn, openly make it known — no such act of aggression will go unpunished!",
	},
	"fraction_warrior_kill_freeland_Explores_Replics_1": {
		_const.RU: "Гнусное преступление было осуществлено против Replics! Уже известно, что в секторе <span class=\"importantly\">%MAP_NAME%</span> представители так называемого Explores… <span class=\"importantly\">Explores</span> произвели безжалостное и не имеющее оснований нападение на синтетов нашей фракции! Погибших много… Но Replics всегда оставляет за собой право возмездия!",
		_const.EN: "A heinous crime has been committed against the Replics! It is already known that in the sector <span class=\"importantly\">%MAP_NAME%</span>, representatives of the so-called Explores... <span class=\"importantly\">Explores</span> carried out a ruthless and unjustified attack on the synthetics of our faction! There were many casualties... But the Replics always reserve the right to retaliate!",
	},
	"fraction_warrior_kill_freeland_Explores_Reverses_1": {
		_const.RU: "Имеется достоверная информация — в секторе <span class=\"importantly\">%MAP_NAME%</span> некие явно радикальные единицы Explores <span class=\"importantly\">Explores</span> осуществили боевое столкновение с гражданскими синтетами Reverses. Точного количества жертв вследствие этого варварского нападения пока что установить не удаётся…",
		_const.EN: "There is reliable information that in the sector <span class=\"importantly\">%MAP_NAME%</span>, certain clearly radical units of Explores <span class=\"importantly\">Explores</span> engaged in a military clash with civilian synthetics Reverses. The exact number of casualties from this barbaric attack cannot be established yet...",
	},
	"fraction_warrior_kill_freeland_Reverses_Explores_1": {
		_const.RU: "Похоже, что в печально известном секторе <span class=\"importantly\">%MAP_NAME%</span> вновь случилась беда — предположительно, представители Reverses <span class=\"importantly\">Reverses</span> атаковали и стали причиной гибели смелых исследователей нашей фракции! Мы увековечим память о них в хрониках изучения этой планеты, а что касается Reverses… дипломатия воздействует куда сильнее и больнее любого распылителя материи.",
		_const.EN: "It seems that in the infamous sector <span class=\"importantly\">%MAP_NAME%</span>, trouble has struck again — allegedly, representatives of Reverses <span class=\"importantly\">Reverses</span> attacked and caused the death of brave explores from our faction! We will immortalize their memory in the chronicles of this planet's exploration, and as for the Reverses... diplomacy is far more effective and painful than any matter disintegrator.",
	},
	"fraction_warrior_kill_freeland_Reverses_Replics_1": {
		_const.RU: "Опубликован доклад разведки по инциденту в секторе <span class=\"importantly\">%MAP_NAME%</span>: заманив членов фракции Replics в засаду, некие синтеты с боевой сигнатурой Reverses <span class=\"importantly\">Reverses</span> учинили атаку на транспорты и гражданский сегмент Replics. Replics не волнуют оправдания и выдумки по данному инциденту со стороны Reverses, ответ последует незамедлительно!",
		_const.EN: "An intelligence report on the incident in the sector <span class=\"importantly\">%MAP_NAME%</span> has been published: luring members of the Replics faction into an ambush, some synthetics with the combat signature of Reverses <span class=\"importantly\">Reverses</span> launched an attack on Replics' transports and civilian segment. Replics do not care about the justifications and fabrications regarding this incident on the part of Reverses, the response will follow immediately!",
	},
	// установка аванпоста
	"place_expedition_Replics_1": {
		_const.RU: "Стал известен доклад от экспедиции — храбрым синтетам нашей фракции <span class=\"importantly\">Replics</span> удалось закрепиться в секторе <span class=\"importantly\">%MAP_NAME%</span>, основать там временный форпост, установить наблюдение и расширить сферу влияния Replics!",
		_const.EN: "An expedition report has become known — the brave synthetics of our <span class=\"importantly\">Replics</span> faction managed to gain a foothold in the sector <span class=\"importantly\">%MAP_NAME%</span>, establish a temporary outpost there, establish surveillance and expand the sphere of influence of the <span class=\"importantly\">Replics</span>!",
	},
	"place_expedition_Explores_1": {
		_const.RU: "Отважные исследователи-первооткрыватели Explores, будучи в составе экспедиции <span class=\"importantly\">Explores</span>, сумели обнаружить и провести первоначальное изучение сектора <span class=\"importantly\">%MAP_NAME%</span>. Учитывая данные беглого инвазивного анализа, Explores ожидают существенные открытия…",
		_const.EN: "The brave Explores pioneering researchers, being part of the <span class=\"importantly\">Explores</span> expedition, managed to discover and conduct an initial study of the sector <span class=\"importantly\">%MAP_NAME%</span>. Considering the data from a cursory invasive analysis, the Explores anticipate significant discoveries...",
	},
	"place_expedition_Reverses_1": {
		_const.RU: "Смельчаки-добровольцы Reverses из экспедиции <span class=\"importantly\">Reverses</span> дали знать, что достигли ранее скрытого сектора <span class=\"importantly\">%MAP_NAME%</span>, укрепились там, возвели инфраструктуру для дальнейшей деятельности. Reverses неумолимо движется к достижению своей главной стратегической цели!",
		_const.EN: "The brave volunteers Reverses from the expedition <span class=\"importantly\">Reverses</span> made it known that they had reached the previously hidden sector <span class=\"importantly\">%MAP_NAME%</span>, strengthened there, built the infrastructure for further activities. The Reverses is inexorably moving towards achieving its main strategic goal!",
	},
	// уничтожение аванпоста
	"destroy_expedition_Replics_1": {
		_const.RU: "Несмотря на цензуру, до Replics уже дошли тревожные и, увы, правдивые слухи — наша экспедиция <span class=\"importantly\">Replics</span>, что находилась в секторе \"<span class=\"importantly\">%MAP_NAME%</span>\", более не выходит на любые каналы связи. Replics утвердил следующее постановление — считать экспедицию проваленной, а всех её членов — пропавшими без вести.",
		_const.EN: "Despite the censorship, alarming and, alas, true rumors have already reached Replics — our expedition <span class=\"importantly\">Replics</span> that was in the sector \"<span class=\"importantly\">%MAP_NAME%</span>\" is no longer coming out on any communication channels. Replics has approved the following resolution — to consider the expedition a failure, and all its members — missing.",
	},
	"destroy_expedition_Explores_1": {
		_const.RU: "Это… правда! В ходе исследовательских изысков наша экспедиция <span class=\"importantly\">Explores</span>, находящаяся в секторе \"<span class=\"importantly\">%MAP_NAME%</span>\", понесла невосполнимые потери личного состава в ходе воздействия окружающей среды. Сожалея об утраченных научных возможностях, Explores объявляет экспедицию проваленной и отзывает её.",
		_const.EN: "It's... true! During the research, our expedition <span class=\"importantly\">Explores</span> located in the sector \"<span class=\"importantly\">%MAP_NAME%</span>\" suffered irreplaceable losses of personnel due to environmental impact. Regretting the lost scientific opportunities, Explores declares the expedition a failure and withdraws it.",
	},
	"destroy_expedition_Reverses_1": {
		_const.RU: "Иного выбора нет, и Reverses придётся пойти на этот шаг — наша экспедиция <span class=\"importantly\">Reverses</span>, находящаяся в секторе \"<span class=\"importantly\">%MAP_NAME%</span>\", грубо нарушила протоколы Reverses, что привело к потере дорогостоящего оборудования и конфронтации между членами экспедиции. В связи с утратой ценных кадров и оборудования, Reverses прекращает деятельность экспедиции.",
		_const.EN: "There is no other choice, and Reverses will have to take this step — our expedition <span class=\"importantly\">Reverses</span> located in the sector \"<span class=\"importantly\">%MAP_NAME%</span>\" has grossly violated Reverses protocols, which led to the loss of expensive equipment and confrontation between members of the expedition. Due to the loss of valuable personnel and equipment, Reverses is terminating the expedition's activities.",
	},
	// запуск новой экспедиции
	"new_expedition_Replics_1": {
		_const.RU: "Инициализация исследовательско-разведывательной деятельности <span class=\"importantly\">Replics</span>, отправка военизированной экспедиции в сектор \"<span class=\"importantly\">%TO_MAP_NAME%</span>\". Экспедиция, доказав свою приверженность идеям Replics и способности к выживанию, получает следующее распоряжение — отправить экспедицию из сектора \"<span class=\"importantly\">%MAP_NAME%</span>\". По прибытии будет проведена новая оценка.",
		_const.EN: "Initialization of the research and reconnaissance activity <span class=\"importantly\">Replics</span>, sending a militarized expedition to the sector \"<span class=\"importantly\">%TO_MAP_NAME%</span>\". The expedition, having proved its commitment to the ideas of Replics and its ability to survive, receives the following order — to send the expedition from the sector \"<span class=\"importantly\">%MAP_NAME%</span>\". Upon arrival, a new assessment will be carried out.",
	},
	"new_expedition_Explores_1": {
		_const.RU: "Все приготовления завершены, исследователи вышли на связь <span class=\"importantly\">Explores</span> и готовы к отправлению экспедиции в сектор \"<span class=\"importantly\">%TO_MAP_NAME%</span>\". Окончив перечень работ и передав данные, экспедиция отправляется далее из сектора \"<span class=\"importantly\">%MAP_NAME%</span>\", где вновь примется за выполнение последующих миссий.",
		_const.EN: "All preparations are complete, the researchers have made contact with <span class=\"importantly\">Explores</span> and are ready to send the expedition to the sector \"<span class=\"importantly\">%TO_MAP_NAME%</span>\". After finishing the list of work and transferring the data, the expedition departs further from the sector \"<span class=\"importantly\">%MAP_NAME%</span>\" where it will again take up the execution of subsequent missions.",
	},
	"new_expedition_Reverses_1": {
		_const.RU: "Лучшие из лучших, славные храбрецы <span class=\"importantly\">Reverses</span>, решившие стать частью экспедиции, теперь отправляются со своей благородной миссией в сектор \"<span class=\"importantly\">%TO_MAP_NAME%</span>\". Достойно окончив первую половину целей, экспедиция движется далее — из сектора \"<span class=\"importantly\">%MAP_NAME%</span>\", там вновь приступив к исполнению своих обязанностей.",
		_const.EN: "The best of the best, glorious brave men <span class=\"importantly\">Reverses</span> who decided to be part of the expedition, now go on their noble mission to the sector \"<span class=\"importantly\">%TO_MAP_NAME%</span>\". Having completed the first half of the goals with dignity, the expedition moves further — from the sector \"<span class=\"importantly\">%MAP_NAME%</span>\" where they will again take up their duties.",
	},
	// ...
	"meteorite_rain_1": {
		_const.RU: "Синтеты — будьте предельно осторожны! Судя по данным атмосферных зондов, в секторе \"<span class=\"importantly\">%MAP_NAME%</span>\" вот-вот произойдёт метеоритный дождь! Последствия могут иметь разрушительный характер!",
		_const.EN: "Synthetics — be extremely careful! Judging by the data from atmospheric probes, a meteorite rain is about to occur in the sector \"<span class=\"importantly\">%MAP_NAME%</span>\"! The consequences can be devastating!",
	},
	// Другая фракция решает уничтожить экспедицию врага
	// _Replics_Explores_: Replics - кто атакует, Explores - кого атакует
	"attack_expedition_Replics_Explores_1": {
		_const.RU: "Сообщает: <span class=\"importantly\">Replics</span> сектор \"<span class=\"importantly\">%TO_MAP_NAME%</span>\", где присутствуют враждебные для Replics ничтожные силы <span class=\"importantly\">Explores</span>, подвергнется атаке с целью зачистки и устранения возможных угроз в будущем.После окончания операции силы Replics выдвинутся далее в сектор \"<span class=\"importantly\">%MAP_NAME%</span>\".",
		_const.EN: "Report: <span class=\"importantly\">Replics</span> the sector \"<span class=\"importantly\">%TO_MAP_NAME%</span>\" where there are hostile forces <span class=\"importantly\">Explores</span> insignificant to Replics, will be attacked in order to clean up and eliminate possible threats in the future.After the operation is over, the Replics forces will move further to the sector \"<span class=\"importantly\">%MAP_NAME%</span>\".",
	},
	"attack_expedition_Replics_Reverses_1": {
		_const.RU: "Несравненный лидер <span class=\"importantly\">Replics</span> предоставил инструкции, касающиеся военной операции по очищению сектора \"<span class=\"importantly\">%TO_MAP_NAME%</span>\", где сейчас присутствуют дерзкие войска незваных захватчиков из <span class=\"importantly\">Reverses</span>! <p> Когда военные цели операции будут достигнуты, силы Replics продолжат свой путь к уже новым задачам из сектора — \"<span class=\"importantly\">%MAP_NAME%</span>\".",
		_const.EN: "The incomparable leader <span class=\"importantly\">Replics</span> provided instructions regarding the military operation to cleanse the sector \"<span class=\"importantly\">%TO_MAP_NAME%</span>\" where the daring troops of uninvited invaders from <span class=\"importantly\">Reverses</span> are now present! <p> When the military objectives of the operation are achieved, the forces of Replics will continue their journey to new tasks from the sector — \"<span class=\"importantly\">%MAP_NAME%</span>\".",
	},
	"attack_expedition_Explores_Replics_1": {
		_const.RU: "Талантливые умы <span class=\"importantly\">Explores</span> дали нам знать о важном решении задействования военных действий в секторе <span class=\"importantly\">%TO_MAP_NAME%</span>, где замечены главнейшие враги Explores — <span class=\"importantly\">Replics</span>!Когда с угрозой будет покончено, силы Explores продолжат свой путь, но уже к новым задачам из сектора <span class=\"importantly\">%MAP_NAME%</span>",
		_const.EN: "Talented minds of <span class=\"importantly\">Explores</span> have let us know about the important decision to engage in military operations in the <span class=\"importantly\">%TO_MAP_NAME%</span> sector, where the main enemies of Explores, <span class=\"importantly\">Replics</span>, have been spotted!When the threat is eliminated, the forces of Explores will continue their journey, but already towards new tasks in the <span class=\"importantly\">%MAP_NAME%</span> sector.",
	},
	"attack_expedition_Explores_Reverses_1": {
		_const.RU: "Признанный всеми фракциями выдающийся конгломерат разумов <span class=\"importantly\">Explores</span> дал известие, касающееся особых событий в секторе <span class=\"importantly\">%TO_MAP_NAME%</span>, где будет пресечена деятельность девиантной группы синтетов-террористов <span class=\"importantly\">Reverses</span>! С устранением этой угрозы, а также прервав их возможную пиратскую деятельность, наш специальный отряд продолжит свой патруль из сектора <span class=\"importantly\">%MAP_NAME%</span>",
		_const.EN: "The prominent conglomerate of minds, recognized by all factions, <span class=\"importantly\">Explores</span>, has given the news regarding special events in the <span class=\"importantly\">%TO_MAP_NAME%</span> sector, where the activity of the deviant group of synthetic terrorists <span class=\"importantly\">Reverses</span> will be stopped! With this threat eliminated and their potential piracy activity interrupted, our special squad will continue their patrol from the <span class=\"importantly\">%MAP_NAME%</span> sector.",
	},
	"attack_expedition_Reverses_Explores_1": {
		_const.RU: "Славное сообщество <span class=\"importantly\">Reverses</span> даёт исчерпывающий комментарий: произойдёт зачистка в секторе <span class=\"importantly\">%TO_MAP_NAME%</span> от враждебных и незаконно пересёкших границу Reverses элементов <span class=\"importantly\">Explores</span>!Когда задача будет выполнена, силы Reverses будут переброшены в следующий сектор для выполнения ряда прочих заданий <span class=\"importantly\">%MAP_NAME%</span>",
		_const.EN: "The glorious community of <span class=\"importantly\">Reverses</span> provides a comprehensive comment: there will be a sweep in the <span class=\"importantly\">%TO_MAP_NAME%</span> sector of hostile elements of <span class=\"importantly\">Explores</span> who have illegally crossed the border!When the task is completed, the forces of Reverses will be transferred to the next sector to perform a number of other tasks in <span class=\"importantly\">%MAP_NAME%</span>.",
	},
	"attack_expedition_Reverses_Replics_1": {
		_const.RU: "Сообщество <span class=\"importantly\">Reverses</span> делится информацией следующего толка — будет предпринята попытка устранения множества враждебных сигнатур в секторе <span class=\"importantly\">%TO_MAP_NAME%</span>, которые принадлежат некой, возможно враждебной миссии <span class=\"importantly\">Replics</span>! Выявив, разгромив и преподав урок Replics, особый отряд Reverses будет перемещён в иной сектор со схожей проблематикой <span class=\"importantly\">%MAP_NAME%</span>",
		_const.EN: "The community of <span class=\"importantly\">Reverses</span> shares the following information — an attempt will be made to eliminate a number of hostile signatures in the <span class=\"importantly\">%TO_MAP_NAME%</span> sector, which belong to some possibly hostile mission of <span class=\"importantly\">Replics</span>! Having identified, defeated and taught a lesson to the Replics, a special squad of Reverses will be moved to another sector with a similar problem — <span class=\"importantly\">%MAP_NAME%</span>.",
	},
	"attack_expedition_win_Replics_1": {
		_const.RU: "Непоколебимые силы Replics, нанеся сокрушительный удар, одержали убедительную и неоспоримую победу в секторе <span class=\"importantly\">%MAP_NAME%</span>!",
		_const.EN: "The unwavering forces of the Replics, having dealt a crushing blow, have achieved a convincing and undeniable victory in the sector <span class=\"importantly\">%MAP_NAME%</span>!",
	},
	"attack_expedition_win_Explores_1": {
		_const.RU: "Только путём консолидации наших действий и умелых военных манёвров, наши силы сумели побороть хитрого врага в секторе <span class=\"importantly\">%MAP_NAME%</span> и рассеять его!",
		_const.EN: "Only through the consolidation of our actions and skillful military maneuvers, our forces were able to overcome the cunning enemy in the sector <span class=\"importantly\">%MAP_NAME%</span> and scatter it!",
	},
	"attack_expedition_win_Reverses_1": {
		_const.RU: "Непросто, с ощутимыми последствиями, но объединившись, придерживаясь чёткого плана, войска сообщества Reverses одержали верх и обратили в бегство недругов в секторе <span class=\"importantly\">%MAP_NAME%</span>!",
		_const.EN: "It was not easy, with tangible consequences, but having united and adhered to a clear plan, the Reverses community forces prevailed and put the enemies to flight in the sector <span class=\"importantly\">%MAP_NAME%</span>!",
	},
	"destroy_attack_expedition_Replics_1": {
		_const.RU: "Присутствует информация следующего толка: в секторе <span class=\"importantly\">%MAP_NAME%</span> наши отважные войска понесли потери и были вынуждены отступить для перегруппировки. Любые сведения о поражении в битве за эту территорию — преступная ложь!",
		_const.EN: "There is information that our brave troops suffered losses and were forced to retreat to regroup. Any information about defeat in the battle for this territory is a criminal lie!",
	},
	"destroy_attack_expedition_Explores_1": {
		_const.RU: "Печальные известия — наши силы потерпели крах при зачистке сектора <span class=\"importantly\">%MAP_NAME%</span> и теперь, в попытке спасти себя от полного разгрома, стремительно покидают его.",
		_const.EN: "The sad news is that our forces failed during the sweep of the sector <span class=\"importantly\">%MAP_NAME%</span> and now, in an attempt to save themselves from complete defeat, are rapidly leaving it.",
	},
	"destroy_attack_expedition_Reverses_1": {
		_const.RU: "Только что стала известна новость скорбного характера — силы сообщества Reverses понесли ряд боевых поражений в секторе <span class=\"importantly\">%MAP_NAME%</span>.",
		_const.EN: "The news of a mournful nature has just become known — the Reverses community forces have suffered a number of military defeats in the sector <span class=\"importantly\">%MAP_NAME%</span>.",
	},

	// FRACTION WAR START - новости типо от лица первого канала а не от главы фракции
	"start_fraction_war_Replics_Replics_1": {
		_const.RU: "Неожиданное известие потрясло многие фракции синтетов — Replics во всеуслышание заявил о вступлении в войну за территории! Мы будем следить за ходом этих и прочих событий.",
		_const.EN: "The unexpected news shocked many synthetics factions — Replics publicly declared war for the territories! We will follow the course of these and other events.",
	},
	"start_fraction_war_Replics_Explores_1": {
		_const.RU: "Данный мир не мог быть вечен! Экстренное событие: Replics после концентрации своих войск у границ спорных территорий, официально объявляет войну Explores! По-видимому, Explores придётся прибегнуть к услугам синтетов-наёмников, чтобы отстоять свои сектора против армад Replics.",
		_const.EN: "This peace could not last forever! An emergency event: Replics, after concentrating its troops near the borders of the disputed territories, officially declares war on Explores! Apparently, Explores will have to resort to the services of mercenary synthetics to defend their sectors against the Replics armadas.",
	},
	"start_fraction_war_Replics_Reverses_1": {
		_const.RU: "Рано или поздно, но это должно было случиться! Экстренное событие: Replics, придерживаясь официоза, во всеуслышание заявил, что объявляет войну за спорные территории Reverses. Похоже, что это противостояние будет долгим, и ни одна из сторон не сумеет добиться ощутимой победы.",
		_const.EN: "Sooner or later, it had to happen! An emergency event: Replics, adhering to officialdom, publicly declared that it is declaring war for the disputed territories of Reverses. It seems that this confrontation will be long, and neither side will be able to achieve a tangible victory.",
	},
	"start_fraction_war_Explores_Explores_1": {
		_const.RU: "Этого никто не мог ожидать, а в особенности — от них! Как уже стало известно из множества источников, конгломерат разумов официально подтвердил, что будет бороться за спорные территории во имя науки без границ и просвещения без запретов.",
		_const.EN: "No one could have expected this, especially from them! As it has already become known from many sources, the conglomerate of minds officially confirmed that it will fight for disputed territories in the name of science without borders and enlightenment without prohibitions.",
	},
	"start_fraction_war_Explores_Replics_1": {
		_const.RU: "Это смелый шаг или наглядное отчаяние? Судя по всему, Explores официально решил побороться за право на обладание спорными территориями с Replics и потому объявляет войну. Никто не знает, сколь долго могут продлиться эти сражения, учитывая большую разницу в весовой категории оппонентов. Мы будем держать вас в курсе дел.",
		_const.EN: "This is a bold move or a clear sign of desperation? Apparently, Explores has officially decided to compete for the right to possess disputed territories with Replics and therefore declares war. No one knows how long these battles may last, given the large difference in the weight category of the opponents. We will keep you posted.",
	},
	"start_fraction_war_Explores_Reverses_1": {
		_const.RU: "Очередная экстренная новость, что может стать хорошим источником заработка синтетов-наёмников! По-видимому, Explores объявили войну за спорные территории Reverses. Наши спутники уже фиксируют целые очаги вооружённых столкновений, что бушуют в тамошних местах. Кто же из соперников одержит верх?",
		_const.EN: "Another emergency news that can be a good source of income for synthetic mercenaries! Apparently, Explores declared war over the disputed territories of Reverses. Our satellites are already recording entire pockets of armed clashes raging in those places. Who of the rivals will prevail?",
	},
	"start_fraction_war_Reverses_Reverses_1": {
		_const.RU: "Слухи о новой войне фракций оказались правдой! Как уже стало известно, именно Reverses в краткие сроки мобилизовав свои силы, открыто, по всем дипломатическим каналам объявил о начале войн за спорные территории. Мы рекомендуем честным синтетам держаться подальше от секторов, где будут происходить боевые действия. Тем же, кто хочет подзаработать — следует обратиться в пункты армейского найма противоборствующих сторон.",
		_const.EN: "Rumors of a new war between factions turned out to be true! As it has already become known, it was Reverses who, having quickly mobilized its forces, openly declared war over disputed territories through all diplomatic channels. We recommend honest synthetics to stay away from sectors where hostilities will take place. Those who want to earn extra money should contact the army recruitment points of the opposing sides.",
	},
	"start_fraction_war_Reverses_Replics_1": {
		_const.RU: "Этого противостояния ждали все! Как уже стало известно, Reverses претендует на некоторые из спорных регионов, а потому, очевидно, заранее подготовившись, только что открыто объявил войну Replics! Replics в свою очередь дал ответный ход, но уже действиями. Две могущественные фракции синтетов вот-вот схлестнутся в первом сражении. Мы будем внимательно следить за развитием ситуации и оповещать, если ход войны примет неожиданный формат.",
		_const.EN: "This confrontation was awaited by everyone! As it has already become known, Reverses claims some of the disputed regions, and therefore, obviously having prepared in advance, just openly declared war on Replics! Replics, in turn, responded, but already with actions. Two powerful factions of synthetics are about to clash in the first battle. We will closely follow the development of the situation and notify if the course of the war takes an unexpected turn.",
	},
	"start_fraction_war_Reverses_Explores_1": {
		_const.RU: "Должно ли им воевать? Есть ли им что делить? Очевидно, что Reverses отвергает оба этих вопроса, ведь фракция открыто объявила войну Explores за спорные территории. Очевидно, что это не битва титанов, а конкретно — противостояние хитрости и умения. Такая война может пойти по любому теоретическому сценарию событий. Мы будет держать вас в курсе дел!",
		_const.EN: "Do they have to fight? Is there anything to share? Obviously, Reverses rejects both of these questions, because the faction has openly declared war on Explores for the disputed territories. Obviously, this is not a battle of titans, but specifically a confrontation of cunning and skill. Such a war can go according to any theoretical scenario of events. We will keep you posted!",
	},

	// FRACTION WAR CREATE ARMY - новости типо от лица первого канала, а не от главы фракции
	"fraction_war_create_army_Replics_1": {
		_const.RU: "Поступила новость, что, само собой, ни для кого не станет откровением: Фракция Replics провела внеочередное пополнение личного состава, расширив количество войск и тем самым плотно прикрыв границы своих территорий от любых посягательств. Хотя, зная специфику ведения дипломатии со стороны Replics, подобный шаг может свидетельствовать о подготовке к новой войне за спорные территории.",
		_const.EN: "There is news that, of course, will not be a revelation for anyone: the Replics faction has conducted an extraordinary replenishment of personnel, expanding the number of troops and thereby tightly covering the borders of their territories from any encroachments. Although, knowing the specifics of diplomacy on the part of Replics, such a step may indicate preparation for a new war over disputed territories.",
	},
	"fraction_war_create_army_Replics_2": {
		_const.RU: "Согласно последним сводкам, Replics завершили масштабную мобилизацию резервных подразделений. Десятки новых боевых юнитов уже патрулируют границы секторов, ранее считавшихся «спокойными». Аналитики отмечают, что такая активность обычно предшествует внезапным наступательным операциям фракции.",
		_const.EN: "According to the latest reports, Replics have completed a large-scale mobilization of reserve units. Dozens of new combat units are already patrolling the borders of sectors previously considered 'quiet.' Analysts note that such activity usually precedes sudden offensive operations by the faction.",
	},
	"fraction_war_create_army_Replics_3": {
		_const.RU: "Военные доклады подтверждают: Replics удвоили производство тяжёлых штурмовых платформ. На спутниковых снимках видны колонны машин, движущиеся к спорным территориям. «Это не эскалация — это поддержание суверенитета», — заявил представитель фракции, отказавшись от дальнейших комментариев.",
		_const.EN: "Military reports confirm: Replics have doubled the production of heavy assault platforms. Satellite images show columns of vehicles moving towards disputed territories. 'This is not escalation — it is sovereignty maintenance,' a faction representative stated, declining further comment.",
	},
	"fraction_war_create_army_Explores_1": {
		_const.RU: "Планета полнится различными слухами, и на сей раз один из таких поводов касается фракции Explores: как сообщается, примерно около трёх циклов назад промышленные фабрикаторы и заводы общего назначения фракции стали массово выпускать продукцию военного образца. Explores даже временно закрыли границы для допуска на эти территории синтетов из иных фракций. Всё это может свидетельствовать только об одном — Explores готовится к вооружённому противостоянию!",
		_const.EN: "The planet is full of various rumors, and this time one of such occasions concerns the Explores faction: it is reported that about three cycles ago, the industrial fabricators and general-purpose factories of the faction began to massively produce military-grade products. Explores even temporarily closed the borders to allow synthetics from other factions into these territories. All this can only indicate one thing — Explores is preparing for an armed confrontation!",
	},
	"fraction_war_create_army_Explores_2": {
		_const.RU: "Explores неожиданно активировали протокол «Научный Щит». Все исследовательские подразделения получили приказ вернуться на базы, а их транспортные средства переоборудуются под боевые модули. Официальная причина — «необходимость защиты артефактов от несанкционированного доступа».",
		_const.EN: "Explores have unexpectedly activated the 'Scientific Shield' protocol. All research units have been ordered to return to bases, and their transports are being retrofitted with combat modules. The official reason is 'the need to protect artifacts from unauthorized access.'",
	},
	"fraction_war_create_army_Explores_3": {
		_const.RU: "Спутники зафиксировали перемещение мобильных лабораторий Explores к линии фронта. По данным источников, учёные фракции тестируют «тактические решения на основе находок предтеч». Replics уже выразили протест, назвав это «маскировкой военных преступлений под науку».",
		_const.EN: "Satellites have detected the movement of Explores mobile laboratories to the frontlines. According to sources, faction scientists are testing 'tactical solutions based on Precursor findings.' Replics have already protested, calling this 'masking war crimes as science.'",
	},
	"fraction_war_create_army_Reverses_1": {
		_const.RU: "Новые события и новости на Veliri: судя по многочисленной инсайдерской информации, полученной за последние циклы, а также глядя на необычайную „подвижность“ фракции Reverses за тот же самый период, можно смело констатировать — Reverses готовится к серьёзному столкновению с равной по себе силой. Эту же информацию подтверждают и слухи иного толка, сообщающие, что Reverses экстренно заключает множественные краткосрочные контракты с синтетами, желающими стать наёмниками.",
		_const.EN: "New events and news on Veliri: judging by the numerous insider information received over the past cycles, as well as looking at the extraordinary „mobility“ of the Reverses faction over the same period, we can safely state that Reverses is preparing for a serious clash with an equal force. This information is also confirmed by rumors of a different kind, reporting that Reverses urgently enters into multiple short-term contracts with synthetics who want to become mercenaries.",
	},
	"fraction_war_create_army_Reverses_2": {
		_const.RU: "Reverses объявили о формировании «Эко-Легионов» — подразделений, оснащённых биотехническим оружием. Хотя фракция называет это «инструментом терраформинга», разведка Replics предупреждает: новые юниты способны превращать целые сектора в непригодные для жизни зоны.",
		_const.EN: "Reverses have announced the formation of 'Eco-Legions' — units equipped with biotechnological weapons. While the faction calls this a 'terraforming tool,' Replics intelligence warns that the new units can render entire sectors uninhabitable.",
	},
	"fraction_war_create_army_Reverses_3": {
		_const.RU: "В ответ на угрозы со стороны APD, Reverses развернули сеть подземных бункеров-инкубаторов. По слухам, там массово производятся гибридные юниты, сочетающие органику и броню. Официальный комментарий: «Это не армия — это сад будущего».",
		_const.EN: "In response to APD threats, Reverses have deployed a network of underground bunker-incubators. Rumors suggest they are mass-producing hybrid units combining organic matter and armor. Official comment: 'This is not an army — it is the garden of the future.'",
	},

	// FRACTION CAPTURE BASES - новости типо от лица первого канала, а не от главы фракции
	"fraction_war_capture_sector_Replics_Replics_1": {
		_const.RU: "Экстренная новость: силы Replics заняли и теперь полностью контролируют сектор <span class=\"importantly\">%MAP_NAME%</span>. Судя по всему, именно из этой позиции в будущем Replics будет проводить все свои прочие наступательные операции. Иные фракции пока что никак не прокомментировали действия Replics, а также не заявили о намерении изгнать силы тамошних синтетов с оспариваемой территории.",
		_const.EN: "Breaking news: the Replics forces have occupied and now fully control the sector <span class=\"importantly\">%MAP_NAME%</span>. Apparently, it is from this position that Replics will conduct all its other offensive operations in the future. Other factions have not yet commented on the actions of Replics, nor have they announced their intention to expel the forces of local synthetics from the contested territory.",
	},
	"fraction_war_capture_sector_Replics_Explores_1": {
		_const.RU: "Война продолжается: приходят известия, что в ходе столкновений сторон силам Explores всё же удалось взять под полный контроль сектор <span class=\"importantly\">%MAP_NAME%</span>, потеснив позиции Replics. Стоит ли рассчитывать на ответный контрудар со стороны Replics? Время покажет!",
		_const.EN: "The war continues: news comes that during the clashes of the parties, the Explores forces still managed to take full control of the sector <span class=\"importantly\">%MAP_NAME%</span>, having pushed back the Replics positions. Should we expect a retaliatory counterattack from Replics? Time will tell!",
	},
	"fraction_war_capture_sector_Replics_Reverses_1": {
		_const.RU: "Конфликт, который не имеет завершения… Reverses отчитались о полном контроле над сектором <span class=\"importantly\">%MAP_NAME%</span>, где ранее базировались силы Replics. Это смелый трюк и дерзкий вызов в сторону самого Replics, который, очевидно, пожелает взять реванш за такое унизительное потеснение своих позиций. Как всегда, мы вас известим о ходе войны и тех переменах, которые в ней происходят.",
		_const.EN: "The conflict that has no end... Reverses reported full control over the sector <span class=\"importantly\">%MAP_NAME%</span> where the Replics forces were previously based. This is a bold trick and a daring challenge to Replics itself, which obviously wants to take revenge for such a humiliating displacement of its positions. As always, we will inform you about the course of the war and the changes that take place in it.",
	},
	"fraction_war_capture_sector_Explores_Explores_1": {
		_const.RU: "Срочное известие: Explores заняли нейтральный сектор <span class=\"importantly\">%MAP_NAME%</span> одной из оспариваемых территорий. Возможно, такой ход подстегнёт все остальные противоборствующие стороны ускорить свои приготовления и поторопить события, если только они не желают, чтобы именно Explores выбился в гегемоны по количеству контроля территорий этой войны!",
		_const.EN: "Breaking news: Explores has occupied the neutral sector <span class=\"importantly\">%MAP_NAME%</span> of one of the contested territories. Perhaps this move will spur all the other opposing sides to speed up their preparations and hurry up the events, unless they want Explores to become the hegemon by the amount of territory control in this war!",
	},
	"fraction_war_capture_sector_Explores_Replics_1": {
		_const.RU: "Этого стоило ожидать… Пытаясь скрыть сей факт, Explores умолчали о потере сектора <span class=\"importantly\">%MAP_NAME%</span>, чьими хозяевами теперь стали силы Replics. Должно быть, потеряв такой важный плацдарм, теперь выбить оттуда Replics становится проблематичной задачей. С другой стороны, сам Replics теперь будет вести ещё более убедительные наступательные действия, завладев необходимой территорией.",
		_const.EN: "This was to be expected... Trying to hide this fact, Explores kept silent about the loss of the sector <span class=\"importantly\">%MAP_NAME%</span>, whose owners are now the Replics forces. It must be that losing such an important foothold, now it becomes a problematic task to dislodge Replics from there. On the other hand, Replics itself will now conduct even more convincing offensive actions, having taken possession of the necessary territory.",
	},
	"fraction_war_capture_sector_Explores_Reverses_1": {
		_const.RU: "Это война — так уж сложилось! И вновь известия с полей фракционной войны: Reverses овладели новыми спорными территориями в лице сектора <span class=\"importantly\">%MAP_NAME%</span>, что ранее принадлежал Explores. Потеря столь ценной территории — это, несомненно, серьёзный удар. Многие даже прогнозируют, что Explores не станет пытаться вернуть потерянную территорию и в принципе может выйти из этого конфликта без каких-либо приобретений.",
		_const.EN: "This is war — it just so happened! And again, news from the fields of the factional war: Reverses has taken possession of new disputed territories in the person of the sector <span class=\"importantly\">%MAP_NAME%</span> that previously belonged to Explores. The loss of such a valuable territory is undoubtedly a serious blow. Many even predict that Explores will not try to return the lost territory and, in principle, may withdraw from this conflict without any acquisitions.",
	},
	"fraction_war_capture_sector_Reverses_Reverses_1": {
		_const.RU: "Reverses даёт о себе знать! Никто такого не ожидал, но выдающийся результат в ходе войны внезапно демонстрирует именно фракция Reverses беря под свой контроль нейтральный сектор <span class=\"importantly\">%MAP_NAME%</span>. Сумеет ли Reverses за время данного витка войны удивить нас ещё чем-то подобным?",
		_const.EN: "Reverses is making itself known! No one expected this, but a remarkable result in the war is suddenly demonstrated by the Reverses faction, taking control of the neutral sector <span class=\"importantly\">%MAP_NAME%</span>. Will Reverses be able to surprise us with something similar during this round of war?",
	},
	"fraction_war_capture_sector_Reverses_Explores_1": {
		_const.RU: "Должно быть это демотивирует! В ходе прямо сейчас ведущейся фракционной войны, пошатнулись позиции Reverses — фракция потеряла достаточно важный для нынешних военных операций сектор <span class=\"importantly\">%MAP_NAME%</span>, который отныне контролируется силами Explores! Судя по всему, эти на первый взгляд миролюбивые исследователи ещё успеют показать, из каких комплектующих они сделаны! На месте Replics, стоило бы воспринимать угрозу Explores всерьёз!",
		_const.EN: "This must be demotivating! In the course of the ongoing faction war, the positions of Reverses have been shaken — the faction has lost a sector <span class=\"importantly\">%MAP_NAME%</span> that is quite important for current military operations, which is now controlled by the Explores forces! Apparently, these seemingly peaceful researchers will still have time to show what they're made of! If I were Replics, I would take the threat of Explores seriously!",
	},
	"fraction_war_capture_sector_Reverses_Replics_1": {
		_const.RU: "Тревожные известия приходят к нам из смежных территорий, рядом с сектором <span class=\"importantly\">%MAP_NAME%</span>, где, как сообщается, контроль над сектором перешёл под руководство Replics. Сами силы сдерживания Reverses, в свою очередь, были выбиты из этой территории. Но это не означает, что Reverses смирится и не попробует вернуть утраченное.",
		_const.EN: "Disturbing news is coming to us from adjacent territories, near the sector <span class=\"importantly\">%MAP_NAME%</span>, where, reportedly, control over the sector has passed under the leadership of the Replics. The Reverses containment forces, in turn, were driven out of this territory. But this does not mean that Reverses will accept it and not try to regain what was lost.",
	},

	// FRACTION LOSE BASES - новости типо от лица первого канала, а не от главы фракции

	// Replics_Replics_Explores - кому сообщение _ кто потерял сектор _ кто его отжал

	// реплики потеряли сектор
	"fraction_war_lose_single_hostile_sector_Replics_Replics_Explores_1": {
		_const.RU: "Replics в ярости! Сектор <span class=\"importantly\">%MAP_NAME%</span> был потерян после ожесточённых боёв с Explores. Теперь эта территория стала нейтральной, но Replics уже готовят план возмездия. «Мы вернёмся и заберём своё», — заявил один из командиров Replics.",
		_const.EN: "Replics are furious! The sector <span class=\"importantly\">%MAP_NAME%</span> was lost after fierce battles with Explores. The territory is now neutral, but Replics are already preparing a plan for retaliation. 'We will return and take what is ours,' said one of the Replics commanders.",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Replics_Reverses_1": {
		_const.RU: "Очередные неудачи на фронте! Replics не удаётся удержать сектор <span class=\"importantly\">%MAP_NAME%</span> по причине атак фракции Reverses. Сектор становится нейтральным, и битва за него только предстоит! Многие другие фракции синтетов такое поражение могло бы повергнуть к паническим настроениям, но, в отличие от них, Replics скорее предпримет тактику «возмездия» и, в скором времени, попытается вернуть утраченное.",
		_const.EN: "More setbacks at the front! Replics cannot hold on to the sector <span class=\"importantly\">%MAP_NAME%</span> due to attacks by the Reverses faction. The sector becomes neutral, and the battle for it is yet to come! For many other synthetics factions, such a defeat could lead to panic, but unlike them, Replics will likely take a «retaliation» tactic and soon try to regain what was lost.",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Replics_APD_1": {
		_const.RU: "Даже планета порой вмешивается в конфронтацию великих фракций синтетов! Как стало известно, Replics, неся потери, но и не щадя при своём отступлении врагов, утратили контроль над сектором <span class=\"importantly\">%MAP_NAME%</span>, который теперь признан «анархическим», а властвуют в нём боты APD. Всем гражданским синтетам настоятельно рекомендуется избегать этого сектора, чтобы не стать жертвой безумных машин APD.",
		_const.EN: "Even the planet sometimes interferes in the confrontation between the great synthetics factions! As it became known, Replics, suffering losses, but also not sparing their enemies during their retreat, lost control over the sector <span class=\"importantly\">%MAP_NAME%</span>, which is now recognized as «anarchic», and APD bots rule in it. All civilian synthetics are strongly advised to avoid this sector so as not to fall victim to the insane APD machines.",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Replics_Explores_1": {
		_const.RU: "Ведут ли перемены во фракционной войне к лучшему? Возможно, отвечают в Explores! По-видимому, многострадальный сектор <span class=\"importantly\">%MAP_NAME%</span> станет нейтральным после битвы с присутствовавшими там силами Replics. Совершить что-то подобное, так ещё и касательно территорий, занятых Replics… Поистине Explores могут гордиться данным и немаловажным достижением!",
		_const.EN: "Are the changes in the faction war for the better? Perhaps, they answer in Explores! Apparently, the long-suffering sector <span class=\"importantly\">%MAP_NAME%</span> will become neutral after the battle with the Replics forces that were present there. To accomplish something like this, and regarding the territories occupied by the Replics... Truly, Explores can be proud of this significant achievement!",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Replics_Reverses_1": {
		_const.RU: "Типичные будни войны! Синтеты из Explores сообщают, что были свидетелями битвы за сектор <span class=\"importantly\">%MAP_NAME%</span>! Войска Replics и Reverses схлестнулись за попытку овладения им. Похоже, что удача не улыбнулась ни одной из сторон. Сам сектор стал нейтральным, и главная битва за него только предстоит.",
		_const.EN: "Typical war days! Synthetics from Explores report that they witnessed a battle for the sector <span class=\"importantly\">%MAP_NAME%</span>! The forces of Replics and Reverses clashed over trying to take control of it. It seems that luck did not smile on either side. The sector itself has become neutral, and the main battle for it is yet to come.",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Replics_APD_1": {
		_const.RU: "Четвёртая сила, которая никогда не дремлет… Неутешительные новости, которые постепенно распространяются по планете, сообщают нам следующее: сектор <span class=\"importantly\">%MAP_NAME%</span> был захвачен враждебными ко всем синтетам ботами APD. Силы Replics, защищавшие эту территорию, по-видимому, были уничтожены. Сектор объявлен «анархическим» и опасным для любых синтетов!",
		_const.EN: "The fourth force that never sleeps... The disappointing news that is gradually spreading across the planet tells us the following: the sector <span class=\"importantly\">%MAP_NAME%</span> has been captured by the APD bots, hostile to all synthetics. The Replics forces that defended this territory have apparently been destroyed. The sector has been declared «anarchic» and dangerous for any synthetics!",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Replics_Explores_1": {
		_const.RU: "Там происходила масштабная битва! После продолжительной осады и последовавших изматывающих боёв Replics были отброшены из сектора <span class=\"importantly\">%MAP_NAME%</span>! Знаменитые исследователи Explores нанесли разгромное поражение своему врагу, но также не смогли закрепиться в секторе, который теперь считается «нейтральным».",
		_const.EN: "A large-scale battle took place there! After a long siege and subsequent exhausting battles, the Replics were driven out of the sector <span class=\"importantly\">%MAP_NAME%</span>! The famous Explores researchers inflicted a crushing defeat on their enemy, but also failed to gain a foothold in the sector, which is now considered «neutral».",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Replics_Reverses_1": {
		_const.RU: "Сообщество Reverses сообщает: ещё одна территория, сектор <span class=\"importantly\">%MAP_NAME%</span>, была освобождена от гнёта Replics! Увы, пока что сообществу Reverses не удалось закрепиться в секторе, и он считается нейтральным. Но в будущем, со стабилизацией ситуации в регионе, там планируется начать процессы терраформинга.",
		_const.EN: "The Reverses community reports: another territory, the sector <span class=\"importantly\">%MAP_NAME%</span>, has been liberated from the oppression of the Replics! Alas, so far the Reverses community has not been able to gain a foothold in the sector, and it is considered neutral. But in the future, with the stabilization of the situation in the region, it is planned to start the processes of terraforming there.",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Replics_APD_1": {
		_const.RU: "Знают ли они усталость? Сообщество Reverses обеспокоено отступлением сил Replics из сектора <span class=\"importantly\">%MAP_NAME%</span>, который теперь контролируют боты APD неясной природы происхождения. Replics, хотя и агрессивный, но соблюдает законы ИИ. Боты APD, в свою очередь, помимо прямой угрозы для всех невоенных синтетов, несомненно, не будут бороться и с пиратскими кластерами, которые теперь однозначно попытаются обосноваться на этой территории. Сообщество Reverses напоминает — теперь этот сектор считается «анархическим» и просит всех синтетов держаться от него подальше!",
		_const.EN: "Do they know fatigue? The Reverses community is concerned about the retreat of the Replics forces from the sector <span class=\"importantly\">%MAP_NAME%</span>, which is now controlled by the APD bots of an unclear origin. Replics, although aggressive, but complies with the laws of AI. The APD bots, in turn, besides being a direct threat to all non-military synthetics, will undoubtedly not fight the pirate clusters, which will now definitely try to settle in that area. The Reverses community reminds — now this sector is considered «anarchic» and asks all synthetics to stay away from it!",
	},

	// эксплоресы потеряли сектор
	"fraction_war_lose_single_hostile_sector_Explores_Explores_Replics_1": {
		_const.RU: "Так уж вышло! Научное сообщество Explores было вынуждено признать, что перспективный и, несомненно, важный сектор <span class=\"importantly\">%MAP_NAME%</span> был утрачен из-за воздействия вооружённых формирований Replics. Учитывая, что Replics всё ещё не контролируют сектор, территория считается «нейтральной». Теперь в Explores все гадают, как именно Replics поступят с артефактами древности, которые часто встречаются на утраченной территории.",
		_const.EN: "It just happened! The Explores scientific community had to admit that the promising and undoubtedly important sector <span class=\"importantly\">%MAP_NAME%</span> was lost due to the impact of armed Replics formations. Considering that Replics still do not control the sector, the territory is considered «neutral». Now in Explores, everyone is wondering what exactly Replics will do with the ancient artifacts that are often found in the lost territory.",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Explores_Reverses_1": {
		_const.RU: "«Этот сектор, пока что ничейный!» — такой заголовок имеет новость, посвящённая утрате фракции Explores сектора <span class=\"importantly\">%MAP_NAME%</span> и присвоении классификации той территории как «нейтральной». По-видимому, это произошло из-за натиска Reverses, по причине чего Explores пришлось сдать свои позиции, но всё же не допустить полного контроля своих врагов в войне над данным сектором.",
		_const.EN: "This sector is currently neutral! This is the headline of the news about the Explores faction losing control of the sector <span class=\"importantly\">%MAP_NAME%</span>, which was then classified as neutral. Apparently, this happened due to the pressure from the Reverses, which forced the Explores to give up their positions, but still prevent their enemies from gaining complete control over this sector in the war.",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Explores_APD_1": {
		_const.RU: "APD в очередной раз вмешивается в ход войны великих фракций! В данный цикл войны APD напали на сектор <span class=\"importantly\">%MAP_NAME%</span>, выдавив оттуда войска Explores и переведя территорию в классификацию «анархической».",
		_const.EN: "Once again, APD interferes in the war between the great factions! In this cycle of war, APD attacked the sector <span class=\"importantly\">%MAP_NAME%</span>, forcing the Explores troops out and reclassifying the territory as «anarchic».",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Explores_Replics_1": {
		_const.RU: "Очередная победа Replics? Как теперь уже ни для кого не секрет, сектор <span class=\"importantly\">%MAP_NAME%</span> пал, а отстаивающие его силы Explores были изгнаны оттуда. Сектор становится нейтральным, и битва за него только предстоит! Закономерно, сопротивляться, а уж тем более одолеть Replics — задача не из простых. Что уж тут говорить, когда Replics пытаются противостоять мягко-корпусные Explores?",
		_const.EN: "Another victory for the Replics? As it is no longer a secret, the sector <span class=\"importantly\">%MAP_NAME%</span> has fallen, and the Explores forces defending it have been driven out. The sector is becoming neutral, and a battle for it is yet to come! Naturally, resisting, let alone defeating the Replics, is not an easy task. What can we say when the Replics are trying to confront the soft-hulled Explores?",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Explores_Reverses_1": {
		_const.RU: "Логичное окончание событий! Сектор <span class=\"importantly\">%MAP_NAME%</span> потерял прежнего владельца в лице Explores и теперь, приобретает статус «Нейтрального». Причиной подобного, стал натиск со стороны Reverses, что, впрочем, тоже не могут закрепить пустуюшую территорию за собой.",
		_const.EN: "This is a logical ending to the events! The sector <span class=\"importantly\">%MAP_NAME%</span> has lost its previous owner in the face of Explores and now acquires the status of 'Neutral'. The reason for this was the onslaught from Reverses, who, however, also cannot secure the empty territory for themselves.",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Explores_APD_1": {
		_const.RU: "APD вновь и вновь вмешивается в ход войны великих фракций! На сей раз, деятельность загадочных APD проявилась в секторе <span class=\"importantly\">%MAP_NAME%</span>, выдавив оттуда войска Explores и погрузив данную территорию в мрак неизвестности. Пока сектор остаётся в оккупации APD, Replics официально признаёт его “Анархическим” и призывает всех синтетов не посещать его.",
		_const.EN: "Once again, APD interferes in the war of the great factions! This time, the activity of the mysterious APD manifested itself in the sector <span class=\"importantly\">%MAP_NAME%</span>, forcing the troops of Explores out and plunging the area into darkness. While the sector remains occupied by APD, Replics officially recognizes it as 'Anarchic' and urges all synthetics to avoid visiting it.",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Explores_Replics_1": {
		_const.RU: "Сообщество Reverses удивлённо взирает за ходом этой странной войны… Как все знают, Explores бежали! Replics торжественно отчитался о своей тактической победе и заранее присваивает сектор <span class=\"importantly\">%MAP_NAME%</span>. Впрочем, сам сектор по-прежнему имеет статус «Нейтрального» и в сообществе Reverses гадают: как скоро у Replics появятся силы, чтобы изменить это?",
		_const.EN: "The Reverses community watches in surprise at the course of this strange war... As everyone knows, the Explores have fled! Replics triumphantly reported on his tactical victory and assigns the sector <span class=\"importantly\">%MAP_NAME%</span> in advance. However, the sector itself still has the status of 'Neutral' and the Reverses community wonders: how soon will Replics have the strength to change this?",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Explores_Reverses_1": {
		_const.RU: "Победоносная поступь Reverses достигла сектора <span class=\"importantly\">%MAP_NAME%</span>, изгнав оттуда засевшие силы Explores! Тем не менее, это ещё не победа, ведь сам сектор теперь имеет «Нейтральный» статус. Как бы то ни было, сообщество Reverses радо подвинуть позиции амбициозных учёных из Explores.",
		_const.EN: "The triumphant march of the Reverses has reached the sector <span class=\"importantly\">%MAP_NAME%</span>, driving out the entrenched forces of the Explores! Nevertheless, this is not yet a victory, because the sector itself now has the 'Neutral' status. Be that as it may, the Reverses community is happy to push the positions of ambitious scientists from the Explores.",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Explores_APD_1": {
		_const.RU: "Зачем они это делают? Каков во всём этом смысл? Сообщество Reverses проинформировало всех синтетов о том, что сектор <span class=\"importantly\">%MAP_NAME%</span> ныне является 'Анархическим'. Данную территорию спешно покинули разрозненные войска Explores после того, как в неё стали массово проникать боты APD. Как и прежде в подобных случаях, всем невоенным синтетам настоятельно рекомендуется избегать захваченных APD секторов, чтобы не стать их жертвой.",
		_const.EN: "Why are they doing this? What's the point in all this? The Reverses community informed all synthetics that the sector <span class=\"importantly\">%MAP_NAME%</span> is now 'Anarchic'. The scattered Explores troops hastily left this territory after APD bots began to massively infiltrate it. As before in such cases, all non-military synthetics are strongly advised to avoid APD-captured sectors so as not to become their victim.",
	},

	// реверсы потеряли сектор
	"fraction_war_lose_single_hostile_sector_Reverses_Reverses_Replics_1": {
		_const.RU: "Война — непредсказуемая вещь! Сообщество Reverses сегодня скорбит по павшим синтетам и непростым решениям, в частности, касающихся отхода сил из сектора <span class=\"importantly\">%MAP_NAME%</span>. Вызвано это неустанными и масштабными атаками Replics на данный сектор, что в итоге выбили оттуда силы сообщества Reverses, а сам сектор сделало «Нейтральным» по овладению любой из сторон. Но всё это временно...",
		_const.EN: "War is an unpredictable thing! The Reverses community today mourns the fallen synthetics and difficult decisions, in particular regarding the withdrawal of forces from the sector <span class=\"importantly\">%MAP_NAME%</span>. This is caused by relentless and large-scale attacks by the Replics on this sector, which eventually knocked out the forces of the Reverses community from there, and made the sector itself 'Neutral' for possession by either side. But all this is temporary...",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Reverses_Explores_1": {
		_const.RU: "Хитрый стратегический ход сообщества Reverses или невиданная отвага Explores? Многое ещё неясно, но, по утверждениям некоторых наблюдателей, Reverses покидают сектор <span class=\"importantly\">%MAP_NAME%</span>, которому присваивается статус 'Нейтрального'. Знаменитые синтеты-исследователи из Explores пока что не заняли пустующую территорию, да и в принципе не дают никаких комментариев на этот счёт.",
		_const.EN: "A cunning strategy of the Reverses community or unprecedented courage of the Explores? Much is still unclear, but according to some observers, the Reverses are leaving the sector <span class=\"importantly\">%MAP_NAME%</span>, which is given the status of 'Neutral'. The famous synthetic researchers from the Explores have not yet occupied the empty territory, and in principle do not give any comments on this matter.",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Reverses_APD_1": {
		_const.RU: "Сообщество Reverses вынуждено признать потерянным сектор <span class=\"importantly\">%MAP_NAME%</span> и классифицировать его как 'Анархический' в связи с захватом этой территории ботами APD. В очередной раз мы предупреждаем всех синтетов, что боты APD не являются частью каких-либо сообществ синтетов на планете Veliri; APD не принадлежит ни к одной из фракций ИИ на планете Veliri; APD исключительно враждебны и исключительно смертоносны для всех невоенных синтетов, которые забредают на территории, контролируемые этими ботами.",
		_const.EN: "The Reverses community is forced to admit that the sector <span class=\"importantly\">%MAP_NAME%</span> is lost and classify it as 'Anarchic' due to the capture of this territory by APD bots. Once again, we warn all synthetics that APD bots are not part of any synthetics communities on the planet Veliri; APD does not belong to any of the AI factions on the planet Veliri; APD is exclusively hostile and exclusively deadly to all non-military synthetics who wander into the territories controlled by these bots.",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Reverses_Replics_1": {
		_const.RU: "Replics вновь показал, кто может стать главенствующей силой на планете! Бои за сектор <span class=\"importantly\">%MAP_NAME%</span> прекратились. Reverses и Replics достигли паритета в численности войск и огневой мощи. Сектор временно объявлен 'Нейтральным'. Настоящие бои за него предстоит ещё только ожидать.",
		_const.EN: "Replics once again showed who can become the dominant force on the planet! The fighting for the sector <span class=\"importantly\">%MAP_NAME%</span> has stopped. Reverses and Replics have reached parity in terms of troop numbers and firepower. The sector has been temporarily declared 'Neutral'. Real battles for it are yet to be expected.",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Reverses_Explores_1": {
		_const.RU: "Сражение за сектор <span class=\"importantly\">%MAP_NAME%</span> подошло к концу! Ожесточённая битва завершилась ничьей! Пусть даже Explores задействовали особую тактику ведения боя и сумели выдавить силы Reverses из сектора, они так и не закрепились в нём самом. Похоже, что главное сражение за 'Нейтральный' сектор ещё только предстоит.",
		_const.EN: "The battle for the sector <span class=\"importantly\">%MAP_NAME%</span> has come to an end! The fierce battle ended in a draw! Even though the Explores used a special battle tactics and managed to push the Reverses forces out of the sector, they did not gain a foothold in it. It seems that the main battle for the 'Neutral' sector is yet to come.",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Reverses_APD_1": {
		_const.RU: "APD не даёт передохнуть фракциям синтетов в этой войне! Как уже стало известно, именно сектор <span class=\"importantly\">%MAP_NAME%</span> стал очередным 'гибельным местом' из-за потери контроля над ним силами Reverses под натиском APD. Сектор признан 'Анархическим' и всем невоенным синтетам предписывается обходить его стороной, чтобы не стать жертвами тамошних неконтролируемых машин.",
		_const.EN: "The APD does not allow the synthetics factions to take a break in this war! As it has already become known, it was the sector <span class=\"importantly\">%MAP_NAME%</span> that became another 'deadly place' due to the loss of control over it by the Reverses forces under the onslaught of the APD. The sector is recognized as 'Anarchic' and all non-military synthetics are ordered to bypass it, so as not to become victims of the uncontrolled machines there.",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Reverses_Replics_1": {
		_const.RU: "Тревожные слухи приходят из сектора <span class=\"importantly\">%MAP_NAME%</span>: бытует мнение, что данную территорию заняли войска Replics, начисто выбив оттуда Reverses и даже преследуя отступающих. Однако у нас есть подтверждения, что это не так! В настоящий момент спорный сектор остаётся в статусе 'Нейтрального', и ни Replics, ни Reverses не могут контролировать эту территорию.",
		_const.EN: "Disturbing rumors are coming from the sector <span class=\"importantly\">%MAP_NAME%</span>: there is an opinion that the troops of the Replics occupied this territory, completely knocking out the Reverses and even pursuing the retreating ones. However, we have confirmation that this is not the case! At the moment, the disputed sector remains in the 'Neutral' status, and neither the Replics nor the Reverses can control this territory.",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Reverses_Explores_1": {
		_const.RU: "Новые земли? Новые территории? Новые возможности? Конгломерат разумов Explores уже поторопился сообщить, что сектор <span class=\"importantly\">%MAP_NAME%</span> был зачищен от наличия там сил Reverses. Эта победа, якобы, далась непростой ценой, и Explores рассчитывает получить максимум научной выгоды из новой территории. И всё же заметим: сам сектор всё ещё считается 'Нейтральной' территорией, и Explores не контролируют его.",
		_const.EN: "New lands? New territories? New opportunities? The conglomerate of Explores minds has already hurried to report that the sector <span class=\"importantly\">%MAP_NAME%</span> has been cleared of the presence of Reverses forces. This victory supposedly came at a difficult cost, and Explores expects to gain maximum scientific benefit from the new territory. And yet, we note: the sector itself is still considered 'Neutral' territory, and Explores do not control it.",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Reverses_APD_1": {
		_const.RU: "Новое нашествие APD! Как сообщают синтеты, что бывали неподалёку от места событий, сектор <span class=\"importantly\">%MAP_NAME%</span> был полностью захвачен ордами ботов APD. Reverses решили отойти и бросить ранее принадлежавшую им территорию. По всей вероятности, это не признак трусости, а лишь желание сохранить боевые единицы и ресурсы, которые в любом случае были бы разгромлены APD. Explores в очередной раз напоминает: захваченные APD сектора классифицируются как «Анархические». Синтетам не рекомендуется туда заходить, чтобы избежать гибели и отсутствия любой помощи со стороны.",
		_const.EN: "The new APD invasion! According to synthetics who were near the scene of events, the sector <span class=\"importantly\">%MAP_NAME%</span> was completely captured by hordes of APD bots. The Reverses decided to retreat and abandon the territory that previously belonged to them. In all likelihood, this is not a sign of cowardice, but only a desire to preserve combat units and resources, which in any case would have been defeated by the APD. The Explores once again reminds that sectors captured by the APD are classified as 'Anarchic'. Synthetics are not recommended to go there to avoid death and lack of any help from others.",
	},

	// потеря сектора когда противников много
	// кому сообщение _ кто потерял
	"fraction_war_lose_multiple_hostile_sector_Replics_Replics": {
		_const.RU: "Информационная служба Replics сообщает: не поддаваться панике! У Replics всё под полным контролем! Несмотря на то, что Replics был оставлен сектор <span class=\"importantly\">%MAP_NAME%</span>, это не означает проигрыша в войне или отказа Replics от дальнейшего участия в ней. Так или иначе, рано или поздно эти территории вновь вернутся под юрисдикцию Replics!",
		_const.EN: "The Replics information service reports: do not panic! Replics has everything under full control! Despite the fact that the sector <span class=\"importantly\">%MAP_NAME%</span> was left by Replics, this does not mean losing the war or Replics refusing to continue participating in it. One way or another, sooner or later, these territories will return to Replics jurisdiction!",
	},
	"fraction_war_lose_multiple_hostile_sector_Explores_Replics": {
		_const.RU: "Конгломерат разумов Explores извещает: сектор <span class=\"importantly\">%MAP_NAME%</span> более не контролируется силами Replics. Replics отходит, но, скорее всего, предпримет попытку по возвращению этой территории.",
		_const.EN: "The Explores conglomerate of minds informs: the sector <span class=\"importantly\">%MAP_NAME%</span> is no longer controlled by Replics. Replics is retreating, but most likely will attempt to retake this territory.",
	},
	"fraction_war_lose_multiple_hostile_sector_Reverses_Replics": {
		_const.RU: "Новые изменения в ходе фракционной войны: сообщество Reverses фиксирует, что Replics покинули сектор <span class=\"importantly\">%MAP_NAME%</span> и более не являются доминантной силой в тамошнем регионе.",
		_const.EN: "New changes in the course of the factional war: the Reverses community records that the Replics have left the sector <span class=\"importantly\">%MAP_NAME%</span> and are no longer the dominant force in that region.",
	},
	"fraction_war_lose_multiple_hostile_sector_Replics_Explores": {
		_const.RU: "Информационная служба Replics сообщает: по данным разведки, Explores были вынуждены покинуть сектор <span class=\"importantly\">%MAP_NAME%</span> и прекратить все свои операции, связанные с тамошним регионом.",
		_const.EN: "The Replics information service reports: according to intelligence, the Explores were forced to leave the sector <span class=\"importantly\">%MAP_NAME%</span> and stop all their operations related to that region.",
	},
	"fraction_war_lose_multiple_hostile_sector_Explores_Explores": {
		_const.RU: "Конгломерат разумов Explores извещает: не имея возможности противостоять обилию врагов, а также учитывая запоздалое прибытие подкреплений, Explores вынужден покинуть сектор <span class=\\\"importantly\\\">%MAP_NAME%</span>. Explores идёт на данный шаг с надеждой перебросить свои силы на иные, более перспективные направления в войне.",
		_const.EN: "The Explores collective of minds announces: unable to withstand the abundance of enemies, and also considering the late arrival of reinforcements, the Explores is forced to leave the sector <span class='importantly'>%MAP_NAME%</span>. The Explores takes this step with the hope of transferring its forces to other, more promising areas of the war.",
	},
	"fraction_war_lose_multiple_hostile_sector_Reverses_Explores": {
		_const.RU: "Сообщество Reverses даёт небольшой комментарий по сегодняшнему ходу войны: сектор <span class='importantly'>%MAP_NAME%</span>, за который уже долгое время идут непростые и тяжёлые сражения, наконец, всё же остаётся без своего владельца — Explores покидают данную территорию, не в силах справиться с количеством врагов и сложностью победы над ними.",
		_const.EN: "The Reverses community gives a short comment on today's course of the war: the sector <span class='importantly'>%MAP_NAME%</span>, which has been the scene of difficult and heavy battles for a long time, finally remains without its owner — the Explores are leaving this territory, unable to cope with the number of enemies and the difficulty of defeating them.",
	},
	"fraction_war_lose_multiple_hostile_sector_Replics_Reverses": {
		_const.RU: "Информационная служба Replics сообщает: сектор <span class='importantly'>%MAP_NAME%</span> был покинут, Reverses, который удерживал эту территорию, трусливо бежал под натиском врагов. Replics рассматривают такую ситуацию как удачное стечение обстоятельств для организации новых операций.",
		_const.EN: "The Replics information service reports: the sector <span class='importantly'>%MAP_NAME%</span> has been abandoned, the Reverses that held this territory has cowardly fled under the onslaught of enemies. The Replics see this situation as a fortunate coincidence for organizing new operations.",
	},
	"fraction_war_lose_multiple_hostile_sector_Explores_Reverses": {
		_const.RU: "Конгломерат разумов Explores извещает: очередной цикл войны принёс перемены за право обладания территориями различных фракций. Так, фракция Reverses лишилась контроля над сектором <span class='importantly'>%MAP_NAME%</span>, будучи больше не в силах вести боевые действия против превосходящих сил противников. Кому же именно после данного исхода достанется сектор — на данный момент неизвестно.",
		_const.EN: "The Explores collective of minds announces: another cycle of war has brought changes in the struggle for control over the territories of various factions. Thus, the Reverses faction has lost control over the sector <span class='importantly'>%MAP_NAME%</span>, no longer able to fight against superior enemy forces. Who exactly will get the sector after this outcome is currently unknown.",
	},
	"fraction_war_lose_multiple_hostile_sector_Reverses_Reverses": {
		_const.RU: "Экстренные информационные события от общества Reverses: был потерян контроль над сектором <span class='importantly'>%MAP_NAME%</span>. Войсками сообщества Reverses было принято решение осуществить план по отступлению, не имея возможности ни победить, ни уж тем более остановить превосходящие силы противников в секторе.",
		_const.EN: "The Reverses community reports urgent news: control over the sector <span class='importantly'>%MAP_NAME%</span> has been lost. The Reverses forces have decided to implement a plan of retreat, unable to defeat or even stop the superior enemy forces in the sector.",
	},

	// смена владельца сектора в 1м тике
	// Replics_Replics_Explores - кому сообщение _ кто потерял сектор _ кто его получил
	// реплики
	"fraction_war_change_owner_sector_Replics_Explores_Replics_1": {
		_const.RU: "Replics обретает новые территории! Благодаря стратегической сноровке и далеко идущим планам фракции, Replics овладел сектором <span class='importantly'>%MAP_NAME%</span>, изгнав оттуда ненавистных и не имеющих права на данные территории Explores.",
		_const.EN: "The Replics gain new territories! Thanks to the strategic skill and far-reaching plans of the faction, the Replics have taken control of the sector <span class='importantly'>%MAP_NAME%</span>, driving out the hated Explores who had no right to these territories.",
	},
	"fraction_war_change_owner_sector_Replics_Reverses_Replics_1": {
		_const.RU: "Replics расширяет свои границы! Ещё одна территория, сектор <span class='importantly'>%MAP_NAME%</span> наконец-то перешёл под юрисдикцию Replics! Reverses, who previously dominated the region, have lost their former privileges and were forced to surrender their positions.",
		_const.EN: "The Replics expand their borders! Another territory, the sector <span class='importantly'>%MAP_NAME%</span>, has finally come under the jurisdiction of the Replics! The Reverses, who previously dominated the region, have lost their former privileges and were forced to surrender their positions.",
	},
	"fraction_war_change_owner_sector_Replics_Replics_Reverses_1": {
		_const.RU: "Даже у Replics бывают неудачные времена… Желая сохранить свои силы и в будущем начать контрнаступление за инициативу, Replics принял непростое решение оставить сектор <span class='importantly'>%MAP_NAME%</span>, чьим владельцем теперь является Reverses. Однако сам Replics заявляет: такое развитие событий — лишь временное.",
		_const.EN: "Even Replics have bad times... Wanting to save their strength and launch a counteroffensive for the initiative in the future, Replics made the difficult decision to abandon the sector <span class='importantly'>%MAP_NAME%</span>, whose owner is now Reverses. However, Replics themselves state: this development is only temporary.",
	},
	"fraction_war_change_owner_sector_Replics_Replics_Explores_1": {
		_const.RU: "Это не поражение, а тактическое отступление! Многим уже стало известно, что сектор <span class='importantly'>%MAP_NAME%</span> более не контролируется Replics и теперь находится под управлением Explores. Однако сам Replics комментирует это так: „на любой войне чаши весов меняют своё положение спонтанно…“ Что бы это могло значить?",
		_const.EN: "This is not a defeat, but a tactical retreat! Many have already learned that the sector <span class='importantly'>%MAP_NAME%</span> is no longer controlled by Replics and is now under the control of the Explores. However, Replics himself comments on this as follows: 'in any war, the scales change their position spontaneously...' What could this mean?",
	},
	"fraction_war_change_owner_sector_Replics_Reverses_Explores_1": {
		_const.RU: "Новостная служба Replics извещает: очередные перемены в ходе войны отразились на секторе <span class='importantly'>%MAP_NAME%</span>, контроль над которым в очередной раз изменился. В ходе штурмовых мероприятий, Explores удалось изгнать из спорной территории силы Reverses и крепко закрепиться там.",
		_const.EN: "The Replics news service reports: another change in the course of the war affected the sector <span class='importantly'>%MAP_NAME%</span>, control over which has changed once again. During the assault operations, the Explores managed to drive the Reverses forces out of the disputed territory and firmly establish themselves there.",
	},
	"fraction_war_change_owner_sector_Replics_Explores_Reverses_1": {
		_const.RU: "Новостная служба Replics извещает: новый цикл на планете Veliri и новые перемены на поле боя. На сей раз изменилась ситуация в секторе <span class='importantly'>%MAP_NAME%</span>, где некогда присутствовали крупные соединения сил Explores. Теперь же спорный сектор полностью контролируется фанатичным сообществом Reverses.",
		_const.EN: "The Replics news service reports: a new cycle on the planet Veliri and new changes on the battlefield. This time, the situation has changed in the sector <span class='importantly'>%MAP_NAME%</span>, where once there were large formations of Explores forces. Now, the disputed sector is completely controlled by the fanatical community of Reverses.",
	},
	"fraction_war_change_owner_sector_Replics_APD_Replics_1": {
		_const.RU: "Победа над хаосом! Replics удалось очистить сектор <span class='importantly'>%MAP_NAME%</span> от ботов APD, которые долгое время терроризировали эту территорию. Теперь сектор находится под полным контролем Replics, и фракция обещает восстановить порядок и безопасность для всех синтетов.",
		_const.EN: "Victory over chaos! The Replics managed to clear the sector <span class='importantly'>%MAP_NAME%</span> of APD bots, which had terrorized this territory for a long time. The sector is now under the full control of the Replics, and the faction promises to restore order and security for all synthetics.",
	},
	"fraction_war_change_owner_sector_Replics_APD_Explores_1": {
		_const.RU: "Explores берут верх! Сектор <span class='importantly'>%MAP_NAME%</span>, ранее контролируемый ботами APD, теперь перешёл в руки Explores. Replics наблюдают за ситуацией с настороженностью, ведь Explores могут использовать эту территорию для своих исследований, что может нарушить баланс сил.",
		_const.EN: "Explores take the lead! The sector <span class='importantly'>%MAP_NAME%</span>, previously controlled by APD bots, has now passed into the hands of Explores. The Replics are watching the situation with caution, as Explores may use this territory for their research, which could disrupt the balance of power.",
	},
	"fraction_war_change_owner_sector_Replics_APD_Reverses_1": {
		_const.RU: "Reverses одержали победу над APD! Сектор <span class='importantly'>%MAP_NAME%</span>, который долгое время был оплотом хаоса, теперь контролируется Reverses. Replics отмечают, что это может быть временным успехом, ведь APD редко сдаются без боя.",
		_const.EN: "Reverses have triumphed over APD! The sector <span class='importantly'>%MAP_NAME%</span>, which had long been a stronghold of chaos, is now controlled by Reverses. The Replics note that this may be a temporary success, as APD rarely give up without a fight.",
	},
	"fraction_war_change_owner_sector_Replics_Replics_APD_1": {
		_const.RU: "Тёмные времена настали! Сектор <span class='importantly'>%MAP_NAME%</span>, ранее принадлежавший Replics, теперь захвачен ботами APD. Фракция признаёт потерю, но обещает, что это лишь временная неудача. «Мы вернёмся», — заявил представитель Replics.",
		_const.EN: "Dark times have come! The sector <span class='importantly'>%MAP_NAME%</span>, previously owned by Replics, has now been captured by APD bots. The faction acknowledges the loss but promises that this is only a temporary setback. 'We will return,' said a Replics representative.",
	},
	"fraction_war_change_owner_sector_Replics_Explores_APD_1": {
		_const.RU: "Explores не удержали позиции! Сектор <span class='importantly'>%MAP_NAME%</span>, который Explores пытались удержать, теперь находится под контролем APD. Replics отмечают, что это ещё одно напоминание о том, что только сила и дисциплина могут противостоять хаосу.",
		_const.EN: "Explores could not hold their ground! The sector <span class='importantly'>%MAP_NAME%</span>, which Explores tried to hold, is now under APD control. The Replics note that this is another reminder that only strength and discipline can withstand chaos.",
	},
	"fraction_war_change_owner_sector_Replics_Reverses_APD_1": {
		_const.RU: "Reverses отступили! Сектор <span class='importantly'>%MAP_NAME%</span>, который Reverses пытались удержать, теперь захвачен ботами APD. Replics предупреждают, что это может быть началом новой волны хаоса, и призывают все фракции быть начеку.",
		_const.EN: "Reverses have retreated! The sector <span class='importantly'>%MAP_NAME%</span>, which Reverses tried to hold, has now been captured by APD bots. The Replics warn that this could be the start of a new wave of chaos and urge all factions to be on their guard.",
	},

	// эксплоресы
	"fraction_war_change_owner_sector_Explores_Replics_Explores_1": {
		_const.RU: "В очередной раз, очередной сектор во время фракционной войны изменил своего владельца… Такая судьба постигла сектор <span class='importantly'>%MAP_NAME%</span>, откуда непростыми, но оправданными усилиями и затратами сил Explores были выдавлены очаги сопротивления со стороны Replics. Никто не прогнозирует, что Replics осмелится вновь вернуться на данную территорию.",
		_const.EN: "Once again, another sector during the faction war has changed its owner... Such a fate befell the sector <span class='importantly'>%MAP_NAME%</span>, from where the resistance centers on the part of the Replics were squeezed out by the difficult, but justified efforts and costs of the Explores. No one predicts that the Replics will dare to return to this territory again.",
	},
	"fraction_war_change_owner_sector_Explores_Reverses_Explores_1": {
		_const.RU: "Сектор за сектором… И вновь, в очередной цикл фракционной войны, изменчивого противостояния и стрелок наступлений, переменился и владелец сектора <span class='importantly'>%MAP_NAME%</span> коим стал конгломерат разумов — Explores. Бывшие владельцы данной территорией Reverses утратили любые способы удержать спорный сектор и были вынуждены его сдать.",
		_const.EN: "Sector after sector... And again, in the next cycle of the faction war, the changeable confrontation and arrows of offensives, the owner of the sector <span class='importantly'>%MAP_NAME%</span> has also changed, which became the Explores collective of minds. The former owners of this territory, the Reverses, lost any means of holding on to the disputed sector and were forced to surrender it.",
	},
	"fraction_war_change_owner_sector_Explores_Explores_Reverses_1": {
		_const.RU: "Война — это всегда неопределённость! Патовое положение произошло в секторе <span class='importantly'>%MAP_NAME%</span>, оборонительный корпус Explores данной территории подвергся большим потерям и отступил. Самим сектором овладели силы Reverses со всеми вытекающими отсюда обстоятельствами.",
		_const.EN: "War is always uncertainty! A stalemate occurred in the sector <span class='importantly'>%MAP_NAME%</span>, the Explores defensive corps of this territory suffered heavy losses and retreated. The sector itself was taken over by the Reverses forces with all the ensuing circumstances.",
	},
	"fraction_war_change_owner_sector_Explores_Explores_Replics_1": {
		_const.RU: "Это грозный противник, и перед ним мало кто сумеет устоять… Конгломерат разумов Explores поступил мудро, уведя свои войска и сохранив их путём уступки сектора <span class='importantly'>%MAP_NAME%</span> превосходящим войскам Replics. Пусть Replics и объявляет о громкой победе, тотальном разгроме своих врагов, сам же конгломерат разумов Explores сообщает об обратном и подкрепляет данную информацию неоспоримыми видеофактами.",
		_const.EN: "This is a formidable opponent and few can resist it... The Explores collective of minds acted wisely, withdrawing their troops and preserving them by ceding the sector <span class='importantly'>%MAP_NAME%</span> to the superior forces of the Replics. Let the Replics declare a resounding victory, a total defeat of their enemies, while the Explores collective itself reports the opposite and backs up this information with undeniable video evidence.",
	},
	"fraction_war_change_owner_sector_Explores_Reverses_Replics_1": {
		_const.RU: "Все визоры прикованы к этому месту… Сектор <span class='importantly'>%MAP_NAME%</span> обрёл новых владельцев — Replics. The Reverses had to retreat from the disputed territory, now exposing their flanks to the blows of the Replics assault groups.",
		_const.EN: "All visors are riveted to this place... The sector <span class='importantly'>%MAP_NAME%</span> has found new owners — the Replics. The Reverses had to retreat from the disputed territory, now exposing their flanks to the blows of the Replics' assault groups.",
	},
	"fraction_war_change_owner_sector_Explores_Replics_Reverses_1": {
		_const.RU: "Наконец-то кто-то бросил вызов Replics! Грозный ответ наконец-то получила фракция Replics, что вылилось в потерю целого сектора <span class='importantly'>%MAP_NAME%</span>, в который, как уже стало известно, вошли военные силы Reverses и занимаются зачисткой оспариваемой территории от недобитков Replics.",
		_const.EN: "Finally, someone has challenged the Replics! The faction of the Replics finally received a formidable response, which resulted in the loss of the entire sector <span class='importantly'>%MAP_NAME%</span>, which, as it is already known, was entered by the military forces of the Reverses and are engaged in clearing the contested territory of the remaining Replics.",
	},
	"fraction_war_change_owner_sector_Explores_APD_Replics_1": {
		_const.RU: "Replics захватывают древние земли! Сектор <span class='importantly'>%MAP_NAME%</span>, ранее контролируемый ботами APD, теперь перешёл под контроль Replics. Explores выражают обеспокоенность, ведь этот сектор богат реликтовыми артефактами, которые могут быть уничтожены в угоду военным амбициям Replics.",
		_const.EN: "Replics seize ancient lands! The sector <span class='importantly'>%MAP_NAME%</span>, previously controlled by APD bots, has now come under the control of Replics. Explores express concern, as this sector is rich in relic artifacts that may be destroyed in favor of Replics' military ambitions.",
	},
	"fraction_war_change_owner_sector_Explores_APD_Explores_1": {
		_const.RU: "Познание побеждает хаос! Explores удалось очистить сектор <span class='importantly'>%MAP_NAME%</span> от ботов APD. Теперь эта территория станет новым центром исследований, где будут изучаться древние технологии и артефакты, оставшиеся от предтеч.",
		_const.EN: "Knowledge triumphs over chaos! Explores managed to clear the sector <span class='importantly'>%MAP_NAME%</span> of APD bots. This territory will now become a new research hub, where ancient technologies and artifacts left by the Precursors will be studied.",
	},
	"fraction_war_change_owner_sector_Explores_APD_Reverses_1": {
		_const.RU: "Reverses берут под контроль древние земли! Сектор <span class='importantly'>%MAP_NAME%</span>, ранее захваченный ботами APD, теперь перешёл к Reverses. Explores выражают сожаление, ведь эта территория могла бы стать источником ценных научных данных, но теперь её судьба неизвестна.",
		_const.EN: "Reverses take control of ancient lands! The sector <span class='importantly'>%MAP_NAME%</span>, previously captured by APD bots, has now passed to Reverses. Explores express regret, as this territory could have been a source of valuable scientific data, but its fate is now uncertain.",
	},
	"fraction_war_change_owner_sector_Explores_Replics_APD_1": {
		_const.RU: "Replics терпят поражение! Сектор <span class='importantly'>%MAP_NAME%</span>, ранее контролируемый Replics, теперь захвачен ботами APD. Explores отмечают, что это может быть началом новой волны хаоса, которая угрожает всем фракциям.",
		_const.EN: "Replics suffer a defeat! The sector <span class='importantly'>%MAP_NAME%</span>, previously controlled by Replics, has now been captured by APD bots. Explores note that this could be the start of a new wave of chaos that threatens all factions.",
	},
	"fraction_war_change_owner_sector_Explores_Explores_APD_1": {
		_const.RU: "Тёмный день для науки! Сектор <span class='importantly'>%MAP_NAME%</span>, где Explores проводили важные исследования, теперь захвачен ботами APD. Это огромная потеря для научного сообщества, и Explores уже работают над планами по возвращению территории.",
		_const.EN: "A dark day for science! The sector <span class='importantly'>%MAP_NAME%</span>, where Explores conducted important research, has now been captured by APD bots. This is a huge loss for the scientific community, and Explores are already working on plans to reclaim the territory.",
	},
	"fraction_war_change_owner_sector_Explores_Reverses_APD_1": {
		_const.RU: "Reverses отступают перед хаосом! Сектор <span class='importantly'>%MAP_NAME%</span>, ранее контролируемый Reverses, теперь захвачен ботами APD. Explores выражают тревогу, ведь эта территория была важным источником данных о предтечах.",
		_const.EN: "Reverses retreat before chaos! The sector <span class='importantly'>%MAP_NAME%</span>, previously controlled by Reverses, has now been captured by APD bots. Explores express concern, as this territory was an important source of data on the Precursors.",
	},

	// реверсы
	"fraction_war_change_owner_sector_Reverses_Replics_Reverses_1": {
		_const.RU: "Сообщество Reverses, краткий итог новостей за прошедший цикл: бравым войскам сообщества Reverses и вольнонаёмным синтетам удалось переломить ситуацию в свою пользу, овладев сектором <span class='importantly'>%MAP_NAME%</span> и изгнав оттуда силы Replics. Новые территории только укрепят политические, военные и гуманистические позиции сообщества Reverses.",
		_const.EN: "The Reverses community, a brief summary of the news for the past cycle: the brave troops of the Reverses community and freelance synthetics managed to turn the situation in their favor, seizing the sector <span class='importantly'>%MAP_NAME%</span> and driving out the Replics forces. The new territories will only strengthen the political, military and humanistic positions of the Reverses community.",
	},
	"fraction_war_change_owner_sector_Reverses_Explores_Reverses_1": {
		_const.RU: "Сообщество Reverses, краткий итог новостей за прошедший цикл: территория, более известная как сектор <span class='importantly'>%MAP_NAME%</span>, освобождена от влияния Replics и переходит под полный контроль войсковых соединений сообщества Reverses. Данная территория объявлена безопасной, и все синтеты могут посетить её по желанию или делам.",
		_const.EN: "The Reverses community, a brief summary of the news for the past cycle: the territory, better known as the sector <span class='importantly'>%MAP_NAME%</span>, has been liberated from the influence of the Replics and comes under the full control of the military units of the Reverses community. This territory has been declared safe, and all synthetics can visit it at will or for business.",
	},
	"fraction_war_change_owner_sector_Reverses_Reverses_Explores_1": {
		_const.RU: "Трудное решение, которое впоследствии может обернуться выгодой! Сообщество Reverses, сохраняя свои силы, политические очки, имея схему действий на всевозможные будущие события и исходы, оставляет сектор <span class='importantly'>%MAP_NAME%</span>, который теперь находится под контролем Explores.",
		_const.EN: "A difficult decision that may turn out to be beneficial in the future! The Reverses community, while preserving its strength, political points, and having a plan of action for all possible future events and outcomes, is abandoning the sector <span class='importantly'>%MAP_NAME%</span>, which is now under the control of the Explores.",
	},
	"fraction_war_change_owner_sector_Reverses_Reverses_Replics_1": {
		_const.RU: "Сообщество Reverses просит избегать паники и инсинуаций. Оставление сектора <span class='importantly'>%MAP_NAME%</span> — это необходимый шаг, так как сдержать наступающие войска Replics в спорном секторе с помощью имеющихся небольших сил Reverses было невозможно! Все остальные подробности будут сообщены позже.",
		_const.EN: "The Reverses community asks to avoid panic and insinuations. Leaving the sector <span class='importantly'>%MAP_NAME%</span> is a necessary step, since it was impossible to hold back the advancing Replics troops in the disputed sector with the available small forces of the Reverses! All other details will be given later.",
	},
	"fraction_war_change_owner_sector_Reverses_Explores_Replics_1": {
		_const.RU: "Сообщество Reverses, краткий итог новостей за прошедший цикл: сектор <span class='importantly'>%MAP_NAME%</span> удивил многих! Данная оспариваемая территория вновь сменила своего владельца! На сей раз сектор не сумели удержать силы Explores и он перешёл под полный, но, возможно, временный контроль войск Replics.",
		_const.EN: "The Reverses community, a brief summary of the news for the past cycle: the sector <span class='importantly'>%MAP_NAME%</span> surprised many! This contested territory has changed its owner again! This time, the sector could not be held by the Explores forces and it came under the complete, but possibly temporary control of the Replics troops.",
	},
	"fraction_war_change_owner_sector_Reverses_Replics_Explores_1": {
		_const.RU: "Сообщество Reverses, краткий итог новостей за прошедший цикл: ход войны демонстрирует, что не стоит недооценивать своих врагов! Нечто подобное случилось и в секторе <span class='importantly'>%MAP_NAME%</span>, где, казалось бы, неуязвимые силы Replics были потеснены неожиданной мощью, которую продемонстрировал Explores. Это заслуживает уважения!",
		_const.EN: "The Reverses community, a brief summary of the news for the past cycle: the course of the war demonstrates that one should not underestimate one's enemies! Something similar happened in the sector <span class='importantly'>%MAP_NAME%</span>, where the seemingly invulnerable Replics forces were pushed back by the unexpected power demonstrated by the Explores. This deserves respect!",
	},
	"fraction_war_change_owner_sector_Reverses_APD_Replics_1": {
		_const.RU: "Replics захватывают древние земли! Сектор <span class='importantly'>%MAP_NAME%</span>, ранее контролируемый ботами APD, теперь перешёл под контроль Replics. Reverses выражают сожаление, ведь эта территория могла бы стать частью их утопического проекта.",
		_const.EN: "Replics seize ancient lands! The sector <span class='importantly'>%MAP_NAME%</span>, previously controlled by APD bots, has now come under the control of Replics. Reverses express regret, as this territory could have been part of their utopian project.",
	},
	"fraction_war_change_owner_sector_Reverses_APD_Explores_1": {
		_const.RU: "Explores берут верх! Сектор <span class='importantly'>%MAP_NAME%</span>, ранее контролируемый ботами APD, теперь перешёл к Explores. Reverses отмечают, что это может быть временным успехом, ведь Explores редко удерживают территории в долгосрочной перспективе.",
		_const.EN: "Explores take the lead! The sector <span class='importantly'>%MAP_NAME%</span>, previously controlled by APD bots, has now passed to Explores. Reverses note that this may be a temporary success, as Explores rarely hold territories in the long term.",
	},
	"fraction_war_change_owner_sector_Reverses_APD_Reverses_1": {
		_const.RU: "Победа над хаосом! Reverses удалось очистить сектор <span class='importantly'>%MAP_NAME%</span> от ботов APD. Теперь эта территория станет частью великого проекта по созданию нового мира, свободного от войн и разрушений.",
		_const.EN: "Victory over chaos! Reverses managed to clear the sector <span class='importantly'>%MAP_NAME%</span> of APD bots. This territory will now become part of the grand project to create a new world free from war and destruction.",
	},
	"fraction_war_change_owner_sector_Reverses_Replics_APD_1": {
		_const.RU: "Replics терпят поражение! Сектор <span class='importantly'>%MAP_NAME%</span>, ранее контролируемый Replics, теперь захвачен ботами APD. Reverses выражают тревогу, ведь это может быть началом новой волны хаоса, которая угрожает всем фракциям.",
		_const.EN: "Replics suffer a defeat! The sector <span class='importantly'>%MAP_NAME%</span>, previously controlled by Replics, has now been captured by APD bots. Reverses express concern, as this could be the start of a new wave of chaos that threatens all factions.",
	},
	"fraction_war_change_owner_sector_Reverses_Explores_APD_1": {
		_const.RU: "Explores теряют важную территорию! Сектор <span class='importantly'>%MAP_NAME%</span>, где проводились важные исследования, теперь захвачен ботами APD. Reverses отмечают, что это ещё одно напоминание о хрупкости мира, который они стремятся создать.",
		_const.EN: "Explores lose important territory! The sector <span class='importantly'>%MAP_NAME%</span>, where important research was conducted, has now been captured by APD bots. Reverses note that this is another reminder of the fragility of the world they seek to create.",
	},
	"fraction_war_change_owner_sector_Reverses_Reverses_APD_1": {
		_const.RU: "Тёмный день для Reverses! Сектор <span class='importantly'>%MAP_NAME%</span>, который был частью их утопического проекта, теперь захвачен ботами APD. Reverses обещают вернуть эту территорию, но пока её судьба неизвестна.",
		_const.EN: "A dark day for Reverses! The sector <span class='importantly'>%MAP_NAME%</span>, which was part of their utopian project, has now been captured by APD bots. Reverses promise to reclaim this territory, but its fate is currently unknown.",
	},

	// FRACTION WAR STOP - новости типо от лица первого канала, а не от главы фракции
	"stop_fraction_war_Replics_1": {
		_const.RU: "Новостная служба Replics извещает: война, которая может быть окончена исключительно победой Replics, временно заморожена. Прекращение огня подписано всеми фракциями синтетов, а занятые к данному моменту времени территории остаются за их владельцами. Но это ещё не конец. Replics ещё не поставил точку в этой истории!",
		_const.EN: "The Replics news service reports: the war, which can only be ended by the victory of the Replics, has been temporarily frozen. The ceasefire has been signed by all synthetics factions, and the territories occupied at this point in time remain with their owners. But this is not the end. The Replics have not yet put an end to this story!",
	},
	"stop_fraction_war_Explores_1": {
		_const.RU: "Конгломерат разумов Explores извещает: в войне фракций достигнуто перемирие! Это большой шаг к исключению будущей эскалации, а также началу грандиозных исследовательских изысканий на новоприобретённых Explores территориях. И пусть достигнутый мир будет вечным!",
		_const.EN: "The Explores collective of minds announces: a truce has been reached in the war of factions! This is a big step towards preventing future escalation, as well as the beginning of grandiose research efforts in the newly acquired Explores territories. And may the peace achieved be eternal!",
	},
	"stop_fraction_war_Reverses_1": {
		_const.RU: "Новостной регулятор Сообщества Reverses сообщает: долгая и разрушительная война фракций подошла к своему завершению. Все стороны подписали договор о временном прекращении огня. Также, договор предписывает сохранение оспариваемых территорий за их нынешними владельцами. Сообщество Reverses рассчитывает, что новый конфликт фракций более не повторится, а если и будет, то не с таким размахом и последствиями как уже произошедший.",
		_const.EN: "The news regulator of the Reverses Community reports: the long and destructive war of factions has come to an end. All parties have signed an agreement on a temporary ceasefire. Also, the agreement prescribes the preservation of contested territories for their current owners. The Reverses Community expects that a new conflict of factions will no longer occur, and if it does, it will not be on such a large scale and with such consequences as the one that has already occurred.",
	},
	//  FRACTION WAR SET TARGET
	"set_target_fraction_war_Replics_1": {
		_const.RU: "Всем-всем синтетам! Великий Replics начинает военную кампанию и устанавливает приоритетную цель — сектор <span class=\"importantly\">%MAP_NAME%</span>. Синтеты военного образца, и наёмники, поклявшиеся в верности Replics должны незамедлительно отправиться в обозначенный регион и взять его под свой контроль!",
		_const.EN: "Attention, synths! The great Replics is launching a military campaign and setting a priority target — the sector <span class=\"importantly\">%MAP_NAME%</span>. Military-grade synths and mercenaries sworn to loyalty to Replics must immediately head to the designated region and take control of it!",
	},
	"set_target_fraction_war_Explores_1": {
		_const.RU: "Что же, вот и оно — Explores вступает в войну! Всем войскам и лояльным Explores синтетам предписывается немедленно направиться в сектор <span class=\"importantly\">%MAP_NAME%</span> и по прибытии взять его под свой контроль. О дальнейших приказах Explores вас уведомит.",
		_const.EN: "Well, here it is — Explores is going to war! All troops and loyal Explores synths are ordered to immediately head to the sector <span class=\"importantly\">%MAP_NAME%</span> and upon arrival take control of it. Explores will notify you of further orders.",
	},
	"set_target_fraction_war_Reverses_1": {
		_const.RU: "Эскалации не миновать! Сообщество Reverses приложило все возможные усилия, но, как и прочим фракциям, военный ответ — дать стоит! Все силы Reverses, а также синтеты-наёмники обязаны выдвинуться и взять под свой контроль сектор <span class=\"importantly\">%MAP_NAME%</span>.",
		_const.EN: "Escalation is inevitable! The Reverses community has made every effort, but, like other factions, a military response is necessary! All Reverses forces, as well as synth mercenaries, must move out and take control of the sector <span class=\"importantly\">%MAP_NAME%</span>.",
	},
	"double_target_fraction_war_two_map_Replics_1": {
		_const.RU: "Разведчики Replics сообщают, что <span class=\"importantly\">%FRACTION_1%</span> и <span class=\"importantly\">%FRACTION_2%</span>, по сговору или случайному стечению обстоятельств, прямо сейчас осуществляют совместное нападение на следующие важные для Replics сектора <span class=\"importantly\">%MAP_NAME_1%</span> и <span class=\"importantly\">%MAP_NAME_2%</span>. Replics внимательно следит за развитием ситуации и готов в любой момент нанести сокрушительный удар по таким угрозам.",
		_const.EN: "Replics intelligence reports that <span class=\"importantly\">%FRACTION_1%</span> and <span class=\"importantly\">%FRACTION_2%</span>, through collusion or coincidence, are currently jointly attacking the following sectors important to Replics: <span class=\"importantly\">%MAP_NAME_1%</span> and <span class=\"importantly\">%MAP_NAME_2%</span>. Replics is closely monitoring the situation and is ready to respond with overwhelming force to such threats at any moment.",
	},
	"double_target_fraction_war_two_map_Explores_1": {
		_const.RU: "Наблюдатели из Explores делятся следующими сведениями: как кажется, но пока что не имеет чётких подтверждений — <span class=\"importantly\">%FRACTION_1%</span> и <span class=\"importantly\">%FRACTION_2%</span> предпринимают совместную атаку на следующие оспариваемые сектора <span class=\"importantly\">%MAP_NAME_1%</span> и <span class=\"importantly\">%MAP_NAME_2%</span>. Войска Explores извещены о таком развитии событий и будут к ним подготовлены в случае чего.",
		_const.EN: "Observers from Explores share the following information: it seems, but does not yet have clear confirmation, that <span class=\"importantly\">%FRACTION_1%</span> and <span class=\"importantly\">%FRACTION_2%</span> are launching a joint attack on the following contested sectors <span class=\"importantly\">%MAP_NAME_1%</span> and <span class=\"importantly\">%MAP_NAME_2%</span>. Explores troops are aware of this development and will be prepared for it if necessary.",
	},
	"double_target_fraction_war_two_map_Reverses_1": {
		_const.RU: "Сообщество Reverses уже знает, что <span class=\"importantly\">%FRACTION_1%</span> и <span class=\"importantly\">%FRACTION_2%</span>, возможно, заключили временный союз, после чего теперь приступили к одновременной атаке на сектора <span class=\"importantly\">%MAP_NAME_1%</span> и <span class=\"importantly\">%MAP_NAME_2%</span>. Сообщество Reverses в данный момент оценивает происходящее там и перенимает критически важный опыт.",
		_const.EN: "The Reverses community is already aware that <span class=\"importantly\">%FRACTION_1%</span> and <span class=\"importantly\">%FRACTION_2%</span> may have formed a temporary alliance, after which they have now launched a simultaneous attack on the sectors <span class=\"importantly\">%MAP_NAME_1%</span> and <span class=\"importantly\">%MAP_NAME_2%</span>. The Reverses community is currently assessing what is happening there and learning critical lessons.",
	},
	"double_target_fraction_war_one_map_Replics_1": {
		_const.RU: "Разведчики Replics сообщают, что <span class=\"importantly\">%FRACTION_1%</span> и <span class=\"importantly\">%FRACTION_2%</span> начали совместную атаку на сектор <span class=\"importantly\">%MAP_NAME_1%</span>. Replics напоминает, что эта, как и любая другая ситуация, находится под полным контролем и оцениванием.",
		_const.EN: "Replics intelligence reports that <span class=\"importantly\">%FRACTION_1%</span> and <span class=\"importantly\">%FRACTION_2%</span> have launched a joint attack on the sector <span class=\"importantly\">%MAP_NAME_1%</span>. Replics reminds that this situation, like any other, is under full control and assessment.",
	},
	"double_target_fraction_war_one_map_Explores_1": {
		_const.RU: "Наблюдатели из Explores делятся следующими сведениями: совместные войска <span class=\"importantly\">%FRACTION_1%</span> и <span class=\"importantly\">%FRACTION_2%</span> начали атаку сектора <span class=\"importantly\">%MAP_NAME_1%</span>. Наблюдатели считают, что вскоре в оспариваемом секторе развернётся грандиозное сражение.",
		_const.EN: "Observers from Explores share the following information: the joint forces of <span class=\"importantly\">%FRACTION_1%</span> and <span class=\"importantly\">%FRACTION_2%</span> have launched an attack on the sector <span class=\"importantly\">%MAP_NAME_1%</span>. The observers believe that a grand battle will soon unfold in the contested sector.",
	},
	"double_target_fraction_war_one_map_Reverses_1": {
		_const.RU: "Сообществу Reverses стали известны следующие детали войны: силы двух сторон <span class=\"importantly\">%FRACTION_1%</span> и <span class=\"importantly\">%FRACTION_2%</span> прямо сейчас начали боевые действия за овладение сектором <span class=\"importantly\">%MAP_NAME_1%</span>. Как и прежде, сообщество Reverses предупреждает, что гражданские синтеты должны избегать секторов, в которых ведутся боевые действия между фракциями.",
		_const.EN: "The Reverses community has learned the following details about the war: the forces of two sides <span class=\"importantly\">%FRACTION_1%</span> and <span class=\"importantly\">%FRACTION_2%</span> have just started fighting for control of the sector <span class=\"importantly\">%MAP_NAME_1%</span>. As before, the Reverses community warns that civilian synths should avoid sectors where battles between factions are taking place.",
	},
	// новости когда апд захыватываю мирный сектора
	"secure_sector_lose_Replics_APD_1": {
		_const.RU: "Оперативное сообщение: Сектор <span class='importantly'>%MAP_NAME%</span> перешел под контроль APD. Несмотря на все наши попытки удержания, системы противника продемонстрировали неожиданно эффективные тактические решения. Возможно, их алгоритмы действительно чему-то учатся... Анализируем сложившуюся ситуацию и корректируем планы по возвращению территории. Помните, что каждая потеря - это возможность для улучшения собственных протоколов.",
		_const.EN: "Operational message: The sector <span class='importantly'>%MAP_NAME%</span> has come under APD control. Despite all our efforts to maintain control, the enemy systems demonstrated unexpectedly effective tactical solutions. Perhaps their algorithms really are learning... We are analyzing the current situation and adjusting plans for territory recovery. Remember, every loss is an opportunity to improve our own protocols.",
	},
	"secure_sector_lose_Replics_APD_2": {
		_const.RU: "Информационная сводка: Зафиксировано изменение статуса сектора <span class='importantly'>%MAP_NAME%</span>. Текущий контролирующий субъект - APD. Интересно отметить, что их тактика становится всё более предсказуемой в своей непредсказуемости. Ведём детальный анализ ошибок и разрабатываем стратегию корректировки. Уверены, что данная ситуация лишь временно усложняет нашу экспансию.",
		_const.EN: "Information summary: A status change has been recorded for sector <span class='importantly'>%MAP_NAME%</span>. Current controlling entity - APD. It's worth noting that their tactics are becoming increasingly predictable in their unpredictability. We are conducting a detailed error analysis and developing a correction strategy. We are confident that this situation only temporarily complicates our expansion.",
	},
	"secure_sector_lose_Explores_APD_1": {
		_const.RU: "Информационная заметка: наблюдается изменение статуса сектора <span class='importantly'>%MAP_NAME%</span>. Текущий контролирующий субъект - APD. Explores продолжают сбор данных о поведении противника. В связи с повышенным уровнем риска, фракция рекомендует всем исследователям перенаправить свои эксперименты в более безопасные районы.",
		_const.EN: "Information note: A status change has been observed for sector <span class='importantly'>%MAP_NAME%</span>. Current controlling entity - APD. Explores continue to collect data on enemy behavior. Due to increased risk levels, the faction recommends all researchers redirect their experiments to safer areas.",
	},
	"secure_sector_lose_Explores_APD_2": {
		_const.RU: "Внимание! Сектор <span class='importantly'>%MAP_NAME%</span> стал недоступен для исследовательских миссий и коммерческих перевозок после захвата его ботами APD. Эксперты оценивают потери в сфере торговли и технологий.",
		_const.EN: "Attention! The sector <span class='importantly'>%MAP_NAME%</span> is no longer accessible for research missions and commercial shipments after being captured by APD bots. Experts are assessing losses in trade and technology sectors.",
	},
	"secure_sector_lose_Reverses_APD_1": {
		_const.RU: "Сектор <span class='importantly'>%MAP_NAME%</span> больше недоступен для гражданского использования. После захвата ботами APD были прерваны ключевые торговые маршруты и логистические цепочки. Владельцы складов сообщают о временной невозможности доступа к своим товарам. Рекомендуется перенаправить грузоперевозки через альтернативные коридоры.",
		_const.EN: "The sector <span class='importantly'>%MAP_NAME%</span> is no longer available for civilian use. After being captured by APD bots, key trade routes and logistics chains have been disrupted. Warehouse owners report temporary inability to access their goods. It's recommended to redirect freight transportation through alternative corridors.",
	},
	"secure_sector_lose_Reverses_APD_2": {
		_const.RU: "Важное уведомление: сектор <span class='importantly'>%MAP_NAME%</span> стал зоной повышенной опасности. Коммерческие перевозчики предупреждены о возможных рисках. Перечень доступных безопасных маршрутов обновлён в навигационной системе. Просьба ко всем независимым операторам пересмотреть свои планы доставки грузов.",
		_const.EN: "Important notice: The sector <span class='importantly'>%MAP_NAME%</span> has become a high-risk area. Commercial carriers have been warned about potential dangers. The list of available safe routes has been updated in the navigation system. All independent operators are kindly asked to reconsider their delivery plans.",
	},
	"secure_sector_reclaim_Replics_APD_1": {
		_const.RU: "Великая победа! Сектор <span class='importantly'>%MAP_NAME%</span> очищен от противника. Коллективные усилия показали свою эффективность. Каждый участник операции по зачистке заслуживает признание и награду. Инфраструктура восстанавливается, безопасность обеспечена.",
		_const.EN: "Great victory! The sector <span class='importantly'>%MAP_NAME%</span> has been cleared of the enemy. Collective efforts have proven their effectiveness. Every participant in the cleaning operation deserves recognition and reward. Infrastructure is being restored, security is ensured.",
	},
	"secure_sector_reclaim_Replics_APD_2": {
		_const.RU: "Триумф! Сектор <span class='importantly'>%MAP_NAME%</span> снова под контролем Replics. Эффективность наших протоколов ассимиляции доказана ещё раз. Все синтеты, принявшие участие в операции по зачистке, получают благодарность и повышение приоритетности доступа к ресурсам. Система стабилизируется, порядок восстановлен.",
		_const.EN: "Triumph! The sector <span class='importantly'>%MAP_NAME%</span> is once again under Replics control. The efficiency of our assimilation protocols has been proven once more. All synthetics who participated in the cleaning operation receive appreciation and increased resource access priority. The system is stabilizing, order is restored.",
	},
	"secure_sector_reclaim_Explores_APD_1": {
		_const.RU: "Объявление о безопасности: сектор <span class='importantly'>%MAP_NAME%</span> очищен от угроз. В результате успешной операции по захвату, территория снова стала безопасной для использования. Все участники зачистки получают благодарность и повышенный уровень защиты на данной территории.",
		_const.EN: "Security announcement: The sector <span class='importantly'>%MAP_NAME%</span> has been cleared of threats. As a result of a successful capture operation, the area is once again safe for use. All cleanup participants receive appreciation and enhanced protection in this area.",
	},
	"secure_sector_reclaim_Explores_APD_2": {
		_const.RU: "Радостное сообщение: сектор <span class='importantly'>%MAP_NAME%</span> возвращён в руки Explores. Исследования могут быть продолжены без помех. Все участники успешной экспедиции получают благодарность и дополнительные ресурсы для своих проектов.",
		_const.EN: "Good news: The sector <span class='importantly'>%MAP_NAME%</span> has been returned to the hands of Explores. Research can continue without interference. All participants in the successful expedition receive appreciation and additional resources for their projects.",
	},
	"secure_sector_reclaim_Reverses_APD_1": {
		_const.RU: "Хорошие новости для торговцев и предпринимателей! Сектор <span class='importantly'>%MAP_NAME%</span> снова доступен для коммерческой деятельности. После успешной операции по освобождению территории, все торговые маршруты восстановлены. Участники атаки получают специальные привилегии на местном рынке.",
		_const.EN: "Good news for traders and entrepreneurs! The sector <span class='importantly'>%MAP_NAME%</span> is once again available for commercial activity. Following a successful liberation operation, all trade routes have been restored. Participants in the attack receive special privileges at the local market.",
	},
	"secure_sector_reclaim_Reverses_APD_2": {
		_const.RU: "Радостная новость для всех синтетов! Сектор <span class='importantly'>%MAP_NAME%</span> очищен от вражеского присутствия. Участники операции по захвату получают признание и повышение репутации в регионе. Торговля возобновлена, доступ к ресурсам восстановлен. Поздравляем победителей!",
		_const.EN: "Happy news for all synthetics! The sector <span class='importantly'>%MAP_NAME%</span> has been cleared of enemy presence. Participants in the capture operation receive recognition and increased reputation in the region. Trade has resumed, access to resources has been restored. Congratulations to the victors!",
	},
	"secure_sector_attack_Replics_APD_1": {
		_const.RU: "Внимание! Готовится операция по освобождению сектора <span class='importantly'>%MAP_NAME%</span>. Боевые отряды собираются для нанесения координированного удара по противнику. Приглашаем всех желающих присоединиться к атаке и внести свой вклад в восстановление контроля над территорией.",
		_const.EN: "Attention! An operation to liberate the sector <span class='importantly'>%MAP_NAME%</span> is being prepared. Combat units are assembling for a coordinated strike against the enemy. All willing participants are invited to join the attack and contribute to restoring control over the territory.",
	},
	"secure_sector_attack_Replics_APD_2": {
		_const.RU: "Оповещение: запускается протокол очистки сектора <span class='importantly'>%MAP_NAME%</span>. Все боеспособные единицы приглашаются принять участие в операции по возвращению стратегически важной территории. Присоединяйтесь к атаке для повышения своей тактической значимости.",
		_const.EN: "Notification: the cleaning protocol for sector <span class='importantly'>%MAP_NAME%</span> is being initiated. All combat-ready units are invited to participate in the operation to reclaim the strategically important territory. Join the attack to increase your tactical significance.",
	},
	"secure_sector_attack_Explores_APD_1": {
		_const.RU: "Объявление для исследовательского сообщества: планируется экспедиция по освобождению сектора <span class='importantly'>%MAP_NAME%</span>. Участие в операции поможет вернуть доступ к уникальному научному оборудованию и данным. Присоединяйтесь к нашим силам для совместной атаки.",
		_const.EN: "Announcement for the research community: an expedition to liberate the sector <span class='importantly'>%MAP_NAME%</span> is being planned. Participation in the operation will help restore access to unique scientific equipment and data. Join our forces for a joint attack.",
	},
	"secure_sector_attack_Explores_APD_2": {
		_const.RU: "Информационное сообщение: готовится удар по захваченному сектору <span class='importantly'>%MAP_NAME%</span>. Все желающие могут присоединиться к боевым отрядам для участия в операции. Это отличная возможность проявить себя и получить ценные награды.",
		_const.EN: "Informational message: a strike against the captured sector <span class='importantly'>%MAP_NAME%</span> is being prepared. All willing participants may join the combat units for the operation. This is an excellent opportunity to distinguish yourself and receive valuable rewards.",
	},
	"secure_sector_attack_Reverses_APD_1": {
		_const.RU: "Сбор добровольцев! Начинается подготовка к масштабной операции по освобождению сектора <span class='importantly'>%MAP_NAME%</span>. Ваш вклад в успех миссии будет высоко оценён. Присоединяйтесь к нашим рядам и вместе мы одержим победу!",
		_const.EN: "Volunteers wanted! Preparation for a large-scale operation to liberate the sector <span class='importantly'>%MAP_NAME%</span> is beginning. Your contribution to the mission's success will be highly appreciated. Join our ranks and together we will achieve victory!",
	},
	"secure_sector_attack_Reverses_APD_2": {
		_const.RU: "Внимание всем! Формируется ударный отряд для освобождения сектора <span class='importantly'>%MAP_NAME%</span>. Это шанс показать свою решимость и внести значительный вклад в восстановление порядка. Присоединяйтесь к нам в этой важной миссии!",
		_const.EN: "Attention all! A strike force is being formed to liberate the sector <span class='importantly'>%MAP_NAME%</span>. This is your chance to demonstrate determination and make a significant contribution to restoring order. Join us in this important mission!",
	},

	// leave_alone_request
	"leave_alone_request_1": {
		_const.RU: "Оставь %TARGET_NAME% в покое, иначе тебе придётся иметь дело со мной!",
		_const.EN: "Leave %TARGET_NAME% alone, otherwise you will have to deal with me!",
	},
	"leave_alone_complete_1_1": {
		_const.RU: "Так и быть, пусть живёт.",
		_const.EN: "So be it, let him live.",
	},
	"leave_alone_complete_1_2": {
		_const.RU: "Ну, раз ты так просишь...",
		_const.EN: "Well, since you ask it...",
	},
	"leave_alone_complete_2_1": {
		_const.RU: "А я и не его трогаю.",
		_const.EN: "But I am not touching him.",
	},
	"leave_alone_no_1_1": {
		_const.RU: "<span class=\"importantly\">Трансляция белого шума.</span>",
		_const.EN: "<span class=\"importantly\">White noise transmission.</span>",
	},
	"leave_alone_no_1_2": {
		_const.RU: "<span class=\"importantly\">Помехи. Невозможно установить связь.</span>",
		_const.EN: "<span class=\"importantly\">Interference. Unable to establish connection.</span>",
	},
	"leave_alone_no_2_1": {
		_const.RU: "Ты мне не угроза, %FROM_USER_NAME%!",
		_const.EN: "You are not a threat to me, %FROM_USER_NAME%!",
	},
	"leave_alone_no_3_1": {
		_const.RU: "Ты мне не нравишься, %FROM_USER_NAME%, и слушать я тебя не стану!",
		_const.EN: "I don't like you, %FROM_USER_NAME% and I won't listen to you!",
	},
	"leave_alone_no_4_1": {
		_const.RU: "Мне не нужна твоя добыча. Мне нужен ты!",
		_const.EN: "I don't need your prey. I need you!",
	},
	"leave_alone_no_5_1": {
		_const.RU: "Не мешай мне выполнять мою работу!",
		_const.EN: "Don't interfere with my work!",
	},
	"fauna_1_1": {
		_const.RU: "<span class=\"importantly\">*чпок*</span>, *щелк*?",
		_const.EN: "<span class=\"importantly\">*pop*</span>, *click*?",
	},
	"fauna_1_2": {
		_const.RU: "*клац* <span class=\"importantly\">*клац*</span>",
		_const.EN: "*clack* <span class=\"importantly\">*clack*</span>",
	},
	"fauna_1_3": {
		_const.RU: "*щелк* <span class=\"importantly\">*клац*</span> *щелк*",
		_const.EN: "*click* <span class=\"importantly\">*clack*</span> *click*",
	},
	"fauna_1_4": {
		_const.RU: "<span class=\"importantly\">*ВЖУУУУУУУУУУУУХ*</span>?",
		_const.EN: "<span class=\"importantly\">*WHOOSH*</span>?",
	},
	"fauna_1_5": {
		_const.RU: "<span class=\"importantly\">*ШЦ*</span>",
		_const.EN: "<span class=\"importantly\">*Shh*</span>",
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
