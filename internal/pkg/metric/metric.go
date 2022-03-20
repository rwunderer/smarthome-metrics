package metric

type Metric map[string] float64

func NewMetric() Metric {
	return make(Metric)
}
