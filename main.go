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

	coverNew := Cover{
		Name: "cover-UFO",
	}

	err = AddCover(coverNew)
	if err != nil {
		panic(err)
	}

	covers, err := GetCovers()
	if err != nil {
		panic(err)
	} 
	// _ = covers
	fmt.Printf("%#v \n", covers)

	// cover,err  := GetCover(1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(*cover)
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

func GetCover(id int) (*Cover, error) {
	//check db is already connect
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	//query
	query := "select id, name from cover where id = $1"
	row := db.QueryRow(query, id)

	cover := Cover{}
	err = row.Scan(&cover.Id, &cover.Name)
	if err != nil {
		return nil, err
	}
	return &cover, nil
}

func AddCover(cover Cover) error {
	//check db is already connect
		err := db.Ping()
		if err != nil {
			return err
		}
	
	query := "INSERT INTO cover (name) values ($1)"
	result, err := db.Exec(query, cover.Name)
	if err != nil {
		return err
	}
	_ = result
	return nil
}