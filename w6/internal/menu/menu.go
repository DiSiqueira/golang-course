package menu

import (
	"encoding/json"
	"fmt"
	"github.com/disiqueira/golang-course/w6/internal/model"
	"github.com/pkg/errors"
	"net/http"
)

type (
	Service struct {
		baseURL string
		country string
		locale string
		jwt string

		httpClient httpExecuter
	}

	httpExecuter interface {
		Do(req *http.Request) (*http.Response, error)
	}

	Result struct {
		Items []struct {
			Link           string      `json:"link"`
			IsActive       bool        `json:"isActive"`
			Courses        []struct {
				Recipe struct {
					Name            string  `json:"name"`
					ImageLink       string  `json:"imageLink"`
					ImagePath       string  `json:"imagePath"`
					AverageRating   float64 `json:"averageRating"`
					Active          bool    `json:"active"`
				} `json:"recipe"`
			} `json:"courses"`
		} `json:"items"`
	}

)

func NewService(baseURL, country, locale, jwt string, httpClient httpExecuter) *Service {
	return &Service{
		baseURL: baseURL,
		country: country,
		locale: locale,
		httpClient: httpClient,
		jwt: jwt,
	}
}

func (s *Service) Search(year int, week int) (*model.Menu, error) {
	finalURL := fmt.Sprintf(
		"%s/gw/menus-service/menus?country=%s&locale=%s&week=%d-W%d",
		s.baseURL,
		s.country,
		s.locale,
		year,
		week)

	req, err := http.NewRequest(http.MethodGet, finalURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "can not create request to menus service")
	}

	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", s.jwt))

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "can not execute request to menus service")
	}

	var result Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, errors.Wrap(err, "can not decode result from menus service")
	}

	menu := &model.Menu{}

	for _, item := range result.Items {
		if !item.IsActive {
			continue
		}

		for _, course := range item.Courses {
			if !course.Recipe.Active {
				continue
			}

			recipe := model.Recipe{
				Name:   course.Recipe.Name,
				URL:    course.Recipe.ImageLink,
				Rating: course.Recipe.AverageRating,
			}

			menu.Recipes = append(menu.Recipes, recipe)
		}
	}

	return menu, nil
}

