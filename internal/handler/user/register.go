package user

import (
	"fmt"
	"gotweet/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.RegisterRequest
	)
	fmt.Println(&req)

	if err := c.Copy().ShouldBindJSON(&req); err != nil {
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
	userID, setStatusCode, err := h.userService.Register(ctx, &req)
	if err != nil {
		c.JSON(setStatusCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(setStatusCode, dto.RegisterResponse{
		ID: userID,
	})
}
