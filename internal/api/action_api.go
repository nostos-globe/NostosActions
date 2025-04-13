package controller

import (
	"main/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActionController struct {
	ActionService *service.ActionService
	AuthClient    *service.AuthClient
	ProfileClient *service.ProfileClient
}

func (c *ActionController) LikeTrip(ctx *gin.Context) {
	tripID := ctx.Param("id")
	tripIDInt, err := strconv.Atoi(tripID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid trip ID"})
		return
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

	likedTrip, err := c.ActionService.LikeTrip(uint(tripIDInt), tokenResponse)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to like trip"})
		return
	}

	ctx.JSON(http.StatusCreated, likedTrip)
}

func (c *ActionController) UnlikeTrip(ctx *gin.Context) {
	tripID := ctx.Param("id")
	tripIDInt, err := strconv.Atoi(tripID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid trip ID"})
		return
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

	unlikedTrip, err := c.ActionService.UnlikeTrip(uint(tripIDInt), tokenResponse)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to unlike trip"})
		return
	}

	ctx.JSON(http.StatusOK, unlikedTrip)
}
