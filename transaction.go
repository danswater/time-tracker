package main

import (
	"fmt"
	"time"
	"errors"
)

// A Transaction describes a single log
type Transaction struct {
	UID            int    `json:UID`
	Active         bool   `json:Active`
	Name           string `json:Name`
	StartTimestamp int    `json:StartTimestamp`
	EndTimestamp   int    `json:EndTimestamp`
}

// ComputeDuration : this will try to compute the duration of a transaction
func (t Transaction) ComputeDuration() (time.Duration, error) {
	if t.Active == true {
		return 0, errors.New("Must stop the track first")
	}

	start64 := int64(t.StartTimestamp)
	end64 := int64(t.EndTimestamp)
	startTimestamp := time.Unix(start64, 0)
	endTimestamp := time.Unix(end64, 0)

	return endTimestamp.Sub(startTimestamp), nil
}

func (t Transaction) String() string {
	return fmt.Sprintf("%v %v %v (%v) (%v)",
		t.UID,
		t.Active,
		t.Name,
		t.StartTimestamp,
		t.EndTimestamp,
	)
}

// NewTransaction : this will create new instance of a transaction
func NewTransaction(uid int, active bool, name string, startTimestamp int) Transaction {
	t := Transaction{}
	t.UID = uid
	t.Active = active
	t.Name = name
	t.StartTimestamp = startTimestamp

	return t
}
