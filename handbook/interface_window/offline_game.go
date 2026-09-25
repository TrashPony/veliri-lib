package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

// OfflineGame - локальный мир в игре: сохранения в окне персонажа Gate, экран смены мира, адрес для кооп-гостей,
// пункты Esc-меню (сохранение, пауза). Veliri-private/docs/OFFLINE_MODE.md.
var OfflineGame = map[string]map[string]string{
	"autosaves": {
		_const.RU:   `Автосохранения`,
		_const.EN:   `Autosaves`,
		_const.ZhCN: `自动存档`,
	},
	"autosaves_hint": {
		_const.RU:   `При заходе на базу, последние 3`,
		_const.EN:   `On entering a base, last 3`,
		_const.ZhCN: `进入基地时保存，保留最近3个`,
	},
	"saves": {
		_const.RU:   `Сохранения`,
		_const.EN:   `Saves`,
		_const.ZhCN: `存档`,
	},
	"saves_hint": {
		_const.RU:   `Сохраниться можно в игре через меню (Esc)`,
		_const.EN:   `Save in game via the menu (Esc)`,
		_const.ZhCN: `可在游戏中通过菜单（Esc）存档`,
	},
	"untitled": {
		_const.RU:   `без названия`,
		_const.EN:   `untitled`,
		_const.ZhCN: `未命名`,
	},
	"load": {
		_const.RU:   `загрузить`,
		_const.EN:   `load`,
		_const.ZhCN: `读取`,
	},
	"sure": {
		_const.RU:   `точно?`,
		_const.EN:   `sure?`,
		_const.ZhCN: `确定？`,
	},
	"world_starts_after": {
		_const.RU:   `Мир персонажа запустится после выбора`,
		_const.EN:   `The character world starts after you choose`,
		_const.ZhCN: `选择后将启动该角色的世界`,
	},
	"lan_address": {
		_const.RU:   `Адрес для друзей`,
		_const.EN:   `Address for friends`,
		_const.ZhCN: `好友连接地址`,
	},
	"title_load_save": {
		_const.RU:   `Загрузка сохранения`,
		_const.EN:   `Loading save`,
		_const.ZhCN: `正在读取存档`,
	},
	"title_start_world": {
		_const.RU:   `Запуск мира`,
		_const.EN:   `Starting world`,
		_const.ZhCN: `正在启动世界`,
	},
	"title_new_world": {
		_const.RU:   `Новый мир`,
		_const.EN:   `New world`,
		_const.ZhCN: `新世界`,
	},
	"start_failed": {
		_const.RU:   `Не удалось запустить мир: %error%`,
		_const.EN:   `Failed to start the world: %error%`,
		_const.ZhCN: `无法启动世界：%error%`,
	},
	"enter_failed": {
		_const.RU:   `Не удалось войти в мир`,
		_const.EN:   `Failed to enter the world`,
		_const.ZhCN: `无法进入世界`,
	},
	"not_responding": {
		_const.RU:   `Мир не отвечает`,
		_const.EN:   `The world is not responding`,
		_const.ZhCN: `世界没有响应`,
	},
	"close": {
		_const.RU:   `Закрыть`,
		_const.EN:   `Close`,
		_const.ZhCN: `关闭`,
	},
	"pause": {
		_const.RU:   `Пауза`,
		_const.EN:   `Pause`,
		_const.ZhCN: `暂停`,
	},
	"save_game": {
		_const.RU:   `Сохранить игру`,
		_const.EN:   `Save game`,
		_const.ZhCN: `保存游戏`,
	},
	"save_name": {
		_const.RU:   `Название сохранения`,
		_const.EN:   `Save name`,
		_const.ZhCN: `存档名称`,
	},
	"save": {
		_const.RU:   `Сохранить`,
		_const.EN:   `Save`,
		_const.ZhCN: `保存`,
	},
	"saving": {
		_const.RU:   `Сохраняем...`,
		_const.EN:   `Saving...`,
		_const.ZhCN: `正在保存...`,
	},
	"saved": {
		_const.RU:   `Сохранено`,
		_const.EN:   `Saved`,
		_const.ZhCN: `已保存`,
	},
	"save_failed": {
		_const.RU:   `Не удалось сохранить`,
		_const.EN:   `Save failed`,
		_const.ZhCN: `保存失败`,
	},
	"local_world": {
		_const.RU:   `Локальный мир`,
		_const.EN:   `Local world`,
		_const.ZhCN: `本地世界`,
	},
	"debug_mode": {
		_const.RU:   `Режим отладки`,
		_const.EN:   `Debug mode`,
		_const.ZhCN: `调试模式`,
	},
	"debug_open": {
		_const.RU:   `Открыть панель`,
		_const.EN:   `Open panel`,
		_const.ZhCN: `打开面板`,
	},
	"save_dungeon": {
		_const.RU:   `На данже сохраниться нельзя`,
		_const.EN:   `You cannot save in a dungeon`,
		_const.ZhCN: `在副本中无法存档`,
	},
}
