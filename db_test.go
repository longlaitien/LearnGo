package main

import (
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println("MYSQL")
	db, err := sql.Open("musql","root:password@tcp(127.0.0.1.3306/test_db)")
	if err != nil{
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT name FROM book")
	if err != nil{
		panic(err.Error())
	}
	for results.Next(){
		var book Book
		err = results.Scan(&book.Name)
		if err!=nil{
			panic(err.Error())
		}
		fmt.Println(book.Name)

	}
}