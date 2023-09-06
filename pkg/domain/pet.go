package domain

type PetKind struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Leg          uint8  `json:"leg"`
	Can_Swim     bool   `json:"can_swim"`
	Can_Fly      bool   `json:"can_fly"`
	Can_Run      bool   `json:"can_run"`
	Is_Dangerous bool   `json:"is_dangerous"`
	Is_Venomous  bool   `json:"is_venomous"`
}

type Pet struct {
	ID            int    `json:"id"`
	Kind_ID       uint8  `json:"kind_id"`
	Owner_ID      uint8  `json:"owner_id"`
	Name          string `json:"name"`
	Age           uint8  `json:"age"`
	Adoption_Date string `json:"adoption_date"`
}
