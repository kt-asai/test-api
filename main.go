package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/applications", appsHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type AppsResponse struct {
	Items []App `json:"items"`
}

func appsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	repository := NewAppRepository()
	service := NewAppSerice(repository)

	apps := service.list()

	ar := AppsResponse{
		Items: apps,
	}

	if err := json.NewEncoder(w).Encode(ar); err != nil {
		log.Fatal(err)
	}
}

// ---------------------------------
// Model

type App struct {
	Name string `json:"name"`
}

type AppRepositoryInterface interface {
	list() []App
}

// ---------------------------------
// Repository

func NewAppRepository() *AppRepository {
	client := "firestore client"

	return &AppRepository{
		client: client,
	}
}

type AppRepository struct {
	client string
}

func (ar AppRepository) list() []App {
	apps := []App{{Name: "vsapi"}, {Name: "vsui"}}

	return apps
}

// ---------------------------------
// Application Service

func NewAppSerice(repository AppRepositoryInterface) *AppService {
	return &AppService{
		repository: repository,
	}
}

type AppService struct {
	repository AppRepositoryInterface
}

func (as AppService) list() []App {
	return as.repository.list()
}
