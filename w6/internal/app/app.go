package app

import (
	"fmt"
	"github.com/disiqueira/golang-course/w6/internal/handler"
	"github.com/disiqueira/golang-course/w6/internal/menu"
	"github.com/go-chi/chi"
	"net/http"
	"time"
)
type (
	App struct {

	}
)

func NewApp() *App {
	return &App{}
}

func (a App) Run() error {
	fmt.Println("Starting Service on 8080")
	baseURL := "https://www-staging.hellofresh.com"
	country := "us"
	locale := "en-US"
	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDIyMDE0NjQsImp0aSI6IjI0Y2QyODBkLWQ2ZDYtNGQyNS1hN2QxLTVkOGFhMmFmNTBmMiIsImlhdCI6MTYwMjE3OTg2NCwiaXNzIjoic2VuZiJ9.Pj-xTA8fb5wc42a_9Y7QC6Qh48vzKeSqAmG-KAgIjYQ"

	httpClient := &http.Client{
		Timeout: 3 * time.Second,
	}

	menuServiceClient := menu.NewService(baseURL, country, locale, jwt, httpClient)
	menuHandler := handler.NewMenu(menuServiceClient)

	router := chi.NewRouter()
	router.Get("/menu/{year}/week/{week}", menuHandler.ServeHTTP)

	return http.ListenAndServe(":8080", router)
}

