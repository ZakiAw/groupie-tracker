package main

type Respect struct {
	ArtistUrl    string `json:"artists"`
	LocationsUrl string `json:"locations"`
	DatesUrl     string `json:"dates"`
	RelationUrl  string `json:"relation"`
}
type ArtistStruct struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}
type LocationStruct struct {
		Locations []string `json:"locations"`
}
type DataStruct struct {
	Dates []string `json:"dates"`
}
type RelationStruct struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}
type FinalStruct struct {
	Artistf   ArtistStruct
	Locationf LocationStruct
	Datef     DataStruct
	Relationf RelationStruct
}
