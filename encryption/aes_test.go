package encryption

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	key := []byte("abcdefghijklmnaofdalfaadfjalfjaf")
	text := "hello world"
	cipher, err := AesEncrypt(text, key)
	t.Logf("%s", cipher)
	assert.NoError(t, err)
	newText, err := AesDecrypt(cipher, key)
	assert.NoError(t, err)
	assert.Equal(t, text, newText)
}
