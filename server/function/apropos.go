package engine

import (
    "html/template"
    "net/http"
    "log"
)

func (apropos *Engine) Apropos(w http.ResponseWriter, r *http.Request) {
    // Correction du chemin du template
    tmpl := template.Must(template.ParseFiles("front/template/apropos.html"))
    
    data := Engine{
        ActiveTab: "apropos", // Pour indiquer que nous sommes sur la page à propos
    }
    
    err := tmpl.Execute(w, data)
    if err != nil {
        log.Printf("Erreur d'exécution du template: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}