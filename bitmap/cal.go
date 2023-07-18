package bitmap

import "bytes"

// GetMask ...
func GetMask() []uint64 {
	r := make([]uint64, bitSize)
	copy(r, bitmask)
	return r
}

// BinaryString
func BinaryString(val uint64) string {
	return uint64ToBinaryString(val)
}

// Uint64ToByteArr ...
func Uint64ToByteArr(data uint64) []byte {
	sb := &bytes.Buffer{}
	sb.Grow(64)
	for index := 0; index < bitSize; index++ {
		if (bitmask[bitSize-1-index] & data) == 0 {
			_ = sb.WriteByte('0')
		} else {
			_ = sb.WriteByte('1')
		}
	}
	return sb.Bytes()
}

// Set 置 1， 从 0 开始
func Set(val uint64, position int) uint64 {
	return val | bitmask[position]
}

// Unset 置 0， 从 0 开始
func Unset(val uint64, position int) uint64 {
	return val & (^bitmask[position])
}

// IsSet ...
func IsSet(val uint64, position int) bool {
	return val&bitmask[position] == 1
}

// NumberOf1 计算 1 的个数
func NumberOf1(val uint64) int {
	count := 0
	for val != 0 {
		count++
		val = val & (val - 1)
	}
	return count
}

// Nagation 取反
func Nagation(val uint64) uint64 {
	return ^val
}
