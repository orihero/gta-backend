package controllers

import (
	"../database"
	"../models"
	"encoding/json"
	"net/http"
)

func GetClients(w http.ResponseWriter, r *http.Request) {
	clients := database.GetClients(1, 20)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(Response{
		Data:      clients,
		IsSuccess: true,
		Errors:    nil,
	})
}

func NewClient(w http.ResponseWriter, r *http.Request) {
	var newClient models.Client
	err := json.NewDecoder(r.Body).Decode(&newClient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	errors := newClient.Validate()
	if errors != nil {
		res := Response{
			Data:      nil,
			IsSuccess: false,
			Errors:    errors,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			panic(err)
		}
		return
	}

	database.AddClient(&newClient)
}
