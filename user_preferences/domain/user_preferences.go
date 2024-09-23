package domain

type UserCuisine struct {
	UserID    int
	CuisineID int
	Cuisine   string
}

type UserCuisines []UserCuisine

type UserDiet struct {
	UserID int
	DietID int
}

type UserDiets []UserDiet

type UserIntolerance struct {
	UserID        int
	IntoleranceID int
}

type UserIntolerances []UserIntolerance
