package build_option

type BuildOption struct {
	X        int `json:"x"`
	Y        int `json:"y"`
	Radius   int `json:"radius"`
	ObjectID int `json:"object_id"`
}

type BaseConfig struct {
	Template     []*BuildOption `json:"template"`
	Radius       int            `json:"radius"`
	X            int            `json:"x"`
	Y            int            `json:"y"`
	Rotate       int            `json:"rotate"`
	BuildTimeOut int64          `json:"build_time_out"`
}
