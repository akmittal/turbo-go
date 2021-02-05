package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/akmittal/turbo-go/pkg/turbo"
	"github.com/go-chi/chi"
)

const temp = `<h1>{{.}}</h1>`

var parsed *template.Template

func main() {
	parsed, _ = template.New("main").Parse(temp)
	mux := chi.NewMux()
	mux.Get("/", getIndex)

	mux.Get("/socket", getSocket)
	http.ListenAndServe(":8000", mux)
}
func getIndex(rw http.ResponseWriter, req *http.Request) {
	temp, _ := template.ParseFiles("examples/timestamp/index.temp.html")
	temp.Execute(rw, nil)
}

func getSocket(rw http.ResponseWriter, req *http.Request) {

	tempChan := make(chan interface{})
	appendMessage := turbo.Stream{
		Action:   turbo.UPDATE,
		Template: parsed,
		Target:   "currenttime",
		Data:     tempChan,
	}

	go sendMessages(tempChan)
	appendMessage.Stream(rw, req)

}
func sendMessages(data chan interface{}) {
	for {
		data <- time.Now().Format("January 02, 2006 15:04:05")
		time.Sleep(time.Second)
	}

}
