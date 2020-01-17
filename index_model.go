package main

import (
	"golang.org/x/text/language"
	"html/template"
)

type I18nMap map[language.Tag]IndexModel

var messages I18nMap

type IndexModel struct {
	Sprache template.HTML
	Titel   string
	Untertitel string
	Nav     NavModel
}

type NavModel struct {
	Touren,
	Kontakt string
}

func InitIndexModel() {
	messages = map[language.Tag]IndexModel{
		language.German: {
			Sprache: `<option value="de" selected>Deutsch</option>
                    <option value="ru">Russisch</option>
                    <option value="ukr">Ukrainisch</option>`,
			Titel: "Titel bla bla",
			Untertitel: "Gemeinsam Österreich entdecken",
			Nav: NavModel{
				Touren:  "Touren",
				Kontakt: "konkat",
			},
		},
		language.Russian: {
			Sprache: `<option value="de">Deutsch</option>
                    <option value="ru" selected>Russisch</option>
                    <option value="ukr">Ukrainisch</option>`,
			Titel: "исследовать",
			Untertitel: "untertitel russ",
			Nav: NavModel{
				Touren:  "tyr",
				Kontakt: "konmakt",
			},
		},
		language.Ukrainian: {
			Sprache: `<option value="de">Deutsch</option>
                    <option value="ru">Russisch</option>
                    <option value="ukr" selected>Ukrainisch</option>`,
			Titel: "titel ukrainisch",
			Untertitel: "untertitel ukr",
			Nav: NavModel{
				Touren:  "tyr ukr",
				Kontakt: "konmakt ukr",
			},
		},
	}
}

func LookUpModel(tag language.Tag) IndexModel {
	return messages[tag]
}
