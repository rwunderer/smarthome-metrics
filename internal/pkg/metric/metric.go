package metric

import (
	"sync"
	"time"
)

type Metric struct {
	value float64
	time  time.Time
}

type Metrics struct {
	metrics map[string]Metric
	lock    sync.RWMutex
}

func NewMetrics() Metrics {
	return make(Metrics)
}

func (Metrics *m) Set(key string, value float64) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.metrics[key] = Metric{
		value: value,
		time:  time.Now(),
	}
}

func (Metrics *m) Get(key string) Metric {
	m.lock.RLock()
	defer m.lock.RUnlock()

	rv := m.metrics[key]
	return rv
}
