package npc_request

type DialogRequest struct {
	UUID         string              `json:"uuid"`
	Type         string              `json:"type"`
	Event        string              `json:"event"`
	Texts        []map[string]string `json:"texts"`
	From         int                 `json:"from"` // unit id
	To           int                 `json:"to"`   // unit id
	Target       int                 `json:"target"`
	FromUser     int                 `json:"from_user"`
	FromUserName string              `json:"from_user_name"`
	ToUser       int                 `json:"to_user"`
	ToUserName   string              `json:"to_user_name"`
	TargetUser   int                 `json:"target_user"`
	Success      int                 `json:"success"`
	Time         int64               `json:"time,omitempty"` // unix utc
	NeedAnswer   bool                `json:"need_answer"`
	NeedCredit   int                 `json:"need_credit"`
	Answers      []map[string]Answer `json:"answers"`
}

type Answer struct {
	Text string `json:"text"`
	ID   int    `json:"id"`
}
