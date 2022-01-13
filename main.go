package main

import (
	// directory root project go yang kita buat
	"Golang_latihan/test_restapi_crud/models" // memanggil package models pada directory models
	"Golang_latihan/test_restapi_crud/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Kontak{})

	r := routes.SetupRoutes(db)
	r.Run()
}
