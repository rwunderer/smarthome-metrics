package metric

import (
	"sync"
	"time"
)

type GraphiteMap map[string]float64

type Metric struct {
	Value float64
	Time  time.Time
}

type Metrics struct {
	prefix  string
	metrics map[string]Metric
	lock    sync.RWMutex
}

func NewMetrics(prefix string) *Metrics {
	m := new(Metrics)
	m.prefix = prefix
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

func (m *Metrics) Iterate(f func(string, string, Metric)) {
	for k, v := range m.metrics {
		f(m.prefix, k, v)
	}
}

func (m *Metrics) GetGraphiteMap() GraphiteMap {
	r := make(GraphiteMap)

	for k, v := range m.metrics {
		r[m.prefix+"."+k] = v.Value
	}

	return r
}
