package domain

import "time"

type PetKind struct {
	ID          uint8  `json:"id"`
	Name        string `json:"name"`
	Leg         uint8  `json:"leg"`
	CanSwim     bool   `json:"can_swim"`
	CanFly      bool   `json:"can_fly"`
	CanRun      bool   `json:"can_run"`
	IsDangerous bool   `json:"is_dangerous"`
	IsVenomous  bool   `json:"is_venomous"`
}

type Pet struct {
	ID           uint8     `json:"id"`
	KindID       uint8     `json:"kind_id"`
	OwnerID      uint8     `json:"owner_id"`
	Name         string    `json:"name"`
	Age          uint8     `json:"age"`
	AdoptionDate time.Time `json:"adoption_date"`
}
