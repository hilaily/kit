package helper

import "testing"

func TestPJSONIndent(t *testing.T) {
	m := map[string]interface{}{
		"a": 1,
		"b": "2",
		"c": map[string]interface{}{
			"c.1": "1",
		},
	}
	PJSON(m)
	PJSONIndent(m)

	PJSONIndent("test", m)
}
