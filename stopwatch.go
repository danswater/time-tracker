package main

import (
	"log"
)

type Stopwatch struct {
	Color     string     `json:Color`
	Id        int        `json:Id`
	Intervals []Interval `json:Intervals`
	Name      string     `json:Name`
}

func NewStopwatch(id int, color string, name string) Stopwatch {
	log.Println("New stopwatch", id, color, name)
	sw := Stopwatch{}
	sw.Id = id
	sw.Color = color
	sw.Name = name
	return sw
}

func NewStopwatchWithIntervals(id int, color string, name string, intervals []Interval) Stopwatch {
	sw := NewStopwatch(id, color, name)
	sw.Intervals = intervals
	return sw
}
