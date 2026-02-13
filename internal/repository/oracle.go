package repository

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/godror/godror"
	"smart-hub/internal/model"
)

type Repo struct {
	DB *sql.DB
}

func NewRepo() (*Repo, error) {
	dsn := "smarthub/smarthub@localhost:1521/XEPDB1"
	db, err := sql.Open("godror", dsn)
	if err != nil {
		return nil, err
	}
	return &Repo{DB: db}, nil
}

func (r *Repo) AddEquipment(e model.Equipment) error {
	_, err := r.DB.ExecContext(context.Background(),
		"INSERT INTO equipment(name,location,type,status) VALUES(:1,:2,:3,:4)",
		e.Name, e.Location, e.Type, "DISCONNECTED")
	return err
}

func (r *Repo) LogActivity(id int, action string) error {
	_, err := r.DB.Exec(
		"INSERT INTO activity(equipment_id,action,activity_time) VALUES(:1,:2,:3)",
		id, action, time.Now())
	return err
}

func (r *Repo) GetActivity() ([]model.Activity, error) {
	rows, err := r.DB.Query(
		"SELECT id,equipment_id,action,activity_time FROM activity")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var acts []model.Activity
	for rows.Next() {
		var a model.Activity
		rows.Scan(&a.ID, &a.EquipmentID, &a.Action, &a.Time)
		acts = append(acts, a)
	}
	return acts, nil
}

