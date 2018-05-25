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

func CreateInterval(interval Interval) Interval {
	statement, err := database.Prepare("INSERT INTO intervals (StopwatchId, StartTime, StopTime) VALUES (?,?,?)")
	if err != nil {
		log.Fatal("Invalid insert query", err)
	}

	stopwatchId := interval.StopwatchId
	startTime := interval.StartTime
	stopTime := interval.StopTime
	log.Println("Saving interval", interval)
	res, err := statement.Exec(stopwatchId, startTime, stopTime)
	if err != nil {
		log.Panic("Unable to save interval", err)
	}

	id, _ := res.LastInsertId()
	interval.Id = id
	return interval
}

func UpdateInterval(interval Interval) Interval {
	statement, err := database.Prepare("UPDATE intervals SET StopTime = ? WHERE Id = ?")
	if err != nil {
		log.Fatal("Invalid updating query", err)
	}

	stopTime := interval.StopTime
	log.Println("Updating interval data", interval)
	_, errExec := statement.Exec(stopTime, interval.Id)
	if errExec != nil {
		log.Panic("Unable to update interval data", errExec)
	}

	return interval;
}


func CreateStopwatch(stopwatch Stopwatch) Stopwatch {
	statement, err := database.Prepare("INSERT INTO stopwatches (Color, Id, Name) VALUES (?,?,?)")
	if err != nil {
		log.Fatal("Invalid insert stopwatch query", err)
	}

	color := stopwatch.Color
	sid := stopwatch.Id
	name := stopwatch.Name
	log.Println("Saving stopwatch", stopwatch)
	_, errExec := statement.Exec(color, sid, name)
	if errExec != nil {
		log.Panic("Unable to save stopwatch ", errExec)
	}

	return stopwatch
}

func CreatePoolData(poolData PoolData) PoolData {
	statement, err := database.Prepare("INSERT INTO pool_datas (CreationDate, LastModDate) VALUES (?,?)")
	if err != nil {
		log.Fatal("Invalid insert pool data query", err)
	}

	creationDate := poolData.CreationDate
	lastModDate := poolData.LastModDate
	log.Println("Creating pool data", poolData)
	res, errExec := statement.Exec(creationDate, lastModDate)
	if errExec != nil {
		log.Panic("Unable to create pool data", errExec)
	}

	id, _ := res.LastInsertId()
	poolData.Id = id
	return poolData;
}

func UpdatePoolData(poolData PoolData) PoolData {
	statement, err := database.Prepare("UPDATE pool_datas SET CreationDate = ?, LastModDate = ?, StopwatchId = ? WHERE Id = ?")
	if err != nil {
		log.Fatal("Invalid updating query", err)
	}

	log.Println("Updating pool data", poolData)
	_, errExec := statement.Exec(poolData.CreationDate, poolData.LastModDate, poolData.StopwatchId, poolData.Id)
	if errExec != nil {
		log.Panic("Unable to update pool data", errExec)
	}

	return poolData;
}

func CreatePool(pool Pool) Pool {
	statement, err := database.Prepare("INSERT INTO pools (EventName, IsReadOnly, Message, PoolDataId, PoolKey, PoolKeyReadOnly) VALUES (?,?,?,?,?,?)")
	if err != nil {
		log.Fatal("Invalid insert query", err)
	}

	eventName := pool.EventName
	isReadOnly := pool.IsReadOnly
	message := pool.Message
	poolKey := pool.PoolKey
	PoolKeyReadOnly := pool.PoolKeyReadOnly
	log.Println("Creating pool",pool)
	_, errExec := statement.Exec(eventName, isReadOnly, message, pool.PoolData.Id, poolKey, PoolKeyReadOnly)
	if errExec != nil {
		log.Fatal("Invalid insert query", errExec)
	}

	return pool
}

