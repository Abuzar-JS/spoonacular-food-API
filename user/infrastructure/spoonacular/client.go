package spoonacular

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Abuzar-JS/go-spoonacular-api/user/domain"
)

type SpoonacularClient struct {
}

type Recipe struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}

func (u Recipe) ToDomain() *domain.Recipe {
	return &domain.Recipe{
		ID:    u.ID,
		Title: u.Title,
		Image: u.Image,
	}
}

type RecipeRows []Recipe

func (r RecipeRows) ToDomain() domain.Recipes {
	recipes := make(domain.Recipes, len(r))

	for i, allRecipes := range r {
		recipes[i] = *allRecipes.ToDomain()
	}
	return recipes
}

func (s *SpoonacularClient) GetSpoonacularRecipe(ctx context.Context, cuisine string) (domain.Recipes, error) {
	apiKey := os.Getenv("SPOONACULAR_API_KEY")
	apiUrl := "https://api.spoonacular.com/recipes/complexSearch"
	url := fmt.Sprintf("%s?cuisine=%s&apiKey=%s", apiUrl, cuisine, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response struct {
		Recipes []Recipe `json:"results"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error: ", err)
	}

	return RecipeRows(response.Recipes).ToDomain(), nil

}
