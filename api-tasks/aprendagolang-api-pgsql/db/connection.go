package db

import (
	"database/sql"
	"fmt"

	"github.com/gilsondev/aprendagolang-api-pgsql/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		// panic(err)
		return nil, err
	}

	err = conn.Ping()

	return conn, err
}
