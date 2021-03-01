package u

import (
	"encoding/json"
	"log"
)

func PrintJson(v interface{}) {
	bs, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bs))
}

func LogJson(log *log.Logger, v interface{}) {
	bs, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bs))
}
