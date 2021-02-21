package database

//func AddModel(client *models.CarModel) {
//	tx, err := DB.Begin()
//	stmt, err := DB.Prepare("insert into clients(firstname, lastname, date_of_birth, avatar, phone_numbers,passport_photos) values(?, ?, ?, ?, ?, ?)")
//	if err != nil {
//		log.Fatal(err)
//	}
//	_, err = stmt.Exec(client.Firstname, client.Lastname, client.DateOfBirth, client.Avatar, client.PhoneNumbers, client.PassportPhotos)
//	if err != nil {
//		log.Fatal(err)
//	}
//	err = tx.Commit()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer stmt.Close()
//}
