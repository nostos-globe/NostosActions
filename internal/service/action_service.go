package service

import (
	"main/internal/db"
	"main/internal/models"
	"time"
)

type ActionService struct {
    ActionRepo *db.ActionRepository
}

func (s *ActionService) LikeTrip(tripIDInt uint, tokenResponse uint) (models.Action, error) {
	action := models.Action{
		TargetID: tripIDInt,
		UserID: tokenResponse,
		ActionType: "like",
		TargetType: "trip",
		ActionDate: time.Now(),
	}
	result, err := s.ActionRepo.CreateAction(action)
	if err != nil {
		return models.Action{}, err	
	}
	return result, nil
}