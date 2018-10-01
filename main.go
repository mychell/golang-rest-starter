package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mychell/go-rest-starter/driver"
	"github.com/mychell/go-rest-starter/handlers"
)

func main() {

	connection, err := driver.ConnectPQ("development")

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	postHandler := handlers.NewPostHandler(connection)
	r.Post("/posts", postHandler.Create)

	http.ListenAndServe(":8005", r)
}
