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

type InputLoginUser struct {
	Msisdn   string `json:"msisdn"`
	Password string `json:"password"`
}

type ResponseLoginUser struct {
	Token string `json:"token"`
}
