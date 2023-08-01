package time

import "time"

func MeasureFuncTime(target func()) time.Duration {
	t := time.Now()
	target()

	return time.Since(t)
}

