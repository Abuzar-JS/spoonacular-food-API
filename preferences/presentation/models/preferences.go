package models

type CreateCuisineRequest struct {
	Name string `json:"name"`
}

type CreateIntoleranceRequest struct {
	Name string `json:"name"`
}

type CreateDietRequest struct {
	Name string `json:"name"`
}

type AddUserPreferencesRequest struct {
	UserID        int   `json:"user_ID"`
	CuisineID     []int `json:"cuisines"`
	IntoleranceID []int `json:"intolerances"`
	DietID        []int `json:"diets"`
}
