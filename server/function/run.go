package engine

import "net/http"

func Run(groupie *Engine) {
	http.HandleFunc("/", groupie.Handler) // Ici, quand on arrive sur la racine, on appelle la fonction Handler
	http.HandleFunc("/apropos", groupie.Apropos)
	http.HandleFunc("/pageartiste", groupie.PageArtiste)

	fs := http.FileServer(http.Dir("front/./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/groupie", groupie.Handler) // Ici, on redirige vers /hangman pour effectuer les fonctions POST
	http.ListenAndServe(":8080", nil)
	// On lance le serveur local sur le port 8080
}
