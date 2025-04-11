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

func (repo *ActionRepository) LikeTrip(tripIDInt uint, tokenResponse uint) (models.Action, error) {
	
}