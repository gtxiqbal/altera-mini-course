package dto

type GrantType string

var (
	RefreshType GrantType = "refresh_token"
)

type LoginDtoReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshDtoReq struct {
	RefreshToken string    `json:"refresh_token" validate:"required"`
	GrantType    GrantType `json:"grant_type" validate:"required"`
}
