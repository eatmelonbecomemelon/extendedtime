package extendedtime

import (
	"time"
)

func GetNaturalRange(t time.Time, rangeType int) (start, end time.Time) {
	year, month, day := t.Date()
	switch rangeType {
	case TimeRangeHour:
		start = time.Date(year, month, day, t.Hour(), 0, 0, 0, t.Location())
		end = start.Add(time.Hour)
	case TimeRangeDay:
		start = time.Date(year, month, day, 0, 0, 0, 0, t.Location())
		end = start.Add(24 * time.Hour)
	case TimeRangeWeek:
		dayShift := -time.Duration(t.Weekday() - time.Sunday)
		start = time.Date(year, month, day, 0, 0, 0, 0, t.Location()).Add(dayShift * 24 * time.Hour)
		end = start.Add(7 * 24 * time.Hour)
	case TimeRangeMonth:
		start = time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
		endYear, endMonth := year, month+1
		if endMonth == 13 {
			endYear++
			endMonth = 1
		}
		end = time.Date(endYear, endMonth, 1, 0, 0, 0, 0, t.Location())
	case TimeRangeYear:
		start = time.Date(year, 1, 1, 0, 0, 0, 0, t.Location())
		end = time.Date(year+1, 1, 1, 0, 0, 0, 0, t.Location())
	}
	return
}

func GetLastRange(t time.Time, rangeType int) (start, end time.Time) {
	end = t
	year, month, day := t.Date()
	hour, minute, second, nanosecond := t.Hour(), t.Minute(), t.Second(), t.Nanosecond()
	switch rangeType {
	case TimeRangeHour:
		start = t.Add(-time.Hour)
	case TimeRangeDay:
		start = t.Add(-24 * time.Hour)
	case TimeRangeWeek:
		start = t.Add(-7 * 24 * time.Hour)
	case TimeRangeMonth:
		endYear, endMonth := year, month-1
		if endMonth == 0 {
			endYear--
			endMonth = 12
		}
		start = time.Date(endYear, endMonth, day, hour, minute, second, nanosecond, t.Location())
	case TimeRangeYear:
		start = time.Date(year-1, month, day, hour, minute, second, nanosecond, t.Location())
	}
	return
}
