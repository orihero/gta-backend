package models

import (
	"database/sql"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type Car struct {
	ID                      int           `json:"id"`
	Price                   string        `json:"price" validate:"required"`
	Color                   string        `json:"color" validate:"required"`
	Mileage                 string        `json:"mileage" validate:"required"`
	CarNumber               string        `json:"car_number" validate:"required"`
	PhoneNumbers            string        `json:"phone_numbers"`
	CarPhotos               string        `json:"car_photos"`
	OwnerPassportPhotos     string        `json:"owner_passport_photos"`
	TechnicalPassportPhotos string        `json:"technical_passport_photos"`
	KeysPhoto               string        `json:"keys_photo"`
	SoldEmployeeID          sql.NullInt32 `json:"sold_employee_id"`
	BoughtClientID          sql.NullInt32 `json:"bought_client_id"`
	CarModelID              sql.NullInt32 `json:"car_model_id"`
	BoughtClient            Client        `json:"bought_client"`
	CarModel                CarModel      `json:"car_model"`
	SoldEmployee            Employee      `json:"sold_employee"`
	CarComments             []CarComment  `json:"car_comments"`
	Expenditures            []Expenditure `json:"expenditures"`
	CreatedAt               string        `json:"created_at"`
	UpdatedAt               string        `json:"updated_at"`
}

func (car Car) Validate() (errors []string) {
	validate := validator.New()
	err := validate.Struct(car)
	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("Field  %s(%s) is %s", err.Field(), err.Type(), err.Tag()))
		}
		return errors
	}
	return nil
}
