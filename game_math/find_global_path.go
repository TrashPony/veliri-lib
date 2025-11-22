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
	Parent *SearchMap
	Entry  *coordinate.Coordinate
}

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

	cPath := cache.getPath(strconv.Itoa(startSectorID) + ":" + strconv.Itoa(endSectorID))
	if cPath != nil {
		return cPath.maps, cPath.path
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

	openPoints, closePoints := make(map[string]*SearchMap), make(map[string]*SearchMap) // создаем 2 карты для посещенных (open) и непосещеных (close) точек
	openPoints[start.ID] = start                                                        // кладем в карту посещенных точек стартовую точку

	// перменная добавляет в цену номер волны
	wave := 0

	for {

		if len(openPoints) == 0 {
			return nil, nil
		}

		wave++
		current := getOpenPoint(openPoints) // Берем точку с мин стоимостью пути
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

	cache.addPath(strconv.Itoa(startSectorID)+":"+strconv.Itoa(endSectorID), path, transitionPoints)
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

		// проверяем что карта существует, и что мы ее уже не обработали
		if closePoints[strconv.Itoa(mp.Id)+entry.Key()] != nil || openPoints[strconv.Itoa(mp.Id)+entry.Key()] != nil {
			continue
		}

		openPoints[strconv.Itoa(mp.Id)+entry.Key()] = &SearchMap{
			ID:     strconv.Itoa(mp.Id) + entry.Key(),
			MapID:  mp.Id,
			Map:    mp,
			Parent: current,
			Entry:  entry,
		}

		// добавяем номер волны что бы все последующие волны имели большую цену
		openPoints[strconv.Itoa(mp.Id)+entry.Key()].Wave = wave

		if current.Entry.Positions != nil && len(current.Entry.Positions) > 0 {
			// учитывать растояние от тп до выхода из него, todo походу это не работает
			openPoints[strconv.Itoa(mp.Id)+entry.Key()].F = GetBetweenDist(entry.X, entry.Y, current.Entry.Positions[0].X, current.Entry.Positions[0].Y)
		}
	}
}

func getOpenPoint(openMaps map[string]*SearchMap) *SearchMap {
	minWall := -1

	var minMap *SearchMap

	for _, p := range openMaps {
		if p.Wave < minWall || minWall < 0 {
			minWall = p.Wave
		}
	}

	minF := -1.0
	for _, p := range openMaps {

		if p.Wave != minWall {
			continue
		}

		if p.F < minF || minF < 0 {
			minMap = p
			minF = p.F
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
