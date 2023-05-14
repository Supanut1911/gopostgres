package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main(){
	var dataSoruce = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", "localhost", 5432, "postgres", "postgres", "godb", "disable")
	sql.Open("postgres",dataSoruce)
}