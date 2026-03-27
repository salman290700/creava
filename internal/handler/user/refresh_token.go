package user

import (
	"gotweet/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RefreshToken(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.RefreshTokenRequest
	)

	if err := c.Copy().ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	user_id := c.GetInt64("userID") 
	token, refresh_token, setStatusCode, err := h.userService.RefreshToken(ctx, &req, user_id)
	if err != nil {
		c.JSON(setStatusCode, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(setStatusCode, dto.RefreshTokenResponse{
		Token:        token,
		RefreshToken: refresh_token,
	})
}
