package locations

import (
	"context"
	"gotweet/internal/dto"
)

func (s *locationService) GetRegenciesbyProvinceId(ctx context.Context, province_id string) ([]dto.ResponseRegencies, error) {
	res, err := s.locationRepo.GetRegenciesByProvinciesId(ctx, province_id)
	if err != nil {
		return nil, err
	}
	var response []dto.ResponseRegencies

	for i := 0; i < len(res)-1; i++ {
		item := res[i]
		var data_response dto.ResponseRegencies
		data_response.Id = item.Id
		data_response.Code = item.Code
		data_response.Name = item.Name
		data_response.ProvinceCode = item.ProvincesCode
		response = append(response, data_response)
	}

	return response, nil
}
