package controller

import (
	"GolangForm/controller/site"
	"GolangForm/model/core"
	"fmt"
	"log"
	"net/http"
	"os"
)

var mux *http.ServeMux

const (
	roleAdmin  = "admin"
	roleMod    = "mod"
	rolePublic = "public"
)

func initHandlers() {
	mux.HandleFunc("/getjwt", core.GetJWT)
	mux.HandleFunc("/api/form/listall", site.ListAllFormsHandler)
	mux.HandleFunc("/api/form/show", site.ShowFormHandler)
	mux.HandleFunc("/api/user/register", site.CreateAcountHandler)
	mux.HandleFunc("/api/user/login", site.LoginHandler)
	mux.Handle("/api/answer/create", core.ValidateJwtOnHandler([]string{os.Getenv("ADMIN")}, site.CreateAnswersHandler))
	mux.Handle("/api/answer/listbysubmitid", core.ValidateJwtOnHandler([]string{os.Getenv("ADMIN")}, site.ListAnswersBySubmitIdHandler))

}

func Start() {
	mux = http.NewServeMux()

	initHandlers()
	fmt.Printf("Mux initialized and listening on server :4000\n")
	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Fatal(err)
	}
}
