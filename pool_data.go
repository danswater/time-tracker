package main

type PoolData struct {
	CreationDate int64       `json:CreationDate`
	LastModDate  int64       `json:LastModDate`
	Stopwatches  []Stopwatch `json:Stopwatches`
}
