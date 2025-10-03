package datetime

import "time"

func ToISO8601(t time.Time) string {
	return t.Format(time.RFC3339)
}
