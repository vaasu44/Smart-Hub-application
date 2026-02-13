package model

import "time"

type Activity struct {
	ID          int       `json:"id"`
	EquipmentID int       `json:"equipment_id"`
	Action      string    `json:"action"`
	Time        time.Time `json:"time"`
}

