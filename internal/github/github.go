package github

import (
    "errors"

    "github.com/prometheus/client_golang/prometheus"
)

// ErrRateLimit happens when we rate limit github API
var ErrRateLimit = errors.New("rate limited, please try again later")

// ErrGitHubAPI happens when github responds with something other than a 2xx
var ErrGitHubAPI = errors.New("failed to talk with github api")

var rateLimits = prometheus.NewCounter(prometheus.CounterOpts{
    Namespace: "star-chart",
    Subsystem: "github",
    Name:      "rate_limit_hits_total",
})

var effectiveEtags = prometheus.NewCounter(prometheus.CounterOpts{
    Namespace: "star-chart",
    Subsystem: "github",
    Name:      "effective_etag_users_total",
})

var tokensCount = prometheus.NewGauge(prometheus.GaugeOpts{
    Namespace: "star-chart",
    Subsystem: "github",
    Name:      "available_tokens",
})

var invalidatedTokens = prometheus.NewGauge(prometheus.GaugeOpts{
    Namespace: "star-chart",
    Subsystem: "github",
    Name:      "invalidated_tokens_total",
})

var rateLimiters = prometheus.NewGauge(prometheus.GaugeOpts{
    Namespace: "star-chart",
    Subsystem: "github",
    Name:      "rate_limit_remaining",
})

func init() {
    prometheus.MustRegister(rateLimits, effectiveEtags, invalidatedTokens, tokensCount, rateLimiters)
}
