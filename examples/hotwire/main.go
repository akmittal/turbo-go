package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/akmittal/turbo-go/examples/hotwire/message"
	"github.com/akmittal/turbo-go/pkg/turbo"
	"github.com/go-chi/chi"
)

var messages []message.Message

func main() {

	mux := chi.NewMux()
	mux.Get("/", getIndex)

	hub := turbo.NewHub()
	msgChan := make(chan interface{})
	mux.Post("/send", func(rw http.ResponseWriter, req *http.Request) {
		sendMessage(msgChan, hub, rw, req)
	})
	go hub.Run()
	mux.Get("/socket", func(rw http.ResponseWriter, req *http.Request) {
		getSocket(msgChan, hub, rw, req)
	})
	http.ListenAndServe(":8000", mux)
}
func getIndex(rw http.ResponseWriter, req *http.Request) {
	temp, err := template.ParseFiles("examples/hotwire/templates/messages.temp.html", "examples/hotwire/templates/base.temp.html")
	if err != nil {
		http.Error(rw, "Error", 400)
	}
	temp.Execute(rw, messages)
}
func sendMessage(msgChan chan interface{}, hub *turbo.Hub, rw http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		http.Error(rw, err.Error(), 401)
		return
	}

	// fmt.Fprintf(rw, "Post from website! r.PostFrom = %v\n", req.PostForm)
	var msg message.Message
	msg.Text = req.FormValue("message")

	messages = append(messages, msg)
	go func() {
		msgChan <- msg
	}()
	fmt.Fprintf(rw, "%s", "Done")

}

func getSocket(msgChan chan interface{}, hub *turbo.Hub, rw http.ResponseWriter, req *http.Request) {
	temp, _ := template.ParseFiles("examples/hotwire/templates/messages.temp.html")
	messageTemp := temp.Lookup("message")

	appendMessage := turbo.Stream{
		Action:   turbo.APPEND,
		Template: messageTemp,
		Target:   "messages",
		Data:     msgChan,
	}

	appendMessage.Stream(hub, rw, req)
}
