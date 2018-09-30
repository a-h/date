package date

import (
	"encoding/json"
	"testing"
	"time"
)

func Test(t *testing.T) {
	tests := []struct {
		input          string
		expectedYMD    time.Time
		expectedYMDHMS time.Time
	}{
		{
			input:          `{"ymd":"2008-12-25","ymdhms":"1742-12-25T13:32:20"}`,
			expectedYMD:    time.Date(2008, time.December, 25, 0, 0, 0, 0, time.UTC),
			expectedYMDHMS: time.Date(1742, time.December, 25, 13, 32, 20, 0, time.UTC),
		},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			var e example
			err := json.Unmarshal([]byte(test.input), &e)
			if err != nil {
				t.Errorf("unexpected error during unmarshalling: %v", err)
			}
			if !test.expectedYMD.Equal(time.Time(e.YMD)) {
				t.Errorf("expected YMD of %v, got %v", test.expectedYMD, e.YMD)
			}
			if !test.expectedYMDHMS.Equal(time.Time(e.YMDHMS)) {
				t.Errorf("expected YMDHMS of %v, got %v", test.expectedYMDHMS, e.YMDHMS)
			}
			output, err := json.Marshal(e)
			if test.input != string(output) {
				t.Errorf("expected JSON output to equal input, but got '%v'", string(output))
			}
		})
	}
}

func TestErrors(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    `{"ymd":"","ymdhms":"1742-02-30T12:30:60"}`,
			expected: `invalid date: '""' did not match yyyy-MM-dd format: invalid length`,
		},
		{
			input:    `{"ymd":"2008-12-25","ymdhms":""}`,
			expected: `invalid date: '""' did not match yyyy-MM-ddThh:mm:ss format: invalid length`,
		},
		{
			input:    `{"ymd":"2008_12_25","ymdhms":"1742-12-25T13:32:20"}`,
			expected: `invalid date: '"2008_12_25"' did not match yyyy-MM-dd format: missing hyphens`,
		},
		{
			input:    `{"ymd":"2008-12-25","ymdhms":"1742_12_25T13:32:20"}`,
			expected: `invalid date: '"1742_12_25T13:32:20"' did not match yyyy-MM-ddThh:mm:ss format: missing hyphens`,
		},
		{
			input:    `{"ymd":"2008-13-25","ymdhms":"1742-12-25T13:32:20"}`,
			expected: `parsing time ""2008-13-25"": month out of range`,
		},
		{
			input:    `{"ymd":"2008-02-30","ymdhms":"1742-12-25T13:32:20"}`,
			expected: `parsing time ""2008-02-30"": day out of range`,
		},
		{
			input:    `{"ymd":"2008-12-25","ymdhms":"1742-13-25T13:32:20"}`,
			expected: `parsing time ""1742-13-25T13:32:20"": month out of range`,
		},
		{
			input:    `{"ymd":"2008-02-14","ymdhms":"1742-02-30T13:32:20"}`,
			expected: `parsing time ""1742-02-30T13:32:20"": day out of range`,
		},
		{
			input:    `{"ymd":"2008-02-14","ymdhms":"1742-02-30T24:32:20"}`,
			expected: `parsing time ""1742-02-30T24:32:20"": hour out of range`,
		},
		{
			input:    `{"ymd":"2008-02-14","ymdhms":"1742-02-30T12:60:20"}`,
			expected: `parsing time ""1742-02-30T12:60:20"": minute out of range`,
		},
		{
			input:    `{"ymd":"2008-02-14","ymdhms":"1742-02-30T12:30:60"}`,
			expected: `parsing time ""1742-02-30T12:30:60"": second out of range`,
		},
		{
			input:    `{"ymd":"a008-02-14","ymdhms":"1742-02-30T12:30:60"}`,
			expected: `invalid date: '"a008-02-14"' did not match yyyy-MM-dd format: invalid digit`,
		},
		{
			input:    `{"ymd":"2008-02-14","ymdhms":"1742-02-30X12:30:60"}`,
			expected: `invalid date: '"1742-02-30X12:30:60"' did not match yyyy-MM-ddThh:mm:ss format: missing T`,
		},
		{
			input:    `{"ymd":"2008-02-14","ymdhms":"1742-02-30T12X30:60"}`,
			expected: `invalid date: '"1742-02-30T12X30:60"' did not match yyyy-MM-ddThh:mm:ss format: missing colons`,
		},
		{
			input:    `{"ymd":"2008-02-14","ymdhms":"1742-02-30T12:30X60"}`,
			expected: `invalid date: '"1742-02-30T12:30X60"' did not match yyyy-MM-ddThh:mm:ss format: missing colons`,
		},
		{
			input:    `{"ymd":"2008-02-14","ymdhms":"a742-02-30T12:30:60"}`,
			expected: `invalid date: '"a742-02-30T12:30:60"' did not match yyyy-MM-ddThh:mm:ss format: invalid digit`,
		},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			var e example
			err := json.Unmarshal([]byte(test.input), &e)
			if err == nil {
				t.Error("expected error not returned")
				return
			}
			if err.Error() != test.expected {
				t.Errorf("expected error '%v' but got: %v", test.expected, err)
			}
		})
	}
}

type example struct {
	YMD    YYYYMMDD       `json:"ymd"`
	YMDHMS YYYYMMDDHHMMSS `json:"ymdhms"`
}

func TestUninitializedValues(t *testing.T) {
	bytes, err := json.Marshal(example{})
	if err != nil {
		t.Errorf("Error marshalling struct with uninitialized fields: %v", err)
	}
	expected := `{"ymd":"0001-01-01","ymdhms":"0001-01-01T00:00:00"}`
	if string(bytes) != expected {
		t.Errorf("expected '%v', got '%v'", expected, string(bytes))
	}
}
