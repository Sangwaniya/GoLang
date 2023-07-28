package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	ID         int
	Name       string
	Department string
	ProjetID   int
}

func main() {
	// db, err := sql.Open("mysql", "mysql://root@/employees")
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/employees")

	if err != nil {
		log.Fatalln("Unable to connect with server", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS employees (id INTEGER PRIMARY KEY, name text, department text, project_id integer);")
	if err != nil {
		log.Fatalln("Unable to create table: ", err)
		return
	}

}
