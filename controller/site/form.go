package site

import (
	"GolangForm/model"
	"encoding/json"
	"net/http"
)

func ListAllFormsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	forms, err := model.ListAllForms()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(forms)
	}
}
