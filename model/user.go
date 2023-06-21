package model

import (
	"GolangForm/model/core"
	"fmt"
	"log"
)

type User struct {
	Id       int
	UserName string
	Password string
	Role     string
	Token    string
}

func CreateAcount(user User) {
	query := `insert into users(user_name,password,role,token) values($1,$2,$3,$4)`

	userToken, err := core.CreateJWT(user.Role)
	if err != nil {
		log.Fatal(err)
	}
	core.DB.Exec(query, user.UserName, user.Password, user.Role, userToken)
}

func Login(inputUser User) (bool, error) {
	query := fmt.Sprintf(`select user_name,password,role from users where id = %v`, inputUser.Id)

	rows, err := core.DB.Query(query)

	if err != nil {
		return false, err
	}

	for rows.Next() {
		var id int
		var userName string
		var password string
		var role string
		var token string

		err := rows.Scan(&id, &userName, &password, &role, &token)
		if err != nil {
			return false, err
		}
		user := User{
			Id:       id,
			UserName: userName,
			Password: password,
			Role:     role,
			Token:    token,
		}
		if inputUser.Password == user.Password && inputUser.UserName == user.UserName {
			return true, nil
		}

	}
	return false, nil
}

//token sistemini entegre et
