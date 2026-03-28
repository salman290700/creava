package post

import (
	"gotweet/internal/middleware"
	"gotweet/internal/service/post"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	api         *gin.Engine
	validate    *validator.Validate
	postService post.PostService
}

func NewPostHandler(
	api *gin.Engine,
	validate *validator.Validate,
	postService post.PostService,
) *Handler {
	return &Handler{
		api:         api,
		validate:    validate,
		postService: postService,
	}
}

func (h *Handler) RoutePostList(secret string) {
	postRouter := h.api.Group("/post")

	postRouter.Use(middleware.AuthMiddleware(secret))

}
