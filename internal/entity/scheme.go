package entity

type Scheme struct {
	Id     string `json:"id"`
	ScName string `json:"sc_name"`
	ScIdx  int    `json:"sc_idx"`
	Lichka string `json:"lichka"`
	LichkaId int `json:"lichka_id"`
	Link   string `json:"link"`
	ChatCheckLink   string `json:"chat_check_link"`
	ChatCheckId int `json:"chat_check_id"`
}
