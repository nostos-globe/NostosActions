package controller

import (
	"main/internal/models"
	"main/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	ActionService  *service.ActionService
	AuthClient    *service.AuthClient
	ProfileClient *service.ProfileClient
}

func (c *AlbumController) CreateAlbum(ctx *gin.Context) {
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Visibility  string `json:"visibility"`
	}

	tokenCookie, err := ctx.Cookie("auth_token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token found"})
		return
	}

	tokenResponse, err := c.AuthClient.GetUserID(tokenCookie)
	if err != nil || tokenResponse == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "failed to find this user"})
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	albumMapper := &models.AlbumMapper{}
	album := albumMapper.ToAlbum(req, tokenResponse)

	createdAlbum, err := c.AlbumService.CreateAlbum(album)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create trip"})
		return
	}

	ctx.JSON(http.StatusCreated, createdAlbum)
}
