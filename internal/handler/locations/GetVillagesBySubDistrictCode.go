package locations

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetVillagesBySubDistrictCode(c *gin.Context) {
	var (
		ctx = c.Request.Context()
	)
	sub_district_code := c.Query("sub_district_code")
	res, err := h.locationService.GetVillagesBySubDistrictCode(ctx, sub_district_code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
