package engine

import (
    "html/template"
    "net/http"
    "log"
)

func (e *Engine) PageArtiste(w http.ResponseWriter, r *http.Request) {
    // Correction du chemin du template
    tmpl := template.Must(template.ParseFiles("front/template/pageartiste.html"))
    
    data := Engine{
        ActiveTab: "pageartiste", // Pour indiquer que nous sommes sur la page artiste
    }
    
    err := tmpl.Execute(w, data)
    if err != nil {
        log.Printf("Erreur d'exécution du template: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}