package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	UserName string
	Password string
}

var tpl *template.Template
var dbUser = map[string]User{}      //key:username | value:User Data
var dbSession = map[string]string{} //key:session id | value:username
var a = 1

func init() {
	tpl = template.Must(template.ParseGlob("view/*"))
}

//index page
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		un := r.FormValue("username") //key就是html的name
		p := r.FormValue("password")
		fmt.Printf("帳號:%s\n密碼:%s\n", un, p)
		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.html", "Login Page")
}

//sign up page
func signUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")
		dbUser[un] = User{UserName: un, Password: p}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "sign-up.html", nil)
}

//welcome page
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "welcome.html", nil)
}
func main() {
	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/sign-up", signUpHandler)
	http.HandleFunc("/welcome", welcomeHandler)
	http.ListenAndServe(":8080", nil)
}
