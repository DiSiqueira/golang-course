package handler

import (
	"encoding/json"
	"github.com/disiqueira/golang-course/w6/internal/model"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type (
	Menu struct {
		menuServiceClient menuService
	}

	menuService interface {
		Search(year int, week int) (*model.Menu, error)
	}
)

func NewMenu(menuServiceClient menuService) *Menu {
	return &Menu{
		menuServiceClient: menuServiceClient,
	}
}

func (m Menu) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	year, err := strconv.Atoi(chi.URLParam(r, "year"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	week, err := strconv.Atoi(chi.URLParam(r, "week"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := m.menuServiceClient.Search(year, week)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

