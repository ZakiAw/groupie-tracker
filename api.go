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
)


func call() {
	Getmosa(ApiLink, &Data)
	Getmosa(Data.ArtistUrl, &Data.Artist)
	Getmosa(Data.LocationsUrl, &Data.Location)
	Getmosa(Data.DatesUrl, &Data.Date)
	Getmosa(Data.RelationUrl, &Data.Relation)
}

func Getmosa(Link string, Respect interface{}) {
	resp, err := http.Get(Link)
	if err != nil {
		fmt.Println("><><><><>< No Response ><><><><>><><")
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("artsfdhgj", err)
		return
	}
	defer resp.Body.Close()

	err = json.Unmarshal(body, Respect)
	if err != nil {
		fmt.Println("errr", err)
		return
	}
}
