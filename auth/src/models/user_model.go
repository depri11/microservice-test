package models

import "github.com/google/uuid"

type InputRegisterUser struct {
	Msisdn   string `json:"msisdn"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseRegisterUser struct {
	ID       uuid.UUID `json:"id"`
	Msisdn   string    `json:"msisdn"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
}
