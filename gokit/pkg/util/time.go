package util

import (
	"fmt"
	"time"
)

// Now

func CurrentUnix() int64 {
	return time.Now().Unix()
}

func CurrentNano() int64 {
	return time.Now().UnixNano()
}

func CurrentMilli() int64 {
	return time.Now().UnixMilli()
}

func CurrentMicro() int64 {
	return time.Now().UnixMicro()
}

func DateInLocation(date time.Time, location *time.Location) (int, time.Month, int) {
	var dateInLocation time.Time
	if location == nil {
		dateInLocation = date.UTC()
	} else {
		dateInLocation = date.In(location)
	}

	return dateInLocation.Date()
}

// Beginning

func BeginningOfDay(date time.Time) time.Time {
	var (
		y, m, d  = DateInLocation(date, nil)
		todayStr = fmt.Sprintf("%04d-%02d-%02dT00:00:00.000000000Z", y, int(m), d)
	)

	date, err := time.Parse(time.RFC3339Nano, todayStr)
	if err != nil {
		panic(fmt.Sprintf("cannot parse date due to: %v", err))
	}

	return date
}

func BeginningOfDayUnix(date time.Time) int64 {
	return BeginningOfDay(date).Unix()
}

func BeginningOfToday() time.Time {
	return BeginningOfDay(time.Now())
}

func BeginningOfDayInLocation(date time.Time, location *time.Location) time.Time {
	y, m, d := DateInLocation(date, location)

	return time.Date(y, m, d, 0, 0, 0, 0, location)
}

func BeginningOfTodayInLocation(location *time.Location) time.Time {
	return BeginningOfDayInLocation(time.Now(), location)
}

// Ending

func EndingOfDay(date time.Time) time.Time {
	var (
		y, m, d  = DateInLocation(date, nil)
		todayStr = fmt.Sprintf("%04d-%02d-%02dT23:59:59.999999999Z", y, int(m), d)
	)

	date, err := time.Parse(time.RFC3339Nano, todayStr)
	if err != nil {
		panic(fmt.Sprintf("cannot parse date due to: %v", err))
	}

	return date
}

func EndingOfDayUnix(date time.Time) int64 {
	return EndingOfDay(date).Unix()
}

func EndingOfToday() time.Time {
	return EndingOfDay(time.Now())
}

func EndingOfDayInLocation(date time.Time, location *time.Location) time.Time {
	if location == nil {
		location, _ = time.LoadLocation("UTC")
	}
	y, m, d := DateInLocation(date, location)

	return time.Date(y, m, d, 23, 59, 59, 999999999, location)
}

func EndingOfTodayInLocation(location *time.Location) time.Time {
	if location == nil {
		location, _ = time.LoadLocation("UTC")
	}
	y, m, d := DateInLocation(time.Now(), location)

	return time.Date(y, m, d, 23, 59, 59, 999999999, location)
}

func EndingOfDayInLocationPreciseToMilli(date time.Time, location *time.Location) time.Time {
	if location == nil {
		location, _ = time.LoadLocation("UTC")
	}
	y, m, d := DateInLocation(date, location)

	return time.Date(y, m, d, 23, 59, 59, 999000000, location)
}

// Time ago

func TimeAgo(numOfDays int) time.Time {
	return time.Now().Add(-time.Duration(numOfDays*24) * time.Hour)
}

func TimeAgoInLocation(numOfDays int, location *time.Location) time.Time {
	return time.Now().In(location).Add(-time.Duration(numOfDays*24) * time.Hour)
}

func TimeAgoUnix(numOfDays int) uint64 {
	return uint64(TimeAgo(numOfDays).Unix())
}

// Parse functions

func MustParseTimeWithDefaultLayout(t string) time.Time {
	r, err := time.Parse("2006-1-2 15:04:05", t)
	if err != nil {
		panic(fmt.Sprintf("cannot parse time for %s", t))
	}

	return r
}

func MustParseTimeDuration(ds string) time.Duration {
	d, err := time.ParseDuration(ds)
	if err != nil {
		panic(fmt.Sprintf("cannot parse time duration for %s", ds))
	}

	return d
}

func MustParseTimeWithLayout(t string, l string) time.Time {
	r, err := time.Parse(l, t)
	if err != nil {
		panic(fmt.Sprintf("cannot parse time for %s", t))
	}

	return r
}

// Other functions

func IsInTodayInLocation(date time.Time, location *time.Location) bool {
	yearOfToday, monthOfToday, dayOfToday := time.Now().In(location).Date()
	yearOfDate, monthOfDate, dayOfDate := date.In(location).Date()

	return yearOfToday == yearOfDate && monthOfToday == monthOfDate && dayOfToday == dayOfDate
}

func GetWeekOrderNumberFromBeginOfYearInLocation(date time.Time, location *time.Location) int {
	daysPassed := date.In(location).YearDay()

	return (daysPassed-1)/7 + 1
}
