package domain

type Intolerance struct {
	ID   int
	Name string
}

type Diet struct {
	ID   int
	Name string
}

type Cuisine struct {
	ID   int
	Name string
}

type Cuisines []Cuisine
