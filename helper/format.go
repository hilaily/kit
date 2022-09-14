package helper

import (
	"encoding/json"
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
	res, err := os.ReadFile(filename)
	CheckErr(err)
	err = json.Unmarshal(res, ptr)
	CheckErr(err)
}

func WriteJSON(filename string, data interface{}) {
	en, err := json.Marshal(data)
	CheckErr(err)
	err = os.WriteFile(filename, en, 0777)
	CheckErr(err)
}

func AppendFile(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.Write(data); err != nil {
		return err
	}
	return nil
}
