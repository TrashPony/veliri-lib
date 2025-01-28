package generate_map_config

type GenerateMapConfig struct {
	ResourceCount      int
	ResourceQuantityK  float64
	StartReservoirs    []int
	DefaultResourceIDs []int
	StartGeysers       int
	MapGeysers         int
	XSize              int
	YSize              int
	ElevatorSize       int
}

func DefaultConfig() GenerateMapConfig {
	return GenerateMapConfig{
		ResourceCount:      2,
		StartReservoirs:    []int{8, 10, 7},
		DefaultResourceIDs: []int{9, 11, 13, 15, 17, 19},
		StartGeysers:       1,
		MapGeysers:         3,
		XSize:              4096,
		YSize:              4096,
		ResourceQuantityK:  7,
		ElevatorSize:       6,
	}
}

func ConfigByLvl(lvl int) GenerateMapConfig {
	if lvl == 0 {
		return GenerateMapConfig{
			ResourceCount:      1,
			StartReservoirs:    []int{8, 10, 7},
			DefaultResourceIDs: []int{9, 11, 13, 15, 17, 19},
			StartGeysers:       1,
			MapGeysers:         1,
			XSize:              3072,
			YSize:              3072,
			ResourceQuantityK:  5,
			ElevatorSize:       3,
		}
	}

	if lvl == 1 {
		return DefaultConfig()
	}

	if lvl == 2 {
		return GenerateMapConfig{
			ResourceCount:      3,
			StartReservoirs:    []int{8, 10, 7},
			DefaultResourceIDs: []int{9, 11, 13, 15, 17, 19},
			StartGeysers:       2,
			MapGeysers:         6,
			XSize:              6144,
			YSize:              6144,
			ResourceQuantityK:  10,
			ElevatorSize:       9,
		}
	}

	return DefaultConfig()
}
