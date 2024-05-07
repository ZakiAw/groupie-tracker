package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var temp, temp2 *template.Template

func main() {
	temp = template.Must(template.ParseFiles("tracker.html"))
	temp2 = template.Must(template.ParseFiles("artist.html"))
	fs := http.FileServer(http.Dir("style"))
	http.Handle("/style/", http.StripPrefix("/style", fs))

	http.HandleFunc("/", MainHandler)
	// http.HandleFunc("/artist/", ArtistHandler)
	http.ListenAndServe(":8080", nil)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	all := collect()
	id := r.FormValue("ID")
	temp.ExecuteTemplate(w, "tracker.html", all)
	fmt.Println(id)
}

//	func ArtistHandler(w http.ResponseWriter, r *http.Request) {
//		all := collect()
//		temp2.ExecuteTemplate(w, "artist.html", all)
//	}
// func main() {
// 	var a1 []FinalStruct
// 	id := 5
// 	all := collect()
// 	for _, i := range all {
// 		if id == i.Artistf.ID {
// 			a1 = append(a1, i)
// 		}
// 	}
// 	fmt.Println(a1)
// }
