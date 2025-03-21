package models

import "github.com/gilsondev/aprendagolang-api-pgsql/db"

func GetAll() (tasks []Task, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT id, title, description FROM tasks`)
	if err != nil {
		return
	}

	for rows.Next() {
		var task Task

		err = rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
		)
		if err != nil {
			continue
		}

		tasks = append(tasks, task)
	}

	return
}
