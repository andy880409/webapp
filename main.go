package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	UserName string
	Password string
}

var tpl *template.Template
var dbUser = map[string]User{}      //key:username | value:User Data
var dbSession = map[string]string{} //key:session id | value:username

func init() {
	tpl = template.Must(template.ParseGlob("view/*"))
}

//檢查用戶是否存在
func checkUserIsExist(u string) bool {
	_, isExist := dbUser[u]
	return isExist
}

//檢查密碼是否一致
func checkPassword(p1 string, p2 string) error {
	if p1 == p2 {
		return nil
	}
	return errors.New("密碼不對")
}

//驗證身分
func auth(username string, password string) error {
	if checkUserIsExist(username) {
		return checkPassword(password, dbUser[username].Password)
	}
	return errors.New("此用戶不存在")
}

//設sessionID在cookie中，並把session完成
func setSession(w http.ResponseWriter, un string) {
	sID, _ := uuid.NewV4()
	c := &http.Cookie{
		Name:  "sessionID",
		Value: sID.String(),
	}
	http.SetCookie(w, c)
	dbSession[c.Value] = un
}
func clearSession(w http.ResponseWriter) {
	c := &http.Cookie{
		Name:   "sessionID",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
}

//index page
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var data string
	if r.Method == http.MethodPost {
		un := r.FormValue("username") //key就是html的name
		p := r.FormValue("password")
		fmt.Printf("帳號:%s\n密碼:%s\n", un, p)
		err := auth(un, p)
		if err == nil {
			setSession(w, un)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		data = err.Error()
	}
	tpl.ExecuteTemplate(w, "login.html", data)
}

//sign up page
func signupHandler(w http.ResponseWriter, r *http.Request) {
	var data string
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")
		if _, ok := dbUser[un]; !ok {
			dbUser[un] = User{UserName: un, Password: p}
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		data = fmt.Sprintf("帳號:%s 已經存在了!", un)
	}
	tpl.ExecuteTemplate(w, "signup.html", data)
}

//log out
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

//index page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	//檢查cookie內是否有sessionID
	c, err := r.Cookie("sessionID")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	un := dbSession[c.Value]
	u := dbUser[un]
	tpl.ExecuteTemplate(w, "index.html", u)
}
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.ListenAndServe(":8080", nil)
}
