package service

import (
	"main/internal/db"
	"main/internal/events"
	"main/internal/models"
	"time"
)

type ActionService struct {
	ActionRepo *db.ActionRepository
	Events     *events.Publisher
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

	if s.Events != nil {
		event := events.ContentLikedEvent{
			SourceID:   userID,
			TargetID:   targetID,
			TargetType: "trip",
			CreatedAt:  time.Now(),
		}
		_ = s.Events.Publish("content.liked", event)
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

	if s.Events != nil {
		event := events.ContentLikedEvent{
			SourceID:   userID,
			TargetID:   targetID,
			TargetType: "media",
			CreatedAt:  time.Now(),
		}
		_ = s.Events.Publish("content.favourited", event)
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

func (s *ActionService) IsMediaFavorite(userID uint, targetID uint) (bool, error) {
	fav := models.Like{
		SourceID:   userID,
		TargetID:   targetID,
		TargetType: "media",
	}

	result, err := s.ActionRepo.IsMediaFavorite(fav)
	if err != nil {
		return false, err
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

func (s *ActionService) FollowUser(userID uint, targetID uint) (models.Action, error) {
	action := models.Action{
		TargetID:   targetID,
		UserID:     userID,
		ActionType: "follow",
		TargetType: "user",
		ActionDate: time.Now(),
	}

	result, err := s.CreateAction(action)
	if err != nil {
		return models.Action{}, err
	}

	if s.Events != nil {
		event := events.UserFollowedEvent{
			FollowerID: userID,
			FollowedID: targetID,
			CreatedAt:  time.Now(),
		}
		_ = s.Events.Publish("user.followed", event)
	}

	return result, nil
}

func (s *ActionService) UnFollowUser(userID uint, targetID uint) (models.Action, error) {
	action := models.Action{
		TargetID:   targetID,
		UserID:     userID,
		ActionType: "unfollow",
		TargetType: "user",
		ActionDate: time.Now(),
	}

	result, err := s.CreateAction(action)
	if err != nil {
		return models.Action{}, err
	}

	return result, nil
}

func (s *ActionService) GetLikesByUserID(userID uint) ([]models.Like, error) {
	likes, err := s.ActionRepo.GetLikesByUserID(userID)
	if err != nil {
		return []models.Like{}, err
	}
	return likes, nil
}
