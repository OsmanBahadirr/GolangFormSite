package model

import (
	"GolangForm/model/core"
	"fmt"
)

type Question struct {
	Id         int    `json:"id"`
	FormId     int    `json:"formId"`
	Text       string `json:"text"`
	QuestionNo int    `json:"questionNo"`
}

func UpdateQuestion(inputQuestion Question) {
	query := `update questions set form_id = $1, text = $2, question_id = $3 where id = $4`

	_, err := core.DB.Exec(query, inputQuestion.FormId, inputQuestion.Text, inputQuestion.QuestionNo, inputQuestion.Id)
	if err != nil {
		fmt.Println(err.Error())
	}
}
