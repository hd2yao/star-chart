package github

import "errors"

// ErrRateLimit happens when we rate limit github API
var ErrRateLimit = errors.New("rate limited, please try again later")

// ErrGitHubAPI happens when github responds with something other than a 2xx
var ErrGitHubAPI = errors.New("failed to talk with github api")
