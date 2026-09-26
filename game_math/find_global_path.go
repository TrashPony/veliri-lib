package game_math

import (
	"github.com/TrashPony/veliri-lib/game_objects/coordinate"
	"github.com/TrashPony/veliri-lib/game_objects/info_map"
	"strconv"
	"sync"
)

type SearchMap struct {
	ID     string
	MapID  int
	Map    *info_map.InfoMap
	Wave   int
	F      float64
	G      float64 // длина пути по картам от старта до точки появления на этой карте (px) + штраф за каждый переход
	Parent *SearchMap
	Entry  *coordinate.Coordinate
}

// globalJumpCost штраф за переход между секторами (px): при почти равной длине выбирается путь с меньшим числом переходов
const globalJumpCost = 300.0

var cache = &cachePaths{}

// поиск по графам, нормально не отдебажан
type cachePaths struct {
	cachePath map[string]*cachePath
	mx        sync.RWMutex
}

type cachePath struct {
	maps []*SearchMap
	path []*coordinate.Coordinate
}

func RemoveCacheMap() {
	cache.mx.Lock()
	defer cache.mx.Unlock()
	cache.cachePath = make(map[string]*cachePath)
}

func (c *cachePaths) addPath(key string, maps []*SearchMap, path []*coordinate.Coordinate) {
	c.mx.Lock()
	defer c.mx.Unlock()

	if c.cachePath == nil {
		c.cachePath = make(map[string]*cachePath)
	}

	c.cachePath[key] = &cachePath{
		maps: maps,
		path: path,
	}
}

func (c *cachePaths) getPath(key string) *cachePath {
	c.mx.RLock()
	defer c.mx.RUnlock()
	return c.cachePath[key]
}

/*
typePath:
NotBattle - по всем секторам кроме боевых
NotFree - по всем секторам кроме свободных
Fraction - по всем фракционным и безопасным
OnlyFraction - по всем только фракционным
*/

func FindGlobalPath(store map[int]*info_map.InfoMap, startSectorID, endSectorID int, typePath, fraction string,
	questHandler map[int][]*coordinate.Coordinate, pID, cID int, pFraction string) ([]*SearchMap, []*coordinate.Coordinate) { // возращает ячейки пеереходов из сектора в сектор

	// ключ кеша учитывает все, от чего зависит путь: тип пути, фракцию и доступ (игрок, корпорация, фракция игрока). Раньше ключом был
	// только старт:финиш, и путь, найденный для одного типа (или с доступом одного игрока), отдавался всем. Квестовые переходы
	// (questHandler) у каждого свои - такие пути не кешируются
	cacheKey := strconv.Itoa(startSectorID) + ":" + strconv.Itoa(endSectorID) + ":" + typePath + ":" + fraction + ":" +
		strconv.Itoa(pID) + ":" + strconv.Itoa(cID) + ":" + pFraction
	if len(questHandler) == 0 {
		if cPath := cache.getPath(cacheKey); cPath != nil {
			return cPath.maps, cPath.path
		}
	}

	startSector := store[startSectorID]
	endSector := store[endSectorID]

	if startSector == nil || endSector == nil {
		return nil, nil
	}

	var path []*SearchMap
	var noSortedPath []*SearchMap

	start := &SearchMap{
		ID:    strconv.Itoa(startSector.Id),
		MapID: startSector.Id,
		Map:   startSector,
		Entry: &coordinate.Coordinate{X: startSector.XSize / 2, Y: startSector.YSize / 2}}

	end := &SearchMap{
		ID:    strconv.Itoa(endSector.Id),
		MapID: endSector.Id,
		Map:   endSector}

	if startSectorID == endSectorID {
		path = append(path, end)
		return path, nil
	}

	// Дейкстра по длине пути в пикселях: узел - карта + переход, которым на нее попали (точка появления), цена ребра - путь внутри
	// карты от точки появления до следующего перехода + штраф за переход. Раньше поиск шел по волнам (число переходов), расстояние
	// внутри секторов не учитывалось, а равные по числу переходов пути выбирались случайно (порядок обхода map) - маршрут
	// выходил длинным и мог отличаться от вызова к вызову
	openPoints, closePoints := make(map[string]*SearchMap), make(map[string]*SearchMap)
	openPoints[start.ID] = start

	wave := 0

	for {

		if len(openPoints) == 0 {
			return nil, nil
		}

		wave++
		current := getOpenPoint(openPoints) // точка с минимальной длиной пути
		if current.MapID == end.MapID {     // если текущая точка и есть конец начинаем генерить путь
			for !(current.MapID == start.MapID) {
				current = current.Parent
				if !(current.MapID == start.MapID) {
					// если текущая точка попрежнему не стартовая то добавляем ее в путь
					noSortedPath = append(noSortedPath, current)
				}
			}
			break
		}

		parseNeighbours(current, openPoints, closePoints, store, end, wave, typePath, fraction, questHandler, pID, cID, pFraction)
	}

	// сразу добавим в путь стартовую точку т.к. нам с нее нужен будет переход
	path = append(path, start)
	for i := len(noSortedPath); i > 0; i-- {
		path = append(path, noSortedPath[i-1])
	}
	// и послуюднюю что бы знать куда прыгать в конце пути
	path = append(path, end)

	var transitionPoints []*coordinate.Coordinate
	for i := 0; i < len(path); i++ {
		if i+1 < len(path) {
			transitionPoint := GetEntryTySector(path[i].Map, questHandler, path[i+1].MapID)
			transitionPoint.MapID = path[i].Map.Id
			transitionPoints = append(transitionPoints, transitionPoint)
		}
	}

	if len(questHandler) == 0 {
		cache.addPath(cacheKey, path, transitionPoints)
	}
	return path, transitionPoints
}

