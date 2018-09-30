package date

import (
	"fmt"
	"time"
	"unicode"
)

const formatStringYYYYMMDDHHMMSS = "2006-01-02T15:04:05"
const parseJSONYYYYMMDDHHMMSS = "\"2006-01-02T15:04:05\""

// NewYYYYMMDDHHMMSS creates a new yyyy-MM-ddThh:mm:ss date.
func NewYYYYMMDDHHMMSS(t time.Time) YYYYMMDDHHMMSS {
	return YYYYMMDDHHMMSS{
		Year:   t.Year(),
		Month:  t.Month(),
		Day:    t.Day(),
		Hour:   t.Hour(),
		Minute: t.Minute(),
		Second: t.Second(),
	}
}

// YYYYMMDDHHMMSS provides a yyyy-MM-ddThh:mm:ss date.
type YYYYMMDDHHMMSS struct {
	Year   int
	Month  time.Month
	Day    int
	Hour   int
	Minute int
	Second int
}

// MarshalJSON outputs JSON.
func (d YYYYMMDDHHMMSS) MarshalJSON() ([]byte, error) {
	if d.Year == 0 && d.Month == 0 && d.Day == 0 {
		// Initialise to a minimum date, it's not possible to have a month of zero.
		d.Year = 0
		d.Month = time.January
		d.Day = 1
	}
	t := time.Date(d.Year, d.Month, d.Day, d.Hour, d.Minute, d.Second, 0, time.UTC)
	return []byte("\"" + t.Format(formatStringYYYYMMDDHHMMSS) + "\""), nil
}

// UnmarshalJSON handles incoming JSON.
func (d *YYYYMMDDHHMMSS) UnmarshalJSON(b []byte) (err error) {
	if err = checkJSONYYYYMMDDHHMMSS(string(b)); err != nil {
		return
	}
	t, err := time.ParseInLocation(parseJSONYYYYMMDDHHMMSS, string(b), time.UTC)
	if err != nil {
		return
	}
	d.Year = t.Year()
	d.Month = t.Month()
	d.Day = t.Day()
	d.Hour = t.Hour()
	d.Minute = t.Minute()
	d.Second = t.Second()
	return
}

// Time returns a time.Time from the year, month, day, hours, minutes and seconds.
func (d YYYYMMDDHHMMSS) Time() time.Time {
	return time.Date(d.Year, d.Month, d.Day, d.Hour, d.Minute, d.Second, 0, time.UTC)
}

// ErrInvalidDateFormatYYYYMMDDHHMMSS is returned when the field cannot be parsed.
type ErrInvalidDateFormatYYYYMMDDHHMMSS error

func checkJSONYYYYMMDDHHMMSS(s string) error {
	// "1234-67-90T23:56:78"
	if len(s) != 21 {
		return newErrInvalidDateFormatYYYYMMDDHHMMSS(s, "invalid length")
	}
	if s[0] != '"' || s[20] != '"' {
		return newErrInvalidDateFormatYYYYMMDDHHMMSS(s, "not a quoted string")
	}
	for i, c := range s[1:20] {
		if i == 4 || i == 7 {
			if c != '-' { // -1 offset because we started the loop at position 1.
				return newErrInvalidDateFormatYYYYMMDDHHMMSS(s, "missing hyphens")
			}
			continue
		}
		if i == 10 {
			if c != 'T' {
				return newErrInvalidDateFormatYYYYMMDDHHMMSS(s, "missing T")
			}
			continue
		}
		if i == 13 || i == 16 {
			if c != ':' {
				return newErrInvalidDateFormatYYYYMMDDHHMMSS(s, "missing colons")
			}
			continue
		}
		if !unicode.IsDigit(c) {
			return newErrInvalidDateFormatYYYYMMDDHHMMSS(s, "invalid digit")
		}
	}
	return nil
}

func newErrInvalidDateFormatYYYYMMDDHHMMSS(input, explanation string) ErrInvalidDateFormatYYYYMMDDHHMMSS {
	return ErrInvalidDateFormatYYYYMMDDHHMMSS(fmt.Errorf("invalid date: '%s' did not match yyyy-MM-ddThh:mm:ss format: %v", input, explanation))
}
