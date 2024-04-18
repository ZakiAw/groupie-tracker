package main

import (
	"html/template"
	"net/http"
)
var temp *template.Template


func main() {
	temp = template.Must(template.ParseGlob("tracker.html"))
	http.HandleFunc("/", MainHandler)
	http.ListenAndServe(":8080", nil)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "tracker.html",nil)
}
