package engine

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-resty/resty/v2"
)

func (groupie *Engine) PageArtiste(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("front/template/pageartiste.html"))

	// Récupérer l'ID de l'artiste depuis l'URL
	idParam := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idParam)

	// Trouver l'artiste correspondant
	var selectedArtist *ApiStruct
	for _, artist := range groupie.Api {
		if artist.ID == id {
			selectedArtist = &artist
			break
		}
	}

	// Récuperer et assigner les données pour la page
	pageArtist := PageArtistStruct{
		ID:             selectedArtist.ID,
		Image:          selectedArtist.Image,
		Name:           selectedArtist.Name,
		Members:        selectedArtist.Members,
		CreationDate:   selectedArtist.CreationDate,
		FirstAlbum:     selectedArtist.FirstAlbum,
		DatesLocations: selectedArtist.DatesLocations,
	}
	groupie.PageArtistStruct = pageArtist

	// Récupérer le mot pour une recherche dans le Home
	mot := r.FormValue("mot")
	groupie.Recherche.Mot = mot
	groupie.Filtres()

	// Vérifier si c'est une requete AJAX
	if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		// On regarde le type de données demandé
		dataType := r.URL.Query().Get("type")

		switch dataType {
		case "artist":
			// Renvoie le JSON de l'artiste
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(pageArtist)
			return

		case "location":
			// Renvoie le JSON des localisations

			// Liste pour stocker les noms des localisations à géolocaliser
			locations := []string{}

			// Récuperer les locations a géolocaliser
			for location := range groupie.PageArtistStruct.DatesLocations {
				locations = append(locations, location)
			}

			// Liste pour stocker les coordonnées géographiques obtenues
			coordinates := []Location{}

			// URL de l'API Nominatim pour la géolocalisation
			apiURL := "https://nominatim.openstreetmap.org/search"

			// Client HTTP avec une configuration de l'User-Agent
			client := resty.New()
			client.SetHeader("User-Agent", "MyTrackerApp/1.0 (contact: contact@mysite.com)")

			// Boucle sur chaque localisation pour obtenir les coordonnées avec l'API
			for _, location := range locations {
				resp, _ := client.R().
					SetQueryParams(map[string]string{
						"q":      location, // Nom de la localisation
						"format": "json",   // Format de la réponse
						"limit":  "1",      // Limite à une seule réponse par localisation
					}).
					Get(apiURL)

				// Parse la réponse JSON
				var results []Location
				_ = json.Unmarshal(resp.Body(), &results)

				// Si au moins un résultat est trouvé, on récupère le premier
				if len(results) > 0 {
					coordinates = append(coordinates, results[0])
				}
			}

			// Définit le type de contenu comme JSON
			w.Header().Set("Content-Type", "application/json")
			// Encode et renvoie les coordonnées sous forme de JSON
			json.NewEncoder(w).Encode(coordinates)
			return

		}
	}

	// Si la requête n'est pas une demande AJAX, on renvoie la page HTML
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, groupie)
}
