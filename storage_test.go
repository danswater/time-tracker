package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	InitializeStorage("./_test.db")
	code := m.Run()
	os.Remove("./_test.db")
	os.Exit(code)
}

func TestSave(t *testing.T) {
	t1 := NewTransaction(1)
	Save(t1)

	tns := Load(1)
	if len(tns) != 1 {
		t.Fatal("Wrong number of transactions" )
	}
}
