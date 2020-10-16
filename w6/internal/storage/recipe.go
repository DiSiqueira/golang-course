package storage

import (
	"database/sql"
	"github.com/disiqueira/golang-course/w6/internal/model"
)

type (
	RecipeRepository struct {
		dbClient *sql.DB
	}
)

func NewRecipeRepository(dbClient *sql.DB) *RecipeRepository {
	return &RecipeRepository{dbClient: dbClient}
}

func (r RecipeRepository) Save(recipe model.Recipe) error {
	query := `
		INSERT INTO recipe
		 (
		 	 "name",
		 	  url, 
		 	  rating
		  ) VALUES(
			 $1, $2, $3
		  );
`

	stmt, err := r.dbClient.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(recipe.Name, recipe.URL, recipe.Rating)
	return err
}

func (r RecipeRepository) Find(url string) (*model.Recipe, error) {
	query := `
		SELECT 
			"name",
			url, 
			rating
		FROM
			recipe
		WHERE
			url = $1
`

	var recipe model.Recipe

	row := r.dbClient.QueryRow(query, url)
	err := row.Scan(&recipe.Name, &recipe.URL, &recipe.Rating)
	return &recipe, err
}

func (r RecipeRepository) Exist(url string) bool {
	_, err := r.Find(url)
	return err == nil
}
