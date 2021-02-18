package models

type Car struct {
	ID                      int
	Price                   float64
	Color                   string
	Mileage                 float64
	CarNumber               string
	PhoneNumbers            string
	CarPhotos               string
	OwnerPassportPhotos     string
	TechnicalPassportPhotos string
	KeysPhoto               string
	SoldEmployeeID          int
	BoughtClientID          int
	BoughtClient            Client
	CarModel                CarModel
	SoldEmployee            Employee
	CarComments             []CarComment
	Expenditures            []Expenditure
}
