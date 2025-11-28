package timex

import "time"

// DateTimeLayout 标准日期时间格式：年-月-日 时:分:秒（例如：2023-10-01 14:30:00）
const DateTimeLayout = "2006-01-02 15:04:05"

// DateLayout 标准日期格式：年-月-日（例如：2023-10-01）
const DateLayout = "2006-01-02"

// NowDateTime 获取当前时间的标准日期时间字符串
// 返回格式遵循 DateTimeLayout（2006-01-02 15:04:05）
func NowDateTime() string {
	return time.Now().Format(DateTimeLayout)
}

// NowDate 获取当前时间的标准日期字符串
// 返回格式遵循 DateLayout（2006-01-02）
func NowDate() string {
	return time.Now().Format(DateLayout)
}

// FormatDate 将时间相关类型格式化为标准日期字符串
// 支持的输入类型：time.Time、*time.Time
// 格式化规则：
// 1. 若输入为 time.Time 且非零值，返回 DateLayout 格式字符串
// 2. 若输入为 *time.Time 且指针非空、指向的时间非零值，返回 DateLayout 格式字符串
// 3. 输入为零值、nil 指针或不支持的类型，返回空字符串
func FormatDate(i any) string {
	switch v := i.(type) {
	case time.Time:
		if v.IsZero() {
			return ""
		}
		return v.Format(DateTimeLayout)
	case *time.Time:
		if v != nil {
			if v.IsZero() {
				return ""
			}
			return v.Format(DateTimeLayout)
		}
	}
	return ""
}
