package engine

import (
	"html/template"
	"net/http"
)

func (groupie *Engine) Handler(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier index.html
	tmpl := template.Must(template.ParseFiles("front/template/home.html"))

	// Je crée une variable qui définit ma structure
	data := Engine{

	}

	tmpl.Execute(w, data)
}
