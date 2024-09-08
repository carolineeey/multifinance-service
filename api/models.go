package api

type ErrorResponse struct {
	Error   error  `json:"error"`
	Message string `json:"message"`
}

// SetUserRequest represents the payload to add or edit a customer.
type SetUserRequest struct {
	Nik       int    `json:"nik"`
	FullName  string `json:"full_name"`
	LegalName string `json:"legal_name"`
	BirtPlace string `json:"birth_place"`
	BirthDate int64  `json:"birth_date"`
	Salary    int    `json:"salary"`
}

// CustomerInfo represents the detail of a customer.
type CustomerInfo struct {
	Nik        int    `json:"nik"`
	FullName   string `json:"full_name"`
	LegalName  string `json:"legal_name"`
	BirtPlace  string `json:"birth_place"`
	BirthDate  int64  `json:"birth_date"`
	Salary     int    `json:"salary"`
	KtpPhotoId string `json:"ktp_photo_id"`
	SelfieId   string `json:"selfie_id"`
}
