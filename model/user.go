package model

import (
	"GolangForm/model/core"
	"fmt"
	"log"
)

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

func CreateAcount(user User) bool {
	query := `insert into users(user_name,password,role,token) values($1,$2,$3,$4)`

	userToken, err := core.CreateJWT(user.Role)
	if err != nil {
		fmt.Print(err.Error())
		return false
	}

	_, err = core.DB.Exec(query, user.UserName, user.Password, user.Role, userToken)

	if err != nil {
		log.Fatal(err)
	}

	return true
}

func Login(inputUser User) (bool, User) {

	query := `select user_name,password,role from users where user_name = $1`

	var user User
	rows, err := core.DB.Query(query, inputUser.UserName)

	if err != nil {
		fmt.Println()
		return false, user
	}

	for rows.Next() {
		var id int
		var userName string
		var password string
		var role string
		var token string

		err := rows.Scan(&userName, &password, &role)
		if err != nil {
			fmt.Println(err)
			return false, user
		}
		user = User{
			Id:       id,
			UserName: userName,
			Password: password,
			Role:     role,
			Token:    token,
		}
		if inputUser.Password == user.Password && inputUser.UserName == user.UserName {
			user.Token, err = core.CreateJWT(user.Role)
			if err != nil {
				fmt.Println(err)
				return false, user

			}
			return true, user
		}

	}
	return false, user
}

//token sistemini entegre et
