package db

import (
	"main/internal/models"

	"gorm.io/gorm"
)

type ActionRepository struct {
	DB *gorm.DB
}

func (repo *ActionRepository) CreateAction(action models.Action) (models.Action, error) {
	result := repo.DB.Table("activity.actions").Create(&action)
	return action, result.Error
}

func (repo *ActionRepository) GetActionsByUserID(userID uint) ([]models.Action, error) {
	var actions []models.Action
	result := repo.DB.Table("activity.actions").Where("user_id = ?", userID).Find(&actions)
	return actions, result.Error
}

func (repo *ActionRepository) LikeTrip(like models.Like) (models.Like, error) {
	result := repo.DB.Table("activity.likes").Create(&like)
	return like, result.Error
}

func (repo *ActionRepository) UnlikeTrip(like models.Like) (models.Like, error) {
	result := repo.DB.Table("activity.likes").Where("source_id =? AND target_id =? AND target_type =?", like.SourceID, like.TargetID, like.TargetType).Delete(&like)
	return like, result.Error
}

func (repo *ActionRepository) FindLike(like models.Like) (models.Like, error) {
	var foundLike models.Like
	result := repo.DB.Table("activity.likes").Where("source_id = ? AND target_id = ? AND target_type = ?", like.SourceID, like.TargetID, like.TargetType).First(&foundLike)
	return foundLike, result.Error
}
