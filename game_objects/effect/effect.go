package effect

type Effect struct {
	UUID        string `json:"u"` // equip_slot_number или еще чет
	Name        string `json:"-"`
	Parameter   string `json:"p"`
	Quantity    int    `json:"q"`
	Percentages bool   `json:"pr"`
	AddTime     int64  `json:"-"`
	TimeOut     int    `json:"-"`
	Subtract    bool   `json:"st"`
	// Все что ниже это для скилов
	WeaponType   string `json:"wt"`
	Fraction     string `json:"-"`
	StandardSize int    `json:"ss"`
	// ячейки эквипа, что бы удалять потом эффекты
	SlotType   int `json:"-"`
	SlotNumber int `json:"-"`
}

func (e *Effect) ToAccept(startValue float64, parameter string) float64 {
	if e.Parameter == parameter {

		diffValue := 0.0

		if e.Percentages {
			diffValue = startValue * (float64(e.Quantity) / 100)
		} else {
			diffValue = float64(e.Quantity)
		}

		if e.Subtract {
			startValue -= diffValue
		} else {
			startValue += diffValue
		}
	}

	return startValue
}
