package app

import (
	"database/sql"
	"fmt"
	"github.com/disiqueira/golang-course/w6/internal/cache"
	"github.com/disiqueira/golang-course/w6/internal/handler"
	"github.com/disiqueira/golang-course/w6/internal/menu"
	"github.com/disiqueira/golang-course/w6/internal/signal"
	"github.com/disiqueira/golang-course/w6/internal/storage"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
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
	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjM0MWYzMzZjLTA3NWYtNDQ2Ny04M2E5LWYxYTAzODc3M2U0NiIsInVzZXJuYW1lIjoic2ZzZnNAaGYuY29tIiwiZW1haWwiOiJzZnNmc0BoZi5jb20iLCJjb3VudHJ5IjoidXMiLCJibG9ja2VkIjpmYWxzZSwibWV0YWRhdGEiOnsibmFtZSI6IkRpZWdvIERpcXVlaXJhIiwicGFzc3dvcmRsZXNzIjpmYWxzZX0sInJvbGVzIjpbXSwiZXhwIjoxNjAyODgyNjc1LCJqdGkiOiI0YjA4MjQwNC1hNzliLTQwNGMtYjkwZi03NjEyZGNkOGNjNmQiLCJpYXQiOjE2MDI4ODA4NzUsImlzcyI6IjEyMzQiLCJzdWIiOiIzNDFmMzM2Yy0wNzVmLTQ0NjctODNhOS1mMWEwMzg3NzNlNDYifQ.gwwtUE1tI86-vkueZqoDOZPJ91zM7v3Rp4BP0XFtWtA"

	httpClient := &http.Client{
		Timeout: 3 * time.Second,
	}

	dsn := "postgres://mps:hello123@menu-proxy-database:5432/menu_proxy_service?sslmode=disable"
	dbClient, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	recipeRepository := storage.NewRecipeRepository(dbClient)
	recipeCache := cache.NewRecipe(recipeRepository)
	menuServiceClient := menu.NewService(baseURL, country, locale, jwt, httpClient)
	menuHandler := handler.NewMenu(menuServiceClient, recipeCache)

	router := chi.NewRouter()
	router.Get("/menu/{year}/week/{week}", menuHandler.ServeHTTP)

	go func() {
		if err := http.ListenAndServe(":8080", router); err != nil {
			log.Fatal(err)
		}
	}()

	signal.Term()
	fmt.Println("Termination signal received.")

	return nil
}
