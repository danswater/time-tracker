package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func InitializeStorage(fileName string) {
	log.Println("Opening database", fileName)
	var err error
	database, err = sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal("Unable to open database", err)
	}

	executeFile("init.sql")
}

func executeFile(fileName string) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panic("Unable to find init file")
	}

	log.Println("Executing init script", fileName)

	cmds := strings.Split(string(bytes), ";")
	for _, cmd := range cmds {
		cmd = strings.TrimSpace(cmd)
		if len(cmd) == 0 {
			continue
		}
		statement, err := database.Prepare(cmd)
		if err != nil {
			log.Fatal("Unable to execute statement:", cmd)
		}

		statement.Exec()
	}
}

func SaveInterval(interval Interval) int64 {
	statement, err := database.Prepare("INSERT INTO intervals (StartTime, StopTime) VALUES (?,?)")
	if err != nil {
		log.Fatal("Invalid insert query", err)
	}

	startTime := interval.StartTime
	stopTime := interval.StopTime
	log.Println("Saving interval", interval)
	res, err := statement.Exec(startTime, stopTime)
	if err != nil {
		log.Panic("Unable to save interval", err)
	}

	id, _ := res.LastInsertId()
	return id;
}

func SaveStopwatch(stopwatch Stopwatch, intervalId int64) int64 {
	statement, err := database.Prepare("INSERT INTO stopwatches (Color, Id, IntervalId, Name) VALUES (?,?,?,?)")
	if err != nil {
		log.Fatal("Invalid insert query", err)
	}

	color := stopwatch.Color
	sid := stopwatch.Id
	name := stopwatch.Name
	log.Println("Saving stopwatch", stopwatch)
	res, err := statement.Exec(color, sid, intervalId, name)
	if err != nil {
		log.Panic("Unable to save stopwatch", err)
	}

	id, _ := res.LastInsertId()
	return id;
}


