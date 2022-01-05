package number

import (
	"strconv"
)

// FormatFloat 小数点保留指定位数
func FormatFloat(value float64, decimal int) float64 {
	value, _ = EFormatFloat(value, decimal)
	return value
}

func EFormatFloat(value float64, decimal int) (float64, error) {
	pp := strconv.FormatFloat(value, 'f', decimal, 64)
	return strconv.ParseFloat(pp, 64)
}
