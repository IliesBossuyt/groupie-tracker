package engine

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ArtistResponse struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

func (groupie *Engine) GetLocation() {
	resp, _ := http.Get(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%d", groupie.ID))

	var location LocationResponse
	json.NewDecoder(resp.Body).Decode(&location)

	if groupie.ArtistStruct.Location == nil {
		groupie.ArtistStruct.Location = make(map[int][]string)
	}
	groupie.ArtistStruct.Location[groupie.ID] = append(groupie.ArtistStruct.Location[groupie.ID], location.Location...)
}

func (groupie *Engine) GetDates() {
	resp, _ := http.Get(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%d", groupie.ID))

	var date DateResponse
	json.NewDecoder(resp.Body).Decode(&date)

	if groupie.ArtistStruct.Date == nil {
		groupie.ArtistStruct.Date = make(map[int][]string)
	}
	for i, d := range date.Dates {
		date.Dates[i] = strings.Replace(d, "*", "", -1)
	}
	groupie.ArtistStruct.Date[groupie.ID] = append(groupie.ArtistStruct.Date[groupie.ID], date.Dates...)
}

func (groupie *Engine) GetRelation() {
	resp, _ := http.Get(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", groupie.ID))

	var relation RelationResponse
	json.NewDecoder(resp.Body).Decode(&relation)

	if groupie.ArtistStruct.Relation == nil {
		groupie.ArtistStruct.Relation = make(map[int][]string)
	}
	var relations []string
	for location, dates := range relation.DatesLocations {
		for _, date := range dates {
			relations = append(relations, fmt.Sprintf("%s: %s", location, date))
		}
		groupie.ArtistStruct.Relation[groupie.ID] = relations
	}
}

func (groupie *Engine) Artist() {
	resp, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	body, _ := io.ReadAll(resp.Body)
	var ArtistResponse []ArtistResponse
	json.Unmarshal(body, &ArtistResponse)

	for _, artist := range ArtistResponse {
		// Vérifier si le mot entrée est égal a l'API
		if strings.HasPrefix(strings.ToLower(artist.Name), strings.ToLower(groupie.Mot)) ||
			strings.HasPrefix(strings.ToLower(artist.Image), strings.ToLower(groupie.Mot)) ||
			containsPrefix(artist.Members, groupie.Mot) ||
			strings.HasPrefix(strings.ToLower(fmt.Sprintf("%d", artist.CreationDate)), strings.ToLower(groupie.Mot)) ||
			strings.HasPrefix(strings.ToLower(artist.FirstAlbum), strings.ToLower(groupie.Mot)) {
			fmt.Printf("ID: %d, Image: %s, Name: %s, Members: %v, Creation Date: %d, First Album: %s\n", artist.ID, artist.Image, artist.Name, artist.Members, artist.CreationDate, artist.FirstAlbum)
			// AJouter les données dans la structure
			groupie.ArtistStruct.ID = append(groupie.ArtistStruct.ID, artist.ID)
			groupie.ArtistStruct.Image = append(groupie.ArtistStruct.Image, artist.Image)
			groupie.ArtistStruct.Name = append(groupie.ArtistStruct.Name, artist.Name)

			if groupie.ArtistStruct.Members == nil {
				groupie.ArtistStruct.Members = make(map[int][]string)
			}
			groupie.ArtistStruct.Members[artist.ID] = append(groupie.ArtistStruct.Members[artist.ID], artist.Members...)

			groupie.ArtistStruct.CreationDate = append(groupie.ArtistStruct.CreationDate, artist.CreationDate)
			groupie.ArtistStruct.FirstAlbum = append(groupie.ArtistStruct.FirstAlbum, artist.FirstAlbum)
			// Recherhe dans les URL
			groupie.ID = artist.ID
			groupie.GetLocation()
			groupie.GetDates()
			groupie.GetRelation()
		}
	}
}

type LocationIndex struct {
	Index []LocationResponse `json:"index"`
}

type LocationResponse struct {
	ID       int      `json:"id"`
	Location []string `json:"locations"`
}

func (groupie *Engine) Location() {
	resp, _ := http.Get("https://groupietrackers.herokuapp.com/api/locations")

	body, _ := io.ReadAll(resp.Body)
	var LocationResponse LocationIndex
	json.Unmarshal(body, &LocationResponse)

	// Afficher les résultats triés
	for _, location := range LocationResponse.Index {
		if containsPrefix(location.Location, groupie.Mot) {
			fmt.Printf("ID: %d, Location: %s\n", location.ID, location.Location)

			groupie.LocationStruct.ID = append(groupie.LocationStruct.ID, location.ID)
			if groupie.LocationStruct.Location == nil {
				groupie.LocationStruct.Location = make(map[int][]string)
			}
			groupie.LocationStruct.Location[location.ID] = append(groupie.LocationStruct.Location[location.ID], location.Location...)
		}
	}
}

type DateIndex struct {
	Index []DateResponse `json:"index"`
}

type DateResponse struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

func (groupie *Engine) Date() {
	resp, _ := http.Get("https://groupietrackers.herokuapp.com/api/dates")

	body, _ := io.ReadAll(resp.Body)
	var DateResponse DateIndex
	json.Unmarshal(body, &DateResponse)

	// Afficher les résultats triés
	for _, date := range DateResponse.Index {
		for i, d := range date.Dates {
			date.Dates[i] = strings.Replace(d, "*", "", -1) // supprimer les *
		}
		if containsPrefix(date.Dates, groupie.Mot) {
			fmt.Printf("ID: %d, Dates: %s\n", date.ID, date.Dates)

			groupie.DateStruct.ID = append(groupie.DateStruct.ID, date.ID)
			if groupie.DateStruct.Dates == nil {
				groupie.DateStruct.Dates = make(map[int][]string)
			}
			groupie.DateStruct.Dates[date.ID] = append(groupie.DateStruct.Dates[date.ID], date.Dates...)
		}
	}
}

type RelationIndex struct {
	Index []RelationResponse `json:"index"`
}

type RelationResponse struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func (groupie *Engine) Relation() {
	resp, _ := http.Get("https://groupietrackers.herokuapp.com/api/relation")

	body, _ := io.ReadAll(resp.Body)
	var RelationResponse RelationIndex
	json.Unmarshal(body, &RelationResponse)
	// Initialisation de la map pour DatesLocations
	groupie.RelationStruct.DatesLocations = make(map[int][]string)

	// Traiter chaque relation de l'API
	for _, relation := range RelationResponse.Index {
		if containsStringInMap(relation.DatesLocations, groupie.Mot) || containsStringInMapKeys(relation.DatesLocations, groupie.Mot) {
			fmt.Printf("ID: %d, Dates: %s\n", relation.ID, relation.DatesLocations)
			// Ajouter l'ID à la liste
			groupie.RelationStruct.ID = append(groupie.RelationStruct.ID, relation.ID)

			// Créer une liste de dates à partir des localisations
			var relations []string
			for location, dates := range relation.DatesLocations {
				for _, date := range dates {
					// Ajouter chaque date dans la liste
					relations = append(relations, fmt.Sprintf("%s: %s", location, date))
				}
			}
			// Ajouter les relations (localisation + dates) à la map par ID
			groupie.RelationStruct.DatesLocations[relation.ID] = relations
		}
	}
}

func containsPrefix(slice []string, prefix string) bool {
	for _, s := range slice {
		if strings.HasPrefix(strings.ToLower(s), strings.ToLower(prefix)) {
			return true
		}
	}
	return false
}

func containsStringInMap(m map[string][]string, str string) bool {
	for _, v := range m {
		for _, s := range v {
			if strings.HasPrefix(strings.ToLower(s), strings.ToLower(str)) {
				return true
			}
		}
	}
	return false
}

func containsStringInMapKeys(m map[string][]string, str string) bool {
	for k := range m {
		if strings.HasPrefix(strings.ToLower(k), strings.ToLower(str)) {
			return true
		}
	}
	return false
}
