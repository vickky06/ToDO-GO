package model

func CreateTodo(name string, todo string) error {
	insertQ, err := con.Query(
		"INSERT INTO TODO VALUES($1, $2)", name, todo,
	)
	if err != nil {
		return err
	}
	defer insertQ.Close()
	return nil
}
