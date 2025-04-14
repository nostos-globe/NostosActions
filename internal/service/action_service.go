package service

import (
	"main/internal/db"
	"main/internal/models"
	"time"
)

type ActionService struct {
	ActionRepo *db.ActionRepository
}

func (s *ActionService) CreateAction(action models.Action) (models.Action, error) {

	result, err := s.ActionRepo.CreateAction(action)
	if err != nil {
		return models.Action{}, err
	}
	return result, nil
}

func (s *ActionService) LikeTrip(userID uint, targetID uint) (models.Like, error) {
	like := models.Like{
		SourceID:   userID,
		TargetID:   targetID,
		TargetType: "trip",
	}
	result, err := s.ActionRepo.LikeTrip(like)
	if err != nil {
		return models.Like{}, err
	}

	action := models.Action{
		TargetID:   targetID,
		UserID:     userID,
		ActionType: "like",
		TargetType: "trip",
		ActionDate: time.Now(),
	}

	_, err = s.ActionRepo.CreateAction(action)
	if err != nil {
		return models.Like{}, err
	}

	return result, nil
}

func (s *ActionService) UnlikeTrip(userID uint, targetID uint) (models.Like, error) {
	like := models.Like{
		SourceID:   userID,
		TargetID:   targetID,
		TargetType: "trip",
	}
	result, err := s.ActionRepo.UnlikeTrip(like)
	if err != nil {
		return models.Like{}, err
	}

	action := models.Action{
		TargetID:   targetID,
		UserID:     userID,
		ActionType: "unlike",
		TargetType: "trip",
		ActionDate: time.Now(),
	}

	_, err = s.ActionRepo.CreateAction(action)
	if err != nil {
		return models.Like{}, err
	}

	return result, nil
}

func (s *ActionService) FavMedia(userID uint, targetID uint) (models.Like, error) {
	fav := models.Like{
		SourceID:   userID,
		TargetID:   targetID,
		TargetType: "media",
	}
	result, err := s.ActionRepo.FavMedia(fav)
	if err != nil {
		return models.Like{}, err
	}

	action := models.Action{
		TargetID:   targetID,
		UserID:     userID,
		ActionType: "favourite",
		TargetType: "media",
		ActionDate: time.Now(),
	}
	_, err = s.ActionRepo.CreateAction(action)
	if err != nil {
		return models.Like{}, err
	}

	return result, nil
}

func (s *ActionService) UnFavMedia(userID uint, targetID uint) (models.Like, error) {
	fav := models.Like{
		SourceID:   userID,
		TargetID:   targetID,
		TargetType: "media",
	}
	result, err := s.ActionRepo.UnFavMedia(fav)
	if err != nil {
		return models.Like{}, err
	}

	action := models.Action{
		TargetID:   targetID,
		UserID:     userID,
		ActionType: "unfavourite",
		TargetType: "media",
		ActionDate: time.Now(),
	}
	_, err = s.ActionRepo.CreateAction(action)
	if err != nil {
		return models.Like{}, err
	}

	return result, nil
}

func (s *ActionService) GetTripLikes(tripID uint) ([]models.Like, error) {
	result, err := s.ActionRepo.GetTripLikes(tripID)
	if err != nil {
		return []models.Like{}, err
	}

	return result, nil
}
