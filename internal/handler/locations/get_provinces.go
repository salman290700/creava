package locations

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProvinces(c *gin.Context) {
	var (
		ctx = c.Request.Context()
	)

	res, err := h.locationService.GetProvincies(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, &res)
}
