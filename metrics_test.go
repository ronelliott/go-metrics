package metrics_test

import (
	"testing"
	"time"

	"github.com/ronelliott/go-metrics"
	"github.com/stretchr/testify/assert"
)

func TestMetrics_Cleanup(t *testing.T) {
	m := metrics.New(time.Second * 2)

	m.Update("bucket", 2, time.Now())
	assert.Equal(t, 2, m.Get("bucket"))

	time.Sleep(time.Second * 3)
	m.Update("bucket", 3, time.Now())
	assert.Equal(t, 5, m.Get("bucket"))

	m.Cleanup("bucket")
	assert.Equal(t, 3, m.Get("bucket"))

	assert.Equal(t, 0, m.Get("bucket2"))
	m.Cleanup("bucket2")
	assert.Equal(t, 0, m.Get("bucket2"))
}

func TestMetrics_Get(t *testing.T) {
	m := metrics.New(time.Second * 2)
	m.Update("bucket", 20, time.Now())
	assert.Equal(t, 20, m.Get("bucket"))
}

func TestMetrics_Remove(t *testing.T) {
	m := metrics.New(time.Second * 2)
	m.Update("bucket", 14, time.Now())
	assert.Equal(t, 14, m.Get("bucket"))
	m.Remove("bucket")
	assert.Equal(t, 0, m.Get("bucket"))
}

func TestMetrics_Update(t *testing.T) {
	m := metrics.New(time.Second * 2)
	m.Update("bucket", 83, time.Now())
	assert.Equal(t, 83, m.Get("bucket"))
}
