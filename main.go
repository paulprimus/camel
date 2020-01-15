package main

import (
	"encoding/json"
	"golang.org/x/text/language"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

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

	InitMessages()
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
	model := LookUp(language.German)
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, model)
	if err != nil {
		log.Fatal(err)
	}
}

type test struct {
	Sprache string //`json:'sprache'`
}

func languageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("languageHandler")
	var err error
	if r.Method == http.MethodPost {
		var answer []byte
		sprache := &test{}
		if err = json.NewDecoder(r.Body).Decode(sprache); err != nil {
			log.Fatal("Fehler: JSON konnte nicht gelesen werden! {}", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if answer, err = json.Marshal(sprache); err != nil {
			log.Fatal("Fehler: Antwort konnte nicht erzeugt werden! {}", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(answer)
	} else {
		http.Error(w, "Falsche HTTP Methode", http.StatusInternalServerError)
	}
}
