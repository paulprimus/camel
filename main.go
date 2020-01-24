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
var activeLanguage language.Tag = language.German

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
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/language", languageHandler)
	http.HandleFunc("/touren", tourenHandler)
	if err = http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("loading index")
	lang := locateVisitor(r)
	model := LookUpIndexModel(lang)
	tmpl = template.Must(template.ParseFiles("templates/index.html", "templates/touren.html"))
	if err := tmpl.ExecuteTemplate(w,"index.html", model); err != nil {
		log.Fatal(err)
	}
}

func locateVisitor(r *http.Request) language.Tag {
	return language.German
}

func languageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("languageHandler")
	if r.Method == http.MethodPost {
		var sprache string = r.FormValue("sprache")
		var page string = r.FormValue("template")
		var err error
		if activeLanguage, err = language.Parse(sprache); err != nil {
			log.Println("Sprache konnte nicht ermittelt werden!")
		}
		log.Println(activeLanguage)
		model := LookUpIndexModel(activeLanguage)

		if err = tmpl.ExecuteTemplate(w, page, model); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Falsche HTTP Methode", http.StatusMethodNotAllowed)
	}
	log.Println("Done!")
}

func tourenHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("loading touren")
	tm := lookUpTourenModel(activeLanguage)
	if err := tmpl.ExecuteTemplate(w, "touren.html", tm); err != nil {
		log.Println("Fehler beim parsen vom Template touren.html", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
