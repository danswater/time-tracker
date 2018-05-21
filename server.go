package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/googollee/go-socket.io"
	"encoding/json"
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
	so.Join("Tracker")
	so.On("EventStopwatchNew", func(msg string) {
		log.Println("received", msg)

		t := Transaction{}
		json.Unmarshal([]byte(msg), &t)

		res, _ := json.Marshal(t)
		so.Emit("EventStopwatchNew", string(res))
	})
	so.On("disconnection", func() {
		log.Println("on disconnect")
	})
}
