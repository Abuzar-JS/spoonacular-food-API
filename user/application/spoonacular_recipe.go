package application

import (
	"context"

	"github.com/Abuzar-JS/go-spoonacular-api/user/domain"
	"github.com/Abuzar-JS/go-spoonacular-api/user/domain/user"
)

type SpoonacularRecipe func(ctx context.Context, cuisine string) (domain.Recipes, error)

func NewSpoonacularRecipe(
	RecipeRepo user.RecipeRepository,
) SpoonacularRecipe {
	return func(ctx context.Context, cuisine string) (domain.Recipes, error) {

		recipes, err := RecipeRepo.GetSpoonacularRecipe(ctx, cuisine)

		if err != nil {
			return nil, err
		}

		return recipes, nil

	}
}
