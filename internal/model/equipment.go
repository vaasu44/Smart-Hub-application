package model

type Equipment struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Type     string `json:"type"`
	Status   string `json:"status"`
}
