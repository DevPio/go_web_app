package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var port string = ":3000"

type Book struct {
	id               int
	title            string
	author           string
	publication_year int
	genre            string
	isbn             string
	price            float64
	copies_available int
}

func main() {

	db, _ := sql.Open("mysql", "root:root@tcp(db:3306)/goweb")

	response, err := db.Query("SELECT * FROM books")
	// response, _ := db.Query("SELECT COUNT(IFNULL(code, 1)) FROM goweb;")

	if err != nil {
		log.Fatal(err)
	}
	for response.Next() {
		var book Book

		response.Scan(&book.id, &book.title, &book.author, &book.publication_year, &book.genre, &book.isbn, &book.price, &book.copies_available)

		fmt.Println(book)
	}

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// 	render := render.RenderTemplate("./views/index.tpml")

	// 	render.Execute(w, nil)
	// })

	// log.Fatal(http.ListenAndServe(port, nil))

}
