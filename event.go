package main

type Event struct {
	Color     string `json:Color`
	EventName string `json:EventName`
	Id        int    `json:Id`
	Name      string `json:Name`
	Payload   string `json:Payload`
	Time      int64  `json:Time`
}
