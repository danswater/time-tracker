package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/googollee/go-socket.io"
	"encoding/json"
	"time"
	"math/rand"
)

// StartServer : start new server
func StartServer() {
	// TODO ML Refactor into multiple files

	// Socket server
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal("err")
	}

	server.On("connection", SocketHandler)
	server.On("error", func( so socketio.Socket, err error) {
		log.Println("on disconnect")
	})

	r := mux.NewRouter()
	r.Handle("/socket.io/", server)
	r.PathPrefix("/").HandlerFunc(FileHandler)

	port := ":8080"
	log.Println("Starting to listen on port", port)

	http.ListenAndServe(port, r)
}

// FileHandler : handle static files
func FileHandler(w http.ResponseWriter, r *http.Request) {
	_, err := os.Stat("static/")
	if err != nil {
		log.Println("File not found ",err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	f := http.FileServer(http.Dir("./static/"))
	f.ServeHTTP(w, r)
	return
}

// SocketHandler : handle socket io
func SocketHandler(so socketio.Socket) {
	log.Println("on connection")

	so.Join("Stopwatch")

	so.On("subscribe", func(msg string) {
		event := Event{}
		json.Unmarshal([]byte(msg), &event)

		if event.Payload == "NEW" {
			// send poolNew
			unix := time.Now().Unix()
			poolData := PoolData{}
			poolData.CreationDate = unix
			poolData.LastModDate = unix

			pool := Pool{}
			pool.EventName = "poolNew"
			pool.IsReadOnly = false
			pool.PoolData = poolData
			pool.PoolKey = randomString()
			pool.PoolKeyReadOnly = randomString()

			SavePool(pool)

			res, _ := json.Marshal(pool)
			so.Emit("poolNew", string(res))
			so.BroadcastTo("Stopwatch", "poolNew", string(res))
		} else {
			p := Pool{}
			p.PoolKey = event.Payload

			pool := LoadPool(p)
			pool.EventName = "subscribeAccepted"

			res, _ := json.Marshal(pool)
			so.Emit( "subscribeAccepted", string(res))
			so.BroadcastTo("Stopwatch", "subscribeAccepted", string(res))
		}
	})

	so.On("stopwatchNew", func(msg string) {
		log.Println("received", msg)

		event := Event{}
		json.Unmarshal([]byte(msg), &event)

		// send poolChanged
		intervals := make([]Interval, 0, 16)

		interval := Interval{}
		interval.StartTime = event.Time

		intervals = append( intervals, interval )

		stopwatches := make([]Stopwatch, 0, 16)

		stopwatch := Stopwatch{}
		stopwatch.Id = event.Id
		stopwatch.Color = event.Color
		stopwatch.Name = event.Name
		stopwatch.Intervals = intervals

		stopwatches = append( stopwatches, stopwatch )

		unix := time.Now().Unix()
		poolData := PoolData{}
		poolData.CreationDate = unix
		poolData.LastModDate = unix
		poolData.Stopwatches = stopwatches

		pool := Pool{}
		pool.EventName = "poolChanged"
		pool.IsReadOnly = false
		pool.PoolData = poolData
		pool.PoolKey = randomString()
		pool.PoolKeyReadOnly = randomString()

		res, _ := json.Marshal(pool)

		log.Println("sending data as poolChanged", msg)
		so.Emit("poolChanged", string(res))
		so.BroadcastTo("Stopwatch", "poolChanged", string(res))
	})

	so.On("disconnection", func() {
		log.Println("on disconnect")
	})
}

func randomString() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	n := 8

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
