package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Msisdn   uint64    `json:"msisdn"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}
