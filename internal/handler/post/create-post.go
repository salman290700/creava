package post

import (
	"gotweet/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePosthandler(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.CreatePostRequest
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

	userID := c.GetInt64("userID")
	id_post, setStatusCode, err := h.postService.CreatePost(ctx, &req, userID)
	if err != nil {
		c.JSON(setStatusCode, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(setStatusCode, &dto.CreatePostResponse{
		ID: id_post,
	})
}
