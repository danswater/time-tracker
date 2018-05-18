package main

import (
	"fmt"
)

func main() {
	t1 := NewTransaction(1)
	t1.StartTrack()
	t1.StopTrack()

	result, err := t1.ComputeDuration()
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
