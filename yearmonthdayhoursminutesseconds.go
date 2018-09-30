package date

import (
	"fmt"
	"time"
	"unicode"
)

const formatStringYYYYMMDDHHMMSS = "2006-01-02T15:04:05"
const parseJSONYYYYMMDDHHMMSS = "\"2006-01-02T15:04:05\""

// YYYYMMDDHHMMSS provides a yyyy-MM-ddThh:mm:ss date.
type YYYYMMDDHHMMSS time.Time

// MarshalJSON outputs JSON.
func (d YYYYMMDDHHMMSS) MarshalJSON() ([]byte, error) {
	return []byte("\"" + time.Time(d).Format(formatStringYYYYMMDDHHMMSS) + "\""), nil
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
	*d = YYYYMMDDHHMMSS(t)
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
func (d YYYYMMDDHHMMSS) String() string {
	return time.Time(d).String()
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
