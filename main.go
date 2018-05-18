package main

import (
	"fmt"
	"time"
	"errors"
)

// A Transaction describes a single log
type Transaction struct {
	Uid            int
	Active         bool
	StartTimestamp time.Time
	EndTimestamp   time.Time
}

func (t *Transaction) StartTrack(uid int) {
	t.Uid = uid
	t.Active = true
	t.StartTimestamp = time.Now()
}

func (t *Transaction) StopTrack() {
	t.Active = false
	t.EndTimestamp = time.Now()
}

func (t Transaction) ComputeDuration() (time.Duration, error) {
	if t.Active == true {
		return 0, errors.New("Must stop the tack first")
	}
	return t.EndTimestamp.Sub(t.StartTimestamp), nil
}

func main() {
	t1 := Transaction{}
	t1.StartTrack(1)
	t1.StopTrack()

	result, err := t1.ComputeDuration()
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
