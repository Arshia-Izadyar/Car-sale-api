package dto

type OtpDTO struct {
	Value string
	Used  bool
}

type GetOtpRequest struct {
	MobileNumber string `json:"mobile_number" binding:"required,phone,min=11,max=11"`
}
