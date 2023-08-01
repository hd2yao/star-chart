package github

import (
    "errors"
)

var (
    errNoMorePages  = errors.New("no more pages to get")
    ErrTooManyStars = errors.New("repo has too many stargazers, github won't allow us to list all stars")
)
