package main

import (
	//"encoding/json"
	"golang.org/x/text/language"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

var tmpl *template.Template

func main() {
	var logfile *os.File
	var err error
	if logfile, err = os.OpenFile("camel.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644); err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()
	logout := io.MultiWriter(os.Stdout, logfile)
	log.SetOutput(logout)
	log.Println("WebServer startet")

	InitIndexModel()
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/main", handler)
	http.HandleFunc("/main/language", languageHandler)
	if err = http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("loading index")
	locateVisitor(r)
	model := LookUpModel(language.German)
	tmpl = template.Must(template.ParseFiles("templates/index.html"))
	if err := tmpl.Execute(w, model); err != nil {
		log.Fatal(err)
	}
}

func locateVisitor(r *http.Request) {
	log.Println(r.Header)
}

func languageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("languageHandler")
	if r.Method == http.MethodPost {
		var err error
		var sprache language.Tag
		if sprache, err = language.Parse(r.FormValue("sprache")); err != nil {
			log.Println("Sprache konnte nicht ermittelt werden!")
		}
		log.Println("Sprache: ", sprache)
		model := LookUpModel(sprache)
		if err = tmpl.ExecuteTemplate(w, "index.html", model); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Falsche HTTP Methode", http.StatusMethodNotAllowed)
	}
	log.Println("Done!")
}
