package entity

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Msisdn   string
	Name     string
	Username string
	Password string
}
