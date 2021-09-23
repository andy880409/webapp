package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUserName = "root"
	dbPassword = "asd880409"
	addr       = "127.0.0.1"
	port       = "3306"
	dbName     = "testdb"
)

func connectDB() (db *sql.DB) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUserName, dbPassword, addr, port, dbName)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("Connect to MySQL Database is Failed:", err)
		return
	}
	return db
}
func createTable(db *sql.DB) {
	sql := `CREATE TABLE IF NOT EXISTS user(
		username VARCHAR(64) PRIMARY KEY,
		password VARCHAR(64)
		);`
	if _, err := db.Exec(sql); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Created Table Successfully.")
}
func addUser(db *sql.DB, un string, pa string) {
	stmt, err := db.Prepare("INSERT INTO user(username,password) VALUES (?,?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(un, pa)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Get Last Insert Id failed,err:", err)
		return
	}
	fmt.Println("Last Insert Id is:", lastInsertId)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Get RowsAffected failed,err:", err)
		return
	}
	fmt.Println("Affected Row:", rowsAffected)
}
func selectUser(db *sql.DB, u User) (*User, error) {
	user := &User{}
	row := db.QueryRow("SELECT username,password FROM user WHERE username=?", u.UserName)
	err := row.Scan(&user.UserName, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}