func parseNeighbours(current *SearchMap, openPoints, closePoints map[string]*SearchMap, store map[int]*info_map.InfoMap,
	end *SearchMap, wave int, typePath, fraction string, questHandler map[int][]*coordinate.Coordinate, pID, cID int, pFraction string) {

	delete(openPoints, current.ID)    // удаляем ячейку из не посещенных
	closePoints[current.ID] = current // добавляем в массив посещенные

	// надо взять все переходы и это будут соседи
	entrySectors := GetAllEntrySectors(current.Map, questHandler)

	for _, entry := range entrySectors {
		mp := store[entry.ToMapID]

		if mp == nil {
			continue
		}

		if end.MapID != mp.Id {

			// проверка доступа по клану и игроку
			if entry.AccessManager.Access && !(entry.AccessManager.GetAccess("corporation"+strconv.Itoa(cID), "player"+strconv.Itoa(pID), "fraction"+pFraction)) {
				continue
			}

			if mp.Fraction != fraction && typePath == "OnlyFraction" {
				continue
			}

			if mp.PossibleBattle && typePath == "NotBattle" {
				continue
			}

			if mp.FreeLand && typePath == "NotFree" {
				continue
			}

			if typePath == "Fraction" && mp.PossibleBattle && mp.Fraction != fraction {
				continue
			}
		}

		// в стартовый сектор не возвращаемся: восстановление пути идет до первой встречи стартовой карты
		root := current
		for root.Parent != nil {
			root = root.Parent
		}
		if mp.Id == root.MapID {
			continue
		}

		id := strconv.Itoa(mp.Id) + entry.Key()
		if closePoints[id] != nil {
			continue
		}

		// путь внутри текущей карты: от точки, где на нее попали, до этого перехода
		ax, ay := arrivalPoint(current, questHandler)
		g := current.G + GetBetweenDist(ax, ay, entry.X, entry.Y) + globalJumpCost

		if old := openPoints[id]; old != nil && old.G <= g {
			continue
		}

		openPoints[id] = &SearchMap{
			ID:     id,
			MapID:  mp.Id,
			Map:    mp,
			Parent: current,
			Entry:  entry,
			Wave:   wave,
			G:      g,
			F:      g,
		}
	}
}

// arrivalPoint где оказываемся на карте узла: старт - центр карты старта, иначе точка появления перехода (Positions), если ее нет -
// переход обратно на карту, с которой пришли (появляемся рядом с ним), иначе центр карты
func arrivalPoint(n *SearchMap, questHandler map[int][]*coordinate.Coordinate) (int, int) {
	if n.Parent == nil {
		if n.Entry != nil {
			return n.Entry.X, n.Entry.Y
		}
		return n.Map.XSize / 2, n.Map.YSize / 2
	}

	if n.Entry != nil && len(n.Entry.Positions) > 0 && n.Entry.Positions[0] != nil {
		return n.Entry.Positions[0].X, n.Entry.Positions[0].Y
	}

	if back := GetEntryTySector(n.Map, questHandler, n.Parent.MapID); back != nil {
		return back.X, back.Y
	}

	return n.Map.XSize / 2, n.Map.YSize / 2
}

// getOpenPoint открытая точка с минимальной длиной пути; при равенстве - по ID, что бы результат не зависел от порядка обхода map
func getOpenPoint(openMaps map[string]*SearchMap) *SearchMap {
	var minMap *SearchMap
	for _, p := range openMaps {
		if minMap == nil || p.G < minMap.G || (p.G == minMap.G && p.ID < minMap.ID) {
			minMap = p
		}
	}

	return minMap
}

func GetEntryTySector(mp *info_map.InfoMap, questHandler map[int][]*coordinate.Coordinate, sectorID int) *coordinate.Coordinate {
	point := mp.GetEntryTySector(sectorID)
	if point != nil {
		return point
	}

	if questHandler != nil && len(questHandler[mp.Id]) > 0 {
		for _, entry := range questHandler[mp.Id] {
			if entry.Handler == "sector" && entry.ToMapID == sectorID {
				return entry
			}
		}
	}

	return nil
}

func GetAllEntrySectors(mp *info_map.InfoMap, questHandler map[int][]*coordinate.Coordinate) []*coordinate.Coordinate {
	handlers := mp.GetAllEntrySectors()

	if questHandler != nil && len(questHandler[mp.Id]) > 0 {
		handlers = append(handlers, questHandler[mp.Id]...)
	}

	return handlers
}
