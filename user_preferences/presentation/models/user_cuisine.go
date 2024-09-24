package models

type CreateUserCuisineRequest struct {
	UserID    int `json:"user_id"`
	CuisineID int `json:"cuisine_id"`
}
