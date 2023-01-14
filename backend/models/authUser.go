package models

import (
	"context"
)

// register user
type RegisterUser struct {
	ID           int64  `json:"id"`
	NamaDepan    string `json:"nama_depan" binding:"required"`
	NamaBelakang string `json:"nama_belakang" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Phone        string `json:"phone" binding:"required"`
	Password     string `json:"password" binding:"required,min=8"`
	Role         string `json:"role"`
}

// login user
type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthRepository interface {
	Login(ctx context.Context, login *LoginUser) (User, error)
	Register(ctx context.Context, register *RegisterUser) (RegisterUser, error)
}
