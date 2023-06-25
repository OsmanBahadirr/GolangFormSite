package site

import (
	"GolangForm/model"
	"encoding/json"
	"net/http"
)

func UpdateQuestionHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var inputQuestion model.Question

	err := decoder.Decode(&inputQuestion)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	model.UpdateQuestion(inputQuestion)
}
