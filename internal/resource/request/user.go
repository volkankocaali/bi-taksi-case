package request

type UserLoginSchema struct {
	Username string `json:"username" form:"username" validate:"required,min=5,max=15"`
	Password string `json:"password" form:"password" validate:"required,min=6,max=20"`
}

type UserRegisterSchema struct {
	Username             string `json:"username" form:"username" validate:"required,min=5,max=15"`
	Password             string `json:"password" form:"password" validate:"required,min=6,max=20"`
	PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" validate:"required,min=6,max=20"`
}
