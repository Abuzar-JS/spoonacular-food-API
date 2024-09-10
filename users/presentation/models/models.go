package models

type CreateUserRequest struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Password string `json:"password"`
}
