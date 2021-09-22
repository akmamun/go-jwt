package models

import "time"

type User struct {
	Email         *string   `json:"email" validate:"email,required"`
	Password      *string   `json:"password" validate:"required,min=6"`
	Token         *string   `json:"token"`
	Refresh_token *string   `json:"refresh_token"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}
