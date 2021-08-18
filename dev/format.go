package dev

import "encoding/json"

// JSON ...
func JSON(i interface{}) []byte {
	en, _ := json.Marshal(i)
	return en
}

func JSONIndent(i interface{}) []byte {
	en, _ := json.MarshalIndent(i, "", "  ")
	return en
}
