package models

import "github.com/gilsondev/aprendagolang-api-pgsql/db"

func Delete(id TaskID) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, nil
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM tasks WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
