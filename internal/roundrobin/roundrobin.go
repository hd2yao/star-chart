// Package roundrobin provides round robin invalidation-aware load balancing of github tokens.
package roundrobin

import (
    "sync"

    "github.com/apex/log"
)

// RoundRobiner can pick a token from a list of tokens
type RoundRobiner interface {
    Pick() (*Token, error)
}

type noTokensRoundRobin struct{}

func (rr *noTokensRoundRobin) Pick() (*Token, error) {
    return nil, nil
}

// Token is a github token
type Token struct {
    token string
    valid bool
    lock  sync.RWMutex
}

// NewToken from its string representation
func NewToken(token string) *Token {
    return &Token{
        token: token,
        valid: true,
    }
}

// String returns the last 3 chars for the token
func (t *Token) String() string {
    return t.token[len(t.token)-3:]
}

// Key returns the actual token
func (t *Token) Key() string {
    return t.token
}

// OK returns true if the token is valid
func (t *Token) OK() bool {
    t.lock.RLock()
    defer t.lock.RUnlock()
    return t.valid
}

// Invalidate invalidates the token
func (t *Token) Invalidate() {
    log.Warnf("invalidated token '...%s'", t)
    t.lock.Lock()
    defer t.lock.Unlock()
    t.valid = false
}
