package main

import "golang.org/x/text/language"

type I18nMap map[language.Tag]IndexModel

var messages I18nMap

type IndexModel struct {
	Titel string
	Nav   NavModel
}

type NavModel struct {
	Touren,
	Kontakt string
}

func InitIndexModel() {
	messages = map[language.Tag]IndexModel{
		language.German: IndexModel{
			Titel: "Touren",
			Nav: NavModel{
				Touren:  "asdfsadf",
				Kontakt: "konkat",
			},
		},
		language.Russian: IndexModel{
			Titel: "russsisch",
			Nav: NavModel{
				Touren:  "tyr",
				Kontakt: "konmakt",
			},
		},
	}
}

func LookUpModel(tag language.Tag) IndexModel {
	return messages[tag]
}
