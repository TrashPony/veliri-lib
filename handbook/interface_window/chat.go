package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Chat = map[string]map[string]string{
	// Chat
	"window_name_1": {
		_const.RU:   `Чат`,
		_const.EN:   `Chat`,
		_const.ZhCN: `聊天`,
	},
	"local": {
		_const.RU:   `Локальный`,
		_const.EN:   `Local`,
		_const.ZhCN: `本地`,
	},
	"local_greetings": {
		_const.RU:   `Вы входите на территорию `,
		_const.EN:   `You enter the territory `,
		_const.ZhCN: `您进入了领土 `,
	},
	"local_greetings_wasteland": {
		_const.RU:   `Пустоши...`,
		_const.EN:   `Wasteland...`,
		_const.ZhCN: `荒地...`,
	},
	// CreateNewGroup
	"window_name_2": {
		_const.RU:   `Создание нового канала`,
		_const.EN:   `Creating a new channel`,
		_const.ZhCN: `创建新频道`,
	},
	"text_1": {
		_const.RU:   `Имя канала...`,
		_const.EN:   `Channel name...`,
		_const.ZhCN: `频道名称...`,
	},
	"text_2": {
		_const.RU:   `Пароль (если пусто то без пароля)`,
		_const.EN:   `Password (if empty then no password)`,
		_const.ZhCN: `密码（如果为空则无密码）`,
	},
	"text_3": {
		_const.RU:   `Загрузить`,
		_const.EN:   `Download`,
		_const.ZhCN: `下载`,
	},
	"text_4": {
		_const.RU:   `Приветственное сообщение`,
		_const.EN:   `Welcome message`,
		_const.ZhCN: `欢迎消息`,
	},
	"button_1": {
		_const.RU:   `Создать`,
		_const.EN:   `Create`,
		_const.ZhCN: `创建`,
	},
	"button_2": {
		_const.RU:   `Закрыть`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
	// Messages
	"chanel_local": {
		_const.RU:   `локация`,
		_const.EN:   `local`,
		_const.ZhCN: `本地`,
	},
	"chanel_group": {
		_const.RU:   `группа`,
		_const.EN:   `group`,
		_const.ZhCN: `组`,
	},
	"chanel_battle": {
		_const.RU:   `бой`,
		_const.EN:   `battle`,
		_const.ZhCN: `战斗`,
	},
	"chanel_corporation": {
		_const.RU:   `кластер`,
		_const.EN:   `corporation`,
		_const.ZhCN: `公司`,
	},
	"chanel_private": {
		_const.RU:   `личка`,
		_const.EN:   `private`,
		_const.ZhCN: `私聊`,
	},
	"chanel_unknown": {
		_const.RU:   `Неизвестный`,
		_const.EN:   `Unknown`,
		_const.ZhCN: `未知`,
	},
	// MessagesWrapper
	"default_greetings": {
		_const.RU:   `Добро пожаловать!`,
		_const.EN:   `Welcome!`,
		_const.ZhCN: `欢迎！`,
	},
	// UserLine
	"text_5": {
		_const.RU:   `Выберите роль`,
		_const.EN:   `Select a role`,
		_const.ZhCN: `选择角色`,
	},
	"text_6": {
		_const.RU:   `Место:`,
		_const.EN:   `Place:`,
		_const.ZhCN: `地点：`,
	},
	// UserSubMenu
	"sub_1": {
		_const.RU:   `Подробнее`,
		_const.EN:   `More details`,
		_const.ZhCN: `更多详情`,
	},
	"sub_2": {
		_const.RU:   `Написать`,
		_const.EN:   `Write to chat`,
		_const.ZhCN: `发送消息`,
	},
	"sub_3": {
		_const.RU:   `Отправить письмо`,
		_const.EN:   `Send mail`,
		_const.ZhCN: `发送邮件`,
	},
	"sub_4": {
		_const.RU:   `Сделать лидером`,
		_const.EN:   `Make a leader`,
		_const.ZhCN: `设为领袖`,
	},
	"sub_5": {
		_const.RU:   `Пригласить в группу`,
		_const.EN:   `Invite to group`,
		_const.ZhCN: `邀请加入组`,
	},
	"sub_6": {
		_const.RU:   `Добавить в друзья`,
		_const.EN:   `Add as Friend`,
		_const.ZhCN: `添加好友`,
	},
	"sub_7": {
		_const.RU:   `Добавить контакт`,
		_const.EN:   `Add contact`,
		_const.ZhCN: `添加联系人`,
	},
	"sub_8": {
		_const.RU:   `Выгнать`,
		_const.EN:   `Kick out`,
		_const.ZhCN: `踢出`,
	},
	"sub_9": {
		_const.RU:   `Покинуть`,
		_const.EN:   `Leave`,
		_const.ZhCN: `离开`,
	},
	"sub_10": {
		_const.RU:   `Перевод кредитов`,
		_const.EN:   `Transfer of credits`,
		_const.ZhCN: `转账`,
	},
	"sub_11": {
		_const.RU:   `Закрыть`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
	// ViewAllGroups
	"button_3": {
		_const.RU:   `Отмена`,
		_const.EN:   `Cancel`,
		_const.ZhCN: `取消`,
	},
	"button_4": {
		_const.RU:   `Войти`,
		_const.EN:   `Enter`,
		_const.ZhCN: `进入`,
	},
	"placeholder_1": {
		_const.RU:   `введите пароль`,
		_const.EN:   `enter the password`,
		_const.ZhCN: `输入密码`,
	},
	"text_7": {
		_const.RU:   `Основные группы:`,
		_const.EN:   `Main groups:`,
		_const.ZhCN: `主要组：`,
	},
	"text_8": {
		_const.RU:   `Пользовательские группы:`,
		_const.EN:   `Player groups:`,
		_const.ZhCN: `玩家组：`,
	},
	"button_5": {
		_const.RU:   `Создать`,
		_const.EN:   `Create`,
		_const.ZhCN: `创建`,
	},
	"button_6": {
		_const.RU:   `Закрыть`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
}
