package turbo

import (
	"html/template"
	"log"
	"net/http"

	"github.com/akmittal/turbo-go/internal/util"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options
// Stream Create a new Turbo stream with Action and Data channel.
type Stream struct {
	Action   Action
	Template *template.Template
	Target   string
	Data     chan interface{}
}

// Stream start streaming messages to all hub clients
func (s *Stream) Stream(hub *Hub, rw http.ResponseWriter, req *http.Request) {

	var turboTemplate, err = util.WrapTemplateInTurbo(s.Template.Name())
	if err != nil {
		http.Error(rw, "Error", 500)
	}
	conn, err := upgrader.Upgrade(rw, req, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	hub.register <- client

	temp, err := s.Template.New("userTemplate").Parse(turboTemplate)
	go client.writePump()

	if err != nil {
		http.Error(rw, "Error parsing template", 500)
	}
	for datum := range s.Data {

		turbo := Turbo{
			Action:   s.Action,
			Template: temp,
			Target:   s.Target,
			Data:     datum,
		}
		turbo.sendSocket(hub)

	}
}
