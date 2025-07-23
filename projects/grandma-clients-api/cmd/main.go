package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/valdir-alves3000/go-developer-training-dio/projets/grandma-clients-api/config"
	"github.com/valdir-alves3000/go-developer-training-dio/projets/grandma-clients-api/internal/handlers"
	"github.com/valdir-alves3000/go-developer-training-dio/projets/grandma-clients-api/internal/repositories"
	"github.com/valdir-alves3000/go-developer-training-dio/projets/grandma-clients-api/internal/usecases"
)

const (
	port = ":8080"
)

func main() {
	_, err := config.Initialize()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	router := initializeRouter()
	initializeRoutes(router)

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func initializeRouter() *mux.Router {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handleHomePage)
	return router
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	userRepo := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	tmpl, err := template.ParseFiles("./template/index.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	people, err := userUsecase.FindAll()
	if err != nil {
		http.Error(w, "Unable to fetch users", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, people); err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
		return
	}
}

func initializeRoutes(router *mux.Router) {
	db := config.GetDB()
	userRepo := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userHandler := handlers.NewUserHandler(userUsecase)

	apiRouter := router.PathPrefix("/api/users").Subrouter()
	apiRouter.HandleFunc("", userHandler.CreateUser).Methods("POST")
	apiRouter.HandleFunc("", userHandler.ListUsers).Methods("GET")
	apiRouter.HandleFunc("/{id}", userHandler.GetUser).Methods("GET")
	apiRouter.HandleFunc("/{id}", userHandler.UpdateUser).Methods("PUT")
	apiRouter.HandleFunc("/{id}", userHandler.DeleteUser).Methods("DELETE")
}
