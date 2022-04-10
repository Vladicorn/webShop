package database

import (
	"casic/src/models"
)

func SelectDBProducts() []models.Product {
	rows, _ := DB.Query("SELECT id, title, description, image,price FROM products")
	defer rows.Close()
	products := []models.Product{}
	for rows.Next() {
		product := models.Product{}
		rows.Scan(&product.Id, &product.Title, &product.Description, &product.Image, &product.Price)

		products = append(products, product)
	}
	return products
}

func InsertDBProduct(product *models.Product) error {
	_, err := DB.Query("INSERT INTO products (title, description, image,price)VALUES ($1, $2, $3, $4)", product.Title, product.Description, product.Image, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func SelectDBProduct(id uint) models.Product {
	rows, _ := DB.Query("SELECT id, title, description, image,price FROM products WHERE id = $1;", id)

	defer rows.Close()
	//var User models.User
	product := models.Product{}
	for rows.Next() {
		rows.Scan(&product.Id, &product.Title, &product.Description, &product.Image, &product.Price)
	}
	return product
}

func UpdateDBProduct(product *models.Product) error {
	_, err := DB.Query("UPDATE products SET title = $2, description = $3, image = $4, price = $5 WHERE id = $1", product.Id, product.Title, product.Description, product.Image, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func DeleteDBProduct(id uint) error {
	_, err := DB.Query("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
