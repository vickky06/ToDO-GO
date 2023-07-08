package model

func DeleteTodo(name string) error {
	deleteQ, err := con.Query("DELETE FROM TODO WHERE name=($1)", name)

	if err != nil {
		print("Err", err)
		return err
	}
	defer deleteQ.Close()
	return nil
}
