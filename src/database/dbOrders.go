package database

import (
	"casic/src/models"
	"fmt"
)

/*
func SelectDBOrders() []models.Order {
	rows, _ := DB.Query("SELECT orders.id, orders.transaction_id, orders.user_id, orders.code,orders.ambassador_email, orders.first_name, orders.last_name, orders.email, orders.adress, orders.city, orders.country, orders.zip, orders.complete, order_items.id, order_items.order_id,order_items.product_title,order_items.price,order_items.quantity,order_items.admin_revenue,order_items.ambassador_revenue FROM orders JOIN order_items ON order_items.order_id =  orders.id")
	defer rows.Close()
	orders := []models.Order{}
	for rows.Next() {
		order := models.Order{}
		order_item := models.OrderItem{}
		rows.Scan(&order.Id, &order.TransactionId, &order.UserId, &order.Code, &order.AmbassadorEmail, &order.FirstName, &order.LastName, &order.Email, &order.Adress, &order.City, &order.Country, &order.Zip, &order.Complete, &order_item.Id, &order_item.Id, &order_item.OrderId, &order_item.ProductTitle, &order_item.Price, &order_item.Quantity, &order_item.AdminRevenue, &order_item.AmbassadorRevenue)
		fmt.Println(rows)
		orders = append(orders, order)
	}
	return orders
}*/

func SelectDBOrders() []models.Order {
	rows, _ := DB.Query("SELECT orders.id, orders.transaction_id, orders.user_id, orders.code,orders.ambassador_email, orders.first_name, orders.last_name, orders.email, orders.adress, orders.city, orders.country, orders.zip, orders.complete FROM orders")
	defer rows.Close()
	orders := []models.Order{}
	for rows.Next() {
		order := models.Order{}
		rows.Scan(&order.Id, &order.TransactionId, &order.UserId, &order.Code, &order.AmbassadorEmail, &order.FirstName, &order.LastName, &order.Email, &order.Adress, &order.City, &order.Country, &order.Zip, &order.Complete)
		rows2, _ := DB.Query("SELECT order_items.id, order_items.order_id, order_items.product_title, order_items.price,order_items.quantity, order_items.admin_revenue, order_items.ambassador_revenue FROM order_items WHERE order_id = $1", order.Id)
		order_items := []models.OrderItem{}
		//fmt.Println(rows2)
		for rows2.Next() {
			order_item := models.OrderItem{}
			rows2.Scan(&order_item.Id, &order_item.OrderId, &order_item.ProductTitle, &order_item.Price, &order_item.Quantity, &order_item.AdminRevenue, &order_item.AmbassadorRevenue)
			order_items = append(order_items, order_item)
		}
		order.OrderItem = order_items
		orders = append(orders, order)
	}
	return orders
}

func InsertDBOrder(order *models.Order) error {
	_, err := DB.Query("INSERT INTO orders (transaction_id, user_id, code, ambassador_email, first_name, last_name, email, complete)VALUES ($1, $2, $3, $4, $5, $6, $7,$8)", order.TransactionId, order.UserId, order.Code, order.AmbassadorEmail, order.FirstName, order.LastName, order.Email, order.Complete)
	if err != nil {
		return err
	}
	return nil
}

func InsertDBOrderItem(order *models.OrderItem) error {
	_, err := DB.Query("INSERT INTO order_items (order_id, product_title, price, quantity, admin_revenue, ambassador_revenue)VALUES ($1, $2, $3, $4, $5, $6)", order.OrderId, order.ProductTitle, order.Price, order.Quantity, order.AdminRevenue, order.AmbassadorRevenue)
	if err != nil {
		return err
	}
	return nil
}

func SelectDBOrder(id uint) []models.Order {
	rows, _ := DB.Query("SELECT orders.id, orders.transaction_id, orders.user_id, orders.code,orders.ambassador_email, orders.first_name, orders.last_name, orders.email, orders.adress, orders.city, orders.country, orders.zip, orders.complete FROM orders WHERE id = $1 AND complete = true", id)
	defer rows.Close()
	orders := []models.Order{}
	fmt.Println(rows)
	for rows.Next() {
		order := models.Order{}
		rows.Scan(&order.Id, &order.TransactionId, &order.UserId, &order.Code, &order.AmbassadorEmail, &order.FirstName, &order.LastName, &order.Email, &order.Adress, &order.City, &order.Country, &order.Zip, &order.Complete)
		rows2, _ := DB.Query("SELECT order_items.id, order_items.order_id, order_items.product_title, order_items.price,order_items.quantity, order_items.admin_revenue, order_items.ambassador_revenue FROM order_items WHERE order_id = $1", order.Id)
		order_items := []models.OrderItem{}
		for rows2.Next() {
			order_item := models.OrderItem{}
			rows2.Scan(&order_item.Id, &order_item.OrderId, &order_item.ProductTitle, &order_item.Price, &order_item.Quantity, &order_item.AdminRevenue, &order_item.AmbassadorRevenue)
			order_items = append(order_items, order_item)
		}
		order.OrderItem = order_items
		orders = append(orders, order)
	}
	return orders
}
