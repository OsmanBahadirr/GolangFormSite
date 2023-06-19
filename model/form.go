package model

type Form struct {
	Id    int
	Title string
}

func ListAllForms() ([]Form, error) {
	var forms []Form

	query := `select id, title from forms`
	rows, err := db.Query(query)
	if err != nil {
		return forms, err
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var title string

		err = rows.Scan(&id, &title)

		if err != nil {
			return forms, err
		}

		form := Form{
			Id:    id,
			Title: title,
		}
		forms = append(forms, form)
	}
	return forms, nil
}
