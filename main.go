package main

import (
	"html/template"
	"net/http"
)

var temp *template.Template

func main() {
	call()
	temp = template.Must(template.ParseGlob("tracker.html"))
	fs := http.FileServer(http.Dir("style"))
	http.Handle("/style/", http.StripPrefix("/style", fs))
	http.HandleFunc("/", MainHandler)
	http.ListenAndServe(":8080", nil)
	
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "tracker.html", Data)
}
