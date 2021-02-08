## Turbo-go

Build hotwire applications using Go



## Example
Examples are in [examples](http://github.com/akmittal/turbo-go/tree/master/examples) directory

## Install turbo-go
``` text
go get github.com/akmittal/turbo-go
```

## API
```github.com/akmittal/turbo-go/pkg/turbo```

Send single templae update
``` go
messageTemp, err := template.New("message").parse(`<div>{{.}}</div>`)
data := time.Now()
turbo := turbo.Turbo{
		Action:   turbo.APPEND, // Action can be UPDATE, APPEND, PREPEND, REPLACE, REMOVE
		Template: messageTemp,
		Target:   "messages",
		Data:     data,
	}
```

Send stream of templates 
Create hub
``` go 
func main(){
	hub := turbo.NewHub()
    go hub.Run()
    mux.Get("/socket", func(rw http.ResponseWriter, req *http.Request) {
		getSocket(msgChan, hub, rw, req)
	})
}

func getSocket(msgChan chan interface{}, hub *turbo.Hub, rw http.ResponseWriter, req *http.Request) {
	temp, _ := template.ParseFiles("templates/messages.temp.html")
	messageTemp := temp.Lookup("message")

	appendMessage := turbo.Stream{
		Action:   turbo.APPEND,
		Template: messageTemp,
		Target:   "messages",
		Data:     msgChan,
	}

	appendMessage.Stream(hub, rw, req)
}


```