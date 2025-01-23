package engine

import (
	"html/template"
	"net/http"
)

func (groupie *Engine) Accueil(w http.ResponseWriter, r *http.Request) {
	// Correction du chemin du template
	tmpl := template.Must(template.ParseFiles("front/template/accueil.html"))

	data := Engine{}

	tmpl.Execute(w, data)

}
