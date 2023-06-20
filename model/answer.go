package model

import "GolangForm/model/core"

type Answer struct {
	Id         int    `json:"id"`
	SubmitId   int    `json:"submitId"`
	QuestionId int    `json:"questionId"`
	Text       string `json:"text"`
	QuestionNo int    `json:"questionNo"`
}

func CreateAnswer(answer Answer, inputSubmitId int, inputQuestionNo int) error {

	//query := fmt.Sprintf(`insert into answers(submit_id,question_id,text) values (%s,%d,%s);`, inputSubmitId, answer.QuestionId, answer.Text)
	query := `insert into answers(submit_id,question_id,text,question_no) values ($1,$2,$3,$4);`
	_, err := core.DB.Exec(query, inputSubmitId, answer.QuestionId, answer.Text, inputQuestionNo)

	if err != nil {
		return err
	}
	return nil
}

func ListAnswersBySubmitId(inputSubmitId int) ([]Answer, error) {
	var answers []Answer

	query := "select id,submit_id,question_id,text,question_no from answers where submit_id = $1;"

	rows, err := core.DB.Query(query, inputSubmitId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id int
		var submitId int
		var questionId int
		var text string
		var questionNo int

		err = rows.Scan(&id, &submitId, &questionId, &text, &questionNo)

		if err != nil {
			return nil, err
		}

		answer := Answer{
			Id:         id,
			SubmitId:   submitId,
			QuestionId: questionId,
			Text:       text,
			QuestionNo: questionNo,
		}
		answers = append(answers, answer)

	}

	return answers, nil
}
