package netx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPPortIsOpen(t *testing.T) {
	open, err := TCPPortIsOpen("www.baidu.com:80", 1)
	assert.NoError(t, err)
	assert.True(t, open)

	open, err = TCPPortIsOpen("127.0.0.1:8", 1)
	assert.False(t, open)
	t.Logf("err: %s\n", err.Error())
}
