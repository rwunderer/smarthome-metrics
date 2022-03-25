package metric

import (
	"sync"
	"time"
)

type Metric struct {
	Value float64
	Time  time.Time
}

type Metrics struct {
	metrics map[string]Metric
	lock    sync.RWMutex
}

func NewMetrics() *Metrics {
	m := new(Metrics)
	m.metrics = make(map[string]Metric)

	return m
}

func (m *Metrics) Set(key string, value float64) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.metrics[key] = Metric{
		Value: value,
		Time:  time.Now(),
	}
}

func (m *Metrics) Get(key string) Metric {
	m.lock.RLock()
	defer m.lock.RUnlock()

	rv := m.metrics[key]
	return rv
}

func (m *Metrics) Iterate(f func(string, Metric)) {
	for k, v := range m.metrics {
		f(k, v)
	}
}
