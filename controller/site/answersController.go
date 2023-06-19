package site

import (
	"GolangForm/model"
	"encoding/json"
	"net/http"
	"strconv"
)

func CreateAnswersHandler(w http.ResponseWriter, r *http.Request) {

	inputSubmitId := r.URL.Query().Get("submitId")
	intInputSubmitId, _ := strconv.Atoi(inputSubmitId)

	decoder := json.NewDecoder(r.Body)
	var answers []model.Answer
	err := decoder.Decode(&answers)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {

		for i := 0; i < len(answers); i++ {
			err = model.CreateAnswer(answers[i], intInputSubmitId, i+1)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
		}
	}
}

func ListAnswersBySubmitIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	inputSubmitId, _ := strconv.Atoi(r.URL.Query().Get("submitId"))

	answers, err := model.ListAnswersBySubmitId(inputSubmitId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(answers)
	}
}
