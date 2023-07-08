package model

import views "../views"

func ReadAll() ([]views.PostRequest, error) {
	rows, err := con.Query("Select * from ToDo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.ToDo)
		todos = append(todos, data)
		// todos
	}
	return todos, nil

}
