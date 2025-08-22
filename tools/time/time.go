package time

import (
	"fmt"
	"time"
)

// region convert
func ISO(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}
func ToMill(str string) int64 {
	parse, _ := time.Parse("2006-01-02 15:04:05", str)
	return parse.UnixMilli()
}
func ToTimeString(mill int64) string {
	milli := time.UnixMilli(mill)
	return milli.Format("2006-01-02 15:04:05")
}
func ZeroTime() int64 {
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return date.UnixMilli()
}
func DateTimeToTimeSpan(dt string) int64 {
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation(time.DateTime, dt, loc)
	if err != nil {
		return 0
	}
	return t.Unix()
}
func TimeSpanToDateTime(ts int64) string {
	return time.Unix(ts, 0).Format(time.DateTime)
}

// endregion
// region pare
func FromString(layout, value string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		loc = time.Local
	}

	t, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		return t, fmt.Errorf("invalid time string: %w", err)
	}
	return t, nil
}
func FromSec(sec int64, loc *time.Location) time.Time {
	t := time.Unix(sec, 0)
	if loc != nil {
		t = t.In(loc)
	}
	return t
}
func FromMs(ms int64, loc *time.Location) time.Time {
	t := time.UnixMilli(ms)
	if loc != nil {
		t = t.In(loc)
	}
	return t
}
func FromUs(us int64, loc *time.Location) time.Time {
	t := time.UnixMicro(us)
	if loc != nil {
		t = t.In(loc)
	}
	return t
}

func FromNs(ns int64, loc *time.Location) time.Time {
	sec := ns / 1e9
	nsec := ns % 1e9
	t := time.Unix(sec, nsec)
	if loc != nil {
		t = t.In(loc)
	}
	return t
}

func GetBeijingTime(layout, value string) (t time.Time, err error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return
	}
	return time.ParseInLocation(layout, value, loc)
}

func GetTimezoneTime(timezone string, layout, value string) (t time.Time, err error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return
	}
	return time.ParseInLocation(layout, value, loc)
}

//endregion
