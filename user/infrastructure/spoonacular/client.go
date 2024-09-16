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

type RecipeRow struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func (u RecipeRow) ToDomain() *domain.Recipe {
	return &domain.Recipe{
		ID:    u.ID,
		Title: u.Title,
	}
}

type RecipeRows []RecipeRow

func (r RecipeRows) ToDomain() domain.Recipes {
	recipes := make(domain.Recipes, len(r))

	for i, allRecipes := range r {
		recipes[i] = *allRecipes.ToDomain()
	}
	return recipes
}

func (s *SpoonacularClient) GetSpoonacularRecipe(ctx context.Context, cuisine string) (domain.Recipes, error) {
	apiKey := os.Getenv("SPOONACULAR_API_KEY")
	url := fmt.Sprintf("https://api.spoonacular.com/recipes/complexSearch?cuisine=%s&apiKey=%s", cuisine, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response struct {
		Recipes []RecipeRow `json:"results"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return RecipeRows(response.Recipes).ToDomain(), nil
}
