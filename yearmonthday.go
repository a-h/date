package date

import (
	"fmt"
	"time"
	"unicode"
)

const formatStringYYYYMMDD = "2006-01-02"
const parseJSONYYYYMMDD = "\"2006-01-02\""

// YYYYMMDD provides a year, month and day type.
type YYYYMMDD time.Time

// MarshalJSON outputs JSON.
func (d YYYYMMDD) MarshalJSON() ([]byte, error) {
	return []byte("\"" + time.Time(d).Format(formatStringYYYYMMDD) + "\""), nil
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
	*d = YYYYMMDD(t)
	return
}

// String // String returns the time formatted using the format string
//	"2006-01-02 15:04:05.999999999 -0700 MST"
//
// If the time has a monotonic clock reading, the returned string
// includes a final field "m=±<value>", where value is the monotonic
// clock reading formatted as a decimal number of seconds.
//
// The returned string is meant for debugging; for a stable serialized
// representation, use t.MarshalText, t.MarshalBinary, or t.Format
// with an explicit format string.
func (d YYYYMMDD) String() string {
	return time.Time(d).String()
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
