package model

type Question struct {
	Id         int    `json:"id"`
	FormId     int    `json:"formId"`
	Text       string `json:"text"`
	QuestionNo int    `json:"questionNo"`
}
