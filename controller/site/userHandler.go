package site

import (
	"GolangForm/model"
	"encoding/json"
	"net/http"
)

func CreateAcountHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user model.User

	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		ok := model.CreateAcount(user)

		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var userToLogin model.User

	err := decoder.Decode(&userToLogin)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		ok, user := model.Login(userToLogin)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.Header().Set("Token", user.Token)

	}
}

//test et
