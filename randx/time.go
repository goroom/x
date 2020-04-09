package randx

import "time"

func (r *Rand) TimeRange(min, max time.Duration) time.Duration {
	return time.Duration(r.Int63n(int64(max-min)) + int64(min))
}
