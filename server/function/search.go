package engine

import (
	"html/template"
	"net/http"
	
)

func (groupie *Engine) Search(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    search := r.FormValue("recherche")
    // Implémenter votre logique de recherche ici
    
    // Rediriger vers la page des résultats ou afficher les résultats
    data := Engine{
        // Ajouter les résultats de recherche ici
		Mot: search,
    }

    tmpl := template.Must(template.ParseFiles("front/template/home.html"))
    tmpl.Execute(w, data)
}