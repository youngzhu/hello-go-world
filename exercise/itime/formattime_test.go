package itime_test

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	// date, _ := time.Parse(time.RFC3339, "2021-09-18T00:00:00+08:00")
	date, _ := time.Parse("2006-01-02", "2021-09-18")

	yearExpect := 2021
	yearGot := date.Year()
	if yearGot != yearExpect {
		t.Errorf("year: got %d; expect %d", yearGot, yearExpect)
	}

	monthExpect := time.September
	monthGot := date.Month()
	if monthGot != monthExpect {
		t.Errorf("month: got %d; expect %d", monthGot, monthExpect)
	}

	dayExpect := 18
	dayGot := date.Day()
	if dayGot != dayExpect {
		t.Errorf("day: got %d; expect %d", dayGot, dayExpect)
	}
}

func TestNewDate(t *testing.T) {
	date := time.Date(2021, 9, 18, 13, 14, 59, 123456, time.UTC)

	yearExpect := 2021
	yearGot := date.Year()
	if yearGot != yearExpect {
		t.Errorf("year: got %d; expect %d", yearGot, yearExpect)
	}

	monthExpect := time.September
	monthGot := date.Month()
	if monthGot != monthExpect {
		t.Errorf("month: got %d; expect %d", monthGot, monthExpect)
	}

	dayExpect := 18
	dayGot := date.Day()
	if dayGot != dayExpect {
		t.Errorf("day: got %d; expect %d", dayGot, dayExpect)
	}

	hourExpect := 13
	hourGot := date.Hour()
	if hourGot != hourExpect {
		t.Errorf("hour: got %d; expect %d", hourGot, hourExpect)
	}

	minuteExpect := 14
	minuteGot := date.Minute()
	if minuteGot != minuteExpect {
		t.Errorf("minute: got %d; expect %d", minuteGot, minuteExpect)
	}

	secondExpect := 59
	secondGot := date.Second()
	if secondGot != secondExpect {
		t.Errorf("second: got %d; expect %d", secondGot, secondExpect)
	}

	nanosExpect := 123456
	nanosGot := date.Nanosecond()
	if nanosGot != nanosExpect {
		t.Errorf("nanos: got %d; expect %d", nanosGot, nanosExpect)
	}

	locExpect := time.UTC
	locGot := date.Location()
	if locGot != locExpect {
		t.Errorf("loc: got %s; expect %s", locGot, locExpect)
	}

}
