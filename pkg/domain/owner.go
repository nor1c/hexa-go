package domain

type Owner struct {
	ID       uint8  `json:"id"`
	Name     string `json:"name"`
	Age      uint8  `json:"age"`
	Address  string `json:"address"`
	IsActive int    `json:"is_active"`
}
