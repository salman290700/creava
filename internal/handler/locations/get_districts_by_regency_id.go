package locations

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetDistrictsByRegencyId(c *gin.Context) {
	var (
		ctx = c.Request.Context()
	)
	regency_id := c.Query("regency_id")
	res, err := h.locationService.GetDistrictsByRegencyId(ctx, regency_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
