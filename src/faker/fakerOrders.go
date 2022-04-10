package main

import (
	"casic/src/database"
	"casic/src/models"
	"fmt"
	"log"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func main() {

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	for i := 1; i < 30; i++ {
		var order models.Order
		order.UserId = uint(rand.Intn(30))
		order.Code = faker.Username()
		order.AmbassadorEmail = faker.Username()
		order.FirstName = faker.FirstName()
		order.LastName = faker.LastName()
		order.Email = faker.Email()
		order.Complete = true

		if err := database.InsertDBOrder(&order); err != nil {
			fmt.Println(err)
		}

		price := float64(rand.Intn(90) * 10)
		quantity := uint(rand.Intn(5))
		for j := 0; j < rand.Intn(4)+1; j++ {
			var OrderItem models.OrderItem
			OrderItem.OrderId = uint(i)
			OrderItem.ProductTitle = faker.Word()
			OrderItem.Price = price
			OrderItem.Quantity = quantity
			OrderItem.AdminRevenue = 0.9 * price * float64(quantity)
			OrderItem.AmbassadorRevenue = 0.1 * price * float64(quantity)
			if err := database.InsertDBOrderItem(&OrderItem); err != nil {
				fmt.Println(err)
			}

		}

	}
}
