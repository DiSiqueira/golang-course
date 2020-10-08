package model

type (
	Menu struct {
		Recipes []Recipe `json:"recipes"`
	}

	Recipe struct {
		Name   string  `json:"name"`
		URL    string  `json:"url"`
		Rating float64 `json:"rating"`
	}
)
