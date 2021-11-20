package mathutil

import (
	"fmt"
	"goutils/fmtutil"
	"time"
)

func IsNUmber(c byte) bool {
	return c >= '0' && c <= '9'
}

func Percent(val, total int) float64 {
	if total == 0 {
		return float64(0)
	}
	return (float64(val) / float64(total)) * 100
}

func ElapsedTime(startTime time.Time) string {
	return fmt.Sprintf("%.3f", time.Since(startTime).Seconds()*1000)
}

func DataSize(size uint64) string {
	return fmtutil.DataSize(size)
}

func HowLongAgo(sec int64) string {
	return fmtutil.HowLongAgo(sec)
}