func LoadInterval(id int64) Interval {
	var StartTime int64
	var StopTime int64
	var StopwatchId int

	err := database.QueryRow("SELECT StartTime, StopTime, StopwatchId FROM intervals WHERE Id", id).Scan(&StartTime, &StopTime, &StopwatchId)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No pool data row found.")
	case err != nil:
		log.Fatal(err)
	}

	i := NewIntervalWithId(id, StopwatchId, StartTime, StopTime)

	return i
}

func LoadIntervalsByStopwatchId(id int) []Interval {
	is := make([]Interval, 0, 16)
	log.Println("TEST", id)
	rows, err := database.Query("SELECT Id, StartTime, StopTime FROM intervals WHERE StopwatchId", id)
	if err != nil {
		log.Fatal("Unable to execute interval query ", err)
	}

	for rows.Next() {
		var Id int64
		var StartTime int64
		var StopTime int64

		rows.Scan(&Id, &StartTime, &StopTime)
		i := NewIntervalWithId(Id, id, StartTime, StopTime)
		is = append(is, i)
	}

	return is
}

func LoadStopwatch(id int) Stopwatch {
	var Color string
	var Id int
	var Name string

	err := database.QueryRow("SELECT Color, Id, Name FROM stopwatches WHERE Id = ?", id).Scan(&Color, &Id, &Name)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No pool data row found.")
	case err != nil:
		log.Fatal(err)
	}

	intervals := LoadIntervalsByStopwatchId(id)
	sps := NewStopwatchWithIntervals(Id, Color, Name, intervals)

	return sps
}

func LoadPoolData(id int64) PoolData {
	var Id int64
	var CreationDate int64
	var LastModDate int64
	var StopwatchId int

	err := database.QueryRow("SELECT Id, CreationDate, LastModDate, StopwatchId FROM pool_datas WHERE Id = ?", id).Scan(&Id, &CreationDate, &LastModDate, &StopwatchId)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No pool data row found.")
	case err != nil:
		log.Fatal(err)
	}

	var pd PoolData
	if StopwatchId != 0 {
		stopwatch := LoadStopwatch(StopwatchId)
		pd = NewPoolWithStopwatches(Id, CreationDate, LastModDate, StopwatchId, stopwatch)
	} else {
		pd = NewPoolWithId(Id, CreationDate, LastModDate)
	}

	return pd
}

func LoadPoolDataByStopwatchId(id int64) PoolData {
	var Id int64
	var CreationDate int64
	var LastModDate int64
	var StopwatchId int

	err := database.QueryRow("SELECT Id, CreationDate, LastModDate, StopwatchId FROM pool_datas WHERE Id = ?", id).Scan(&Id, &CreationDate, &LastModDate, &StopwatchId)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No pool data row found.")
	case err != nil:
		log.Fatal(err)
	}

	var pd PoolData
	if StopwatchId != 0 {
		stopwatches := LoadStopwatch(StopwatchId)
		pd = NewPoolWithStopwatches(Id, CreationDate, LastModDate, StopwatchId, stopwatches)
	} else {
		pd = NewPoolWithId(Id, CreationDate, LastModDate)
	}

	return pd
}

func LoadPool(pool Pool) Pool {
	var Id int64
	var EventName string
	var IsReadOnly bool
	var Message string
	var PoolDataId int64
	var PoolKey string
	var PoolKeyReadOnly string

	err := database.QueryRow("SELECT Id, EventName, IsReadOnly, Message, PoolDataId, PoolKey, PoolKeyReadOnly FROM pools WHERE poolKey = ?", pool.PoolKey).Scan(&Id, &EventName, &IsReadOnly, &Message, &PoolDataId, &PoolKey, &PoolKeyReadOnly)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No pool row found.")
	case err != nil:
		log.Fatal(err)
	}

	pd := LoadPoolData(PoolDataId)
	p := NewPool(EventName, IsReadOnly, Message, pd, PoolKey, PoolKeyReadOnly)

	return p
}


