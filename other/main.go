package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

type (
	response struct {
		ErrorCode string `json:"code,omitempty"`
		Message string `json:"message"`
	}
)

type (
	mux struct {
		routes map[string]map[string]http.HandlerFunc
	}
)

func (m mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handlers, ok := m.routes[req.URL.Path]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	handler, ok := handlers[req.Method]
	if !ok {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, req)
}

func main() {
	m := &mux{
		routes: make(map[string]map[string]http.HandlerFunc),
	}

	go func() {
		if err := http.ListenAndServe(":8080", m); err != nil {
			panic(err)
		}
	}()

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt)

	<-stopSignal
	fmt.Println("kill signal received.")

	os.Exit(0)
}

func handlerHello(w http.ResponseWriter, _ *http.Request) {
	r := response{
		Message: "Fresh",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(r)
}

func handlerFresh(w http.ResponseWriter, _ *http.Request) {
	r := response{
		Message: "Hello",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(r)
}
