package user

import (
	"encoding/json"
	"fmt"
	"gotweet/internal/dto"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) LoginWithGoogle(c *gin.Context) {
	// Variables
	var (
		ctx = c.Request.Context()
		req dto.LoginWithGoogleReq
	)
	// Bind Json
	if err := c.Copy().ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Validate JSON
	if err := h.validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Get data from Google API
	linkTokenProfile := fmt.Sprintf(`https://www.googleapis.com/oauth2/v3/userinfo?access_token=%s`, req.Token)
	req_google_api, err := http.NewRequest("POST", linkTokenProfile, nil)
	client := &http.Client{}
	resp, err := client.Do(req_google_api)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // Read the response body into a byte slice
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	var res_user_google dto.UserResGoogle
	// Bind JSON
	err = json.Unmarshal(body, &res_user_google)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := c.Copy().ShouldBindJSON(&res_user_google); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	// validate JSON
	if err := h.validate.Struct(res_user_google); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Service
	token, refresh_token, setStatusCode, err := h.userService.LoginWithGoogle(ctx, &res_user_google)
	// Response
	c.JSON(setStatusCode, dto.LoginResponse{
		Token:        token,
		RefreshToken: refresh_token,
	})
}
