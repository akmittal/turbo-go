package room

import (
	"github.com/akmittal/turbo-go/examples/hotwire/message"
)

type Room struct {
	Name     string `json:"name"`
	Messages []message.Message
}

func New(text string) Room {

	return Room{text, []message.Message{}}
}
func (r Room) AddMessage(text string) {
	r.Messages = append(r.Messages, message.New(text))
}
