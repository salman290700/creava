package locations

import (
	"context"
	"gotweet/internal/dto"
)

func (s *locationService) GetSubdistrictsByDistrictId(ctx context.Context, district_id string) ([]dto.ResponseSubdistricts, error) {
	res, err := s.locationRepo.GetSubDistrictsByDistrictId(ctx, district_id)
	if err != nil {
		return nil, err
	}
	var response []dto.ResponseSubdistricts
	for i := 0; i < len(res)-1; i++ {
		item := res[i]

		var data dto.ResponseSubdistricts
		data.ID = item.Id
		data.Name = item.Name
		data.Code = item.Code
		data.DistrictCode = item.DistrictCode
		response = append(response, data)
	}
	return response, nil
}
