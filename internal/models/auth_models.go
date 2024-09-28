package models

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	Id        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type AuthMethodModel struct {
	Id                 uuid.UUID `json:"id"`
	UserId             uuid.UUID `json:"user_id"`
	AuthType           string    `json:"auth_type"`
	AuthProviderUserId uuid.UUID `json:"auth_provider_user_id"`
	Email              string    `json:"email"`
	Password           string    `json:"password"`
}
