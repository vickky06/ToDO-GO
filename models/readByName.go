package model

import views "../views"

func ReadByName(name string) ([]views.PostRequest, error) {
	rows, err := con.Query("Select * from ToDo where name = ($1)", name)
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
