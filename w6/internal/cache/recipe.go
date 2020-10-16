package cache

import (
	"fmt"
	"github.com/disiqueira/golang-course/w6/internal/model"
)

type (
	Recipe struct {
		fallback   model.RecipeRepository
		memoryData map[string]*model.Recipe
	}
)

func NewRecipe(fallback model.RecipeRepository) *Recipe {
	return &Recipe{
		fallback:   fallback,
		memoryData: map[string]*model.Recipe{},
	}
}

func (r *Recipe) Exist(url string) bool {
	_, err := r.Find(url)
	return err == nil
}

func (r *Recipe) Find(url string) (*model.Recipe, error) {
	recipe, ok := r.memoryData[url]
	if ok {
		fmt.Println("cache hit", url)
		return recipe, nil
	}
	fmt.Println("cache miss", url)

	recipe, err := r.fallback.Find(url)
	if err != nil {
		return nil, err
	}

	r.memoryData[recipe.URL] = recipe

	return recipe, nil
}

func (r *Recipe) Save(recipe model.Recipe) error {
	return r.fallback.Save(recipe)
}
