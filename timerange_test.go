package extendedtime

import (
	"testing"
	"time"
)

func TestGetNaturalRange(t *testing.T) {
	layout := "2006-01-02T15:04:05.999999999"
	assertEqual := func(start, end time.Time, expectStart, expectEnd string, timRange int) {
		if start.Format(layout) != expectStart || end.Format(layout) != expectEnd {
			t.Errorf("timerange :%d: test failed: is start correct:%v, is end correct: %v", timRange, start.Format(layout) == expectStart, end.Format(layout) == expectEnd)
		}
	}

	now, _ := time.Parse(layout, "2018-02-02T14:01:12.582999999")

	start, end := GetNaturalRange(now, TimeRangeHour)
	assertEqual(start, end, "2018-02-02T14:00:00", "2018-02-02T15:00:00", TimeRangeHour)

	start, end = GetNaturalRange(now, TimeRangeDay)
	assertEqual(start, end, "2018-02-02T00:00:00", "2018-02-03T00:00:00", TimeRangeDay)

	start, end = GetNaturalRange(now, TimeRangeWeek)
	assertEqual(start, end, "2018-01-28T00:00:00", "2018-02-04T00:00:00", TimeRangeWeek)

	start, end = GetNaturalRange(now, TimeRangeMonth)
	assertEqual(start, end, "2018-02-01T00:00:00", "2018-03-01T00:00:00", TimeRangeMonth)

	start, end = GetNaturalRange(now, TimeRangeYear)
	assertEqual(start, end, "2018-01-01T00:00:00", "2019-01-01T00:00:00", TimeRangeYear)

}
