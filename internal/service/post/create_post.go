package post

import (
	"context"
	"gotweet/internal/dto"
	"gotweet/internal/model"
	"net/http"
	"time"
)

func (s *postService) CreatePost(ctx context.Context, dto *dto.CreatePostRequest, userID int64) (int64, int, error) {
	now := time.Now()
	data := &model.PostModel{
		UserID:    userID,
		Title:     dto.Title,
		Content:   dto.Content,
		CreatedAt: now,
		UpdatedAt: now,
	}
	id, err := s.postRepo.CreatePost(ctx, data)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	return id, http.StatusAccepted, nil
}
