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

func TestSavePoolData( t *testing.T) {
	pd := PoolData{}
	pd.CreationDate = 312313123
	pd.LastModDate = 321321313

	stopwatchId := int64(1)

	id := SavePoolData( pd, &stopwatchId)
	if id == 0 {
		t.Fatal("Id must automatically increment")
	}
}

func TestSavePool( t *testing.T) {
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

	SavePool(pool)
}

func TestLoadPool( t *testing.T) {
	pool := Pool{}
	pool.EventName = "poolNew"
	pool.IsReadOnly = false
	pool.PoolKey = "dasdas3123"
	pool.PoolKeyReadOnly = "dsadsdasd43"

	LoadPool(pool)
}
