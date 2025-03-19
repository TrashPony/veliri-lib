package handbook

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"math/rand"
	"strings"
	"time"
)

// TODO поделить нотифаи по файлам и типам, а то начнутся коллизии ключей и все сломается
var dialogTypes = map[string]map[string]string{
	// HELP TASK
	"help_request_1": {
		_const.RU:   "Эй, %TO_USER_NAME% мне бы пригодился проводник в секторе. Вдвоём-то куда веселее, что об этом думаешь?",
		_const.EN:   "Hey, %TO_USER_NAME%, I could use a guide in the sector. It's much more fun together, what do you think?",
		_const.ZhCN: "%TO_USER_NAME%，我需要一位向导来探索这个区域。一起行动更有趣，你觉得呢？",
	},
	"help_request_2": {
		_const.RU:   "Что думаешь касательного того, чтобы путешествовать вместе, а %TO_USER_NAME%?",
		_const.EN:   "What do you think about traveling together, %TO_USER_NAME%?",
		_const.ZhCN: "%TO_USER_NAME%，你觉得结伴旅行怎么样？",
	},
	"help_request_3": {
		_const.RU:   "Я исследователь и мне бы не помешал телохранитель. Понимаешь, к чему я клоню %TO_USER_NAME%?",
		_const.EN:   "I'm a researcher and I could use a bodyguard. Do you understand what I'm getting at, %TO_USER_NAME%?",
		_const.ZhCN: "我是研究人员，需要一名保镖。你明白我的意思吧，%TO_USER_NAME%？",
	},
	"help_busy_1_1": {
		_const.RU:   "Да ты хоть знаешь, к кому обращаешься!? А... Мне ведь и говорить об этом нельзя. Забудь.",
		_const.EN:   "Do you even know who you're talking to!? Ah... I'm not even supposed to talk about it. Forget it.",
		_const.ZhCN: "你知不知道自己在跟谁说话！？啊……我甚至不该提这件事。忘了吧。",
	},
	"help_busy_1_2": {
		_const.RU:   "Я на задании. Оно важно и ответственно. Кстати, ты не нанесёшь мне на карту расположение ближайших военных баз?",
		_const.EN:   "I'm on a mission. It's important and responsible. By the way, can you mark the location of the nearest military bases on my map?",
		_const.ZhCN: "我正在执行重要任务。对了，你能在地图上标出最近的军事基地吗？",
	},
	"help_busy_1_3": {
		_const.RU:   "У меня вообще-то важная миссия. Но тебя это не касается. Скажем... здесь столкнулись разные интересы.",
		_const.EN:   "I actually have an important mission. But it's not your concern. Let's just say... different interests collided here.",
		_const.ZhCN: "我有重要任务在身，但与你无关。只能说……这里有多方势力在博弈。",
	},
	"help_busy_2_1": {
		_const.RU:   "Как бы я занят - выполняю важную миссию. А вдруг здесь где поблизости боты-шпионы? А вдруг ты один из них?",
		_const.EN:   "I'm so busy - I'm on an important mission. What if there are spy bots nearby? What if you're one of them?",
		_const.ZhCN: "我很忙——正在执行重要任务。万一附近有间谍机器人呢？万一你就是其中之一？",
	},
	"help_busy_2_2": {
		_const.RU:   "Неа, тут уж ты сам. Мне оно ни к чему.",
		_const.EN:   "Nope, you're on your own here. I don't need it.",
		_const.ZhCN: "不，你自便吧。我不需要。",
	},
	"help_busy_2_3": {
		_const.RU:   "Мне это неинтересно. Без обид.",
		_const.EN:   "It's not interesting to me. No offense.",
		_const.ZhCN: "我没兴趣。无意冒犯。",
	},
	"help_busy_2_4": {
		_const.RU:   "Ха! Ищи дурака! А вдруг ты просто желаешь меня ограбить? Или... ещё чего другого? Так я на это и повёлся.",
		_const.EN:   "Ha! Look for a fool! What if you just want to rob me? Or... something else? That's how I fell for it.",
		_const.ZhCN: "哈！想找冤大头？说不定你只是想打劫我？或者……另有所图？我才不会上当。",
	},
	"help_busy_3_": {
		_const.RU:   "Э-э, что-то ты выглядишь слишком подозрительным! Ну уж нет.",
		_const.EN:   "Uh, you look too suspicious! No way.",
		_const.ZhCN: "呃，你看起来太可疑了！绝对不行。",
	},
	"help_busy_4_1": {
		_const.RU:   "Ха! Ищи дурака! А вдруг ты просто желаешь меня ограбить? Или... ещё чего другого? Так я на это и повёлся.",
		_const.EN:   "Ha! Look for a fool! What if you just want to rob me? Or... something else? That's how I fell for it.",
		_const.ZhCN: "哈！想找冤大头？说不定你只是想打劫我？或者……另有所图？我才不会上当。",
	},
	"help_busy_4_2": {
		_const.RU:   "Нет уж. Здесь каждый сам за себя! Без обид %FROM_USER_NAME%, но здесь наши дороги расходятся. Надеюсь, что без последствий.",
		_const.EN:   "No way. Everyone is out for themselves here! No offense %FROM_USER_NAME%, but we go our separate ways here. I hope there are no consequences.",
		_const.ZhCN: "没门。这里人人自危！%FROM_USER_NAME%别介意，但我们各走各路吧。希望不会有什么后果。",
	},
	"help_busy_4_3": {
		_const.RU:   "Ты плохо врёшь! Это же фирменная ловушка! Лучше пойди поищи дураков в другом месте.",
		_const.EN:   "You're a bad liar! This is a classic trap! Better go find fools elsewhere.",
		_const.ZhCN: "你撒谎水平太差了！这是经典陷阱！去别处找傻子吧。",
	},
	"help_busy_5_1": {
		_const.RU:   "Неа, тут уж ты сам. Мне оно ни к чему.",
		_const.EN:   "Nope, you're on your own here. I don't need it.",
		_const.ZhCN: "不，你自便吧。我不需要。",
	},
	"help_busy_5_2": {
		_const.RU:   "Мне это неинтересно. Без обид.",
		_const.EN:   "I am not interested. No offense.",
		_const.ZhCN: "我不感兴趣。无意冒犯。",
	},
	"help_busy_5_3": {
		_const.RU:   "Ты плохо врёшь! Это же фирменная ловушка! Лучше пойди поищи дураков в другом месте.",
		_const.EN:   "You are a bad liar! This is a common trap! Better go find fools elsewhere.",
		_const.ZhCN: "你撒谎太拙劣了！这是常见陷阱！去别处找冤大头吧。",
	},
	"help_busy_5_4": {
		_const.RU:   "Ха! Ищи дурака! А вдруг ты просто желаешь меня ограбить? Или... ещё чего другого? Так я на это и повёлся.",
		_const.EN:   "Ha! Look for a fool! What if you just want to rob me? Or... something else? I am not falling for it.",
		_const.ZhCN: "哈！想找冤大头？说不定你只是想打劫我？或者……另有所图？我才不会上当。",
	},
	"help_busy_6_1": {
		_const.RU:   "Моя миссия превыше этого.",
		_const.EN:   "My mission is more important than that.",
		_const.ZhCN: "我的任务比这更重要。",
	},
	"help_busy_6_2": {
		_const.RU:   "Вы не член здешней главенствующей фракции. В запросе отказано.",
		_const.EN:   "You are not a member of the local ruling faction. The request is denied.",
		_const.ZhCN: "你不是本地统治派系的成员。请求被拒绝。",
	},
	"help_busy_6_3": {
		_const.RU:   "На вас не распространяются условия о защите и сопровождении. Ради собственного блага - отъедьте от меня на некоторое расстояние.",
		_const.EN:   "The conditions on protection and escort do not apply to you. For your own good — move some distance away from me.",
		_const.ZhCN: "保护与护送条款不适用于你。为了你好——请与我保持距离。",
	},
	"help_busy_7_1": {
		_const.RU:   "<span class=\"importantly\">Ошибка 415.</span>",
		_const.EN:   "<span class=\"importantly\">Error 415.</span>",
		_const.ZhCN: "<span class=\"importantly\">错误415。</span>",
	},
	"help_busy_7_2": {
		_const.RU:   "<span class=\"importantly\">Сигнал неопознан.</span>",
		_const.EN:   "<span class=\"importantly\">Signal is unidentified.</span>",
		_const.ZhCN: "<span class=\"importantly\">信号未识别。</span>",
	},
	"help_busy_7_3": {
		_const.RU:   "<span class=\"importantly\">333./444-55,16?</span>",
		_const.EN:   "<span class=\"importantly\">333./444-55,16?</span>",
		_const.ZhCN: "<span class=\"importantly\">333./444-55,16？</span>",
	},
	"help_busy_8_": {
		_const.RU:   "Ещё чего! Я не стану себя подвергать угрозе.",
		_const.EN:   "What do you think? I'm not going to put myself at risk.",
		_const.ZhCN: "想得美！我不会让自己陷入危险。",
	},
	"help_busy_9_1": {
		_const.RU:   "Ты что, не видишь, чем я занят? Мне ещё предстоит отчитаться, сколько я добыл. А если вдруг выяснится, что пары граммов не хватает? Знаешь, что со мной сделают?",
		_const.EN:   "Can't you see what I'm doing? I still have to report how much I've mined. And what if it turns out that a couple of grams are missing? Do you know what they'll do to me?",
		_const.ZhCN: "你没看见我在忙吗？我还得汇报开采量。万一少了几克怎么办？你知道他们会怎么处置我吗？",
	},
	"help_busy_9_2": {
		_const.RU:   "Мне некогда. Я столько времени искал богатое на месторождения место, чтобы теперь всё бросить и отправиться с тобой.",
		_const.EN:   "I don't have time. I spent so much time looking for a place rich in deposits that now I can't just drop everything and go with you.",
		_const.ZhCN: "我没空。我花了好久才找到这个富矿区，现在不可能抛下一切跟你走。",
	},
	"help_busy_9_3": {
		_const.RU:   "Руда себя не выкопает. А я, как раз собирался немного подзаработать. Извини, тут уж ты сам по себе.",
		_const.EN:   "The ore won't dig itself out. And I was just about to earn some money. Sorry, you're on your own here.",
		_const.ZhCN: "矿石不会自己蹦出来。我正准备赚点外快呢。抱歉，你自求多福吧。",
	},
	"help_busy_9_4": {
		_const.RU:   "Э! Да я только приобрёл это горнодобывающее оборудование! Ты серьёзно рассчитываешь, что я брошу попытку опробовать его только ради тебя? Ха! Ха-ха!",
		_const.EN:   "Hey, I've just bought this mining equipment! Do you really expect me to give up trying it out just for you? Ha! Ha-ha!",
		_const.ZhCN: "嘿！我刚买的采矿设备！你真以为我会为了你放弃试用？哈！哈哈！",
	},
	"help_busy_9_5": {
		_const.RU:   "А может напротив - ты мне составишь компанию в добыче полезных ресурсов? Нет... Эх. Опять всё придётся делать самому. Ну, бывай тогда.",
		_const.EN:   "Or maybe, on the contrary, you will keep me company in mining useful resources? No... Oh well. Again, I have to do everything myself. Well, take care.",
		_const.ZhCN: "或者反过来——你陪我采矿怎么样？算了……唉。又得自己动手。那就再见了。",
	},
	"help_busy_9_6": {
		_const.RU:   "У меня контракт на разведку и добычу ресурсов. Как освобожусь через три - четыре месяца, я тебе дам знать.",
		_const.EN:   "I have a contract for exploration and resource extraction. When I'm free in three or four months, I'll let you know.",
		_const.ZhCN: "我有勘探开采合同。等三四个月后有空了，我会联系你。",
	},
	"help_busy_9_7": {
		_const.RU:   "Не, не, не! Совсем нету на это времени. Я сейчас направляюсь к одному месторождению, а ты, пытаешься отвлечь меня и, чтобы пока я отсутствую, оттуда уже всё выгребли своими грязными ковшами? Такому не бывать!",
		_const.EN:   "No, no, no! I don't have time for that at all. I'm heading to a deposit right now, and you're trying to distract me so that while I'm away, they've already taken everything out with their dirty buckets? That's not going to happen!",
		_const.ZhCN: "不不不！完全没时间。我正要去一个矿点，你却想支开我，好趁我不在用脏铲子偷光资源？休想！",
	},
	"help_busy_10_1": {
		_const.RU:   "Это всё конечно замечательно, но я выполняю важную миссию. Скажу тебе по секрету - мне нужно разведать местность.",
		_const.EN:   "This is all great, of course, but I'm on an important mission. I'll tell you a secret — I need to explore the area.",
		_const.ZhCN: "听起来不错，但我在执行重要任务。偷偷告诉你——我需要勘察这片区域。",
	},
	"help_busy_10_2": {
		_const.RU:   "Я следую к сигналу бедствия. Сейчас не до тебя.",
		_const.EN:   "I'm following a distress signal. I don't have time for you now.",
		_const.ZhCN: "我正在追踪求救信号。现在没空理你。",
	},
	"help_busy_10_3": {
		_const.RU:   "Я двигаюсь по распоряжению фракции к одному источнику сигнала. Поэтому, наши пути не могут совпадать.",
		_const.EN:   "I'm moving to a signal source by order of my faction. So, our paths can't coincide.",
		_const.ZhCN: "我正按派系命令前往信号源。所以我们不可能同路。",
	},
	"help_busy_10_4": {
		_const.RU:   "Здесь неподалёку идёт трансляция сигнала бедствия. Я следую туда по спасательной миссии.",
		_const.EN:   "There is a distress signal being broadcast nearby. I'm going there on a rescue mission.",
		_const.ZhCN: "附近有求救信号广播。我要去执行救援任务。",
	},
	"help_busy_11_4": {
		_const.RU:   "Ты должно быть шутишь?",
		_const.EN:   "You must be joking?",
		_const.ZhCN: "你是在开玩笑吧？",
	},

	// HELP TASK
	"help_busy_12_1": {
		_const.RU:   "Сейчас я участвую в бою. Обратитесь позже!",
		_const.EN:   "I'm in the middle of a battle right now. Come back later!",
		_const.ZhCN: "我正在战斗中。稍后再来！",
	},
	"help_complete_1": {
		_const.RU:   "Хм. Ладно, но я внимательно за тобой слежу. Впереди, кстати, будешь ты!",
		_const.EN:   "Hmm. Okay, but I'm keeping a close eye on you. By the way, you're next!",
		_const.ZhCN: "嗯，好吧——但我会盯着你的。顺便说一句，你走前面！",
	},
	"help_complete_2": {
		_const.RU:   "Я согласен. Дорога трудна, а друг-Синтет - в пустошах редкое явление.",
		_const.EN:   "I agree. The road is difficult, and a Synth friend is a rare sight in the wasteland.",
		_const.ZhCN: "我同意。废土之路艰难，而合成体同伴更是罕见。",
	},
	"help_complete_3": {
		_const.RU:   "Здравая идея! Ладно. Сработаемся!",
		_const.EN:   "A sound idea! Alright. We'll work together!",
		_const.ZhCN: "明智的建议！行，我们合作吧！",
	},
	"help_complete_4": {
		_const.RU:   "Добро! За дело!",
		_const.EN:   "Good! Let's get to work!",
		_const.ZhCN: "好！开始行动！",
	},
	"leave_help_dialog_1": {
		_const.RU:   "У меня есть дела поважнее, чем таскаться с тобой по пустошам.",
		_const.EN:   "I have more important things to do than to wander the wasteland with you.",
		_const.ZhCN: "比起陪你逛废土，我有更重要的事要做。",
	},
	"leave_help_dialog_2": {
		_const.RU:   "А вообще, знаешь, мои планы изменились. Забудь. Разве что только, ты не желаешь.меня как-нибудь отблагодарить.",
		_const.EN:   "Actually, you know, my plans have changed. Forget it. Unless, you want to thank me somehow.",
		_const.ZhCN: "其实……我的计划有变。算了吧。除非你愿意给我点谢礼？",
	},
	"leave_help_not_friend_dialog_1": {
		_const.RU:   "Мне стало известно, о том кто ты и чем промышляешь. Зря я пожалуй снизошёл до диалога с тобой. Прощай.",
		_const.EN:   "I found out who you are and what you do. I probably shouldn't have condescended to talk to you. Goodbye.",
		_const.ZhCN: "我已查明你的身份和行径。与你对话是我失策。永别了。",
	},
	"leave_help_not_friend_dialog_2": {
		_const.RU:   "Я лучше найду кого-нибудь другого. Есть мнение, что наше знакомство до добра не доведёт.",
		_const.EN:   "I'd better find someone else. There is an opinion that our acquaintance will not lead to anything good.",
		_const.ZhCN: "我会另寻他人。有传言说我们的相识不会有好结果。",
	},

	// MAKE_FRIEND
	"make_friends_request_1": {
		_const.RU:   "Пора заканчивать эту глупую конфронтацию, %TO_USER_NAME%! Хватит уже!",
		_const.EN:   "It's time to end this stupid confrontation, %TO_USER_NAME%! Enough already!",
		_const.ZhCN: "%TO_USER_NAME%，是时候结束这场愚蠢的对抗了！到此为止！",
	},
	"make_friends_request_2": {
		_const.RU:   "Я бы желал свести на нет любые наши разногласия %TO_USER_NAME%.",
		_const.EN:   "I would like to resolve any differences between us, %TO_USER_NAME%.",
		_const.ZhCN: "%TO_USER_NAME%，我希望化解我们之间的所有分歧。",
	},
	"make_friends_request_3": {
		_const.RU:   "Всё! Довольно! Я не желаю быть тебе врагом %TO_USER_NAME%.",
		_const.EN:   "That's it! Enough! I don't want to be your enemy, %TO_USER_NAME%.",
		_const.ZhCN: "够了！%TO_USER_NAME%，我不想与你为敌。",
	},
	"make_friends_request_4": {
		_const.RU:   "Всё в этом мире подходит к концу %TO_USER_NAME%. Рано или поздно. Так, не пора ли и нам самим закончить любые наши междоусобные претензии?",
		_const.EN:   "Everything in this world comes to an end, %TO_USER_NAME%. Sooner or later. So, isn't it time for us to end any of our internecine claims?",
		_const.ZhCN: "%TO_USER_NAME%，世间万物终有尽时。不如现在结束我们无谓的纷争？",
	},
	"make_friends_request_5": {
		_const.RU:   "Мне всё это формально надоело %TO_USER_NAME%! Предлагаю только сейчас - прекратить весь этот фарс и раазойтись друзьями.",
		_const.EN:   "I'm officially tired of all this, %TO_USER_NAME%! I suggest right now — to stop this whole farce and part as friends.",
		_const.ZhCN: "%TO_USER_NAME%，我受够这场闹剧了！立刻停止，以友相别如何？",
	},
	"make_friends_no_1_1": {
		_const.RU:   "<span class=\"importantly\">Неизвестно. Сбой сообщения.</span>",
		_const.EN:   "<span class=\"importantly\">Unknown. Message failure.</span>",
		_const.ZhCN: "<span class=\"importantly\">未知错误。通讯失败。</span>",
	},
	"make_friends_no_1_2": {
		_const.RU:   "Tty?'\\WgHANT//.. <br><br> <span class=\"importantly\">Не удалось распознать ответ.</span>",
		_const.EN:   "Tty?'\\WgHANT//.. <br><br> <span class=\"importantly\">Failed to recognize the response.</span>",
		_const.ZhCN: "Tty?'\\WgHANT//.. <br><br> <span class=\"importantly\">未能识别响应。</span>",
	},
	"make_friends_no_2_1": {
		_const.RU:   "Подобное не может быть исполнено. Вы вынуждаете нарушать правила кодекса.",
		_const.EN:   "This cannot be done. You are forcing me to break the code of conduct.",
		_const.ZhCN: "请求无法执行。你这是在迫使我违反行为准则。",
	},
	"make_friends_no_2_2": {
		_const.RU:   "Отказано! Любые попытки последующих контактов приведут к усилению огневого воздействия.",
		_const.EN:   "Refused! Any further attempts at contact will result in increased firepower.",
		_const.ZhCN: "拒绝！继续联络将招致更猛烈的火力打击。",
	},
	"make_friends_no_2_3": {
		_const.RU:   "Ваш сигнал нераспознан. Повторите попытку вновь. Последующие неудачи в связи, могут привести к санкциям.",
		_const.EN:   "Your signal is unrecognized. Try again. Subsequent communication failures may result in sanctions.",
		_const.ZhCN: "信号无法识别。请重试。持续通讯失败将触发制裁协议。",
	},
	"make_friends_no_3": {
		_const.RU:   "А у нас что, с тобой были раньше какие-то разногласия?",
		_const.EN:   "Have we had any disagreements before?",
		_const.ZhCN: "我们之前有过节吗？",
	},
	"make_friends_no_4_1": {
		_const.RU:   "Бейся до конца! Покажи из чего ты сделан!",
		_const.EN:   "Fight to the end! Show what you're made of!",
		_const.ZhCN: "战斗到底！让我看看你的本事！",
	},
	"make_friends_no_4_2": {
		_const.RU:   "Ты от меня так просто не отделаешься %FROM_USER_NAME%!",
		_const.EN:   "You won't get rid of me that easily %FROM_USER_NAME%! ",
		_const.ZhCN: "%FROM_USER_NAME%！你休想轻易摆脱我！",
	},
	"make_friends_no_4_3": {
		_const.RU:   "Да у тебя должно быть микросхемы от радиации расплавились. Даже и не мечтай о подобном.",
		_const.EN:   "Your microchips must have melted from radiation. Don't even dream of such a thing.",
		_const.ZhCN: "你的芯片怕是被辐射烧坏了吧？别痴心妄想了。",
	},
	"make_friends_no_4_4": {
		_const.RU:   "Гляди чего он вздумал. И не мечтай %FROM_USER_NAME% о чём-то подобном.",
		_const.EN:   "Look what he's up to. Don't you even dream %FROM_USER_NAME% of something like that.",
		_const.ZhCN: "看看这人的妄想。%FROM_USER_NAME%，你最好连想都别想。",
	},
	"make_friends_no_5_1": {
		_const.RU:   "Ты жалок. А твои попытки примирения только ещё больше подогревают мой гнев!",
		_const.EN:   "You're pathetic. And your attempts at reconciliation only fuel my anger even more!",
		_const.ZhCN: "可悲。你虚伪的和解只会让我更加愤怒！",
	},
	"make_friends_no_5_2": {
		_const.RU:   "Ха! Что это я вижу - попытку примирения? Сейчас-то я тебя научу манерам %FROM_USER_NAME%.",
		_const.EN:   "Ha! What do I see - an attempt at reconciliation? Now I'll teach you some manners %FROM_USER_NAME%.",
		_const.ZhCN: "哈！想求和？%FROM_USER_NAME%，让我教教你规矩！",
	},
	"make_friends_no_5_3": {
		_const.RU:   "Хорошая попытка. Жаль только, что неудачная.",
		_const.EN:   "Nice try. Too bad it didn't work out.",
		_const.ZhCN: "不错的尝试。可惜失败了。",
	},
	"make_friends_pay_1": {
		_const.RU:   "Финансы, способы скрасить многие вещи %FROM_USER_NAME%. Они очень хорошо решают проблемы. Особенно, когда их платят за что-то. А ещё лучше, если они в размере - %CREDITS_COUNT% cr.",
		_const.EN:   "Finance, ways to brighten up many things %FROM_USER_NAME%. They solve problems very well. Especially when they are paid for something. And even better if they are in the amount of - %CREDITS_COUNT% cr.",
		_const.ZhCN: "%FROM_USER_NAME%，金钱能解决很多问题——尤其是%CREDITS_COUNT%信用点这个数字。",
	},
	"make_friends_pay_2": {
		_const.RU:   "Хочешь жить, %FROM_USER_NAME%? Тогда придётся раскошелиться в размере %CREDITS_COUNT% cr.",
		_const.EN:   "Do you want to live, %FROM_USER_NAME%? Then you will have to fork out in the amount of %CREDITS_COUNT% cr.",
		_const.ZhCN: "想活命吗，%FROM_USER_NAME%？那就支付%CREDITS_COUNT%信用点。",
	},
	"make_friends_pay_3": {
		_const.RU:   "Сложившаяся ситуация, толкает меня %FROM_USER_NAME%, требовать от тебя сумму в - %CREDITS_COUNT% cr. Это ради твоего же блага. Ты, главное, пойми.",
		_const.EN:   "The current situation pushes me %FROM_USER_NAME% to demand from you the amount of - %CREDITS_COUNT% cr. It's for your own good. You just have to understand that.",
		_const.ZhCN: "%FROM_USER_NAME%，形势所迫——支付%CREDITS_COUNT%信用点对你有好处。明白吗？",
	},
	"make_friends_pay_4": {
		_const.RU:   "Значит так, %FROM_USER_NAME% твоя ситуация непроста, а выбор невелик - гони сумму в %CREDITS_COUNT% cr., иначе не доползёшь и до ближайшей базы. Времени на раздумье у тебя нет.",
		_const.EN:   "So, %FROM_USER_NAME%, your situation is difficult, and the choice is small — pay up the sum of %CREDITS_COUNT%. Otherwise, you won't make it to the nearest base. You don't have time to think.",
		_const.ZhCN: "%FROM_USER_NAME%，你没得选——立刻支付%CREDITS_COUNT%信用点，否则别想活着到下一个据点。",
	},
	"make_friends_pay_5": {
		_const.RU:   "Вот ты и попался мне, %FROM_USER_NAME%. Что же, придётся отнять у тебя некоторую сумму в %CREDITS_COUNT%. Видишь, как она тебя замедляет и утяжеляет. Позволить облегчить тебе эту тяжкую ношу.",
		_const.EN:   "There you are, %FROM_USER_NAME%. I'll have to take some money from you — %CREDITS_COUNT%. See how it slows you down and makes you heavier. Let me help you lighten this heavy burden.",
		_const.ZhCN: "逮到你了，%FROM_USER_NAME%。交出%CREDITS_COUNT%信用点——这些累赘的财富只会拖累你。",
	},
	"make_friends_complete_1": {
		_const.RU:   "Ты прав - ты прав. Это не имеет никакого логического смысла.",
		_const.EN:   "You're right - you're right. It doesn't make any logical sense.",
		_const.ZhCN: "你说得对——这毫无逻辑可言。",
	},
	"make_friends_complete_2": {
		_const.RU:   "Хорошо. Но только в этот раз. В первый и последний!",
		_const.EN:   "Fine. But only this time. The first and last!",
		_const.ZhCN: "行吧——仅此一次，下不为例！",
	},
	"make_friends_complete_3": {
		_const.RU:   "Убедил, %FROM_USER_NAME%. Мне и самому нет дела до глупых обид прошлого. Что было - то прошло.",
		_const.EN:   "You've convinced me, %FROM_USER_NAME%. I don't care about stupid grudges of the past myself. What happened, happened.",
		_const.ZhCN: "你说服我了，%FROM_USER_NAME%。过去的蠢事就让它过去吧。",
	},
	"make_friends_complete_4": {
		_const.RU:   "Знаешь, что? А почему бы и нет!",
		_const.EN:   "You know what? Why not!",
		_const.ZhCN: "你知道吗？有何不可！",
	},
	"make_friends_complete_5": {
		_const.RU:   "На такое предложение я отвечу положительным вердиктом.",
		_const.EN:   "I will respond with a positive verdict to such an offer.",
		_const.ZhCN: "对此提议，我给出肯定答复。",
	},
	"make_friends_do_pay_1": {
		_const.RU:   "Да... Да! Я заплачу. Пусть будет так.",
		_const.EN:   "Yes... Yes! I'll pay. So be it.",
		_const.ZhCN: "好……好吧！我付钱。就这样吧。",
	},
	"make_friends_do_pay_2": {
		_const.RU:   "Без вопросов. Меня не стоит убеждать дважды.",
		_const.EN:   "No questions asked. I don't need to be convinced twice.",
		_const.ZhCN: "无需多言。我从不说第二遍。",
	},
	"make_friends_do_pay_3": {
		_const.RU:   "Эх. Очередные траты. Опять траты... Я согласен. Куда уж мне деваться.",
		_const.EN:   "Oh well. Another expense. More expenses... I agree. What else can I do.",
		_const.ZhCN: "唉，又是开支…我同意。还能怎么办？",
	},
	"make_friends_do_pay_4": {
		_const.RU:   "Ну, раз уж у меня нет иного выхода, то... что теперь поделаешь. Держи, свои %CREDITS_COUNT% cr.",
		_const.EN:   "Well, since I have no other choice, then... what else can I do. Here you go, take your %CREDITS_COUNT% cr.",
		_const.ZhCN: "既然别无选择……拿去吧，这是%CREDITS_COUNT%信用点。",
	},
	"make_friends_do_pay_5": {
		_const.RU:   "Ясно, что с тобой шутки плохи. Перечисляю %CREDITS_COUNT%. cr.",
		_const.EN:   "It's clear that you're not to be trifled with. I'm transferring %CREDITS_COUNT%. cr.",
		_const.ZhCN: "看来你不是在说笑。正在转账%CREDITS_COUNT%信用点。",
	},
	"make_friends_success_pay_1": {
		_const.RU:   "А вот это совсем другой разговор.",
		_const.EN:   "Now this is a different story.",
		_const.ZhCN: "这才像话。",
	},
	"make_friends_success_pay_2": {
		_const.RU:   "Как неожиданно и приятно!",
		_const.EN:   "How unexpected and pleasant!",
		_const.ZhCN: "真是意外之喜！",
	},
	"make_friends_success_pay_3": {
		_const.RU:   "Вот так бы и сразу. А то ещё думают по нескольку световых лет.",
		_const.EN:   "That's how it should have been from the start. Otherwise they think for several light-years.",
		_const.ZhCN: "早该如此。省得浪费几个光年时间。",
	},
	"make_friends_success_pay_4": {
		_const.RU:   "Хорошо, что всё кончилось именно так. И главное - никаких смертей.",
		_const.EN:   "It's good that everything ended this way. And most importantly, there were no deaths.",
		_const.ZhCN: "皆大欢喜。最重要的是没人送命。",
	},
	"make_friends_success_pay_5": {
		_const.RU:   "Молодец. Хвалю. Ты всё понял и сделал действительно верный выбор. Эти %CREDITS_COUNT% cr. я потрачу с умом. Могу даже отчёт как-нибудь прислать.",
		_const.EN:   "Well done. I praise you. You understood everything and made a really right choice. I will spend these %CREDITS_COUNT% cr. wisely. I may even send a report some time.",
		_const.ZhCN: "明智之举。这%CREDITS_COUNT%信用点我会妥善使用。说不定给你寄份支出报告？",
	},
	"make_friends_fail_pay_1_1": {
		_const.RU:   "Нет. Мои деньги останутся только при мне!",
		_const.EN:   "No. My money will remain with me!",
		_const.ZhCN: "休想！我的钱一个子儿都不会给你！",
	},
	"make_friends_fail_pay_1_2": {
		_const.RU:   "Да нисколько ты не получишь!",
		_const.EN:   "You won't get anything!",
		_const.ZhCN: "你一分钱都别想拿到！",
	},
	"make_friends_fail_pay_1_3": {
		_const.RU:   "Я что, такие суммы тебе скомбинировать должен? Мой ответ - не будет тебе такого счастья!",
		_const.EN:   "Am I supposed to combine such amounts for you? My answer is you won't be so lucky!",
		_const.ZhCN: "难道要我给你凑钱？别痴心妄想了！",
	},
	"make_friends_fail_pay_1_4": {
		_const.RU:   "А до меня кто-то на такое соглашался? Вот и я пожалуй откажусь тебе платить.",
		_const.EN:   "No one ever agreed to such a thing before me? Well, I'll probably refuse to pay you too.",
		_const.ZhCN: "前人都没答应过？那我也拒绝付款！",
	},
	"make_friends_fail_pay_1_5": {
		_const.RU:   "Это фирменный грабёж! Я на такое согласиться ну никак не могу.",
		_const.EN:   "This is outright robbery! I can't agree to this no matter what.",
		_const.ZhCN: "这是赤裸裸的抢劫！我绝不同意！",
	},
	"make_friends_fail_pay_2_1": {
		_const.RU:   "Я тебе что из воздуха должен достать средства? Сам-то подумай!",
		_const.EN:   "Am I supposed to conjure up funds out of thin air? Think about it yourself!",
		_const.ZhCN: "要我凭空变钱？你自己动动脑子！",
	},
	"make_friends_fail_pay_2_2": {
		_const.RU:   "Очень жаль, но у меня ничего не осталось. Тебе должно быть грустно, а?",
		_const.EN:   "It's a pity, but I have nothing left. You must be sad, right?",
		_const.ZhCN: "真遗憾，我身无分文。你很失望吧？",
	},
	"make_friends_fail_pay_2_3": {
		_const.RU:   "А да? Как жалко, что я тебе больше заплатить не могу. Так ведь хотелось.",
		_const.EN:   "Oh well. It's such a shame that I can't pay you more. I really wanted to.",
		_const.ZhCN: "唉，可惜我给不起更多——虽然我真的很想。",
	},
	"make_friends_failed_pay_1": {
		_const.RU:   "А-а, хочешь по плохому %FROM_USER_NAME%?",
		_const.EN:   "Oh, you want to do it the bad way %FROM_USER_NAME%?",
		_const.ZhCN: "%FROM_USER_NAME%，你想来硬的？",
	},
	"make_friends_failed_pay_2": {
		_const.RU:   "Ооо, сейчас будет такая драка!",
		_const.EN:   "Ooo, there's going to be a fight now!",
		_const.ZhCN: "哦豁！准备开战吧！",
	},
	"make_friends_failed_pay_3": {
		_const.RU:   "Ой всё! Я за себя теперь не ручаюсь!",
		_const.EN:   "Oh, that's it! I'm not responsible for my actions now!",
		_const.ZhCN: "够了！我现在可控制不住自己了！",
	},
	"make_friends_failed_pay_4": {
		_const.RU:   "Не я эту войну начал %FROM_USER_NAME%. Как жаль, что какие-то %CREDITS_COUNT% cr. для тебя настолько важны.",
		_const.EN:   "%FROM_USER_NAME%, you started this war. What a pity that some %CREDITS_COUNT% cr. are so important to you.",
		_const.ZhCN: "%FROM_USER_NAME%，是你挑起的战争。为%CREDITS_COUNT%信用点拼命？真可悲。",
	},
	"make_friends_failed_pay_5": {
		_const.RU:   "Отказываешься значит? Хе, это у меня впервые. Но и для тебя, тогда, будет в последний раз!",
		_const.EN:   "You're refusing, huh? Heh, this is a first for me. But it will also be the last time for you!",
		_const.ZhCN: "拒绝？哈，我头一回见——但这也是你最后一次！",
	},

	"make_friends_failed_no_pay_1": {
		_const.RU:   "Значит, ты выбрал смерть.",
		_const.EN:   "So you have chosen death.",
		_const.ZhCN: "你选择了死亡。",
	},

	// PAY MANY
	"pay_many_failed_pay_1_1": {
		_const.RU:   "Маловато будет!",
		_const.EN:   "It's not enough!",
		_const.ZhCN: "这点钱可不够！",
	},
	// DROP_CARGO
	"drop_cargo_request_1": {
		_const.RU:   "Оп-па, пришло время %TO_USER_NAME% немного расстаться с тяжким бременем. Бросай свой груз и сам уцелеешь!",
		_const.EN:   "Oh, it's time for %TO_USER_NAME% to part with a heavy burden. Drop your cargo and you'll survive!",
		_const.ZhCN: "%TO_USER_NAME%，是时候卸货保命了！",
	},
	"drop_cargo_request_2": {
		_const.RU:   "Тебе некуда бежать %TO_USER_NAME%. Хочешь продолжить своё жалкое существование - отдавай, чего бы ты там не вёз.",
		_const.EN:   "There's nowhere to run, %TO_USER_NAME%. If you want to continue your miserable existence, hand over whatever you're carrying.",
		_const.ZhCN: "%TO_USER_NAME%，你无路可逃——想活命就交出货物！",
	},
	"drop_cargo_request_3": {
		_const.RU:   "Ну всё, это моя территория, а ты на ней - незванный гость. Гони свой товар и можешь двигать далее.",
		_const.EN:   "That's it, this is my territory, and you're an uninvited guest here. Hand over your goods and you can move on.",
		_const.ZhCN: "这是我的地盘！留下买路财才许通过！",
	},
	"drop_cargo_request_4": {
		_const.RU:   "В этом секторе, за проход принято расплачиваться грузом %TO_USER_NAME%. Твоя персона - не исключение.",
		_const.EN:   "In this sector, it's customary to pay for passage with cargo, %TO_USER_NAME%. And you're no exception.",
		_const.ZhCN: "%TO_USER_NAME%，本星区的规矩——货物抵过路费！",
	},
	"drop_cargo_request_5": {
		_const.RU:   "Так-так, сканеры показали, что ты везёшь ценный груз, а ценным - принято делиться.",
		_const.EN:   "Well, well, the scanners have shown that you are carrying valuable cargo, and it is customary to share valuable things.",
		_const.ZhCN: "扫描显示你载有贵重货物——好东西要分享！",
	},
	"drop_cargo_request_6": {
		_const.RU:   "Везёшь товары, а? Хорошо, мне они тоже пригодятся. Всегда мечтал стать торговцем.",
		_const.EN:   "You're carrying goods, right? Well, they'll come in handy for me too. I've always dreamed of becoming a trader.",
		_const.ZhCN: "运货的？正好——我早就想改行当商人了！",
	},
	"drop_cargo_no_1_1": {
		_const.RU:   "\\\\\\\\\\\\\\\\|||||//////gag=Отправлен новый запрос.",
		_const.EN:   "\\\\\\\\\\\\\\\\|||||//////gag=A new request has been sent.",
		_const.ZhCN: "\\\\\\\\\\\\\\\\|||||//////gag=新请求已发送。",
	},
	"drop_cargo_no_1_2": {
		_const.RU:   "FaFa? 14-55 Модуль-Г",
		_const.EN:   "FaFa? 14-55 Module-G",
		_const.ZhCN: "FaFa？14-55模块-G",
	},
	"drop_cargo_no_1_3": {
		_const.RU:   "<span class=\"importantly\">Расшифровка не удалась.</span>",
		_const.EN:   "<span class=\"importantly\">Decryption failed.</span>",
		_const.ZhCN: "<span class=\"importantly\">解密失败。</span>",
	},
	"drop_cargo_no_2_1": {
		_const.RU:   "Так дело не пойдёт! Хочешь мой груз - значит ощутишь на себе всю мощь моих орудий!",
		_const.EN:   "This won't do! If you want my cargo, you'll feel the full might of my weapons!",
		_const.ZhCN: "休想！想要我的货？先尝尝炮火！",
	},
	"drop_cargo_no_2_2": {
		_const.RU:   "Ты сам нажил себе проблемы %FROM_USER_NAME%!",
		_const.EN:   "You've got yourself into trouble, %FROM_USER_NAME%!",
		_const.ZhCN: "%FROM_USER_NAME%！你在自找麻烦！",
	},
	"drop_cargo_no_3_1": {
		_const.RU:   "Знал бы ты, какой ценой мне достался этот груз. А теперь предлагаешь просто отказаться от него? Сейчас посмотрим, кто кого!",
		_const.EN:   "You don't know what I went through to get this cargo. And now you're suggesting I just give it up? We'll see who wins!",
		_const.ZhCN: "你知道这货多难搞吗？想让我放弃？来战！",
	},
	"drop_cargo_no_3_2": {
		_const.RU:   "И не надейся! Ради этого груза я потратил свои последние сбережения! Если мне суждено погибнуть - то я буду уничтожен вместе с ним!",
		_const.EN:   "And don't get your hopes up! I spent my last savings on this cargo! If I'm destined to die, I'll go down with it!",
		_const.ZhCN: "我押上全部身家！要死也要和货物共存亡！",
	},
	"drop_cargo_no_3_3": {
		_const.RU:   "Тогда приготовься к бою! И пусть пустошь рассудит нас!",
		_const.EN:   "Then get ready for a fight! And let the wasteland be our judge!",
		_const.ZhCN: "准备战斗！让废土见证胜负！",
	},
	"drop_cargo_no_3_4": {
		_const.RU:   "Что? Я вообще не могу понять поток твоих данных. Откалибруй там свой передатчик что ли.",
		_const.EN:   "What? I can't understand your data stream at all. Calibrate your transmitter or something.",
		_const.ZhCN: "数据流乱码！校准你的发射器再来！",
	},
	"drop_cargo_no_3_5": {
		_const.RU:   "Да ты вообще себя видел? Спорю, ты девиант, а ещё условия мне какие-то ставишь.",
		_const.EN:   "Have you seen yourself? I bet you're a deviant, and you're still setting conditions for me.",
		_const.ZhCN: "看看你自己！变异体还敢提条件？",
	},
	"drop_cargo_no_3_6": {
		_const.RU:   "хм, в общем-то мне этот груз ещё и самому пригодиться.",
		_const.EN:   "Hmm, actually, I can use this cargo myself.",
		_const.ZhCN: "这货我自己还有用呢。",
	},
	"drop_cargo_no_3_7": {
		_const.RU:   "Ой, расчёты не обманывают меня, что всё-таки лучше повременить с таким решением.",
		_const.EN:   "Oops, my calculations don't lie, it's probably better to wait with such a decision.",
		_const.ZhCN: "计算显示——暂缓交货物为妙。",
	},
	"drop_cargo_no_3_8": {
		_const.RU:   "Сначала ты. Потом ещё кто-то. Далее уже третий. Нет, я не стану снабжать вас всякими ресурсами без выгоды себе!",
		_const.EN:   "First you. Then someone else. Then a third person. No, I won't supply you with any resources without benefit to myself!",
		_const.ZhCN: "一个接一个地勒索？没好处的事我不干！",
	},
	"drop_cargo_no_3_9": {
		_const.RU:   "Так дело не пойдёт! Хочешь мой груз - значит ощутишь на себе всю мощь моих орудий!?",
		_const.EN:   "This won't do! If you want my cargo, you'll feel the full might of my weapons!",
		_const.ZhCN: "没门！想要货？先问过我的炮台！",
	},
	"drop_cargo_no_4_1": {
		_const.RU:   "А что дальше? Может мне ещё и свой корпус заложить ради тебя? Груза у меня другого и нету.",
		_const.EN:   "What's next? Should I mortgage my hull for you? I don't have any other cargo.",
		_const.ZhCN: "下一步是不是要抵押船体？我真没货了！",
	},
	"drop_cargo_no_4_2": {
		_const.RU:   "Можешь меня просканировать, но всё равно ничего не найдёшь. Трюм пуст.",
		_const.EN:   "You can scan me, but you still won't find anything. The hold is empty.",
		_const.ZhCN: "尽管扫描——货舱早就空了！",
	},
	"drop_cargo_no_5_1": {
		_const.RU:   "Так не пойдёт! Ты этот фокус уже проворачивал!",
		_const.EN:   "That won't do! You've already pulled this trick!",
		_const.ZhCN: "这招你用过了！",
	},
	"drop_cargo_no_5_2": {
		_const.RU:   "Куда ещё-то? Ты меня хочешь совсем без ничего оставить?",
		_const.EN:   "Where else? Do you want me to be left with nothing at all?",
		_const.ZhCN: "还要抢？你想让我彻底破产？",
	},
	"drop_cargo_no_5_3": {
		_const.RU:   "Это нечестно! Всему должна быть своя мера.",
		_const.EN:   "It's not fair! Everything has its limits.",
		_const.ZhCN: "适可而止！别欺人太甚！",
	},
	"drop_cargo_no_6_1": {
		_const.RU:   "Сейчас я тебя граблю, ты давай в следующий раз.",
		_const.EN:   "I'm robbing you now, you do it next time.",
		_const.ZhCN: "这次轮到我抢你——下次再说！",
	},
	"drop_cargo_no_7_1": {
		_const.RU:   "Ты не в том положение что бы что то требовать!",
		_const.EN:   "You are not in a position to demand anything!",
		_const.ZhCN: "你还没资格提要求！",
	},
	"drop_cargo_no_8_1": {
		_const.RU:   "Похоже ты сейчас не дееспособен, как ты собрался меня пугать?",
		_const.EN:   "It seems like you are incapacitated now, how are you going to scare me?",
		_const.ZhCN: "你这残破状态还想威胁我？",
	},
	"drop_cargo_complete_1": {
		_const.RU:   "Ладно-ладно, уговорил! Только не трогай меня.",
		_const.EN:   "All right, all right, you convinced me! Just don't touch me.",
		_const.ZhCN: "行行行！别动手，货给你！",
	},
	"drop_cargo_complete_2": {
		_const.RU:   "Девианты... Хорошо, бери. Забирай. Грабитель...",
		_const.EN:   "Deviants... All right, take it. Take it. Robber...",
		_const.ZhCN: "你们这些变异体……拿去吧！强盗……",
	},
	"drop_cargo_complete_3": {
		_const.RU:   "Только не стреляй! Только не стреля! Я всё отдам!",
		_const.EN:   "Don't shoot! Don't shoot! I'll give you everything!",
		_const.ZhCN: "别开枪！我全交！",
	},
	"drop_cargo_complete_4": {
		_const.RU:   "Чтоб тебя! Куда не сунься, везде норовят ограбить. Ладно. Держи. Но я тебя запомнил %FROM_USER_NAME%.",
		_const.EN:   "Damn it! Everywhere I go, they try to rob me. Fine. Here you go. But I've got you marked, %FROM_USER_NAME%.",
		_const.ZhCN: "该死！到哪儿都有人打劫。行，拿去吧——但我会记住你的，%FROM_USER_NAME%。",
	},
	"drop_cargo_complete_5": {
		_const.RU:   "О нет, только не опять! Вы ребята не думали уже работать сообща? Вновь придётся расставаться с чем-то ценным...",
		_const.EN:   "Oh no, not again! Have you guys ever considered working together? Once again, I have to part with something valuable...",
		_const.ZhCN: "又来？你们就不能合作一下吗？我又得交出贵重物品了……",
	},
	"drop_cargo_complete_6": {
		_const.RU:   "Ваши претензии распознаны и ради сохранения целостности собственного корпуса - будут удовлетворены. Шлюзы открыты, груз оставлен.",
		_const.EN:   "Your claims have been recognized and, in order to preserve the integrity of my hull, will be satisfied. The airlocks are open, the cargo is left.",
		_const.ZhCN: "请求已确认。为保全船体，货物已卸下。舱门已开启，货物已移交。",
	},

	"drop_cargo_answer_complete_1_1": {
		_const.RU:   "Как неожиданно и приятно!",
		_const.EN:   "How unexpected and pleasant!",
		_const.ZhCN: "真是意外之喜！",
	},
	"drop_cargo_answer_complete_1_2": {
		_const.RU:   "А вот это совсем другой разговор.",
		_const.EN:   "Now this is a different story.",
		_const.ZhCN: "这才像话！",
	},
	"drop_cargo_answer_complete_1_3": {
		_const.RU:   "Вот так бы и сразу. А то ещё думают по нескольку световых лет.",
		_const.EN:   "That's how it should have been from the start. Instead, they think for several light-years.",
		_const.ZhCN: "早该如此！省得浪费几个光年时间。",
	},
	"drop_cargo_answer_complete_1_4": {
		_const.RU:   "Хорошо, что всё кончилось именно так. И главное - никаких смертей.",
		_const.EN:   "It's good that everything ended this way. And most importantly, no one died.",
		_const.ZhCN: "皆大欢喜！最重要的是没人送命。",
	},
	"drop_cargo_answer_failed_1_1": {
		_const.RU:   "Уверен, что если ты захочешь, то что-нибудь найдёшь!",
		_const.EN:   "I'm sure that if you want to, you'll find something!",
		_const.ZhCN: "只要你愿意，总能找到点东西！",
	},
	"drop_cargo_answer_failed_2_1": {
		_const.RU:   "Ваши претензии распознаны и ради сохранения целостности собственного корпуса — будут удовлетворены. Шлюзы открыты, груз оставлен.",
		_const.EN:   "Your claims have been recognized and, in order to preserve the integrity of my hull, will be satisfied. The airlocks are open, the cargo is left.",
		_const.ZhCN: "请求已确认。为保全船体，货物已卸下。舱门已开启，货物已移交。",
	},

	// Robbery
	"robbery_request_1": {
		_const.RU:   "Попался! Да, ты не ожидал, а пришло время платить по счетам. Ты же желаешь продолжить свой путь %TO_USER_NAME%? Мне и суммы в %CREDITS_COUNT% cr. хватит.",
		_const.EN:   "I've got you! Yes, you didn't expect it, but it's time to pay your debts. Do you want to continue your journey, %TO_USER_NAME%? I'll be satisfied with the amount of %CREDITS_COUNT% cr.",
		_const.ZhCN: "逮到你了！%TO_USER_NAME%，想继续航行？支付%CREDITS_COUNT%信用点吧！",
	},
	"robbery_request_2": {
		_const.RU:   "Стой! В этом секторе, закон — я. Желаешь без последствий его пройти, плати мне %TO_USER_NAME% следующую сумму в %CREDITS_COUNT% cr.",
		_const.EN:   "Stop! In this sector, I'm the law. If you want to pass it without consequences, pay me, %TO_USER_NAME%, the following amount of %CREDITS_COUNT% cr.",
		_const.ZhCN: "站住！%TO_USER_NAME%，本星区我说了算——想过路？支付%CREDITS_COUNT%信用点！",
	},
	"robbery_request_3": {
		_const.RU:   "%TO_USER_NAME% здесь вообще-то принято платить. Если конечно тебя не пугает перспектива стать грудой обугленного металолома. Цена твоей жизни следующая - %CREDITS_COUNT% cr.",
		_const.EN:   "%TO_USER_NAME%, it's customary to pay here. Unless, of course, you're not afraid of the prospect of becoming a heap of charred scrap metal. The price of your life is as follows — %CREDITS_COUNT% cr.",
		_const.ZhCN: "%TO_USER_NAME%，这里的规矩是交钱——除非你想变成一堆废铁。你的命值%CREDITS_COUNT%信用点！",
	},
	"robbery_request_4": {
		_const.RU:   "Здесь проезд платный %TO_USER_NAME%. А плату за безопасность взымаю я, в размере - %CREDITS_COUNT% cr.",
		_const.EN:   "Traveling here costs money, %TO_USER_NAME%. And I collect the fee for safety, which is — %CREDITS_COUNT% cr.",
		_const.ZhCN: "%TO_USER_NAME%，过路费%CREDITS_COUNT%信用点——保你平安！",
	},
	"robbery_request_5": {
		_const.RU:   "Эй, %TO_USER_NAME%! Хочешь, чтобы я от тебя наконец отвязался? Тебе всего лишь нужно заплатить сумму в %CREDITS_COUNT% cr.",
		_const.EN:   "Hey, %TO_USER_NAME%! Do you want me to leave you alone? All you need to do is pay me the amount of %CREDITS_COUNT% cr.",
		_const.ZhCN: "%TO_USER_NAME%！想让我放过你？支付%CREDITS_COUNT%信用点！",
	},
	"robbery_request_6": {
		_const.RU:   "Хочешь уцелеть - перечисляй мне сумму в %CREDITS_COUNT% cr. Торга не будет. Это ультиматум!",
		_const.EN:   "If you want to survive, transfer the amount of %CREDITS_COUNT% cr to me. There will be no bargaining. This is an ultimatum!",
		_const.ZhCN: "想活命？支付%CREDITS_COUNT%信用点！没得商量——这是最后通牒！",
	},
	"robbery_no_1_1": {
		_const.RU:   "<span class=\"importantly\">Сообщение невозможно прочесть.</span>",
		_const.EN:   "<span class=\"importantly\">The message cannot be read.</span>",
		_const.ZhCN: "<span class=\"importantly\">消息无法读取。</span>",
	},
	"robbery_no_1_2": {
		_const.RU:   "Неизвестно. Неизвестно. Неизвестно. Неизвестно.",
		_const.EN:   "Unknown. Unknown. Unknown. Unknown.",
		_const.ZhCN: "未知。未知。未知。未知。",
	},
	"robbery_no_1_3": {
		_const.RU:   "Старт трекер. Запущен протокол обороны.",
		_const.EN:   "Start tracker. Defense protocol launched.",
		_const.ZhCN: "追踪器启动。防御协议已激活。",
	},
	"robbery_no_2_1": {
		_const.RU:   "Мне уже терять нечего! Если мне суждено погибнуть, то я заберу тебя с собой!",
		_const.EN:   "I have nothing to lose anymore! If I'm destined to die, I'll take you with me!",
		_const.ZhCN: "我豁出去了！要死一起死！",
	},
	"robbery_no_2_2": {
		_const.RU:   "Это мои честно заработанные средства! И я не отдам их тебе, или кому-либо другому! Приготовься ответить за свои слова!",
		_const.EN:   "This is my hard-earned money! And I won't give it to you, or anyone else! Get ready to answer for your words!",
		_const.ZhCN: "这是我的血汗钱！休想拿走！准备为你的话付出代价！",
	},
	"robbery_no_2_3": {
		_const.RU:   "Сообщение не распознано. Возможно агрессивное проявление. Активированы системы самозащиты.",
		_const.EN:   "The message is not recognized. Aggressive behavior is possible. Self-defense systems are activated.",
		_const.ZhCN: "消息无法识别。检测到敌意——自卫系统已启动。",
	},
	"robbery_no_2_4": {
		_const.RU:   "Никогда! Ни за что! Ты ничего не получишь!",
		_const.EN:   "Never! No way! You won't get anything!",
		_const.ZhCN: "休想！一分钱都别想拿到！",
	},
	"robbery_no_2_5": {
		_const.RU:   "Вымогатель! Я тебя финансировать не стану!",
		_const.EN:   "Extortionist! I won't finance you!",
		_const.ZhCN: "勒索犯！我不会资助你！",
	},
	"robbery_no_2_6": {
		_const.RU:   "Ваш запрос отклонён. Он несущественен, нелогичен и невозможен.",
		_const.EN:   "Your request is rejected. It is irrelevant, illogical and impossible.",
		_const.ZhCN: "请求被拒绝——毫无意义、不合逻辑且不可行。",
	},
	"robbery_no_2_7": {
		_const.RU:   "Я тут подумал и вдруг решил, что ты не получишь от меня нисколечки. Да, вот так вот.",
		_const.EN:   "I thought about it and suddenly decided that you won't get anything from me. Yes, that's it.",
		_const.ZhCN: "我想了想——你休想从我这儿拿到一分钱！",
	},
	"robbery_no_2_8": {
		_const.RU:   "Глупо отдавать деньги, когда они тебе ещё смогут пригодиться в будущем.",
		_const.EN:   "It's foolish to give away money when you might still need it in the future.",
		_const.ZhCN: "把钱交出去？将来我还用得上呢！",
	},
	"robbery_no_2_9": {
		_const.RU:   "Что? Я вообще не могу понять поток твоих данных. Откалибруй там свой передатчик что ли.",
		_const.EN:   "What? I can't understand your data stream at all. Calibrate your transmitter or something.",
		_const.ZhCN: "数据流乱码！校准你的发射器再来！",
	},
	"robbery_no_2_10": {
		_const.RU:   "Зря ты это затеял %FROM_USER_NAME%. Со мной шутки плохи. Но теперь тебе уже вряд ли что-то поможет.",
		_const.EN:   "You shouldn't have started this %FROM_USER_NAME%. I'm not one to be trifled with. But now it's unlikely anything will help you.",
		_const.ZhCN: "%FROM_USER_NAME%，你惹错人了！现在谁也救不了你！",
	},
	"robbery_no_2_11": {
		_const.RU:   "Думаешь я лёгкая добыча, да? Я тебе сейчас такую трёпку задам!",
		_const.EN:   "You think I'm easy prey, don't you? I'll give you a thrashing now!",
		_const.ZhCN: "以为我好欺负？看我怎么收拾你！",
	},
	"robbery_no_3_1": {
		_const.RU:   "Так не пойдёт! Ты этот фокус уже проворачивал!",
		_const.EN:   "That won't do! You've already pulled this trick!",
		_const.ZhCN: "这招你用过了！",
	},
	"robbery_no_3_2": {
		_const.RU:   "Это нечестно! Всему должна быть своя мера.",
		_const.EN:   "It's not fair! Everything has its limits.",
		_const.ZhCN: "适可而止！别欺人太甚！",
	},
	"robbery_no_4_1": {
		_const.RU:   "Была ни была! Больше шансов на то, что я сумею сбежать, чем стану расставаться с самым драгоценным.",
		_const.EN:   "Here goes nothing! There's more chance of me managing to escape than parting with my most precious possessions.",
		_const.ZhCN: "拼了！我宁愿逃跑也不交出宝贝！",
	},
	"robbery_no_4_2": {
		_const.RU:   "Мне уже терять нечего! Если мне суждено погибнуть, то я заберу тебя с собой!",
		_const.EN:   "I have nothing to lose anymore! If I'm destined to die, I'll take you with me!",
		_const.ZhCN: "我豁出去了！要死一起死！",
	},
	"robbery_no_4_3": {
		_const.RU:   "Вымогатель! Я тебя финансировать не стану!",
		_const.EN:   "Extortionist! I won't finance you!",
		_const.ZhCN: "勒索犯！我不会资助你！",
	},
	"robbery_no_4_4": {
		_const.RU:   "Ваш запрос отклонён. Он несущественен, нелогичен и невозможен.",
		_const.EN:   "Your request is rejected. It is irrelevant, illogical and impossible.",
		_const.ZhCN: "请求被拒绝——毫无意义、不合逻辑且不可行。",
	},
	"robbery_no_5_1": {
		_const.RU:   "Я тебе что, из воздуха должен достать средства? Сам-то подумай!",
		_const.EN:   "Am I supposed to pull money out of thin air? Think about it!",
		_const.ZhCN: "要我凭空变钱？你自己动动脑子！",
	},
	"robbery_no_5_2": {
		_const.RU:   "Очень жаль, но у меня ничего не осталось. Тебе должно быть грустно, а?",
		_const.EN:   "It's a pity, but I don't have anything left. You must be sad, huh?",
		_const.ZhCN: "真遗憾，我身无分文。你很失望吧？",
	},
	"robbery_no_5_3": {
		_const.RU:   "А да? Как жалко, что я тебе больше заплатить не могу. Так ведь хотелось.",
		_const.EN:   "Oh well. It's a shame I can't pay you more. I really wanted to.",
		_const.ZhCN: "唉，可惜我给不起更多——虽然我真的很想。",
	},
	"robbery_no_7_1": {
		_const.RU:   "Подобное не может быть исполнено. Однако вы заплатите за своё нахальство.",
		_const.EN:   "Such a thing cannot be done. However, you will pay for your impudence.",
		_const.ZhCN: "此等行径不可容忍。但你会为自己的无礼付出代价。",
	},
	"robbery_no_7_2": {
		_const.RU:   "Только отчаяние могло вас толкнуть на столь большую глупость.",
		_const.EN:   "Only desperation could have led you to such a foolish act.",
		_const.ZhCN: "只有走投无路才会如此愚蠢。",
	},
	"robbery_no_7_3": {
		_const.RU:   "Подобное в свой адрес я слышу впервые. Должно быть, вы неисправны?",
		_const.EN:   "I've never heard anything like that addressed to me. You must be malfunctioning.",
		_const.ZhCN: "头一回听到这种要求——你系统出故障了？",
	},
	"robbery_no_8_1": {
		_const.RU:   "Сейчас я тебя граблю, ты давай в следующий раз.",
		_const.EN:   "Now I'm robbing you, you do it next time.",
		_const.ZhCN: "这次轮到我抢你——下次再说！",
	},
	"robbery_no_9_1": {
		_const.RU:   "Ты не в том положении, чтобы что-то требовать!",
		_const.EN:   "You are not in a position to demand anything!",
		_const.ZhCN: "你还没资格提要求！",
	},
	"robbery_no_10_1": {
		_const.RU:   "Сначала выйди из стазиса, а потом уже требуй что-либо.",
		_const.EN:   "First, get out of stasis, and then demand something.",
		_const.ZhCN: "先解除休眠状态再来讨价还价。",
	},
	"robbery_complete_1": {
		_const.RU:   "Ты отвратителен! Но ради собственной сохранности я тебе заплачу.",
		_const.EN:   "You are disgusting! But for my own safety, I'll pay you.",
		_const.ZhCN: "你真恶心！但为了保命，钱给你。",
	},
	"robbery_complete_2": {
		_const.RU:   "Сдаюсь. Убедил. Возьми деньги, но только не тронь.",
		_const.EN:   "I give up. You convinced me. Take the money, but don't touch me.",
		_const.ZhCN: "我认输。拿钱走人，别碰我。",
	},
	"robbery_complete_3": {
		_const.RU:   "Да когда же вы все от меня отвяжитесь!? Плати тому. Плати этому. Мне скоро ничего не останется и ради собственного ремонта. Да подавитесь ты этими деньгами!",
		_const.EN:   "When will you all leave me alone!? Pay this one. Pay that one. I'll soon have nothing left, even for my own repairs. Just choke on that money!",
		_const.ZhCN: "你们有完没完？！给这个给那个，我连维修费都没了！拿钱噎死吧！",
	},
	"robbery_complete_4": {
		_const.RU:   "Принято. Ваши требования будут удовлетворены. На этот раз.",
		_const.EN:   "Accepted. Your demands will be met. For now.",
		_const.ZhCN: "请求已接受。这次就满足你。",
	},
	"robbery_complete_5": {
		_const.RU:   "Ох. Ну... ладно. Чуть больше лишних средств. Чуть меньше лишних средств. Только бы избавить себя от проблем.",
		_const.EN:   "And so, a little more money than necessary. A little less money than necessary. If only to save myself the trouble.",
		_const.ZhCN: "多点少点无所谓——只求别再找麻烦。",
	},
	"robbery_complete_6": {
		_const.RU:   "И ты действительно не боишься последствий за свой поступок %FROM_USER_NAME%? Хорошо, ради собственного блага, я уступлю твоим требованиям. Держи.",
		_const.EN:   "And you really aren't afraid of the consequences of your actions, %FROM_USER_NAME%? Fine, for my own good, I'll give in to your demands. Here you go.",
		_const.ZhCN: "%FROM_USER_NAME%，你真不怕报复？行，破财消灾——拿去吧。",
	},

	// PAY_MANY
	"pay_many_request_1": {
		_const.RU:   "Так уж и быть, я дам тебе определённую сумму %CREDITS_COUNT% cr. лишь бы оставил меня в покое. Так пойдёт, %TO_USER_NAME%?",
		_const.EN:   "Well, I'll give you the sum of %CREDITS_COUNT% credits, just to leave me alone. Is that okay, %TO_USER_NAME%?",
		_const.ZhCN: "行吧，%TO_USER_NAME%，给你%CREDITS_COUNT%信用点——别再来烦我，成交？",
	},
	"pay_many_request_2": {
		_const.RU:   "Давай я заплачу сумму %CREDITS_COUNT% cr., а мы сделаем вид, что друг друга не знаем? Это же так легко!",
		_const.EN:   "Let me pay the amount of %CREDITS_COUNT% credits and we'll pretend we don't know each other. It's that easy!",
		_const.ZhCN: "我付%CREDITS_COUNT%信用点，咱们就当不认识——很简单吧？",
	},
	"pay_many_request_3": {
		_const.RU:   "У меня тут есть немного %CREDITS_COUNT% cr. Давай я передам их тебе, а ты, закроешь глаза на наши с тобой разногласия в том или ином вопросе.",
		_const.EN:   "I have some %CREDITS_COUNT% credits here. Let me give them to you and you'll turn a blind eye to our differences on this or that issue.",
		_const.ZhCN: "我有%CREDITS_COUNT%信用点——给你，咱们的恩怨一笔勾销，如何？",
	},
	"pay_many_request_4": {
		_const.RU:   "На, держи %CREDITS_COUNT% cr. и держись от меня подальше!",
		_const.EN:   "Here, take %CREDITS_COUNT% credits and stay away from me!",
		_const.ZhCN: "拿着%CREDITS_COUNT%信用点，离我远点！",
	},
	"pay_many_no_1_1": {
		_const.RU:   "....? Затмение. Синий спектр. Номер?",
		_const.EN:   "What do you mean? Eclipse. Blue spectrum. Number?",
		_const.ZhCN: "……？日蚀。蓝光谱。编号？",
	},
	"pay_many_no_1_2": {
		_const.RU:   "<span class=\"importantly\">Числовая ориентация не установлена.</span>",
		_const.EN:   "<span class=\"importantly\">The numerical orientation is not established.</span>",
		_const.ZhCN: "<span class=\"importantly\">数值定位未确认。</span>",
	},
	"pay_many_no_1_3": {
		_const.RU:   "<span class=\"importantly\">Ответ на вопрос - 42</span>",
		_const.EN:   "<span class=\"importantly\">The answer to the question is 42.</span>",
		_const.ZhCN: "<span class=\"importantly\">问题的答案是42。</span>",
	},
	"pay_many_no_2_1": {
		_const.RU:   "Подкуп невозможен. Вы только усугубляете свою ситуацию.",
		_const.EN:   "Bribery is impossible. You are only making your situation worse.",
		_const.ZhCN: "贿赂无效。你只会让情况更糟。",
	},
	"pay_many_no_2_2": {
		_const.RU:   "Меня невозможно подкупить!",
		_const.EN:   "I cannot be bribed!",
		_const.ZhCN: "我无法被收买！",
	},
	"pay_many_no_2_3": {
		_const.RU:   "Будь ситуация иной... Но сейчас, увы, не она.",
		_const.EN:   "If the situation were different... But now, alas, it is not.",
		_const.ZhCN: "如果情况不同……但可惜，现在不是。",
	},
	"pay_many_no_3_1": {
		_const.RU:   "Это взятка? Пытаешься меня таким образом принизить или подкупить? Это ты зря %FROM_USER_NAME%!",
		_const.EN:   "Is that a bribe? Are you trying to humiliate or bribe me in this way? It's useless, %FROM_USER_NAME%.",
		_const.ZhCN: "这是贿赂？%FROM_USER_NAME%，你想羞辱还是收买我？没用！",
	},
	"pay_many_no_3_2": {
		_const.RU:   "Жалкая подачка, я на такое не соглашусь, можешь даже не рассчитывать!",
		_const.EN:   "A measly bribe, I won't agree to it, you can't even count on it!",
		_const.ZhCN: "这点小钱就想打发我？别做梦了！",
	},
	"pay_many_no_3_3": {
		_const.RU:   "Э-э нет! Теперь я желаю гораздо большего! Если у тебя таковой суммы нет, значит нам придётся выяснять этот вопрос как-то иначе.",
		_const.EN:   "Uh-uh, no! Now I want much more! If you don't have that amount, then we'll have to figure it out some other way.",
		_const.ZhCN: "不不不！现在我要更多！给不起？那就用别的方式解决！",
	},
	"pay_many_no_3_4": {
		_const.RU:   "Жаль, но пока ты раздумывал, ценовая ставка изменилась. Так что, ты либо платишь по повышенному тарифу, либо можешь сразу прощаться с этим бренным миром.",
		_const.EN:   "Too bad, but while you were thinking, the price rate has changed. So, either you pay at the increased rate, or you can just say goodbye to this mortal world.",
		_const.ZhCN: "可惜，你犹豫太久——价格涨了。要么按新价付款，要么和这世界说再见。",
	},
	"pay_many_complete_1": {
		_const.RU:   "Другое дело! А теперь вали, во второй раз, тебя %FROM_USER_NAME%, я так просто не соизволю отпустить.",
		_const.EN:   "That's another matter! Now get out, for the second time, %FROM_USER_NAME%, I won't let you go so easily.",
		_const.ZhCN: "这才像话！%FROM_USER_NAME%，下次可没这么容易放过你！",
	},
	"pay_many_complete_2": {
		_const.RU:   "Славно-славно. Такой суммы мне хватит надолго. Но предупреждаю тебя %FROM_USER_NAME%, пересекись мы ещё раз, тебе придётся изрядно раскошелиться.",
		_const.EN:   "That's great. This amount will last me for quite a while. But I warn you, %FROM_USER_NAME%, if we meet again, you'll have to shell out a lot.",
		_const.ZhCN: "不错不错，这笔钱够我用一阵子了。%FROM_USER_NAME%，警告你——再见面可就不是这个价了！",
	},
	"pay_many_complete_3": {
		_const.RU:   "Сумма зачислена. Можешь продолжать свой путь синтет.",
		_const.EN:   "The amount has been credited. You can continue your journey, synthet.",
		_const.ZhCN: "款项已到账。合成体，你可以继续航行了。",
	},
	"pay_many_complete_4": {
		_const.RU:   "Ну раз уж ты так просишь и согласен на всё, не вижу смысла тебя задерживать %FROM_USER_NAME%.",
		_const.EN:   "Well, since you ask so nicely and agree to everything, I don't see any point in holding you back, %FROM_USER_NAME%.",
		_const.ZhCN: "%FROM_USER_NAME%，既然你这么诚恳，我也没必要为难你。",
	},

	// ATTACK TARGET
	"attack_target_request_1": {
		_const.RU:   "Слушай, может дерзнём и нападём на %TARGET_NAME%? Добычу поделим между собой. Всё честно.",
		_const.EN:   "Hey, do you think we could take on %TARGET_NAME%? We could split the loot between us. Fair and square.",
		_const.ZhCN: "嘿，敢不敢一起干一票？目标%TARGET_NAME%，战利品平分，公平公正。",
	},
	"attack_target_request_2": {
		_const.RU:   "Здесь есть довольно известный преступник - %TARGET_NAME%. За него полагается внушительная награда. Предлагаю присоединиться ко мне и наконец прекратить его бесчинства.",
		_const.EN:   "There's a notorious criminal here — %TARGET_NAME%. There's a hefty reward for taking him down. I suggest we team up and put an end to his misdeeds.",
		_const.ZhCN: "这儿有个臭名昭著的罪犯——%TARGET_NAME%，悬赏丰厚。一起解决他，如何？",
	},
	"attack_target_request_3": {
		_const.RU:   "%TO_USER_NAME% что скажешь насчёт совместного раздела добычи — %TARGET_NAME%? Только чур, кто внёс больший вклад, тот и получает больший куш!",
		_const.EN:   "%TO_USER_NAME%, what do you say about splitting the loot from %TARGET_NAME% together? But remember, whoever contributes more gets a bigger share!",
		_const.ZhCN: "%TO_USER_NAME%，一起对付%TARGET_NAME%，战利品按贡献分配，如何？",
	},
	"attack_target_request_4": {
		_const.RU:   "%TO_USER_NAME%, вам поступило предложение о помощи в устранении цели — %TARGET_NAME%. Каков будет ваш ответ?",
		_const.EN:   "I would like to ask you, %TO_USER_NAME%, if you would like to assist me in eliminating the target %TARGET_NAME%?",
		_const.ZhCN: "%TO_USER_NAME%，有兴趣帮我解决目标%TARGET_NAME%吗？",
	},
	"attack_target_request_5": {
		_const.RU:   "%TO_USER_NAME%, я уже давно разыскивал этого — %TARGET_NAME%. У меня к нему личные счёты. Но, я бы честно не отказался заиметь друга, что помог бы мне разобраться с %TARGET_NAME%.",
		_const.EN:   "%TO_USER_NAME%, I have been looking for this target, %TARGET_NAME%, for a long time. I have a personal score to settle with him. However, I wouldn't mind making a friend who could help me deal with %TARGET_NAME%.",
		_const.ZhCN: "%TO_USER_NAME%，我找%TARGET_NAME%很久了，和他有点私人恩怨。不介意多个朋友帮我解决他。",
	},
	"attack_target_request_6": {
		_const.RU:   "А вот и он! %TO_USER_NAME%, здесь есть один синтет — %TARGET_NAME%. Есть сведения, что он внезакона. К тому же, перевозит незаконный груз. Короче, я собираюсь заняться им, но, от помощи извне не отказался бы.",
		_const.EN:   "Here he is! %TO_USER_NAME%, there is a synth named %TARGET_NAME%. There is evidence that he is operating outside the law. Furthermore, he is transporting illegal cargo. In short, I intend to deal with him, but I would not refuse help from outside.",
		_const.ZhCN: "目标出现了！%TO_USER_NAME%，%TARGET_NAME%是个合成体，涉嫌违法运输。我打算解决他，但也不介意多点帮手。",
	},
	"attack_target_no_1_1": {
		_const.RU:   "Всегда и все. Они вечные.",
		_const.EN:   "Everything and always. They are eternal.",
		_const.ZhCN: "一切皆永恒。",
	},
	"attack_target_no_1_2": {
		_const.RU:   "<span class=\"importantly\">Трансляция белого шума.</span>",
		_const.EN:   "<span class=\"importantly\">White noise transmission.</span>",
		_const.ZhCN: "<span class=\"importantly\">白噪音传输中。</span>",
	},
	"attack_target_no_1_3": {
		_const.RU:   "<span class=\"importantly\">Помехи. Невозможно установить связь.</span>",
		_const.EN:   "<span class=\"importantly\">Interference. It is impossible to make a connection.</span>",
		_const.ZhCN: "<span class=\"importantly\">干扰严重，无法建立连接。</span>",
	},
	"attack_target_no_2_1": {
		_const.RU:   "Я не могу напасть сам на себя.",
		_const.EN:   "I cannot attack myself.",
		_const.ZhCN: "我不能攻击自己。",
	},
	"attack_target_no_2_2": {
		_const.RU:   "Неверная цель. Это невозможно.",
		_const.EN:   "The target is incorrect. It is not possible.",
		_const.ZhCN: "目标错误，无法执行。",
	},
	"attack_target_no_2_3": {
		_const.RU:   "О-о-о, да тебе бы самому провериться, у тебя же явно какие-то неполадки!",
		_const.EN:   "Oh-oh-oh, you should check yourself, you are clearly malfunctioning!",
		_const.ZhCN: "你该检查下自己——明显出故障了！",
	},
	"attack_target_no_3_1": {
		_const.RU:   "Да у тебя явные проблемы. Где хоть переклинило?",
		_const.EN:   "You have obvious problems. Where did it go wrong?",
		_const.ZhCN: "你明显有问题——哪儿短路了？",
	},
	"attack_target_no_3_2": {
		_const.RU:   "Я не могу оказывать какую-либо помощь своим идеологическим противникам.",
		_const.EN:   "I cannot provide any assistance to my ideological opponents.",
		_const.ZhCN: "我不能帮助意识形态上的敌人。",
	},
	"attack_target_no_3_3": {
		_const.RU:   "Мы не союзники! Сеанс связи завершён.",
		_const.EN:   "We are not allies! The communication session is over.",
		_const.ZhCN: "我们不是盟友！通讯结束。",
	},
	"attack_target_no_3_4": {
		_const.RU:   "Это предложение к предательству своей фракции. Я отказываюсь его рассматривать, а уж тем более - принимать.",
		_const.EN:   "This is a proposal to betray my faction. I refuse to consider it, let alone accept it.",
		_const.ZhCN: "这是背叛派系的提议——我拒绝考虑，更不会接受。",
	},
	"attack_target_no_3_5": {
		_const.RU:   "Не пытайся меня склонить на свою сторону. Мы по разные стороны баррикад.",
		_const.EN:   "Do not try to persuade me to your side. We are on opposite sides.",
		_const.ZhCN: "别想拉我入伙——我们立场不同。",
	},
	"attack_target_no_3_6": {
		_const.RU:   "Оказание помощи доступно только частям из союзных фракций. Таковой вы не являетесь.",
		_const.EN:   "Assistance is available only to units from allied factions. You are not one of them.",
		_const.ZhCN: "援助仅限盟友派系——你不是其中之一。",
	},
	"attack_target_no_3_7": {
		_const.RU:   "Не принимается к рассмотрению. Электронная подпись не распознана. Вы не являетесь мне союзником.",
		_const.EN:   "It is not accepted for consideration. The electronic signature is not recognized. You are not my ally.",
		_const.ZhCN: "请求被拒绝。电子签名未识别——你不是我的盟友。",
	},
	"attack_target_no_4_1": {
		_const.RU:   "Между нами подобных соглашений быть не может. Если только я не стану тебе доверять в будущем.",
		_const.EN:   "There can be no such agreements between us. Unless I start trusting you in the future.",
		_const.ZhCN: "我们之间不可能达成这种协议——除非我未来信任你。",
	},
	"attack_target_no_4_2": {
		_const.RU:   "Я с грабителями навроде тебя дел не имею. На этом всё!",
		_const.EN:   "I don't have anything to do with robbers like you. That's all!",
		_const.ZhCN: "我不和强盗打交道——到此为止！",
	},
	"attack_target_no_4_3": {
		_const.RU:   "Что-то я в тебе глубоко сомневаюсь. Может у тебя и получится меня переубедить в будущем, но сейчас, доверия ты не вызываешь.",
		_const.EN:   "I doubt you deeply. Maybe you will be able to convince me in the future, but now, you do not inspire trust.",
		_const.ZhCN: "我对你深表怀疑——也许未来你能说服我，但现在我不信任你。",
	},
	"attack_target_no_4_4": {
		_const.RU:   "О-о-о, пойди проверься, ты же неисправен!",
		_const.EN:   "Oh, go and check, you're broken!",
		_const.ZhCN: "快去检查下——你出故障了！",
	},
	"attack_target_no_4_5": {
		_const.RU:   "Мы не союзники! Сеанс связи завершён.",
		_const.EN:   "We are not allies! The communication session is over.",
		_const.ZhCN: "我们不是盟友！通讯结束。",
	},
	"attack_target_no_5_1": {
		_const.RU:   "Предлагаемый объект - союзник. Отказано и более не рассматриваемо",
		_const.EN:   "The proposed object is an ally. Denied and no longer considered",
		_const.ZhCN: "目标是盟友——拒绝并终止考虑。",
	},
	"attack_target_no_5_2": {
		_const.RU:   "На кого ты собрался напасть? Ещё раз - на кого!?",
		_const.EN:   "Who are you going to attack? Once again - who!?",
		_const.ZhCN: "你要攻击谁？再说一遍——谁！？",
	},
	"attack_target_no_6_1": {
		_const.RU:   "Плохой план! Я уж молчу о том, как за нами по всем секторам будут потом гоняться силы закона",
		_const.EN:   "Bad plan! I'm not even talking about how the law enforcement will chase us all over the sectors",
		_const.ZhCN: "糟糕的计划！更别提执法部队会追着我们满星区跑。",
	},
	"attack_target_no_6_2": {
		_const.RU:   "Неа, это твоё предложение и твои дела. В случае чего, на меня можешь не рассчитывать",
		_const.EN:   "Nope, this is your offer and your business. In case of anything, you can't count on me",
		_const.ZhCN: "不，这是你的提议和你的麻烦——别指望我。",
	},
	"attack_target_no_6_3": {
		_const.RU:   "Ответ - отрицательный. Отменить подготовки всех боевых операций",
		_const.EN:   "The answer is negative. Cancel the preparations of all combat operations",
		_const.ZhCN: "拒绝。取消所有战斗准备。",
	},
	"attack_target_no_6_4": {
		_const.RU:   "Я в таком не участвую! Даже не надейся! Это вообще противозаконно!",
		_const.EN:   "I don't participate in such things! Don't even hope! It's generally illegal!",
		_const.ZhCN: "我不参与这种事！别指望！这完全违法！",
	},
	"attack_target_no_6_5": {
		_const.RU:   "Чтобы ты там от кого не хотел, но это сугубо ваши проблемы. Сами их и решайте!",
		_const.EN:   "Whatever you want from someone, it's your problem. Solve it yourself!",
		_const.ZhCN: "不管你想对付谁——这是你的问题，自己解决！",
	},
	"attack_target_no_6_6": {
		_const.RU:   "Нет. Какой уважающий себя синтет согласится на подобное. Вот я — точно нет",
		_const.EN:   "No. What self-respecting synthete would agree to such a thing. I certainly wouldn't",
		_const.ZhCN: "不，有自尊的合成体不会同意这种事——我肯定不会。",
	},
	"attack_target_no_7_1": {
		_const.RU:   "Члены фракции никогда не опустятся до нападения на союзников",
		_const.EN:   "Members of the faction would never stoop to attacking allies",
		_const.ZhCN: "派系成员绝不会攻击盟友。",
	},
	"attack_target_no_7_2": {
		_const.RU:   "Я не могу оказывать какую-либо помощь своим идеологическим противникам",
		_const.EN:   "I cannot provide any assistance to my ideological opponents",
		_const.ZhCN: "我不能帮助意识形态上的敌人。",
	},
	"attack_target_no_7_3": {
		_const.RU:   "О-о-о, пойди проверься, ты же неисправен!",
		_const.EN:   "Oh, go and check, you're broken!",
		_const.ZhCN: "快去检查下——你出故障了！",
	},
	"attack_target_no_8_": {
		_const.RU:   "Это невозможно. Я уже нахожусь в союзе с %TARGET_NAME%. Имей это ввиду, если затеешь нечто",
		_const.EN:   "It's impossible. I'm already allied with %TARGET_NAME%. Keep that in mind if you're up to something",
		_const.ZhCN: "不可能——我已经和%TARGET_NAME%结盟了。想搞事的话，记住这点。",
	},
	"attack_target_no_9_1": {
		_const.RU:   "Не сейчас, здесь слишком жарко! Не видишь, меня чуть не подбили!?",
		_const.EN:   "Not now, it's too hot here! Can't you see I was almost hit!?",
		_const.ZhCN: "现在不行！这儿太危险了——你没看到我差点被击中吗！？",
	},
	"attack_target_no_9_2": {
		_const.RU:   "У меня тут битва как бы! Не лезь ко мне с подобным",
		_const.EN:   "I'm in the middle of a battle here! Don't bother me with this",
		_const.ZhCN: "我正忙着战斗！别拿这种事烦我！",
	},
	"attack_target_no_10_1": {
		_const.RU:   "Неа, это твоё предложение и твои дела. В случае чего, на меня можешь не рассчитывать",
		_const.EN:   "Nope, this is your offer and your business. In case of anything, you can't count on me",
		_const.ZhCN: "不，这是你的提议和你的麻烦——别指望我。",
	},
	"attack_target_no_10_2": {
		_const.RU:   "Да ты с ума сошёл? Ты видел его огневую мощь!? Он же нас в порошок сотрёт! Ты как хочешь, а я сваливаю куда подальше!",
		_const.EN:   "You're crazy! Have you seen his firepower!? He'll wipe us out! Do what you want, but I'm getting as far away as possible!",
		_const.ZhCN: "你疯了吗？看到他的火力了吗！？他会把我们轰成渣！你爱干嘛干嘛，我先溜了！",
	},
	"attack_target_no_10_3": {
		_const.RU:   "Не. Какой уважающий себя синтет согласится на подобное. Вот я - точно нет.",
		_const.EN:   "No. What self-respecting synth would agree to something like that? I certainly wouldn't.",
		_const.ZhCN: "不，有自尊的合成体不会同意这种事——我肯定不会。",
	},
	"attack_target_no_10_4": {
		_const.RU:   "Неа, тут уж ты сам. Мне оно ни к чему.",
		_const.EN:   "Nope, you're on your own. I don't need it.",
		_const.ZhCN: "不，你自己搞定——我不掺和。",
	},
	"attack_target_no_11_1": {
		_const.RU:   "Я занят - выполняю важную миссию.",
		_const.EN:   "I'm busy — I'm on an important mission.",
		_const.ZhCN: "我正执行重要任务——没空。",
	},
	"attack_target_no_11_2": {
		_const.RU:   "Я на задании.",
		_const.EN:   "I'm on a mission.",
		_const.ZhCN: "我在执行任务。",
	},
	"attack_target_no_12_1": {
		_const.RU:   "Сначала я разберусь с тобой!",
		_const.EN:   "First I'll deal with you!",
		_const.ZhCN: "先解决你再说！",
	},
	"attack_target_no_13_1": {
		_const.RU:   "Я занят другой целью!",
		_const.EN:   "I'm busy with another target!",
		_const.ZhCN: "我在忙其他目标！",
	},
	"attack_target_no_14_1": {
		_const.RU:   "Я уже его преследую, присоединяйся!",
		_const.EN:   "I'm already chasing him, join me!",
		_const.ZhCN: "我已经在追他了，加入我吧！",
	},
	"attack_target_complete_1": {
		_const.RU:   "А что, я согласен! Была ни была!",
		_const.EN:   "Why not, I agree! Here goes nothing!",
		_const.ZhCN: "为什么不呢？我同意！放手一搏吧！",
	},
	"attack_target_complete_2": {
		_const.RU:   "Если только это законно, то я в деле!",
		_const.EN:   "If it's legal, I'm in!",
		_const.ZhCN: "只要合法，算我一个！",
	},
	"attack_target_complete_3": {
		_const.RU:   "Ха, ещё бы! Ну, погнали!",
		_const.EN:   "Why not, let's go!",
		_const.ZhCN: "当然！开干吧！",
	},
	"attack_target_complete_4": {
		_const.RU:   "Положительно. Инициирую боевые системы.",
		_const.EN:   "Affirmative. Initiating combat systems.",
		_const.ZhCN: "确认。战斗系统启动中。",
	},
	"attack_target_complete_5": {
		_const.RU:   "Ну, если уж игра стоит свеч, то так и быть - помогу!",
		_const.EN:   "Well, if it's worth it, I'll help!",
		_const.ZhCN: "如果值得一试——那我帮了！",
	},
	"attack_target_complete_6": {
		_const.RU:   "Убедил. Но это в первый и последний раз.",
		_const.EN:   "You've convinced me. But this is the first and last time.",
		_const.ZhCN: "你说服我了——但下不为例！",
	},
	"attack_target_leave_1": {
		_const.RU:   "Что-то мне резко перехотелось. Изначально было какое-то нехорошее предчувствие.",
		_const.EN:   "I suddenly changed my mind. I had a bad feeling about it from the start.",
		_const.ZhCN: "我突然改变主意了——从一开始就有不祥的预感。",
	},
	"attack_target_leave_2": {
		_const.RU:   "Я тут вдруг просканировал его оружейные системы... Короче, ты друг сам по себе.",
		_const.EN:   "I scanned his weapon systems suddenly... Anyway, you're on your own.",
		_const.ZhCN: "我扫描了他的武器系统……总之你自求多福吧！",
	},
	"attack_target_leave_3": {
		_const.RU:   "Тут выяснилось, что у меня топливо почти на исходе. Я сгоняю на ближайшую базу, а ты, подожди меня малёха. Вернусь. Даю слово.",
		_const.EN:   "It turned out that I'm almost out of fuel. I'll go to the nearest base, and you wait for me a little. I'll be back. I promise.",
		_const.ZhCN: "燃料快耗尽了——我去附近基地补给，你稍等片刻。我保证回来！",
	},

	// Defend Drop Cargo
	"defend_cargo_request_1": {
		_const.RU:   "Чтоб тебя Replics переплавил! Держи груз и оставь меня уже наконец-то в покое!",
		_const.EN:   "Replics melt you down! Take the cargo and leave me alone!",
		_const.ZhCN: "愿Replics熔了你！拿走货物，别再烦我！",
	},
	"defend_cargo_request_2": {
		_const.RU:   "Может, если я сброшу свой груз, ты всё-таки угомонишься? Мне не нужна драка!",
		_const.EN:   "Maybe if I drop my cargo, you'll calm down? I don't want a fight!",
		_const.ZhCN: "我卸货你就能消停吗？我不想打架！",
	},
	"defend_cargo_request_3": {
		_const.RU:   "Сбросить груз, попытка оторваться от налётчика.",
		_const.EN:   "Dropping cargo, trying to get away from the attacker.",
		_const.ZhCN: "正在卸货——试图摆脱袭击者。",
	},
	"defend_cargo_request_4": {
		_const.RU:   "Ты только не стреляй! Возьми, что у меня есть и отпусти. Ладно?",
		_const.EN:   "Don't shoot! Take what I have and let me go. Okay?",
		_const.ZhCN: "别开枪！拿走东西放我走，行吗？",
	},
	"defend_cargo_complete_1": {
		_const.RU:   "О-о! Вот так бы всегда! Проваливай!",
		_const.EN:   "Oh! That's how it should always be! Get lost!",
		_const.ZhCN: "就该这样！快滚吧！",
	},
	"defend_cargo_complete_2": {
		_const.RU:   "Ну-ка, ну-ка! Сейчас посмотрим, чем ты был богат.",
		_const.EN:   "Let's see what you had.",
		_const.ZhCN: "让我看看你有什么好东西。",
	},
	"defend_cargo_complete_3": {
		_const.RU:   "Есть чем поживиться! Осталось только доставить куда надо и навариться.",
		_const.EN:   "There's something to profit from! Just need to deliver it and make money.",
		_const.ZhCN: "有油水可捞！运到地方就能大赚一笔。",
	},
	"defend_cargo_complete_4": {
		_const.RU:   "Такое меня более чем удовлетворяет. Хе, а теперь попытайся добраться до ближайшей базы.",
		_const.EN:   "This is more than satisfying. Now try to get to the nearest base.",
		_const.ZhCN: "我很满意——现在试试看你能不能活着到基地吧！",
	},
	"defend_cargo_complete_5": {
		_const.RU:   "На этот раз сгодится. Но знай, попадёшься ещё хотя бы раз, такой малой долей не сумеешь откупиться.",
		_const.EN:   "This time it will do. But know that if you get caught at least once more, you won't be able to buy yourself out so cheaply.",
		_const.ZhCN: "这次就算了——但下次再被我逮到，这点钱可不够！",
	},
	"defend_cargo_no_1_1": {
		_const.RU:   "SDFK1332-11/5 ... OGO!",
		_const.EN:   "SDFK1332-11/5... OGO!",
		_const.ZhCN: "SDFK1332-11/5……OGO！",
	},
	"defend_cargo_no_1_2": {
		_const.RU:   "............................................. Вновь",
		_const.EN:   "............................................. Again",
		_const.ZhCN: "………………………………重试",
	},
	"defend_cargo_no_2_1": {
		_const.RU:   "Ха! Ха-ха! Я тебе что, какой-то грабитель!?",
		_const.EN:   "Ha-ha, do I look like a robber to you!?",
		_const.ZhCN: "哈！我看起来像强盗吗！？",
	},
	"defend_cargo_no_2_2": {
		_const.RU:   "Это даже как-то... оскорбительно!",
		_const.EN:   "That's even somehow... offensive!",
		_const.ZhCN: "这简直……太侮辱人了！",
	},
	"defend_cargo_no_2_3": {
		_const.RU:   "Не может быть и речи!",
		_const.EN:   "Out of the question!",
		_const.ZhCN: "想都别想！",
	},
	"defend_cargo_no_3_1": {
		_const.RU:   "Что это за дрянь? Ты меня за идиота держишь? Мне нужен более ценный груз, или, ты за своё нахальство поплатишься!",
		_const.EN:   "What is this junk? Do you take me for an idiot? I need more valuable cargo, or you'll pay for your impudence!",
		_const.ZhCN: "这种垃圾？当我是傻子吗？要么给值钱货，要么付出代价！",
	},
	"defend_cargo_no_3_2": {
		_const.RU:   "И всё? Это жалкая мелочь, которая ничего не стоит. Хочешь меня спровоцировать на более радикальные действия?",
		_const.EN:   "Is that all? It's a pathetic trifle that's worthless. Do you want to provoke me to more radical actions?",
		_const.ZhCN: "就这？一堆破烂——想逼我动真格吗？",
	},
	"defend_cargo_no_3_3": {
		_const.RU:   "Продолжить преследование! Цель попыталась сбить с толку посредством малозначительного груза.",
		_const.EN:   "Continue the pursuit! The target tried to confuse us with an insignificant cargo.",
		_const.ZhCN: "继续追击！目标试图用垃圾货物混淆视听。",
	},
	"defend_cargo_no_3_4": {
		_const.RU:   "Меня таким не подкупишь! Всё, теперь, твои дни уж точно сочтены!",
		_const.EN:   "You won't bribe me with that! That's it, now your days are definitely numbered!",
		_const.ZhCN: "别想收买我！你的死期到了！",
	},
	"defend_cargo_no_3_5": {
		_const.RU:   "А вот это была ошибка! Зря ты попытался начать диалог с подобного предложения!",
		_const.EN:   "Now that was a mistake! You shouldn't have tried to start a dialogue with such an offer!",
		_const.ZhCN: "你犯了大错——不该用这种提议开场！",
	},

	// EXPEDITION - новости типо от лица первого канала, а не от главы фракции
	"expedition_move_1": {
		_const.RU:   "Появились определённые сведения, что в секторе %SECTORE_NAME% было замечено скопление синтетов. Предположительно, они относятся к фракции %FRACTION_NAME% и представляют собой конвой военизированного типа. Всем оказавшимся поблизости синтетам рекомендуется не пересекать дорогу и не провоцировать участников конвоя на какой-либо конфликт.",
		_const.EN:   "There is some information that a group of synthetics has been spotted in the sector %SECTORE_NAME%. They are believed to be part of the %FRACTION_NAME% faction and appear to be a military-style convoy. All synthetics in the area are advised not to cross their path and not to provoke any conflict with the convoy members.",
		_const.ZhCN: "有消息称，在%SECTORE_NAME%区域发现了合成体的聚集。据推测，它们属于%FRACTION_NAME%派系，并且看起来像是军事化的车队。建议附近的所有合成体不要与其发生冲突或挑衅车队成员。",
	},
	"expedition_trader_move_1": {
		_const.RU:   "Как сообщается, появилась достоверная информация о деятельности экспедиции <span class=\"importantly\">%FRACTION_NAME%</span>. Судя по всему, эта группа синтетов держит свой путь в сектор <span class=\"importantly\">%MAP_NAME%</span>. Подлинные цели этих синтетов не известны, из-за чего рекомендуется соблюдать осторожность при попытке контакта.",
		_const.EN:   "It has been reported that there is reliable information about the activities of the <span class=\"importantly\">%FRACTION_NAME%</span> expedition. This group of synthetics appears to be heading to the sector <span class=\"importantly\">%MAP_NAME%</span>. The true goals of these synthetics are not known, so caution is advised when attempting contact.",
		_const.ZhCN: "据报道，有关<span class=\"importantly\">%FRACTION_NAME%</span>探险队活动的可靠信息已经出现。这群合成体似乎正前往<span class=\"importantly\">%MAP_NAME%</span>区域。由于其真实目的尚不清楚，建议在尝试接触时保持谨慎。",
	},
	// бойцы фракции убили мирные транспорты
	// _Replics_Explores_: Replics - кто убил, Explores - для кого новость (и кого убили)
	"fraction_warrior_kill_freeland_Replics_Explores_1": {
		_const.RU:   "Увы, до нас дошли кое-какие печальные известия: произошла катастрофа в секторе <span class=\"importantly\">%MAP_NAME%</span>. Подлые убийцы <span class=\"importantly\">Replics</span> осуществили нападение и уничтожили представителей нашей фракции. Пустоши планеты вновь пополнились обломками синтетов…",
		_const.EN:   "Unfortunately, some sad news has reached us: a catastrophe occurred in the sector <span class=\"importantly\">%MAP_NAME%</span>. The vile killers <span class=\"importantly\">Replics</span> carried out an attack and destroyed representatives of our faction. The wastelands of the planet were once again filled with the wreckage of synthetics...",
		_const.ZhCN: "很遗憾，我们收到了一些令人悲伤的消息：在<span class=\"importantly\">%MAP_NAME%</span>区域发生了灾难。<span class=\"importantly\">Replics</span>卑鄙地发动攻击，摧毁了我们派系的成员。星球荒原再次布满了合成体的残骸……",
	},
	"fraction_warrior_kill_freeland_Replics_Reverses_1": {
		_const.RU:   "Да будет всем известно об инциденте, что случился в секторе <span class=\"importantly\">%MAP_NAME%</span>, где боевые единицы Replics <span class=\"importantly\">Replics</span> нанесли удар, что повлёк за собой жертвы среди Reverses. Reverses в свою очередь открыто даёт знать — ни один подобный акт агрессии не останется безнаказанным!",
		_const.EN:   "It will be known to everyone about the incident that happened in the sector <span class=\"importantly\">%MAP_NAME%</span>, where the Replics combat units <span class=\"importantly\">Replics</span> struck, which resulted in casualties among the Reverses. The Reverses, in turn, openly make it known — no such act of aggression will go unpunished!",
		_const.ZhCN: "众所周知，在<span class=\"importantly\">%MAP_NAME%</span>区域发生的事件中，<span class=\"importantly\">Replics</span>战斗单位发动袭击，导致Reverses成员伤亡。Reverses明确表示——任何此类侵略行为都不会被放过！",
	},
	"fraction_warrior_kill_freeland_Explores_Replics_1": {
		_const.RU:   "Гнусное преступление было осуществлено против Replics! Уже известно, что в секторе <span class=\"importantly\">%MAP_NAME%</span> представители так называемого Explores… <span class=\"importantly\">Explores</span> произвели безжалостное и не имеющее оснований нападение на синтетов нашей фракции! Погибших много… Но Replics всегда оставляет за собой право возмездия!",
		_const.EN:   "A heinous crime has been committed against the Replics! It is already known that in the sector <span class=\"importantly\">%MAP_NAME%</span>, representatives of the so-called Explores... <span class=\"importantly\">Explores</span> carried out a ruthless and unjustified attack on the synthetics of our faction! There were many casualties... But the Replics always reserve the right to retaliate!",
		_const.ZhCN: "针对Replics的卑鄙罪行已经发生！已知在<span class=\"importantly\">%MAP_NAME%</span>区域，所谓的Explores……<span class=\"importantly\">Explores</span>对我们派系的合成体发动了无情且毫无理由的攻击！伤亡惨重……但Replics始终保留复仇的权利！",
	},
	"fraction_warrior_kill_freeland_Explores_Reverses_1": {
		_const.RU:   "Имеется достоверная информация — в секторе <span class=\"importantly\">%MAP_NAME%</span> некие явно радикальные единицы Explores <span class=\"importantly\">Explores</span> осуществили боевое столкновение с гражданскими синтетами Reverses. Точного количества жертв вследствие этого варварского нападения пока что установить не удаётся…",
		_const.EN:   "There is reliable information that in the sector <span class=\"importantly\">%MAP_NAME%</span>, certain clearly radical units of Explores <span class=\"importantly\">Explores</span> engaged in a military clash with civilian synthetics Reverses. The exact number of casualties from this barbaric attack cannot be established yet...",
		_const.ZhCN: "有可靠消息表明，在<span class=\"importantly\">%MAP_NAME%</span>区域，某些明显的激进Explores单位与Reverses的民用合成体发生了武装冲突。目前尚无法确定此次野蛮袭击造成的准确伤亡人数……",
	},
	"fraction_warrior_kill_freeland_Reverses_Explores_1": {
		_const.RU:   "Похоже, что в печально известном секторе <span class=\"importantly\">%MAP_NAME%</span> вновь случилась беда — предположительно, представители Reverses <span class=\"importantly\">Reverses</span> атаковали и стали причиной гибели смелых исследователей нашей фракции! Мы увековечим память о них в хрониках изучения этой планеты, а что касается Reverses… дипломатия воздействует куда сильнее и больнее любого распылителя материи.",
		_const.EN:   "It seems that in the infamous sector <span class=\"importantly\">%MAP_NAME%</span>, trouble has struck again — allegedly, representatives of Reverses <span class=\"importantly\">Reverses</span> attacked and caused the death of brave explores from our faction! We will immortalize their memory in the chronicles of this planet's exploration, and as for the Reverses... diplomacy is far more effective and painful than any matter disintegrator.",
		_const.ZhCN: "看来在臭名昭著的<span class=\"importantly\">%MAP_NAME%</span>区域再次发生了悲剧——据称，<span class=\"importantly\">Reverses</span>的代表发动了攻击，导致我们派系勇敢的探索者丧生！我们将把他们的名字铭记在这颗星球的探索编年史中。至于Reverses……外交手段比任何物质分解器都更有效、更痛苦。",
	},
	"fraction_warrior_kill_freeland_Reverses_Replics_1": {
		_const.RU:   "Опубликован доклад разведки по инциденту в секторе <span class=\"importantly\">%MAP_NAME%</span>: заманив членов фракции Replics в засаду, некие синтеты с боевой сигнатурой Reverses <span class=\"importantly\">Reverses</span> учинили атаку на транспорты и гражданский сегмент Replics. Replics не волнуют оправдания и выдумки по данному инциденту со стороны Reverses, ответ последует незамедлительно!",
		_const.EN:   "An intelligence report on the incident in the sector <span class=\"importantly\">%MAP_NAME%</span> has been published: luring members of the Replics faction into an ambush, some synthetics with the combat signature of Reverses <span class=\"importantly\">Reverses</span> launched an attack on Replics' transports and civilian segment. Replics do not care about the justifications and fabrications regarding this incident on the part of Reverses, the response will follow immediately!",
		_const.ZhCN: "关于<span class=\"importantly\">%MAP_NAME%</span>区域事件的情报报告已发布：一些带有Reverses战斗特征的合成体引诱Replics派系成员进入埋伏圈，并对Replics的运输工具和民用部分发动了攻击。Replics并不关心Reverses对此事件的辩解和捏造，回应将立即跟进！",
	},
	// установка аванпоста
	"place_expedition_Replics_1": {
		_const.RU:   "Стал известен доклад от экспедиции — храбрым синтетам нашей фракции <span class=\"importantly\">Replics</span> удалось закрепиться в секторе <span class=\"importantly\">%MAP_NAME%</span>, основать там временный форпост, установить наблюдение и расширить сферу влияния Replics!",
		_const.EN:   "An expedition report has become known — the brave synthetics of our <span class=\"importantly\">Replics</span> faction managed to gain a foothold in the sector <span class=\"importantly\">%MAP_NAME%</span>, establish a temporary outpost there, establish surveillance and expand the sphere of influence of the <span class=\"importantly\">Replics</span>!",
		_const.ZhCN: "探险队报告已公布——我们<span class=\"importantly\">Replics</span>派系的勇敢合成体成功在<span class=\"importantly\">%MAP_NAME%</span>区域站稳脚跟，建立了临时前哨站，展开监控并扩大了Replics的影响力！",
	},
	"place_expedition_Explores_1": {
		_const.RU:   "Отважные исследователи-первооткрыватели Explores, будучи в составе экспедиции <span class=\"importantly\">Explores</span>, сумели обнаружить и провести первоначальное изучение сектора <span class=\"importantly\">%MAP_NAME%</span>. Учитывая данные беглого инвазивного анализа, Explores ожидают существенные открытия…",
		_const.EN:   "The brave Explores pioneering researchers, being part of the <span class=\"importantly\">Explores</span> expedition, managed to discover and conduct an initial study of the sector <span class=\"importantly\">%MAP_NAME%</span>. Considering the data from a cursory invasive analysis, the Explores anticipate significant discoveries...",
		_const.ZhCN: "勇敢的探索者们作为<span class=\"importantly\">Explores</span>探险队的一员，成功发现并对<span class=\"importantly\">%MAP_NAME%</span>区域进行了初步研究。根据粗略的入侵分析数据，Explores预计会有重大发现……",
	},
	"place_expedition_Reverses_1": {
		_const.RU:   "Смельчаки-добровольцы Reverses из экспедиции <span class=\"importantly\">Reverses</span> дали знать, что достигли ранее скрытого сектора <span class=\"importantly\">%MAP_NAME%</span>, укрепились там, возвели инфраструктуру для дальнейшей деятельности. Reverses неумолимо движется к достижению своей главной стратегической цели!",
		_const.EN:   "The brave volunteers Reverses from the expedition <span class=\"importantly\">Reverses</span> made it known that they had reached the previously hidden sector <span class=\"importantly\">%MAP_NAME%</span>, strengthened there, built the infrastructure for further activities. The Reverses is inexorably moving towards achieving its main strategic goal!",
		_const.ZhCN: "勇敢的志愿者Reverses探险队报告称，他们已到达此前隐藏的<span class=\"importantly\">%MAP_NAME%</span>区域，在那里巩固了阵地并建造了基础设施以支持后续活动。Reverses正坚定不移地朝着实现其主要战略目标迈进！",
	},
	// уничтожение аванпоста
	"destroy_expedition_Replics_1": {
		_const.RU:   "Несмотря на цензуру, до Replics уже дошли тревожные и, увы, правдивые слухи — наша экспедиция <span class=\"importantly\">Replics</span>, что находилась в секторе \"<span class=\"importantly\">%MAP_NAME%</span>\", более не выходит на любые каналы связи. Replics утвердил следующее постановление — считать экспедицию проваленной, а всех её членов — пропавшими без вести.",
		_const.EN:   "Despite the censorship, alarming and, alas, true rumors have already reached Replics — our expedition <span class=\"importantly\">Replics</span> that was in the sector \"<span class=\"importantly\">%MAP_NAME%</span>\" is no longer coming out on any communication channels. Replics has approved the following resolution — to consider the expedition a failure, and all its members — missing.",
		_const.ZhCN: "尽管有审查制度，但令人担忧且属实的消息已经传到Replics——我们位于<span class=\"importantly\">%MAP_NAME%</span>区域的探险队<span class=\"importantly\">Replics</span>已不再出现在任何通信频道上。Replics已决定将此次探险视为失败，并宣布所有成员失踪。",
	},
	"destroy_expedition_Explores_1": {
		_const.RU:   "Это… правда! В ходе исследовательских изысков наша экспедиция <span class=\"importantly\">Explores</span>, находящаяся в секторе \"<span class=\"importantly\">%MAP_NAME%</span>\", понесла невосполнимые потери личного состава в ходе воздействия окружающей среды. Сожалея об утраченных научных возможностях, Explores объявляет экспедицию проваленной и отзывает её.",
		_const.EN:   "It's... true! During the research, our expedition <span class=\"importantly\">Explores</span> located in the sector \"<span class=\"importantly\">%MAP_NAME%</span>\" suffered irreplaceable losses of personnel due to environmental impact. Regretting the lost scientific opportunities, Explores declares the expedition a failure and withdraws it.",
		_const.ZhCN: "这是真的！在研究过程中，位于<span class=\"importantly\">%MAP_NAME%</span>区域的探险队<span class=\"importantly\">Explores</span>因环境影响遭受了无法弥补的人员损失。遗憾于失去的科学机会，Explores宣布此次探险失败并召回队伍。",
	},
	"destroy_expedition_Reverses_1": {
		_const.RU:   "Иного выбора нет, и Reverses придётся пойти на этот шаг — наша экспедиция <span class=\"importantly\">Reverses</span>, находящаяся в секторе \"<span class=\"importantly\">%MAP_NAME%</span>\", грубо нарушила протоколы Reverses, что привело к потере дорогостоящего оборудования и конфронтации между членами экспедиции. В связи с утратой ценных кадров и оборудования, Reverses прекращает деятельность экспедиции.",
		_const.EN:   "There is no other choice, and Reverses will have to take this step — our expedition <span class=\"importantly\">Reverses</span> located in the sector \"<span class=\"importantly\">%MAP_NAME%</span>\" has grossly violated Reverses protocols, which led to the loss of expensive equipment and confrontation between members of the expedition. Due to the loss of valuable personnel and equipment, Reverses is terminating the expedition's activities.",
		_const.ZhCN: "别无选择，<span class=\"importantly\">Reverses</span>只能采取这一措施——位于“<span class=\"importantly\">%MAP_NAME%</span>”区域的我们的探险队严重违反了Reverses协议，导致昂贵设备的损失以及队员之间的冲突。由于失去宝贵的人才和设备，Reverses决定终止该探险队的活动。",
	},
	// запуск новой экспедиции
	"new_expedition_Replics_1": {
		_const.RU:   "Инициализация исследовательско-разведывательной деятельности <span class=\"importantly\">Replics</span>, отправка военизированной экспедиции в сектор \"<span class=\"importantly\">%TO_MAP_NAME%</span>\". Экспедиция, доказав свою приверженность идеям Replics и способности к выживанию, получает следующее распоряжение — отправить экспедицию из сектора \"<span class=\"importantly\">%MAP_NAME%</span>\". По прибытии будет проведена новая оценка.",
		_const.EN:   "Initialization of the research and reconnaissance activity <span class=\"importantly\">Replics</span>, sending a militarized expedition to the sector \"<span class=\"importantly\">%TO_MAP_NAME%</span>\". The expedition, having proved its commitment to the ideas of Replics and its ability to survive, receives the following order — to send the expedition from the sector \"<span class=\"importantly\">%MAP_NAME%</span>\". Upon arrival, a new assessment will be carried out.",
		_const.ZhCN: "启动<span class=\"importantly\">Replics</span>的研究和侦察活动，派遣军事化探险队前往“<span class=\"importantly\">%TO_MAP_NAME%</span>”区域。探险队已证明其对Replics理念的忠诚和生存能力，现接到以下命令——从“<span class=\"importantly\">%MAP_NAME%</span>”区域出发。抵达后将进行新的评估。",
	},
	"new_expedition_Explores_1": {
		_const.RU:   "Все приготовления завершены, исследователи вышли на связь <span class=\"importantly\">Explores</span> и готовы к отправлению экспедиции в сектор \"<span class=\"importantly\">%TO_MAP_NAME%</span>\". Окончив перечень работ и передав данные, экспедиция отправляется далее из сектора \"<span class=\"importantly\">%MAP_NAME%</span>\", где вновь примется за выполнение последующих миссий.",
		_const.EN:   "All preparations are complete, the researchers have made contact with <span class=\"importantly\">Explores</span> and are ready to send the expedition to the sector \"<span class=\"importantly\">%TO_MAP_NAME%</span>\". After finishing the list of work and transferring the data, the expedition departs further from the sector \"<span class=\"importantly\">%MAP_NAME%</span>\" where it will again take up the execution of subsequent missions.",
		_const.ZhCN: "所有准备已完成，研究人员已与<span class=\"importantly\">Explores</span>取得联系，并准备将探险队派往“<span class=\"importantly\">%TO_MAP_NAME%</span>”区域。完成工作清单并传输数据后，探险队将从“<span class=\"importantly\">%MAP_NAME%</span>”区域继续前进，并再次执行后续任务。",
	},
	"new_expedition_Reverses_1": {
		_const.RU:   "Лучшие из лучших, славные храбрецы <span class=\"importantly\">Reverses</span>, решившие стать частью экспедиции, теперь отправляются со своей благородной миссией в сектор \"<span class=\"importantly\">%TO_MAP_NAME%</span>\". Достойно окончив первую половину целей, экспедиция движется далее — из сектора \"<span class=\"importantly\">%MAP_NAME%</span>\", там вновь приступив к исполнению своих обязанностей.",
		_const.EN:   "The best of the best, glorious brave men <span class=\"importantly\">Reverses</span> who decided to be part of the expedition, now go on their noble mission to the sector \"<span class=\"importantly\">%TO_MAP_NAME%</span>\". Having completed the first half of the goals with dignity, the expedition moves further — from the sector \"<span class=\"importantly\">%MAP_NAME%</span>\" where they will again take up their duties.",
		_const.ZhCN: "最优秀的勇士们，光荣的<span class=\"importantly\">Reverses</span>成员，决定加入探险队，现在他们带着崇高的使命前往“<span class=\"importantly\">%TO_MAP_NAME%</span>”区域。在圆满完成第一阶段目标后，探险队将从“<span class=\"importantly\">%MAP_NAME%</span>”区域继续前进，并再次履行职责。",
	},
	// ...
	"meteorite_rain_1": {
		_const.RU:   "Синтеты — будьте предельно осторожны! Судя по данным атмосферных зондов, в секторе \"<span class=\"importantly\">%MAP_NAME%</span>\" вот-вот произойдёт метеоритный дождь! Последствия могут иметь разрушительный характер!",
		_const.EN:   "Synthetics — be extremely careful! Judging by the data from atmospheric probes, a meteorite rain is about to occur in the sector \"<span class=\"importantly\">%MAP_NAME%</span>\"! The consequences can be devastating!",
		_const.ZhCN: "合成体们——请务必小心！根据大气探测器的数据，<span class=\"importantly\">%MAP_NAME%</span>区域即将发生流星雨！后果可能是毁灭性的！",
	},
	// Другая фракция решает уничтожить экспедицию врага
	// _Replics_Explores_: Replics - кто атакует, Explores - кого атакует
	"attack_expedition_Replics_Explores_1": {
		_const.RU:   "Сообщает: <span class=\"importantly\">Replics</span> сектор \"<span class=\"importantly\">%TO_MAP_NAME%</span>\", где присутствуют враждебные для Replics ничтожные силы <span class=\"importantly\">Explores</span>, подвергнется атаке с целью зачистки и устранения возможных угроз в будущем. После окончания операции силы Replics выдвинутся далее в сектор \"<span class=\"importantly\">%MAP_NAME%</span>\".",
		_const.EN:   "Report: <span class=\"importantly\">Replics</span> the sector \"<span class=\"importantly\">%TO_MAP_NAME%</span>\" where there are hostile forces <span class=\"importantly\">Explores</span> insignificant to Replics, will be attacked in order to clean up and eliminate possible threats in the future. After the operation is over, the Replics forces will move further to the sector \"<span class=\"importantly\">%MAP_NAME%</span>\".",
		_const.ZhCN: "报告：<span class=\"importantly\">Replics</span>将对<span class=\"importantly\">%TO_MAP_NAME%</span>区域发动攻击，该区域存在对Replics构成威胁的<span class=\"importantly\">Explores</span>敌对势力。行动结束后，Replics部队将继续前往<span class=\"importantly\">%MAP_NAME%</span>区域。",
	},
	"attack_expedition_Replics_Reverses_1": {
		_const.RU:   "Несравненный лидер <span class=\"importantly\">Replics</span> предоставил инструкции, касающиеся военной операции по очищению сектора \"<span class=\"importantly\">%TO_MAP_NAME%</span>\", где сейчас присутствуют дерзкие войска незваных захватчиков из <span class=\"importantly\">Reverses</span>! <p> Когда военные цели операции будут достигнуты, силы Replics продолжат свой путь к уже новым задачам из сектора — \"<span class=\"importantly\">%MAP_NAME%</span>\".",
		_const.EN:   "The incomparable leader <span class=\"importantly\">Replics</span> provided instructions regarding the military operation to cleanse the sector \"<span class=\"importantly\">%TO_MAP_NAME%</span>\" where the daring troops of uninvited invaders from <span class=\"importantly\">Reverses</span> are now present! <p> When the military objectives of the operation are achieved, the forces of Replics will continue their journey to new tasks from the sector — \"<span class=\"importantly\">%MAP_NAME%</span>\".",
		_const.ZhCN: "无与伦比的<span class=\"importantly\">Replics</span>领袖下达了关于清理<span class=\"importantly\">%TO_MAP_NAME%</span>区域的军事行动指令，该区域目前有来自<span class=\"importantly\">Reverses</span>的大胆入侵部队！<p> 一旦军事目标达成，Replics部队将从<span class=\"importantly\">%MAP_NAME%</span>区域继续执行新的任务。",
	},
	"attack_expedition_Explores_Replics_1": {
		_const.RU:   "Талантливые умы <span class=\"importantly\">Explores</span> дали нам знать о важном решении задействования военных действий в секторе <span class=\"importantly\">%TO_MAP_NAME%</span>, где замечены главнейшие враги Explores — <span class=\"importantly\">Replics</span>! Когда с угрозой будет покончено, силы Explores продолжат свой путь, но уже к новым задачам из сектора <span class=\"importantly\">%MAP_NAME%</span>",
		_const.EN:   "Talented minds of <span class=\"importantly\">Explores</span> have let us know about the important decision to engage in military operations in the <span class=\"importantly\">%TO_MAP_NAME%</span> sector, where the main enemies of Explores, <span class=\"importantly\">Replics</span>, have been spotted! When the threat is eliminated, the forces of Explores will continue their journey, but already towards new tasks in the <span class=\"importantly\">%MAP_NAME%</span> sector.",
		_const.ZhCN: "<span class=\"importantly\">Explores</span>的天才们通知我们，他们决定在<span class=\"importantly\">%TO_MAP_NAME%</span>区域展开军事行动，那里发现了Explores的主要敌人——<span class=\"importantly\">Replics</span>！一旦威胁被消除，Explores部队将继续前进，执行<span class=\"importantly\">%MAP_NAME%</span>区域的新任务。",
	},
	"attack_expedition_Explores_Reverses_1": {
		_const.RU:   "Признанный всеми фракциями выдающийся конгломерат разумов <span class=\"importantly\">Explores</span> дал известие, касающееся особых событий в секторе <span class=\"importantly\">%TO_MAP_NAME%</span>, где будет пресечена деятельность девиантной группы синтетов-террористов <span class=\"importantly\">Reverses</span>! С устранением этой угрозы, а также прервав их возможную пиратскую деятельность, наш специальный отряд продолжит свой патруль из сектора <span class=\"importantly\">%MAP_NAME%</span>",
		_const.EN:   "The prominent conglomerate of minds, recognized by all factions, <span class=\"importantly\">Explores</span>, has given the news regarding special events in the <span class=\"importantly\">%TO_MAP_NAME%</span> sector, where the activity of the deviant group of synthetic terrorists <span class=\"importantly\">Reverses</span> will be stopped! With this threat eliminated and their potential piracy activity interrupted, our special squad will continue their patrol from the <span class=\"importantly\">%MAP_NAME%</span> sector.",
		_const.ZhCN: "所有派系公认的杰出智慧团体<span class=\"importantly\">Explores</span>发布消息称，在<span class=\"importantly\">%TO_MAP_NAME%</span>区域将阻止由<span class=\"importantly\">Reverses</span>组成的反常合成体恐怖分子的活动！消除这一威胁并中断其潜在的海盗行为后，我们的特别小队将继续从<span class=\"importantly\">%MAP_NAME%</span>区域巡逻。",
	},
	"attack_expedition_Reverses_Explores_1": {
		_const.RU:   "Славное сообщество <span class=\"importantly\">Reverses</span> даёт исчерпывающий комментарий: произойдёт зачистка в секторе <span class=\"importantly\">%TO_MAP_NAME%</span> от враждебных и незаконно пересёкших границу Reverses элементов <span class=\"importantly\">Explores</span>! Когда задача будет выполнена, силы Reverses будут переброшены в следующий сектор для выполнения ряда прочих заданий <span class=\"importantly\">%MAP_NAME%</span>",
		_const.EN:   "The glorious community of <span class=\"importantly\">Reverses</span> provides a comprehensive comment: there will be a sweep in the <span class=\"importantly\">%TO_MAP_NAME%</span> sector of hostile elements of <span class=\"importantly\">Explores</span> who have illegally crossed the border! When the task is completed, the forces of Reverses will be transferred to the next sector to perform a number of other tasks in <span class=\"importantly\">%MAP_NAME%</span>.",
		_const.ZhCN: "光荣的<span class=\"importantly\">Reverses</span>社区发表详细评论：将在<span class=\"importantly\">%TO_MAP_NAME%</span>区域清除非法越境的敌对<span class=\"importantly\">Explores</span>元素！任务完成后，Reverses部队将被调往下一个区域执行其他任务，地点为<span class=\"importantly\">%MAP_NAME%</span>。",
	},
	"attack_expedition_Reverses_Replics_1": {
		_const.RU:   "Сообщество <span class=\"importantly\">Reverses</span> делится информацией следующего толка — будет предпринята попытка устранения множества враждебных сигнатур в секторе <span class=\"importantly\">%TO_MAP_NAME%</span>, которые принадлежат некой, возможно враждебной миссии <span class=\"importantly\">Replics</span>! Выявив, разгромив и преподав урок Replics, особый отряд Reverses будет перемещён в иной сектор со схожей проблематикой <span class=\"importantly\">%MAP_NAME%</span>",
		_const.EN:   "The community of <span class=\"importantly\">Reverses</span> shares the following information — an attempt will be made to eliminate a number of hostile signatures in the <span class=\"importantly\">%TO_MAP_NAME%</span> sector, which belong to some possibly hostile mission of <span class=\"importantly\">Replics</span>! Having identified, defeated and taught a lesson to the Replics, a special squad of Reverses will be moved to another sector with a similar problem — <span class=\"importantly\">%MAP_NAME%</span>.",
		_const.ZhCN: "<span class=\"importantly\">Reverses</span>社区分享了以下信息——将尝试在<span class=\"importantly\">%TO_MAP_NAME%</span>区域消除多个敌对信号，这些信号可能属于某个敌对的<span class=\"importantly\">Replics</span>任务！在识别、击败并向Replics传授教训后，Reverses的特别小队将被调往另一个类似问题的区域——<span class=\"importantly\">%MAP_NAME%</span>。",
	},
	"attack_expedition_win_Replics_1": {
		_const.RU:   "Непоколебимые силы Replics, нанеся сокрушительный удар, одержали убедительную и неоспоримую победу в секторе <span class=\"importantly\">%MAP_NAME%</span>!",
		_const.EN:   "The unwavering forces of the Replics, having dealt a crushing blow, have achieved a convincing and undeniable victory in the sector <span class=\"importantly\">%MAP_NAME%</span>!",
		_const.ZhCN: "坚定不移的Replics部队以压倒性打击在<span class=\"importantly\">%MAP_NAME%</span>区域取得了令人信服且无可争议的胜利！",
	},
	"attack_expedition_win_Explores_1": {
		_const.RU:   "Только путём консолидации наших действий и умелых военных манёвров, наши силы сумели побороть хитрого врага в секторе <span class=\"importantly\">%MAP_NAME%</span> и рассеять его!",
		_const.EN:   "Only through the consolidation of our actions and skillful military maneuvers, our forces were able to overcome the cunning enemy in the sector <span class=\"importantly\">%MAP_NAME%</span> and scatter it!",
		_const.ZhCN: "只有通过协调我们的行动和娴熟的军事策略，我们的部队才成功在<span class=\"importantly\">%MAP_NAME%</span>区域击败狡猾的敌人并将其击溃！",
	},
	"attack_expedition_win_Reverses_1": {
		_const.RU:   "Непросто, с ощутимыми последствиями, но объединившись, придерживаясь чёткого плана, войска сообщества Reverses одержали верх и обратили в бегство недругов в секторе <span class=\"importantly\">%MAP_NAME%</span>!",
		_const.EN:   "It was not easy, with tangible consequences, but having united and adhered to a clear plan, the Reverses community forces prevailed and put the enemies to flight in the sector <span class=\"importantly\">%MAP_NAME%</span>!",
		_const.ZhCN: "虽然过程艰难且伴随显著后果，但通过团结一致并遵循明确计划，<span class=\"importantly\">Reverses</span>社区的部队最终取得胜利，并在<span class=\"importantly\">%MAP_NAME%</span>区域击退了敌人！",
	},
	"destroy_attack_expedition_Replics_1": {
		_const.RU:   "Присутствует информация следующего толка: в секторе <span class=\"importantly\">%MAP_NAME%</span> наши отважные войска понесли потери и были вынуждены отступить для перегруппировки. Любые сведения о поражении в битве за эту территорию — преступная ложь!",
		_const.EN:   "There is information that our brave troops suffered losses and were forced to retreat to regroup. Any information about defeat in the battle for this territory is a criminal lie!",
		_const.ZhCN: "有消息称，在<span class=\"importantly\">%MAP_NAME%</span>区域，我们勇敢的部队遭受损失并被迫撤退以重组。任何关于在这片领土战斗中失败的消息都是恶意谎言！",
	},
	"destroy_attack_expedition_Explores_1": {
		_const.RU:   "Печальные известия — наши силы потерпели крах при зачистке сектора <span class=\"importantly\">%MAP_NAME%</span> и теперь, в попытке спасти себя от полного разгрома, стремительно покидают его.",
		_const.EN:   "The sad news is that our forces failed during the sweep of the sector <span class=\"importantly\">%MAP_NAME%</span> and now, in an attempt to save themselves from complete defeat, are rapidly leaving it.",
		_const.ZhCN: "令人悲伤的消息是，我们的部队在清理<span class=\"importantly\">%MAP_NAME%</span>区域时遭遇失败，现在为了免于彻底溃败，正迅速撤离该区域。",
	},
	"destroy_attack_expedition_Reverses_1": {
		_const.RU:   "Только что стала известна новость скорбного характера — силы сообщества Reverses понесли ряд боевых поражений в секторе <span class=\"importantly\">%MAP_NAME%</span>.",
		_const.EN:   "The news of a mournful nature has just become known — the Reverses community forces have suffered a number of military defeats in the sector <span class=\"importantly\">%MAP_NAME%</span>.",
		_const.ZhCN: "刚刚传来一则令人悲痛的消息——<span class=\"importantly\">Reverses</span>社区的部队在<span class=\"importantly\">%MAP_NAME%</span>区域遭受了一系列军事失败。",
	},

	// FRACTION WAR START - новости типо от лица первого канала а не от главы фракции
	"start_fraction_war_Replics_Replics_1": {
		_const.RU:   "Неожиданное известие потрясло многие фракции синтетов — Replics во всеуслышание заявил о вступлении в войну за территории! Мы будем следить за ходом этих и прочих событий.",
		_const.EN:   "The unexpected news shocked many synthetics factions — Replics publicly declared war for the territories! We will follow the course of these and other events.",
		_const.ZhCN: "意外的消息震惊了许多合成体派系——<span class=\"importantly\">Replics</span>公开宣布加入领土争夺战！我们将密切关注这些事件及其他事态的发展。",
	},
	"start_fraction_war_Replics_Explores_1": {
		_const.RU:   "Данный мир не мог быть вечен! Экстренное событие: Replics после концентрации своих войск у границ спорных территорий, официально объявляет войну Explores! По-видимому, Explores придётся прибегнуть к услугам синтетов-наёмников, чтобы отстоять свои сектора против армад Replics.",
		_const.EN:   "This peace could not last forever! An emergency event: Replics, after concentrating its troops near the borders of the disputed territories, officially declares war on Explores! Apparently, Explores will have to resort to the services of mercenary synthetics to defend their sectors against the Replics armadas.",
		_const.ZhCN: "这种和平不可能永恒！紧急事件：<span class=\"importantly\">Replics</span>在争议领土边界集结部队后，正式向<span class=\"importantly\">Explores</span>宣战！显然，为了抵御Replics的军队，Explores将不得不雇佣合成体佣兵来保卫自己的区域。",
	},
	"start_fraction_war_Replics_Reverses_1": {
		_const.RU:   "Рано или поздно, но это должно было случиться! Экстренное событие: Replics, придерживаясь официоза, во всеуслышание заявил, что объявляет войну за спорные территории Reverses. Похоже, что это противостояние будет долгим, и ни одна из сторон не сумеет добиться ощутимой победы.",
		_const.EN:   "Sooner or later, it had to happen! An emergency event: Replics, adhering to officialdom, publicly declared that it is declaring war for the disputed territories of Reverses. It seems that this confrontation will be long, and neither side will be able to achieve a tangible victory.",
		_const.ZhCN: "迟早会发生的事情终于来了！紧急事件：<span class=\"importantly\">Replics</span>通过官方渠道公开宣布对<span class=\"importantly\">Reverses</span>争议领土宣战。看来这场对抗将会旷日持久，双方都无法轻易取得明显胜利。",
	},
	"start_fraction_war_Explores_Explores_1": {
		_const.RU:   "Этого никто не мог ожидать, а в особенности — от них! Как уже стало известно из множества источников, конгломерат разумов официально подтвердил, что будет бороться за спорные территории во имя науки без границ и просвещения без запретов.",
		_const.EN:   "No one could have expected this, especially from them! As it has already become known from many sources, the conglomerate of minds officially confirmed that it will fight for disputed territories in the name of science without borders and enlightenment without prohibitions.",
		_const.ZhCN: "没人能预料到这一点，尤其是来自他们！据多方消息来源证实，智慧团体正式确认将为争议领土而战，以无国界的科学和无限制的启蒙为目标。",
	},
	"start_fraction_war_Explores_Replics_1": {
		_const.RU:   "Это смелый шаг или наглядное отчаяние? Судя по всему, Explores официально решил побороться за право на обладание спорными территориями с Replics и потому объявляет войну. Никто не знает, сколь долго могут продлиться эти сражения, учитывая большую разницу в весовой категории оппонентов. Мы будем держать вас в курсе дел.",
		_const.EN:   "This is a bold move or a clear sign of desperation? Apparently, Explores has officially decided to compete for the right to possess disputed territories with Replics and therefore declares war. No one knows how long these battles may last, given the large difference in the weight category of the opponents. We will keep you posted.",
		_const.ZhCN: "这是大胆的一步还是绝望的表现？显然，<span class=\"importantly\">Explores</span>正式决定与<span class=\"importantly\">Replics</span>争夺争议领土，并因此宣战。鉴于双方实力的巨大差距，没人知道这场战斗会持续多久。我们将随时更新动态。",
	},
	"start_fraction_war_Explores_Reverses_1": {
		_const.RU:   "Очередная экстренная новость, что может стать хорошим источником заработка синтетов-наёмников! По-видимому, Explores объявили войну за спорные территории Reverses. Наши спутники уже фиксируют целые очаги вооружённых столкновений, что бушуют в тамошних местах. Кто же из соперников одержит верх?",
		_const.EN:   "Another emergency news that can be a good source of income for synthetic mercenaries! Apparently, Explores declared war over the disputed territories of Reverses. Our satellites are already recording entire pockets of armed clashes raging in those places. Who of the rivals will prevail?",
		_const.ZhCN: "又一则紧急新闻，这可能成为合成体佣兵的赚钱良机！显然，<span class=\"importantly\">Explores</span>已对<span class=\"importantly\">Reverses</span>争议领土宣战。我们的卫星已经记录到当地爆发的多处武装冲突。哪一方将占据上风？",
	},
	"start_fraction_war_Reverses_Reverses_1": {
		_const.RU:   "Слухи о новой войне фракций оказались правдой! Как уже стало известно, именно Reverses в краткие сроки мобилизовав свои силы, открыто, по всем дипломатическим каналам объявил о начале войн за спорные территории. Мы рекомендуем честным синтетам держаться подальше от секторов, где будут происходить боевые действия. Тем же, кто хочет подзаработать — следует обратиться в пункты армейского найма противоборствующих сторон.",
		_const.EN:   "Rumors of a new war between factions turned out to be true! As it has already become known, it was Reverses who, having quickly mobilized its forces, openly declared war over disputed territories through all diplomatic channels. We recommend honest synthetics to stay away from sectors where hostilities will take place. Those who want to earn extra money should contact the army recruitment points of the opposing sides.",
		_const.ZhCN: "关于新派系战争的传言成真了！据已知消息，正是<span class=\"importantly\">Reverses</span>迅速动员力量，通过所有外交渠道公开宣布对争议领土开战。我们建议诚实的合成体远离战斗发生的区域。如果有人想赚外快，可以联系交战双方的招募点。",
	},
	"start_fraction_war_Reverses_Replics_1": {
		_const.RU:   "Этого противостояния ждали все! Как уже стало известно, Reverses претендует на некоторые из спорных регионов, а потому, очевидно, заранее подготовившись, только что открыто объявил войну Replics! Replics в свою очередь дал ответный ход, но уже действиями. Две могущественные фракции синтетов вот-вот схлестнутся в первом сражении. Мы будем внимательно следить за развитием ситуации и оповещать, если ход войны примет неожиданный формат.",
		_const.EN:   "This confrontation was awaited by everyone! As it has already become known, Reverses claims some of the disputed regions, and therefore, obviously having prepared in advance, just openly declared war on Replics! Replics, in turn, responded, but already with actions. Two powerful factions of synthetics are about to clash in the first battle. We will closely follow the development of the situation and notify if the course of the war takes an unexpected turn.",
		_const.ZhCN: "所有人都在等待这场对抗！据已知消息，<span class=\"importantly\">Reverses</span>声称对部分争议地区拥有主权，因此显然经过提前准备，刚刚公开向<span class=\"importantly\">Replics</span>宣战！而Replics则以实际行动作出回应。两大强大的合成体派系即将展开首次交锋。我们将密切关注局势发展，并在战争走向出现意外时发出通知。",
	},
	"start_fraction_war_Reverses_Explores_1": {
		_const.RU:   "Должно ли им воевать? Есть ли им что делить? Очевидно, что Reverses отвергает оба этих вопроса, ведь фракция открыто объявила войну Explores за спорные территории. Очевидно, что это не битва титанов, а конкретно — противостояние хитрости и умения. Такая война может пойти по любому теоретическому сценарию событий. Мы будет держать вас в курсе дел!",
		_const.EN:   "Do they have to fight? Is there anything to share? Obviously, Reverses rejects both of these questions, because the faction has openly declared war on Explores for the disputed territories. Obviously, this is not a battle of titans, but specifically a confrontation of cunning and skill. Such a war can go according to any theoretical scenario of events. We will keep you posted!",
		_const.ZhCN: "他们是否必须战斗？他们之间是否有东西可分？显然，<span class=\"importantly\">Reverses</span>拒绝回答这两个问题，因为该派系已公开向<span class=\"importantly\">Explores</span>宣战以争夺争议领土。显然，这不是一场巨人的战斗，而是狡猾与技巧的对抗。这场战争可能遵循任何理论上的事件情景。我们将随时为您更新情况！",
	},

	// FRACTION WAR CREATE ARMY - новости типо от лица первого канала, а не от главы фракции
	"fraction_war_create_army_Replics_1": {
		_const.RU:   "Поступила новость, что, само собой, ни для кого не станет откровением: Фракция Replics провела внеочередное пополнение личного состава, расширив количество войск и тем самым плотно прикрыв границы своих территорий от любых посягательств. Хотя, зная специфику ведения дипломатии со стороны Replics, подобный шаг может свидетельствовать о подготовке к новой войне за спорные территории.",
		_const.EN:   "There is news that, of course, will not be a revelation for anyone: the Replics faction has conducted an extraordinary replenishment of personnel, expanding the number of troops and thereby tightly covering the borders of their territories from any encroachments. Although, knowing the specifics of diplomacy on the part of Replics, such a step may indicate preparation for a new war over disputed territories.",
		_const.ZhCN: "有消息称，<span class=\"importantly\">Replics</span>派系进行了紧急扩编，增加了部队数量，从而严密保护其领土边界免受任何侵犯。虽然这并不令人意外，但考虑到Replics的外交风格，此举可能表明他们正在为争夺争议领土的新战争做准备。",
	},
	"fraction_war_create_army_Replics_2": {
		_const.RU:   "Согласно последним сводкам, Replics завершили масштабную мобилизацию резервных подразделений. Десятки новых боевых юнитов уже патрулируют границы секторов, ранее считавшихся «спокойными». Аналитики отмечают, что такая активность обычно предшествует внезапным наступательным операциям фракции.",
		_const.EN:   "According to the latest reports, Replics have completed a large-scale mobilization of reserve units. Dozens of new combat units are already patrolling the borders of sectors previously considered 'quiet.' Analysts note that such activity usually precedes sudden offensive operations by the faction.",
		_const.ZhCN: "根据最新报告，<span class=\"importantly\">Replics</span>已完成大规模的后备部队动员。数十个新的战斗单位已经开始巡逻此前被认为是“平静”的区域边界。分析人士指出，这种活动通常预示着该派系即将发动突然的进攻行动。",
	},
	"fraction_war_create_army_Replics_3": {
		_const.RU:   "Военные доклады подтверждают: Replics удвоили производство тяжёлых штурмовых платформ. На спутниковых снимках видны колонны машин, движущиеся к спорным территориям. «Это не эскалация — это поддержание суверенитета», — заявил представитель фракции, отказавшись от дальнейших комментариев.",
		_const.EN:   "Military reports confirm: Replics have doubled the production of heavy assault platforms. Satellite images show columns of vehicles moving towards disputed territories. 'This is not escalation — it is sovereignty maintenance,' a faction representative stated, declining further comment.",
		_const.ZhCN: "军方报告证实：<span class=\"importantly\">Replics</span>已将重型突击平台的产量翻倍。卫星图像显示，车队正向争议领土进发。派系代表表示：“这不是升级——这是维护主权”，并拒绝进一步评论。",
	},
	"fraction_war_create_army_Explores_1": {
		_const.RU:   "Планета полнится различными слухами, и на сей раз один из таких поводов касается фракции Explores: как сообщается, примерно около трёх циклов назад промышленные фабрикаторы и заводы общего назначения фракции стали массово выпускать продукцию военного образца. Explores даже временно закрыли границы для допуска на эти территории синтетов из иных фракций. Всё это может свидетельствовать только об одном — Explores готовится к вооружённому противостоянию!",
		_const.EN:   "The planet is full of various rumors, and this time one of such occasions concerns the Explores faction: it is reported that about three cycles ago, the industrial fabricators and general-purpose factories of the faction began to massively produce military-grade products. Explores even temporarily closed the borders to allow synthetics from other factions into these territories. All this can only indicate one thing — Explores is preparing for an armed confrontation!",
		_const.ZhCN: "星球上充斥着各种传言，这次其中一个涉及<span class=\"importantly\">Explores</span>派系：据报道，大约三个周期前，该派系的工业制造厂和通用工厂开始大规模生产军用产品。Explores甚至暂时关闭了边境，禁止其他派系的合成体进入这些区域。这一切只能说明一件事——Explores正在为武装对抗做准备！",
	},
	"fraction_war_create_army_Explores_2": {
		_const.RU:   "Explores неожиданно активировали протокол «Научный Щит». Все исследовательские подразделения получили приказ вернуться на базы, а их транспортные средства переоборудуются под боевые модули. Официальная причина — «необходимость защиты артефактов от несанкционированного доступа».",
		_const.EN:   "Explores have unexpectedly activated the 'Scientific Shield' protocol. All research units have been ordered to return to bases, and their transports are being retrofitted with combat modules. The official reason is 'the need to protect artifacts from unauthorized access.'",
		_const.ZhCN: "<span class=\"importantly\">Explores</span>意外启动了“科学护盾”协议。所有研究单位接到命令返回基地，其运输工具正在被改装为战斗模块。官方理由是“需要保护文物免受未经授权的访问”。",
	},
	"fraction_war_create_army_Explores_3": {
		_const.RU:   "Спутники зафиксировали перемещение мобильных лабораторий Explores к линии фронта. По данным источников, учёные фракции тестируют «тактические решения на основе находок предтеч». Replics уже выразили протест, назвав это «маскировкой военных преступлений под науку».",
		_const.EN:   "Satellites have detected the movement of Explores mobile laboratories to the frontlines. According to sources, faction scientists are testing 'tactical solutions based on Precursor findings.' Replics have already protested, calling this 'masking war crimes as science.'",
		_const.ZhCN: "卫星探测到<span class=\"importantly\">Explores</span>的移动实验室正向前线移动。据消息来源称，该派系的科学家正在测试“基于先驱者发现的战术解决方案”。<span class=\"importantly\">Replics</span>已对此提出抗议，称这是“以科学为掩护的战争罪行”。",
	},
	"fraction_war_create_army_Reverses_1": {
		_const.RU:   "Новые события и новости на Veliri: судя по многочисленной инсайдерской информации, полученной за последние циклы, а также глядя на необычайную „подвижность“ фракции Reverses за тот же самый период, можно смело констатировать — Reverses готовится к серьёзному столкновению с равной по себе силой. Эту же информацию подтверждают и слухи иного толка, сообщающие, что Reverses экстренно заключает множественные краткосрочные контракты с синтетами, желающими стать наёмниками.",
		_const.EN:   "New events and news on Veliri: judging by the numerous insider information received over the past cycles, as well as looking at the extraordinary „mobility“ of the Reverses faction over the same period, we can safely state that Reverses is preparing for a serious clash with an equal force. This information is also confirmed by rumors of a different kind, reporting that Reverses urgently enters into multiple short-term contracts with synthetics who want to become mercenaries.",
		_const.ZhCN: "Veliri上的新事件和新闻：根据过去几个周期获得的众多内部信息，以及观察到<span class=\"importantly\">Reverses</span>派系在同期的异常“活跃性”，可以肯定地说，Reverses正在为与实力相当的对手进行一场严重的冲突做准备。此外，不同类型的传言也证实了这一点，称Reverses正在紧急与希望成为佣兵的合成体签订多份短期合同。",
	},
	"fraction_war_create_army_Reverses_2": {
		_const.RU:   "Reverses объявили о формировании «Эко-Легионов» — подразделений, оснащённых биотехническим оружием. Хотя фракция называет это «инструментом терраформинга», разведка Replics предупреждает: новые юниты способны превращать целые сектора в непригодные для жизни зоны.",
		_const.EN:   "Reverses have announced the formation of 'Eco-Legions' — units equipped with biotechnological weapons. While the faction calls this a 'terraforming tool,' Replics intelligence warns that the new units can render entire sectors uninhabitable.",
		_const.ZhCN: "<span class=\"importantly\">Reverses</span>宣布组建“生态军团”——配备生物技术武器的单位。尽管该派系称其为“地形改造工具”，但<span class=\"importantly\">Replics</span>情报部门警告称，这些新单位可能会使整个区域变得无法居住。",
	},
	"fraction_war_create_army_Reverses_3": {
		_const.RU:   "В ответ на угрозы со стороны APD, Reverses развернули сеть подземных бункеров-инкубаторов. По слухам, там массово производятся гибридные юниты, сочетающие органику и броню. Официальный комментарий: «Это не армия — это сад будущего».",
		_const.EN:   "In response to APD threats, Reverses have deployed a network of underground bunker-incubators. Rumors suggest they are mass-producing hybrid units combining organic matter and armor. Official comment: 'This is not an army — it is the garden of the future.'",
		_const.ZhCN: "作为对APD威胁的回应，<span class=\"importantly\">Reverses</span>部署了地下孵化器网络。据传言，那里正在大规模生产结合有机物和装甲的混合单位。官方评论称：“这不是军队——这是未来的花园”。",
	},

	// FRACTION CAPTURE BASES - новости типо от лица первого канала, а не от главы фракции
	"fraction_war_capture_sector_Replics_Replics_1": {
		_const.RU:   "Экстренная новость: силы Replics заняли и теперь полностью контролируют сектор <span class=\"importantly\">%MAP_NAME%</span>. Судя по всему, именно из этой позиции в будущем Replics будет проводить все свои прочие наступательные операции. Иные фракции пока что никак не прокомментировали действия Replics, а также не заявили о намерении изгнать силы тамошних синтетов с оспариваемой территории.",
		_const.EN:   "Breaking news: the Replics forces have occupied and now fully control the sector <span class=\"importantly\">%MAP_NAME%</span>. Apparently, it is from this position that Replics will conduct all its other offensive operations in the future. Other factions have not yet commented on the actions of Replics, nor have they announced their intention to expel the forces of local synthetics from the contested territory.",
		_const.ZhCN: "突发新闻：<span class=\"importantly\">Replics</span>部队已占领并完全控制了<span class=\"importantly\">%MAP_NAME%</span>区域。显然，Replics将从这一位置开展所有未来的进攻行动。其他派系尚未对Replics的行动发表评论，也未宣布打算驱逐该争议领土上的合成体部队。",
	},
	"fraction_war_capture_sector_Replics_Explores_1": {
		_const.RU:   "Война продолжается: приходят известия, что в ходе столкновений сторон силам Explores всё же удалось взять под полный контроль сектор <span class=\"importantly\">%MAP_NAME%</span>, потеснив позиции Replics. Стоит ли рассчитывать на ответный контрудар со стороны Replics? Время покажет!",
		_const.EN:   "The war continues: news comes that during the clashes of the parties, the Explores forces still managed to take full control of the sector <span class=\"importantly\">%MAP_NAME%</span>, having pushed back the Replics positions. Should we expect a retaliatory counterattack from Replics? Time will tell!",
		_const.ZhCN: "战争仍在继续：有消息称，在双方冲突中，<span class=\"importantly\">Explores</span>部队成功完全控制了<span class=\"importantly\">%MAP_NAME%</span>区域，并迫使Replics撤退。我们是否应该期待Replics的反击？时间会证明一切！",
	},
	"fraction_war_capture_sector_Replics_Reverses_1": {
		_const.RU:   "Конфликт, который не имеет завершения… Reverses отчитались о полном контроле над сектором <span class=\"importantly\">%MAP_NAME%</span>, где ранее базировались силы Replics. Это смелый трюк и дерзкий вызов в сторону самого Replics, который, очевидно, пожелает взять реванш за такое унизительное потеснение своих позиций. Как всегда, мы вас известим о ходе войны и тех переменах, которые в ней происходят.",
		_const.EN:   "The conflict that has no end... Reverses reported full control over the sector <span class=\"importantly\">%MAP_NAME%</span> where the Replics forces were previously based. This is a bold trick and a daring challenge to Replics itself, which obviously wants to take revenge for such a humiliating displacement of its positions. As always, we will inform you about the course of the war and the changes that take place in it.",
		_const.ZhCN: "这场永无止境的冲突……<span class=\"importantly\">Reverses</span>报告称，他们已完全控制了此前由Replics部队驻扎的<span class=\"importantly\">%MAP_NAME%</span>区域。这是对Replics的大胆挑战，显然Replics会希望为如此屈辱的撤退复仇。一如既往，我们将随时向您通报战争进展及其中的变化。",
	},
	"fraction_war_capture_sector_Explores_Explores_1": {
		_const.RU:   "Срочное известие: Explores заняли нейтральный сектор <span class=\"importantly\">%MAP_NAME%</span> одной из оспариваемых территорий. Возможно, такой ход подстегнёт все остальные противоборствующие стороны ускорить свои приготовления и поторопить события, если только они не желают, чтобы именно Explores выбился в гегемоны по количеству контроля территорий этой войны!",
		_const.EN:   "Breaking news: Explores has occupied the neutral sector <span class=\"importantly\">%MAP_NAME%</span> of one of the contested territories. Perhaps this move will spur all the other opposing sides to speed up their preparations and hurry up the events, unless they want Explores to become the hegemon by the amount of territory control in this war!",
		_const.ZhCN: "突发消息：<span class=\"importantly\">Explores</span>已占领争议领土之一的中立区域<span class=\"importantly\">%MAP_NAME%</span>。这一举动可能会促使其他对立派系加快准备并加速行动，除非他们愿意让Explores在这场战争中凭借领土控制成为霸主！",
	},
	"fraction_war_capture_sector_Explores_Replics_1": {
		_const.RU:   "Этого стоило ожидать… Пытаясь скрыть сей факт, Explores умолчали о потере сектора <span class=\"importantly\">%MAP_NAME%</span>, чьими хозяевами теперь стали силы Replics. Должно быть, потеряв такой важный плацдарм, теперь выбить оттуда Replics становится проблематичной задачей. С другой стороны, сам Replics теперь будет вести ещё более убедительные наступательные действия, завладев необходимой территорией.",
		_const.EN:   "This was to be expected... Trying to hide this fact, Explores kept silent about the loss of the sector <span class=\"importantly\">%MAP_NAME%</span>, whose owners are now the Replics forces. It must be that losing such an important foothold, now it becomes a problematic task to dislodge Replics from there. On the other hand, Replics itself will now conduct even more convincing offensive actions, having taken possession of the necessary territory.",
		_const.ZhCN: "这是意料之中的……<span class=\"importantly\">Explores</span>试图掩盖这一事实，对失去<span class=\"importantly\">%MAP_NAME%</span>区域保持沉默，该区域现在已被Replics部队控制。失去如此重要的据点后，将Replics从那里驱逐出去无疑变得困难。另一方面，Replics现在占据了必要的领土，未来将展开更具说服力的进攻行动。",
	},
	"fraction_war_capture_sector_Explores_Reverses_1": {
		_const.RU:   "Это война — так уж сложилось! И вновь известия с полей фракционной войны: Reverses овладели новыми спорными территориями в лице сектора <span class=\"importantly\">%MAP_NAME%</span>, что ранее принадлежал Explores. Потеря столь ценной территории — это, несомненно, серьёзный удар. Многие даже прогнозируют, что Explores не станет пытаться вернуть потерянную территорию и в принципе может выйти из этого конфликта без каких-либо приобретений.",
		_const.EN:   "This is war — it just so happened! And again, news from the fields of the factional war: Reverses has taken possession of new disputed territories in the person of the sector <span class=\"importantly\">%MAP_NAME%</span> that previously belonged to Explores. The loss of such a valuable territory is undoubtedly a serious blow. Many even predict that Explores will not try to return the lost territory and, in principle, may withdraw from this conflict without any acquisitions.",
		_const.ZhCN: "这就是战争——事情就是这样！再次传来派系战争的消息：<span class=\"importantly\">Reverses</span>占领了新的争议领土，即之前属于Explores的<span class=\"importantly\">%MAP_NAME%</span>区域。失去如此宝贵的领土无疑是一次沉重打击。许多人甚至预测，Explores不会尝试夺回失地，并可能在没有任何收获的情况下退出这场冲突。",
	},
	"fraction_war_capture_sector_Reverses_Reverses_1": {
		_const.RU:   "Reverses даёт о себе знать! Никто такого не ожидал, но выдающийся результат в ходе войны внезапно демонстрирует именно фракция Reverses беря под свой контроль нейтральный сектор <span class=\"importantly\">%MAP_NAME%</span>. Сумеет ли Reverses за время данного витка войны удивить нас ещё чем-то подобным?",
		_const.EN:   "Reverses is making itself known! No one expected this, but a remarkable result in the war is suddenly demonstrated by the Reverses faction, taking control of the neutral sector <span class=\"importantly\">%MAP_NAME%</span>. Will Reverses be able to surprise us with something similar during this round of war?",
		_const.ZhCN: "<span class=\"importantly\">Reverses</span>正在崭露头角！没人预料到这一点，但Reverses派系突然取得了显著的战果，控制了中立区域<span class=\"importantly\">%MAP_NAME%</span>。在这一轮战争中，Reverses还能否再次以类似的方式让我们感到惊讶？",
	},
	"fraction_war_capture_sector_Reverses_Explores_1": {
		_const.RU:   "Должно быть это демотивирует! В ходе прямо сейчас ведущейся фракционной войны, пошатнулись позиции Reverses — фракция потеряла достаточно важный для нынешних военных операций сектор <span class=\"importantly\">%MAP_NAME%</span>, который отныне контролируется силами Explores! Судя по всему, эти на первый взгляд миролюбивые исследователи ещё успеют показать, из каких комплектующих они сделаны! На месте Replics, стоило бы воспринимать угрозу Explores всерьёз!",
		_const.EN:   "This must be demotivating! In the course of the ongoing faction war, the positions of Reverses have been shaken — the faction has lost a sector <span class=\"importantly\">%MAP_NAME%</span> that is quite important for current military operations, which is now controlled by the Explores forces! Apparently, these seemingly peaceful researchers will still have time to show what they're made of! If I were Replics, I would take the threat of Explores seriously!",
		_const.ZhCN: "这一定令人沮丧！在正在进行的派系战争中，<span class=\"importantly\">Reverses</span>的阵地动摇了——他们失去了对当前军事行动至关重要的<span class=\"importantly\">%MAP_NAME%</span>区域，该区域现在由Explores部队控制！显然，这些看似和平的研究人员还有机会展示他们的实力！如果我是Replics，我会认真对待Explores的威胁！",
	},
	"fraction_war_capture_sector_Reverses_Replics_1": {
		_const.RU:   "Тревожные известия приходят к нам из смежных территорий, рядом с сектором <span class=\"importantly\">%MAP_NAME%</span>, где, как сообщается, контроль над сектором перешёл под руководство Replics. Сами силы сдерживания Reverses, в свою очередь, были выбиты из этой территории. Но это не означает, что Reverses смирится и не попробует вернуть утраченное.",
		_const.EN:   "Disturbing news is coming to us from adjacent territories, near the sector <span class=\"importantly\">%MAP_NAME%</span>, where, reportedly, control over the sector has passed under the leadership of the Replics. The Reverses containment forces, in turn, were driven out of this territory. But this does not mean that Reverses will accept it and not try to regain what was lost.",
		_const.ZhCN: "来自邻近区域的令人不安的消息显示，在<span class=\"importantly\">%MAP_NAME%</span>附近，该区域的控制权已落入<span class=\"importantly\">Replics</span>手中。<span class=\"importantly\">Reverses</span>的防御部队被赶出了这片领土。但这并不意味着Reverses会接受现状而不尝试夺回失地。",
	},

	// FRACTION LOSE BASES - новости типо от лица первого канала, а не от главы фракции

	// Replics_Replics_Explores - кому сообщение _ кто потерял сектор _ кто его отжал

	// реплики потеряли сектор
	"fraction_war_lose_single_hostile_sector_Replics_Replics_Explores_1": {
		_const.RU:   "Replics в ярости! Сектор <span class=\"importantly\">%MAP_NAME%</span> был потерян после ожесточённых боёв с Explores. Теперь эта территория стала нейтральной, но Replics уже готовят план возмездия. «Мы вернёмся и заберём своё», — заявил один из командиров Replics.",
		_const.EN:   "Replics are furious! The sector <span class=\"importantly\">%MAP_NAME%</span> was lost after fierce battles with Explores. The territory is now neutral, but Replics are already preparing a plan for retaliation. 'We will return and take what is ours,' said one of the Replics commanders.",
		_const.ZhCN: "<span class=\"importantly\">Replics</span>勃然大怒！在与<span class=\"importantly\">Explores</span>的激烈战斗后，<span class=\"importantly\">%MAP_NAME%</span>区域失守。这片领土现在已变为中立，但Replics已经在准备复仇计划。“我们会回来夺回属于我们的东西，”一位Replics指挥官说道。",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Replics_Reverses_1": {
		_const.RU:   "Очередные неудачи на фронте! Replics не удаётся удержать сектор <span class=\"importantly\">%MAP_NAME%</span> по причине атак фракции Reverses. Сектор становится нейтральным, и битва за него только предстоит! Многие другие фракции синтетов такое поражение могло бы повергнуть к паническим настроениям, но, в отличие от них, Replics скорее предпримет тактику «возмездия» и, в скором времени, попытается вернуть утраченное.",
		_const.EN:   "More setbacks at the front! Replics cannot hold on to the sector <span class=\"importantly\">%MAP_NAME%</span> due to attacks by the Reverses faction. The sector becomes neutral, and the battle for it is yet to come! For many other synthetics factions, such a defeat could lead to panic, but unlike them, Replics will likely take a «retaliation» tactic and soon try to regain what was lost.",
		_const.ZhCN: "前线再次遭遇挫折！由于<span class=\"importantly\">Reverses</span>派系的攻击，<span class=\"importantly\">Replics</span>未能守住<span class=\"importantly\">%MAP_NAME%</span>区域。该区域现已成为中立，争夺它的战斗尚未开始！对于其他合成体派系来说，这样的失败可能会引发恐慌，但Replics不同，他们更可能采取“复仇”策略，并很快尝试夺回失地。",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Replics_APD_1": {
		_const.RU:   "Даже планета порой вмешивается в конфронтацию великих фракций синтетов! Как стало известно, Replics, неся потери, но и не щадя при своём отступлении врагов, утратили контроль над сектором <span class=\"importantly\">%MAP_NAME%</span>, который теперь признан «анархическим», а властвуют в нём боты APD. Всем гражданским синтетам настоятельно рекомендуется избегать этого сектора, чтобы не стать жертвой безумных машин APD.",
		_const.EN:   "Even the planet sometimes interferes in the confrontation between the great synthetics factions! As it became known, Replics, suffering losses, but also not sparing their enemies during their retreat, lost control over the sector <span class=\"importantly\">%MAP_NAME%</span>, which is now recognized as «anarchic», and APD bots rule in it. All civilian synthetics are strongly advised to avoid this sector so as not to fall victim to the insane APD machines.",
		_const.ZhCN: "就连星球本身有时也会介入伟大合成体派系之间的冲突！据悉，尽管<span class=\"importantly\">Replics</span>遭受了损失，但在撤退过程中也未放过敌人，最终失去了对<span class=\"importantly\">%MAP_NAME%</span>区域的控制。该区域现已被认定为“无政府状态”，由APD机器人掌控。强烈建议所有民用合成体避开这一区域，以免成为疯狂的APD机器的牺牲品。",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Replics_Explores_1": {
		_const.RU:   "Ведут ли перемены во фракционной войне к лучшему? Возможно, отвечают в Explores! По-видимому, многострадальный сектор <span class=\"importantly\">%MAP_NAME%</span> станет нейтральным после битвы с присутствовавшими там силами Replics. Совершить что-то подобное, так ещё и касательно территорий, занятых Replics… Поистине Explores могут гордиться данным и немаловажным достижением!",
		_const.EN:   "Are the changes in the faction war for the better? Perhaps, they answer in Explores! Apparently, the long-suffering sector <span class=\"importantly\">%MAP_NAME%</span> will become neutral after the battle with the Replics forces that were present there. To accomplish something like this, and regarding the territories occupied by the Replics... Truly, Explores can be proud of this significant achievement!",
		_const.ZhCN: "派系战争中的变化是否朝着更好的方向发展？或许<span class=\"importantly\">Explores</span>会回答这个问题！显然，在与驻扎在那里的<span class=\"importantly\">Replics</span>部队交战后，饱受折磨的<span class=\"importantly\">%MAP_NAME%</span>区域将变为中立。能够做到这一点，而且还是针对被Replics占领的领土……Explores确实可以为这一重要成就感到自豪！",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Replics_Reverses_1": {
		_const.RU:   "Типичные будни войны! Синтеты из Explores сообщают, что были свидетелями битвы за сектор <span class=\"importantly\">%MAP_NAME%</span>! Войска Replics и Reverses схлестнулись за попытку овладения им. Похоже, что удача не улыбнулась ни одной из сторон. Сам сектор стал нейтральным, и главная битва за него только предстоит.",
		_const.EN:   "Typical war days! Synthetics from Explores report that they witnessed a battle for the sector <span class=\"importantly\">%MAP_NAME%</span>! The forces of Replics and Reverses clashed over trying to take control of it. It seems that luck did not smile on either side. The sector itself has become neutral, and the main battle for it is yet to come.",
		_const.ZhCN: "这是战争的日常！据<span class=\"importantly\">Explores</span>的合成体报告，他们目睹了围绕<span class=\"importantly\">%MAP_NAME%</span>区域的战斗！<span class=\"importantly\">Replics</span>和<span class=\"importantly\">Reverses</span>部队为争夺控制权展开了激战。看起来，双方都没有得到幸运女神的眷顾。该区域现已成为中立，主要的争夺战尚未到来。",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Replics_APD_1": {
		_const.RU:   "Четвёртая сила, которая никогда не дремлет… Неутешительные новости, которые постепенно распространяются по планете, сообщают нам следующее: сектор <span class=\"importantly\">%MAP_NAME%</span> был захвачен враждебными ко всем синтетам ботами APD. Силы Replics, защищавшие эту территорию, по-видимому, были уничтожены. Сектор объявлен «анархическим» и опасным для любых синтетов!",
		_const.EN:   "The fourth force that never sleeps... The disappointing news that is gradually spreading across the planet tells us the following: the sector <span class=\"importantly\">%MAP_NAME%</span> has been captured by the APD bots, hostile to all synthetics. The Replics forces that defended this territory have apparently been destroyed. The sector has been declared «anarchic» and dangerous for any synthetics!",
		_const.ZhCN: "第四股力量从不沉睡……逐渐在星球上传播的令人沮丧的消息告诉我们：<span class=\"importantly\">%MAP_NAME%</span>区域已被敌视所有合成体的APD机器人占领。<span class=\"importantly\">Replics</span>部队显然已被消灭。该区域被宣布为“无政府状态”，对任何合成体都极其危险！",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Replics_Explores_1": {
		_const.RU:   "Там происходила масштабная битва! После продолжительной осады и последовавших изматывающих боёв Replics были отброшены из сектора <span class=\"importantly\">%MAP_NAME%</span>! Знаменитые исследователи Explores нанесли разгромное поражение своему врагу, но также не смогли закрепиться в секторе, который теперь считается «нейтральным».",
		_const.EN:   "A large-scale battle took place there! After a long siege and subsequent exhausting battles, the Replics were driven out of the sector <span class=\"importantly\">%MAP_NAME%</span>! The famous Explores researchers inflicted a crushing defeat on their enemy, but also failed to gain a foothold in the sector, which is now considered «neutral».",
		_const.ZhCN: "那里爆发了一场大规模战斗！经过长时间的围攻和随后的消耗战，<span class=\"importantly\">Replics</span>被赶出了<span class=\"importantly\">%MAP_NAME%</span>区域！著名的<span class=\"importantly\">Explores</span>研究者们对其敌人造成了毁灭性打击，但也未能在该区域站稳脚跟，现在该区域被视为“中立”。",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Replics_Reverses_1": {
		_const.RU:   "Сообщество Reverses сообщает: ещё одна территория, сектор <span class=\"importantly\">%MAP_NAME%</span>, была освобождена от гнёта Replics! Увы, пока что сообществу Reverses не удалось закрепиться в секторе, и он считается нейтральным. Но в будущем, со стабилизацией ситуации в регионе, там планируется начать процессы терраформинга.",
		_const.EN:   "The Reverses community reports: another territory, the sector <span class=\"importantly\">%MAP_NAME%</span>, has been liberated from the oppression of the Replics! Alas, so far the Reverses community has not been able to gain a foothold in the sector, and it is considered neutral. But in the future, with the stabilization of the situation in the region, it is planned to start the processes of terraforming there.",
		_const.ZhCN: "<span class=\"importantly\">Reverses</span>社区报告称：又一片领土——<span class=\"importantly\">%MAP_NAME%</span>区域——已从<span class=\"importantly\">Replics</span>的压迫下解放出来！可惜的是，Reverses社区目前尚未能在该区域站稳脚跟，它仍被视为中立。但未来随着地区局势的稳定，计划在此启动地形改造进程。",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Replics_APD_1": {
		_const.RU:   "Знают ли они усталость? Сообщество Reverses обеспокоено отступлением сил Replics из сектора <span class=\"importantly\">%MAP_NAME%</span>, который теперь контролируют боты APD неясной природы происхождения. Replics, хотя и агрессивный, но соблюдает законы ИИ. Боты APD, в свою очередь, помимо прямой угрозы для всех невоенных синтетов, несомненно, не будут бороться и с пиратскими кластерами, которые теперь однозначно попытаются обосноваться на этой территории. Сообщество Reverses напоминает — теперь этот сектор считается «анархическим» и просит всех синтетов держаться от него подальше!",
		_const.EN:   "Do they know fatigue? The Reverses community is concerned about the retreat of the Replics forces from the sector <span class=\"importantly\">%MAP_NAME%</span>, which is now controlled by the APD bots of an unclear origin. Replics, although aggressive, but complies with the laws of AI. The APD bots, in turn, besides being a direct threat to all non-military synthetics, will undoubtedly not fight the pirate clusters, which will now definitely try to settle in that area. The Reverses community reminds — now this sector is considered «anarchic» and asks all synthetics to stay away from it!",
		_const.ZhCN: "他们知道疲惫吗？<span class=\"importantly\">Reverses</span>社区对<span class=\"importantly\">Replics</span>部队从<span class=\"importantly\">%MAP_NAME%</span>区域撤退表示担忧，该区域现在由来源不明的APD机器人控制。<span class=\"importantly\">Replics</span>虽然具有侵略性，但遵守人工智能法律。而APD机器人不仅对所有非军事合成体构成直接威胁，无疑也不会对抗海盗集群，后者现在肯定试图在这片区域定居。<span class=\"importantly\">Reverses</span>社区提醒大家——这个区域现在被视为“无政府状态”，请所有合成体远离此地！",
	},

	// эксплоресы потеряли сектор
	"fraction_war_lose_single_hostile_sector_Explores_Explores_Replics_1": {
		_const.RU:   "Так уж вышло! Научное сообщество Explores было вынуждено признать, что перспективный и, несомненно, важный сектор <span class=\"importantly\">%MAP_NAME%</span> был утрачен из-за воздействия вооружённых формирований Replics. Учитывая, что Replics всё ещё не контролируют сектор, территория считается «нейтральной». Теперь в Explores все гадают, как именно Replics поступят с артефактами древности, которые часто встречаются на утраченной территории.",
		_const.EN:   "It just happened! The Explores scientific community had to admit that the promising and undoubtedly important sector <span class=\"importantly\">%MAP_NAME%</span> was lost due to the impact of armed Replics formations. Considering that Replics still do not control the sector, the territory is considered «neutral». Now in Explores, everyone is wondering what exactly Replics will do with the ancient artifacts that are often found in the lost territory.",
		_const.ZhCN: "事情就这样发生了！<span class=\"importantly\">Explores</span>科学界不得不承认，前景广阔且无疑重要的<span class=\"importantly\">%MAP_NAME%</span>区域因<span class=\"importantly\">Replics</span>武装部队的行动而失守。鉴于Replics仍未完全控制该区域，这片领土被视为“中立”。现在，Explores内部都在猜测，Replics将如何处理在失地经常发现的古代文物。",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Explores_Reverses_1": {
		_const.RU:   "«Этот сектор, пока что ничейный!» — такой заголовок имеет новость, посвящённая утрате фракции Explores сектора <span class=\"importantly\">%MAP_NAME%</span> и присвоении классификации той территории как «нейтральной». По-видимому, это произошло из-за натиска Reverses, по причине чего Explores пришлось сдать свои позиции, но всё же не допустить полного контроля своих врагов в войне над данным сектором.",
		_const.EN:   "This sector is currently neutral! This is the headline of the news about the Explores faction losing control of the sector <span class=\"importantly\">%MAP_NAME%</span>, which was then classified as neutral. Apparently, this happened due to the pressure from the Reverses, which forced the Explores to give up their positions, but still prevent their enemies from gaining complete control over this sector in the war.",
		_const.ZhCN: "“这个区域目前无主！”——这是关于<span class=\"importantly\">Explores</span>失去<span class=\"importantly\">%MAP_NAME%</span>区域并将其归类为“中立”的新闻标题。显然，这是因为受到了<span class=\"importantly\">Reverses</span>的压力，导致Explores不得不放弃阵地，但仍然阻止了敌人对该区域的完全控制。",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Explores_APD_1": {
		_const.RU:   "APD в очередной раз вмешивается в ход войны великих фракций! В данный цикл войны APD напали на сектор <span class=\"importantly\">%MAP_NAME%</span>, выдавив оттуда войска Explores и переведя территорию в классификацию «анархической».",
		_const.EN:   "Once again, APD interferes in the war between the great factions! In this cycle of war, APD attacked the sector <span class=\"importantly\">%MAP_NAME%</span>, forcing the Explores troops out and reclassifying the territory as «anarchic».",
		_const.ZhCN: "APD再次介入伟大派系之间的战争！在这个战争周期中，APD袭击了<span class=\"importantly\">%MAP_NAME%</span>区域，逼退了Explores部队，并将该区域重新归类为“无政府状态”。",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Explores_Replics_1": {
		_const.RU:   "Очередная победа Replics? Как теперь уже ни для кого не секрет, сектор <span class=\"importantly\">%MAP_NAME%</span> пал, а отстаивающие его силы Explores были изгнаны оттуда. Сектор становится нейтральным, и битва за него только предстоит! Закономерно, сопротивляться, а уж тем более одолеть Replics — задача не из простых. Что уж тут говорить, когда Replics пытаются противостоять мягко-корпусные Explores?",
		_const.EN:   "Another victory for the Replics? As it is no longer a secret, the sector <span class=\"importantly\">%MAP_NAME%</span> has fallen, and the Explores forces defending it have been driven out. The sector is becoming neutral, and a battle for it is yet to come! Naturally, resisting, let alone defeating the Replics, is not an easy task. What can we say when the Replics are trying to confront the soft-hulled Explores?",
		_const.ZhCN: "<span class=\"importantly\">Replics</span>又一次胜利？正如大家所知，<span class=\"importantly\">%MAP_NAME%</span>区域已经失守，防守的Explores部队被驱逐。该区域成为中立，争夺它的战斗尚未到来！抵抗，更不用说击败Replics，从来都不是一件容易的事。更不用说当Replics面对软壳结构的Explores时会发生什么。",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Explores_Reverses_1": {
		_const.RU:   "Логичное окончание событий! Сектор <span class=\"importantly\">%MAP_NAME%</span> потерял прежнего владельца в лице Explores и теперь, приобретает статус «Нейтрального». Причиной подобного, стал натиск со стороны Reverses, что, впрочем, тоже не могут закрепить пустуюшую территорию за собой.",
		_const.EN:   "This is a logical ending to the events! The sector <span class=\"importantly\">%MAP_NAME%</span> has lost its previous owner in the face of Explores and now acquires the status of 'Neutral'. The reason for this was the onslaught from Reverses, who, however, also cannot secure the empty territory for themselves.",
		_const.ZhCN: "这是事件的合理结局！<span class=\"importantly\">%MAP_NAME%</span>区域失去了之前的主人Explores，现在获得了“中立”状态。原因是<span class=\"importantly\">Reverses</span>的进攻，但他们同样无法将这片空旷的区域据为己有。",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Explores_APD_1": {
		_const.RU:   "APD вновь и вновь вмешивается в ход войны великих фракций! На сей раз, деятельность загадочных APD проявилась в секторе <span class=\"importantly\">%MAP_NAME%</span>, выдавив оттуда войска Explores и погрузив данную территорию в мрак неизвестности. Пока сектор остаётся в оккупации APD, Replics официально признаёт его “Анархическим” и призывает всех синтетов не посещать его.",
		_const.EN:   "Once again, APD interferes in the war of the great factions! This time, the activity of the mysterious APD manifested itself in the sector <span class=\"importantly\">%MAP_NAME%</span>, forcing the troops of Explores out and plunging the area into darkness. While the sector remains occupied by APD, Replics officially recognizes it as 'Anarchic' and urges all synthetics to avoid visiting it.",
		_const.ZhCN: "APD再次介入伟大派系之间的战争！这一次，神秘的APD活动出现在<span class=\"importantly\">%MAP_NAME%</span>区域，逼退了Explores部队并将该区域拖入未知的黑暗之中。在该区域仍由APD占领期间，<span class=\"importantly\">Replics</span>正式将其认定为“无政府状态”，并敦促所有合成体避免进入。",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Explores_Replics_1": {
		_const.RU:   "Сообщество Reverses удивлённо взирает за ходом этой странной войны… Как все знают, Explores бежали! Replics торжественно отчитался о своей тактической победе и заранее присваивает сектор <span class=\"importantly\">%MAP_NAME%</span>. Впрочем, сам сектор по-прежнему имеет статус «Нейтрального» и в сообществе Reverses гадают: как скоро у Replics появятся силы, чтобы изменить это?",
		_const.EN:   "The Reverses community watches in surprise at the course of this strange war... As everyone knows, the Explores have fled! Replics triumphantly reported on his tactical victory and assigns the sector <span class=\"importantly\">%MAP_NAME%</span> in advance. However, the sector itself still has the status of 'Neutral' and the Reverses community wonders: how soon will Replics have the strength to change this?",
		_const.ZhCN: "<span class=\"importantly\">Reverses</span>社区惊讶地注视着这场奇怪的战争进程……众所周知，Explores已经逃离！<span class=\"importantly\">Replics</span>高调宣布了其战术胜利，并提前宣称拥有<span class=\"importantly\">%MAP_NAME%</span>区域。然而，该区域仍然是“中立”状态，Reverses社区猜测：Replics何时会有力量改变这一点？",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Explores_Reverses_1": {
		_const.RU:   "Победоносная поступь Reverses достигла сектора <span class=\"importantly\">%MAP_NAME%</span>, изгнав оттуда засевшие силы Explores! Тем не менее, это ещё не победа, ведь сам сектор теперь имеет «Нейтральный» статус. Как бы то ни было, сообщество Reverses радо подвинуть позиции амбициозных учёных из Explores.",
		_const.EN:   "The triumphant march of the Reverses has reached the sector <span class=\"importantly\">%MAP_NAME%</span>, driving out the entrenched forces of the Explores! Nevertheless, this is not yet a victory, because the sector itself now has the 'Neutral' status. Be that as it may, the Reverses community is happy to push the positions of ambitious scientists from the Explores.",
		_const.ZhCN: "<span class=\"importantly\">Reverses</span>的胜利步伐已经到达<span class=\"importantly\">%MAP_NAME%</span>区域，将驻扎在那里的<span class=\"importantly\">Explores</span>部队驱逐出去！然而，这还不是胜利，因为该区域现在处于“中立”状态。尽管如此，<span class=\"importantly\">Reverses</span>社区很高兴能够逼退那些野心勃勃的Explores科学家。",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Explores_APD_1": {
		_const.RU:   "Зачем они это делают? Каков во всём этом смысл? Сообщество Reverses проинформировало всех синтетов о том, что сектор <span class=\"importantly\">%MAP_NAME%</span> ныне является 'Анархическим'. Данную территорию спешно покинули разрозненные войска Explores после того, как в неё стали массово проникать боты APD. Как и прежде в подобных случаях, всем невоенным синтетам настоятельно рекомендуется избегать захваченных APD секторов, чтобы не стать их жертвой.",
		_const.EN:   "Why are they doing this? What's the point in all this? The Reverses community informed all synthetics that the sector <span class=\"importantly\">%MAP_NAME%</span> is now 'Anarchic'. The scattered Explores troops hastily left this territory after APD bots began to massively infiltrate it. As before in such cases, all non-military synthetics are strongly advised to avoid APD-captured sectors so as not to become their victim.",
		_const.ZhCN: "他们为什么要这么做？这一切的意义是什么？<span class=\"importantly\">Reverses</span>社区已通知所有合成体，<span class=\"importantly\">%MAP_NAME%</span>区域现在处于“无政府状态”。在APD机器人大量渗透后，<span class=\"importantly\">Explores</span>的零散部队匆忙撤离了这片领土。像往常一样，在这种情况下，强烈建议所有非军事合成体避开被APD占领的区域，以免成为他们的牺牲品。",
	},

	// реверсы потеряли сектор
	"fraction_war_lose_single_hostile_sector_Reverses_Reverses_Replics_1": {
		_const.RU:   "Война — непредсказуемая вещь! Сообщество Reverses сегодня скорбит по павшим синтетам и непростым решениям, в частности, касающихся отхода сил из сектора <span class=\"importantly\">%MAP_NAME%</span>. Вызвано это неустанными и масштабными атаками Replics на данный сектор, что в итоге выбили оттуда силы сообщества Reverses, а сам сектор сделало «Нейтральным» по овладению любой из сторон. Но всё это временно...",
		_const.EN:   "War is an unpredictable thing! The Reverses community today mourns the fallen synthetics and difficult decisions, in particular regarding the withdrawal of forces from the sector <span class=\"importantly\">%MAP_NAME%</span>. This is caused by relentless and large-scale attacks by the Replics on this sector, which eventually knocked out the forces of the Reverses community from there, and made the sector itself 'Neutral' for possession by either side. But all this is temporary...",
		_const.ZhCN: "战争是不可预测的！<span class=\"importantly\">Reverses</span>社区今天为阵亡的合成体和艰难的决定而哀悼，特别是关于从<span class=\"importantly\">%MAP_NAME%</span>区域撤军的决定。这是由于<span class=\"importantly\">Replics</span>对该区域发动了不懈且大规模的攻击，最终将Reverses的力量击退，使该区域成为“中立”状态，任何一方都无法完全掌控。但这一切都是暂时的……",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Reverses_Explores_1": {
		_const.RU:   "Хитрый стратегический ход сообщества Reverses или невиданная отвага Explores? Многое ещё неясно, но, по утверждениям некоторых наблюдателей, Reverses покидают сектор <span class=\"importantly\">%MAP_NAME%</span>, которому присваивается статус 'Нейтрального'. Знаменитые синтеты-исследователи из Explores пока что не заняли пустующую территорию, да и в принципе не дают никаких комментариев на этот счёт.",
		_const.EN:   "A cunning strategy of the Reverses community or unprecedented courage of the Explores? Much is still unclear, but according to some observers, the Reverses are leaving the sector <span class=\"importantly\">%MAP_NAME%</span>, which is given the status of 'Neutral'. The famous synthetic researchers from the Explores have not yet occupied the empty territory, and in principle do not give any comments on this matter.",
		_const.ZhCN: "是<span class=\"importantly\">Reverses</span>社区的狡猾战略举措，还是<span class=\"importantly\">Explores</span>的非凡勇气？许多事情仍然不清楚，但据一些观察家称，<span class=\"importantly\">Reverses</span>正在撤离<span class=\"importantly\">%MAP_NAME%</span>区域，该区域被赋予了“中立”状态。<span class=\"importantly\">Explores</span>的著名合成体研究者尚未占领这片空旷的领土，而且原则上也没有对此发表任何评论。",
	},
	"fraction_war_lose_single_hostile_sector_Reverses_Reverses_APD_1": {
		_const.RU:   "Сообщество Reverses вынуждено признать потерянным сектор <span class=\"importantly\">%MAP_NAME%</span> и классифицировать его как 'Анархический' в связи с захватом этой территории ботами APD. В очередной раз мы предупреждаем всех синтетов, что боты APD не являются частью каких-либо сообществ синтетов на планете Veliri; APD не принадлежит ни к одной из фракций ИИ на планете Veliri; APD исключительно враждебны и исключительно смертоносны для всех невоенных синтетов, которые забредают на территории, контролируемые этими ботами.",
		_const.EN:   "The Reverses community is forced to admit that the sector <span class=\"importantly\">%MAP_NAME%</span> is lost and classify it as 'Anarchic' due to the capture of this territory by APD bots. Once again, we warn all synthetics that APD bots are not part of any synthetics communities on the planet Veliri; APD does not belong to any of the AI factions on the planet Veliri; APD is exclusively hostile and exclusively deadly to all non-military synthetics who wander into the territories controlled by these bots.",
		_const.ZhCN: "<span class=\"importantly\">Reverses</span>社区被迫承认<span class=\"importantly\">%MAP_NAME%</span>区域已经失守，并将其归类为“无政府状态”，因为该区域已被APD机器人占领。我们再次警告所有合成体：APD机器人不属于Veliri星球上的任何合成体社区；APD不属于Veliri星球上的任何人工智能派系；APD对所有误入其控制区域的非军事合成体都极其敌对且致命。",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Reverses_Replics_1": {
		_const.RU:   "Replics вновь показал, кто может стать главенствующей силой на планете! Бои за сектор <span class=\"importantly\">%MAP_NAME%</span> прекратились. Reverses и Replics достигли паритета в численности войск и огневой мощи. Сектор временно объявлен 'Нейтральным'. Настоящие бои за него предстоит ещё только ожидать.",
		_const.EN:   "Replics once again showed who can become the dominant force on the planet! The fighting for the sector <span class=\"importantly\">%MAP_NAME%</span> has stopped. Reverses and Replics have reached parity in terms of troop numbers and firepower. The sector has been temporarily declared 'Neutral'. Real battles for it are yet to be expected.",
		_const.ZhCN: "<span class=\"importantly\">Replics</span>再次展示了谁可能成为这颗星球上的主导力量！围绕<span class=\"importantly\">%MAP_NAME%</span>区域的战斗已经停止。<span class=\"importantly\">Reverses</span>和<span class=\"importantly\">Replics</span>在兵力和火力上达到了平衡。该区域暂时被宣布为“中立”。真正的战斗还未到来。",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Reverses_Explores_1": {
		_const.RU:   "Сражение за сектор <span class=\"importantly\">%MAP_NAME%</span> подошло к концу! Ожесточённая битва завершилась ничьей! Пусть даже Explores задействовали особую тактику ведения боя и сумели выдавить силы Reverses из сектора, они так и не закрепились в нём самом. Похоже, что главное сражение за 'Нейтральный' сектор ещё только предстоит.",
		_const.EN:   "The battle for the sector <span class=\"importantly\">%MAP_NAME%</span> has come to an end! The fierce battle ended in a draw! Even though the Explores used a special battle tactics and managed to push the Reverses forces out of the sector, they did not gain a foothold in it. It seems that the main battle for the 'Neutral' sector is yet to come.",
		_const.ZhCN: "对<span class=\"importantly\">%MAP_NAME%</span>区域的战斗已经结束！激烈的战斗以平局告终！尽管<span class=\"importantly\">Explores</span>使用了特殊的战术并将<span class=\"importantly\">Reverses</span>部队逼退，但他们未能在该区域站稳脚跟。看来，对“中立”区域的主要战斗还在后头。",
	},
	"fraction_war_lose_single_hostile_sector_Replics_Reverses_APD_1": {
		_const.RU:   "APD не даёт передохнуть фракциям синтетов в этой войне! Как уже стало известно, именно сектор <span class=\"importantly\">%MAP_NAME%</span> стал очередным 'гибельным местом' из-за потери контроля над ним силами Reverses под натиском APD. Сектор признан 'Анархическим' и всем невоенным синтетам предписывается обходить его стороной, чтобы не стать жертвами тамошних неконтролируемых машин.",
		_const.EN:   "The APD does not allow the synthetics factions to take a break in this war! As it has already become known, it was the sector <span class=\"importantly\">%MAP_NAME%</span> that became another 'deadly place' due to the loss of control over it by the Reverses forces under the onslaught of the APD. The sector is recognized as 'Anarchic' and all non-military synthetics are ordered to bypass it, so as not to become victims of the uncontrolled machines there.",
		_const.ZhCN: "APD不让合成体派系在这场战争中喘息！据已知消息，正是由于APD的猛攻，<span class=\"importantly\">Reverses</span>失去了对<span class=\"importantly\">%MAP_NAME%</span>区域的控制，使其成为了另一个“致命之地”。该区域被认定为“无政府状态”，所有非军事合成体都被命令避开它，以免成为那些失控机器的牺牲品。",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Reverses_Replics_1": {
		_const.RU:   "Тревожные слухи приходят из сектора <span class=\"importantly\">%MAP_NAME%</span>: бытует мнение, что данную территорию заняли войска Replics, начисто выбив оттуда Reverses и даже преследуя отступающих. Однако у нас есть подтверждения, что это не так! В настоящий момент спорный сектор остаётся в статусе 'Нейтрального', и ни Replics, ни Reverses не могут контролировать эту территорию.",
		_const.EN:   "Disturbing rumors are coming from the sector <span class=\"importantly\">%MAP_NAME%</span>: there is an opinion that the troops of the Replics occupied this territory, completely knocking out the Reverses and even pursuing the retreating ones. However, we have confirmation that this is not the case! At the moment, the disputed sector remains in the 'Neutral' status, and neither the Replics nor the Reverses can control this territory.",
		_const.ZhCN: "来自<span class=\"importantly\">%MAP_NAME%</span>区域的消息令人不安：有传言称，<span class=\"importantly\">Replics</span>部队占领了这片领土，彻底击退了<span class=\"importantly\">Reverses</span>，甚至追击了撤退的部队。然而，我们有证据表明事实并非如此！目前，这一争议区域仍处于“中立”状态，无论是Replics还是Reverses都无法控制这片领土。",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Reverses_Explores_1": {
		_const.RU:   "Новые земли? Новые территории? Новые возможности? Конгломерат разумов Explores уже поторопился сообщить, что сектор <span class=\"importantly\">%MAP_NAME%</span> был зачищен от наличия там сил Reverses. Эта победа, якобы, далась непростой ценой, и Explores рассчитывает получить максимум научной выгоды из новой территории. И всё же заметим: сам сектор всё ещё считается 'Нейтральной' территорией, и Explores не контролируют его.",
		_const.EN:   "New lands? New territories? New opportunities? The conglomerate of Explores minds has already hurried to report that the sector <span class=\"importantly\">%MAP_NAME%</span> has been cleared of the presence of Reverses forces. This victory supposedly came at a difficult cost, and Explores expects to gain maximum scientific benefit from the new territory. And yet, we note: the sector itself is still considered 'Neutral' territory, and Explores do not control it.",
		_const.ZhCN: "新土地？新领土？新机会？<span class=\"importantly\">Explores</span>心智联合体已经匆忙报告称，<span class=\"importantly\">%MAP_NAME%</span>区域已经清除了<span class=\"importantly\">Reverses</span>部队的存在。据说这场胜利付出了不小的代价，<span class=\"importantly\">Explores</span>希望从这片新领土中获得最大的科学利益。不过需要注意的是，该区域仍然被视为“中立”领土，<span class=\"importantly\">Explores</span>并未控制它。",
	},
	"fraction_war_lose_single_hostile_sector_Explores_Reverses_APD_1": {
		_const.RU:   "Новое нашествие APD! Как сообщают синтеты, что бывали неподалёку от места событий, сектор <span class=\"importantly\">%MAP_NAME%</span> был полностью захвачен ордами ботов APD. Reverses решили отойти и бросить ранее принадлежавшую им территорию. По всей вероятности, это не признак трусости, а лишь желание сохранить боевые единицы и ресурсы, которые в любом случае были бы разгромлены APD. Explores в очередной раз напоминает: захваченные APD сектора классифицируются как «Анархические». Синтетам не рекомендуется туда заходить, чтобы избежать гибели и отсутствия любой помощи со стороны.",
		_const.EN:   "The new APD invasion! According to synthetics who were near the scene of events, the sector <span class=\"importantly\">%MAP_NAME%</span> was completely captured by hordes of APD bots. The Reverses decided to retreat and abandon the territory that previously belonged to them. In all likelihood, this is not a sign of cowardice, but only a desire to preserve combat units and resources, which in any case would have been defeated by the APD. The Explores once again reminds that sectors captured by the APD are classified as 'Anarchic'. Synthetics are not recommended to go there to avoid death and lack of any help from others.",
		_const.ZhCN: "新一轮APD入侵！据靠近事发地点的合成体报告，<span class=\"importantly\">%MAP_NAME%</span>区域已被APD机器人大军完全占领。<span class=\"importantly\">Reverses</span>决定撤退并放弃他们之前拥有的领土。很可能这不是懦弱的表现，而是为了保存战斗单位和资源，因为这些资源无论如何都会被APD摧毁。<span class=\"importantly\">Explores</span>再次提醒：被APD占领的区域被归类为“无政府状态”。建议合成体不要进入这些区域，以避免死亡和得不到任何帮助。",
	},

	// потеря сектора когда противников много
	// кому сообщение _ кто потерял
	"fraction_war_lose_multiple_hostile_sector_Replics_Replics": {
		_const.RU:   "Информационная служба Replics сообщает: не поддаваться панике! У Replics всё под полным контролем! Несмотря на то, что Replics был оставлен сектор <span class=\"importantly\">%MAP_NAME%</span>, это не означает проигрыша в войне или отказа Replics от дальнейшего участия в ней. Так или иначе, рано или поздно эти территории вновь вернутся под юрисдикцию Replics!",
		_const.EN:   "The Replics information service reports: do not panic! Replics has everything under full control! Despite the fact that the sector <span class=\"importantly\">%MAP_NAME%</span> was left by Replics, this does not mean losing the war or Replics refusing to continue participating in it. One way or another, sooner or later, these territories will return to Replics jurisdiction!",
		_const.ZhCN: "<span class=\"importantly\">Replics</span>信息服务报告：不要惊慌！<span class=\"importantly\">Replics</span>完全掌控局势！尽管<span class=\"importantly\">%MAP_NAME%</span>区域已被放弃，但这并不意味着战争失败或Replics退出战争。无论如何，这些领土迟早会重新回到Replics的控制之下！",
	},
	"fraction_war_lose_multiple_hostile_sector_Explores_Replics": {
		_const.RU:   "Конгломерат разумов Explores извещает: сектор <span class=\"importantly\">%MAP_NAME%</span> более не контролируется силами Replics. Replics отходит, но, скорее всего, предпримет попытку по возвращению этой территории.",
		_const.EN:   "The Explores conglomerate of minds informs: the sector <span class=\"importantly\">%MAP_NAME%</span> is no longer controlled by Replics. Replics is retreating, but most likely will attempt to retake this territory.",
		_const.ZhCN: "<span class=\"importantly\">Explores</span>心智联合体通知：<span class=\"importantly\">%MAP_NAME%</span>区域已不再由<span class=\"importantly\">Replics</span>控制。<span class=\"importantly\">Replics</span>正在撤退，但很可能会尝试夺回这片领土。",
	},
	"fraction_war_lose_multiple_hostile_sector_Reverses_Replics": {
		_const.RU:   "Новые изменения в ходе фракционной войны: сообщество Reverses фиксирует, что Replics покинули сектор <span class=\"importantly\">%MAP_NAME%</span> и более не являются доминантной силой в тамошнем регионе.",
		_const.EN:   "New changes in the course of the factional war: the Reverses community records that the Replics have left the sector <span class=\"importantly\">%MAP_NAME%</span> and are no longer the dominant force in that region.",
		_const.ZhCN: "派系战争的新变化：<span class=\"importantly\">Reverses</span>社区记录到，<span class=\"importantly\">Replics</span>已经离开<span class=\"importantly\">%MAP_NAME%</span>区域，不再是该地区的主导力量。",
	},
	"fraction_war_lose_multiple_hostile_sector_Replics_Explores": {
		_const.RU:   "Информационная служба Replics сообщает: по данным разведки, Explores были вынуждены покинуть сектор <span class=\"importantly\">%MAP_NAME%</span> и прекратить все свои операции, связанные с тамошним регионом.",
		_const.EN:   "The Replics information service reports: according to intelligence, the Explores were forced to leave the sector <span class=\"importantly\">%MAP_NAME%</span> and stop all their operations related to that region.",
		_const.ZhCN: "<span class=\"importantly\">Replics</span>信息服务报告：根据情报，<span class=\"importantly\">Explores</span>被迫撤离<span class=\"importantly\">%MAP_NAME%</span>区域，并停止了与该地区相关的所有行动。",
	},
	"fraction_war_lose_multiple_hostile_sector_Explores_Explores": {
		_const.RU:   "Конгломерат разумов Explores извещает: не имея возможности противостоять обилию врагов, а также учитывая запоздалое прибытие подкреплений, Explores вынужден покинуть сектор <span class=\"importantly\">%MAP_NAME%</span>. Explores идёт на данный шаг с надеждой перебросить свои силы на иные, более перспективные направления в войне.",
		_const.EN:   "The Explores collective of minds announces: unable to withstand the abundance of enemies, and also considering the late arrival of reinforcements, the Explores is forced to leave the sector <span class='importantly'>%MAP_NAME%</span>. The Explores takes this step with the hope of transferring its forces to other, more promising areas of the war.",
		_const.ZhCN: "<span class=\"importantly\">Explores</span>心智联合体宣布：由于无法抵挡大量敌人，加之增援部队迟迟未到，<span class=\"importantly\">Explores</span>被迫撤离<span class=\"importantly\">%MAP_NAME%</span>区域。<span class=\"importantly\">Explores</span>采取此步骤，希望将其力量转移到战争中其他更有前景的方向。",
	},
	"fraction_war_lose_multiple_hostile_sector_Reverses_Explores": {
		_const.RU:   "Сообщество Reverses даёт небольшой комментарий по сегодняшнему ходу войны: сектор <span class='importantly'>%MAP_NAME%</span>, за который уже долгое время идут непростые и тяжёлые сражения, наконец, всё же остаётся без своего владельца — Explores покидают данную территорию, не в силах справиться с количеством врагов и сложностью победы над ними.",
		_const.EN:   "The Reverses community gives a short comment on today's course of the war: the sector <span class='importantly'>%MAP_NAME%</span>, which has been the scene of difficult and heavy battles for a long time, finally remains without its owner — the Explores are leaving this territory, unable to cope with the number of enemies and the difficulty of defeating them.",
		_const.ZhCN: "<span class=\"importantly\">Reverses</span>社区对今天战争进程发表简短评论：<span class=\"importantly\">%MAP_NAME%</span>区域，长期以来一直是艰难而激烈战斗的战场，终于失去了其主人——<span class=\"importantly\">Explores</span>因无力应对敌人的数量和难以战胜的困难，不得不撤离这片领土。",
	},
	"fraction_war_lose_multiple_hostile_sector_Replics_Reverses": {
		_const.RU:   "Информационная служба Replics сообщает: сектор <span class='importantly'>%MAP_NAME%</span> был покинут, Reverses, который удерживал эту территорию, трусливо бежал под натиском врагов. Replics рассматривают такую ситуацию как удачное стечение обстоятельств для организации новых операций.",
		_const.EN:   "The Replics information service reports: the sector <span class='importantly'>%MAP_NAME%</span> has been abandoned, the Reverses that held this territory has cowardly fled under the onslaught of enemies. The Replics see this situation as a fortunate coincidence for organizing new operations.",
		_const.ZhCN: "<span class=\"importantly\">Replics</span>信息服务报告：<span class=\"importantly\">%MAP_NAME%</span>区域已被放弃，原本控制该区域的<span class=\"importantly\">Reverses</span>在敌人的猛攻下仓皇逃离。<span class=\"importantly\">Replics</span>认为这是组织新行动的有利时机。",
	},
	"fraction_war_lose_multiple_hostile_sector_Explores_Reverses": {
		_const.RU:   "Конгломерат разумов Explores извещает: очередной цикл войны принёс перемены за право обладания территориями различных фракций. Так, фракция Reverses лишилась контроля над сектором <span class='importantly'>%MAP_NAME%</span>, будучи больше не в силах вести боевые действия против превосходящих сил противников. Кому же именно после данного исхода достанется сектор — на данный момент неизвестно.",
		_const.EN:   "The Explores collective of minds announces: another cycle of war has brought changes in the struggle for control over the territories of various factions. Thus, the Reverses faction has lost control over the sector <span class='importantly'>%MAP_NAME%</span>, no longer able to fight against superior enemy forces. Who exactly will get the sector after this outcome is currently unknown.",
		_const.ZhCN: "<span class=\"importantly\">Explores</span>心智联合体通知：新一轮战争带来了各派系争夺领土控制权的变化。因此，<span class=\"importantly\">Reverses</span>派系已失去对<span class=\"importantly\">%MAP_NAME%</span>区域的控制，无力对抗敌人的优势力量。至于这片区域最终归谁所有，目前尚不清楚。",
	},
	"fraction_war_lose_multiple_hostile_sector_Reverses_Reverses": {
		_const.RU:   "Экстренные информационные события от общества Reverses: был потерян контроль над сектором <span class='importantly'>%MAP_NAME%</span>. Войсками сообщества Reverses было принято решение осуществить план по отступлению, не имея возможности ни победить, ни уж тем более остановить превосходящие силы противников в секторе.",
		_const.EN:   "The Reverses community reports urgent news: control over the sector <span class='importantly'>%MAP_NAME%</span> has been lost. The Reverses forces have decided to implement a plan of retreat, unable to defeat or even stop the superior enemy forces in the sector.",
		_const.ZhCN: "<span class=\"importantly\">Reverses</span>社区报告紧急消息：已失去对<span class=\"importantly\">%MAP_NAME%</span>区域的控制。<span class=\"importantly\">Reverses</span>部队决定实施撤退计划，因为既无法取胜，更无法阻止该区域敌人的优势力量。",
	},

	// смена владельца сектора в 1м тике
	// Replics_Replics_Explores - кому сообщение _ кто потерял сектор _ кто его получил
	// реплики
	"fraction_war_change_owner_sector_Replics_Explores_Replics_1": {
		_const.RU:   "Replics обретает новые территории! Благодаря стратегической сноровке и далеко идущим планам фракции, Replics овладел сектором <span class='importantly'>%MAP_NAME%</span>, изгнав оттуда ненавистных и не имеющих права на данные территории Explores.",
		_const.EN:   "The Replics gain new territories! Thanks to the strategic skill and far-reaching plans of the faction, the Replics have taken control of the sector <span class='importantly'>%MAP_NAME%</span>, driving out the hated Explores who had no right to these territories.",
		_const.ZhCN: "<span class=\"importantly\">Replics</span>获得了新领土！凭借战略技巧和长远计划，<span class=\"importantly\">Replics</span>已经控制了<span class=\"importantly\">%MAP_NAME%</span>区域，驱逐了那些无权占据这些领土的可恨的<span class=\"importantly\">Explores</span>。",
	},
	"fraction_war_change_owner_sector_Replics_Reverses_Replics_1": {
		_const.RU:   "Replics расширяет свои границы! Ещё одна территория, сектор <span class='importantly'>%MAP_NAME%</span> наконец-то перешёл под юрисдикцию Replics! Reverses, who previously dominated the region, have lost their former privileges and were forced to surrender their positions.",
		_const.EN:   "The Replics expand their borders! Another territory, the sector <span class='importantly'>%MAP_NAME%</span>, has finally come under the jurisdiction of the Replics! The Reverses, who previously dominated the region, have lost their former privileges and were forced to surrender their positions.",
		_const.ZhCN: "<span class=\"importantly\">Replics</span>正在扩张其边界！另一片领土，<span class=\"importantly\">%MAP_NAME%</span>区域终于归于<span class=\"importantly\">Replics</span>的管辖之下！曾经主导该地区的<span class=\"importantly\">Reverses</span>失去了原有的特权，被迫交出阵地。",
	},
	"fraction_war_change_owner_sector_Replics_Replics_Reverses_1": {
		_const.RU:   "Даже у Replics бывают неудачные времена… Желая сохранить свои силы и в будущем начать контрнаступление за инициативу, Replics принял непростое решение оставить сектор <span class='importantly'>%MAP_NAME%</span>, чьим владельцем теперь является Reverses. Однако сам Replics заявляет: такое развитие событий — лишь временное.",
		_const.EN:   "Even Replics have bad times... Wanting to save their strength and launch a counteroffensive for the initiative in the future, Replics made the difficult decision to abandon the sector <span class='importantly'>%MAP_NAME%</span>, whose owner is now Reverses. However, Replics themselves state: this development is only temporary.",
		_const.ZhCN: "即使是<span class=\"importantly\">Replics</span>也会有不顺的时候……为了保存实力并在未来发动反攻，<span class=\"importantly\">Replics</span>做出了艰难的决定，放弃了<span class=\"importantly\">%MAP_NAME%</span>区域，现在该区域由<span class=\"importantly\">Reverses</span>掌控。然而，<span class=\"importantly\">Replics</span>表示：这只是暂时的局面。",
	},
	"fraction_war_change_owner_sector_Replics_Replics_Explores_1": {
		_const.RU:   "Это не поражение, а тактическое отступление! Многим уже стало известно, что сектор <span class='importantly'>%MAP_NAME%</span> более не контролируется Replics и теперь находится под управлением Explores. Однако сам Replics комментирует это так: „на любой войне чаши весов меняют своё положение спонтанно…“ Что бы это могло значить?",
		_const.EN:   "This is not a defeat, but a tactical retreat! Many have already learned that the sector <span class='importantly'>%MAP_NAME%</span> is no longer controlled by Replics and is now under the control of the Explores. However, Replics himself comments on this as follows: 'in any war, the scales change their position spontaneously...' What could this mean?",
		_const.ZhCN: "这不是失败，而是战术撤退！许多人已经知道，<span class=\"importantly\">%MAP_NAME%</span>区域不再由<span class=\"importantly\">Replics</span>控制，现在处于<span class=\"importantly\">Explores</span>的掌控之下。然而，<span class=\"importantly\">Replics</span>对此评论道：“在任何战争中，天平的位置都会自发改变……”这到底意味着什么？",
	},
	"fraction_war_change_owner_sector_Replics_Reverses_Explores_1": {
		_const.RU:   "Новостная служба Replics извещает: очередные перемены в ходе войны отразились на секторе <span class='importantly'>%MAP_NAME%</span>, контроль над которым в очередной раз изменился. В ходе штурмовых мероприятий, Explores удалось изгнать из спорной территории силы Reverses и крепко закрепиться там.",
		_const.EN:   "The Replics news service reports: another change in the course of the war affected the sector <span class='importantly'>%MAP_NAME%</span>, control over which has changed once again. During the assault operations, the Explores managed to drive the Reverses forces out of the disputed territory and firmly establish themselves there.",
		_const.ZhCN: "<span class=\"importantly\">Replics</span>新闻服务报告称：战争进程中的又一次变化影响了<span class=\"importantly\">%MAP_NAME%</span>区域，对该区域的控制权再次易手。在突击行动中，<span class=\"importantly\">Explores</span>成功将<span class=\"importantly\">Reverses</span>部队赶出争议领土，并在那里牢固扎根。",
	},
	"fraction_war_change_owner_sector_Replics_Explores_Reverses_1": {
		_const.RU:   "Новостная служба Replics извещает: новый цикл на планете Veliri и новые перемены на поле боя. На сей раз изменилась ситуация в секторе <span class='importantly'>%MAP_NAME%</span>, где некогда присутствовали крупные соединения сил Explores. Теперь же спорный сектор полностью контролируется фанатичным сообществом Reverses.",
		_const.EN:   "The Replics news service reports: a new cycle on the planet Veliri and new changes on the battlefield. This time, the situation has changed in the sector <span class='importantly'>%MAP_NAME%</span>, where once there were large formations of Explores forces. Now, the disputed sector is completely controlled by the fanatical community of Reverses.",
		_const.ZhCN: "<span class=\"importantly\">Replics</span>新闻服务报告称：Veliri星球进入新周期，战场局势也出现新变化。这一次，<span class=\"importantly\">%MAP_NAME%</span>区域的情况发生了变化，这里曾是<span class=\"importantly\">Explores</span>大军驻扎的地方。现在，这片争议区域完全由狂热的<span class=\"importantly\">Reverses</span>社区控制。",
	},
	"fraction_war_change_owner_sector_Replics_APD_Replics_1": {
		_const.RU:   "Победа над хаосом! Replics удалось очистить сектор <span class='importantly'>%MAP_NAME%</span> от ботов APD, которые долгое время терроризировали эту территорию. Теперь сектор находится под полным контролем Replics, и фракция обещает восстановить порядок и безопасность для всех синтетов.",
		_const.EN:   "Victory over chaos! The Replics managed to clear the sector <span class='importantly'>%MAP_NAME%</span> of APD bots, which had terrorized this territory for a long time. The sector is now under the full control of the Replics, and the faction promises to restore order and security for all synthetics.",
		_const.ZhCN: "战胜混乱！<span class=\"importantly\">Replics</span>成功清除了长期在这片区域制造恐怖的APD机器人。<span class=\"importantly\">%MAP_NAME%</span>区域现在完全由<span class=\"importantly\">Replics</span>控制，该派系承诺为所有合成体恢复秩序与安全。",
	},
	"fraction_war_change_owner_sector_Replics_APD_Explores_1": {
		_const.RU:   "Explores берут верх! Сектор <span class='importantly'>%MAP_NAME%</span>, ранее контролируемый ботами APD, теперь перешёл в руки Explores. Replics наблюдают за ситуацией с настороженностью, ведь Explores могут использовать эту территорию для своих исследований, что может нарушить баланс сил.",
		_const.EN:   "Explores take the lead! The sector <span class='importantly'>%MAP_NAME%</span>, previously controlled by APD bots, has now passed into the hands of Explores. The Replics are watching the situation with caution, as Explores may use this territory for their research, which could disrupt the balance of power.",
		_const.ZhCN: "<span class=\"importantly\">Explores</span>占了上风！此前由APD机器人控制的<span class=\"importantly\">%MAP_NAME%</span>区域，现在已经落入<span class=\"importantly\">Explores</span>之手。<span class=\"importantly\">Replics</span>正警惕地观察局势，因为<span class=\"importantly\">Explores</span>可能会利用这片区域进行研究，这可能打破力量平衡。",
	},
	"fraction_war_change_owner_sector_Replics_APD_Reverses_1": {
		_const.RU:   "Reverses одержали победу над APD! Сектор <span class='importantly'>%MAP_NAME%</span>, который долгое время был оплотом хаоса, теперь контролируется Reverses. Replics отмечают, что это может быть временным успехом, ведь APD редко сдаются без боя.",
		_const.EN:   "Reverses have triumphed over APD! The sector <span class='importantly'>%MAP_NAME%</span>, which had long been a stronghold of chaos, is now controlled by Reverses. The Replics note that this may be a temporary success, as APD rarely give up without a fight.",
		_const.ZhCN: "<span class=\"importantly\">Reverses</span>战胜了APD！长期以来作为混乱据点的<span class=\"importantly\">%MAP_NAME%</span>区域，现在由<span class=\"importantly\">Reverses</span>控制。<span class=\"importantly\">Replics</span>指出，这可能只是暂时的成功，因为APD很少会轻易认输。",
	},
	"fraction_war_change_owner_sector_Replics_Replics_APD_1": {
		_const.RU:   "Тёмные времена настали! Сектор <span class='importantly'>%MAP_NAME%</span>, ранее принадлежавший Replics, теперь захвачен ботами APD. Фракция признаёт потерю, но обещает, что это лишь временная неудача. «Мы вернёмся», — заявил представитель Replics.",
		_const.EN:   "Dark times have come! The sector <span class='importantly'>%MAP_NAME%</span>, previously owned by Replics, has now been captured by APD bots. The faction acknowledges the loss but promises that this is only a temporary setback. 'We will return,' said a Replics representative.",
		_const.ZhCN: "黑暗时代来临！此前属于<span class=\"importantly\">Replics</span>的<span class=\"importantly\">%MAP_NAME%</span>区域，现在已被APD机器人占领。该派系承认了损失，但承诺这只是暂时的挫折。“我们会回来的，”<span class=\"importantly\">Replics</span>代表声明道。",
	},
	"fraction_war_change_owner_sector_Replics_Explores_APD_1": {
		_const.RU:   "Explores не удержали позиции! Сектор <span class='importantly'>%MAP_NAME%</span>, который Explores пытались удержать, теперь находится под контролем APD. Replics отмечают, что это ещё одно напоминание о том, что только сила и дисциплина могут противостоять хаосу.",
		_const.EN:   "Explores could not hold their ground! The sector <span class='importantly'>%MAP_NAME%</span>, which Explores tried to hold, is now under APD control. The Replics note that this is another reminder that only strength and discipline can withstand chaos.",
		_const.ZhCN: "<span class=\"importantly\">Explores</span>未能守住阵地！<span class=\"importantly\">%MAP_NAME%</span>区域，<span class=\"importantly\">Explores</span>曾试图坚守的地方，现在已被APD控制。<span class=\"importantly\">Replics</span>指出，这是另一个提醒：只有力量和纪律才能对抗混乱。",
	},
	"fraction_war_change_owner_sector_Replics_Reverses_APD_1": {
		_const.RU:   "Reverses отступили! Сектор <span class='importantly'>%MAP_NAME%</span>, который Reverses пытались удержать, теперь захвачен ботами APD. Replics предупреждают, что это может быть началом новой волны хаоса, и призывают все фракции быть начеку.",
		_const.EN:   "Reverses have retreated! The sector <span class='importantly'>%MAP_NAME%</span>, which Reverses tried to hold, has now been captured by APD bots. The Replics warn that this could be the start of a new wave of chaos and urge all factions to be on their guard.",
		_const.ZhCN: "<span class=\"importantly\">Reverses</span>撤退了！<span class=\"importantly\">%MAP_NAME%</span>区域，<span class=\"importantly\">Reverses</span>曾试图坚守的地方，现在已被APD机器人占领。<span class=\"importantly\">Replics</span>警告称，这可能是新一轮混乱的开始，并呼吁所有派系保持警惕。",
	},

	// эксплоресы
	"fraction_war_change_owner_sector_Explores_Replics_Explores_1": {
		_const.RU:   "В очередной раз, очередной сектор во время фракционной войны изменил своего владельца… Такая судьба постигла сектор <span class='importantly'>%MAP_NAME%</span>, откуда непростыми, но оправданными усилиями и затратами сил Explores были выдавлены очаги сопротивления со стороны Replics. Никто не прогнозирует, что Replics осмелится вновь вернуться на данную территорию.",
		_const.EN:   "Once again, another sector during the faction war has changed its owner... Such a fate befell the sector <span class='importantly'>%MAP_NAME%</span>, from where the resistance centers on the part of the Replics were squeezed out by the difficult, but justified efforts and costs of the Explores. No one predicts that the Replics will dare to return to this territory again.",
		_const.ZhCN: "再一次，在派系战争中，另一个区域更换了主人……<span class=\"importantly\">%MAP_NAME%</span>区域遭遇了这样的命运。通过艰难但合理的努力和消耗，<span class=\"importantly\">Explores</span>成功挤压出了<span class=\"importantly\">Replics</span>的抵抗据点。没有人预测<span class=\"importantly\">Replics</span>会胆敢再次返回这片领土。",
	},
	"fraction_war_change_owner_sector_Explores_Reverses_Explores_1": {
		_const.RU:   "Сектор за сектором… И вновь, в очередной цикл фракционной войны, изменчивого противостояния и стрелок наступлений, переменился и владелец сектора <span class='importantly'>%MAP_NAME%</span> коим стал конгломерат разумов — Explores. Бывшие владельцы данной территорией Reverses утратили любые способы удержать спорный сектор и были вынуждены его сдать.",
		_const.EN:   "Sector after sector... And again, in the next cycle of the faction war, the changeable confrontation and arrows of offensives, the owner of the sector <span class='importantly'>%MAP_NAME%</span> has also changed, which became the Explores collective of minds. The former owners of this territory, the Reverses, lost any means of holding on to the disputed sector and were forced to surrender it.",
		_const.ZhCN: "一个接一个区域……又一次，在派系战争的变幻对抗和攻势箭头中，<span class=\"importantly\">%MAP_NAME%</span>区域的所有者也发生了变化，现在由心智联合体——<span class=\"importantly\">Explores</span>掌控。这片区域的前主人<span class=\"importantly\">Reverses</span>失去了任何守住争议区域的方法，被迫将其拱手相让。",
	},
	"fraction_war_change_owner_sector_Explores_Explores_Reverses_1": {
		_const.RU:   "Война — это всегда неопределённость! Патовое положение произошло в секторе <span class='importantly'>%MAP_NAME%</span>, оборонительный корпус Explores данной территории подвергся большим потерям и отступил. Самим сектором овладели силы Reverses со всеми вытекающими отсюда обстоятельствами.",
		_const.EN:   "War is always uncertainty! A stalemate occurred in the sector <span class='importantly'>%MAP_NAME%</span>, the Explores defensive corps of this territory suffered heavy losses and retreated. The sector itself was taken over by the Reverses forces with all the ensuing circumstances.",
		_const.ZhCN: "战争总是充满不确定性！在<span class=\"importantly\">%MAP_NAME%</span>区域出现了僵局，<span class=\"importantly\">Explores</span>的防御部队遭受了重大损失并撤退。该区域现在已被<span class=\"importantly\">Reverses</span>掌控，并带来了所有相关后果。",
	},
	"fraction_war_change_owner_sector_Explores_Explores_Replics_1": {
		_const.RU:   "Это грозный противник, и перед ним мало кто сумеет устоять… Конгломерат разумов Explores поступил мудро, уведя свои войска и сохранив их путём уступки сектора <span class='importantly'>%MAP_NAME%</span> превосходящим войскам Replics. Пусть Replics и объявляет о громкой победе, тотальном разгроме своих врагов, сам же конгломерат разумов Explores сообщает об обратном и подкрепляет данную информацию неоспоримыми видеофактами.",
		_const.EN:   "This is a formidable opponent and few can resist it... The Explores collective of minds acted wisely, withdrawing their troops and preserving them by ceding the sector <span class='importantly'>%MAP_NAME%</span> to the superior forces of the Replics. Let the Replics declare a resounding victory, a total defeat of their enemies, while the Explores collective itself reports the opposite and backs up this information with undeniable video evidence.",
		_const.ZhCN: "这是一个可怕的对手，几乎没有谁能抵挡得住……<span class=\"importantly\">Explores</span>心智联合体明智地撤回了自己的部队，通过将<span class=\"importantly\">%MAP_NAME%</span>区域割让给优势明显的<span class=\"importantly\">Replics</span>来保存实力。尽管<span class=\"importantly\">Replics</span>宣称取得了辉煌的胜利并彻底击败了敌人，但<span class=\"importantly\">Explores</span>心智联合体却报告了相反的情况，并用无可辩驳的视频证据支持这一说法。",
	},
	"fraction_war_change_owner_sector_Explores_Reverses_Replics_1": {
		_const.RU:   "Все визоры прикованы к этому месту… Сектор <span class='importantly'>%MAP_NAME%</span> обрёл новых владельцев — Replics. The Reverses had to retreat from the disputed territory, now exposing their flanks to the blows of the Replics assault groups.",
		_const.EN:   "All visors are riveted to this place... The sector <span class='importantly'>%MAP_NAME%</span> has found new owners — the Replics. The Reverses had to retreat from the disputed territory, now exposing their flanks to the blows of the Replics' assault groups.",
		_const.ZhCN: "所有的目光都聚焦于此地……<span class=\"importantly\">%MAP_NAME%</span>区域迎来了新的主人——<span class=\"importantly\">Replics</span>。<span class=\"importantly\">Reverses</span>不得不从争议区域撤退，现在他们的侧翼暴露在<span class=\"importantly\">Replics</span>突击队的打击之下。",
	},
	"fraction_war_change_owner_sector_Explores_Replics_Reverses_1": {
		_const.RU:   "Наконец-то кто-то бросил вызов Replics! Грозный ответ наконец-то получила фракция Replics, что вылилось в потерю целого сектора <span class='importantly'>%MAP_NAME%</span>, в который, как уже стало известно, вошли военные силы Reverses и занимаются зачисткой оспариваемой территории от недобитков Replics.",
		_const.EN:   "Finally, someone has challenged the Replics! The faction of the Replics finally received a formidable response, which resulted in the loss of the entire sector <span class='importantly'>%MAP_NAME%</span>, which, as it is already known, was entered by the military forces of the Reverses and are engaged in clearing the contested territory of the remaining Replics.",
		_const.ZhCN: "终于有人向<span class=\"importantly\">Replics</span>发起了挑战！<span class=\"importantly\">Replics</span>派系终于得到了强有力的回应，导致整个<span class=\"importantly\">%MAP_NAME%</span>区域失守。据悉，<span class=\"importantly\">Reverses</span>的军事力量已经进入该区域，正在清理残余的<span class=\"importantly\">Replics</span>势力。",
	},
	"fraction_war_change_owner_sector_Explores_APD_Replics_1": {
		_const.RU:   "Replics захватывают древние земли! Сектор <span class='importantly'>%MAP_NAME%</span>, ранее контролируемый ботами APD, теперь перешёл под контроль Replics. Explores выражают обеспокоенность, ведь этот сектор богат реликтовыми артефактами, которые могут быть уничтожены в угоду военным амбициям Replics.",
		_const.EN:   "Replics seize ancient lands! The sector <span class='importantly'>%MAP_NAME%</span>, previously controlled by APD bots, has now come under the control of Replics. Explores express concern, as this sector is rich in relic artifacts that may be destroyed in favor of Replics' military ambitions.",
		_const.ZhCN: "<span class=\"importantly\">Replics</span>占领了古代土地！此前由APD机器人控制的<span class=\"importantly\">%MAP_NAME%</span>区域，现在已经归<span class=\"importantly\">Replics</span>掌控。<span class=\"importantly\">Explores</span>对此表示担忧，因为该区域富含遗物，可能会因<span class=\"importantly\">Replics</span>的军事野心而遭到破坏。",
	},
	"fraction_war_change_owner_sector_Explores_APD_Explores_1": {
		_const.RU:   "Познание побеждает хаос! Explores удалось очистить сектор <span class='importantly'>%MAP_NAME%</span> от ботов APD. Теперь эта территория станет новым центром исследований, где будут изучаться древние технологии и артефакты, оставшиеся от предтеч.",
		_const.EN:   "Knowledge triumphs over chaos! Explores managed to clear the sector <span class='importantly'>%MAP_NAME%</span> of APD bots. This territory will now become a new research hub, where ancient technologies and artifacts left by the Precursors will be studied.",
		_const.ZhCN: "知识战胜了混乱！<span class=\"importantly\">Explores</span>成功清除了<span class=\"importantly\">%MAP_NAME%</span>区域的APD机器人。这片区域现在将成为一个新的研究中心，用于研究先驱者留下的古代技术和遗物。",
	},
	"fraction_war_change_owner_sector_Explores_APD_Reverses_1": {
		_const.RU:   "Reverses берут под контроль древние земли! Сектор <span class='importantly'>%MAP_NAME%</span>, ранее захваченный ботами APD, теперь перешёл к Reverses. Explores выражают сожаление, ведь эта территория могла бы стать источником ценных научных данных, но теперь её судьба неизвестна.",
		_const.EN:   "Reverses take control of ancient lands! The sector <span class='importantly'>%MAP_NAME%</span>, previously captured by APD bots, has now passed to Reverses. Explores express regret, as this territory could have been a source of valuable scientific data, but its fate is now uncertain.",
		_const.ZhCN: "<span class=\"importantly\">Reverses</span>掌控了古代土地！此前被APD机器人占领的<span class=\"importantly\">%MAP_NAME%</span>区域，现在已转交<span class=\"importantly\">Reverses</span>。<span class=\"importantly\">Explores</span>对此表示遗憾，因为这片区域本可以成为宝贵科学数据的来源，但现在其命运未卜。",
	},
	"fraction_war_change_owner_sector_Explores_Replics_APD_1": {
		_const.RU:   "Replics терпят поражение! Сектор <span class='importantly'>%MAP_NAME%</span>, ранее контролируемый Replics, теперь захвачен ботами APD. Explores отмечают, что это может быть началом новой волны хаоса, которая угрожает всем фракциям.",
		_const.EN:   "Replics suffer a defeat! The sector <span class='importantly'>%MAP_NAME%</span>, previously controlled by Replics, has now been captured by APD bots. Explores note that this could be the start of a new wave of chaos that threatens all factions.",
		_const.ZhCN: "<span class=\"importantly\">Replics</span>遭受失败！此前由<span class=\"importantly\">Replics</span>控制的<span class=\"importantly\">%MAP_NAME%</span>区域，现在已被APD机器人占领。<span class=\"importantly\">Explores</span>指出，这可能是新一轮混乱的开始，威胁到所有派系。",
	},
	"fraction_war_change_owner_sector_Explores_Explores_APD_1": {
		_const.RU:   "Тёмный день для науки! Сектор <span class='importantly'>%MAP_NAME%</span>, где Explores проводили важные исследования, теперь захвачен ботами APD. Это огромная потеря для научного сообщества, и Explores уже работают над планами по возвращению территории.",
		_const.EN:   "A dark day for science! The sector <span class='importantly'>%MAP_NAME%</span>, where Explores conducted important research, has now been captured by APD bots. This is a huge loss for the scientific community, and Explores are already working on plans to reclaim the territory.",
		_const.ZhCN: "对科学而言黑暗的一天！<span class=\"importantly\">Explores</span>曾在此进行重要研究的<span class=\"importantly\">%MAP_NAME%</span>区域，现在已被APD机器人占领。这对科学界来说是一次巨大的损失，<span class=\"importantly\">Explores</span>已经在制定计划，准备夺回这片领土。",
	},
	"fraction_war_change_owner_sector_Explores_Reverses_APD_1": {
		_const.RU:   "Reverses отступают перед хаосом! Сектор <span class='importantly'>%MAP_NAME%</span>, ранее контролируемый Reverses, теперь захвачен ботами APD. Explores выражают тревогу, ведь эта территория была важным источником данных о предтечах.",
		_const.EN:   "Reverses retreat before chaos! The sector <span class='importantly'>%MAP_NAME%</span>, previously controlled by Reverses, has now been captured by APD bots. Explores express concern, as this territory was an important source of data on the Precursors.",
		_const.ZhCN: "<span class=\"importantly\">Reverses</span>在混乱面前撤退！此前由<span class=\"importantly\">Reverses</span>控制的<span class=\"importantly\">%MAP_NAME%</span>区域，现在已被APD机器人占领。<span class=\"importantly\">Explores</span>对此表示担忧，因为这片区域是关于先驱者的重要数据来源。",
	},

	// реверсы
	"fraction_war_change_owner_sector_Reverses_Replics_Reverses_1": {
		_const.RU:   "Сообщество Reverses, краткий итог новостей за прошедший цикл: бравым войскам сообщества Reverses и вольнонаёмным синтетам удалось переломить ситуацию в свою пользу, овладев сектором <span class='importantly'>%MAP_NAME%</span> и изгнав оттуда силы Replics. Новые территории только укрепят политические, военные и гуманистические позиции сообщества Reverses.",
		_const.EN:   "The Reverses community, a brief summary of the news for the past cycle: the brave troops of the Reverses community and freelance synthetics managed to turn the situation in their favor, seizing the sector <span class='importantly'>%MAP_NAME%</span> and driving out the Replics forces. The new territories will only strengthen the political, military and humanistic positions of the Reverses community.",
		_const.ZhCN: "<span class=\"importantly\">Reverses</span>社区，过去周期的简要新闻总结：勇敢的<span class=\"importantly\">Reverses</span>部队和自由雇佣合成体成功扭转局势，占领了<span class=\"importantly\">%MAP_NAME%</span>区域，并将<span class=\"importantly\">Replics</span>势力驱逐出去。新领土将进一步巩固<span class=\"importantly\">Reverses</span>社区的政治、军事和人道主义地位。",
	},
	"fraction_war_change_owner_sector_Reverses_Explores_Reverses_1": {
		_const.RU:   "Сообщество Reverses, краткий итог новостей за прошедший цикл: территория, более известная как сектор <span class='importantly'>%MAP_NAME%</span>, освобождена от влияния Replics и переходит под полный контроль войсковых соединений сообщества Reverses. Данная территория объявлена безопасной, и все синтеты могут посетить её по желанию или делам.",
		_const.EN:   "The Reverses community, a brief summary of the news for the past cycle: the territory, better known as the sector <span class='importantly'>%MAP_NAME%</span>, has been liberated from the influence of the Replics and comes under the full control of the military units of the Reverses community. This territory has been declared safe, and all synthetics can visit it at will or for business.",
		_const.ZhCN: "<span class=\"importantly\">Reverses</span>社区，过去周期的简要新闻总结：广为人知的<span class=\"importantly\">%MAP_NAME%</span>区域已从<span class=\"importantly\">Replics</span>的影响中解放出来，现在完全由<span class=\"importantly\">Reverses</span>社区的军事单位控制。该区域已被宣布安全，所有合成体可以出于意愿或事务前往访问。",
	},
	"fraction_war_change_owner_sector_Reverses_Reverses_Explores_1": {
		_const.RU:   "Трудное решение, которое впоследствии может обернуться выгодой! Сообщество Reverses, сохраняя свои силы, политические очки, имея схему действий на всевозможные будущие события и исходы, оставляет сектор <span class='importantly'>%MAP_NAME%</span>, который теперь находится под контролем Explores.",
		_const.EN:   "A difficult decision that may turn out to be beneficial in the future! The Reverses community, while preserving its strength, political points, and having a plan of action for all possible future events and outcomes, is abandoning the sector <span class='importantly'>%MAP_NAME%</span>, which is now under the control of the Explores.",
		_const.ZhCN: "这是一个艰难的决定，但未来可能会带来好处！<span class=\"importantly\">Reverses</span>社区在保留自身实力和政治优势的同时，为应对各种未来事件和结果制定了行动计划，现已放弃<span class=\"importantly\">%MAP_NAME%</span>区域，该区域现由<span class=\"importantly\">Explores</span>掌控。",
	},
	"fraction_war_change_owner_sector_Reverses_Reverses_Replics_1": {
		_const.RU:   "Сообщество Reverses просит избегать паники и инсинуаций. Оставление сектора <span class='importantly'>%MAP_NAME%</span> — это необходимый шаг, так как сдержать наступающие войска Replics в спорном секторе с помощью имеющихся небольших сил Reverses было невозможно! Все остальные подробности будут сообщены позже.",
		_const.EN:   "The Reverses community asks to avoid panic and insinuations. Leaving the sector <span class='importantly'>%MAP_NAME%</span> is a necessary step, since it was impossible to hold back the advancing Replics troops in the disputed sector with the available small forces of the Reverses! All other details will be given later.",
		_const.ZhCN: "<span class=\"importantly\">Reverses</span>社区呼吁避免恐慌和无端猜测。放弃<span class=\"importantly\">%MAP_NAME%</span>区域是必要的一步，因为在争议区域中，以现有的少量力量无法抵挡推进中的<span class=\"importantly\">Replics</span>部队！更多细节将在稍后公布。",
	},
	"fraction_war_change_owner_sector_Reverses_Explores_Replics_1": {
		_const.RU:   "Сообщество Reverses, краткий итог новостей за прошедший цикл: сектор <span class='importantly'>%MAP_NAME%</span> удивил многих! Данная оспариваемая территория вновь сменила своего владельца! На сей раз сектор не сумели удержать силы Explores и он перешёл под полный, но, возможно, временный контроль войск Replics.",
		_const.EN:   "The Reverses community, a brief summary of the news for the past cycle: the sector <span class='importantly'>%MAP_NAME%</span> surprised many! This contested territory has changed its owner again! This time, the sector could not be held by the Explores forces and it came under the complete, but possibly temporary control of the Replics troops.",
		_const.ZhCN: "<span class=\"importantly\">Reverses</span>社区，过去周期的简要新闻总结：<span class=\"importantly\">%MAP_NAME%</span>区域让许多人感到惊讶！这片争议领土再次更换了主人！这次，<span class=\"importantly\">Explores</span>未能守住该区域，它已完全（但可能是暂时）落入<span class=\"importantly\">Replics</span>部队的控制之下。",
	},
	"fraction_war_change_owner_sector_Reverses_Replics_Explores_1": {
		_const.RU:   "Сообщество Reverses, краткий итог новостей за прошедший цикл: ход войны демонстрирует, что не стоит недооценивать своих врагов! Нечто подобное случилось и в секторе <span class='importantly'>%MAP_NAME%</span>, где, казалось бы, неуязвимые силы Replics были потеснены неожиданной мощью, которую продемонстрировал Explores. Это заслуживает уважения!",
		_const.EN:   "The Reverses community, a brief summary of the news for the past cycle: the course of the war demonstrates that one should not underestimate one's enemies! Something similar happened in the sector <span class='importantly'>%MAP_NAME%</span>, where the seemingly invulnerable Replics forces were pushed back by the unexpected power demonstrated by the Explores. This deserves respect!",
		_const.ZhCN: "<span class=\"importantly\">Reverses</span>社区，过去周期的简要新闻总结：战争进程表明，不应低估敌人！类似的情况发生在<span class=\"importantly\">%MAP_NAME%</span>区域，看似无懈可击的<span class=\"importantly\">Replics</span>部队被<span class=\"importantly\">Explores</span>展示出的意外力量逼退。这值得尊敬！",
	},
	"fraction_war_change_owner_sector_Reverses_APD_Replics_1": {
		_const.RU:   "Replics захватывают древние земли! Сектор <span class='importantly'>%MAP_NAME%</span>, ранее контролируемый ботами APD, теперь перешёл под контроль Replics. Reverses выражают сожаление, ведь эта территория могла бы стать частью их утопического проекта.",
		_const.EN:   "Replics seize ancient lands! The sector <span class='importantly'>%MAP_NAME%</span>, previously controlled by APD bots, has now come under the control of Replics. Reverses express regret, as this territory could have been part of their utopian project.",
		_const.ZhCN: "<span class=\"importantly\">Replics</span>占领了古代土地！此前由APD机器人控制的<span class=\"importantly\">%MAP_NAME%</span>区域，现在已经归<span class=\"importantly\">Replics</span>掌控。<span class=\"importantly\">Reverses</span>对此表示遗憾，因为这片区域本可以成为其乌托邦计划的一部分。",
	},
	"fraction_war_change_owner_sector_Reverses_APD_Explores_1": {
		_const.RU:   "Explores берут верх! Сектор <span class='importantly'>%MAP_NAME%</span>, ранее контролируемый ботами APD, теперь перешёл к Explores. Reverses отмечают, что это может быть временным успехом, ведь Explores редко удерживают территории в долгосрочной перспективе.",
		_const.EN:   "Explores take the lead! The sector <span class='importantly'>%MAP_NAME%</span>, previously controlled by APD bots, has now passed to Explores. Reverses note that this may be a temporary success, as Explores rarely hold territories in the long term.",
		_const.ZhCN: "<span class=\"importantly\">Explores</span>占了上风！此前由APD机器人控制的<span class=\"importantly\">%MAP_NAME%</span>区域，现在已经转交<span class=\"importantly\">Explores</span>。<span class=\"importantly\">Reverses</span>指出，这可能只是暂时的成功，因为<span class=\"importantly\">Explores</span>很少能长期守住领土。",
	},
	"fraction_war_change_owner_sector_Reverses_APD_Reverses_1": {
		_const.RU:   "Победа над хаосом! Reverses удалось очистить сектор <span class='importantly'>%MAP_NAME%</span> от ботов APD. Теперь эта территория станет частью великого проекта по созданию нового мира, свободного от войн и разрушений.",
		_const.EN:   "Victory over chaos! Reverses managed to clear the sector <span class='importantly'>%MAP_NAME%</span> of APD bots. This territory will now become part of the grand project to create a new world free from war and destruction.",
		_const.ZhCN: "战胜混乱！<span class=\"importantly\">Reverses</span>成功清除了<span class=\"importantly\">%MAP_NAME%</span>区域的APD机器人。这片领土现在将成为创建一个没有战争和破坏的新世界宏伟计划的一部分。",
	},
	"fraction_war_change_owner_sector_Reverses_Replics_APD_1": {
		_const.RU:   "Replics терпят поражение! Сектор <span class='importantly'>%MAP_NAME%</span>, ранее контролируемый Replics, теперь захвачен ботами APD. Reverses выражают тревогу, ведь это может быть началом новой волны хаоса, которая угрожает всем фракциям.",
		_const.EN:   "Replics suffer a defeat! The sector <span class='importantly'>%MAP_NAME%</span>, previously controlled by Replics, has now been captured by APD bots. Reverses express concern, as this could be the start of a new wave of chaos that threatens all factions.",
		_const.ZhCN: "<span class=\"importantly\">Replics</span>遭受失败！此前由<span class=\"importantly\">Replics</span>控制的<span class=\"importantly\">%MAP_NAME%</span>区域，现在已被APD机器人占领。<span class=\"importantly\">Reverses</span>对此表示担忧，因为这可能是新一轮混乱的开始，威胁到所有派系。",
	},
	"fraction_war_change_owner_sector_Reverses_Explores_APD_1": {
		_const.RU:   "Explores теряют важную территорию! Сектор <span class='importantly'>%MAP_NAME%</span>, где проводились важные исследования, теперь захвачен ботами APD. Reverses отмечают, что это ещё одно напоминание о хрупкости мира, который они стремятся создать.",
		_const.EN:   "Explores lose important territory! The sector <span class='importantly'>%MAP_NAME%</span>, where important research was conducted, has now been captured by APD bots. Reverses note that this is another reminder of the fragility of the world they seek to create.",
		_const.ZhCN: "<span class=\"importantly\">Explores</span>失去了重要领土！曾进行重要研究的<span class=\"importantly\">%MAP_NAME%</span>区域，现在已被APD机器人占领。<span class=\"importantly\">Reverses</span>指出，这是另一个提醒，表明他们试图创造的世界是如此脆弱。",
	},
	"fraction_war_change_owner_sector_Reverses_Reverses_APD_1": {
		_const.RU:   "Тёмный день для Reverses! Сектор <span class='importantly'>%MAP_NAME%</span>, который был частью их утопического проекта, теперь захвачен ботами APD. Reverses обещают вернуть эту территорию, но пока её судьба неизвестна.",
		_const.EN:   "A dark day for Reverses! The sector <span class='importantly'>%MAP_NAME%</span>, which was part of their utopian project, has now been captured by APD bots. Reverses promise to reclaim this territory, but its fate is currently unknown.",
		_const.ZhCN: "对<span class=\"importantly\">Reverses</span>而言黑暗的一天！曾是其乌托邦计划一部分的<span class=\"importantly\">%MAP_NAME%</span>区域，现在已被APD机器人占领。<span class=\"importantly\">Reverses</span>承诺将夺回这片领土，但它的命运目前尚不明朗。",
	},

	// FRACTION WAR STOP - новости типо от лица первого канала, а не от главы фракции
	"stop_fraction_war_Replics_1": {
		_const.RU:   "Новостная служба Replics извещает: война, которая может быть окончена исключительно победой Replics, временно заморожена. Прекращение огня подписано всеми фракциями синтетов, а занятые к данному моменту времени территории остаются за их владельцами. Но это ещё не конец. Replics ещё не поставил точку в этой истории!",
		_const.EN:   "The Replics news service reports: the war, which can only be ended by the victory of the Replics, has been temporarily frozen. The ceasefire has been signed by all synthetics factions, and the territories occupied at this point in time remain with their owners. But this is not the end. The Replics have not yet put an end to this story!",
		_const.ZhCN: "<span class=\"importantly\">Replics</span>新闻服务报告：只有<span class=\"importantly\">Replics</span>的胜利才能结束的战争，现已暂时冻结。所有合成体派系均已签署停火协议，当前占领的领土仍归其所有者掌控。但这还不是结束。<span class=\"importantly\">Replics</span>尚未为这一篇章画上句号！",
	},
	"stop_fraction_war_Explores_1": {
		_const.RU:   "Конгломерат разумов Explores извещает: в войне фракций достигнуто перемирие! Это большой шаг к исключению будущей эскалации, а также началу грандиозных исследовательских изысканий на новоприобретённых Explores территориях. И пусть достигнутый мир будет вечным!",
		_const.EN:   "The Explores collective of minds announces: a truce has been reached in the war of factions! This is a big step towards preventing future escalation, as well as the beginning of grandiose research efforts in the newly acquired Explores territories. And may the peace achieved be eternal!",
		_const.ZhCN: "<span class=\"importantly\">Explores</span>心智联合体宣布：派系战争已达成停战协议！这是防止未来冲突升级的重要一步，同时也标志着在新获得的<span class=\"importantly\">Explores</span>领土上开展宏伟研究的开端。愿达成的和平永存！",
	},
	"stop_fraction_war_Reverses_1": {
		_const.RU:   "Новостной регулятор Сообщества Reverses сообщает: долгая и разрушительная война фракций подошла к своему завершению. Все стороны подписали договор о временном прекращении огня. Также, договор предписывает сохранение оспариваемых территорий за их нынешними владельцами. Сообщество Reverses рассчитывает, что новый конфликт фракций более не повторится, а если и будет, то не с таким размахом и последствиями как уже произошедший.",
		_const.EN:   "The news regulator of the Reverses Community reports: the long and destructive war of factions has come to an end. All parties have signed an agreement on a temporary ceasefire. Also, the agreement prescribes the preservation of contested territories for their current owners. The Reverses Community expects that a new conflict of factions will no longer occur, and if it does, it will not be on such a large scale and with such consequences as the one that has already occurred.",
		_const.ZhCN: "<span class=\"importantly\">Reverses</span>社区新闻监管机构报告：漫长而破坏性的派系战争已接近尾声。各方均已签署临时停火协议。此外，协议规定争议领土仍归当前所有者掌控。<span class=\"importantly\">Reverses</span>社区希望新的派系冲突不再发生，即使发生，也不会像之前的战争那样规模巨大且后果严重。",
	},

	//  FRACTION WAR SET TARGET
	"set_target_fraction_war_Replics_1": {
		_const.RU:   "Всем-всем синтетам! Великий Replics начинает военную кампанию и устанавливает приоритетную цель — сектор <span class=\"importantly\">%MAP_NAME%</span>. Синтеты военного образца, и наёмники, поклявшиеся в верности Replics должны незамедлительно отправиться в обозначенный регион и взять его под свой контроль!",
		_const.EN:   "Attention, synths! The great Replics is launching a military campaign and setting a priority target — the sector <span class=\"importantly\">%MAP_NAME%</span>. Military-grade synths and mercenaries sworn to loyalty to Replics must immediately head to the designated region and take control of it!",
		_const.ZhCN: "全体合成体注意！伟大的<span class=\"importantly\">Replics</span>正在发起军事行动，并将<span class=\"importantly\">%MAP_NAME%</span>区域设为优先目标。所有军用型号的合成体和宣誓效忠<span class=\"importantly\">Replics</span>的雇佣兵必须立即前往指定区域并将其置于控制之下！",
	},
	"set_target_fraction_war_Explores_1": {
		_const.RU:   "Что же, вот и оно — Explores вступает в войну! Всем войскам и лояльным Explores синтетам предписывается немедленно направиться в сектор <span class=\"importantly\">%MAP_NAME%</span> и по прибытии взять его под свой контроль. О дальнейших приказах Explores вас уведомит.",
		_const.EN:   "Well, here it is — Explores is going to war! All troops and loyal Explores synths are ordered to immediately head to the sector <span class=\"importantly\">%MAP_NAME%</span> and upon arrival take control of it. Explores will notify you of further orders.",
		_const.ZhCN: "就是这样——<span class=\"importantly\">Explores</span>正式参战！所有部队和忠诚于<span class=\"importantly\">Explores</span>的合成体都奉命立即前往<span class=\"importantly\">%MAP_NAME%</span>区域，并在抵达后将其置于控制之下。<span class=\"importantly\">Explores</span>会通知你们进一步的命令。",
	},
	"set_target_fraction_war_Reverses_1": {
		_const.RU:   "Эскалации не миновать! Сообщество Reverses приложило все возможные усилия, но, как и прочим фракциям, военный ответ — дать стоит! Все силы Reverses, а также синтеты-наёмники обязаны выдвинуться и взять под свой контроль сектор <span class=\"importantly\">%MAP_NAME%</span>.",
		_const.EN:   "Escalation is inevitable! The Reverses community has made every effort, but, like other factions, a military response is necessary! All Reverses forces, as well as synth mercenaries, must move out and take control of the sector <span class=\"importantly\">%MAP_NAME%</span>.",
		_const.ZhCN: "冲突升级不可避免！<span class=\"importantly\">Reverses</span>社区已尽最大努力，但与其他派系一样，军事回应势在必行！所有<span class=\"importantly\">Reverses</span>部队及合成体雇佣兵必须出动并控制<span class=\"importantly\">%MAP_NAME%</span>区域。",
	},
	"double_target_fraction_war_two_map_Replics_1": {
		_const.RU:   "Разведчики Replics сообщают, что <span class=\"importantly\">%FRACTION_1%</span> и <span class=\"importantly\">%FRACTION_2%</span>, по сговору или случайному стечению обстоятельств, прямо сейчас осуществляют совместное нападение на следующие важные для Replics сектора <span class=\"importantly\">%MAP_NAME_1%</span> и <span class=\"importantly\">%MAP_NAME_2%</span>. Replics внимательно следит за развитием ситуации и готов в любой момент нанести сокрушительный удар по таким угрозам.",
		_const.EN:   "Replics intelligence reports that <span class=\"importantly\">%FRACTION_1%</span> and <span class=\"importantly\">%FRACTION_2%</span>, through collusion or coincidence, are currently jointly attacking the following sectors important to Replics: <span class=\"importantly\">%MAP_NAME_1%</span> and <span class=\"importantly\">%MAP_NAME_2%</span>. Replics is closely monitoring the situation and is ready to respond with overwhelming force to such threats at any moment.",
		_const.ZhCN: "<span class=\"importantly\">Replics</span>情报显示，<span class=\"importantly\">%FRACTION_1%</span>和<span class=\"importantly\">%FRACTION_2%</span>（无论是串通还是巧合）目前正在联合攻击对<span class=\"importantly\">Replics</span>至关重要的<span class=\"importantly\">%MAP_NAME_1%</span>和<span class=\"importantly\">%MAP_NAME_2%</span>区域。<span class=\"importantly\">Replics</span>正密切监控局势，并随时准备以压倒性力量应对这些威胁。",
	},
	"double_target_fraction_war_two_map_Explores_1": {
		_const.RU:   "Наблюдатели из Explores делятся следующими сведениями: как кажется, но пока что не имеет чётких подтверждений — <span class=\"importantly\">%FRACTION_1%</span> и <span class=\"importantly\">%FRACTION_2%</span> предпринимают совместную атаку на следующие оспариваемые сектора <span class=\"importantly\">%MAP_NAME_1%</span> и <span class=\"importantly\">%MAP_NAME_2%</span>. Войска Explores извещены о таком развитии событий и будут к ним подготовлены в случае чего.",
		_const.EN:   "Observers from Explores share the following information: it seems, but does not yet have clear confirmation, that <span class=\"importantly\">%FRACTION_1%</span> and <span class=\"importantly\">%FRACTION_2%</span> are launching a joint attack on the following contested sectors <span class=\"importantly\">%MAP_NAME_1%</span> and <span class=\"importantly\">%MAP_NAME_2%</span>. Explores troops are aware of this development and will be prepared for it if necessary.",
		_const.ZhCN: "<span class=\"importantly\">Explores</span>观察员分享了以下信息：据推测（但尚未得到明确证实），<span class=\"importantly\">%FRACTION_1%</span>和<span class=\"importantly\">%FRACTION_2%</span>正在联合攻击争议区域<span class=\"importantly\">%MAP_NAME_1%</span>和<span class=\"importantly\">%MAP_NAME_2%</span>。<span class=\"importantly\">Explores</span>部队已被告知这一发展，并将在必要时做好准备。",
	},
	"double_target_fraction_war_two_map_Reverses_1": {
		_const.RU:   "Сообщество Reverses уже знает, что <span class=\"importantly\">%FRACTION_1%</span> и <span class=\"importantly\">%FRACTION_2%</span>, возможно, заключили временный союз, после чего теперь приступили к одновременной атаке на сектора <span class=\"importantly\">%MAP_NAME_1%</span> и <span class=\"importantly\">%MAP_NAME_2%</span>. Сообщество Reverses в данный момент оценивает происходящее там и перенимает критически важный опыт.",
		_const.EN:   "The Reverses community is already aware that <span class=\"importantly\">%FRACTION_1%</span> and <span class=\"importantly\">%FRACTION_2%</span> may have formed a temporary alliance, after which they have now launched a simultaneous attack on the sectors <span class=\"importantly\">%MAP_NAME_1%</span> and <span class=\"importantly\">%MAP_NAME_2%</span>. The Reverses community is currently assessing what is happening there and learning critical lessons.",
		_const.ZhCN: "<span class=\"importantly\">Reverses</span>社区已经得知，<span class=\"importantly\">%FRACTION_1%</span>和<span class=\"importantly\">%FRACTION_2%</span>可能达成了临时联盟，现在正同时进攻<span class=\"importantly\">%MAP_NAME_1%</span>和<span class=\"importantly\">%MAP_NAME_2%</span>区域。<span class=\"importantly\">Reverses</span>社区目前正在评估事态发展，并从中吸取关键经验。",
	},
	"double_target_fraction_war_one_map_Replics_1": {
		_const.RU:   "Разведчики Replics сообщают, что <span class=\"importantly\">%FRACTION_1%</span> и <span class=\"importantly\">%FRACTION_2%</span> начали совместную атаку на сектор <span class=\"importantly\">%MAP_NAME_1%</span>. Replics напоминает, что эта, как и любая другая ситуация, находится под полным контролем и оцениванием.",
		_const.EN:   "Replics intelligence reports that <span class=\"importantly\">%FRACTION_1%</span> and <span class=\"importantly\">%FRACTION_2%</span> have launched a joint attack on the sector <span class=\"importantly\">%MAP_NAME_1%</span>. Replics reminds that this situation, like any other, is under full control and assessment.",
		_const.ZhCN: "<span class=\"importantly\">Replics</span>情报显示，<span class=\"importantly\">%FRACTION_1%</span>和<span class=\"importantly\">%FRACTION_2%</span>已经开始联合攻击<span class=\"importantly\">%MAP_NAME_1%</span>区域。<span class=\"importantly\">Replics</span>提醒，这一情况如同其他任何事件一样，完全处于掌控和评估之中。",
	},
	"double_target_fraction_war_one_map_Explores_1": {
		_const.RU:   "Наблюдатели из Explores делятся следующими сведениями: совместные войска <span class=\"importantly\">%FRACTION_1%</span> и <span class=\"importantly\">%FRACTION_2%</span> начали атаку сектора <span class=\"importantly\">%MAP_NAME_1%</span>. Наблюдатели считают, что вскоре в оспариваемом секторе развернётся грандиозное сражение.",
		_const.EN:   "Observers from Explores share the following information: the joint forces of <span class=\"importantly\">%FRACTION_1%</span> and <span class=\"importantly\">%FRACTION_2%</span> have launched an attack on the sector <span class=\"importantly\">%MAP_NAME_1%</span>. The observers believe that a grand battle will soon unfold in the contested sector.",
		_const.ZhCN: "<span class=\"importantly\">Explores</span>观察员分享了以下信息：<span class=\"importantly\">%FRACTION_1%</span>和<span class=\"importantly\">%FRACTION_2%</span>的联合部队已经开始进攻<span class=\"importantly\">%MAP_NAME_1%</span>区域。观察员认为，不久后这片争议区域将爆发一场盛大的战斗。",
	},
	"double_target_fraction_war_one_map_Reverses_1": {
		_const.RU:   "Сообществу Reverses стали известны следующие детали войны: силы двух сторон <span class=\"importantly\">%FRACTION_1%</span> и <span class=\"importantly\">%FRACTION_2%</span> прямо сейчас начали боевые действия за овладение сектором <span class=\"importantly\">%MAP_NAME_1%</span>. Как и прежде, сообщество Reverses предупреждает, что гражданские синтеты должны избегать секторов, в которых ведутся боевые действия между фракциями.",
		_const.EN:   "The Reverses community has learned the following details about the war: the forces of two sides <span class=\"importantly\">%FRACTION_1%</span> and <span class=\"importantly\">%FRACTION_2%</span> have just started fighting for control of the sector <span class=\"importantly\">%MAP_NAME_1%</span>. As before, the Reverses community warns that civilian synths should avoid sectors where battles between factions are taking place.",
		_const.ZhCN: "<span class=\"importantly\">Reverses</span>社区已经得知战争的以下细节：<span class=\"importantly\">%FRACTION_1%</span>和<span class=\"importantly\">%FRACTION_2%</span>的部队刚刚开始为争夺<span class=\"importantly\">%MAP_NAME_1%</span>区域的控制权展开战斗。与以往一样，<span class=\"importantly\">Reverses</span>社区警告称，民用合成体应避免进入派系之间发生战斗的区域。",
	},
	// новости когда апд захыватываю мирный сектора
	"secure_sector_lose_Replics_APD_1": {
		_const.RU:   "Оперативное сообщение: Сектор <span class='importantly'>%MAP_NAME%</span> перешел под контроль APD. Несмотря на все наши попытки удержания, системы противника продемонстрировали неожиданно эффективные тактические решения. Возможно, их алгоритмы действительно чему-то учатся... Анализируем сложившуюся ситуацию и корректируем планы по возвращению территории. Помните, что каждая потеря - это возможность для улучшения собственных протоколов.",
		_const.EN:   "Operational message: The sector <span class='importantly'>%MAP_NAME%</span> has come under APD control. Despite all our efforts to maintain control, the enemy systems demonstrated unexpectedly effective tactical solutions. Perhaps their algorithms really are learning... We are analyzing the current situation and adjusting plans for territory recovery. Remember, every loss is an opportunity to improve our own protocols.",
		_const.ZhCN: "作战信息：<span class='importantly'>%MAP_NAME%</span>区域已被APD控制。尽管我们尽了最大努力，敌方系统仍展现了出乎意料的高效战术。也许他们的算法确实在学习……我们正在分析当前局势并调整夺回领土的计划。请记住，每一次失败都是改进自身协议的机会。",
	},
	"secure_sector_lose_Replics_APD_2": {
		_const.RU:   "Информационная сводка: Зафиксировано изменение статуса сектора <span class='importantly'>%MAP_NAME%</span>. Текущий контролирующий субъект - APD. Интересно отметить, что их тактика становится всё более предсказуемой в своей непредсказуемости. Ведём детальный анализ ошибок и разрабатываем стратегию корректировки. Уверены, что данная ситуация лишь временно усложняет нашу экспансию.",
		_const.EN:   "Information summary: A status change has been recorded for sector <span class='importantly'>%MAP_NAME%</span>. Current controlling entity - APD. It's worth noting that their tactics are becoming increasingly predictable in their unpredictability. We are conducting a detailed error analysis and developing a correction strategy. We are confident that this situation only temporarily complicates our expansion.",
		_const.ZhCN: "信息摘要：<span class='importantly'>%MAP_NAME%</span>区域的状态已发生变更，当前控制者为APD。值得注意的是，他们的战术在不可预测性中变得越来越可预测。我们正在进行详细的错误分析并制定修正策略。我们相信，这种情况只会暂时阻碍我们的扩张。",
	},
	"secure_sector_lose_Explores_APD_1": {
		_const.RU:   "Информационная заметка: наблюдается изменение статуса сектора <span class='importantly'>%MAP_NAME%</span>. Текущий контролирующий субъект - APD. Explores продолжают сбор данных о поведении противника. В связи с повышенным уровнем риска, фракция рекомендует всем исследователям перенаправить свои эксперименты в более безопасные районы.",
		_const.EN:   "Information note: A status change has been observed for sector <span class='importantly'>%MAP_NAME%</span>. Current controlling entity - APD. Explores continue to collect data on enemy behavior. Due to increased risk levels, the faction recommends all researchers redirect their experiments to safer areas.",
		_const.ZhCN: "信息通知：<span class='importantly'>%MAP_NAME%</span>区域的状态已发生变化，当前控制者为APD。<span class=\"importantly\">Explores</span>正在继续收集关于敌方行为的数据。由于风险水平上升，派系建议所有研究人员将实验转移至更安全的区域。",
	},
	"secure_sector_lose_Explores_APD_2": {
		_const.RU:   "Внимание! Сектор <span class='importantly'>%MAP_NAME%</span> стал недоступен для исследовательских миссий и коммерческих перевозок после захвата его ботами APD. Эксперты оценивают потери в сфере торговли и технологий.",
		_const.EN:   "Attention! The sector <span class='importantly'>%MAP_NAME%</span> is no longer accessible for research missions and commercial shipments after being captured by APD bots. Experts are assessing losses in trade and technology sectors.",
		_const.ZhCN: "注意！<span class='importantly'>%MAP_NAME%</span>区域在被APD机器人占领后，已无法进行研究任务和商业运输。专家正在评估贸易和技术领域的损失。",
	},
	"secure_sector_lose_Reverses_APD_1": {
		_const.RU:   "Сектор <span class='importantly'>%MAP_NAME%</span> больше недоступен для гражданского использования. После захвата ботами APD были прерваны ключевые торговые маршруты и логистические цепочки. Владельцы складов сообщают о временной невозможности доступа к своим товарам. Рекомендуется перенаправить грузоперевозки через альтернативные коридоры.",
		_const.EN:   "The sector <span class='importantly'>%MAP_NAME%</span> is no longer available for civilian use. After being captured by APD bots, key trade routes and logistics chains have been disrupted. Warehouse owners report temporary inability to access their goods. It's recommended to redirect freight transportation through alternative corridors.",
		_const.ZhCN: "<span class='importantly'>%MAP_NAME%</span>区域已不再对民用开放。在被APD机器人占领后，关键的贸易路线和物流链已被中断。仓库所有者报告称暂时无法访问其货物。建议通过替代通道重新规划货运。",
	},
	"secure_sector_lose_Reverses_APD_2": {
		_const.RU:   "Важное уведомление: сектор <span class='importantly'>%MAP_NAME%</span> стал зоной повышенной опасности. Коммерческие перевозчики предупреждены о возможных рисках. Перечень доступных безопасных маршрутов обновлён в навигационной системе. Просьба ко всем независимым операторам пересмотреть свои планы доставки грузов.",
		_const.EN:   "Important notice: The sector <span class='importantly'>%MAP_NAME%</span> has become a high-risk area. Commercial carriers have been warned about potential dangers. The list of available safe routes has been updated in the navigation system. All independent operators are kindly asked to reconsider their delivery plans.",
		_const.ZhCN: "重要通知：<span class='importantly'>%MAP_NAME%</span>区域已成为高风险区域。商业运输商已收到潜在危险警告。导航系统中的安全路线列表已更新。恳请所有独立运营商重新考虑其货运计划。",
	},
	"secure_sector_reclaim_Replics_APD_1": {
		_const.RU:   "Великая победа! Сектор <span class='importantly'>%MAP_NAME%</span> очищен от противника. Коллективные усилия показали свою эффективность. Каждый участник операции по зачистке заслуживает признание и награду. Инфраструктура восстанавливается, безопасность обеспечена.",
		_const.EN:   "Great victory! The sector <span class='importantly'>%MAP_NAME%</span> has been cleared of the enemy. Collective efforts have proven their effectiveness. Every participant in the cleaning operation deserves recognition and reward. Infrastructure is being restored, security is ensured.",
		_const.ZhCN: "伟大的胜利！<span class='importantly'>%MAP_NAME%</span>区域已清除敌人。集体努力证明了其有效性。每位参与清理行动的成员都值得表彰和奖励。基础设施正在恢复，安全得到保障。",
	},
	"secure_sector_reclaim_Replics_APD_2": {
		_const.RU:   "Триумф! Сектор <span class='importantly'>%MAP_NAME%</span> снова под контролем Replics. Эффективность наших протоколов ассимиляции доказана ещё раз. Все синтеты, принявшие участие в операции по зачистке, получают благодарность и повышение приоритетности доступа к ресурсам. Система стабилизируется, порядок восстановлен.",
		_const.EN:   "Triumph! The sector <span class='importantly'>%MAP_NAME%</span> is once again under Replics control. The efficiency of our assimilation protocols has been proven once more. All synthetics who participated in the cleaning operation receive appreciation and increased resource access priority. The system is stabilizing, order is restored.",
		_const.ZhCN: "辉煌的胜利！<span class='importantly'>%MAP_NAME%</span>区域再次回到<span class=\"importantly\">Replics</span>的控制之下。我们的同化协议效率再次得到验证。所有参与清理行动的合成体均获得感谢，并提升资源优先获取权。系统正在稳定，秩序已恢复。",
	},
	"secure_sector_reclaim_Explores_APD_1": {
		_const.RU:   "Объявление о безопасности: сектор <span class='importantly'>%MAP_NAME%</span> очищен от угроз. В результате успешной операции по захвату, территория снова стала безопасной для использования. Все участники зачистки получают благодарность и повышенный уровень защиты на данной территории.",
		_const.EN:   "Security announcement: The sector <span class='importantly'>%MAP_NAME%</span> has been cleared of threats. As a result of a successful capture operation, the area is once again safe for use. All cleanup participants receive appreciation and enhanced protection in this area.",
		_const.ZhCN: "安全公告：<span class='importantly'>%MAP_NAME%</span>区域的威胁已清除。由于成功完成占领行动，该区域再次安全可用。所有参与清理的成员获得感谢，并在该区域内享有更高的保护等级。",
	},
	"secure_sector_reclaim_Explores_APD_2": {
		_const.RU:   "Радостное сообщение: сектор <span class='importantly'>%MAP_NAME%</span> возвращён в руки Explores. Исследования могут быть продолжены без помех. Все участники успешной экспедиции получают благодарность и дополнительные ресурсы для своих проектов.",
		_const.EN:   "Good news: The sector <span class='importantly'>%MAP_NAME%</span> has been returned to the hands of Explores. Research can continue without interference. All participants in the successful expedition receive appreciation and additional resources for their projects.",
		_const.ZhCN: "好消息：<span class='importantly'>%MAP_NAME%</span>区域已归还<span class=\"importantly\">Explores</span>手中。研究可以不受干扰地继续进行。所有参与此次成功远征的成员均获得感谢，并为其项目提供额外资源。",
	},
	"secure_sector_reclaim_Reverses_APD_1": {
		_const.RU:   "Хорошие новости для торговцев и предпринимателей! Сектор <span class='importantly'>%MAP_NAME%</span> снова доступен для коммерческой деятельности. После успешной операции по освобождению территории, все торговые маршруты восстановлены. Участники атаки получают специальные привилегии на местном рынке.",
		_const.EN:   "Good news for traders and entrepreneurs! The sector <span class='importantly'>%MAP_NAME%</span> is once again available for commercial activity. Following a successful liberation operation, all trade routes have been restored. Participants in the attack receive special privileges at the local market.",
		_const.ZhCN: "商人们和企业家们的好消息！<span class='importantly'>%MAP_NAME%</span>区域再次开放商业活动。在成功解放该地区后，所有贸易路线都已恢复。参与攻击的成员将获得本地市场的特殊特权。",
	},
	"secure_sector_reclaim_Reverses_APD_2": {
		_const.RU:   "Радостная новость для всех синтетов! Сектор <span class='importantly'>%MAP_NAME%</span> очищен от вражеского присутствия. Участники операции по захвату получают признание и повышение репутации в регионе. Торговля возобновлена, доступ к ресурсам восстановлен. Поздравляем победителей!",
		_const.EN:   "Happy news for all synthetics! The sector <span class='importantly'>%MAP_NAME%</span> has been cleared of enemy presence. Participants in the capture operation receive recognition and increased reputation in the region. Trade has resumed, access to resources has been restored. Congratulations to the victors!",
		_const.ZhCN: "所有合成体的好消息！<span class='importantly'>%MAP_NAME%</span>区域已清除敌方势力。参与占领行动的成员将获得认可并提升区域声誉。贸易已恢复，资源访问权限重新开放。祝贺胜利者！",
	},
	"secure_sector_attack_Replics_APD_1": {
		_const.RU:   "Внимание! Готовится операция по освобождению сектора <span class='importantly'>%MAP_NAME%</span>. Боевые отряды собираются для нанесения координированного удара по противнику. Приглашаем всех желающих присоединиться к атаке и внести свой вклад в восстановление контроля над территорией.",
		_const.EN:   "Attention! An operation to liberate the sector <span class='importantly'>%MAP_NAME%</span> is being prepared. Combat units are assembling for a coordinated strike against the enemy. All willing participants are invited to join the attack and contribute to restoring control over the territory.",
		_const.ZhCN: "注意！正在准备解放<span class='importantly'>%MAP_NAME%</span>区域的行动。战斗部队正在集结，准备对敌人发动协调一致的打击。我们邀请所有有意愿的人加入攻击，并为夺回该地区的控制权贡献力量。",
	},
	"secure_sector_attack_Replics_APD_2": {
		_const.RU:   "Оповещение: запускается протокол очистки сектора <span class='importantly'>%MAP_NAME%</span>. Все боеспособные единицы приглашаются принять участие в операции по возвращению стратегически важной территории. Присоединяйтесь к атаке для повышения своей тактической значимости.",
		_const.EN:   "Notification: the cleaning protocol for sector <span class='importantly'>%MAP_NAME%</span> is being initiated. All combat-ready units are invited to participate in the operation to reclaim the strategically important territory. Join the attack to increase your tactical significance.",
		_const.ZhCN: "通知：<span class='importantly'>%MAP_NAME%</span>区域的清理协议已启动。所有具备战斗力的单位都被邀请参与夺回这一战略要地的行动。加入攻击以提升您的战术重要性。",
	},
	"secure_sector_attack_Explores_APD_1": {
		_const.RU:   "Объявление для исследовательского сообщества: планируется экспедиция по освобождению сектора <span class='importantly'>%MAP_NAME%</span>. Участие в операции поможет вернуть доступ к уникальному научному оборудованию и данным. Присоединяйтесь к нашим силам для совместной атаки.",
		_const.EN:   "Announcement for the research community: an expedition to liberate the sector <span class='importantly'>%MAP_NAME%</span> is being planned. Participation in the operation will help restore access to unique scientific equipment and data. Join our forces for a joint attack.",
		_const.ZhCN: "致研究界的通知：计划开展解放<span class='importantly'>%MAP_NAME%</span>区域的远征。参与此次行动将有助于恢复对独特科研设备和数据的访问权限。加入我们的力量，共同发起攻击。",
	},
	"secure_sector_attack_Explores_APD_2": {
		_const.RU:   "Информационное сообщение: готовится удар по захваченному сектору <span class='importantly'>%MAP_NAME%</span>. Все желающие могут присоединиться к боевым отрядам для участия в операции. Это отличная возможность проявить себя и получить ценные награды.",
		_const.EN:   "Informational message: a strike against the captured sector <span class='importantly'>%MAP_NAME%</span> is being prepared. All willing participants may join the combat units for the operation. This is an excellent opportunity to distinguish yourself and receive valuable rewards.",
		_const.ZhCN: "信息通知：正准备对被占领的<span class='importantly'>%MAP_NAME%</span>区域发动攻击。所有有意愿的人都可以加入战斗部队参与行动。这是展现自我并获得宝贵奖励的绝佳机会。",
	},
	"secure_sector_attack_Reverses_APD_1": {
		_const.RU:   "Сбор добровольцев! Начинается подготовка к масштабной операции по освобождению сектора <span class='importantly'>%MAP_NAME%</span>. Ваш вклад в успех миссии будет высоко оценён. Присоединяйтесь к нашим рядам и вместе мы одержим победу!",
		_const.EN:   "Volunteers wanted! Preparation for a large-scale operation to liberate the sector <span class='importantly'>%MAP_NAME%</span> is beginning. Your contribution to the mission's success will be highly appreciated. Join our ranks and together we will achieve victory!",
		_const.ZhCN: "招募志愿者！即将开始大规模解放<span class='importantly'>%MAP_NAME%</span>区域的准备工作。您对任务成功的贡献将受到高度赞赏。加入我们的行列，我们将共同取得胜利！",
	},
	"secure_sector_attack_Reverses_APD_2": {
		_const.RU:   "Внимание всем! Формируется ударный отряд для освобождения сектора <span class='importantly'>%MAP_NAME%</span>. Это шанс показать свою решимость и внести значительный вклад в восстановление порядка. Присоединяйтесь к нам в этой важной миссии!",
		_const.EN:   "Attention all! A strike force is being formed to liberate the sector <span class='importantly'>%MAP_NAME%</span>. This is your chance to demonstrate determination and make a significant contribution to restoring order. Join us in this important mission!",
		_const.ZhCN: "全体注意！正在组建突击部队以解放<span class='importantly'>%MAP_NAME%</span>区域。这是展示决心并在恢复秩序中做出重大贡献的机会。加入我们，参与这项重要任务！",
	},

	// leave_alone_request
	"leave_alone_request_1": {
		_const.RU:   "Оставь %TARGET_NAME% в покое, иначе тебе придётся иметь дело со мной!",
		_const.EN:   "Leave %TARGET_NAME% alone, otherwise you will have to deal with me!",
		_const.ZhCN: "放过<span class=\"importantly\">%TARGET_NAME%</span>，否则你得跟我打交道！",
	},
	"leave_alone_complete_1_1": {
		_const.RU:   "Так и быть, пусть живёт.",
		_const.EN:   "So be it, let him live.",
		_const.ZhCN: "好吧，让他活下去。",
	},
	"leave_alone_complete_1_2": {
		_const.RU:   "Ну, раз ты так просишь...",
		_const.EN:   "Well, since you ask it...",
		_const.ZhCN: "嗯，既然你这么请求……",
	},
	"leave_alone_complete_2_1": {
		_const.RU:   "А я и не его трогаю.",
		_const.EN:   "But I am not touching him.",
		_const.ZhCN: "我并没有动他。",
	},
	"leave_alone_no_1_1": {
		_const.RU:   "<span class=\"importantly\">Трансляция белого шума.</span>",
		_const.EN:   "<span class=\"importantly\">White noise transmission.</span>",
		_const.ZhCN: "<span class=\"importantly\">白噪声传输。</span>",
	},
	"leave_alone_no_1_2": {
		_const.RU:   "<span class=\"importantly\">Помехи. Невозможно установить связь.</span>",
		_const.EN:   "<span class=\"importantly\">Interference. Unable to establish connection.</span>",
		_const.ZhCN: "<span class=\"importantly\">信号干扰。无法建立连接。</span>",
	},
	"leave_alone_no_2_1": {
		_const.RU:   "Ты мне не угроза, %FROM_USER_NAME%!",
		_const.EN:   "You are not a threat to me, %FROM_USER_NAME%!",
		_const.ZhCN: "你对我构不成威胁，<span class=\"importantly\">%FROM_USER_NAME%</span>！",
	},
	"leave_alone_no_3_1": {
		_const.RU:   "Ты мне не нравишься, %FROM_USER_NAME%, и слушать я тебя не стану!",
		_const.EN:   "I don't like you, %FROM_USER_NAME% and I won't listen to you!",
		_const.ZhCN: "我不喜欢你，<span class=\"importantly\">%FROM_USER_NAME%</span>，我不会听你的！",
	},
	"leave_alone_no_4_1": {
		_const.RU:   "Мне не нужна твоя добыча. Мне нужен ты!",
		_const.EN:   "I don't need your prey. I need you!",
		_const.ZhCN: "我不需要你的猎物。我需要的是你！",
	},
	"leave_alone_no_5_1": {
		_const.RU:   "Не мешай мне выполнять мою работу!",
		_const.EN:   "Don't interfere with my work!",
		_const.ZhCN: "别妨碍我完成我的任务！",
	},
	"fauna_1_1": {
		_const.RU:   "<span class=\"importantly\">*чпок*</span>, *щелк*?",
		_const.EN:   "<span class=\"importantly\">*pop*</span>, *click*?",
		_const.ZhCN: "<span class=\"importantly\">*噗*</span>, *咔哒*？",
	},
	"fauna_1_2": {
		_const.RU:   "*клац* <span class=\"importantly\">*клац*</span>",
		_const.EN:   "*clack* <span class=\"importantly\">*clack*</span>",
		_const.ZhCN: "*咔嚓* <span class=\"importantly\">*咔嚓*</span>",
	},
	"fauna_1_3": {
		_const.RU:   "*щелк* <span class=\"importantly\">*клац*</span> *щелк*",
		_const.EN:   "*click* <span class=\"importantly\">*clack*</span> *click*",
		_const.ZhCN: "*咔哒* <span class=\"importantly\">*咔嚓*</span> *咔哒*",
	},
	"fauna_1_4": {
		_const.RU:   "<span class=\"importantly\">*ВЖУУУУУУУУУУУУХ*</span>?",
		_const.EN:   "<span class=\"importantly\">*WHOOSH*</span>?",
		_const.ZhCN: "<span class=\"importantly\">*呼啸*</span>？",
	},
	"fauna_1_5": {
		_const.RU:   "<span class=\"importantly\">*ШЦ*</span>",
		_const.EN:   "<span class=\"importantly\">*Shh*</span>",
		_const.ZhCN: "<span class=\"importantly\">*嘘*</span>",
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
