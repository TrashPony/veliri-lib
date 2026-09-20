package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var GroupMenu = map[string]map[string]string{
	"window_name": {
		_const.RU:   `Отряд`,
		_const.EN:   `Squad`,
		_const.ZhCN: `队`,
	},
	"text_1": {
		_const.RU:   `Вы не состоите в отряде...`,
		_const.EN:   `You are not a member of the squad...`,
		_const.ZhCN: `你不是队伍的一员...`,
	},
	"party": {
		_const.RU:   `Пати`,
		_const.EN:   `Party`,
		_const.ZhCN: `队伍`,
	},
	"team": {
		_const.RU:   `Команда`,
		_const.EN:   `Team`,
		_const.ZhCN: `团队`,
	},
}
