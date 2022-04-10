package models

type Order struct {
	Id              uint        `json:"id"`
	TransactionId   string      `json:"transactionid"`
	UserId          uint        `json:"user_id"`
	Code            string      `json:"code"`
	AmbassadorEmail string      `json:"ambassador_email"`
	FirstName       string      `json:"first_name"`
	LastName        string      `json:"last_name"`
	Email           string      `json:"email"`
	Adress          string      `json:"adress"`
	City            string      `json:"city"`
	Country         string      `json:"country"`
	Zip             string      `json:"zip"`
	Complete        bool        `json:"complete"`
	OrderItem       []OrderItem `gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	Id                uint    `json:"id"`
	OrderId           uint    `json:"order_id"`
	ProductTitle      string  `json:"product_title"`
	Price             float64 `json:"price"`
	Quantity          uint    `json:"quantity"`
	AdminRevenue      float64 `json:"admin_revenue"`
	AmbassadorRevenue float64 `json:"ambassador_revenue"`
}
