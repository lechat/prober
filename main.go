package main

import (
	"fmt"
	"time"

	"net/http"

	"github.com/lechat/prober/probes"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	sprobe := probes.NewRandomProbe()
	for _, coll := range sprobe.Collectors() {
		prometheus.MustRegister(coll)
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Done!")
				return
			case t := <-ticker.C:
				fmt.Println("Current time: ", t)
				sprobe.Run()
			}
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
