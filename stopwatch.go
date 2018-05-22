package main

type Stopwatch struct {
	Color     string     `json:Color`
	Id        int        `json:Id`
	Intervals []Interval `json:Intervals`
	Name      string     `json:Name`
}
