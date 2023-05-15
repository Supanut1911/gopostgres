package main

import (
	"database/sql"
	"errors"
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

	_ = coverNew

	// err = AddCover(coverNew)
	// if err != nil {
		// panic(err)
	// }

	//UPFATE
	// err = UpdateCover(9, "cover-ZEO")

	//DELETE
	err = DeleteCover(9)
	if err != nil {
		panic(err)
	}

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

	affected, err := result.RowsAffected()
	if err != nil {
		return err	
	}
	if affected <= 0 {
		return errors.New("can not insert")
	}

	return nil
}

func UpdateCover(id int, name string) error {
	//check db is already connect
	err := db.Ping()
	if err != nil {
		return err
	}

	//find cover
	cover, err := GetCover(id)
	if err != nil {
		return err
	}


	//query update cover
	query := "UPDATE cover set name = $1 where id = $2"
	result, err := db.Exec(query, name, cover.Id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("can not insert")
	}
	return nil
}

func DeleteCover(id int) error {
	//check db is already connect
		err := db.Ping()
		if err != nil {
			return err
	}
	
	//query del cover
	query := "DELETE FROM cover WHERE id = $1"
	result, err := db.Exec(query, id)
	
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("can not delete")
	}

	return nil
	


		
}