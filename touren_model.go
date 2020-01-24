package main

import (
	"golang.org/x/text/language"
)


var tourenModelMap map[language.Tag]TourenModel

type TourenModel struct {

	Titel   string
	Untertitel string
	Nav     NavModel
	Spaziergaenge FuehrungSpaziergaenge
}

func lookUpTourenModel(tag language.Tag) TourenModel {
	if tourenModelMap == nil {
		tourenModelMap = initModelMap()
	}
	return tourenModelMap[tag]
}

func initModelMap() map[language.Tag]TourenModel {
	tourenModelMap = map[language.Tag]TourenModel {
		language.German : {
			Titel:"Titel DE",
			Untertitel:"Untertitel DE",
			Nav: NavModel{
				Touren:  "Touren",
				Uebermich: "Über mich",
				Sprache: `<option value="de" selected>Deutsch</option>
                    <option value="ru">Русский</option>
                    <option value="ukr">Українська</option>`,
			},
		},
		language.Russian : {
			Titel:"Titel Ru",
			Untertitel:"Untertitel Ru",
			Nav: NavModel{
				Touren:  "tyr",
				Uebermich: "Обо мне",
				Sprache: `<option value="de">Deutsch</option>
                    <option value="ru" selected>Русский</option>
                    <option value="ukr">Українська</option>`,
			},
		},
		language.Ukrainian: {
			Titel:      "titel ukrainisch",
			Untertitel: "untertitel ukr",
			Nav: NavModel{
				Touren:    "tyr ukr",
				Uebermich: "Про мене",
				Sprache: `<option value="de">Deutsch</option>
                    <option value="ru">Русский</option>
                    <option value="ukr" selected>Українська</option>`,
			},
		},
	}
	return tourenModelMap
}
