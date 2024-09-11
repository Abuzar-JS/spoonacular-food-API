package models

type CreateUserRequest struct {
	Name     string `json:"name"`
	Cuisine  string `json:"cuisine"`
	Password string `json:"password"`
}
