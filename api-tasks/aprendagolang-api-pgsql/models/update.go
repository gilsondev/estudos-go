package models

import "github.com/gilsondev/aprendagolang-api-pgsql/db"

func Update(id TaskID, task Task) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE tasks SET title = $1, description = $2, done = $3 WHERE id = $4`, task.Title, task.Description, task.Done, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
