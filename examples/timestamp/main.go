package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/akmittal/turbo-go/pkg/turbo"
	"github.com/go-chi/chi"
)

const temp = `<h1>{{.}}</h1>`

func main() {

	mux := chi.NewMux()
	mux.Get("/", getIndex)

	hub := turbo.NewHub()
	go hub.Run()
	mux.Get("/socket", func(rw http.ResponseWriter, req *http.Request) {
		getSocket(hub, rw, req)
	})
	http.ListenAndServe(":8000", mux)
}
func getIndex(rw http.ResponseWriter, req *http.Request) {
	temp, _ := template.ParseFiles("examples/timestamp/index.temp.html")
	temp.Execute(rw, nil)
}

func getSocket(hub *turbo.Hub, rw http.ResponseWriter, req *http.Request) {

	parsed, err := template.New("main").Parse(temp)
	if err != nil {
		http.Error(rw, "Error", 500)
	}

	tempChan := make(chan interface{})
	appendMessage := turbo.Stream{
		Action:   turbo.UPDATE,
		Template: parsed,
		Target:   "currenttime",
		Data:     tempChan,
	}

	go sendMessages(tempChan)
	appendMessage.Stream(hub, rw, req)

}
func sendMessages(data chan interface{}) {
	for {
		data <- time.Now().Format("January 02, 2006 15:04:05")
		time.Sleep(time.Second)
	}

}
