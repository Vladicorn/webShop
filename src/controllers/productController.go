package controllers

import (
	"casic/src/database"
	"casic/src/models"
	"context"
	"encoding/json"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Products(c *fiber.Ctx) error {
	products := database.SelectDBProducts()
	return c.JSON(products)
}

func CreateProduct(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	price, _ := strconv.ParseFloat(data["price"], 64)

	product := models.Product{
		Title:       data["title"],
		Description: data["description"],
		Image:       data["image"],
		Price:       price,
	}

	if err := database.InsertDBProduct(&product); err != nil {
		return err
	}
	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := database.SelectDBProduct(uint(id))
	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	price, _ := strconv.ParseFloat(data["price"], 64)
	product := models.Product{
		Id:          uint(id),
		Title:       data["title"],
		Description: data["description"],
		Image:       data["image"],
		Price:       price,
	}

	if err := database.UpdateDBProduct(&product); err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	if err := database.DeleteDBProduct(uint(id)); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

func ProductFont(c *fiber.Ctx) error {
	var products []models.Product
	ctx := context.Background()

	result, err := database.Cache.Get(ctx, "product_front").Result()
	if err != nil {
		products = database.SelectDBProducts()
		bytes, err := json.Marshal(products)
		if err != nil {
			log.Println("tyt 1     ", err)
		}
		if errKey := database.Cache.Set(ctx, "product_front", bytes, 30*time.Minute).Err(); errKey != nil {
			log.Println("tyt 2    ", errKey)
		}
	} else {
		json.Unmarshal([]byte(result), &products)
	}

	return c.JSON(products)
}

func ProductBackend(c *fiber.Ctx) error {
	var products []models.Product
	ctx := context.Background()

	result, err := database.Cache.Get(ctx, "product_back").Result()
	if err != nil {
		products = database.SelectDBProducts()
		bytes, err := json.Marshal(products)
		if err != nil {
			log.Println("tyt 1     ", err)
		}
		if errKey := database.Cache.Set(ctx, "product_back", bytes, 30*time.Minute).Err(); errKey != nil {
			log.Println("tyt 2    ", errKey)
		}
	} else {
		json.Unmarshal([]byte(result), &products)
	}

	query := c.Query("c")
	query = strings.ToLower(query)
	var filterproducts []models.Product
	if query != "" {
		for _, product := range products {
			if strings.Contains(strings.ToLower(product.Title), query) || strings.Contains(strings.ToLower(product.Description), query) {
				filterproducts = append(filterproducts, product)
			}
		}

	} else {
		filterproducts = products
	}

	if query != "" {
		for _, product := range products {
			if strings.Contains(strings.ToLower(product.Title), query) || strings.Contains(strings.ToLower(product.Description), query) {
				filterproducts = append(filterproducts, product)
			}
		}
	} else {
		filterproducts = products
	}

	sortQuery := c.Query("sort")
	sortQuery = strings.ToLower(sortQuery)

	if sortQuery == "asc" {
		sort.Slice(filterproducts, func(i, j int) bool {
			return filterproducts[i].Price < filterproducts[j].Price
		})
	} else if sortQuery == "dec" {
		sort.Slice(filterproducts, func(i, j int) bool {
			return filterproducts[i].Price > filterproducts[j].Price
		})
	}

	page, _ := strconv.Atoi(c.Query("page"))
	totalPage := len(filterproducts)
	perPage := 9
	var data []models.Product
	if totalPage <= page*perPage {

	}

	return c.JSON(filterproducts)
}
