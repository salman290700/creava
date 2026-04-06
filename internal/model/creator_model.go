package model

import "time"

type (
	// CreatorProiling only for query
	CreatorProfiling struct {
		CreatorID      int64
		Name           string
		Email          string
		HashedPassword string
		PhoneNumber    string
		Province       string
		Regency        string
		District       string
		SubDistrict    string
		Villages       string
		Address        string
		PostallCode    int64
	}

	CreatorData struct {
		ID          int64
		Creator     int64
		Email       int64
		PhoneNumber int64
		Address     int64
		Version     int64
		Status      int64
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	Creator struct {
		ID        int64
		Name      string
		Password  string
		Version   int64
		CreatedAt time.Time
	}

	CreatorAddress struct {
		ID        int64
		CreatorId int64
		Address   Address
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
	}

	CreatorStatus struct {
		ID        int64
		CreatorID int64
		Status    int64
		CreatedAt time.Time
	}

	// Addres in Indonesia
	Address struct {
		ID              int64
		ProvinceCode    string
		RegencyCode     string
		DistrictCode    string
		SubDistrictCode string
		VillageCode     string
		Address         string
		PostalCode      int64
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

	ID struct {
		ID int64
	}
)
