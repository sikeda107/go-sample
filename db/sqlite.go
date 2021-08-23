package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Run() {
	fmt.Println("open databse SQLite ...")
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("drop table ...")
	if _, err := db.Exec(`DROP TABLE IF EXISTS employees;`); err != nil {
		panic(err)
	}
	fmt.Println("create table ...")
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS employees (
		id      INT             NOT NULL,
		first_name  VARCHAR(14)     NOT NULL,
		last_name   VARCHAR(16)     NOT NULL,
		dept_id INT,
		PRIMARY KEY (id),
		FOREIGN KEY (dept_id)  REFERENCES departments (id)
	);`); err != nil {
		panic(err)
	}
	fmt.Println("insert values ...")
	if _, err := db.Exec(`INSERT INTO employees (id, first_name, last_name, dept_id) VALUES
	(1,"foo","bar",1),
	(2,"foo","baz",2),
	(3,"hoge","huga",2);`); err != nil {
		panic(err)
	}
	fmt.Println("select values ...")
	rows, err := db.Query(`SELECT * 
	FROM employees
	;`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var id, deptId int
	var firstName, lastName string
	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName, &deptId)
		if err != nil {
			panic(nil)
		}
		fmt.Println(id, firstName, lastName, deptId)
	}
}
