package model

import (
	"fmt"
)

type Form struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func ListAllForms() ([]Form, error) {
	var forms []Form

	query := `select id, title from forms`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var title string

		err = rows.Scan(&id, &title)

		if err != nil {
			return nil, err
		}

		form := Form{
			Id:    id,
			Title: title,
		}
		forms = append(forms, form)
	}
	return forms, nil
}

func ShowForm(inputFormId string) ([]Question, error) {
	var questions []Question

	query := fmt.Sprintf(`select id,form_id,text from questions where form_id = %s`, inputFormId)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var formId int
		var text string
		var questionNo int

		err := rows.Scan(&id, &formId, &text, &questionNo)

		if err != nil {
			return nil, err
		}
		question := Question{
			Id:         id,
			FormId:     formId,
			Text:       text,
			QuestionNo: questionNo,
		}
		questions = append(questions, question)
	}

	return questions, nil
}
