package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// load movies db file
	db, err := sql.Open("sqlite3", "./movies.db")
	if err != nil {
		panic(err)
	}

	// query to select top 10 shortest runtime movies
	rows, err := db.Query("SELECT title, director, year, runtime FROM movies ORDER BY runtime limit 10")
	if err != nil {
		panic(err)
	}

	var title string
	var director string
	var runtime int
	var year int

	fmt.Println("------------------------------------")
	for rows.Next() {
		err = rows.Scan(&title, &director, &year, &runtime)
		if err != nil {
			panic(err)
		}

		// print each movie
		fmt.Printf("Title: %v\n", title)
		fmt.Printf("Director: %v\n", director)
		fmt.Printf("Year: %v\n", year)
		fmt.Printf("Runtime: %v\n", runtime)
		fmt.Println("------------------------------------")
	}

	rows.Close()
	db.Close()
}
