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
