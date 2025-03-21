package models

import "github.com/gilsondev/aprendagolang-api-pgsql/db"

func Get(id TaskID) (task Task, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT id, title, description FROM tasks WHERE id = $1`, id)
	err = row.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
	)

	return
}
