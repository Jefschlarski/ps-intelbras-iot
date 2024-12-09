package db

import (
	"database/sql"
	"fmt"

	"github.com/Jefschlarski/ps-intelbras-iot/monitor_api/pkg/config"
	_ "github.com/lib/pq"
)

func ConnectDB(dbConfig config.DBConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Pass, dbConfig.Database)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + dbConfig.Database)

	return db, nil
}
