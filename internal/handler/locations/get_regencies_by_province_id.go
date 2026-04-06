package locations

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type province_code struct {
	Province_code string `uri:"province_code" binding:"required"`
}

func (h *Handler) GetRegenciesByProvinceId(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		// req province_code
	)
	// province_id := c.("province_code")

	// if err := c.ShouldBindUri(&req); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"messages": err.Error(),
	// 	})
	// 	return
	// }
	// fmt.Println(req.Province_code)
	province_id := c.Query("province_code")
	fmt.Println(province_id)
	res, err := h.locationService.GetRegenciesbyProvinceId(ctx, province_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
