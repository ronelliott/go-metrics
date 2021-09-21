package metrics

import (
	"sync"
	"time"
)

// MetricsDatum represents an entry within the metrics data structure
type MetricsDatum struct {
	// The value for this entry
	Value int
	// When this entry occurred
	When time.Time
}

// MetricsImpl implements the Metrics interface
type MetricsImpl struct {
	data    map[string][]MetricsDatum
	lock    sync.Mutex
	timeout time.Duration
}

// New creates a new instance of MetricsImpl
func New(timeout time.Duration) Metrics {
	return &MetricsImpl{
		data:    map[string][]MetricsDatum{},
		lock:    sync.Mutex{},
		timeout: timeout,
	}
}

// Cleanup will remove old values past the timeout within the specified bucket
func (metrics *MetricsImpl) Cleanup(bucket string) {
	metrics.lock.Lock()
	defer metrics.lock.Unlock()

	values, ok := metrics.data[bucket]
	if !ok {
		return
	}

	timedout := time.Now().Add(-metrics.timeout)
	newValues := []MetricsDatum{}

	for _, value := range values {
		if timedout.Before(value.When) {
			newValues = append(newValues, value)
		}
	}

	metrics.data[bucket] = newValues
}

// Get returns the value for the given bucket, if it is defined. If the bucket
// is not defined, this operation will return 0 (zero).
func (metrics *MetricsImpl) Get(bucket string) int {
	metrics.lock.Lock()
	defer metrics.lock.Unlock()

	total := 0

	if values, ok := metrics.data[bucket]; ok {
		for _, value := range values {
			total += value.Value
		}
	}

	return total
}

// Remove will remove the bucket, and all of it's associated data, if it is
// defined. If it is not defined, this operation does nothing.
func (metrics *MetricsImpl) Remove(bucket string) {
	metrics.lock.Lock()
	defer metrics.lock.Unlock()
	delete(metrics.data, bucket)
}

// Add or subtract the given value to the given bucket. If the bucket does not
// exist, it will be created. The time given will be the time associated with
// the given quantity change.
func (metrics *MetricsImpl) Update(bucket string, value int, when time.Time) {
	metrics.lock.Lock()
	defer metrics.lock.Unlock()

	datum := MetricsDatum{
		Value: value,
		When:  when,
	}

	_, ok := metrics.data[bucket]
	if ok {
		metrics.data[bucket] = append(metrics.data[bucket], datum)
	} else {
		metrics.data[bucket] = []MetricsDatum{datum}
	}
}
