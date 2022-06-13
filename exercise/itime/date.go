package itime

import "time"

// 处理日期问题

const day0 = "2006-01-02"

const (
	OneDay = time.Hour * 24
)

func ParseFromStr(dateStr string) (time.Time, error) {
	return time.Parse(day0, dateStr)
}

func ParseToStr(date time.Time) string {
	return date.Format(day0)
}

func AddDay(date time.Time, days int) time.Time {
	if days < 1 {
		return date
	}
	return date.Add(OneDay * time.Duration(days))
}
