package dto

import "time"

type UserResp struct {
	Id         int       `json:"id,omitempty"`
	Username   string    `json:"username,omitempty"`
	Email      string    `json:"email,omitempty"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type CreateUserReq struct {
	Username        string `json:"username,omitempty" validate:"required"`
	Password        string `json:"password,omitempty" validate:"required"`
	ConfirmPassword string `json:"confirm_password,omitempty" validate:"required,eqfield=Password"`
	Email           string `json:"email,omitempty"`
}

type UpdateUserReq struct {
	Id              int    `json:"id,omitempty" validate:"required"`
	Username        string `json:"username,omitempty" validate:"required"`
	Password        string `json:"password,omitempty" validate:"required_with=ConfirmPassword"`
	ConfirmPassword string `json:"confirm_password,omitempty" validate:"required_with=Password,eqfield=Password"`
	Email           string `json:"email,omitempty"`
}

type UserLoginReq struct {
	Username string `json:"username,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

type UserLoginResp struct {
	AccessToken string `json:"accessToken"`
}
