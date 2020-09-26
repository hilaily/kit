package datetime

import "time"

const (
	Tpl2006_01_02_15_04_05 = "2006-01-02 15:04:05"
	Tpl2006_01_02          = "2006-01-02"
	Tpl20060102            = "20060102"
	Tpl01_02               = "01-02"
	Tpl0102                = "0102"
)

// FormatTime 格式化成通用字符串格式
func Format2006_01_02_15_04_05(t time.Time) string {
	return t.Format(Tpl2006_01_02_15_04_05)
}

// Parse 解析通用格式的时间字符串
func Parse2006_01_02_15_04_05(t string) (time.Time, error) {
	return time.Parse(Tpl2006_01_02_15_04_05, t)
}
