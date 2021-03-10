package u

import (
	"encoding/json"
	"log"
	"testing"
)


func PringTypeValue(v interface{}) {
    log.Printf("%T: %#v", v)
}

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

func TlogJson(t *testing.T, v interface{}) {
	bs, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(bs))
}

func TfatalJson(t *testing.T, v interface{}) {
	bs, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	t.Fatal(string(bs))
}
