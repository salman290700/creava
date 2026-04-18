package creatordonation

import (
	"gotweet/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCreatorDonation(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.CreatorDonationRequest
	)

	if err := c.Copy().ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	userID := c.GetInt64("userID")

	id_donation, setStatusCode, err := h.creatorDonationService.CreateCreatorDonation(ctx, &req, userID)
	if err != nil {
		c.JSON(setStatusCode, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(setStatusCode, &dto.CreatorDonationRespose{
		ID: id_donation,
	})
}
