package datetime_test

import (
	"testing"
	"time"

	"github.com/palebluedot4/code-snippets/cookbook/datetime"
)

var (
	utcTime       = time.Date(2025, 1, 2, 3, 4, 5, 123456789, time.UTC)
	utcTimeNoNano = time.Date(2025, 1, 2, 3, 4, 5, 0, time.UTC)

	fixedZone           = time.FixedZone("UTC+8", 8*60*60)
	fixedZoneTime       = time.Date(2025, 1, 2, 11, 4, 5, 123456789, fixedZone)
	fixedZoneTimeNoNano = time.Date(2025, 1, 2, 11, 4, 5, 0, fixedZone)

	zeroTime = time.Time{}
)

func TestToISO8601(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input time.Time
		want  string
	}{
		{
			name:  "UTC time with nanoseconds (should be truncated)",
			input: utcTime,
			want:  "2025-01-02T03:04:05Z",
		},
		{
			name:  "time in a fixed zone with nanoseconds (should be truncated)",
			input: fixedZoneTime,
			want:  "2025-01-02T11:04:05+08:00",
		},
		{
			name:  "UTC time without nanoseconds",
			input: utcTimeNoNano,
			want:  "2025-01-02T03:04:05Z",
		},
		{
			name:  "time in a fixed zone without nanoseconds",
			input: fixedZoneTimeNoNano,
			want:  "2025-01-02T11:04:05+08:00",
		},
		{
			name:  "zero value of time.Time",
			input: zeroTime,
			want:  "0001-01-01T00:00:00Z",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := datetime.ToISO8601(tt.input)
			if got != tt.want {
				t.Errorf("ToISO8601() got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestToISO8601Nano(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input time.Time
		want  string
	}{
		{
			name:  "UTC time with nanoseconds",
			input: utcTime,
			want:  "2025-01-02T03:04:05.123456789Z",
		},
		{
			name:  "time in a fixed zone with nanoseconds",
			input: fixedZoneTime,
			want:  "2025-01-02T11:04:05.123456789+08:00",
		},
		{
			name:  "UTC time without nanoseconds",
			input: utcTimeNoNano,
			want:  "2025-01-02T03:04:05Z",
		},
		{
			name:  "time in a fixed zone without nanoseconds",
			input: fixedZoneTimeNoNano,
			want:  "2025-01-02T11:04:05+08:00",
		},
		{
			name:  "zero value of time.Time",
			input: zeroTime,
			want:  "0001-01-01T00:00:00Z",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := datetime.ToISO8601Nano(tt.input)
			if got != tt.want {
				t.Errorf("ToISO8601Nano() got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestToDateTime(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input time.Time
		want  string
	}{
		{
			name:  "UTC time with nanoseconds",
			input: utcTime,
			want:  "2025-01-02 03:04:05",
		},
		{
			name:  "time in a fixed zone with nanoseconds",
			input: fixedZoneTime,
			want:  "2025-01-02 11:04:05",
		},
		{
			name:  "UTC time without nanoseconds",
			input: utcTimeNoNano,
			want:  "2025-01-02 03:04:05",
		},
		{
			name:  "time in a fixed zone without nanoseconds",
			input: fixedZoneTimeNoNano,
			want:  "2025-01-02 11:04:05",
		},
		{
			name:  "zero value of time.Time",
			input: zeroTime,
			want:  "0001-01-01 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := datetime.ToDateTime(tt.input)
			if got != tt.want {
				t.Errorf("ToDateTime() got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestToDateOnly(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input time.Time
		want  string
	}{
		{
			name:  "UTC time with nanoseconds",
			input: utcTime,
			want:  "2025-01-02",
		},
		{
			name:  "time in a fixed zone with nanoseconds",
			input: fixedZoneTime,
			want:  "2025-01-02",
		},
		{
			name:  "UTC time without nanoseconds",
			input: utcTimeNoNano,
			want:  "2025-01-02",
		},
		{
			name:  "time in a fixed zone without nanoseconds",
			input: fixedZoneTimeNoNano,
			want:  "2025-01-02",
		},
		{
			name:  "zero value of time.Time",
			input: zeroTime,
			want:  "0001-01-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := datetime.ToDateOnly(tt.input)
			if got != tt.want {
				t.Errorf("ToDateOnly() got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestToTimeOnly(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input time.Time
		want  string
	}{
		{
			name:  "UTC time with nanoseconds",
			input: utcTime,
			want:  "03:04:05",
		},
		{
			name:  "time in a fixed zone with nanoseconds",
			input: fixedZoneTime,
			want:  "11:04:05",
		},
		{
			name:  "UTC time without nanoseconds",
			input: utcTimeNoNano,
			want:  "03:04:05",
		},
		{
			name:  "time in a fixed zone without nanoseconds",
			input: fixedZoneTimeNoNano,
			want:  "11:04:05",
		},
		{
			name:  "zero value of time.Time",
			input: zeroTime,
			want:  "00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := datetime.ToTimeOnly(tt.input)
			if got != tt.want {
				t.Errorf("ToTimeOnly() got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestToHTTPHeader(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input time.Time
		want  string
	}{
		{
			name:  "from UTC time with nanoseconds",
			input: utcTime,
			want:  "Thu, 02 Jan 2025 03:04:05 GMT",
		},
		{
			name:  "from fixed zone time (should convert to GMT)",
			input: fixedZoneTime,
			want:  "Thu, 02 Jan 2025 03:04:05 GMT",
		},
		{
			name:  "from UTC time without nanoseconds",
			input: utcTimeNoNano,
			want:  "Thu, 02 Jan 2025 03:04:05 GMT",
		},
		{
			name:  "from time in a fixed zone without nanoseconds",
			input: fixedZoneTimeNoNano,
			want:  "Thu, 02 Jan 2025 03:04:05 GMT",
		},
		{
			name:  "from zero value of time.Time",
			input: zeroTime,
			want:  "Mon, 01 Jan 0001 00:00:00 GMT",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := datetime.ToHTTPHeader(tt.input)
			if got != tt.want {
				t.Errorf("ToHTTPHeader() got %q, want %q", got, tt.want)
			}
		})
	}
}
