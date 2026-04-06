package locations

import (
	"gotweet/internal/service/locations"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	api             *gin.Engine
	validate        *validator.Validate
	locationService locations.LocationsService
}

func NewLocationHandler(
	api *gin.Engine,
	validate *validator.Validate,
	locationService locations.LocationsService,
) *Handler {
	return &Handler{
		api:             api,
		validate:        validate,
		locationService: locationService,
	}
}

func (h *Handler) RoutePostList(secretKey string) {
	locationsRouter := h.api.Group("/locations")
	locationsRouter.GET("/get-provinces", h.GetProvinces)
	locationsRouter.GET("/get-regencies-by-province-code", h.GetRegenciesByProvinceId)
	locationsRouter.GET("/get-district-by-regency-id/", h.GetDistrictsByRegencyId)
	locationsRouter.GET("/get-sub-district-by-district-id", h.GetSubdistrictsByDistrictId)

}
