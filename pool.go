package main

type Pool struct {
	EventName       string   `json:EventName`
	IsReadOnly      bool     `json:IsReadOnly`
	Message         string   `json:Message`
	PoolData        PoolData `json:PoolData`
	PoolKey         string   `json:PoolKey`
	PoolKeyReadOnly string   `json:PoolKeyReadOnly`
}
