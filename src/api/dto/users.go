package dto

type RegisterUserByUsername struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email"`
	Password  string `json:"password" binding:"password"`
}
type RegisterLoginByPhone struct {
	Phone string `json:"phone" binding:"required,phone,min=11,max=11"`
	Otp   string `json:"otp" binding:"required"`
}
type LoginByUsername struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=3"`
}
