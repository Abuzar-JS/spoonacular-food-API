package models

type CreateUserRequest struct {
	Name     string `json:"name"`
	Cuisine  string `json:"cuisine"`
	Password string `json:"password"`
}

type GetUserRequest struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Cuisine string `json:"cuisine"`
}
