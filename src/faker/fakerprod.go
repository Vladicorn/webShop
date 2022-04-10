package main

import (
	"casic/src/database"
	"casic/src/models"
	"log"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func main() {

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 40; i++ {
		product := models.Product{
			Title:       faker.Username(),
			Description: faker.Username(),
			Image:       faker.URL(),
			Price:       float64(rand.Intn(90) * 10),
		}

		database.InsertDBProduct(&product)
	}
}
