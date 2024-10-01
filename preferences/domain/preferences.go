package domain

type UserPreferences struct {
	UserID           int
	UserCuisines     UserCuisines
	UserDiets        UserDiets
	UserIntolerances UserIntolerances
}

type Intolerance struct {
	ID   int
	Name string
}

type Intolerances []Intolerance

type Diet struct {
	ID   int
	Name string
}

type Diets []Diet

type Cuisine struct {
	ID   int
	Name string
}

type Cuisines []Cuisine

type UserCuisine struct {
	UserID    int
	CuisineID int
	Cuisine   string
}

type UserCuisines []UserCuisine

type UserDiet struct {
	UserID int
	DietID int
	Diet   string
}

type UserDiets []UserDiet

type UserIntolerance struct {
	UserID        int
	IntoleranceID int
	Intolerance   string
}

type UserIntolerances []UserIntolerance
