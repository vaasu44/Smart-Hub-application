package service

import (
	"smart-hub/internal/model"
	"smart-hub/internal/repository"
)

type HubService struct {
	repo *repository.Repo
}

func NewHubService(r *repository.Repo) *HubService {
	return &HubService{repo: r}
}

func (h *HubService) AddEquipment(e model.Equipment) error {
	return h.repo.AddEquipment(e)
}

func (h *HubService) Control(id int, action string) error {
	return h.repo.LogActivity(id, action)
}

func (h *HubService) Activity() ([]model.Activity, error) {
	return h.repo.GetActivity()
}

