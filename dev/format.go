package dev

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

// Deprecated: use helper package
func IsGoTest() bool {
	return strings.HasSuffix(os.Args[0], ".test")
}

// Deprecated: use helper package
// JSON ...
func JSON(i interface{}) []byte {
	en, _ := json.Marshal(i)
	return en
}

// Deprecated: use helper package
func JSONIndent(i interface{}) []byte {
	en, _ := json.MarshalIndent(i, "", "  ")
	return en
}

// Deprecated: use helper package
func ReadJSON(filename string, ptr interface{}) {
	res, err := ioutil.ReadFile(filename)
	CheckErr(err)
	err = json.Unmarshal(res, ptr)
	CheckErr(err)
}

// Deprecated: use helper package
func WriteJSON(filename string, data interface{}) {
	en, err := json.Marshal(data)
	CheckErr(err)
	err = ioutil.WriteFile(filename, en, 0777)
	CheckErr(err)
}
