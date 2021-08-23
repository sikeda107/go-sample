package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sikeda107/go-sample/db"
	"github.com/sikeda107/go-sample/example"
)

func hello() string {
	return "Hello, world."
}

func greet(n int) string {
	result := ""
	switch n {
	case 1:
		result = "Good Morning"
	case 2:
		result = "Hello"
	case 3:
		result = "Good Evening"
	default:
		result = "Hello, world."
	}
	return result
}

func main() {
	fmt.Println(hello())
	fmt.Println(greet(3))
	c := example.Calculator{A: 10, B: 20}
	fmt.Println(c.Add1())
	fmt.Println(c.Add2(30))
	db.Run()
}
