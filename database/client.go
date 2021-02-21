package database

import (
	"../models"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func GetClients(page, pageSize int) *[]models.Client {
	var clients []models.Client
	smt, err := DB.Prepare(`select * from clients limit ?`)
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
		cl := models.Client{}
		err = rows.Scan(&cl.ID, &cl.Firstname, &cl.Lastname, &cl.DateOfBirth, &cl.Avatar, &cl.PhoneNumbers, &cl.PassportPhotos, &cl.CreatedAt, &cl.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		clients = append(clients, cl)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return &clients
}

func AddClient(client *models.Client) {
	tx, err := DB.Begin()
	stmt, err := DB.Prepare("insert into clients(firstname, lastname, date_of_birth, avatar, phone_numbers,passport_photos) values(?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(client.Firstname, client.Lastname, client.DateOfBirth, client.Avatar, client.PhoneNumbers, client.PassportPhotos)
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
}

func UpdateClient(client *models.Client) {

}

var js = `{
    "firstname": "lol",
    "lastname": "foo",
    "date_of_birth": "10-12-2020",
    "avatar": {
        "path": "https://www.slashfilm.com/wp/wp-content/images/avatar-2-story.jpg"
    },
    "phone_numbers": [
        {
            "value": "+998932300500"
        },
        {
            "value": "+998998970597"
        }
    ]
}`
