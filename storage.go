package main

import (
	"database/sql"
	"log"
	"io/ioutil"
	"strings"
	"time"

	// This is the usual way to include an SQL driver in golang. Actually we are not using
	// any imports from the package explicitly.
	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

// InitializeStorage : init storage
func InitializeStorage(fileName string) {
	log.Println("Opening database", fileName)
	var err error
	database, _ = sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal("Unable to open database", err)
	}
	ExecuteFile("init.sql")
}

// ExecuteFile : execute sql statements
func ExecuteFile(fileName string) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panic("Unable to find init file anywhere: fileName=", fileName)
		return
	}
	log.Println("Executing init script", fileName)

	cmds := strings.Split(string(bytes), ";")
	for _,cmd := range cmds {
		cmd = strings.TrimSpace(cmd)
		if len(cmd) == 0 {
			continue
		}

		statement, err := database.Prepare(cmd)
		if err != nil {
			log.Fatal("Unable to prepare statement: ", cmd)
			return
		}

		_, errExec := statement.Exec()
		if errExec != nil {
			log.Fatal("Unbable to execute statement", errExec)
			return
		}
	}
}

// Save : persist transaction in into a db
func Save(t Transaction) {
	statement, err := database.Prepare(
		"INSERT INTO transactions (UID, Active, StartTimestamp, EndTimestamp) VALUES (?,?,?,?)",
	)

	if err != nil {
		log.Fatal("Invalid insert query", err)
	}

	log.Println("Saving transaction", t)
	statement.Exec(t.UID, t.Active, t.StartTimestamp, t.EndTimestamp)
}

// Load : get data from db
func Load(uid int) []Transaction {
	ts := make([]Transaction, 0, 16)

	rows, err := database.Query( "SELECT UID, Active, StartTimestamp, EndTimestamp FROM transactions WHERE UID = ?", uid)

	if err != nil {
		log.Fatal("Unable to execute query", err)
	}

	for rows.Next() {
		var uid int
		var active bool
		var StartTimestamp time.Time
		var EndTimestamp time.Time

		rows.Scan(&uid, &active, &StartTimestamp, &EndTimestamp)
		t := Transaction{uid, active, StartTimestamp, EndTimestamp}
		ts = append(ts, t)
	}

	return ts
}
