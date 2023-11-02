package models

import (
	"time"
)

type User struct {
	Id           uint       `json:"id,omitempty"`
	Username     string     `json:"username"`
	PasswordHash string     `json:"-"`
	Role         string     `json:"role"`
	Created      *time.Time `json:"-" gorm:"autoCreateTime"`
	Updated      *time.Time `json:"-" gorm:"autoUpdateTime"`
}
