package locations

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetSubdistrictsByDistrictId(c *gin.Context) {
	var (
		ctx = c.Request.Context()
	)
	district_id := c.Query("district_id")
	res, err := h.locationService.GetSubdistrictsByDistrictId(ctx, district_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
