package datastructures

import "errors"

var (
	ErrNotFound   = errors.New("value is not found")
	ErrNoElements = errors.New("is empty")
)
