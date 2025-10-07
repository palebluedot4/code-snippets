package datetime

import "time"

func ToISO8601(t time.Time) string {
	return t.Format(time.RFC3339)
}

func ToISO8601Nano(t time.Time) string {
	return t.Format(time.RFC3339Nano)
}

func ToDateTime(t time.Time) string {
	return t.Format(time.DateTime)
}

func ToDateOnly(t time.Time) string {
	return t.Format(time.DateOnly)
}
