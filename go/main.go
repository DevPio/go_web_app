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

func main() {

	db, _ := sql.Open("mysql", "root:root@tcp(db:3306)/goweb")

	response, err := db.Query(`CREATE TABLE IF NOT EXISTS books (
		Id INT AUTO_INCREMENT PRIMARY KEY,
		Title VARCHAR(100) NOT NULL,
		Body VARCHAR(255) NOT NULL
	);`)

	fmt.Println(err)
	defer response.Close()
	// response, _ := db.Query("SELECT COUNT(IFNULL(code, 1)) FROM goweb;")

	fmt.Println(db.Stats().InUse)

	router := mux.NewRouter()

	router.HandleFunc("/", start)
	router.HandleFunc("/createPost", renderPost)
	router.HandleFunc("/newPost", createPost).Methods("POST")

	http.ListenAndServe(":3000", router)

}

func start(w http.ResponseWriter, r *http.Request) {
	newPost := Post{
		Id:    1,
		Title: "Nem post",
		Body:  "",
	}

	title := r.FormValue("title")

	if title != "" {
		newPost.Title = title
	}

	render := template.Must(template.ParseFiles("./views/index.html"))

	render.Execute(w, newPost)

}

func renderPost(w http.ResponseWriter, r *http.Request) {

	render := template.Must(template.ParseFiles("./views/post.html"))

	render.Execute(w, nil)
}

func createPost(w http.ResponseWriter, r *http.Request) {

	title := r.PostFormValue("title")
	body := r.PostFormValue("body")
	fmt.Println(title)
	fmt.Println(body)

	http.Redirect(w, r, "http://localhost:3000/", http.StatusAccepted)
}
