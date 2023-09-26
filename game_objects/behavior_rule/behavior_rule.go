package behavior_rule

import (
	"encoding/json"
	"fmt"
	"github.com/TrashPony/veliri-lib/game_objects/coordinate"
	"sync"
	"time"
)

type BehaviorRules struct {
	Rules []*BehaviorRule `json:"rules"`
	Meta  *Meta           `json:"meta"`
}

type BehaviorRule struct {
	Action   string        `json:"action"`
	Meta     *Meta         `json:"meta"`
	PassRule *BehaviorRule `json:"access_rule"`
	StopRule *BehaviorRule `json:"stop_rule"`
	Exit     bool          `json:"exit"`
}

type Patrol struct {
	Path       []*coordinate.Coordinate `json:"path"`        // маршрут состоящий из x,y и радиуса
	ToIDIndex  int                      `json:"to_id_index"` // текущая цель куда двигается группа, по достижение происходит ++, если -1
	AutoChange bool                     `json:"auto_change"` // true, ToIDIndex меняется автоматически как только достигнута цель
}

type SubGroup struct {
	members       []int                  // TODO
	Leader        int                    `json:"leader"`
	Radius        int                    `json:"radius"`
	RegroupRadius int                    `json:"regroup_radius"`
	Regroup       bool                   `json:"regroup"`
	RegroupMapID  int                    `json:"regroup_map_id"`
	NextMapID     int                    `json:"next_map_id"`
	RegroupPos    *coordinate.Coordinate `json:"regroup_pos"`
	Name          string                 `json:"name"`
	mx            sync.RWMutex           `json:"mx"`
}

func (s *SubGroup) AddMember(id int) {
	s.mx.Lock()
	defer s.mx.Unlock()

	if s.members == nil {
		s.members = make([]int, 0)
	}

	s.members = append(s.members, id)
}

func (s *SubGroup) SetMember(ids []int) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.members = ids
}

func (s *SubGroup) RemoveMember(id int) {
	s.mx.Lock()
	defer s.mx.Unlock()

	index := -1
	for i, m := range s.members {
		if m == id {
			index = i
			break
		}
	}

	if index >= 0 {
		s.members = append(s.members[:index], s.members[index+1:]...)
	}
}

func (s *SubGroup) GetMembers() <-chan int {
	s.mx.RLock()

	ids := make(chan int, len(s.members))

	go func() {
		defer func() {
			close(ids)
			s.mx.RUnlock()
		}()

		for _, id := range s.members {
			ids <- id
		}
	}()

	return ids
}

func (s *SubGroup) UnsafeMembers() ([]int, *sync.RWMutex) {
	return s.members, &s.mx
}

func (s *SubGroup) GetCountMembers() int {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return len(s.members)
}

func (s *SubGroup) GetCopyMembers() []int {
	s.mx.RLock()
	defer s.mx.RUnlock()

	ids := make([]int, 0, len(s.members))
	for _, id := range s.members {
		ids = append(ids, id)
	}

	return ids
}

func (s *SubGroup) LoadMembers(data []byte) {
	s.mx.Lock()
	defer s.mx.Unlock()

	ids := make([]int, 0)
	err := json.Unmarshal(data, &ids)
	if err != nil {
		fmt.Println("error load subgroup members")
	}

	s.members = ids
}

type Meta struct {
	Radius                        int               `json:"radius"`
	Parameter                     string            `json:"parameter"`
	Count                         int               `json:"count"`
	Count2                        int               `json:"count_2"`
	Current                       int               `json:"current"`
	Percent                       bool              `json:"percent"`
	MapID                         int               `json:"map_id"`
	BaseID                        int               `json:"base_id"`
	TypeTarget                    string            `json:"type_target"`
	IDTarget                      int               `json:"id_target"`
	IDPlayerTarget                int               `json:"id_player_target"`
	TypeInteraction               string            `json:"type_interaction"`
	TypeID                        int               `json:"type_id"`
	X                             int               `json:"x"`
	Y                             int               `json:"y"`
	UUID                          string            `json:"uuid"`
	Slot                          int               `json:"slot"`
	Try                           int               `json:"try"`
	Type                          string            `json:"type"`
	ForcePass                     bool              `json:"force_pass"`
	Patrol                        *Patrol           `json:"patrol"`
	SubGroup                      *SubGroup         `json:"sub_group"` // [user_id]
	Ready                         bool              `json:"ready"`
	ResourceTask                  *ResourceSendTask `json:"resource_task"`
	SetBackToBaseWhenChangeSector bool              `json:"set_back_to_base_when_change_sector"`
	BackToBase                    bool              `json:"back_to_base"`
	Start                         int64             `json:"start"`
	ID                            int               `json:"id"`
	FixedSector                   bool              `json:"fixed_sector"`
	Global                        bool              `json:"global"`
	TimeOut                       int               `json:"time_out"`
	Stop                          bool              `json:"stop"`
	NotWarTypePath                string            `json:"not_war_type_path"`
	MissionTarget                 int               `json:"mission_target"`
	IgnoreBaseID                  int               `json:"ignore_base_id"`
	IgnoreMapID                   int               `json:"ignore_map_id"`
	robberyIgnore                 map[int]int64
	mx                            sync.RWMutex
}

func (m *Meta) GetRobberyIgnore(unitID int) bool {
	m.mx.Lock()
	defer m.mx.Unlock()

	for id, unixTime := range m.robberyIgnore {
		if time.Now().UnixNano()-unixTime > int64(time.Minute*3) {
			delete(m.robberyIgnore, id)
		}
	}

	_, ok := m.robberyIgnore[unitID]
	return ok
}

func (m *Meta) AddRobberyIgnore(unitID int) {
	m.mx.Lock()
	defer m.mx.Unlock()
	if m.robberyIgnore == nil {
		m.robberyIgnore = make(map[int]int64)
	}

	m.robberyIgnore[unitID] = time.Now().UnixNano()
}

func (m *Meta) GetJsonRobberyIgnore() []byte {
	m.mx.Lock()
	defer m.mx.Unlock()

	data, err := json.Marshal(m.robberyIgnore)
	if err != nil {
		fmt.Println("error get json robbery ignore")
	}

	return data
}

func (m *Meta) LoadRobberyIgnore(data []byte) {
	m.mx.Lock()
	defer m.mx.Unlock()

	ids := make(map[int]int64, 0)
	err := json.Unmarshal(data, &ids)
	if err != nil {
		fmt.Println("error load robbery ignore")
	}

	m.robberyIgnore = ids
}

type ResourceSendTask struct {
	GetResourceBaseID int  `json:"get_resource_base_id"`
	GetResource       bool `json:"get_resource"`
	ToBaseID          int  `json:"to_base_id"`
	ToObjectID        int  `json:"to_object_id"`
	ToMapID           int  `json:"to_map_id"`
}

func (m *Meta) GetCopy() *Meta {
	copyMeta := *m
	copyMeta.mx = sync.RWMutex{}
	return &copyMeta
}
