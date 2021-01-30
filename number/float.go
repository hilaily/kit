package number

import (
	"strconv"
)

// FloatPrecision float 小数点保留指定位数
func FloatPrecision(value float64, prec int) float64 {
	pp := strconv.FormatFloat(value, 'f', prec, 64)
	value, _ = strconv.ParseFloat(pp, 64)
	return value
}
