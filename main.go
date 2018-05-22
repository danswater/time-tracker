package main

import (
	"os"
)

func main() {
	databasePath := os.Getenv("TRACKER_DATABASE")
	if databasePath == "" {
		databasePath = "./data.db"
	}

	InitializeStorage(databasePath)

	StartServer()
}
