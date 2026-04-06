package locations

import (
	"context"
	"gotweet/internal/dto"
	"gotweet/internal/repository/locations"

	"github.com/gin-gonic/gin"
)

type LocationsService interface {
	GetProvincies(ctx context.Context) ([]dto.Provincies_res, error)
	GetRegenciesbyProvinceId(ctx context.Context, province_id string) ([]dto.ResponseRegencies, error)
	GetDistrictsByRegencyId(ctx context.Context, regency_id string) ([]dto.ResponseDistricts, error)
	GetSubdistrictsByDistrictId(ctx context.Context, district_id string) ([]dto.ResponseSubdistricts, error)
	GetVillagesBySubDistrictCode(ctx context.Context, sub_district_id string) ([]dto.ResponseVillages, error)
}
type locationService struct {
	api          *gin.Engine
	locationRepo locations.LocationsRepository
}

func NewLocationService(api *gin.Engine, locationRepo locations.LocationsRepository) LocationsService {
	return &locationService{
		api:          api,
		locationRepo: locationRepo,
	}
}
