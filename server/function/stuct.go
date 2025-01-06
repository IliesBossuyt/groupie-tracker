package engine

// Je crée ma structure
type Engine struct {
	Mot       string
	ID        int
	ActiveTab string
	Title     string

	ArtistStruct struct {
		ID           []int
		Image        []string
		Name         []string
		Members      map[int][]string
		CreationDate []int
		FirstAlbum   []string
		Location     map[int][]string
		Date         map[int][]string
		Relation     map[int][]string
	}

	LocationStruct struct {
		ID       []int
		Location map[int][]string
	}

	DateStruct struct {
		ID    []int
		Dates map[int][]string
	}

	RelationStruct struct {
		ID             []int
		DatesLocations map[int][]string
	}

	Filtre struct {
		NameArtiste  bool
		CreationDate bool
		FirstAlbum   bool
		Members      bool
	}
}
