package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
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
	    passport_photos text
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
	    owner_id int,
	    price float, 
	    color text,
	    mileage float,
	    car_number text, 
	    car_photos string,
	    technical_passport_photos string,
	    keys_photos string
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
	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}
	return err
}
