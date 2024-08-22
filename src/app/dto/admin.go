package dto

type IAdmin interface {
	Validate() error
}

type RegisterSuperAdminDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"number"`
	Fullname string `json:"fullname" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterAdminDTO struct {
	RegisterSuperAdminDTO
	Role string `json:"role" validate:"required"`
}

type LoginAdminReqDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginAdminRestDTO struct {
	Token string `json:"token"`
}

func (r RegisterSuperAdminDTO) Validate() error {
	return validate.Struct(r)
}

func (r RegisterAdminDTO) Validate() error {
	return validate.Struct(r)
}

func (r LoginAdminReqDTO) Validate() error {
	return validate.Struct(r)
}
