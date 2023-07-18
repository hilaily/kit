package helper

// revive:disable:deep-exit

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

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

// Dump ...
func Dump(i ...interface{}) {
	str := strings.Builder{}
	for _, v := range i {
		en, err := json.Marshal(v)
		if err != nil {
			log.Panicln(err)
		}
		_, _ = str.Write(en)
		_ = str.WriteByte('\n')
	}
	log.Println(str.String())
}

// IDump ...
func IDump(i ...interface{}) {
	str := strings.Builder{}
	for _, v := range i {
		en, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			log.Panicln(err)
		}
		_, _ = str.Write(en)
		_ = str.WriteByte('\n')
	}
	log.Println(str.String())
}
