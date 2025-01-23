package engine

import (
	"math/rand"
	"strconv"
	"strings"
)

// Image aléatoire pour le carrousel
func (groupie *Engine) RandomImg() {
	// Créer une liste de la taille de groupie.Api
	image := make([]int, len(groupie.Api))
	for i := range image {
		image[i] = i
	}

	// Mélanger les image
	rand.Shuffle(len(image), func(i, j int) {
		image[i], image[j] = image[j], image[i]
	})

	// Prendre les 20 premieres image mélangés
	var randomImages []RandomImage
	for i := 0; i < len(image) && i < 20; i++ {
		randomIndex := image[i]
		randomArtist := groupie.Api[randomIndex]
		randomImage := RandomImage{
			ID:    randomArtist.ID,
			Image: randomArtist.Image,
		}
		randomImages = append(randomImages, randomImage)
	}

	groupie.RandomImages = randomImages
}

// Reinitialiser la structure Artists
func (groupie *Engine) Reset() {
	groupie.Artists = nil
}

// Récuperer les données de l'api depuis la structure Api
func (groupie *Engine) Clone() {
	for _, add := range groupie.Api {
		artist := ArtistStruct(add)
		groupie.Artists = append(groupie.Artists, artist)
	}
}

// Fonction pour savoir si un slice contient un prefix
func containsPrefix(slice []string, prefix string) bool {
	for _, s := range slice {
		if strings.Contains(strings.ToLower(s), strings.ToLower(prefix)) {
			return true
		}
	}
	return false
}

// Remplacer les tirets et underscores par des espaces
func Underscore(datesLocations map[string][]string) map[string][]string {
	updatedDatesLocations := make(map[string][]string)

	for key, value := range datesLocations {
		location := strings.ReplaceAll(key, "-", " ")
		location = strings.ReplaceAll(location, "_", " ")

		// Mettre la première lettre de chaque mot en majuscule
		location = Maj(location)

		updatedDatesLocations[location] = value
	}

	return updatedDatesLocations
}

// Mettre en majuscule la première lettre de chaque mot
func Maj(input string) string {
	mot := strings.Fields(input) // Divise la chaîne en mots
	for i, word := range mot {
		if len(word) > 0 {
			mot[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(mot, " ")
}

// Filtre tri alphabétique
func (groupie *Engine) FiltreMenu(menu string) {
	if menu == "az" {
		groupie.Filtre.NameArtisteAZ = true
		groupie.Filtre.NameArtisteZA = false
	}
	if menu == "za" {
		groupie.Filtre.NameArtisteZA = true
		groupie.Filtre.NameArtisteAZ = false
	}
	if menu == "desactiver" {
		groupie.Filtre.NameArtisteAZ = false
		groupie.Filtre.NameArtisteZA = false
	}
}

// Filtre par date de creation
func (groupie *Engine) FiltreCreationDate(minCreationDate string, maxCreationDate string) {
	minIntCreationDate, _ := strconv.Atoi(minCreationDate)
	maxIntCreationDate, _ := strconv.Atoi(maxCreationDate)
	groupie.GestionFiltre.minCreationDate = minIntCreationDate
	groupie.GestionFiltre.maxCreationDate = maxIntCreationDate

	if minCreationDate == "" {
		groupie.GestionFiltre.minCreationDate = 0
	}
	if maxCreationDate == "" {
		groupie.GestionFiltre.maxCreationDate = 3000
	}

	if minCreationDate != "" || maxCreationDate != "" {
		groupie.Filtre.CreationDate = true
	} else {
		groupie.Filtre.CreationDate = false
	}
}

// Filtre par FirstAlbum
func (groupie *Engine) FiltreFirstAlbum(minFirstAlbum string, maxFirstAlbum string) {
	minIntFirstAlbum, _ := strconv.Atoi(minFirstAlbum)
	maxIntFirstAlbum, _ := strconv.Atoi(maxFirstAlbum)
	groupie.GestionFiltre.minFirstAlbum = minIntFirstAlbum
	groupie.GestionFiltre.maxFirstAlbum = maxIntFirstAlbum

	if minFirstAlbum == "" {
		groupie.GestionFiltre.minFirstAlbum = 0
	}
	if maxFirstAlbum == "" {
		groupie.GestionFiltre.maxFirstAlbum = 3000
	}
	if minFirstAlbum != "" || maxFirstAlbum != "" {
		groupie.Filtre.FirstAlbum = true
	} else {
		groupie.Filtre.FirstAlbum = false
	}
}

// Filtre par nombre de membres
func (groupie *Engine) FiltreMembers(membersString string) {
	members, _ := strconv.Atoi(membersString)
	groupie.GestionFiltre.NombreMembres = members
	if members > 0 {
		groupie.Filtre.Members = true
	} else {
		groupie.Filtre.Members = false
	}
}

// Réinitialisation des filtres
func (groupie *Engine) ResetFiltre() {
	groupie.Filtre.NameArtisteAZ = false
	groupie.Filtre.NameArtisteZA = false
	groupie.GestionFiltre.minCreationDate = 0
	groupie.GestionFiltre.maxCreationDate = 3000
	groupie.GestionFiltre.minFirstAlbum = 0
	groupie.GestionFiltre.maxFirstAlbum = 3000
	groupie.Filtre.Members = false
	groupie.Reset()
	groupie.Clone()
}
