package info_map

import "github.com/TrashPony/veliri-lib/game_objects/coordinate"

type InfoMap struct {
	Id                  int                      `json:"id"`
	Name                string                   `json:"name"`
	XSize               int                      `json:"x_size"`
	YSize               int                      `json:"y_size"`
	Specification       string                   `json:"specification"`
	Global              bool                     `json:"global"`
	XGlobal             int                      `json:"x_global"`
	YGlobal             int                      `json:"y_global"`
	Fraction            string                   `json:"fraction"`
	PossibleBattle      bool                     `json:"possible_battle"`
	Battle              bool                     `json:"battle"`
	FreeLand            bool                     `json:"free_land"`
	HandlersCoordinates []*coordinate.Coordinate `json:"handlers_coordinates"`
}

func (mp *InfoMap) GetEntryTySector(sectorID int) *coordinate.Coordinate {
	for _, entry := range mp.HandlersCoordinates {
		if entry.Handler == "sector" && entry.ToMapID == sectorID {
			return entry
		}
	}
	return nil
}

func (mp *InfoMap) GetAllEntrySectors() []*coordinate.Coordinate {
	entrySectors := make([]*coordinate.Coordinate, 0)
	for _, entry := range mp.HandlersCoordinates {
		if entry.Handler == "sector" {
			entrySectors = append(entrySectors, entry)
		}
	}

	return entrySectors
}
