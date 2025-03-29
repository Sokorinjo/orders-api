package main

import (
	"fmt"
	"net/http"

	// "github.com/go-chi/chi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	//Using &, meaning it's storing it as a pointer and taking it's memory addres, not it's value.
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Get("/home", basicHandler)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to listen to server:", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
