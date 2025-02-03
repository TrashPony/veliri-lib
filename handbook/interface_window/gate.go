package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Gate = map[string]map[string]string{
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
	"close": {
		_const.RU: `Выход`,
		_const.EN: `Exit`,
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
		_const.RU: `Назад`,
		_const.EN: `Back`,
	},
	"enter_name": {
		_const.RU: `Введите имя`,
		_const.EN: `Enter your name`,
	},
	"player_name": {
		_const.RU: `Имя персонажа`,
		_const.EN: "Character `s name",
	},
}
