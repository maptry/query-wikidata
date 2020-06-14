package main

/*
instance of (P31)
subclass of (P279) AUTHORITY CONTROL FOR PLACES (Q19829908)
nuts code (P605)
lau (P782)

contains administrative territorial entity (P150)
European Union (Q458)

de German municipality key (P439)
fr INSEE municipality code (P374)
es INE municipality code (P772)
it ISTAT ID (P635)
po INE ID (Portugal) (P6324)
ne CBS municipality code (P382)
be NIS/INS code (P1567)
lu LAU (P782)
se
fi
dk
ie
pl
hu
cz
so
sl
cr
he
bu
ro
li
el
le

Bundesland
Regierungsbezirk
Landkreis
Gemeinde
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/knakk/sparql"
)

type Municipality struct {
	Name            string
	Population      string
	Area            string
	Elevation       string
	OSMRelationID   string
	MunicipalityKey string
	URI             string
	Coordinates     string
	NUTS            string
}

func main() {
	repo, err := sparql.NewRepo("https://query.wikidata.org/sparql",
		sparql.Timeout(time.Millisecond*50000),
	)
	if err != nil {
		log.Fatal(err)
	}
	res, err := repo.Query(`
SELECT DISTINCT ?municipality ?municipalityLabel ?germmunikey ?population ?coords ?area ?elevation ?osmrelationid ?nuts
WHERE {
  ?municipality wdt:P439 ?germmunikey;
		wdt:P625 ?coords;
                wdt:P1082 ?population;
		wdt:P2046 ?area;
  OPTIONAL { ?municipality wdt:P2044 ?elevation }
  OPTIONAL { ?municipality wdt:P402 ?osmrelationid }
  OPTIONAL { ?municipality wdt:P605 ?nuts }.
  SERVICE wikibase:label { bd:serviceParam wikibase:language "de". }
}
`)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("% v", res.Results.Bindings)
	places := make([]Municipality, 0, len(res.Results.Bindings))
	for _, m := range res.Results.Bindings {
		places = append(places, Municipality{m["municipalityLabel"].Value, m["population"].Value, m["area"].Value, m["elevation"].Value, m["osmrelationid"].Value, m["germmunikey"].Value, m["municipality"].Value, m["coords"].Value, m["nuts"].Value})
	}
	b, err := json.MarshalIndent(places, "", "    ")
	if err != nil {
		log.Fatalf("error while encoding: %s", err.Error())
	}
	fmt.Printf("%s", b)
}
