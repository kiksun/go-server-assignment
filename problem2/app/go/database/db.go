package database

import (
	"database/sql"
	"log"
	"problem1/configs"
)

func ConnectDB() (*sql.DB, error) {
	conf := configs.Get()

	db, err := sql.Open(conf.DB.Driver, conf.DB.DataSource)
	if err != nil {
		panic(err)
	}

	return db, nil
}

func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}

}
