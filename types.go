package metrics

import "time"

// Metrics handles incremental metrics that should expire after some time.
type Metrics interface {
	// Cleanup will remove old values past the timeout within the specified bucket,
	// if the bucket is defined. If the bucket is not defined this operation does
	// nothing.
	Cleanup(string)
	// Get returns the value for the given bucket, if it is defined. If the bucket
	// is not defined, this operation will return 0 (zero).
	Get(string) int
	// Remove will remove the bucket, and all of it's associated data, if it is
	// defined. If it is not defined, this operation does nothing.
	Remove(string)
	// Add or subtract the given value to the given bucket. If the bucket does not
	// exist, it will be created. The time given will be the time associated with
	// the given quantity change.
	Update(string, int, time.Time)
}
