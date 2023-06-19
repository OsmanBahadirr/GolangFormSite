package controller

import (
	"GolangForm/controller/site"
	"fmt"
	"log"
	"net/http"
)

var mux *http.ServeMux

func initHandlers() {
	mux.HandleFunc("/api/form/listall", site.ListAllFormsHandler)
	mux.HandleFunc("/api/form/show", site.ShowFormHandler)
	mux.HandleFunc("/api/answer/create", site.CreateAnswersHandler)
	mux.HandleFunc("/api/answer/listbysubmitid", site.ListAnswersBySubmitIdHandler)
}

func Start() {
	mux = http.NewServeMux()

	initHandlers()
	fmt.Printf("Mux initialized and listening on server :8080\n")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
