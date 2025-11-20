package controller

import (
	"main/internal/models"
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

	likedTrip, err := c.ActionService.LikeTrip(tokenResponse, uint(tripIDInt))
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

	unlikedTrip, err := c.ActionService.UnlikeTrip(tokenResponse, uint(tripIDInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to unlike " + unlikedTrip.TargetType})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "trip unliked successfully"})
}

func (c *ActionController) FavMedia(ctx *gin.Context) {
	mediaID := ctx.Param("id")
	mediaIDInt, err := strconv.Atoi(mediaID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid media ID"})
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

	favMedia, err := c.ActionService.FavMedia(tokenResponse, uint(mediaIDInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fav media"})
		return
	}

	ctx.JSON(http.StatusCreated, favMedia)
}

func (c *ActionController) UnFavMedia(ctx *gin.Context) {
	mediaID := ctx.Param("id")
	mediaIDInt, err := strconv.Atoi(mediaID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid media ID"})
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

	unFavMedia, err := c.ActionService.UnFavMedia(tokenResponse, uint(mediaIDInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to unFav " + unFavMedia.TargetType})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "media unFav successfully"})
}

func (c *ActionController) GetTripLikes(ctx *gin.Context) {
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

	likes, err := c.ActionService.GetTripLikes(uint(tripIDInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get trip likes"})
		return
	}

	var profiles []interface{}
	for _, like := range likes {
		profile, err := c.ProfileClient.GetProfile(tokenCookie, like.SourceID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user"})
			return
		}
		profiles = append(profiles, profile)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"total_likes": len(likes),
		"profiles":    profiles,
	})
}

func (c *ActionController) GetMediaStatus(ctx *gin.Context) {
	mediaID := ctx.Param("id")
	mediaIDInt, err := strconv.Atoi(mediaID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid media ID"})
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

	isFavorite, err := c.ActionService.IsMediaFavorite(tokenResponse, uint(mediaIDInt))
	if err != nil {
		isFavorite = false
	}

	ctx.JSON(http.StatusOK, gin.H{"is_favourite": isFavorite})
}

func (c *ActionController) CreateAction(ctx *gin.Context) {
	var action models.Action
	if err := ctx.ShouldBindJSON(&action); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid action data"})
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

	action.UserID = tokenResponse
	createdAction, err := c.ActionService.CreateAction(action)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create action"})
		return
	}

	ctx.JSON(http.StatusCreated, createdAction)
}

func (c *ActionController) GetMyLikes(ctx *gin.Context) {
	tokenCookie, err := ctx.Cookie("auth_token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token found"})
		return
	}

	userID, err := c.AuthClient.GetUserID(tokenCookie)
	if err != nil || userID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "failed to find this user"})
		return
	}

	likes, err := c.ActionService.GetLikesByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get likes"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"likes": likes,
	})
}

func (c *ActionController) GetLikesByUserID(ctx *gin.Context) {
	tokenCookie, err := ctx.Cookie("auth_token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token found"})
		return
	}

	askerID, err := c.AuthClient.GetUserID(tokenCookie)
	if err != nil || askerID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "failed to find this user"})
		return
	}

	userIDStr := ctx.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	likes, err := c.ActionService.GetLikesByUserID(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get likes"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"likes": likes,
	})
}

func (c *ActionController) FollowUser(ctx *gin.Context) {
	var action models.Action
	if err := ctx.ShouldBindJSON(&action); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid action data"})
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

	action.UserID = tokenResponse
	createdAction, err := c.ActionService.FollowUser(tokenResponse, action.TargetID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create action"})
		return
	}

	ctx.JSON(http.StatusCreated, createdAction)

}

func (c *ActionController) UnFollowUser(ctx *gin.Context) {
	var action models.Action
	if err := ctx.ShouldBindJSON(&action); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid action data"})
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

	action.UserID = tokenResponse
	createdAction, err := c.ActionService.UnFollowUser(tokenResponse, action.TargetID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create action"})
		return
	}

	ctx.JSON(http.StatusCreated, createdAction)

}
