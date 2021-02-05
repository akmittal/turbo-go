package turbo

import (
	"html/template"
	"net/http"

	"github.com/akmittal/turbo-go/internal/util"
)

type Update struct {
	Action   Action
	Template *template.Template
	Target   string
	Data     interface{}
}

func (u *Update) Send(rw http.ResponseWriter, req *http.Request) {
	var turboTemplate, err = util.WrapTemplateInTurbo(u.Template.Name())

	parsed, err := u.Template.New("userTemplate").Parse(turboTemplate)
	if err != nil {
		http.Error(rw, "Error parsing template", 500)
	}

	turbo := Turbo{
		Action:   u.Action,
		Template: parsed,
		Target:   u.Target,
		Data:     u.Data,
	}
	turbo.Send(rw)

}
