package main

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// load movies db file
	db, err := sql.Open("sqlite3", "./movies.db")
	if err != nil {
		panic(err)
	}

	// query to select top 10 shortest runtime movies by given directors and years
	sql := "SELECT title, director, year, runtime FROM movies WHERE director IN (?) AND year IN (?) ORDER BY runtime limit ?"
	directors := []string{"Clint Eastwood", "Christopher Nolan", "Steven Spielberg"}
	years := []int{1981, 1989, 1993, 2000, 2008}
	count := 10

	// expanding '?' placeholders for array args. Not needed for queries without array args
	query, args, err := sqlx.In(sql, directors, years, count)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query(query, args...)
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
