package main

type Respect struct {
	ArtistUrl    string `json:"artists"`
	LocationsUrl string `json:"locations"`
	DatesUrl     string `json:"dates"`
	RelationUrl  string `json:"relation"`
}
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}
type Loc struct {
	Locations []string `json:"locations"`
}
type Da struct {
	Dates []string `json:"dates"`
}
type Rel struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}
type Fin struct {
	Artistsf   []Artist
	Locationsf Loc
	Datesf     Da
	Relf       Rel
}
