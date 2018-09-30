package date

import (
	"fmt"
	"time"
	"unicode"
)

const formatStringYYYYMMDD = "2006-01-02"
const parseJSONYYYYMMDD = "\"2006-01-02\""

// NewYYYYMMDD creates a new YYYYMMDD date.
func NewYYYYMMDD(t time.Time) YYYYMMDD {
	return YYYYMMDD{
		Year:  t.Year(),
		Month: t.Month(),
		Day:   t.Day(),
	}
}

// YYYYMMDD provides a year, month and day type.
type YYYYMMDD struct {
	Year  int
	Month time.Month
	Day   int
}

// MarshalJSON outputs JSON.
func (d YYYYMMDD) MarshalJSON() ([]byte, error) {
	t := time.Date(d.Year, d.Month, d.Day, 0, 0, 0, 0, time.UTC)
	return []byte("\"" + t.Format(formatStringYYYYMMDD) + "\""), nil
}

// UnmarshalJSON handles incoming JSON.
func (d *YYYYMMDD) UnmarshalJSON(b []byte) (err error) {
	if err = checkJSONYYYYMMDD(string(b)); err != nil {
		return
	}
	t, err := time.ParseInLocation(parseJSONYYYYMMDD, string(b), time.UTC)
	if err != nil {
		return
	}
	d.Year = t.Year()
	d.Month = t.Month()
	d.Day = t.Day()
	return
}

// Time returns a time.Time from the year, month and day.
func (d YYYYMMDD) Time() time.Time {
	return time.Date(d.Year, d.Month, d.Day, 0, 0, 0, 0, time.UTC)
}

// ErrInvalidDateFormatYYYYMMDD is returned when the field cannot be parsed.
type ErrInvalidDateFormatYYYYMMDD error

func checkJSONYYYYMMDD(s string) error {
	// "1234-67-90"
	if len(s) != 12 {
		return newErrInvalidDateFormatYYYYMMDD(s, "invalid length")
	}
	if s[0] != '"' || s[11] != '"' {
		return newErrInvalidDateFormatYYYYMMDD(s, "not a quoted string")
	}
	for i, c := range s[1:10] {
		if i == 4 || i == 7 { // -1 offset because we started the loop at position 1.
			if c != '-' {
				return newErrInvalidDateFormatYYYYMMDD(s, "missing hyphens")
			}
			continue
		}
		if !unicode.IsDigit(c) {
			return newErrInvalidDateFormatYYYYMMDD(s, "invalid digit")
		}
	}
	return nil
}

func newErrInvalidDateFormatYYYYMMDD(input, explanation string) ErrInvalidDateFormatYYYYMMDD {
	return ErrInvalidDateFormatYYYYMMDD(fmt.Errorf("invalid date: '%s' did not match yyyy-MM-dd format: %v", input, explanation))
}
