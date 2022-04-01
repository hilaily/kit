package dev

import (
	"encoding/json"
	"io/ioutil"
)

// JSON ...
func JSON(i interface{}) []byte {
	en, _ := json.Marshal(i)
	return en
}

func JSONIndent(i interface{}) []byte {
	en, _ := json.MarshalIndent(i, "", "  ")
	return en
}

func ReadJSON(filename string, ptr interface{}) {
	res, err := ioutil.ReadFile(filename)
	CheckErr(err)
	err = json.Unmarshal(res, ptr)
	CheckErr(err)
}
