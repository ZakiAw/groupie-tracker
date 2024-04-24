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
	Getmosa(ApiLink)
}

func Getmosa(Link string) {
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


	err = json.Unmarshal(body, &art)
	if err != nil {
		fmt.Println("errr", err)
		return
	}
}
