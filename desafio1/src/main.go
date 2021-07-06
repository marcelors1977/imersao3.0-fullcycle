package main

import (
    "html/template"
    "net/http"
)

var tmpl *template.Template

func mainpage(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w,"index.html",nil)
}

func init() {
	tmpl = template.Must(template.ParseGlob("./templates/index.html"))
}

func main() {	
	mux := http.NewServeMux()	
	mux.HandleFunc("/", mainpage)

	fileServer := http.FileServer(http.Dir("./gallery"))
	mux.Handle("/gallery/", http.StripPrefix("/gallery", fileServer))

	http.ListenAndServe( ":8000",mux)
}

