package entity

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Msisdn   int
	Name     string
	Username string
	Password string
}
