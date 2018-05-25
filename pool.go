package main

import(
	"log"
)

type Pool struct {
	EventName       string   `json:EventName`
	IsReadOnly      bool     `json:IsReadOnly`
	Message         string   `json:Message`
	PoolData        PoolData `json:PoolData`
	PoolKey         string   `json:PoolKey`
	PoolKeyReadOnly string   `json:PoolKeyReadOnly`
}

func NewPool(eventName string, isReadOnly bool, message string, poolData PoolData, poolKey string, poolKeyReadOnly string) Pool {
	log.Println("New Pool ", eventName, isReadOnly, message, poolData, poolKey, poolKeyReadOnly)
	p := Pool{}
	p.EventName = eventName
	p.IsReadOnly = isReadOnly
	p.Message = message
	p.PoolData = poolData
	p.PoolKey = poolKey
	p.PoolKeyReadOnly = poolKeyReadOnly
	return p
}
