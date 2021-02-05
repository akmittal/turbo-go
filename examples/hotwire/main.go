package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/akmittal/turbo-go/examples/hotwire/room"
	"github.com/go-chi/chi"
)

var rooms []room.Room

func main() {

	mux := chi.NewMux()
	mux.Get("/", getIndex)
	mux.Get("/room", getRooms)
	mux.Post("/room", createRoom)
	mux.Get("/room/create", createRoomPage)

	mux.Get("/socket", getSocket)
	http.ListenAndServe(":8000", mux)
}
func getRooms(rw http.ResponseWriter, req *http.Request) {
	temp, _ := template.ParseFiles("examples/hotwire/templates/rooms.temp.html", "examples/hotwire/templates/base.temp.html")
	temp.Execute(rw, rooms)
}
func createRoom(rw http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {

		return
	}
	// fmt.Fprintf(rw, "Post from website! r.PostFrom = %v\n", req.PostForm)
	var room room.Room
	room.Name = req.FormValue("name")

	rooms = append(rooms, room)
	http.Redirect(rw, req, "/room", 200)
}
func createRoomPage(rw http.ResponseWriter, req *http.Request) {
	temp, _ := template.ParseFiles("examples/hotwire/templates/create-room.temp.html", "examples/hotwire/templates/base.temp.html")
	temp.Execute(rw, rooms)
}
func getIndex(rw http.ResponseWriter, req *http.Request) {
	temp, _ := template.ParseFiles("examples/timestamp/index.temp.html")
	temp.Execute(rw, nil)
}

func getSocket(rw http.ResponseWriter, req *http.Request) {

	// tempChan := make(chan interface{})
	// appendMessage := turbo.Stream{
	// 	Action:   turbo.UPDATE,
	// 	Template: parsed,
	// 	Target:   "currenttime",
	// 	Data:     tempChan,
	// }

	// go sendMessages(tempChan)
	// appendMessage.Stream(rw, req)

}
func sendMessages(data chan interface{}) {
	for {
		data <- time.Now().Format("January 02, 2006 15:04:05")
		time.Sleep(time.Second)
	}

}
