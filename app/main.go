package main

import (
	"fmt"
	"net/http"
	"html/template"

)

func index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("template/index.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	temp.ExecuteTemplate(w, "index", nil)

}


func handleFunc() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8090", nil)
}

func main() {
	handleFunc()
}
