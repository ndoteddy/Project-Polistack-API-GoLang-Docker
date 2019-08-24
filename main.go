package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// use PORT environment variable, or default to 8080
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	router := mux.NewRouter()
	restaurant = append(restaurant,
		Restaurant{ID: "1", Name: "Ayam Super Nando", Address: "Jalan Bulan 10th Kuala Lumpur", Follower: &Follower{Total: "10", Like: "5"}})
	restaurant = append(restaurant,
		Restaurant{ID: "2", Name: "Kopi Hoho Nando", Address: "Central Boulevard 20th Singapore ", Follower: &Follower{Total: "25", Like: "2"}})

	router.HandleFunc("/restaurant", GetRestaurantsEndpoint).Methods("GET")
	router.HandleFunc("/restaurant/{id}", GetRestaurantEndpoint).Methods("GET")

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe(":"+port, router)
	log.Fatal(err)
}

type Restaurant struct {
	ID       string    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Address  string    `json:"address,omitempty"`
	Follower *Follower `json:"follower,omitempty"`
}

type Follower struct {
	Total string `json:"total,omitempty"`
	Like  string `json:"like,omitempty"`
}

var restaurant []Restaurant

func GetRestaurantEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range restaurant {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Restaurant{})
}

func GetRestaurantsEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(restaurant)
}
