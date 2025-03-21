package models

import "github.com/gilsondev/aprendagolang-api-pgsql/db"

// Insert a task into the database.
//
// task Task - the task to be inserted.
// (id TaskID, err error) - returns the task ID and any error that occurred.
func Insert(task Task) (id TaskID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO tasks (title, description, done) VALUES ($1, $2, $3) RETURNING id`
	err = conn.QueryRow(sql, task.Title, task.Description, task.Done).Scan(&id)

	return
}
