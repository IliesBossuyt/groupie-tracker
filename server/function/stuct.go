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
		Members      []string
		CreationDate []int
		FirstAlbum   []string
		Location     []string
		Date         []string
	}

	LocationStruct struct {
		ID       []int
		Location []string
	}

	DateStruct struct {
		ID    []int
		Dates []string
	}
	
}
