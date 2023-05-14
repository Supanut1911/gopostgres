package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Cover struct {
	Id int
	Name string
}

var db *sql.DB

func main(){
	var err error

	var dataSoruce = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", "localhost", 5432, "postgres", "postgres", "dbpg", "disable")
	db, err = sql.Open("postgres",dataSoruce)
	if err != nil {
		panic(err)
	}

	covers, err := GetCovers()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", covers)
}

func GetCovers() ([]Cover, error){
	//check db is already connect
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	//query
	query := "select * from cover"
	rows, err :=  db.Query(query)
	if err != nil {
		return nil, err
	}

	covers := []Cover{}
	for rows.Next() {
		cover := Cover{}
		err := rows.Scan(&cover.Id ,&cover.Name)
		if err != nil {
			panic(err)
		}
		covers = append(covers, cover)
	}

	return covers, nil
}