package main

import (
	"testing"
)

func TestStartTrack(t *testing.T) {
	t1 := NewTransaction(1)
	t1.StartTrack()
	t1.StopTrack()

	_, err := t1.ComputeDuration()
	if err != nil {
		t.Error(err)
	}
}
