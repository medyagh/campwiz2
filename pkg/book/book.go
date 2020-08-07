package book

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v1"
)

// - key: VILLAGE CAMPER INN RV - Crescent City
//   name: VILLAGE CAMPER INN RV PARK
//   srating: 7
//   desc: Woods and water, that’s what attracts visitors to California’s north coast.
//     Village Camper Inn provides nearby access to big woods and big water. This RV
//     park is on 20 acres of wooded land, about a 10-minute drive away from the giant
//     redwoods along U.S. 199. In addition, you’ll find some ...
// 	 locale: in Crescent Cit
type Entry struct {
	Key     string `yaml:"key"`
	Name    string `yaml:"name"`
	SRating int    `yaml:"srating"`
	Desc    string `yaml:"desc"`
	Locale  string `yaml:"locale"`
}

type Entries struct {
	Entries []Entry
}

var LoadedEntries = make(map[string]Entry)

func Load() error {
	log.Printf("Loading books ...")
	var flatEntries Entries
	f, err := ioutil.ReadFile("book.yaml")
	if err != nil {
		log.Printf("failed to read book file: %v", err)
		return err
	}
	err = yaml.Unmarshal(f, &flatEntries)
	if err != nil {
		log.Printf("failed to unmarshal book: %v", err)
		return err
	}
	log.Printf("Loaded %d entries from book.yaml ...", len(flatEntries.Entries))
	for _, m := range flatEntries.Entries {
		if val, ok := LoadedEntries[m.Key]; ok {
			log.Printf("already loaded. Previous=%+v, New=%+v", val, m)
			continue
		}
		LoadedEntries[m.Key] = m
	}

	log.Printf("Length of loaded books %d", len(LoadedEntries))

	return nil
}
