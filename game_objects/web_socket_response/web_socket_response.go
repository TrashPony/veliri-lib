package web_socket_response

import "sync"

type Response struct {
	Responses         []*Response `json:"-"`
	Event             string      `json:"event,omitempty"`
	Service           string      `json:"service,omitempty"`
	Error             string      `json:"error,omitempty"`
	Data              interface{} `json:"data,omitempty"`
	MapID             int         `json:"-"`
	UserID            int         `json:"-"`
	PlayerID          int         `json:"-"`
	OnlyData          bool        `json:"-"`
	BinaryMsg         []byte      `json:"-"`
	X                 int         `json:"-"`
	Y                 int         `json:"-"`
	ToX               int         `json:"-"`
	ToY               int         `json:"-"`
	CheckTo           bool        `json:"-"`
	Radius            int         `json:"-"`
	ID                int         `json:"-"`
	IgnoreWatchJammer bool        `json:"-"`
	All               bool        `json:"-"`
	CheckView         bool        `json:"-"`
	RadarIsView       bool        `json:"-"`
	NodeName          string      `json:"-"`
	ExcludePlayerID   int         `json:"-"`
}

type ResponsesStore struct {
	ResponsesGroup []*ResponsesGroup
	mx             sync.RWMutex
}

type ResponsesGroup struct {
	Key        string             `json:"key"`
	Responses  *GameLoopResponses `json:"Responses"`
	Type       string             `json:"type"` // move - проверяем функцией "moveResponsesToUser", view - "ResponseToBin"
	Attributes map[string]string  `json:"attributes"`
}

type GameLoopResponses struct {
	Responses []Response
	mx        sync.RWMutex
}

func (m *GameLoopResponses) AddResponse(msg Response) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.Responses = append(m.Responses, msg)
}

func (ms *ResponsesStore) AddResponse(typeMsg, typeCheck string, msg Response, attributes map[string]string) {
	ms.mx.Lock()
	defer ms.mx.Unlock()

	ms.getResponseGroup(typeMsg, typeCheck, attributes).Responses.AddResponse(msg)
}

func (ms *ResponsesStore) Clear() {
	//ms.ResponsesGroup = nil
	for _, m := range ms.ResponsesGroup {
		m.Responses.Responses = m.Responses.Responses[:0]
	}
}

func (ms *ResponsesStore) getResponseGroup(typeMsg, typeCheck string, attributes map[string]string) *ResponsesGroup {
	for _, m := range ms.ResponsesGroup {
		if m.Key == typeMsg {
			return m
		}
	}

	newGroup := &ResponsesGroup{
		Key:        typeMsg,
		Responses:  &GameLoopResponses{Responses: make([]Response, 0)},
		Type:       typeCheck,
		Attributes: attributes,
	}

	ms.ResponsesGroup = append(ms.ResponsesGroup, newGroup)
	return newGroup
}

func (ms *ResponsesStore) GetResponseGroups() []*ResponsesGroup {
	return ms.ResponsesGroup
}
