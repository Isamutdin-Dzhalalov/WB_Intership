package DB

import (
	"database/sql"
	"fmt"
	"main/setting"
)

var db *sql.DB

func ConnectWithDb() (*sql.DB, error) {
	var err error
	
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", setting.Config.PgUser, setting.Config.PgPassword, setting.Config.PgNameDB)
	db, err = sql.Open("postgres", connStr)

/*
	if err != nil {
		return db, err
	}
*/

	return db, err
}
