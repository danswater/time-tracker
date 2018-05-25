package main

import (
	"log"
)

type Interval struct {
	Id          int64 `json:Id`
	StopwatchId int   `json:StopwatchId`
	StartTime   int64 `json:StartTime`
	StopTime    int64 `json:StopTime`
}

func NewInterval(stopwatchId int, startTime int64, stopTime int64) Interval {
	log.Println("New interval", stopwatchId, startTime, stopTime)
	i := Interval{}
	i.StopwatchId = stopwatchId
	i.StartTime = startTime
	i.StopTime = stopTime
	return i
}

func NewIntervalWithId(id int64, stopwatchId int, startTime int64, stopTime int64) Interval {
	log.Println("New interval with id", id, stopwatchId, startTime, stopTime)
	i := NewInterval(stopwatchId, startTime, stopTime)
	i.Id = id
	return i
}
