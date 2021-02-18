package main

import (
	"./database"
	"./router"
	"net/http"
)

func main() {
	_ = database.ConfigureDatabase()
	defer database.DB.Close()
	r :=router.NewRouter()
	_ = http.ListenAndServe(":8090", r)
}
