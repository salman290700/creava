package locations

import (
	"context"
	"gotweet/internal/dto"
)

func (s *locationService) GetVillagesBySubDistrictCode(ctx context.Context, sub_district_id string) ([]dto.ResponseVillages, error) {
	var (
		data     dto.ResponseVillages
		response []dto.ResponseVillages
	)
	rows, err := s.locationRepo.GetVillagesBySubdistrictId(ctx, sub_district_id)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(rows); i++ {
		item := rows[i]
		data.ID = item.Id
		data.Code = item.Code
		data.Name = item.Name
		data.SubDistrictCode = item.SubDistrictCode
		response = append(response, data)
	}
	return response, nil
}
