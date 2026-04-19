package model

import "time"

type (
	// CreatorProiling only for query
	CreatorProfiling struct {
		CreatorID      int64
		Name           string
		Username       string
		Email          string
		HashedPassword string
		PhoneNumber    string
		Image_url      string
		Province       string
		Regency        string
		District       string
		SubDistrict    string
		Villages       string
		Address        string
		PostallCode    int64
		Status         string
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}

	CreatorData struct {
		ID          int64
		Creator     int64
		Username    string
		Email       int64
		PhoneNumber int64
		ImageUrl    int64
		Address     int64
		Version     int64
		Status      int64
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	Creator struct {
		ID        int64
		Name      string
		Username  string
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
		ID        int64
		ContactID int64
		CreatorId int64
		CreatedAt time.Time
		UpdatedAt time.Time
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
	Contact struct {
		ID          int64
		PhoneNumber string
		CreatedAt   time.Time
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
		UpdatedAt time.Time
	}

	ID struct {
		ID int64
	}

	// Data only for program
	RefreshTokenCreatorModel struct {
		ID           int64
		CreatorID    int64
		RefreshToken string
		CreatedAt    time.Time
		UpdatedAt    time.Time
		ExpiredAt    time.Time
	}

	CreatorBalance struct {
		ID               int64
		CreatorID        int64
		Balance          float64
		PendingBalance   float64
		WithdrawnBalance float64
		Currency         string
		Version          int64
		CreatedAt        time.Time
		UpdatedAt        time.Time
	}

	CreatorTransaction struct {
		ID               int64
		Creator_id       int64
		Balance_id       int64
		Transaction_type string
		Amount           float64
		Currency         string
		Reference_id     int64
		Description      string
		Status           string
		Version          int64
		Created_at       time.Time
		Updated_at       time.Time
	}

	CreatorDonation struct {
		ID                int64
		SenderCreatorID   int64
		ReceiverCreatorID int64
		TransactionID     int64
		Amount            float64
		Currency          string
		DonorName         string
		DonorEmail        string
		DonorMessage      string
		Message           string
		PaymentMethod     string
		ExternalReference int64
		Status            string
		Version           int64
		CreatedAt         time.Time
		UpdatedAt         time.Time
	}

	CreatorImage struct {
		ID        int64
		CreatorID int64
		ImageUrl  int64
		CreatedAt time.Time
	}

	Image struct {
		ID        int64
		Image     string
		CreatedAt time.Time
	}
)
