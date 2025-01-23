package engine

import (
	"encoding/json"
	"html/template"
	"net/http"
)

func (groupie *Engine) Home(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("front/template/home.html"))

	// Envoie des données pour les images aleatoires du carroussel
	data := Engine{
		RandomImages: groupie.RandomImages,
	}

	// Génération des images aléatoires
	groupie.RandomImg()

	// Affichage des données de l'artiste dynamiquement avec XMLHttpRequest
	if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"artists": groupie.Artists,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if r.Method == http.MethodPost {

		// Filtre tri alphabétique
		menu := r.FormValue("menu")
		groupie.FiltreMenu(menu)

		// Filtre par date de creation
		minCreationDate := r.FormValue("minCreationDate")
		maxCreationDate := r.FormValue("maxCreationDate")
		groupie.FiltreCreationDate(minCreationDate, maxCreationDate)

		// Filtre par date de premiere album
		minFirstAlbum := r.FormValue("minFirstAlbum")
		maxFirstAlbum := r.FormValue("maxFirstAlbum")
		groupie.FiltreFirstAlbum(minFirstAlbum, maxFirstAlbum)

		// Filtre par nombre de membres
		membersString := r.FormValue("members")
		groupie.FiltreMembers(membersString)

		// Réinitialisation des filtres
		if r.FormValue("reset") == "reset" {
			groupie.ResetFiltre()
		}

		// Réinitalisation de la structure Artist, et remise à 0 des données depuis la structure API
		groupie.Reset()
		groupie.Clone()
		// Recherche de l'utilisateur
		mot := r.FormValue("mot")
		groupie.Recherche.Mot = mot
		groupie.Filtres()
	}

	tmpl.Execute(w, data)
}
