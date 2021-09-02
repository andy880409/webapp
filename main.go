package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//index page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, "Login Page")
}

//authentication page
func authHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	fmt.Printf("帳號:%s\n密碼:%s\n", username, password)
	http.Redirect(w, r, "/welcome", http.StatusSeeOther)
}

//sign up page
func signUpHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("sign-up.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, t)
}

//welcome page
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("welcome.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, t)
}
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/sign-up", signUpHandler)
	http.HandleFunc("/welcome", welcomeHandler)
	http.ListenAndServe(":8080", nil)
}
