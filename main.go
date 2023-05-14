package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main(){
	var dataSoruce = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", "localhost", 5432, "postgres", "postgres", "dbpg", "disable")
	db, err := sql.Open("postgres",dataSoruce)
	if err != nil {
		panic(err)
	}

	//check db is already connect
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//query
	query := "select * from cover"
	rows, err :=  db.Query(query)
	if err != nil {
		panic(err)
	}

	id := 0
	name := ""
	ok := rows.Next()
	if ok {
		rows.Scan(&id ,&name)
	}

	rows.Close()

	println(id, name)
}