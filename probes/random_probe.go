package probes

import (
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type randomProbe struct {
	counter *prometheus.CounterVec
	histo   *prometheus.HistogramVec
}

// Only const label required here. Prober runner will set them.
func NewRandomProbe(constLabels prometheus.Labels) Probe {
	rand.Seed(time.Now().UnixNano())
	cntVec := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   "randomprobe",
			Name:        "counto",
			Help:        "just a counter",
			ConstLabels: constLabels,
		},
		nil,
	)
	histVec := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace:   "randomprobe",
			Name:        "histo",
			Help:        "histogram for random values",
			ConstLabels: constLabels,
			Buckets:     []float64{0.1, 0.5, 0.9},
		},
		nil,
	)

	return &randomProbe{
		counter: cntVec,
		histo:   histVec,
	}
}

// Return all metric collectors as a slice, so caller can register them with prometheus
func (p *randomProbe) Collectors() []prometheus.Collector {
	return []prometheus.Collector{p.counter, p.histo}
}

func (collector *randomProbe) Run() error {
	collector.counter.With(nil).Inc()
	collector.histo.With(nil).Observe(rand.Float64())

	return nil
}
