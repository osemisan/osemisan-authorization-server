package templates

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type RenderTemplateError struct {
	msg string
	err error
}

func (e *RenderTemplateError) Error() string {
	return fmt.Sprintf("cannot render template: %s (%s)", e.msg, e.err.Error())
}

func (e *RenderTemplateError) Unwrap() error {
	return e.err
}

func Render(name string, w http.ResponseWriter, data any) error {
	wd, err := os.Getwd()
	if err != nil {
		return &RenderTemplateError{ msg: "get working dir", err: err } 
	}
	f := wd + "/data/" + name + ".html"
	tmpl, err := template.ParseFiles(f)
	if err != nil {
		return &RenderTemplateError{ msg: "parse template file", err: err } 
	}
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, data)
	return nil
}
