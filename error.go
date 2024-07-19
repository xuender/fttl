package fttl

import "errors"

var (
	ErrIsDir    = errors.New("is a directory")
	ErrNotFound = errors.New("not found")
)
