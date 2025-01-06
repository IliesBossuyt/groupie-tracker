package engine

import (
	"fmt"
	"html/template"
	"net/http"
)

func (groupie *Engine) Handler(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier index.html
	tmpl := template.Must(template.ParseFiles("front/template/home.html"))

	// Je crée une variable qui définit ma structure
	data := Engine{
		ArtistStruct: struct {
			ID           []int
			Image        []string
			Name         []string
			Members      map[int][]string
			CreationDate []int
			FirstAlbum   []string
			Location     map[int][]string
			Date         map[int][]string
			Relation     map[int][]string
		}{
			ID:           groupie.ArtistStruct.ID,
			Image:        groupie.ArtistStruct.Image,
			Name:         groupie.ArtistStruct.Name,
			Members:      groupie.ArtistStruct.Members,
			CreationDate: groupie.ArtistStruct.CreationDate,
			FirstAlbum:   groupie.ArtistStruct.FirstAlbum,
			Location:     groupie.ArtistStruct.Location,
			Date:         groupie.ArtistStruct.Date,
			Relation:     groupie.ArtistStruct.Relation,
		},
		LocationStruct: struct {
			ID       []int
			Location map[int][]string
		}{
			ID:       groupie.LocationStruct.ID,
			Location: groupie.LocationStruct.Location,
		},
		DateStruct: struct {
			ID    []int
			Dates map[int][]string
		}{
			ID:    groupie.DateStruct.ID,
			Dates: groupie.DateStruct.Dates,
		},
		RelationStruct: struct {
			ID             []int
			DatesLocations map[int][]string
		}{
			ID:             groupie.RelationStruct.ID,
			DatesLocations: groupie.RelationStruct.DatesLocations,
		},
	}

	if r.Method == "POST" {
		groupie.ArtistStruct.ID = nil
		groupie.ArtistStruct.Image = nil
		groupie.ArtistStruct.Name = nil
		groupie.ArtistStruct.Members = nil
		groupie.ArtistStruct.CreationDate = nil
		groupie.ArtistStruct.FirstAlbum = nil
		groupie.ArtistStruct.Location = nil
		groupie.ArtistStruct.Date = nil
		groupie.LocationStruct.ID = nil
		groupie.LocationStruct.Location = nil
		groupie.DateStruct.ID = nil
		groupie.DateStruct.Dates = nil
		groupie.RelationStruct.ID = nil
		groupie.RelationStruct.DatesLocations = nil
		mot := r.FormValue("mot")
		groupie.Mot = mot
		fmt.Println(groupie.Mot)

		if !groupie.Filtre.NameArtiste && !groupie.Filtre.CreationDate && !groupie.Filtre.FirstAlbum {
			groupie.Artist()
			groupie.Location()
			groupie.Date()
			groupie.Relation()
		}
		if groupie.Filtre.NameArtiste {
			groupie.FiltreArtist()
		}
		if groupie.Filtre.CreationDate {
			groupie.FiltreCreationDate()
		}
		if groupie.Filtre.FirstAlbum {
			groupie.FiltreFirstAlbum()
		}
		if groupie.Filtre.Members {
			groupie.FiltreMembers()
		}
	}
	if r.Method == "POST" {
		buttonValue := r.FormValue("filtre")
		if buttonValue == "NameArtiste" {
			groupie.Filtre.NameArtiste = !groupie.Filtre.NameArtiste
			fmt.Println(groupie.Filtre.NameArtiste)
		}
		if buttonValue == "CreationDate" {
			groupie.Filtre.CreationDate = !groupie.Filtre.CreationDate
			fmt.Println(groupie.Filtre.CreationDate)
		}
		if buttonValue == "FirstAlbum" {
			groupie.Filtre.FirstAlbum = !groupie.Filtre.FirstAlbum
			fmt.Println(groupie.Filtre.FirstAlbum)
		}
		if buttonValue == "Members" {
			groupie.Filtre.Members = !groupie.Filtre.Members
			fmt.Println(groupie.Filtre.Members)
		}
	}
	tmpl.Execute(w, data)
}


func (e *Engine) SetActiveTab(tab string) {
	e.ActiveTab = tab
}

func (e Engine) GetActiveTab() string {
	return e.ActiveTab
}