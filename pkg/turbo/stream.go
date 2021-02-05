package turbo

import (
	"html/template"
	"net/http"

	"github.com/akmittal/turbo-go/internal/util"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

type Stream struct {
	Action   Action
	Template *template.Template
	Target   string
	Data     chan interface{}
}

func (s *Stream) Stream(rw http.ResponseWriter, req *http.Request) {

	var turboTemplate, err = util.WrapTemplateInTurbo(s.Template.Name())
	if err != nil {
		http.Error(rw, "Error", 500)
	}
	c, err := upgrader.Upgrade(rw, req, nil)
	defer c.Close()
	if err != nil {
		http.Error(rw, "Error", 500)
	}

	temp, err := s.Template.New("userTemplate").Parse(turboTemplate)
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
		turbo.SendSocket(c)

	}
}
