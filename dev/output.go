package dev

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// Deprecated: use helper package
// CheckErr ...
func CheckErr(err error, msg ...string) {
	if err != nil {
		if len(msg) > 0 {
			m := strings.Join(msg, ";")
			err = fmt.Errorf("%w, %s", err, m)
		}
		panic(err)
	}
}

// Deprecated: use helper package
// PJSON ...
func PJSON(data ...interface{}) {
	en, err := json.Marshal(data)
	CheckErr(err)
	log.Println(string(en))
}

// Deprecated: use helper package
// PJSONIndent ...
func PJSONIndent(data ...interface{}) {
	en, err := json.MarshalIndent(data, "", "  ")
	CheckErr(err)
	log.Println(string(en))
}

// Deprecated: use helper package
// Dump ...
func Dump(i ...interface{}) {
	str := strings.Builder{}
	for _, v := range i {
		en, err := json.Marshal(v)
		if err != nil {
			log.Panicln(err)
		}
		str.Write(en)
		str.WriteByte('\n')
	}
	log.Println(str.String())
}

// Deprecated: use helper package
// IDump ...
func IDump(i ...interface{}) {
	str := strings.Builder{}
	for _, v := range i {
		en, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			log.Panicln(err)
		}
		str.Write(en)
		str.WriteByte('\n')
	}
	log.Println(str.String())
}
