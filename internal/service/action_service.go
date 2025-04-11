package service

import (
	"main/internal/db"
	"main/internal/models"
)

type ActionService struct {
    ActionRepo *db.ActionRepository
}