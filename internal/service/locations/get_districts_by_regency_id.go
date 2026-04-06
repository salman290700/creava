package locations

import (
	"context"
	"gotweet/internal/dto"
)

func (s *locationService) GetDistrictsByRegencyId(ctx context.Context, regency_id string) ([]dto.ResponseDistricts, error) {
	res, err := s.locationRepo.GetDistrictsByRegencyId(ctx, regency_id)
	if err != nil {
		return nil, err
	}
	var response []dto.ResponseDistricts

	for i := 0; i < len(res)-1; i++ {
		item := res[i]
		var data dto.ResponseDistricts
		data.ID = item.Id
		data.Code = item.Code
		data.RegencyCode = item.RegencyCode
		data.Name = item.Name

		response = append(response, data)
	}
	return response, nil
}
