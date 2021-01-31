// InitBackoffPolicy is a backoff-policy intervals ranging up to 5 seconds.
func InitSimpleBackoffPolicy() []int {
	backoffIntervals := []int{0, 100, 500, 3000, 5000}
	return backoffIntervals
}

// Randomization of interval with a jitter function for uniform distribution
// between two successive retry attempts
func BackoffDuration(n int, intervals []int) time.Duration {
	if n >= len(intervals) {
		n = len(intervals) - 1
	}

	return time.Duration(jitter(intervals[n], n)) * time.Millisecond
}

func jitter(interval int, attempt int) int {
	if interval == 0 {
		return 0
	}
	// generates a range of [0-8199] milliseconds
	return 100*int(math.Pow(2, float64(attempt))) + rand.Intn(interval)
}