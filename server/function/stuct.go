package engine

// Ce qu'on va utiliser dans les fonctions
type Engine struct {
	Api              []ApiStruct
	Artists          []ArtistStruct
	Filtre           FiltreStruct
	GestionFiltre    GestionFiltreStruct
	Recherche        Recherche
	Relations        Relation
	RelationEntry    RelationEntry
	PageArtistStruct PageArtistStruct
	RandomImages     []RandomImage
}

// Structure pour la génération d'image aléatoire
type RandomImage struct {
	ID    int
	Image string
}

// Mot qu'on récupere depuis le front
type Recherche struct {
	Mot string
}

// Structure de ce qu'on récupere depuis l'api
type ApiStruct struct {
	ID             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Structure pour la relation qu'on récupere depuis l'api
type Relation struct {
	Index []RelationEntry `json:"index"`
}

type RelationEntry struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Structure pour la page artiste
type PageArtistStruct struct {
	ID             int
	Image          string
	Name           string
	Members        []string
	CreationDate   int
	FirstAlbum     string
	DatesLocations map[string][]string
}

// Filtre et envoie au front
type ArtistStruct struct {
	ID             int
	Image          string
	Name           string
	Members        []string
	CreationDate   int
	FirstAlbum     string
	DatesLocations map[string][]string
}

// Filtre true or false
type FiltreStruct struct {
	NameArtisteAZ bool
	NameArtisteZA bool
	CreationDate  bool
	FirstAlbum    bool
	Members       bool
}

// Gestion des filtres
type GestionFiltreStruct struct {
	NombreMembres   int
	minCreationDate int
	maxCreationDate int
	minFirstAlbum   int
	maxFirstAlbum   int
}

// Structure pour la géolocalisation
type Location struct {
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	DisplayName string `json:"display_name"`
}