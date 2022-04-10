package database

import (
	"casic/src/models"
)

func InsertDB(user *models.User) error {
	_, err := DB.Query("INSERT INTO usrs (firstname, lastname, email,password, isambassador)VALUES ($1, $2, $3, $4, $5)", user.FirstName, user.LastName, user.Email, user.Password, user.IsAmbassador)
	if err != nil {
		return err
	}
	return nil
}

func UpdateDB(user *models.User) error {
	_, err := DB.Query("UPDATE usrs SET firstname = $2, lastname = $3, email = $4 WHERE id = $1", user.Id, user.FirstName, user.LastName, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func UpdateDBPswd(user *models.User) error {
	_, err := DB.Query("UPDATE usrs SET password = $2 WHERE id = $1", user.Id, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func SelectDB(email string) models.User {
	rows, _ := DB.Query("SELECT id, firstname, lastname, email,password,isambassador FROM usrs WHERE email = $1;", email)
	defer rows.Close()
	//var User models.User
	user := models.User{}
	for rows.Next() {
		rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.IsAmbassador)
	}
	return user
}

func SelectDBId(id uint) models.User {
	rows, _ := DB.Query("SELECT id, firstname, lastname, email,password,isambassador FROM usrs WHERE id = $1;", id)
	defer rows.Close()
	//var User models.User
	user := models.User{}
	//fmt.Println(id)
	for rows.Next() {
		rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.IsAmbassador)
	}

	return user
}

func SelectDBAmbassador(ambassador bool) []models.User {
	rows, _ := DB.Query("SELECT id, firstname, lastname, email,password,isambassador FROM usrs WHERE isambassador = $1;", ambassador)
	defer rows.Close()
	//var User models.User
	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.IsAmbassador)
		users = append(users, user)
	}
	return users
}
