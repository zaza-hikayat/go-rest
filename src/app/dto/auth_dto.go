package dto

type IAuthDTO interface {
	Validate() error
}

type LoginReqDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5,max=20"`
}

type LoginResDTO struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type RegisterReqDTO struct {
	Name     string `json:"name" validate:"required,min=5,max=120"`
	Email    string `json:"email" validate:"required,min=5,max=20,email"`
	Password string `json:"password" validate:"required,min=5,max=20"`
	Address  string `json:"address"`
}

type ForgotPasswordReqDTO struct {
	Email string `json:"email" validate:"required,min=5,max=20,email"`
}

func (f ForgotPasswordReqDTO) Validate() error {
	return validate.Struct(f)
}

func (l *LoginReqDTO) Validate() error {
	return validate.Struct(l)
}

func (r RegisterReqDTO) Validate() error {
	return validate.Struct(r)
}
