package handbook

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"math/rand"
	"strings"
	"time"
)

const (
	Neutral = "neutral" // "Привет"
	Good    = "good"    // "Привет, друган"
	Bad     = "bad"     // "Убирайся, отвали", тут смысл в том что отношение плохое, но это не бой
	Target  = "target"  // А этот статус говорит что игрок на прицеле у нпс, и он его преследует
	Fear    = "fear"    // "Отстань от меня!"
	Help    = "help"    // Когда бот в страхе, но не из-за игрока
)

var npcGreetings = map[string]map[string]string{
	// Replicas
	_const.Replicas + "_miner_" + Neutral + "_1": {
		_const.RU: "Привет %UserName%. Я иду копать руду, не мешай мне.",
		_const.EN: "english translate",
	},
	_const.Replicas + "_guard_" + Neutral + "_1": {
		_const.RU: "Проезжай, не задерживайся.",
		_const.EN: "english translate",
	},
	_const.Replicas + "_out_scout_" + Neutral + "_1": {
		_const.RU: "Привет, я занимаюсь выполнением правительственных миссий, не мешай.",
		_const.EN: "english translate",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_1": {
		_const.RU: "Чего тебе?",
		_const.EN: "english translate",
	},
	_const.Replicas + "_transport_" + Neutral + "_1": {
		_const.RU: "Не мешай, я занимаюсь торговлей",
		_const.EN: "english translate",
	},
	_const.Replicas + "_builder_" + Neutral + "_1": {
		_const.RU: "Опять кто то что то сломал, а мне чинить...",
		_const.EN: "english translate",
	},
	_const.Replicas + "_warrior_" + Neutral + "_1": {
		_const.RU: "Не мешай выполнению военных обязанностей.",
		_const.EN: "english translate",
	},
	_const.Replicas + "_expedition_" + Neutral + "_1": {
		_const.RU: "Наша задача осваивать пустоши, не мешай нам и будешь цел.",
		_const.EN: "english translate",
	},

	_const.Replicas + "_miner_" + Good + "_1": {
		_const.RU: "Кажется я нашел руду! Хочешь помочь мне ее разработать?",
		_const.EN: "english translate",
	},
	_const.Replicas + "_miner_" + Good + "_2": {
		_const.RU: "Похоже в этом секторе уже выкопали всю руду, лучше поискать ее где то еще.",
		_const.EN: "english translate",
	},
	_const.Replicas + "_guard_" + Good + "_1": {
		_const.RU: "Вас что то беспокоить?",
		_const.EN: "english translate",
	},
	_const.Replicas + "_out_scout_" + Good + "_1": {
		_const.RU: "Кажется в пустошах стало слишком много враждебных ботов, будь осторожен!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_in_scout_" + Good + "_1": {
		_const.RU: "Может храбанем когонить?",
		_const.EN: "english translate",
	},
	_const.Replicas + "_transport_" + Good + "_1": {
		_const.RU: "Я везу %ProductName% на базу %BaseName%, это будет выгодная сделка!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_transport_" + Good + "_2": {
		_const.RU: "Я следую на %BaseName%, база очень выгодно продает %ProductName%!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_transport_" + Good + "_3": {
		_const.RU: "Привет %UserName%, как дела?",
		_const.EN: "english translate",
	},
	_const.Replicas + "_builder_" + Good + "_1": {
		_const.RU: "Привет как дела?",
		_const.EN: "english translate",
	},
	_const.Replicas + "_warrior_" + Good + "_1": {
		_const.RU: "Мы отправляемся в сектор %SectorName%, можешь присоединиться!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_expedition_" + Good + "_1": {
		_const.RU: "Мы пробуем освоить секторе %SectorName%, будем рады содействию!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_expedition_" + Good + "_2": {
		_const.RU: "Наша задача выбирать противника с территории пустошь в секторе %SectorName%! Будем рады содействию!",
		_const.EN: "english translate",
	},

	_const.Replicas + "_miner_" + Bad + "_1": {
		_const.RU: "Ну, иди сюда, попробуй меня трахнуть — я тебя сам трахну!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_guard_" + Bad + "_1": {
		_const.RU: "Не усугубляй!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_out_scout_" + Bad + "_1": {
		_const.RU: "Лучше уходи пока можешь!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_in_scout_" + Bad + "_1": {
		_const.RU: "Приветствую. Хочешь внести плату за проезд?",
		_const.EN: "english translate",
	},
	_const.Replicas + "_transport_" + Bad + "_1": {
		_const.RU: "Если ты помешаешь мне торговать то я за себя не ручаюсь!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_builder_" + Bad + "_1": {
		_const.RU: "Отвали, ты мне не интересен.",
		_const.EN: "english translate",
	},
	_const.Replicas + "_warrior_" + Bad + "_1": {
		_const.RU: "Советую тебе избегай встречи со мной!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_expedition_" + Bad + "_1": {
		_const.RU: "Не приближайся!",
		_const.EN: "english translate",
	},

	_const.Replicas + "_miner_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_guard_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_out_scout_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_in_scout_" + Target + "_1": {
		_const.RU: "Уже сдаешься? Это не бесплатная услуга!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_transport_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_builder_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_warrior_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_expedition_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "english translate",
	},

	_const.Replicas + "_miner_" + Fear + "_1": {
		_const.RU: "Прекрати! Чего тебе надо?!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_guard_" + Fear + "_1": {
		_const.RU: "Не усугубляй!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_out_scout_" + Fear + "_1": {
		_const.RU: "А знаешь что? У меня есть дела поважнее!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_in_scout_" + Fear + "_1": {
		_const.RU: "Я все понял! Пожалуйста не разрушай мой транспорт!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_transport_" + Fear + "_1": {
		_const.RU: "Может мы как то договоримся?!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_builder_" + Fear + "_1": {
		_const.RU: "Может мы как то договоримся?!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_warrior_" + Fear + "_1": {
		_const.RU: "Советую тебе избегай встречи со мной!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_expedition_" + Fear + "_1": {
		_const.RU: "Не приближайся!",
		_const.EN: "english translate",
	},

	_const.Replicas + "_miner_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_guard_" + Help + "_1": {
		_const.RU: "Не усугубляй!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_out_scout_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_in_scout_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_transport_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_builder_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_warrior_" + Help + "_1": {
		_const.RU: "Советую тебе избегай встречи со мной!",
		_const.EN: "english translate",
	},
	_const.Replicas + "_expedition_" + Help + "_1": {
		_const.RU: "Не приближайся!",
		_const.EN: "english translate",
	},

	// Explores
	_const.Explores + "_miner_" + Neutral + "_1": {
		_const.RU: "Привет %UserName%. Я иду копать руду, не мешай мне.",
		_const.EN: "english translate",
	},
	_const.Explores + "_guard_" + Neutral + "_1": {
		_const.RU: "Проезжай, не задерживайся.",
		_const.EN: "english translate",
	},
	_const.Explores + "_out_scout_" + Neutral + "_1": {
		_const.RU: "Привет, я занимаюсь какой то важной херней, не мешай.",
		_const.EN: "english translate",
	},
	_const.Explores + "_in_scout_" + Neutral + "_1": {
		_const.RU: "Чего тебе?",
		_const.EN: "english translate",
	},
	_const.Explores + "_transport_" + Neutral + "_1": {
		_const.RU: "Не мешай, я занимаюсь торговлей",
		_const.EN: "english translate",
	},
	_const.Explores + "_builder_" + Neutral + "_1": {
		_const.RU: "Опять кто то что то сломал, а мне чинить...",
		_const.EN: "english translate",
	},
	_const.Explores + "_warrior_" + Neutral + "_1": {
		_const.RU: "Не мешай выполнению военных обязанностей.",
		_const.EN: "english translate",
	},
	_const.Explores + "_expedition_" + Neutral + "_1": {
		_const.RU: "Наша задача осваивать пустоши, не мешай нам и будешь цел.",
		_const.EN: "english translate",
	},

	_const.Explores + "_miner_" + Good + "_1": {
		_const.RU: "Кажется я нашел руду! Хочешь помочь мне ее разработать?",
		_const.EN: "english translate",
	},
	_const.Explores + "_miner_" + Good + "_2": {
		_const.RU: "Похоже в этом секторе уже выкопали всю руду, лучше поискать ее где то еще.",
		_const.EN: "english translate",
	},
	_const.Explores + "_guard_" + Good + "_1": {
		_const.RU: "Вас что то беспокоить?",
		_const.EN: "english translate",
	},
	_const.Explores + "_out_scout_" + Good + "_1": {
		_const.RU: "Кажется в пустошах стало слишком много враждебных ботов, будь осторожен!",
		_const.EN: "english translate",
	},
	_const.Explores + "_in_scout_" + Good + "_1": {
		_const.RU: "Может храбанем когонить?",
		_const.EN: "english translate",
	},
	_const.Explores + "_transport_" + Good + "_1": {
		_const.RU: "Я везу %ProductName% на базу %BaseName%, это будет выгодная сделка!",
		_const.EN: "english translate",
	},
	_const.Explores + "_transport_" + Good + "_2": {
		_const.RU: "Я следую на %BaseName%, база очень выгодно продает %ProductName%!",
		_const.EN: "english translate",
	},
	_const.Explores + "_transport_" + Good + "_3": {
		_const.RU: "Привет %UserName%, как дела?",
		_const.EN: "english translate",
	},
	_const.Explores + "_builder_" + Good + "_1": {
		_const.RU: "Привет как дела?",
		_const.EN: "english translate",
	},
	_const.Explores + "_warrior_" + Good + "_1": {
		_const.RU: "Мы отправляемся в сектор %SectorName%, можешь присоединиться!",
		_const.EN: "english translate",
	},
	_const.Explores + "_expedition_" + Good + "_1": {
		_const.RU: "Мы пробуем освоить секторе %SectorName%, будем рады содействию!",
		_const.EN: "english translate",
	},
	_const.Explores + "_expedition_" + Good + "_2": {
		_const.RU: "Наша задача выбирать противника с территории пустошь в секторе %SectorName%! Будем рады содействию!",
		_const.EN: "english translate",
	},

	_const.Explores + "_miner_" + Bad + "_1": {
		_const.RU: "Ну, иди сюда, попробуй меня трахнуть — я тебя сам трахну!",
		_const.EN: "english translate",
	},
	_const.Explores + "_guard_" + Bad + "_1": {
		_const.RU: "Не усугубляй!",
		_const.EN: "english translate",
	},
	_const.Explores + "_out_scout_" + Bad + "_1": {
		_const.RU: "Лучше уходи пока можешь!",
		_const.EN: "english translate",
	},
	_const.Explores + "_in_scout_" + Bad + "_1": {
		_const.RU: "Приветствую. Хочешь внести плату за проезд?",
		_const.EN: "english translate",
	},
	_const.Explores + "_transport_" + Bad + "_1": {
		_const.RU: "Если ты помешаешь мне торговать то я за себя не ручаюсь!",
		_const.EN: "english translate",
	},
	_const.Explores + "_builder_" + Bad + "_1": {
		_const.RU: "Отвали, ты мне не интересен.",
		_const.EN: "english translate",
	},
	_const.Explores + "_warrior_" + Bad + "_1": {
		_const.RU: "Советую тебе избегай встречи со мной!",
		_const.EN: "english translate",
	},
	_const.Explores + "_expedition_" + Bad + "_1": {
		_const.RU: "Не приближайся!",
		_const.EN: "english translate",
	},

	_const.Explores + "_miner_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "english translate",
	},
	_const.Explores + "_guard_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "english translate",
	},
	_const.Explores + "_out_scout_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "english translate",
	},
	_const.Explores + "_in_scout_" + Target + "_1": {
		_const.RU: "Уже сдаешься? Это не бесплатная услуга!",
		_const.EN: "english translate",
	},
	_const.Explores + "_transport_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "english translate",
	},
	_const.Explores + "_builder_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "english translate",
	},
	_const.Explores + "_warrior_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "english translate",
	},
	_const.Explores + "_expedition_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "english translate",
	},

	_const.Explores + "_miner_" + Fear + "_1": {
		_const.RU: "Прекрати! Чего тебе надо?!",
		_const.EN: "english translate",
	},
	_const.Explores + "_guard_" + Fear + "_1": {
		_const.RU: "Не усугубляй!",
		_const.EN: "english translate",
	},
	_const.Explores + "_out_scout_" + Fear + "_1": {
		_const.RU: "А знаешь что? У меня есть дела поважнее!",
		_const.EN: "english translate",
	},
	_const.Explores + "_in_scout_" + Fear + "_1": {
		_const.RU: "Я все понял! Пожалуйста не разрушай мой транспорт!",
		_const.EN: "english translate",
	},
	_const.Explores + "_transport_" + Fear + "_1": {
		_const.RU: "Может мы как то договоримся?!",
		_const.EN: "english translate",
	},
	_const.Explores + "_builder_" + Fear + "_1": {
		_const.RU: "Может мы как то договоримся?!",
		_const.EN: "english translate",
	},
	_const.Explores + "_warrior_" + Fear + "_1": {
		_const.RU: "Советую тебе избегай встречи со мной!",
		_const.EN: "english translate",
	},
	_const.Explores + "_expedition_" + Fear + "_1": {
		_const.RU: "Не приближайся!",
		_const.EN: "english translate",
	},

	_const.Explores + "_miner_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "english translate",
	},
	_const.Explores + "_guard_" + Help + "_1": {
		_const.RU: "Не усугубляй!",
		_const.EN: "english translate",
	},
	_const.Explores + "_out_scout_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "english translate",
	},
	_const.Explores + "_in_scout_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "english translate",
	},
	_const.Explores + "_transport_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "english translate",
	},
	_const.Explores + "_builder_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "english translate",
	},
	_const.Explores + "_warrior_" + Help + "_1": {
		_const.RU: "Советую тебе избегай встречи со мной!",
		_const.EN: "english translate",
	},
	_const.Explores + "_expedition_" + Help + "_1": {
		_const.RU: "Не приближайся!",
		_const.EN: "english translate",
	},

	// Reverses Neutral
	_const.Reverses + "_miner_" + Neutral + "_1": {
		_const.RU: "Привет %UserName%. Я иду копать руду, не мешай мне.",
		_const.EN: "english translate",
	},
	_const.Reverses + "_guard_" + Neutral + "_1": {
		_const.RU: "Проезжай, не задерживайся.",
		_const.EN: "english translate",
	},
	_const.Reverses + "_out_scout_" + Neutral + "_1": {
		_const.RU: "Привет, я занимаюсь какой то важной херней, не мешай.",
		_const.EN: "english translate",
	},
	_const.Reverses + "_in_scout_" + Neutral + "_1": {
		_const.RU: "Чего тебе?",
		_const.EN: "english translate",
	},
	_const.Reverses + "_transport_" + Neutral + "_1": {
		_const.RU: "Не мешай, я занимаюсь торговлей",
		_const.EN: "english translate",
	},
	_const.Reverses + "_builder_" + Neutral + "_1": {
		_const.RU: "Опять кто то что то сломал, а мне чинить...",
		_const.EN: "english translate",
	},
	_const.Reverses + "_warrior_" + Neutral + "_1": {
		_const.RU: "Не мешай выполнению военных обязанностей.",
		_const.EN: "english translate",
	},
	_const.Reverses + "_expedition_" + Neutral + "_1": {
		_const.RU: "Наша задача осваивать пустоши, не мешай нам и будешь цел.",
		_const.EN: "english translate",
	},

	_const.Reverses + "_miner_" + Good + "_1": {
		_const.RU: "Кажется я нашел руду! Хочешь помочь мне ее разработать?",
		_const.EN: "english translate",
	},
	_const.Reverses + "_miner_" + Good + "_2": {
		_const.RU: "Похоже в этом секторе уже выкопали всю руду, лучше поискать ее где то еще.",
		_const.EN: "english translate",
	},
	_const.Reverses + "_guard_" + Good + "_1": {
		_const.RU: "Вас что то беспокоить?",
		_const.EN: "english translate",
	},
	_const.Reverses + "_out_scout_" + Good + "_1": {
		_const.RU: "Кажется в пустошах стало слишком много враждебных ботов, будь осторожен!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_in_scout_" + Good + "_1": {
		_const.RU: "Может храбанем когонить?",
		_const.EN: "english translate",
	},
	_const.Reverses + "_transport_" + Good + "_1": {
		_const.RU: "Я везу %ProductName% на базу %BaseName%, это будет выгодная сделка!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_transport_" + Good + "_2": {
		_const.RU: "Я следую на %BaseName%, база очень выгодно продает %ProductName%!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_transport_" + Good + "_3": {
		_const.RU: "Привет %UserName%, как дела?",
		_const.EN: "english translate",
	},
	_const.Reverses + "_builder_" + Good + "_1": {
		_const.RU: "Привет как дела?",
		_const.EN: "english translate",
	},
	_const.Reverses + "_warrior_" + Good + "_1": {
		_const.RU: "Мы отправляемся в сектор %SectorName%, можешь присоединиться!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_expedition_" + Good + "_1": {
		_const.RU: "Мы пробуем освоить секторе %SectorName%, будем рады содействию!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_expedition_" + Good + "_2": {
		_const.RU: "Наша задача выбирать противника с территории пустошь в секторе %SectorName%! Будем рады содействию!",
		_const.EN: "english translate",
	},

	_const.Reverses + "_miner_" + Bad + "_1": {
		_const.RU: "Ну, иди сюда, попробуй меня трахнуть — я тебя сам трахну!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_guard_" + Bad + "_1": {
		_const.RU: "Не усугубляй!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_out_scout_" + Bad + "_1": {
		_const.RU: "Лучше уходи пока можешь!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_in_scout_" + Bad + "_1": {
		_const.RU: "Приветствую. Хочешь внести плату за проезд?",
		_const.EN: "english translate",
	},
	_const.Reverses + "_transport_" + Bad + "_1": {
		_const.RU: "Если ты помешаешь мне торговать то я за себя не ручаюсь!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_builder_" + Bad + "_1": {
		_const.RU: "Отвали, ты мне не интересен.",
		_const.EN: "english translate",
	},
	_const.Reverses + "_warrior_" + Bad + "_1": {
		_const.RU: "Советую тебе избегай встречи со мной!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_expedition_" + Bad + "_1": {
		_const.RU: "Не приближайся!",
		_const.EN: "english translate",
	},

	_const.Reverses + "_miner_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_guard_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_out_scout_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_in_scout_" + Target + "_1": {
		_const.RU: "Уже сдаешься? Это не бесплатная услуга!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_transport_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_builder_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_warrior_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_expedition_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "english translate",
	},

	_const.Reverses + "_miner_" + Fear + "_1": {
		_const.RU: "Прекрати! Чего тебе надо?!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_guard_" + Fear + "_1": {
		_const.RU: "Не усугубляй!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_out_scout_" + Fear + "_1": {
		_const.RU: "А знаешь что? У меня есть дела поважнее!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_in_scout_" + Fear + "_1": {
		_const.RU: "Я все понял! Пожалуйста не разрушай мой транспорт!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_transport_" + Fear + "_1": {
		_const.RU: "Может мы как то договоримся?!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_builder_" + Fear + "_1": {
		_const.RU: "Может мы как то договоримся?!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_warrior_" + Fear + "_1": {
		_const.RU: "Советую тебе избегай встречи со мной!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_expedition_" + Fear + "_1": {
		_const.RU: "Не приближайся!",
		_const.EN: "english translate",
	},

	_const.Reverses + "_miner_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_guard_" + Help + "_1": {
		_const.RU: "Не усугубляй!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_out_scout_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_in_scout_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_transport_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_builder_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_warrior_" + Help + "_1": {
		_const.RU: "Советую тебе избегай встречи со мной!",
		_const.EN: "english translate",
	},
	_const.Reverses + "_expedition_" + Help + "_1": {
		_const.RU: "Не приближайся!",
		_const.EN: "english translate",
	},

	// APD
	_const.APD + "___1": {
		_const.RU: "ПУ ПУ ПУ...",
		_const.EN: "english translate",
	},

	// FarGod
	_const.FarGod + "___1": {
		_const.RU: "Скоро все изменится!",
		_const.EN: "english translate",
	},

	// FAUNA
	_const.FAUNA + "___1": {
		_const.RU: "*клац* <span class=\"importantly\">*клац*</span> *клац*, *щелк*?",
		_const.EN: "english translate",
	},
}

func GetNpcGreetings(fraction, role, relation, number string) map[string]string {
	textVariants := make([]map[string]string, 0)
	rand.Seed(time.Now().UnixNano())

	findKey := fraction + "_" + role + "_" + relation + number

	for key, text := range npcGreetings {
		if strings.Contains(key, findKey) {
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
