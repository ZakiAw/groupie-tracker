package main

import (
	"html/template"
	"net/http"
)

var temp, temp2 *template.Template

func main() {
	call()
	temp = template.Must(template.ParseGlob("tracker.html"))
	temp2 = template.Must(template.ParseGlob("artist.html"))
	fs := http.FileServer(http.Dir("style"))
	http.Handle("/style/", http.StripPrefix("/style", fs))
	http.HandleFunc("/", MainHandler)
	http.HandleFunc("/artist/", ArtistHandler)
	http.ListenAndServe(":8080", nil)

}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "tracker.html", Final)
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	temp2.ExecuteTemplate(w, "artist.html", Final)
}
