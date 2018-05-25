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

func TestCreatePool( t *testing.T) {
	// send poolNew
	rawPd := PoolData{}
	rawPd.Id = 1

	rawP := Pool{}
	rawP.EventName = "poolNew"
	rawP.IsReadOnly = false
	rawP.PoolData = rawPd
	rawP.PoolKey = "dasdas3123"
	rawP.PoolKeyReadOnly = "dsadsdasd43"

	pool := CreatePool(rawP)
	if pool.PoolKey != rawP.PoolKey {
		t.Fatal("Both pool keys must be the same")
	}
}

func TestLoadPool( t *testing.T) {
	rawP := Pool{}
	rawP.PoolKey = "dasdas3123"

	pool := LoadPool(rawP)
	if pool.EventName == "" {
		t.Fatal("pool should contain event name")
	}
}

func TestCreatePoolData( t *testing.T) {
	rawPd := PoolData{}
	rawPd.CreationDate = 312313123
	rawPd.LastModDate = 321321313

	pd := CreatePoolData(rawPd)
	if pd.Id == 0 {
		t.Fatal("Id must automatically increment")
	}
}

func TestUpdatePoolData(t *testing.T) {
	unix := time.Now().Unix()
	poolData := PoolData{}
	poolData.Id = 1
	poolData.LastModDate = unix


	res := UpdatePoolData(poolData)
	if res.Id == 0 {
		t.Fatal("Must return non zero value")
	}
}

func TestCreateStopwatch( t *testing.T) {
	rawSp := Stopwatch{}
	rawSp.Color = "#fff"
	rawSp.Id = 321321313
	rawSp.Name = "Hello"

	sp := CreateStopwatch(rawSp)
	if sp.Id == 0 {
		t.Fatal("Id must automatically increment")
	}
}

func TestLoadStopwatch( t *testing.T) {
	sp := LoadStopwatch(321321313)
	if sp.Id == 0 {
		t.Fatal("Stopwatch should container more then 0")
	}
}

func TestCreateInterval(t *testing.T) {
	rawI := Interval{}
	rawI.StopwatchId = 1
	rawI.StartTime = 1526986686
	rawI.StopTime = 1526986702

	i := CreateInterval(rawI)
	if i.Id == 0 {
		t.Fatal("Id must automatically increment")
	}
}
