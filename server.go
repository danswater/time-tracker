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
			// create pool data
			unix := time.Now().Unix()
			rawPd := NewPoolData(unix, unix)
			pd := CreatePoolData(rawPd)

			// create pool
			rawp := NewPool("poolNew", false, "", pd, randomString(), randomString())
			p := CreatePool(rawp)

			// emit pool new
			res, _ := json.Marshal(p)
			so.Emit("poolNew", string(res))
			so.BroadcastTo("Stopwatch", "poolNew", string(res))
		} else {
			// load pool
			p := Pool{}
			p.PoolKey = event.Payload

			pool := LoadPool(p)
			pool.EventName = "subscribeAccepted"

			// emit subscribe accepted
			res, _ := json.Marshal(pool)
			so.Emit( "subscribeAccepted", string(res))
			so.BroadcastTo("Stopwatch", "subscribeAccepted", string(res))
		}
	})

	so.On("stopwatchNew", func(msg string) {
		log.Println("received", msg)

		event := Event{}
		json.Unmarshal([]byte(msg), &event)

		// load pool
		refP := Pool{}
		refP.PoolKey = event.Payload
		refPool := LoadPool(refP)

		// create stopwatch with intervalid
		rawSp := NewStopwatch(event.Id, event.Color, event.Name)
		sp := CreateStopwatch(rawSp)

		// create interval
		rawI := NewInterval(sp.Id, event.Time, 0)
		CreateInterval(rawI)

		// load and update pool data with stopwatch
		pd := LoadPoolData(refPool.PoolData.Id)
		unix := time.Now().Unix()
		pd.LastModDate = unix
		pd.StopwatchId = sp.Id
		UpdatePoolData(pd)

		// load pool
		p := Pool{}
		p.PoolKey = event.Payload

		pool := LoadPool(p)
		pool.EventName = "poolChanged"

		// emit pool changed
		res, _ := json.Marshal(pool)
		so.Emit("poolChanged", string(res))
		so.BroadcastTo("Stopwatch", "poolChanged", string(res))
	})

	so.On("entityStop", func(msg string) {
		log.Println("received", msg)

		event := Event{}
		json.Unmarshal([]byte(msg), &event)

		// load and update interval stoptime
		id := int64(event.Id)
		i := LoadInterval(id)
		i.StopTime = event.Time

		UpdateInterval(i)

		// load and update pool data with stopwatch
		pd := LoadPoolDataByStopwatchId(int64(i.StopwatchId))
		unix := time.Now().Unix()
		pd.LastModDate = unix
		UpdatePoolData(pd)

		// load pool
		p := Pool{}
		p.PoolKey = event.Payload

		pool := LoadPool(p)
		pool.EventName = "poolChanged"

		// emit pool changed
		res, _ := json.Marshal(pool)
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
