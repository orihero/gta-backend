package database

import (
	"../models"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
)

var DB *sql.DB

func ConfigureDatabase() error {
	var err error
	DB, err = sql.Open("sqlite3", "./db.db")
	if err != nil {
		log.Fatal(err)
	}
	sqlStmt := `
	create table if not exists clients (
	    id integer not null primary key, 
	    firstname text, 
	    lastname text, 
	    date_of_birth text, 
	    avatar text, 
	    phone_numbers text,
	    passport_photos text,
	    created_at timestamp  NOT NULL  DEFAULT current_timestamp,
  		updated_at timestamp  NOT NULL  DEFAULT current_timestamp
	    );
	create table if not exists employee (
	    id integer not null primary key, 
	    firstname text, 
	    lastname text, 
	    date_of_birth text, 
	    avatar text, 
	    phone_numbers text,
	    salary float 
	    );
	create table if not exists  cars (
	    id integer not null primary key, 
	    model_id integer, 
	    owner_id integer,
	    price text, 
	    color text,
	    mileage text,
	    car_number text, 
	    car_photos string,
	    technical_passport_photos string,
	    keys_photos string,
	    created_at timestamp  NOT NULL  DEFAULT current_timestamp,
  		updated_at timestamp  NOT NULL  DEFAULT current_timestamp
	    );
	create table if not exists  expenditure (
	    id integer not null primary key, 
	    amount float,
	    car_id integer,
		foreign key (car_id) references cars(id)
	    );
	create table if not exists  comments (
	    id integer not null primary key, 
	    title text,
	    content text,
	    car_id integer, 
	  	foreign key (car_id) references cars(id)            
	    );
	`
	seed()
	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}
	return err
}

func seed() {
	var logos []models.CarModel
	var model []models.CarModel
	bytes, err := ioutil.ReadFile("car-logos.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(bytes, &logos)
	bytes, err = ioutil.ReadFile("car-models.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(bytes, &model)
	if err != nil {
		fmt.Println(err)
	}
	for _, el := range logos {
		for _, e := range model {
			if el.Name == e.Brand {
				el.Models = e.Models
			}
		}
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on

}
