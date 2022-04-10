package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           uint     `json:"id"`
	FirstName    string   `json:"first_name"`
	LastName     string   `json:"last_name"`
	Email        string   `json:"email"`
	Password     []byte   `json:"-"`
	IsAmbassador bool     `json:"-"`
	Revenue      *float64 `json:"revenue,omitemplty"`
}

func (user *User) SetPswd(password string) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashPassword
}

func (user *User) CheckPswd(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))

}

type Admin User

type Ambassador User

func (admin *Admin) CalculateRevenue(orders []Order) {

	var revenue float64 = 0
	for _, order := range orders {
		for _, orderItem := range order.OrderItem {
			revenue += orderItem.AdminRevenue
		}
	}

	admin.Revenue = &revenue

}

func (ambassador *Ambassador) CalculateRevenue(orders []Order) {

	var revenue float64 = 0
	for _, order := range orders {
		for _, orderItem := range order.OrderItem {
			revenue += orderItem.AmbassadorRevenue
		}
	}

	ambassador.Revenue = &revenue

}
