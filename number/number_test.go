package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloatPrecision(t *testing.T) {
	data := map[float64]float64{
		0.12131: 0.121,
		1.23456: 1.235,
		0.99999: 1,
		0.9995:  1,
		0.99895: 0.999,
	}
	for k, v := range data {
		res := FloatPrecision(k, 3)
		assert.Equal(t, v, res)
		t.Log(res)
	}
}
