package cache

import (
    "github.com/prometheus/client_golang/prometheus"
)

var cacheGets = prometheus.NewCounter(
    prometheus.CounterOpts{
        Namespace: "star-chart",
        Subsystem: "cache",
        Name:      "gets_total",
        Help:      "Total number of successful cache gets",
    },
)

var cachePuts = prometheus.NewCounter(
    prometheus.CounterOpts{
        Namespace: "star-chart",
        Subsystem: "cache",
        Name:      "puts_total",
        Help:      "Total number of successful cache puts",
    },
)

var cacheDeletes = prometheus.NewCounter(
    prometheus.CounterOpts{
        Namespace: "star-chart",
        Subsystem: "cache",
        Name:      "deletes_total",
        Help:      "Total number of successful cache deletes",
    },
)

func init() {
    prometheus.MustRegister(cacheGets, cachePuts)
}
