package locations

import (
	"context"
	"gotweet/internal/dto"
)

func (s *locationService) GetProvincies(ctx context.Context) ([]dto.Provincies_res, error) {
	res, err := s.locationRepo.GetProvinces(ctx)
	if err != nil {
		return nil, err
	}
	var response []dto.Provincies_res
	for i := 0; i < len(res)-1; i++ {
		item := res[i]
		var data dto.Provincies_res
		data.Code = item.Code
		data.ID = item.ID
		data.Name = item.Name
		response = append(response, data)
	}
	return response, nil
}
