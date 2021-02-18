package models

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type Client struct {
	ID             int    `json:"id"`
	Firstname      string `json:"firstname" validate:"required"`
	Lastname       string `json:"lastname" validate:"required"`
	DateOfBirth    string `json:"date_of_birth" validate:"required"`
	Avatar         string `json:"avatar" validate:"required"`
	PhoneNumbers   string `json:"phone_numbers" validate:"required"`
	PassportPhotos string `json:"passport_photos"`
}

//func (client *Client) MarshalJSON() ([]byte,error){
//	return
//}

func (client Client) Validate() (errors []string) {
	validate := validator.New()
	err := validate.Struct(client)
	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("Field  %s(%s) is %s", err.Field(), err.Type(), err.Tag()))
		}
		return errors
	}
	return nil
}
