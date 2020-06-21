package probes

import "github.com/prometheus/client_golang/prometheus"

type Probe interface {
	// Returns a slice of collectors used by a probe
	Collectors() []prometheus.Collector

	// Runs probe itself
	Run() error
}
