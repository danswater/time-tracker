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
	_, errExec := statement.Exec(color, sid, intervalId, name)
	if errExec != nil {
		log.Panic("Unable to save stopwatch ", errExec)
	}

	return int64(sid);
}

func CreatePoolData(poolData PoolData, stopwatchId *int64) int64 {
	statement, err := database.Prepare("INSERT INTO pool_datas (CreationDate, LastModDate, StopwatchId) VALUES (?,?,?)")
	if err != nil {
		log.Fatal("Invalid insert query", err)
	}

	creationDate := poolData.CreationDate
	lastModDate := poolData.LastModDate
	log.Println("Creating pool data", poolData)
	res, errExec := statement.Exec(creationDate, lastModDate, stopwatchId)
	if errExec != nil {
		log.Panic("Unable to create pool data", errExec)
	}

	id, _ := res.LastInsertId()
	return id;
}

func UpdatePoolData(pool Pool, poolData PoolData, stopwatchId *int64) int64 {
	rows, errSelect := database.Query("SELECT PoolDataId FROM pools WHERE poolKey = ?", pool.PoolKey)
	if errSelect != nil {
		log.Fatal("Unable to execute pool query", errSelect)
	}

	var poolDataId int64
	for rows.Next() {
		var PoolDataId int64
		rows.Scan(&PoolDataId)
		poolDataId = PoolDataId
	}

	statement, err := database.Prepare("UPDATE pool_datas SET CreationDate = ?, LastModDate = ?, StopwatchId = ? WHERE Id = ?")
	if err != nil {
		log.Fatal("Invalid updating query", err)
	}

	creationDate := poolData.CreationDate
	lastModDate := poolData.LastModDate
	log.Println("Updating pool data", poolData)
	_, errExec := statement.Exec(creationDate, lastModDate, stopwatchId, poolDataId)
	if errExec != nil {
		log.Panic("Unable to update pool data", errExec)
	}

	return poolDataId;
}

func UpdatePool(pool Pool, poolDataId int64) {
	statement, err := database.Prepare("UPDATE pools SET EventName = ?, IsReadOnly = ?, Message = ?, PoolDataId = ? WHERE PoolKey = ?")
	if err != nil {
		log.Fatal("Invalid update query", err)
	}

	eventName := pool.EventName
	isReadOnly := pool.IsReadOnly
	message := pool.Message
	poolKey := pool.PoolKey
	log.Println("Updating pool",pool)
	_, errExec := statement.Exec(eventName, isReadOnly, message, poolDataId, poolKey)
	if errExec != nil {
		log.Fatal("Invalid update query", errExec)
	}
}

func CreatePool(pool Pool, poolDataId int64) {
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
	_, errExec := statement.Exec(eventName, isReadOnly, message, poolDataId, poolKey, PoolKeyReadOnly)
	if errExec != nil {
		log.Fatal("Invalid insert query", errExec)
	}
}

func LoadInterval(id int64) []Interval {
	is := make([]Interval, 0, 16)

	rows, err := database.Query("SELECT StartTime, StopTime FROM intervals WHERE Id", id)
	if err != nil {
		log.Fatal("Unable to execute interval query ", err)
	}

	for rows.Next() {
		var StartTime int64
		var StopTime int64

		rows.Scan(&StartTime, &StopTime)
		i := Interval{}
		i.StartTime = StartTime
		i.StopTime = StopTime

		is = append(is, i)
	}

	return is
}

func LoadStopwatch(id int64) []Stopwatch {
	sps := make([]Stopwatch, 0, 16)

	rows, err := database.Query("SELECT Color, Id, IntervalId, Name FROM stopwatches WHERE Id = ?", id)
	if err != nil {
		log.Fatal("Unable to execute stopwatch query ", err)
	}

	for rows.Next() {
		var Color string
		var Id int
		var IntervalId int64
		var Name string

		rows.Scan(&Color, &Id, &IntervalId, &Name)
		s := Stopwatch{}
		s.Color = Color
		s.Id = Id
		s.Name = Name

		if IntervalId != 0 {
			s.Intervals = LoadInterval(IntervalId)
		}

		sps = append(sps, s)
	}

	return sps
}

func LoadPoolData(id int64) PoolData {
	rows, err := database.Query("SELECT CreationDate, LastModDate, StopwatchId FROM pool_datas WHERE Id = ?", id)
	if err != nil {
		log.Fatal("Unable to execute pool_data query ", err)
	}

	var pd PoolData
	for rows.Next() {
		var CreationDate int64
		var LastModDate int64
		var StopwatchId int64

		rows.Scan(&CreationDate, &LastModDate, &StopwatchId)
		p := PoolData{}
		p.CreationDate = CreationDate
		p.LastModDate = LastModDate

		if StopwatchId != 0 {
			p.Stopwatches = LoadStopwatch(StopwatchId)
		}

		pd = p
	}

	return pd
}

func LoadPool(pool Pool) Pool {
	log.Println("Executing query")
	rows, err := database.Query("SELECT EventName, IsReadOnly, Message, PoolDataId, PoolKey, PoolKeyReadOnly FROM pools WHERE poolKey = ?", pool.PoolKey)
	if err != nil {
		log.Fatal("Unable to execute pool query", err)
	}

	var po Pool
	var poolDataId int64
	for rows.Next() {
		var EventName string
		var IsReadOnly bool
		var Message string
		var PoolDataId int64
		var PoolKey string
		var PoolKeyReadOnly string

		rows.Scan(&EventName, &IsReadOnly, &Message, &PoolDataId, &PoolKey, &PoolKeyReadOnly)
		p := Pool{}
		p.EventName = EventName
		p.IsReadOnly = IsReadOnly
		p.Message = Message
		p.PoolKey = PoolKey
		p.PoolKeyReadOnly = PoolKeyReadOnly
		po = p

		poolDataId = PoolDataId
	}

	pd := LoadPoolData(poolDataId)
	po.PoolData = pd

	return po
}


