package main

import (
	"fmt"
	"time"
	"errors"
)

// A Transaction describes a single log
type Transaction struct {
	UID            int
	Active         bool
	StartTimestamp time.Time
	EndTimestamp   time.Time
}

// StartTrack : this will start tracking transcation
func (t *Transaction) StartTrack() {
	t.Active = true
	t.StartTimestamp = time.Now()
}

// StopTrack : this will stop tracking transaction
func (t *Transaction) StopTrack() {
	t.Active = false
	t.EndTimestamp = time.Now()
}

// ComputeDuration : this will try to compute the duration of a transaction
func (t Transaction) ComputeDuration() (time.Duration, error) {
	if t.Active == true {
		return 0, errors.New("Must stop the track first")
	}
	return t.EndTimestamp.Sub(t.StartTimestamp), nil
}

func (t Transaction) String() string {
	return fmt.Sprintf("%v %v (%v) (%v)",
		t.UID,
		t.Active,
		t.StartTimestamp.Format("2006-01-02 15:04:05"),
		t.EndTimestamp.Format("2006-01-02 15:04:05"),
	)
}

// NewTransaction : this will create new instance of a transaction
func NewTransaction(uid int) Transaction {
	t1 := Transaction{}
	t1.UID = uid

	return t1
}
