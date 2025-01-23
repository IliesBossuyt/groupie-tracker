package engine

import (
	"encoding/json"
	"io"
	"net/http"
)

func (groupie *Engine) Init() {
	// Recuperation des données de l'api avec requete GET
	resp1, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	resp2, _ := http.Get("https://groupietrackers.herokuapp.com/api/relation")

	// Lit le corps de la réponse HTTP pour les artistes
	body1, _ := io.ReadAll(resp1.Body)
	body2, _ := io.ReadAll(resp2.Body)

	// Désérialise les données JSON et les stocke dans groupie.Api et dans groupie.Relations
	json.Unmarshal(body1, &groupie.Api)
	json.Unmarshal(body2, &groupie.Relations)

	// Ferme les corps des réponses HTTP
	resp1.Body.Close()
	resp2.Body.Close()

	// Fusion des données
	for i, artist := range groupie.Api {
		for _, relation := range groupie.Relations.Index {
			if artist.ID == relation.ID {
				groupie.Api[i].DatesLocations = Underscore(relation.DatesLocations)
				break
			}
		}
	}

	// Ajout des données sans la strucure Artists
	for _, add := range groupie.Api {
		artist := ArtistStruct(add)
		groupie.Artists = append(groupie.Artists, artist)
	}

	// Initialisation des images aléatoire du carrousel
	groupie.RandomImg()

}
