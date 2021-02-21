package controllers

import (
	"../database"
	"../models"
	"encoding/json"
	"net/http"
)

func GetCars(w http.ResponseWriter, r *http.Request) {
	cars := database.GetCars(1, 20)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(Response{
		Data:      cars,
		IsSuccess: true,
		Errors:    nil,
	})
}

func NewCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//errors := car.Validate()
	//if errors != nil {
	//	res := Response{
	//		Data:      nil,
	//		IsSuccess: false,
	//		Errors:    errors,
	//	}
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(http.StatusUnprocessableEntity)
	//	err = json.NewEncoder(w).Encode(res)
	//	if err != nil {
	//		panic(err)
	//	}
	//	return
	//}

	database.AddCar(&car)
}
