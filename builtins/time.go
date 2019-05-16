package builtins

import (
	"time"
)

type Time struct {
	ANSIC       string
	UnixDate    string
	RubyDate    string
	RFC822      string
	RFC822Z     string
	RFC850      string
	RFC1123     string
	RFC1123Z    string
	RFC3339     string
	RFC3339Nano string
	Kitchen     string
	Stamp       string
	StampMilli  string
	StampMicro  string
	StampNano   string
	Nanosecond  int
	Microsecond int
	Millisecond int
	Second      int
	Minute      int
	Hour        int
}

func NewTime() Time {
	return Time{
		ANSIC:       time.ANSIC,
		UnixDate:    time.UnixDate,
		RubyDate:    time.RubyDate,
		RFC822:      time.RFC822,
		RFC822Z:     time.RFC822Z,
		RFC850:      time.RFC850,
		RFC1123:     time.RFC1123,
		RFC1123Z:    time.RFC1123Z,
		RFC3339:     time.RFC3339,
		RFC3339Nano: time.RFC3339Nano,
		Kitchen:     time.Kitchen,
		Stamp:       time.Stamp,
		StampMilli:  time.StampMilli,
		StampMicro:  time.StampMicro,
		StampNano:   time.StampNano,
		Nanosecond:  int(time.Nanosecond),
		Microsecond: int(time.Microsecond),
		Millisecond: int(time.Millisecond),
		Second:      int(time.Second),
		Minute:      int(time.Minute),
		Hour:        int(time.Hour),
	}
}

func (Time) After(d int) <-chan time.Time {
	return time.After(time.Duration(d))
}
func (Time) Sleep(d int) {
	time.Sleep(time.Duration(d))
}

func (Time) Tick(d int) <-chan time.Time {
	return time.Tick(time.Duration(d))
}

func (Time) ParseDuration(s string) (int, error) {
	i, err := time.ParseDuration(s)
	return int(i), err
}
func (Time) Since(t time.Time) int {
	return int(time.Since(t))
}
func (Time) Until(t time.Time) int {
	return int(time.Until(t))
}
func (Time) FixedZone(name string, offset int) *time.Location {
	return time.FixedZone(name, offset)
}
func (Time) LoadLocation(name string) (*time.Location, error) {
	return time.LoadLocation(name)
}
func (Time) LoadLocationFromTZData(name string, data []byte) (*time.Location, error) {
	return time.LoadLocationFromTZData(name, data)
}
func (Time) NewTicker(d int) *time.Ticker {
	return time.NewTicker(time.Duration(d))
}
func (Time) Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
	return time.Date(year, month, day, hour, min, sec, nsec, loc)
}
func (Time) Now() time.Time {
	return time.Now()
}
func (Time) Parse(layout, value string) (time.Time, error) {
	return time.Parse(layout, value)
}
func (Time) ParseInLocation(layout, value string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(layout, value, loc)
}
func (Time) Unix(sec int64, nsec int64) time.Time {
	return time.Unix(sec, nsec)
}
func (Time) NewTimer(d int) *time.Timer {
	return time.NewTimer(time.Duration(d))
}
