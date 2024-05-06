package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	AllData = []byte{}
	m map[string]json.RawMessage
	Data    Respect
	ApiLink = "https://groupietrackers.herokuapp.com/api"
	FinalData FinalStruct
	artistData []ArtistStruct
	locData []LocationStruct
	dateData []DataStruct
	relationData []RelationStruct
)


func call() {
	Getmosa(ApiLink, &Data)
	Getmosa(Data.ArtistUrl, &artistData)
	Getzaki(Data.LocationsUrl,m, &locData)
	Getzaki(Data.DatesUrl,m, &dateData)
	Getzaki(Data.RelationUrl,m, &relationData)
}

func Getmosa(Link string, Output interface{}) {
	linkData, err := http.Get(Link)
	if err != nil {
		fmt.Println("error getting link")
		return
	}
	body, err := io.ReadAll(linkData.Body)
	if err != nil {
		fmt.Println("error reading body", err)
		return
	}
	defer linkData.Body.Close()

	err = json.Unmarshal(body, Output)
	if err != nil {
		fmt.Println("error unmarshling", err)
		return
	}
}

func Getzaki(Link string, m map[string]json.RawMessage, Output interface{}) {
	linkData, err := http.Get(Link)
	if err != nil {
		fmt.Println("error getting link")
		return
	}
	body, err := io.ReadAll(linkData.Body)
	if err != nil {
		fmt.Println("error reading body", err)
		return
	}
	defer linkData.Body.Close()

	err = json.Unmarshal(body, &m)
	if err != nil {
		fmt.Println("error unmarshling", err)
		return
	}
	jk := unmarsh(m)
	err = json.Unmarshal(jk, &Output)
	if err != nil {
		fmt.Println("error unmarshling", err)
		return
	}
}

func unmarsh(m map[string]json.RawMessage) (a []byte){
	for _, raw := range m{
		for _, byte := range raw{
			a = append(a, byte)
		}
	} 
	return a
	}
	func collect()( []FinalStruct){
		call()
		allData := make([]FinalStruct,len(artistData))
		for i := 0 ; i < len(artistData) ; i++{
			allData[i].Artistf = artistData[i]
			allData[i].Locationf = locData[i]
			allData[i].Datef = dateData[i]
			allData[i].Relationf = relationData[i]
		}
		return allData
	}


