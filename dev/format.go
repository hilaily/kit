package dev

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

func IsGoTest() bool {
	return strings.HasSuffix(os.Args[0], ".test")
}

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

func WriteJSON(filename string, data interface{}) {
	en, err := json.Marshal(data)
	CheckErr(err)
	err = ioutil.WriteFile(filename, en, 0777)
	CheckErr(err)
}
