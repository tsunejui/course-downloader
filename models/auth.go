package models

type LoginRequest struct {
	Account  string `json:"account"`
	Confirm  bool   `json:"confirm"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
