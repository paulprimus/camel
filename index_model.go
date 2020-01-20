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
	Spaziergaenge FuehrungSpaziergaenge
}

type NavModel struct {
	Touren,
	Uebermich string
}

type FuehrungSpaziergaenge struct {
	Titel string
	Kurzbeschreibung string
}

func InitIndexModel() {
	messages = map[language.Tag]IndexModel{
		language.German: {
			Sprache: `<option value="de" selected>Deutsch</option>
                    <option value="ru">Русский</option>
                    <option value="ukr">Українська</option>`,
			Titel: "Fremdenführer Kateryna Pacher",
			Untertitel: "Österreich entdecken",
			Nav: NavModel{
				Touren:  "Touren",
				Uebermich: "Über mich",
			},
			Spaziergaenge: FuehrungSpaziergaenge {
				Titel: "Spaziergänge",
				Kurzbeschreibung: "Wien ist klein und die Wege sind kurz. Alleine im Zentrum Wiens gibt es unglaublich viel zum Entdecken",
			},
		},
		language.Russian: {
			Sprache: `<option value="de">Deutsch</option>
                    <option value="ru" selected>Русский</option>
                    <option value="ukr">Українська</option>`,
			Titel: "исследовать",
			Untertitel: "untertitel russ",
			Nav: NavModel{
				Touren:  "tyr",
				Uebermich: "Обо мне",
			},
			Spaziergaenge: FuehrungSpaziergaenge {
				Titel: "Прогулки пешком",
				Kurzbeschreibung: "Лорем ипсум долор сит амет, еу вис менандри медиоцрем улламцорпер. Иус ех фацер новум, ат лорем постулант урбанитас меи. Ет ест меис тинцидунт.",
			},
		},
		language.Ukrainian: {
			Sprache: `<option value="de">Deutsch</option>
                    <option value="ru">Русский</option>
                    <option value="ukr" selected>Українська</option>`,
			Titel: "titel ukrainisch",
			Untertitel: "untertitel ukr",
			Nav: NavModel{
				Touren:  "tyr ukr",
				Uebermich: "Про мене",
			},
			Spaziergaenge: FuehrungSpaziergaenge {
				Titel: "Піші прогулянки",
				Kurzbeschreibung: "Wien ist klein und die Wege sind kurz. Alleine im Zentrum Wiens gibt es unglaublich viel zum Entdecken",
			},
		},
	}
}

func LookUpModel(tag language.Tag) IndexModel {
	return messages[tag]
}
