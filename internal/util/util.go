package util

import (
	"bytes"
	"text/template"
)

const temp = `<turbo-stream action="{{.Action}}" target="{{.Target}}">
<template>
{{template "[[.]]" .Data}}
</template>
</turbo-stream>`

var parsed *template.Template

func init() {
	parsed, _ = template.New("text").Delims("[[", "]]").Parse(temp)

}

func WrapTemplateInTurbo(name string) (string, error) {

	var buf bytes.Buffer

	err := parsed.Execute(&buf, name)
	return buf.String(), err
}
