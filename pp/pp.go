package pp

import (
	"encoding/json"
	"log"
	"strings"
)

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
