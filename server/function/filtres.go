package engine

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func (groupie *Engine) Filtres() {

	// Tri alphabétique A-Z
	if groupie.Filtre.NameArtisteAZ {
		sort.Slice(groupie.Artists, func(i, j int) bool {
			return strings.ToLower(groupie.Artists[i].Name) < strings.ToLower(groupie.Artists[j].Name)
		})
	}

	// Tri alphabétique Z-A
	if groupie.Filtre.NameArtisteZA {
		sort.Slice(groupie.Artists, func(i, j int) bool {
			return strings.ToLower(groupie.Artists[i].Name) > strings.ToLower(groupie.Artists[j].Name)
		})
	}

	// Filtre par mot clé (Name, Members, CreationDate, FirstAlbum, DatesLocations)
	if groupie.Recherche.Mot != "" {
		var filteredArtists []ArtistStruct
		for _, artist := range groupie.Artists {
			if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(groupie.Recherche.Mot)) ||
				containsPrefix(artist.Members, groupie.Recherche.Mot) ||
				strings.Contains(strings.ToLower(fmt.Sprintf("%d", artist.CreationDate)), strings.ToLower(groupie.Recherche.Mot)) ||
				strings.Contains(strings.ToLower(artist.FirstAlbum), strings.ToLower(groupie.Recherche.Mot)) ||
				strings.Contains(strings.ToLower(fmt.Sprintf("%v", artist.DatesLocations)), strings.ToLower(groupie.Recherche.Mot)) {
				filteredArtists = append(filteredArtists, artist)
			}
		}
		groupie.Artists = filteredArtists
	}

	// Filtre par date de création
	if groupie.Filtre.CreationDate {
		var filteredArtists []ArtistStruct
		for _, artist := range groupie.Artists {
			if artist.CreationDate >= groupie.GestionFiltre.minCreationDate && artist.CreationDate <= groupie.GestionFiltre.maxCreationDate {
				filteredArtists = append(filteredArtists, artist)
			}
		}
		groupie.Artists = filteredArtists
	}

	// Filtre par premier album
	if groupie.Filtre.FirstAlbum {
		var filteredArtists []ArtistStruct
		for _, artist := range groupie.Artists {
			layout := "02-01-2006"
			firstAlbumDate, _ := time.Parse(layout, artist.FirstAlbum)
			firstAlbumYear := firstAlbumDate.Year()
			if firstAlbumYear >= groupie.GestionFiltre.minFirstAlbum && firstAlbumYear <= groupie.GestionFiltre.maxFirstAlbum {
				filteredArtists = append(filteredArtists, artist)
			}
		}
		groupie.Artists = filteredArtists
	}

	// Filtre par nombre de membres
	if groupie.Filtre.Members {
		var filteredArtists []ArtistStruct
		for _, artist := range groupie.Artists {
			if len(artist.Members) == groupie.GestionFiltre.NombreMembres {
				filteredArtists = append(filteredArtists, artist)
			}
		}
		groupie.Artists = filteredArtists
	}
}
