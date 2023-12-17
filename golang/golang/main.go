package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Item represents a simple entity for this example
type Item struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var items []Item

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Define API endpoints
	r.HandleFunc("/items", getItems).Methods("GET")
	r.HandleFunc("/items/{id}", getItem).Methods("GET")
	r.HandleFunc("/items", createItem).Methods("POST")
	r.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	// Start the server
	port := 8080
	fmt.Printf("Server started on :%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

// Handler functions

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range items {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Item{}) // Return an empty Item if not found
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newItem Item
	_ = json.NewDecoder(r.Body).Decode(&newItem)
	items = append(items, newItem)
	json.NewEncoder(w).Encode(newItem)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range items {
		if item.ID == params["id"] {
			var updatedItem Item
			_ = json.NewDecoder(r.Body).Decode(&updatedItem)
			items[index] = updatedItem
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}
	json.NewEncoder(w).Encode(&Item{}) // Return an empty Item if not found
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range items {
		if item.ID == params["id"] {
			items = append(items[:index], items[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(items)
}
