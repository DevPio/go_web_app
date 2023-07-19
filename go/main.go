package main

import (
	"database/sql"
	"fmt"

	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

var db, err = sql.Open("mysql", "root:root@tcp(db:3306)/goweb")

func Init() {

	response, err := db.Query(`CREATE TABLE IF NOT EXISTS books (
		Id INT AUTO_INCREMENT PRIMARY KEY,
		Title VARCHAR(100) NOT NULL,
		Body VARCHAR(255) NOT NULL
	);`)

	fmt.Println(err)
	defer response.Close()

}

func main() {

	checkErro(err)
	Init()

	db.Ping()

	router := mux.NewRouter()

	router.HandleFunc("/", start)
	router.HandleFunc("/createPost", renderPost)
	router.HandleFunc("/newPost", createPost).Methods("POST")
	router.HandleFunc("/editPost", editPost)

	http.ListenAndServe(":3000", router)

}

func editPost(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")

	result, error := db.Query("SELECT * FROM books WHERE Id = " + id + ";")

	checkErro(error)

	var post Post

	for result.Next() {
		result.Scan(&post.Id, &post.Title, &post.Body)
	}

	render := template.Must(template.ParseFiles("./views/post.html"))

	render.Execute(w, post)
}

func getAllItem() []Post {

	result, err := db.Query("SELECT * FROM books;")

	checkErro(err)

	posts := make([]Post, 0)

	for result.Next() {
		var post Post

		result.Scan(&post.Id, &post.Title, &post.Body)

		posts = append(posts, post)
	}

	return posts
}

func start(w http.ResponseWriter, r *http.Request) {

	render := template.Must(template.ParseFiles("./views/index.html"))

	render.Execute(w, getAllItem())

}

func addPost(title string, body string) *sql.Stmt {
	smt, error := db.Prepare("INSERT INTO books(Title, Body) values(?,?)")

	checkErro(error)

	smt.Exec(title, body)

	return smt
}

func checkErro(err error) {
	if err != nil {
		panic(err)
	}
}

func renderPost(w http.ResponseWriter, r *http.Request) {

	render := template.Must(template.ParseFiles("./views/post.html"))

	render.Execute(w, nil)
}

func createPost(w http.ResponseWriter, r *http.Request) {

	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	addPost(title, body)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}
