package domain

type User struct {
	ID       int
	Name     string
	Password string
}

type UserResponse struct {
	ID   int
	Name string
}

type Recipe struct {
	ID    int
	Title string
	Image string
}

type Recipes []Recipe
