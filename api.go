package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	Data    Respect
	ApiLink = "https://groupietrackers.herokuapp.com/api"
	Final Fin
)


func call() {
	Getmosa(ApiLink, &Data)
	Getmosa(Data.ArtistUrl, &Final.Artistsf)
	Getmosa(Data.LocationsUrl, &Final.Locationsf)
	Getmosa(Data.DatesUrl, &Final.Datesf)
	Getmosa(Data.RelationUrl, &Final.Relf)
}

func Getmosa(Link string, Respect interface{}) {
	linkData, err := http.Get(Link)
	if err != nil {
		fmt.Println("><><><><>< Wala ma metzaker ><><><><>><><")
		return
	}
	body, err := io.ReadAll(linkData.Body)
	if err != nil {
		fmt.Println("><><><><>< walaaa ><><><><>", err)
		return
	}
	defer linkData.Body.Close()

	err = json.Unmarshal(body, Respect)
	if err != nil {
		fmt.Println("fehadagay", err)
		return
	}
}
