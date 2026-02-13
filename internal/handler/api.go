package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"smart-hub/internal/model"
	"smart-hub/internal/service"
)

type API struct {
	svc *service.HubService
}

func NewAPI(s *service.HubService) *API {
	return &API{svc: s}
}

func (a *API) AddEquipment(w http.ResponseWriter, r *http.Request) {
	var e model.Equipment
	json.NewDecoder(r.Body).Decode(&e)
	a.svc.AddEquipment(e)
	w.WriteHeader(http.StatusCreated)
}

func (a *API) Control(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	action := r.URL.Query().Get("action")
	a.svc.Control(id, action)
	w.WriteHeader(http.StatusOK)
}

func (a *API) Activity(w http.ResponseWriter, r *http.Request) {
	data, _ := a.svc.Activity()
	json.NewEncoder(w).Encode(data)
}

