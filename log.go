package u

import (
	"encoding/json"
	"log"
	"testing"
)

func PringTypeValue(v interface{}) {
    log.Printf("%T: %#v", v)
}

func toJson(v interface{}) (string, error) {
	bs, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return "", err
	}

	return string(bs), nil
}


func J(v interface{}) string {
    s, err := toJson(v)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func LoggerJson(log *log.Logger) func(v interface{}) string {
    return func (v interface{}) string {
        s, err := toJson(v)
        if err != nil {
            log.Fatal(err)
        }
        return s
    }
}

func TestingJson(t *testing.T) func(v interface{}) string {
    return func (v interface{}) string {
        s, err := toJson(v)
        if err != nil {
            t.Fatal(err)
        }
        return s
    }
}

func PrintJson(v interface{}) {
    log.Println(J(v))
}

func LogJson(log *log.Logger, v interface{}) {
	log.Println(LoggerJson(log)(v))
}

func TlogJson(t *testing.T, v interface{}) {
	t.Log(TestingJson(t)(v))
}

func TfatalJson(t *testing.T, v interface{}) {
	t.Fatal(TestingJson(t)(v))
}
