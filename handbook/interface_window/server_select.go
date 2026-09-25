package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

// ServerSelect - экран выбора сервера в Electron и этапы запуска локального мира (Veliri-private/docs/OFFLINE_MODE.md).
// Экран открывается до входа в игру, поэтому фронт читает это окно прямо из handbook.json, а не из стора.
var ServerSelect = map[string]map[string]string{
	"window_name": {
		_const.RU:   `Выбор сервера`,
		_const.EN:   `Server selection`,
		_const.ZhCN: `选择服务器`,
	},
	"local_game": {
		_const.RU:   `Локальная игра`,
		_const.EN:   `Local game`,
		_const.ZhCN: `本地游戏`,
	},
	"experimental": {
		_const.RU:   `Экспериментальная функция: игра запущена на локальном сервере.`,
		_const.EN:   `Experimental feature: The game is running on a local server.`,
		_const.ZhCN: `实验性功能：游戏在本地服务器上运行。`,
	},
	"coop_open": {
		_const.RU:   `Открыть для друзей в локальной сети`,
		_const.EN:   `Open to friends on the local network`,
		_const.ZhCN: `向局域网中的好友开放`,
	},
	"coop_password": {
		_const.RU:   `Пароль для друзей (необязательно)`,
		_const.EN:   `Password for friends (optional)`,
		_const.ZhCN: `好友密码（可选）`,
	},
	"play_local": {
		_const.RU:   `Играть локально`,
		_const.EN:   `Play locally`,
		_const.ZhCN: `本地游戏`,
	},
	"servers": {
		_const.RU:   `Серверы`,
		_const.EN:   `Servers`,
		_const.ZhCN: `服务器`,
	},
	"refresh": {
		_const.RU:   `обновить`,
		_const.EN:   `refresh`,
		_const.ZhCN: `刷新`,
	},
	"pinging": {
		_const.RU:   `Опрашиваем серверы...`,
		_const.EN:   `Pinging servers...`,
		_const.ZhCN: `正在检测服务器...`,
	},
	"list_unavailable": {
		_const.RU:   `Список серверов недоступен`,
		_const.EN:   `Server list is unavailable`,
		_const.ZhCN: `服务器列表不可用`,
	},
	"online": {
		_const.RU:   `онлайн`,
		_const.EN:   `online`,
		_const.ZhCN: `在线`,
	},
	"unavailable": {
		_const.RU:   `недоступен`,
		_const.EN:   `offline`,
		_const.ZhCN: `不可用`,
	},
	"connect_friend": {
		_const.RU:   `Подключиться к другу по адресу`,
		_const.EN:   `Connect to a friend by address`,
		_const.ZhCN: `通过地址连接好友`,
	},
	"friend_password": {
		_const.RU:   `Пароль (если задан)`,
		_const.EN:   `Password (if set)`,
		_const.ZhCN: `密码（如已设置）`,
	},
	"connect": {
		_const.RU:   `Подключиться`,
		_const.EN:   `Connect`,
		_const.ZhCN: `连接`,
	},
	"connect_failed": {
		_const.RU:   `Не удалось подключиться: %error%`,
		_const.EN:   `Connection failed: %error%`,
		_const.ZhCN: `连接失败：%error%`,
	},
	"preview_only": {
		_const.RU:   `Предпросмотр: подключение работает только в клиенте игры`,
		_const.EN:   `Preview: connecting works only in the game client`,
		_const.ZhCN: `预览：仅在游戏客户端中可以连接`,
	},
	"address_format": {
		_const.RU:   `Адрес в виде 192.168.0.10:18090`,
		_const.EN:   `Address like 192.168.0.10:18090`,
		_const.ZhCN: `地址格式如 192.168.0.10:18090`,
	},
	"access_denied": {
		_const.RU:   `Сервер не пустил: неверный пароль или доступ закрыт`,
		_const.EN:   `Access denied: wrong password or server is closed`,
		_const.ZhCN: `访问被拒绝：密码错误或服务器已关闭`,
	},
	"login_failed": {
		_const.RU:   `Не удалось войти на сервер`,
		_const.EN:   `Failed to log in to the server`,
		_const.ZhCN: `无法登录服务器`,
	},

	// этапы запуска локального сервера/мира (offline:status из Electron + login/enter из Gate)
	"status_init_db": {
		_const.RU:   `Первый запуск: создаём базу данных...`,
		_const.EN:   `First launch: creating the database...`,
		_const.ZhCN: `首次启动：正在创建数据库...`,
	},
	"status_start_db": {
		_const.RU:   `Запускаем базу данных...`,
		_const.EN:   `Starting the database...`,
		_const.ZhCN: `正在启动数据库...`,
	},
	"status_restore_static": {
		_const.RU:   `Разворачиваем игровые данные...`,
		_const.EN:   `Unpacking game data...`,
		_const.ZhCN: `正在解包游戏数据...`,
	},
	"status_restore_game": {
		_const.RU:   `Создаём новый мир...`,
		_const.EN:   `Creating a new world...`,
		_const.ZhCN: `正在创建新世界...`,
	},
	"status_create_world": {
		_const.RU:   `Создаём новый мир...`,
		_const.EN:   `Creating a new world...`,
		_const.ZhCN: `正在创建新世界...`,
	},
	"status_stop_world": {
		_const.RU:   `Останавливаем текущий мир...`,
		_const.EN:   `Stopping the current world...`,
		_const.ZhCN: `正在停止当前世界...`,
	},
	"status_load_save": {
		_const.RU:   `Загружаем сохранение...`,
		_const.EN:   `Loading the save...`,
		_const.ZhCN: `正在加载存档...`,
	},
	"status_start_master": {
		_const.RU:   `Запускаем сервер...`,
		_const.EN:   `Starting the server...`,
		_const.ZhCN: `正在启动服务器...`,
	},
	"status_start_node": {
		_const.RU:   `Запускаем мир...`,
		_const.EN:   `Starting the world...`,
		_const.ZhCN: `正在启动世界...`,
	},
	"status_ready": {
		_const.RU:   `Входим...`,
		_const.EN:   `Logging in...`,
		_const.ZhCN: `正在登录...`,
	},
	"status_login": {
		_const.RU:   `Входим в мир...`,
		_const.EN:   `Entering the world...`,
		_const.ZhCN: `正在进入世界...`,
	},
	"status_enter": {
		_const.RU:   `Выходим на карту...`,
		_const.EN:   `Entering the map...`,
		_const.ZhCN: `正在进入地图...`,
	},
}
