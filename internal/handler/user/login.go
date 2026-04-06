package user

import (
	"fmt"
	"gotweet/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.LoginRequst
	)

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

	token, refresh_token, setStatusCode, err := h.userService.Login(ctx, &req)
	fmt.Println(err)
	if err != nil {
		c.JSON(setStatusCode, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(setStatusCode, &dto.LoginResponse{
		Token:        token,
		RefreshToken: refresh_token,
	})
}
