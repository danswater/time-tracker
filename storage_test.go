package main

import (
	"os"
	"testing"
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
