package engine

import (
	"html/template"
	"net/http"
)

func (groupie *Engine) Apropos(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("front/template/apropos.html"))

	data := Engine{}


	tmpl.Execute(w, data)
}
