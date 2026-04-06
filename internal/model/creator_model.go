package model

import "time"

type (
	Creator struct {
		ID        int64
		Name      string
		Password  string
		Version   int64
		CreatedAt time.Time
	}

	CreatorAddress struct {
		ID              int64
		CreatorId       int64
		ProvinceCode    string
		RegencyCode     string
		DistrictCode    string
		SubDistrictCode string
		VillageCode     string
	}

	CreatorContact struct {
		ID          int64
		PhoneNumber string
		CreatorId   int64
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	CreatorEmail struct {
		ID        int64
		Email     int64
		CreatorId int64
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	CreatorStatus struct {
		ID        int64
		CreatorID int64
		Status    int64
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	CreatorData struct {
		ID          int64
		Name        string
		Email       string
		Password    string
		PhoneNumber CreatorContact
		Address     CreatorAddress
		Version     int64
		Status      int64
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	Email struct {
		ID        int64
		Email     string
		Verified  bool
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	Status struct {
		ID        int64
		Status    string
		CreatedAt time.Time
	}
)
