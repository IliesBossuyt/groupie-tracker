package engine

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (groupie *Engine) FiltreArtist() {
	resp, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	body, _ := io.ReadAll(resp.Body)
	var ArtistResponse []ArtistResponse
	json.Unmarshal(body, &ArtistResponse)

	for _, artist := range ArtistResponse {
		// Vérifier si le mot entrée est égal a l'API
		if strings.HasPrefix(strings.ToLower(artist.Name), strings.ToLower(groupie.Mot)) {
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

func (groupie *Engine) FiltreCreationDate() {
	resp, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	body, _ := io.ReadAll(resp.Body)
	var ArtistResponse []ArtistResponse
	json.Unmarshal(body, &ArtistResponse)

	for _, artist := range ArtistResponse {
		// Vérifier si le mot entrée est égal a l'API
		if strings.HasPrefix(strings.ToLower(fmt.Sprintf("%d", artist.CreationDate)), strings.ToLower(groupie.Mot)) {
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

func (groupie *Engine) FiltreFirstAlbum() {
	resp, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	body, _ := io.ReadAll(resp.Body)
	var ArtistResponse []ArtistResponse
	json.Unmarshal(body, &ArtistResponse)

	for _, artist := range ArtistResponse {
		// Vérifier si le mot entrée est égal a l'API
		if strings.HasPrefix(strings.ToLower(artist.FirstAlbum), strings.ToLower(groupie.Mot)) {
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

func (groupie *Engine) FiltreMembers() {

}
