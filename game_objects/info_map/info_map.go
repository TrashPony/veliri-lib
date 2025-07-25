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
	OccupiedFraction    string                   `json:"occupied_fraction"`
	PossibleBattle      bool                     `json:"possible_battle"`
	Battle              int                      `json:"battle"`
	FreeLand            bool                     `json:"free_land"`
	HandlersCoordinates []*coordinate.Coordinate `json:"handlers_coordinates"`
	CorporationID       int                      `json:"corporation_id"`
	Transit             bool                     `json:"transit"`
	IsIndoors           bool                     `json:"is_indoors"`
	GameTime            gameTime                 `json:"-"`
}

type gameTime struct {
	Hours   int
	Minutes int
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
