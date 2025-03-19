package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Gate = map[string]map[string]string{
	"window_name": {
		_const.RU:   `Вход`,
		_const.EN:   `Entry`,
		_const.ZhCN: `入口`,
	},
	"attention_1": {
		_const.RU: `<h3 class="not_card_notify">Внимание!</h3>
						<p class="not_card_notify_text">Похоже в вашем браузере отключено
						  <a class="not_card_notify_text_link" target="_blank"
							 href="https://www.google.com/search?q=%D0%B0%D0%BF%D0%BF%D0%B0%D1%80%D0%B0%D1%82%D0%BD%D0%BE%D0%B5+%D1%83%D1%81%D0%BA%D0%BE%D1%80%D0%B5%D0%BD%D0%B8%D0%B5+%D0%B1%D1%80%D0%B0%D1%83%D0%B7%D0%B5%D1%80%D0%B0">
							аппаратное ускорение</a>, без него может сильно снизится производительность.</p>`,
		_const.EN: `<h3 class="not_card_notify">Attention!</h3>
						<p class="not_card_notify_text">It seems that hardware acceleration is disabled in your browser. Performance may be significantly reduced without it.</p>`,
		_const.ZhCN: `<h3 class="not_card_notify">注意！</h3>
						<p class="not_card_notify_text">您的浏览器似乎禁用了硬件加速。没有它，性能可能会显著降低。</p>`,
	},
	"play": {
		_const.RU:   `Играть`,
		_const.EN:   `Play`,
		_const.ZhCN: `开始游戏`,
	},
	"close": {
		_const.RU:   `Выход`,
		_const.EN:   `Exit`,
		_const.ZhCN: `退出`,
	},
	"createPlayerLarchName": {
		_const.RU:   `В имени персонажа не может быть больше 16ти символов.`,
		_const.EN:   `A character's name cannot contain more than 16 characters.`,
		_const.ZhCN: `角色名称不能超过16个字符。`,
	},
	"createPlayerNameNotAvailable": {
		_const.RU:   `Имя уже занято.`,
		_const.EN:   `The name is already taken.`,
		_const.ZhCN: `该名称已被占用。`,
	},
	"createPlayerSmallName": {
		_const.RU:   `В имени персонажа не может быть меньше 3х символов.`,
		_const.EN:   `A character's name must be at least 3 characters long.`,
		_const.ZhCN: `角色名称必须至少包含3个字符。`,
	},
	"create": {
		_const.RU:   `Создать`,
		_const.EN:   `Create`,
		_const.ZhCN: `创建`,
	},
	"back": {
		_const.RU:   `Назад`,
		_const.EN:   `Back`,
		_const.ZhCN: `返回`,
	},
	"enter_name": {
		_const.RU:   `Введите имя`,
		_const.EN:   `Enter your name`,
		_const.ZhCN: `输入名称`,
	},
	"player_name": {
		_const.RU:   `Имя персонажа`,
		_const.EN:   "Character `s name",
		_const.ZhCN: `角色名称`,
	},
}
