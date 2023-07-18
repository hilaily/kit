package bitmap

import (
	"fmt"
	"strings"
)

const (
	bitSize = 64
)

var bitmask = make([]uint64, bitSize)

func init() {
	for i := range bitmask {
		bitmask[i] = 1 << i
	}
}

// 创建工厂函数
func New(maxCount uint64) *BitMap {
	return &BitMap{bits: make([]uint64, (maxCount+bitSize-1)/bitSize), bitCount: 0, maxCount: maxCount}
}

// NewFromBitMap ...
func NewFromBitMap(data []uint64, maxCount uint64) *BitMap {
	bitCount := 0
	copyData := make([]uint64, 0, len(data)) // 复制一份，一面传入的值被外部无意识修改
	for i, v := range data {
		bitCount += NumberOf1(v)
		copyData[i] = v
	}
	return &BitMap{bits: copyData, bitCount: uint64(bitCount), maxCount: maxCount}
}

// 首字母小写 只能调用 工厂函数 创建
type BitMap struct {
	bits     []uint64
	bitCount uint64 // 已填入数字的数量
	maxCount uint64 // 容量
}

// 填入数字
func (b *BitMap) Set(num uint64) {
	byteIndex, bitPos := b.offset(num)
	// 1 左移 bitPos 位 进行 按位或 (置为 1)
	b.bits[byteIndex] |= bitmask[bitPos]
	b.bitCount++
}

// 清除填入的数字
func (b *BitMap) Unset(num uint64) {
	byteIndex, bitPos := b.offset(num)
	// 重置为空位 (重置为 0)
	b.bits[byteIndex] &= ^bitmask[bitPos]
	b.bitCount--
}

// 数字是否在位图中
func (b *BitMap) IsSet(num uint64) bool {
	byteIndex := num / bitSize
	if byteIndex >= uint64(len(b.bits)) {
		return false
	}
	bitPos := num % bitSize
	// 右移 bitPos 位 和 1 进行 按位与
	return !(b.bits[byteIndex]&bitmask[bitPos] == 0)
}

// 位图的容量
func (b *BitMap) Size() uint64 {
	return uint64(len(b.bits) * bitSize)
}

// 是否空位图
func (b *BitMap) IsEmpty() bool {
	return b.bitCount == 0
}

// 是否已填满
func (b *BitMap) IsFully() bool {
	return b.bitCount == b.maxCount
}

// 已填入的数字个数
func (b *BitMap) Count() uint64 {
	return b.bitCount
}

// 获取填入的数字切片
func (b *BitMap) GetData() []uint64 {
	var data []uint64
	count := b.Size()
	for index := uint64(0); index < count; index++ {
		if b.IsSet(index) {
			data = append(data, index)
		}
	}
	return data
}

// String ...
func (b *BitMap) String() string {
	var sb strings.Builder
	for index := len(b.bits) - 1; index >= 0; index-- {
		_, _ = sb.WriteString(uint64ToBinaryString(b.bits[index]))
		_, _ = sb.WriteString(" ")
	}
	return sb.String()
}

// Values returns a slice of ints
// represented by the values in the bitmap.
func (b *BitMap) Values() []uint64 {
	return b.bits
}

func (b *BitMap) offset(num uint64) (byteIndex uint64, bitPos byte) {
	byteIndex = num / bitSize // 字节索引
	if byteIndex >= uint64(len(b.bits)) {
		panic(fmt.Sprintf(" runtime error: index value %d out of range", byteIndex))
	}
	bitPos = byte(num % bitSize) // bit位置
	return byteIndex, bitPos
}

func uint64ToBinaryString(data uint64) string {
	sb := &strings.Builder{}
	_, _ = sb.WriteString(fmt.Sprintf("(%20d %d)[", data, NumberOf1(data)))
	for index := 0; index < bitSize; index++ {
		if index > 0 && index < bitSize && index%8 == 0 {
			_ = sb.WriteByte(' ')
		}
		if (bitmask[bitSize-1-index] & data) == 0 {
			_, _ = sb.WriteString("0")
		} else {
			_, _ = sb.WriteString("1")
		}
	}
	_ = sb.WriteByte(']')
	return sb.String()
}
