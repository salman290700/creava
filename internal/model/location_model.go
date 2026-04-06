package model

type (
	Provincies struct {
		ID   int64
		Code string
		Name string
	}

	Regencies struct {
		Id            int64
		Code          string
		ProvincesCode string
		Name          string
	}

	Districts struct {
		Id          int64
		Code        string
		RegencyCode string
		Name        string
	}

	SubDistricts struct {
		Id           int64
		Code         string
		DistrictCode string
		Name         string
	}

	Villages struct {
		Id              int64
		Code            string
		SubDistrictCode string
		Name            string
	}
)
