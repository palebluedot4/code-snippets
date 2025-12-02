package datetime_test

import (
	"time"
)

var (
	utcTime       = time.Date(2025, 1, 2, 3, 4, 5, 123456789, time.UTC)
	utcTimeNoNano = time.Date(2025, 1, 2, 3, 4, 5, 0, time.UTC)

	fixedZone           = time.FixedZone("UTC+8", 8*60*60)
	fixedZoneTime       = time.Date(2025, 1, 2, 11, 4, 5, 123456789, fixedZone)
	fixedZoneTimeNoNano = time.Date(2025, 1, 2, 11, 4, 5, 0, fixedZone)

	zeroTime = time.Time{}
)
