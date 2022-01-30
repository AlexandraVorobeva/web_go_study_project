package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type User struct {
	Name  string
	Age   uint16
	Money     int16
	AvgGrades float64
	Happiness float64
	Hobbies []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is %s. He is %d years old.", u.Name, u.Age)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}


func homePage(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, 50, 4.8, 0.8, []string{"football", "dance"}}
	//fmt.Fprintf(w, "<b>Main Text</b>")
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)
	http.ListenAndServe(":8080", nil)
}

func main() {

	handleRequest()
}
