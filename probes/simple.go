package probes

import (
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Probe interface {
	Collectors() []prometheus.Collector
	Run()
}

type randomProbe struct {
	counter *prometheus.CounterVec
	histo   *prometheus.HistogramVec
}

func NewRandomProbe() Probe {
	rand.Seed(time.Now().UnixNano())
	cntVec := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   "example",
			Name:        "counto",
			Help:        "just a counter",
			ConstLabels: nil,
		},
		nil,
	)
	histVec := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace:   "example",
			Name:        "histo",
			Help:        "histogram for random values",
			ConstLabels: nil,
			Buckets:     []float64{0.1, 0.5, 0.9},
		},
		nil,
	)

	return &randomProbe{
		counter: cntVec,
		histo:   histVec,
	}
}

func (p *randomProbe) Collectors() []prometheus.Collector {
	return []prometheus.Collector{p.counter, p.histo}
}

func (collector *randomProbe) Run() {
	collector.counter.With(nil).Inc()
	collector.histo.With(nil).Observe(rand.Float64())
}
