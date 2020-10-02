package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/go-chi/chi"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	userStorage struct {
		list []user
		mtx  *sync.Mutex
	}

	userManager struct {
		db *userStorage
	}
)

func newUserStorage() *userStorage {
	return &userStorage{
		mtx: &sync.Mutex{},
	}
}

func (us *userStorage) insertUser(name string) {
	us.mtx.Lock()
	defer us.mtx.Unlock()

	user := user{
		ID:   len(us.list),
		Name: name,
	}

	us.list = append(us.list, user)
}

func (us *userStorage) deleteUser(id int) {
	us.mtx.Lock()
	defer us.mtx.Unlock()

	for index, user := range us.list {
		if user.ID == id {
			us.list = append(us.list[:index], us.list[index+1:]...)
			return
		}
	}
}

/*
cmd/user-service/main.go
internal/app/app.go
internal/handlers/user/get.go
internal/handlers/user/post.go
internal/storage/db.go
*/

func main() {
	db := newUserStorage()

	userManager := &userManager{
		db: db,
	}

	routerChi := chi.NewRouter()

	routerChi.Get("/user", userManager.userGET)
	routerChi.Post("/user", userManager.userPOST)
	routerChi.Delete("/user/{userID}", userManager.userDELETE)

	if err := http.ListenAndServe(":8080", routerChi); err != nil {
		panic(err)
	}
}

type (
	postRequest struct {
		Name string `json:"name"`
	}
)

func (u *userManager) userPOST(w http.ResponseWriter, r *http.Request) {
	pr := &postRequest{}

	if err := json.NewDecoder(r.Body).Decode(pr); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u.db.insertUser(pr.Name)

	w.WriteHeader(http.StatusAccepted)
}

func (u *userManager) userGET(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u.db.list)
}

func (u *userManager) userDELETE(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	id, err := strconv.Atoi(userID)
	if err != nil {
		panic(err)
	}

	u.db.deleteUser(id)

	w.WriteHeader(http.StatusOK)
}
