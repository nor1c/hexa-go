package domain

type Owner struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Age       uint8  `json:"age"`
	Address   string `json:"address"`
	Is_Active uint8  `json:"is_active"`
}
