package helper

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const templateDir = "./web/templates/"

func executeOrError(t *template.Template, d any, w http.ResponseWriter, e error) {
	if e != nil {
		log.Fatal(e)
	}

	e = t.Execute(w, d)
	if e != nil {
		log.Fatal(e)
	}
}

type Content struct {
	Data     any
	IsLogged interface{}
}

func TemplateParse(templateFile string, data any, w http.ResponseWriter) {

	content := Content{
		Data:     data,
		IsLogged: Session.Values["isLoggedIn"],
	}

	t, err := template.ParseFiles(templateDir+"base.html", templateDir+templateFile)
	executeOrError(t, content, w, err)
}

func TemplateInternalError(data any, w http.ResponseWriter) {
	t, err := template.ParseFiles(templateDir + "error.html")
	executeOrError(t, data, w, err)
}

func TemplateNotFound(w http.ResponseWriter) {
	t, err := template.ParseFiles(templateDir + "error.html")
	w.WriteHeader(http.StatusNotFound)
	data := fmt.Sprintf("Page not found %v", http.StatusNotFound)
	executeOrError(t, data, w, err)
}
