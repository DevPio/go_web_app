package main

import (
	"database/sql"

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

var port string = ":3000"

func Init() {

	response, err := db.Query(`CREATE TABLE IF NOT EXISTS books (
		Id INT AUTO_INCREMENT PRIMARY KEY,
		Title VARCHAR(100) NOT NULL,
		Body VARCHAR(255) NOT NULL
	);`)

	checkErro(err)
	defer response.Close()

}

func main() {

	Init()
	checkErro(err)
	http.ListenAndServe(port, router())

}

func router() *mux.Router {
	router := mux.NewRouter()

	router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))

	router.HandleFunc("/", start)
	router.HandleFunc("/createPost", renderPost)
	router.HandleFunc("/newPost", createPost).Methods("POST")
	router.HandleFunc("/editPost", editPost)
	router.HandleFunc("/editPostDb", editPostDb).Methods("POST")

	return router
}

func editPost(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")

	result := db.QueryRow("SELECT * FROM books WHERE Id =?", id)

	var post Post

	result.Scan(&post.Id, &post.Title, &post.Body)

	render := template.Must(template.ParseFiles("./views/layout.html", "./views/post.html"))

	render.Execute(w, post)
}

func editPostDb(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	body := r.FormValue("title")

	_, error := db.Query("UPDATE books SET Title = '" + title + "', Body = '" + body + "' WHERE Id =" + id + "")

	checkErro(error)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
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

	render := template.Must(template.ParseFiles("./views/layout.html", "./views/list.html"))

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

	render := template.Must(template.ParseFiles("./views/layout.html", "./views/post.html"))

	render.Execute(w, nil)
}

func createPost(w http.ResponseWriter, r *http.Request) {

	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	addPost(title, body)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}
