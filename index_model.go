package main

import (
	"golang.org/x/text/language"
)

type IndexModelMap map[language.Tag]IndexModel

var messages IndexModelMap

type IndexModel struct {
	Titel   string
	Untertitel string
	Nav     NavModel
	Spaziergaenge FuehrungSpaziergaenge
}


type FuehrungSpaziergaenge struct {
	Titel string
	Kurzbeschreibung string
}

func InitIndexModel() {
	messages = map[language.Tag]IndexModel{
		language.German: {
			Titel: "Fremdenführer Kateryna Pacher",
			Untertitel: "Österreich entdecken",
			Nav: NavModel{
				Touren:  "Touren",
				Uebermich: "Über mich",
				Sprache: `<option value="de" selected>Deutsch</option>
                    <option value="ru">Русский</option>
                    <option value="ukr">Українська</option>`,
			},
			Spaziergaenge: FuehrungSpaziergaenge {
				Titel: "Spaziergänge",
				Kurzbeschreibung: "Wien ist klein und die Wege sind kurz. Alleine im Zentrum Wiens gibt es unglaublich viel zum Entdecken",
			},
		},
		language.Russian: {
			Titel: "исследовать",
			Untertitel: "untertitel russ",
			Nav: NavModel{
				Touren:  "tyr",
				Uebermich: "Обо мне",
				Sprache: `<option value="de">Deutsch</option>
                    <option value="ru" selected>Русский</option>
                    <option value="ukr">Українська</option>`,
			},
			Spaziergaenge: FuehrungSpaziergaenge {
				Titel: "Прогулки пешком",
				Kurzbeschreibung: "Лорем ипсум долор сит амет, еу вис менандри медиоцрем улламцорпер. Иус ех фацер новум, ат лорем постулант урбанитас меи. Ет ест меис тинцидунт.",
			},
		},
		language.Ukrainian: {
			Titel: "titel ukrainisch",
			Untertitel: "untertitel ukr",
			Nav: NavModel{
				Touren:  "tyr ukr",
				Uebermich: "Про мене",
				Sprache: `<option value="de">Deutsch</option>
                    <option value="ru">Русский</option>
                    <option value="ukr" selected>Українська</option>`,
			},
			Spaziergaenge: FuehrungSpaziergaenge {
				Titel: "Піші прогулянки",
				Kurzbeschreibung: "Wien ist klein und die Wege sind kurz. Alleine im Zentrum Wiens gibt es unglaublich viel zum Entdecken",
			},
		},
	}
}

func LookUpIndexModel(tag language.Tag) IndexModel {
	return messages[tag]
}
