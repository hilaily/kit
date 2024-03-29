package datetime

import "time"

const (
	Tpl2006_01_02_15_04_05 = "2006-01-02 15:04:05"
	Tpl2006_01_02          = "2006-01-02"
	Tpl0102_1504           = "0102-1504"
	Tpl20060102            = "20060102"
	Tpl01_02               = "01-02"
	Tpl0102                = "0102"

	TPLYY_M_D_H_M_S = "2006-01-02 15:04:05"
	TPLYY_M_D       = "2006-01-02"
	TPLYYMD         = "20060102"
	TPLMD_HM        = "0102-1504"
	TPLMD           = "0102"
	TPLHM           = "15-04"
)

func STRYY_M_D_H_M_S() string {
	return time.Now().Format(TPLYY_M_D_H_M_S)
}

func STRYY_M_D() string {
	return time.Now().Format(TPLYY_M_D)
}

func STRYYMD() string {
	return time.Now().Format(TPLYYMD)
}

func STRMD_HM() string {
	return time.Now().Format(TPLMD_HM)
}
func STRMD() string {
	return time.Now().Format(TPLMD)
}

func STRHM() string {
	return time.Now().Format(TPLHM)
}

// FormatTime 格式化成通用字符串格式
func Format2006_01_02_15_04_05(t time.Time) string {
	return t.Format(Tpl2006_01_02_15_04_05)
}

func Format0102_1504(t time.Time) string {
	return t.Format(Tpl0102_1504)
}

// Parse 解析通用格式的时间字符串
func Parse2006_01_02_15_04_05(t string) (time.Time, error) {
	return time.Parse(Tpl2006_01_02_15_04_05, t)
}
