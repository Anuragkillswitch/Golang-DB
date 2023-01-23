package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var driver *Driver

func main() {
	var err error
	driver, err = New("data", nil)
	if err != nil {
		// Handle error
	}

	r := mux.NewRouter()

	r.HandleFunc("/{collection}/{resource}", handleCreateOrUpdate).Methods("POST")
	r.HandleFunc("/{collection}/{resource}", handleRead).Methods("GET")
	r.HandleFunc("/{collection}", handleReadAll).Methods("GET")

	http.ListenAndServe(":3000", r)
}

func handleCreateOrUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collection := vars["collection"]
	resource := vars["resource"]

	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
		return
	}

	if err := driver.Write(collection, resource, data); err != nil {
		http.Error(w, "Error writing data to the database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func handleRead(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collection := vars["collection"]
	resource := vars["resource"]

	var data map[string]interface{}
	if err := driver.Read(collection, resource, &data); err != nil {
		http.Error(w, "Error reading data from the database", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding JSON data", http.StatusInternalServerError)
		return
	}
}

func handleReadAll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collection := vars["collection"]

	resources, err := driver.ReadAll(collection)
	if err != nil {
		http.Error(w, "Error reading data from the database", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(resources); err != nil {
		http.Error(w, "Error encoding JSON data", http.StatusInternalServerError)
		return
	}
}
