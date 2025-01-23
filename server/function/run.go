package engine

import (
	"fmt"
	"net/http"
)

func Run(groupie *Engine) {
	// Définir les différentes routes
	http.HandleFunc("/", groupie.Accueil)
	http.HandleFunc("/home", groupie.Home)
	http.HandleFunc("/apropos", groupie.Apropos)
	http.HandleFunc("/pageartiste", groupie.PageArtiste)
	// Lancer le Init
	groupie.Init()

	fs := http.FileServer(http.Dir("front/./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/groupie", groupie.Home)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Serveur lancé sur http://localhost:8080/")
	// On lance le serveur local sur le port 8080
}
