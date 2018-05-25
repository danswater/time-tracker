package main

import(
	"log"
)

type PoolData struct {
	Id           int64       `json:Id`
	CreationDate int64       `json:CreationDate`
	LastModDate  int64       `json:LastModDate`
	StopwatchId  int         `json:StopwatchId`
	Stopwatch    Stopwatch   `json:Stopwatch`
}


func NewPoolData(creationDate int64, lastModDate int64) PoolData {
	log.Println("New PoolData", creationDate, lastModDate)
	pd := PoolData{}
	pd.CreationDate = creationDate
	pd.LastModDate = lastModDate
	return pd
}

func NewPoolWithId(id int64, creationDate int64, lastModDate int64) PoolData {
	log.Println("New PoolData with Id", id, creationDate, lastModDate)
	pd := NewPoolData(creationDate, lastModDate)
	pd.Id = id
	return pd
}

func NewPoolWithStopwatches(id int64, creationDate int64, lastModDate int64, stopwatchId int, stopwatch Stopwatch) PoolData {
	log.Println("New PoolData with Stopwatch", id, creationDate, lastModDate)
	pd := NewPoolWithId(id, creationDate, lastModDate)
	pd.StopwatchId = stopwatchId
	pd.Stopwatch = stopwatch
	return pd
}
