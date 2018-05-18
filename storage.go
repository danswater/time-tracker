package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"

	// This is the usual way to include an SQL driver in golang. Actually we are not using
	// any imports from the package explictly.
	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

// InitializeStorage : init storage
func InitializeStorage() {
	database, _ = sql.Open("sqlite3", "./demo.db")
	ExecuteFile("init.sql")
}

// ExecuteFile : execute sql statements
func ExecuteFile(fileName string) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("No initialization file found.")
		return
	}

	cmds := strings.Split(string(bytes), ";")
	for _, cmd := range cmds {
		cmd = strings.TrimSpace(cmd)
		if len(cmd) == 0 {
			continue
		}

		statement, err := database.Prepare(cmd)
		if err != nil {
			fmt.Println("Unable to execute statement: ", cmd)
			return
		}
		statement.Exec()
	}
}

// Save : persist transaction in into a db
func Save(t Transaction) {
	fmt.Println("Saving: ", t)
}

// Load : get data from db
func Load(year, month int) Transaction {
	fmt.Println("Loading for", year, month)
	return Transaction{}
}
