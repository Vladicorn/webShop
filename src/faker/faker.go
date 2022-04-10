package main

import (
	"casic/src/database"
	"casic/src/models"
	"log"

	"github.com/bxcodec/faker/v3"
)

func main() {

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 40; i++ {
		user := models.User{
			FirstName:    faker.FirstName(),
			LastName:     faker.LastName(),
			Email:        faker.Email(),
			IsAmbassador: true,
		}
		user.SetPswd("1234")
		database.InsertDB(&user)
	}
}
