package models

import (
	"encoding/json"
)

type User struct {
	ID              uint   `json:"id" gorm:"column:id;type:autoIncrement;primaryKey"`
	Name            string `json:"name" gorm:"column:name"`
	Email           string `json:"email" gorm:"column:email"`
	EmailVerifiedAt string `json:"email_verified_at" gorm:"column:email_verified_at"`
	Password        string `json:"password" gorm:"column:password"`
	RememberToken   string `json:"remember_token" gorm:"column:remember_token"`
	CreatedAt       string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       string `json:"updated_at" gorm:"column:updated_at"`
	DeleteAt        string `json:"delete_at" gorm:"column:delete_at"`
}

// Required for Redis
func (u *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
