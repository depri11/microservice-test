package models

type InputRegisterUser struct {
	Msisdn   int    `json:"msisdn"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseRegisterUser struct {
	ID       string `json:"id"`
	Msisdn   int    `json:"msisdn"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
