package dev

import (
	"encoding/json"
	"log"
)

// CheckErr ...
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// PJSON ...
func PJSON(data ...interface{}) {
	en, err := json.Marshal(data)
	CheckErr(err)
	log.Println(string(en))
}

// PJSONIndent ...
func PJSONIndent(data ...interface{}) {
	en, err := json.MarshalIndent(data, "", "  ")
	CheckErr(err)
	log.Println(string(en))
}
