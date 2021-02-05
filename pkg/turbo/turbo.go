package turbo

import (
	"bytes"
	"html/template"
	"net/http"
)

type Action string

const (
	APPEND  Action = "append"
	PREPEND Action = "prepend"
	REPLACE Action = "replace"
	UPDATE  Action = "update"
	REMOVE  Action = "remove"
)

var parsedTemp *template.Template

type Turbo struct {
	Action   Action
	Template *template.Template
	Target   string
	Data     interface{}
}

func (h *Turbo) SetHeader(rw http.ResponseWriter) {
	rw.Header().Add("Content-type", "text/vnd.turbo-stream.html")
}

func (h *Turbo) Send(rw http.ResponseWriter) {
	rw.Header().Add("Content-type", "text/vnd.turbo-stream.html")
	h.Template.Execute(rw, h)
}
func (h *Turbo) sendSocket(hub *Hub) {
	var buf bytes.Buffer
	h.Template.Execute(&buf, h)
	hub.broadcast <- buf.Bytes()
}
