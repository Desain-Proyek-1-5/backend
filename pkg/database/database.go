package database

import (
	"database/sql"
	"log"
)

type DBInstance struct {
	db *sql.DB
}

type RetrievedData struct {
	Data *sql.Rows
}

func NewDatabase(databaseURL string) (*DBInstance, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		panic(err)
	}
	return &DBInstance{db}, nil
}

func (a *DBInstance) AddData(Query string) error {
	_, err := a.db.Query(Query)
	return err

}

func (a *DBInstance) UpdateData(Query string) error {
	a.db.Query("SET SQL_SAFE_UPDATES = 0;")
	_, err := a.db.Query(Query)
	a.db.Query("SET SQL_SAFE_UPDATES = 1;")
	return err

}

func (a *DBInstance) DeleteData(Query string) error {
	a.db.Query("SET FOREIGN_KEY_CHECKS=0;")
	_, err := a.db.Query(Query)
	a.db.Query("SET FOREIGN_KEY_CHECKS=1;")
	return err
}

func (a *DBInstance) RetrieveData(Query string) (*RetrievedData, error) {
	rows, err := a.db.Query(Query)
	if err != nil {
		return nil, err
	}
	return &RetrievedData{rows}, nil
}

func (a *DBInstance) CheckIfExists(Query string) bool {
	log.Println("Check for duplicates")
	var exists bool
	err := a.db.QueryRow(Query).Scan(&exists)
	if err != nil {
		log.Println("error : ", err.Error())
		return true
	}
	log.Printf("Duplicate returns : %t", exists)
	return exists

}
