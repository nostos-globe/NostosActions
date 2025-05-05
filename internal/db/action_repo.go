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

func (repo *ActionRepository) FavMedia(fav models.Like) (models.Like, error) {
	result := repo.DB.Table("activity.likes").Create(&fav)
	return fav, result.Error
}

func (repo *ActionRepository) UnFavMedia(fav models.Like) (models.Like, error) {
	result := repo.DB.Table("activity.likes").Where("source_id =? AND target_id =? AND target_type =?", fav.SourceID, fav.TargetID, fav.TargetType).Delete(&fav)
	return fav, result.Error
}

func (repo *ActionRepository) FindLike(like models.Like) (models.Like, error) {
	var foundLike models.Like
	result := repo.DB.Table("activity.likes").Where("source_id = ? AND target_id = ? AND target_type = ?", like.SourceID, like.TargetID, like.TargetType).First(&foundLike)
	return foundLike, result.Error
}

func (repo *ActionRepository) GetTripLikes(tripID uint) ([]models.Like, error) {
	var likes []models.Like
	result := repo.DB.Table("activity.likes").Where("target_id =? AND target_type =?", tripID, "trip").Find(&likes)
	return likes, result.Error
}

func (repo *ActionRepository) GetMyTripLikes(userID uint) ([]models.Like, error) {
	var likes []models.Like
	result := repo.DB.Table("activity.likes").Where("source_id = ? AND target_type =?", userID, "trip").Find(&likes)
	return likes, result.Error
}

func (repo *ActionRepository) IsMediaFavorite(fav models.Like) (bool, error) {
	var foundLike models.Like
	result := repo.DB.Table("activity.likes").Where("source_id =? AND target_id =? AND target_type =?", fav.SourceID, fav.TargetID, fav.TargetType).First(&foundLike)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
