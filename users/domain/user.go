package domain

type User struct {
	ID       int
	Name     string
	Cuisine  string
	Password string 
}

type UserResponse struct{
	ID int
	Name string
	Cuisine string
}
