package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// way to read environment variables
	// by default it reads from shell
	// to allow  it read from .env add godotenv
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in the environment")
	}
	// it creates a router :- main router
	router := chi.NewRouter()

	// default cors object, you can get from docs
	router.Use(
		cors.Handler(
			cors.Options{
				AllowedOrigins:   []string{"https://*", "http://*"},
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders:   []string{"*"},
				ExposedHeaders:   []string{"Link"},
				AllowCredentials: false,
				MaxAge:           300,
			}))

	// sub router, v1Rputer.Group can have one middleware auth check for every route that comes under that group
	v1Router := chi.NewRouter()
	// HandleFunc will give same response for any type of req like get, post etc
	//v1Router.HandleFunc("/health", handlerReadiness)
	v1Router.Get("/health", handlerReadiness)
	v1Router.Get("/err", handlerError)
	router.Mount("/v1", v1Router)
	// now we will create a http server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	// it will be blocked at this function
	// as now server starts serving request
	log.Printf("Server starting on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Port:", portString)
}
