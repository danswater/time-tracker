package main

import (
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	InitializeStorage("./test.db")
	code := m.Run()
	os.Remove("./test.db")
	os.Exit(code)
}

func TestSaveInterval(t *testing.T) {
	interval := Interval{}
	interval.StartTime = 1526986686
	interval.StopTime = 1526986702

	id := SaveInterval(interval)
	if id == 0 {
		t.Fatal("Id must automatically increment")
	}
}

func TestSaveStopwatch( t *testing.T) {
	stopwatch := Stopwatch{}
	stopwatch.Color = "#fff"
	stopwatch.Id = 321321313
	stopwatch.Name = "Hello"

	intervalId := int64(1)
	id := SaveStopwatch( stopwatch, intervalId)
	if id == 0 {
		t.Fatal("Id must automatically increment")
	}
}

func TestCreatePool( t *testing.T) {
	// send poolNew
	unix := time.Now().Unix()
	pd := PoolData{}
	pd.CreationDate = unix
	pd.LastModDate = unix

	pool := Pool{}
	pool.EventName = "poolNew"
	pool.IsReadOnly = false
	pool.PoolData = pd
	pool.PoolKey = "dasdas3123"
	pool.PoolKeyReadOnly = "dsadsdasd43"

	CreatePool(pool, 1)
}

func TestCreatePoolData( t *testing.T) {
	pd := PoolData{}
	pd.CreationDate = 312313123
	pd.LastModDate = 321321313

	stopwatchId := int64(1)

	id := CreatePoolData( pd, &stopwatchId)
	if id == 0 {
		t.Fatal("Id must automatically increment")
	}
}

func TestUpdatePoolData(t *testing.T) {
	stopwatches := make([]Stopwatch, 0, 16)

	stopwatch := Stopwatch{}
	stopwatch.Id = 1
	stopwatch.Color = "#fff"
	stopwatch.Name = "hello"

	stopwatches = append( stopwatches, stopwatch )

	unix := time.Now().Unix()
	poolData := PoolData{}
	poolData.LastModDate = unix
	poolData.Stopwatches = stopwatches

	p := Pool{}
	p.PoolKey = "dasdas3123"

	stopwatchId := int64(2)
	res := UpdatePoolData(p, poolData, &stopwatchId)
	if res == 0 {
		t.Fatal("Must return non zero value")
	}
}

func TestLoadPool( t *testing.T) {
	pool := Pool{}
	pool.EventName = "poolNew"
	pool.IsReadOnly = false
	pool.PoolKey = "dasdas3123"
	pool.PoolKeyReadOnly = "dsadsdasd43"

	LoadPool(pool)
}
