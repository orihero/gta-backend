package database

import (
	"../models"
	"log"
)

func GetCars(page, pageSize int) *[]models.Car {
	var clients []models.Car
	smt, err := DB.Prepare(`select * from cars limit ?`)
	if err != nil {
		log.Fatal(err)
	}
	defer smt.Close()
	rows, err := smt.Query(page * pageSize)
	if err != nil || rows == nil {
		return &clients
	}
	defer rows.Close()
	for rows.Next() {
		car := models.Car{}
		err = rows.Scan(&car.ID, &car.CarModelID, &car.BoughtClientID, &car.Price, &car.Color, &car.Mileage, &car.CarNumber, &car.CarPhotos, &car.TechnicalPassportPhotos, &car.KeysPhoto, &car.CreatedAt, &car.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		clients = append(clients, car)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return &clients
}

func AddCar(car *models.Car) {
	tx, err := DB.Begin()
	stmt, err := DB.Prepare("insert into cars(price,mileage,color,car_number,car_photos,technical_passport_photos,keys_photos) values(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(car.Price, car.Mileage, car.Color, car.CarNumber, car.CarPhotos, car.TechnicalPassportPhotos, car.KeysPhoto)
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
}
