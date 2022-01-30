package main

import (
	"fmt"
	"net/http"
	"html/template"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

)

type Article struct {
	Id uint16
	Title, Anons, FullText string
}



func index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("template/index.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}


	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/go-web-site")
	if err != nil{
		panic(err)
	}

	defer db.Close()

	res, err := db.Query("SELECT * FROM `articles`")
	if err != nil{
		panic(err)
	}

	var posts = []Article{}

	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)
		if err != nil{
			panic(err)
		}
		posts = append(posts, post)

	}


	temp.ExecuteTemplate(w, "index", posts)
}

func create(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("template/create.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	temp.ExecuteTemplate(w, "create", nil)
}

func save_article(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	if title == "" || anons == "" || full_text == "" {
		fmt.Fprintf(w, "Не все поля заполнены")
	} else {
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/go-web-site")
		if err != nil{
			panic(err)
		}

		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`, `anons`, `full_text`) " +
			"VALUES('%s', '%s', '%s')", title, anons, full_text))
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}


func handleFunc() {
	http.HandleFunc("/", index)
	http.HandleFunc("/create/", create)
	http.HandleFunc("/save_article/", save_article)
	http.ListenAndServe(":8090", nil)
}

func main() {
	handleFunc()
}
