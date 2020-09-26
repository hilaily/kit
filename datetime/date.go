package datetime

import "time"

// FormatDate 格式化只含有日期
func Format2006_01_02(t time.Time) string {
	return t.Format(Tpl2006_01_02)
}

// ParseDate 解析日期
func Parse2006_01_02(t string) (time.Time, error) {
	return time.Parse(Tpl2006_01_02, t)
}

// DateToday 今天的凌晨
func DateToday() time.Time {
	now := time.Now()
	timestamp := now.Unix() - int64(now.Second()) - int64(60*now.Minute()) - int64(60*60*now.Hour())
	return time.Unix(timestamp, 0)
}

// DateYestoday 昨天的凌晨
func DateYestoday() time.Time {
	today := DateToday()
	return today.AddDate(0, 0, -1)
}
