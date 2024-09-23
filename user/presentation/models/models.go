package models

type CreateUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
