package dto

type Provincies_res struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type ResponseRegencies struct {
	Id           int64  `json:"id"`
	Code         string `json:"code"`
	ProvinceCode string `json:"province_code"`
	Name         string `json:"name"`
}

type ResponseDistricts struct {
	ID          int64
	Code        string `json:"code"`
	Name        string `json:"name"`
	RegencyCode string `json:"regency_code"`
}

type ResponseSubdistricts struct {
	ID           int64
	Code         string `json:"code"`
	Name         string `json:"name"`
	DistrictCode string `json:"district_code"`
}

type ResponseVillages struct {
	ID              int64
	Code            string `json:"code"`
	Name            string `json:"name"`
	SubDistrictCode string `json:"sub_district_code"`
}
