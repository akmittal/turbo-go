package message

type Message struct {
	text string
}

func New(text string) Message {
	return Message{text}
}
